package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	mq2 "min-tiktok/common/mq"
	"min-tiktok/services/publish/internal/config"
	"min-tiktok/services/publish/internal/mq"
	"min-tiktok/services/publish/internal/server"
	"min-tiktok/services/publish/internal/svc"
	"min-tiktok/services/publish/publish"
)

var configFile = flag.String("f", "etc/publish.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		publish.RegisterPublishServer(grpcServer, server.NewPublishServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// max-size
	// 50m
	s.AddOptions(grpc.MaxRecvMsgSize((1 << 20) * 50))
	//
	// register center
	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}

	// -------------------- init --------------------
	if err := mq.InitVideoSummery(ctx); err != nil {
		panic(err)
	}

	if err := mq2.InitGorse(&c.Gorse, &c.RabbitMQ, &c.MySQL); err != nil {
		panic(err)
	}

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
