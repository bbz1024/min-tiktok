package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relationclient"
	"min-tiktok/services/user/userclient"

	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListRequest) (resp *types.RelationFollowListResponse, err error) {
	// check user_id exist
	exist, err := l.svcCtx.UserRpc.CheckUserExist(l.ctx, &userclient.UserExistRequest{
		UserId: req.UserID,
	})
	if err != nil {
		l.Errorw("call rpc UserRpc.CheckUserExist", logx.Field("err", err))
		return nil, err
	}
	if !exist.Exist {
		l.Infow("user not found", logx.Field("user_id", req.UserID))
		return &types.RelationFollowListResponse{
			StatusCode: code.UserNotFoundCode,
			StatusMsg:  code.UserNotFoundMsg,
		}, nil
	}
	res, err := l.svcCtx.RelationRpc.GetFollowList(
		l.ctx,
		&relationclient.FollowListRequest{
			ActorId: req.ActorID,
			UserId:  req.UserID,
		})
	resp = new(types.RelationFollowListResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		return
	}
	resp = new(types.RelationFollowListResponse)
	if res.StatusCode != code.OK {
		resp.StatusCode = uint32(res.StatusCode)
		resp.StatusMsg = res.StatusMsg
		return
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
