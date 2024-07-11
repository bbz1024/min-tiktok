package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhenghaoz/gorse/client"
	"min-tiktok/models/user"
	"min-tiktok/services/auths/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   user.UsersModel
	Rdb         *redis.Redis
	GorseClient *client.GorseClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	rdb, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		UserModel:   user.NewUsersModel(mysqlConn),
		Rdb:         rdb,
		GorseClient: client.NewGorseClient(c.Gorse.GorseAddr, c.Gorse.GorseApikey),
	}
}
