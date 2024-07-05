package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"min-tiktok/models/user"
	"min-tiktok/services/user/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UsersModel
	Rdb       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUsersModel(mysqlConn),
		Rdb: redis.New(c.RedisConf.Host, func(r *redis.Redis) {
			r.Type = c.RedisConf.Type
			r.Pass = c.RedisConf.Pass
		}),
	}
}
