package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/favorite/favorite"
	"min-tiktok/services/favorite/internal/svc"
	"min-tiktok/services/feed/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "service")),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {

	// get user favorite set
	key := fmt.Sprintf(keys.UserFavoriteKey, in.UserId)
	members, err := l.svcCtx.Rdb.SmembersCtx(l.ctx, key)
	if err != nil && errors.Is(err, redis.Nil) {
		l.Errorw("redis error", logx.Field("err", err))
		return nil, err
	}
	resp := new(favorite.FavoriteListResponse)
	if len(members) == 0 {
		return resp, nil
	}
	res, err := l.svcCtx.FeedRpc.ListVideosBySet(l.ctx, &feed.ListVideosBySetRequest{
		VideoIdSet: members,
		ActorId:    in.ActorId,
	})
	if err != nil {
		l.Errorw("call rpc FeedRpc.ListVideosBySet", logx.Field("err", err))
		return nil, err
	}

	for _, v := range res.VideoList {
		resp.VideoList = append(resp.VideoList, &favorite.Video{
			Id: v.Id,
			Author: &favorite.UserInfo{
				Id:              v.Author.Id,
				Name:            v.Author.Name,
				FollowCount:     v.Author.FollowCount,
				FollowerCount:   v.Author.FollowerCount,
				IsFollow:        v.Author.IsFollow,
				Avatar:          v.Author.Avatar,
				BackgroundImage: v.Author.BackgroundImage,
				Signature:       v.Author.Signature,
				WorkCount:       v.Author.WorkCount,
				FavoriteCount:   v.Author.FavoriteCount,
				TotalFavorited:  v.Author.TotalFavorited,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	return resp, nil
}
