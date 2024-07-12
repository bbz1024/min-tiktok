package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/models/user"
	"min-tiktok/services/favorite/internal/config"
	"min-tiktok/services/feed/feedclient"
	"min-tiktok/services/feedback/feedbackclient"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   user.UsersModel
	Rdb         *redis.Redis
	FeedRpc     feedclient.Feed
	FeedBackRpc feedbackclient.Feedback
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
		FeedRpc:     feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc)),
		FeedBackRpc: feedbackclient.NewFeedback(zrpc.MustNewClient(c.FeedBackRpc)),
	}
}
