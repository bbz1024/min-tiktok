package logic

import (
	"context"
	"fmt"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/relation/internal/svc"
	"min-tiktok/services/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnfollowLogic) Unfollow(in *relation.RelationActionRequest) (*relation.RelationActionResponse, error) {

	// actor follow user
	followKey := fmt.Sprintf(keys.UserFollow, in.ActorId)
	// check user follow user
	isFollow, err := l.svcCtx.Rdb.SismemberCtx(l.ctx, followKey, in.UserId)
	if err != nil {
		return nil, err
	}
	if !isFollow {
		return &relation.RelationActionResponse{
			StatusCode: code.IsNotFollowCode,
			StatusMsg:  code.IsNotFollowMsg,
		}, nil
	}
	if _, err := l.svcCtx.Rdb.SremCtx(l.ctx, followKey, in.UserId); err != nil {
		return nil, err
	}

	// remove actor from user follow list
	followerKey := fmt.Sprintf(keys.UserFollower, in.UserId)
	if _, err := l.svcCtx.Rdb.SremCtx(l.ctx, followerKey, in.ActorId); err != nil {
		return nil, err
	}
	// user follow count - 1
	useKey := fmt.Sprintf(keys.UserInfoKey, in.ActorId)
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, useKey, keys.FollowCount, -1); err != nil {
		return nil, err
	}
	// actor follower count - 1
	actorKey := fmt.Sprintf(keys.UserInfoKey, in.UserId)
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, actorKey, keys.FollowerCount, -1); err != nil {
		return nil, err
	}
	return &relation.RelationActionResponse{}, nil
}
