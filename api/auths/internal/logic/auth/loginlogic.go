package auth

import (
	"context"
	"min-tiktok/api/auths/internal/svc"
	"min-tiktok/api/auths/internal/types"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/auths/auths"

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
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc AuthsRpc.Login error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = res.StatusCode
		resp.StatusMsg = res.StatusMsg
		return
	}
	resp.Token = res.Token
	resp.UserID = res.UserId
	return

}
