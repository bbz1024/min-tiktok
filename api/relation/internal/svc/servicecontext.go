package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/relation/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/relation/relationclient"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	WithMiddleware rest.Middleware
	RelationRpc    relationclient.Relation
	UserRpc        userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.WrapperAuthMiddleware(c.AuthsRpc),
		WithMiddleware: middleware.WithMiddleware,
		RelationRpc:    relationclient.NewRelation(zrpc.MustNewClient(c.RelationRpc)),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
