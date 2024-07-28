package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"math"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/user/internal/config"
	"min-tiktok/services/user/internal/mq"
	"min-tiktok/services/user/internal/server"
	"min-tiktok/services/user/internal/svc"
	"min-tiktok/services/user/user"
	"time"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

const (
	// 线程池大小
	workPoolSize = 10
	// 7天
	triggerTime = time.Hour * 24 * 7
	// 秒，分、时、日、月、周
	period = "*/10 * * * *"
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}
	if err := mq.InitDeed(ctx); err != nil {
		panic(err)
	}
	// --------------------	start cron  --------------------
	schedule := cron.New()
	deed := NewDeed(ctx)
	if err := schedule.AddJob(period, deed); err != nil {
		panic(err)
	}
	schedule.Start()
	defer schedule.Stop()
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

type Deed struct {
	svc          *svc.ServiceContext
	workPoolSize int
}

func NewDeed(svc *svc.ServiceContext) *Deed {
	return &Deed{
		svc:          svc,
		workPoolSize: workPoolSize,
	}
}

func (d *Deed) Run() {
	allUserID, err := d.svc.UserModel.QueryAllUserID(context.Background())
	if err != nil {
		logx.Errorf("query all user id error: %v", err)
		return
	}
	workPool := threading.NewTaskRunner(d.workPoolSize)
	for _, userID := range allUserID {
		userid := userID
		workPool.Schedule(func() {
			key := fmt.Sprintf(keys.UserDeedKey, userid)
			// check exist key
			exists, err := d.svc.Rdb.Exists(key)
			if err != nil {
				logx.Errorf("check exist key error: %v", err)
				return
			}
			if !exists {
				return
			}
			// check meet the condition
			pairs, err := d.svc.Rdb.ZrangebyscoreWithScoresAndLimit(key, 0, math.MaxInt, 0, 1)
			if err != nil && !errors.Is(err, redis.Nil) {
				logx.Errorf("zrangebyscore error: %v", err)
				return
			}
			if len(pairs) == 0 {
				logx.Infow("user  not meet the condition", logx.Field("userid", userid))
				return
			}
			lastTime := time.Unix(pairs[0].Score, 0)
			if time.Since(lastTime) < triggerTime {
				logx.Infow("user  not meet the condition", logx.Field("userid", userid))
				return
			}

			// get all videoId
			videoIds, err := d.svc.Rdb.Zrange(key, 0, -1)
			if err != nil && !errors.Is(err, redis.Nil) {
				logx.Errorf("zrange error: %v", err)
				return
			}
			if err := mq.GetInstance().Product(mq.DeedReq{
				VideoIds: videoIds,
				UserID:   userid,
			}); err != nil {
				logx.Errorf("product error: %v", err)
				return
			}
			// clear deed history
			if _, err = d.svc.Rdb.Del(key); err != nil {
				logx.Errorf("del error: %v", err)
				return
			}
			logx.Infow("user deed success", logx.Field("userid", userid))
		})
	}
	workPool.Wait()

}
