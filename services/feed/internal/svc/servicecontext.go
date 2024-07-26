package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/models/video"
	"min-tiktok/services/feed/internal/config"
	"min-tiktok/services/feedback/feedbackclient"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config      config.Config
	Rdb         *redis.Redis
	VideoModel  video.VideoModel
	UserRpc     userclient.User
	FeedBackRpc feedbackclient.Feedback
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	rdb, err := redis.NewRedis(c.CacheConf[0].RedisConf)

	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		VideoModel:  video.NewVideoModel(mysqlConn, c.CacheConf),
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Rdb:         rdb,
		FeedBackRpc: feedbackclient.NewFeedback(zrpc.MustNewClient(c.FeedBackRpc)),
	}
}
