package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/models/video"
	"min-tiktok/models/videoInfo"
	"min-tiktok/services/feed/feedclient"
	"min-tiktok/services/publish/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	VideoModel     video.VideoModel
	VideoInfoModel videoInfo.VideoinfoModel
	Rdb            *redis.Redis
	FeedRpc        feedclient.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {

	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	mysqlConn2 := sqlx.NewMysql(c.MySQL.DataSource)
	videoInfoModel := videoInfo.NewVideoinfoModel(mysqlConn)
	videoModel := video.NewVideoModel(mysqlConn2, c.CacheConf)
	rdb, err := redis.NewRedis(c.CacheConf[0].RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		VideoModel:     videoModel,
		FeedRpc:        feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc)),
		Rdb:            rdb,
		VideoInfoModel: videoInfoModel,
	}
}
