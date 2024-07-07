package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/user/userclient"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/services/feed/feed"
	"min-tiktok/services/feed/internal/svc"
)

type ListVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosLogic {
	return &ListVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func isUnixMilliTimestamp(s string) (time.Time, bool) {
	timestamp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.UnixMilli(timestamp), false
	}

	startTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Now().AddDate(100, 0, 0)

	t := time.UnixMilli(timestamp)
	res := t.After(startTime) && t.Before(endTime)

	return t, res
}

func (l *ListVideosLogic) ListVideos(in *feed.ListFeedRequest) (*feed.ListFeedResponse, error) {
	latestTime := time.Now()
	if in.LatestTime != "" {
		// Check if request.LatestTime is a timestamp
		t, ok := isUnixMilliTimestamp(in.LatestTime)
		if ok {
			latestTime = t
		}
		// if not ok return lately videos
	}
	videoList, err := l.svcCtx.VideoModel.ListVideoByCreateTime(l.ctx, latestTime)
	if err != nil {
		logx.Errorw("query video list failed", logx.Field("err", err))
		return &feed.ListFeedResponse{
			StatusCode: code.ServerError,
			StatusMsg:  code.ServerErrorMsg,
		}, err
	}
	runner := threading.NewTaskRunner(10)
	video := make([]*feed.Video, 0, len(videoList))
	for _, v := range videoList {
		videoInfo := &feed.Video{
			Id:       uint32(v.Id),
			PlayUrl:  v.Playurl,
			CoverUrl: v.Coverurl,
			Title:    v.Title,
		}
		res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userclient.UserRequest{
			UserId:  uint32(v.Userid),
			ActorId: in.ActorId,
		})
		if err != nil {
			logx.Errorw("query user info failed", logx.Field("err", err))
			return &feed.ListFeedResponse{
				StatusCode: code.ServerError,
				StatusMsg:  code.ServerErrorMsg,
			}, err
		}
		videoInfo.Author = &feed.UserInfo{
			Id:              res.User.Id,
			Name:            res.User.Name,
			FollowCount:     res.User.FollowCount,
			FollowerCount:   res.User.FollowerCount,
			IsFollow:        res.User.IsFollow,
			Avatar:          res.User.Avatar,
			BackgroundImage: res.User.BackgroundImage,
			Signature:       res.User.Signature,
			TotalFavorited:  res.User.TotalFavorited,
			WorkCount:       res.User.WorkCount,
			FavoriteCount:   res.User.FavoriteCount,
		}
		video = append(video, videoInfo)
	}
	runner.Wait()
	return &feed.ListFeedResponse{
		VideoList: video,
	}, nil
}
