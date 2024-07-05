package auth

import (
	"context"
	"min-tiktok/services/auths/auths"

	"min-tiktok/api/auths/internal/svc"
	"min-tiktok/api/auths/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// call rpc
	res, err := l.svcCtx.AuthsRpc.Login(l.ctx, &auths.LoginRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return &types.LoginResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}, err
	}

	return &types.LoginResp{
		Token:      res.Token,
		UserID:     int64(res.UserId),
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}, nil
}
