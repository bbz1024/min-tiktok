// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "min-tiktok/api/auths/internal/handler/auth"
	"min-tiktok/api/auths/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.WithMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: auth.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/register",
					Handler: auth.RegisterHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/user"),
	)
}
