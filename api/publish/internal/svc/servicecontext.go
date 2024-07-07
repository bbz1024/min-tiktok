package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/publish/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/auths/authsclient"
	"min-tiktok/services/publish/publishclient"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	AuthsRpc       authsclient.Auths
	PublishRpc     publishclient.Publish
	UserRpc        userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.WrapperAuthMiddleware(c.AuthsRpc),
		PublishRpc:     publishclient.NewPublish(zrpc.MustNewClient(c.PublishRpc)),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
