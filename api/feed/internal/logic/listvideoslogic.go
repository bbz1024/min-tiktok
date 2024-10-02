package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/breaker"
	"min-tiktok/common/consts/code"
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
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "api")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var b = breaker.NewBreaker(breaker.WithName("ListVideos"))

func (l *ListVideosLogic) handleCircuitBreakerError(err error) {
	l.Errorw("circuit breaker open", logx.Field("error", err))
}
func (l *ListVideosLogic) ListVideos(req *types.ListVideosReq) (resp *types.ListVideosResp, err error) {
	var res *feedclient.ListFeedResponse
	err = b.DoWithAcceptable(func() error {
		var err error
		// not login
		if req.ActorId == 0 {
			res, err = l.svcCtx.FeedRpc.ListVideos(l.ctx,
				&feedclient.ListFeedRequest{
					LatestTime: req.LatestTime,
				},
			)
		} else {
			// recommend
			res, err = l.svcCtx.FeedRpc.ListRecommendVideos(l.ctx,
				&feedclient.ListRecommendRequest{
					ActorId: req.ActorId,
				},
			)
		}
		return err
	}, func(err error) bool {
		// TODO 熔断错误
		//if err != nil {
		//	// circuit
		//	l.Errorw("open circuit breaker", logx.Field("err", err))
		//	// 熔断，快速返回错误
		//	res = &feedclient.ListFeedResponse{
		//		StatusCode: code.TooManyRequestCode,
		//		StatusMsg:  code.TooManyRequestMsg,
		//	}
		//	// true 代表错误可接受，不会进行重试，false 代表错误不可接受，进行重试
		//	return true
		//}
		return false
	})
	resp = new(types.ListVideosResp)
	if err != nil {
		resp.StatusMsg = code.ServerErrorMsg
		resp.StatusCode = code.ServerError
		l.Errorw("call rpc FeedRpc.ListVideos error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusMsg = res.StatusMsg
		resp.StatusCode = res.StatusCode
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
