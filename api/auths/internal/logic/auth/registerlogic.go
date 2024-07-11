package auth

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/auths/auths"

	"min-tiktok/api/auths/internal/svc"
	"min-tiktok/api/auths/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// call rpc
	// res is nil if exist err
	res, err := l.svcCtx.AuthsRpc.Register(l.ctx, &auths.RegisterRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	resp = new(types.RegisterResp)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc AuthsRpc.Register error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = res.StatusCode
		resp.StatusMsg = res.StatusMsg
		return
	}
	// register success before put in bloom
	l.svcCtx.UserFilter.Add([]byte(req.UserName))
	resp.Token = res.Token
	resp.UserID = res.UserId
	return
}
