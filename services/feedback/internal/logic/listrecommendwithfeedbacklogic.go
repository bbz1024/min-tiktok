package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/keys"
	"min-tiktok/common/consts/variable"
	"min-tiktok/services/feedback/feedback"
	"min-tiktok/services/feedback/internal/svc"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRecommendWithFeedbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRecommendWithFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRecommendWithFeedbackLogic {
	return &ListRecommendWithFeedbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "service")),
	}
}

func (l *ListRecommendWithFeedbackLogic) ListRecommendWithFeedback(in *feedback.ListRecommendRequest) (*feedback.ListRecommendResponse, error) {
	key := fmt.Sprintf(keys.UserOffsetKey, in.ActorId)
	offsetStr, err := l.svcCtx.Rdb.Get(key)
	var offset int
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			l.Errorw("get offset error", logx.Field("err", err))
			return nil, err
		}
		offset = 0
	}
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			l.Errorw("parse offset error", logx.Field("err", err))
			return nil, err
		}
	}

	recommendSet, err := l.svcCtx.Recommend.GetItemRecommend(
		l.ctx,
		strconv.Itoa(int(in.ActorId)),
		[]string{},
		variable.ReadFeedBack,
		"5m",
		int(in.Count),
		offset,
	)
	// update offset
	length := len(recommendSet)
	if length > 0 {
		offset += length
		err = l.svcCtx.Rdb.SetCtx(l.ctx, key, strconv.Itoa(offset+length))
		if err != nil {
			l.Errorw("set offset error", logx.Field("err", err))
			return nil, err
		}
	}
	// reset offset because of the end of the list
	if length < int(in.Count) {
		err = l.svcCtx.Rdb.SetCtx(l.ctx, key, "0")
		if err != nil {
			l.Errorw("set offset error", logx.Field("err", err))
			return nil, err
		}
	}

	if err != nil {
		l.Errorw("feedback error", logx.Field("err", err))
		return nil, err
	}

	return &feedback.ListRecommendResponse{
		VideoIds: recommendSet,
	}, nil
}
