package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"strconv"

	"min-tiktok/services/favorite/favorite"
	"min-tiktok/services/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *favorite.FavoriteRequest) (*favorite.FavoriteResponse, error) {
	// valid VideoId
	// valid ActorId
	favoriteKey := fmt.Sprintf(keys.UserFavoriteKey, in.ActorId)
	userKey := fmt.Sprintf(keys.UserInfoKey, in.ActorId)
	videoInfoKey := fmt.Sprintf(keys.VideoInfoKey, in.VideoId)
	// get video author
	id, err := l.svcCtx.Rdb.HgetCtx(l.ctx, videoInfoKey, keys.VideoAuthorID)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if id == "" {
		l.Infow("author not found", logx.Field("videoID", in.VideoId))
		return &favorite.FavoriteResponse{
			StatusCode: code.UserNotFoundCode,
			StatusMsg:  code.UserNotFoundMsg,
		}, nil
	}
	authorId, _ := strconv.ParseInt(id, 10, 64)
	authorKey := fmt.Sprintf(keys.UserInfoKey, authorId)

	// note: should use lua
	increment := 1
	switch in.ActionType {
	case favorite.ActionType_FAVORITE:
		// already favorite
		ok, err := l.svcCtx.Rdb.SismemberCtx(l.ctx, favoriteKey, in.ActorId)
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		}
		if ok {
			return &favorite.FavoriteResponse{
				StatusCode: code.FavoriteRepeatCode,
				StatusMsg:  code.FavoriteRepeatMsg,
			}, nil
		}
		// 1. add videoId to favorite set
		if _, err := l.svcCtx.Rdb.SaddCtx(l.ctx, favoriteKey, in.ActorId); err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		}
	case favorite.ActionType_CANCEL_FAVORITE:
		ok, err := l.svcCtx.Rdb.SismemberCtx(l.ctx, favoriteKey, in.ActorId)
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		}
		if !ok {
			return &favorite.FavoriteResponse{
				StatusCode: code.FavoriteNotFoundCode,
				StatusMsg:  code.FavoriteNotFoundMsg,
			}, nil
		}
		if _, err := l.svcCtx.Rdb.SremCtx(l.ctx, favoriteKey, in.ActorId); err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		}
		increment = -1
	}
	// 2. incr/dec user favorite count
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, userKey, keys.FavoriteCount, increment); err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	// 3. incr/dec author  total favorite count
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, authorKey, keys.TotalFavorite, increment); err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	// 4. incr/dec video favorite count
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, videoInfoKey, keys.VideoFavoriteCount, increment); err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	return &favorite.FavoriteResponse{}, nil
}
