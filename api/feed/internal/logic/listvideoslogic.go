package logic

import (
	"context"
	"min-tiktok/services/feed/feedclient"

	"min-tiktok/api/feed/internal/svc"
	"min-tiktok/api/feed/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosLogic {
	return &ListVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVideosLogic) ListVideos(req *types.ListVideosReq) (resp *types.ListVideosResp, err error) {
	var res *feedclient.ListFeedResponse
	// not login
	if req.ActorId == 0 {
		res, err = l.svcCtx.FeedRpc.ListVideos(
			l.ctx,
			&feedclient.ListFeedRequest{
				LatestTime: req.LatestTime,
				ActorId:    req.ActorId,
			},
		)
	} else {
		// recommend
		res, err = l.svcCtx.FeedRpc.ListVideos(
			l.ctx,
			&feedclient.ListFeedRequest{
				LatestTime: req.LatestTime,
				ActorId:    req.ActorId,
			},
		)
	}
	resp = new(types.ListVideosResp)
	resp.StatusMsg = res.StatusMsg
	resp.StatusCode = res.StatusCode
	if err != nil {
		l.Errorw("call rpc FeedRpc.ListVideos error ", logx.Field("err", err))
		return
	}
	var videoList []*types.Video
	for _, v := range res.VideoList {
		videoList = append(videoList, &types.Video{
			Id: v.Id,
			Author: &types.User{
				ID:              v.Author.Id,
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
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}

	resp.NextTime = res.NextTime
	resp.VideoList = videoList
	return
}
