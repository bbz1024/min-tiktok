package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/services/favorite/internal/config"
	"min-tiktok/services/feed/feedclient"
	"min-tiktok/services/feedback/feedbackclient"
)

type ServiceContext struct {
	Config      config.Config
	Rdb         *redis.Redis
	FeedRpc     feedclient.Feed
	FeedBackRpc feedbackclient.Feedback
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb, err := redis.NewRedis(c.CacheConf[0].RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		Rdb:         rdb,
		FeedRpc:     feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc)),
		FeedBackRpc: feedbackclient.NewFeedback(zrpc.MustNewClient(c.FeedBackRpc)),
	}
}
