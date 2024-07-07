package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/models/user"
	"min-tiktok/models/video"
	"min-tiktok/services/feed/internal/config"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config     config.Config
	Rdb        *redis.Redis
	UserModel  user.UsersModel
	VideoModel video.VideoModel
	UserRpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	rdb, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		UserModel:  user.NewUsersModel(mysqlConn),
		VideoModel: video.NewVideoModel(mysqlConn),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Rdb:        rdb,
	}
}
