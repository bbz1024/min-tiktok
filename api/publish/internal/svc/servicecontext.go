package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/publish/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/auths/authsclient"
	"min-tiktok/services/publish/publishclient"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	WithMiddleware rest.Middleware
	AuthsRpc       authsclient.Auths
	PublishRpc     publishclient.Publish
}

func NewServiceContext(c config.Config) *ServiceContext {
	zrpc.SetClientSlowThreshold(time.Millisecond * time.Duration(c.Timeout))
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.WrapperAuthMiddleware(c.AuthsRpc),
		WithMiddleware: middleware.WithMiddleware,
		PublishRpc:     publishclient.NewPublish(zrpc.MustNewClient(c.PublishRpc)),
	}
}
