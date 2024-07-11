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

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Follow 关注
func (l *FollowLogic) Follow(in *relation.RelationActionRequest) (*relation.RelationActionResponse, error) {
	/*
		user_id: 作者id
		actor_id: 我的id
	*/
	followKey := fmt.Sprintf(keys.UserFollow, in.ActorId)
	// check user is follow actor
	isFollow, err := l.svcCtx.Rdb.SismemberCtx(l.ctx, followKey, in.UserId)
	if err != nil {
		return nil, err
	}
	if isFollow {
		return &relation.RelationActionResponse{
			StatusCode: code.IsFollowCode,
			StatusMsg:  code.IsFollowMsg,
		}, nil
	}

	// actor follow user
	if _, err := l.svcCtx.Rdb.SaddCtx(l.ctx, followKey, in.UserId); err != nil {
		return nil, err
	}
	// put actor in user follower set
	followerKey := fmt.Sprintf(keys.UserFollower, in.UserId)
	if _, err := l.svcCtx.Rdb.SaddCtx(l.ctx, followerKey, in.ActorId); err != nil {
		return nil, err
	}
	// user follow count + 1
	useKey := fmt.Sprintf(keys.UserInfoKey, in.ActorId)
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, useKey, keys.FollowCount, 1); err != nil {
		return nil, err
	}
	// actor follower count + 1
	actorKey := fmt.Sprintf(keys.UserInfoKey, in.UserId)
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, actorKey, keys.FollowerCount, 1); err != nil {
		return nil, err
	}
	return &relation.RelationActionResponse{}, nil
}
