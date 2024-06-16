package main

import (
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/ulyssesorz/douyin/cmd/user/service"
	"github.com/ulyssesorz/douyin/kitex/kitex_gen/user/userservice"
	"github.com/ulyssesorz/douyin/pkg/etcd"
	"github.com/ulyssesorz/douyin/pkg/middleware"
	"github.com/ulyssesorz/douyin/pkg/viper"
	"github.com/ulyssesorz/douyin/pkg/zap"
)

var (
	config      = viper.Init("user")
	serviceName = config.Viper.GetString("server.name")
	serviceAddr = fmt.Sprintf("%s:%d", config.Viper.GetString("server.host"), config.Viper.GetInt("server.port"))
	etcdAddr    = fmt.Sprintf("%s:%d", config.Viper.GetString("etcd.host"), config.Viper.GetInt("etcd.port"))
	signingKey  = config.Viper.GetString("JWT.signingKey")
	logger      = zap.InitLogger()
)

func init() {
	service.Init(signingKey)
}

func main() {
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		logger.Fatalln(err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Fatalln(err.Error())
	}

	// 初始化etcd
	s := userservice.NewServer(new(service.UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithRegistry(r),
		//server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		//server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)

	if err := s.Run(); err != nil {
		logger.Fatalf("%v stopped with error: %v", serviceName, err.Error())
	}
}
