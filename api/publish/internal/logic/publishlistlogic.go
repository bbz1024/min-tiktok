package logic

import (
	"context"
	"min-tiktok/api/publish/internal/svc"
	"min-tiktok/api/publish/internal/types"
	"min-tiktok/services/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListReq) (resp *types.PublishListResp, err error) {
	res, err := l.svcCtx.PublishRpc.ListVideo(l.ctx, &publish.ListVideoReq{
		ActorId: req.ActorId,
		UserId:  req.UserId,
	})
	resp = new(types.PublishListResp)
	resp.StatusMsg = res.StatusMsg
	resp.StatusCode = res.StatusCode
	if err != nil {
		l.Errorw("call rpc PublishRpc.PublishList error ", logx.Field("err", err))
		return
	}
	var videoList = make([]types.Video, 0, len(res.VideoList))
	for _, v := range res.VideoList {
		videoList = append(videoList, types.Video{
			Id: v.Id,
			Author: &types.User{
				ID:              v.Author.Id,
				Name:            v.Author.Name,
				FollowCount:     v.Author.FollowCount,
				FollowerCount:   v.Author.FollowerCount,
				IsFollow:        v.Author.IsFollow,
				Avatar:          v.Author.Avatar,
				BackgroundImage: v.Author.BackgroundImage,
				TotalFavorited:  v.Author.TotalFavorited,
				WorkCount:       v.Author.WorkCount,
				Signature:       v.Author.Signature,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	resp.VideoList = videoList
	return
}
