package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	client "min-tiktok/common/util/gorse"
	"min-tiktok/services/feedback/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	Recommend *client.GorseClient
	Rdb       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb, err := redis.NewRedis(c.CacheConf[0].RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		Recommend: client.NewGorseClient(c.Gorse.GorseAddr, c.Gorse.GorseApikey),
		Rdb:       rdb,
	}
}
