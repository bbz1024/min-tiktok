package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	user2 "min-tiktok/models/user"
	"min-tiktok/services/user/internal/svc"
	"min-tiktok/services/user/user"
)

type CheckUserExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserExistLogic {
	return &CheckUserExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserExistLogic) CheckUserExist(in *user.UserExistRequest) (*user.UserExistResponse, error) {
	var userinfo *user2.Users
	var err error
	// check is  zero value
	if in.UserId == 0 {
		goto notExist
	}
	// check with bloom
	//if !l.svcCtx.UserBloom.TestString(strconv.Itoa(int(in.UserId))) {
	//	goto notExist
	//}
	// check with db
	userinfo, err = l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.UserId))
	if err != nil {
		goto notExist
	}
	if userinfo == nil || userinfo.Id == 0 {
		goto notExist
	}
	return &user.UserExistResponse{
		Exist: true,
	}, nil
notExist:
	return &user.UserExistResponse{
		Exist: false,
	}, nil
}
