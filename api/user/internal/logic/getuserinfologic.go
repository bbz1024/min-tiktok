package logic

import (
	"context"
	"min-tiktok/services/user/user"

	"min-tiktok/api/user/internal/svc"
	"min-tiktok/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	// call rpc
	info, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserRequest{
		ActorId: uint32(req.ActorID),
		UserId:  uint32(req.UserID),
	})
	if err != nil {
		return &types.GetUserInfoResponse{
			StatusCode: info.StatusCode,
			StatusMsg:  info.StatusMsg,
		}, err
	}
	return &types.GetUserInfoResponse{
		User: types.User{
			ID:              int64(info.User.Id),
			Name:            info.User.Name,
			FollowCount:     int64(info.User.FollowCount),
			FollowerCount:   int64(info.User.FollowerCount),
			IsFollow:        info.User.IsFollow,
			Avatar:          info.User.Avatar,
			BackgroundImage: info.User.BackgroundImage,
			Signature:       info.User.Signature,
			TotalFavorited:  int64(info.User.TotalFavorited),
			WorkCount:       int64(info.User.WorkCount),
			FavoriteCount:   int64(info.User.FavoriteCount),
		},
	}, nil
}
