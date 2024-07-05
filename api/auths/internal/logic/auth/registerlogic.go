package auth

import (
	"context"
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
	if err != nil {
		return &types.RegisterResp{
			StatusCode: res.StatusCode,
			StatusMsg:  res.StatusMsg,
		}, err
	}

	// register success before put in bloom
	l.svcCtx.UserFilter.Add([]byte(req.UserName))
	return &types.RegisterResp{
		Token:      res.Token,
		UserID:     int64(res.UserId),
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}, nil
}
