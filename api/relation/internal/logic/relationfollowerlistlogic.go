package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relationclient"

	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(req *types.RelationFollowerListRequest) (resp *types.RelationFollowerListResponse, err error) {
	res, err := l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &relationclient.FollowerListRequest{
		ActorId: req.ActorID,
		UserId:  req.UserID,
	})
	resp = new(types.RelationFollowerListResponse)
	if err != nil {
		resp.StatusMsg = code.ServerErrorMsg
		resp.StatusCode = code.ServerError
		return
	}
	if res.StatusCode != 0 {
		resp.StatusMsg = res.StatusMsg
		resp.StatusCode = uint32(res.StatusCode)
		return resp, nil
	}
	resp.UserList = make([]*types.User, len(res.UserList))
	for i, user := range res.UserList {
		resp.UserList[i] = &types.User{
			ID:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
	}
	return
}
