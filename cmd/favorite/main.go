package main

import (
	"fmt"
	"net"

	"github.com/ulyssesorz/douyin/cmd/favorite/service"

	"github.com/ulyssesorz/douyin/kitex/kitex_gen/favorite/favoriteservice"
	"github.com/ulyssesorz/douyin/pkg/etcd"
	"github.com/ulyssesorz/douyin/pkg/middleware"
	"github.com/ulyssesorz/douyin/pkg/viper"
	"github.com/ulyssesorz/douyin/pkg/zap"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

var (
	config      = viper.Init("favorite")
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
	defer service.FavoriteMq.Destroy()
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		logger.Fatalln(err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		logger.Fatalln(err.Error())
	}

	// 初始化etcd
	s := favoriteservice.NewServer(new(service.FavoriteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithRegistry(r),
		//server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		// server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)

	if err := s.Run(); err != nil {
		logger.Fatalf("%v stopped with error: %v", serviceName, err.Error())
	}
}
