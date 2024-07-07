package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/feed/feed"
	"min-tiktok/services/publish/internal/svc"
	"min-tiktok/services/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoLogic {
	return &ListVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListVideo 获取用户发布列表
func (l *ListVideoLogic) ListVideo(in *publish.ListVideoReq) (*publish.ListVideoResp, error) {
	res, err := l.svcCtx.FeedRpc.ListVideosByUserID(l.ctx, &feed.ListVideosByUserIDRequest{
		UserId:  in.UserId,
		ActorId: in.ActorId,
	})
	if err != nil {
		logx.Errorw("query feed error", logx.Field("err", err))
		return &publish.ListVideoResp{
			StatusCode: code.ServerError,
			StatusMsg:  code.ServerErrorMsg,
		}, err
	}
	var videos = make([]*publish.Video, 0, len(res.VideoList))
	for _, v := range res.VideoList {
		videos = append(videos, &publish.Video{
			Id:            v.Id,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
			Author: &publish.UserInfo{
				Id:              v.Author.Id,
				Name:            v.Author.Name,
				FollowCount:     v.Author.FollowCount,
				FollowerCount:   v.Author.FollowerCount,
				IsFollow:        v.Author.IsFollow,
				Avatar:          v.Author.Avatar,
				BackgroundImage: v.Author.BackgroundImage,
				Signature:       v.Author.Signature,
				TotalFavorited:  v.Author.TotalFavorited,
				WorkCount:       v.Author.WorkCount,
				FavoriteCount:   v.Author.FavoriteCount,
			},
		})
	}
	return &publish.ListVideoResp{
		VideoList: videos,
	}, nil
}
