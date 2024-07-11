package logic

import (
	"context"

	"min-tiktok/services/favorite/favorite"
	"min-tiktok/services/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountUserFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountUserFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountUserFavoriteLogic {
	return &CountUserFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountUserFavoriteLogic) CountUserFavorite(in *favorite.CountUserFavoriteRequest) (*favorite.CountUserFavoriteResponse, error) {
	// todo: add your logic here and delete this line

	return &favorite.CountUserFavoriteResponse{}, nil
}
