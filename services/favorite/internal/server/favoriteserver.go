// Code generated by goctl. DO NOT EDIT.
// Source: favorite.proto

package server

import (
	"context"

	"min-tiktok/services/favorite/favorite"
	"min-tiktok/services/favorite/internal/logic"
	"min-tiktok/services/favorite/internal/svc"
)

type FavoriteServer struct {
	svcCtx *svc.ServiceContext
	favorite.UnimplementedFavoriteServer
}

func NewFavoriteServer(svcCtx *svc.ServiceContext) *FavoriteServer {
	return &FavoriteServer{
		svcCtx: svcCtx,
	}
}

func (s *FavoriteServer) FavoriteAction(ctx context.Context, in *favorite.FavoriteRequest) (*favorite.FavoriteResponse, error) {
	l := logic.NewFavoriteActionLogic(ctx, s.svcCtx)
	return l.FavoriteAction(in)
}

func (s *FavoriteServer) FavoriteList(ctx context.Context, in *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	l := logic.NewFavoriteListLogic(ctx, s.svcCtx)
	return l.FavoriteList(in)
}

func (s *FavoriteServer) IsFavorite(ctx context.Context, in *favorite.IsFavoriteRequest) (*favorite.IsFavoriteResponse, error) {
	l := logic.NewIsFavoriteLogic(ctx, s.svcCtx)
	return l.IsFavorite(in)
}

func (s *FavoriteServer) CountFavorite(ctx context.Context, in *favorite.CountFavoriteRequest) (*favorite.CountFavoriteResponse, error) {
	l := logic.NewCountFavoriteLogic(ctx, s.svcCtx)
	return l.CountFavorite(in)
}

func (s *FavoriteServer) CountUserFavorite(ctx context.Context, in *favorite.CountUserFavoriteRequest) (*favorite.CountUserFavoriteResponse, error) {
	l := logic.NewCountUserFavoriteLogic(ctx, s.svcCtx)
	return l.CountUserFavorite(in)
}

func (s *FavoriteServer) CountUserTotalFavorited(ctx context.Context, in *favorite.CountUserTotalFavoritedRequest) (*favorite.CountUserTotalFavoritedResponse, error) {
	l := logic.NewCountUserTotalFavoritedLogic(ctx, s.svcCtx)
	return l.CountUserTotalFavorited(in)
}
