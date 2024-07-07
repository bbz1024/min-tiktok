package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"min-tiktok/models/user"
	"min-tiktok/services/auths/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UsersModel
	Rdb       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	rdb, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUsersModel(mysqlConn),
		Rdb:       rdb,
	}
}
