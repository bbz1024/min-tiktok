package logic

import (
	"context"
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
	// 1. get latest time
	latestTime := time.Now()
	if in.LatestTime != "" {
		// Check if request.LatestTime is a timestamp
		t, ok := isUnixMilliTimestamp(in.LatestTime)
		if ok {
			latestTime = t
		}
		// if not ok return lately videos
	}
	// 2. query video list by create time
	videoList, err := l.svcCtx.VideoModel.ListVideoByCreateTime(l.ctx, latestTime)
	if err != nil {
		logx.Errorw("query video list failed", logx.Field("err", err))
		return nil, err
	}

	videos, err := FetchVideoDetails(l.ctx, videoList, in.ActorId, l.svcCtx.UserRpc, l.svcCtx.Rdb)
	if err != nil {
		l.Errorw("fetch video details failed", logx.Field("err", err))
		return nil, err
	}
	resp := &feed.ListFeedResponse{
		VideoList: videos,
	}

	return resp, nil
}
