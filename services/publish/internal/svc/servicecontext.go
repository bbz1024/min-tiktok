package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/models/video"
	"min-tiktok/services/feed/feedclient"
	"min-tiktok/services/publish/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel video.VideoModel
	Rdb        *redis.Redis
	FeedRpc    feedclient.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	rdb, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		VideoModel: video.NewVideoModel(mysqlConn),
		FeedRpc:    feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc)),
		Rdb:        rdb,
	}
}