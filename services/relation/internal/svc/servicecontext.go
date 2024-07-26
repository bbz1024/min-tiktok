package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/services/relation/internal/config"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	Rdb     *redis.Redis
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb, err := redis.NewRedis(c.CacheConf[0].RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:  c,
		Rdb:     rdb,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
