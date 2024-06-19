package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/registry"

	"github.com/ulyssesorz/douyin/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	ttlKey     = "KITEX_ETCD_REGISTRY_LEASE_TTL"
	defaultTTL = 60
)

type registerMeta struct {
	leaseID clientv3.LeaseID
	ctx     context.Context
	cancel  context.CancelFunc
}

type etcdRegistry struct {
	etcdClient *clientv3.Client
	leaseTTL   int64
	meta       *registerMeta
}

func NewEtcdRegistry(endpoints []string) (registry.Registry, error) {
	return NewEtcdRegistryWithAuth(endpoints, "", "")
}

func NewEtcdRegistryWithAuth(endpoints []string, username, password string) (registry.Registry, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		Username:  username,
		Password:  password,
	})
	if err != nil {
		return nil, err
	}
	return &etcdRegistry{
		etcdClient: etcdClient,
		leaseTTL:   getTTL(),
	}, nil
}

// 授权租约
func (e *etcdRegistry) grantLease() (clientv3.LeaseID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := e.etcdClient.Grant(ctx, e.leaseTTL)
	if err != nil {
		return clientv3.NoLease, err
	}
	return resp.ID, nil
}

// 定期更新租约，使其保持有效
func (e *etcdRegistry) keepalive(meta *registerMeta) error {
	logger := zap.InitLogger()
	keepAlive, err := e.etcdClient.KeepAlive(meta.ctx, meta.leaseID)
	if err != nil {
		return err
	}
	go func() {
		logger.Infof("start keepalive lease %x for etcd registry", meta.leaseID)
		// 租约撤销（keepalive通道关闭）或服务结束（done通过发消息）结束租约
		for range keepAlive {
			select {
			case <-meta.ctx.Done():
				break
			default:
			}
		}
		logger.Infof("stop keepalive lease %x for etcd registry", meta.leaseID)
	}()
	return nil
}

// 检验注册信息
func validateRegistryInfo(info *registry.Info) error {
	if info.ServiceName == "" {
		return fmt.Errorf("missing service name in Register")
	}
	if info.Addr == nil {
		return fmt.Errorf("missing addr in Register")
	}
	return nil
}

func getTTL() int64 {
	var ttl int64 = defaultTTL
	if str, ok := os.LookupEnv(ttlKey); ok {
		if t, err := strconv.Atoi(str); err == nil {
			ttl = int64(t)
		}
	}
	return ttl
}

func (e *etcdRegistry) register(info *registry.Info, leaseID clientv3.LeaseID) error {
	val, err := json.Marshal(&instanceInfo{
		Network: info.Addr.Network(),
		Address: info.Addr.String(),
		Weight:  info.Weight,
		Tags:    info.Tags,
	})
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = e.etcdClient.Put(ctx, serviceKey(info.ServiceName, info.Addr.String()), string(val), clientv3.WithLease(leaseID))
	return err
}

// 注册
func (e *etcdRegistry) Register(info *registry.Info) error {
	if err := validateRegistryInfo(info); err != nil {
		return err
	}
	leaseID, err := e.grantLease()
	if err != nil {
		return err
	}

	if err := e.register(info, leaseID); err != nil {
		return err
	}
	meta := registerMeta{
		leaseID: leaseID,
	}
	meta.ctx, meta.cancel = context.WithCancel(context.Background())
	if err := e.keepalive(&meta); err != nil {
		return err
	}
	e.meta = &meta
	return nil
}

func (e *etcdRegistry) deregister(info *registry.Info) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err := e.etcdClient.Delete(ctx, serviceKey(info.ServiceName, info.Addr.String()))
	return err
}

// 取消注册
func (e *etcdRegistry) Deregister(info *registry.Info) error {
	if info.ServiceName == "" {
		return fmt.Errorf("missing service name in Deregister")
	}
	if err := e.deregister(info); err != nil {
		return err
	}
	e.meta.cancel()
	return nil
}
