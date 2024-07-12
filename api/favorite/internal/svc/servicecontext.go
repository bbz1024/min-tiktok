package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/favorite/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/favorite/favoriteclient"
	"min-tiktok/services/user/userclient"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	WithMiddleware rest.Middleware
	FavoriteRpc    favoriteclient.Favorite
	UserRpc        userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.WrapperAuthMiddleware(c.AuthsRpc),
		WithMiddleware: middleware.WithMiddleware,
		FavoriteRpc:    favoriteclient.NewFavorite(zrpc.MustNewClient(c.FavoriteRpc)),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
