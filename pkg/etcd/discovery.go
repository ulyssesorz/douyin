package etcd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/ulyssesorz/douyin/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	defaultWeight = 10
)

type etcdResolver struct {
	etcdClient *clientv3.Client
}

func NewEtcdResolver(endpoints []string) (discovery.Resolver, error) {
	return NewEtcdResolverWithAuth(endpoints, "", "")
}

func NewEtcdResolverWithAuth(endpoints []string, username, password string) (discovery.Resolver, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		Username:  username,
		Password:  password,
	})
	if err != nil {
		return nil, err
	}
	return &etcdResolver{
		etcdClient: etcdClient,
	}, nil
}

func (e *etcdResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

// 从etcd解析出服务
func (e *etcdResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	logger := zap.InitLogger()
	prefix := serviceKeyPrefix(desc)
	resp, err := e.etcdClient.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return discovery.Result{}, err
	}
	var (
		info instanceInfo
		eps  []discovery.Instance
	)
	// etcd本质上是一个分布式kv存储，服务发现就是找出全部相关实例
	for _, kv := range resp.Kvs {
		err := json.Unmarshal(kv.Value, &info)
		if err != nil {
			logger.Warnf("fail to unmarshal with err: %v, ignore key: %v", err, string(kv.Key))
			continue
		}
		weight := info.Weight
		if weight <= 0 {
			weight = defaultWeight
		}
		eps = append(eps, discovery.NewInstance(info.Network, info.Address, weight, info.Tags))
	}
	if len(eps) == 0 {
		return discovery.Result{}, fmt.Errorf("no instance remains for %v", desc)
	}
	return discovery.Result{
		Cacheable: true,
		CacheKey:  desc,
		Instances: eps,
	}, nil
}

func (e *etcdResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

func (e *etcdResolver) Name() string {
	return "etcd"
}
