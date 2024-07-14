package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/publish/internal/config"
	"min-tiktok/services/publish/internal/mq"
	"min-tiktok/services/publish/internal/server"
	"min-tiktok/services/publish/internal/svc"
	"min-tiktok/services/publish/publish"
	"os"
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
	//
	// max-size 50m
	s.AddOptions(grpc.MaxRecvMsgSize((1 << 20) * 50))
	//
	// register center
	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}

	// -------------------- init --------------------
	if err := mq.InitExtractVideo(ctx, 3, 3); err != nil {
		panic(err)
	}
	if err := mq.InitChatVideo(ctx, 3, 3); err != nil {
		panic(err)
	}

	defer s.Stop()
	//genVideoInfo(ctx)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
func genVideoInfo(ctx *svc.ServiceContext) {
	//get all video
	ids, err := ctx.VideoModel.GetVideoIds(context.Background())
	if err != nil {
		return
	}
	userID := 1
	for _, id := range ids {
		key := fmt.Sprintf(keys.UserInfoKey, userID)
		if _, err := ctx.Rdb.HincrbyCtx(context.Background(), key, keys.WorkCount, 1); err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("incr user work count ", logx.Field("err", err))
		}

		// 4. add video id to user video list
		key = fmt.Sprintf(keys.UserWorkKey, userID)
		if _, err := ctx.Rdb.SaddCtx(context.Background(), key, id); err != nil {
			logx.Errorw("add video id to user video list ", logx.Field("err", err))
		}
		videoInfoKey := fmt.Sprintf(keys.VideoInfoKey, id)
		if err := ctx.Rdb.HsetCtx(context.Background(), videoInfoKey, keys.VideoAuthorID, fmt.Sprintf("%d", userID)); err != nil {
			logx.Errorw("set video author id ", logx.Field("err", err))
		}
		//if err := mq.GetChatVideo().Product(mq.ChatVideoReq{
		//	VideoID: id}); err != nil {
		//	logx.Error(err)
		//}
	}
	os.Exit(0)
}
