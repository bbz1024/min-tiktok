package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/feed/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/feed/feedclient"
)

type ServiceContext struct {
	Config config.Config
	// middleware
	AuthMiddleware rest.Middleware
	// rpc
	FeedRpc feedclient.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.WrapperAuthMiddleware(c.AuthsRpc),
		FeedRpc:        feedclient.NewFeed(zrpc.MustNewClient(c.FeedRpc)),
	}
}
