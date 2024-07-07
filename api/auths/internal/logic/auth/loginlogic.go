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
	resp = new(types.LoginResp)
	resp.StatusMsg = res.StatusMsg
	resp.StatusCode = res.StatusCode
	if err != nil {
		l.Errorw("call rpc AuthsRpc.Login error ", logx.Field("err", err))
		return
	}
	resp.Token = res.Token
	resp.UserID = res.UserId
	return

}
