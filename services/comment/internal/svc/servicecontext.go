package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/models/comment"
	"min-tiktok/services/comment/internal/config"
	"min-tiktok/services/feedback/feedbackclient"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config       config.Config
	Rdb          *redis.Redis
	CommentModel comment.CommentModel
	UserRpc      userclient.User
	FeedbackRpc  feedbackclient.Feedback
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb, err := redis.NewRedis(c.CacheConf[0].RedisConf)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		Rdb:          rdb,
		CommentModel: comment.NewCommentModel(sqlx.NewMysql(c.MySQL.DataSource)),
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		FeedbackRpc:  feedbackclient.NewFeedback(zrpc.MustNewClient(c.FeedBackRpc)),
	}
}
