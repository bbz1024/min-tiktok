package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"min-tiktok/services/auths/auths"
	"min-tiktok/services/auths/internal/config"
	"min-tiktok/services/auths/internal/server"
	"min-tiktok/services/auths/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/auths.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		auths.RegisterAuthsServer(grpcServer, server.NewAuthsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
