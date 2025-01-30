package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/auths/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/auths/authsclient"
)

type ServiceContext struct {
	Config         config.Config
	AuthsRpc       authsclient.Auths
	WithMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthsRpc:       authsclient.NewAuths(zrpc.MustNewClient(c.AuthsRpc)),
		WithMiddleware: middleware.WithMiddleware,
	}
}
