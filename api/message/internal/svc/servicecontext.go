package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/message/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/message/messageclient"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config config.Config
	// middleware
	AuthMiddleware rest.Middleware
	WithMiddleware rest.Middleware
	MessageRpc     messageclient.Message
	UserRpc        userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.WrapperAuthMiddleware(c.AuthsRpc),
		WithMiddleware: middleware.WithMiddleware,
		MessageRpc:     messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
