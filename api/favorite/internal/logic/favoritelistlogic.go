package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/favorite/favorite"

	"min-tiktok/api/favorite/internal/svc"
	"min-tiktok/api/favorite/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.ListRequest) (resp *types.ListResponse, err error) {
	res, err := l.svcCtx.FavoriteRpc.FavoriteList(l.ctx, &favorite.FavoriteListRequest{
		UserId:  req.UserID,
		ActorId: req.ActorID,
	})
	resp = new(types.ListResponse)

	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc FavoriteRpc.FavoriteList error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = uint32(res.StatusCode)
		resp.StatusMsg = res.StatusMsg
		return
	}
	for _, v := range res.VideoList {
		resp.VideoList = append(resp.VideoList, types.Video{
			Id: v.Id,
			Author: &types.User{
				ID:              v.Author.Id,
				Name:            v.Author.Name,
				FollowCount:     v.Author.FollowCount,
				FollowerCount:   v.Author.FollowerCount,
				IsFollow:        v.Author.IsFollow,
				Avatar:          v.Author.Avatar,
				Signature:       v.Author.Signature,
				TotalFavorited:  v.Author.TotalFavorited,
				WorkCount:       v.Author.WorkCount,
				FavoriteCount:   v.Author.FavoriteCount,
				BackgroundImage: v.Author.BackgroundImage,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	return
}
