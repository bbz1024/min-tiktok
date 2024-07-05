package logic

import (
	"context"
	"errors"
	"fmt"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"min-tiktok/common/cryptx"
	"min-tiktok/common/uid"
	"min-tiktok/models/user"
	"min-tiktok/services/auths/auths"
	"min-tiktok/services/auths/internal/svc"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *auths.LoginRequest) (*auths.LoginResponse, error) {
	// get user info
	userinfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return &auths.LoginResponse{
				StatusCode: code.UserNotFoundCode,
				StatusMsg:  code.UserNotFoundMsg,
			}, nil
		}
		// other err
		return nil, err
	}
	if !cryptx.PasswordVerify(in.Password, userinfo.Password) {
		return &auths.LoginResponse{
			StatusCode: code.UserPasswordErrorCode,
			StatusMsg:  code.UserPasswordErrorMsg,
		}, nil
	}
	// 生成token，使用session（这里简单的使用uuid）,存储在redis
	token := uid.GenUid(l.ctx, int(userinfo.Id))
	key := fmt.Sprintf(keys.UserTokenKey, token)
	if err := l.svcCtx.Rdb.SetCtx(l.ctx, key, strconv.FormatUint(userinfo.Id, 10)); err != nil {
		return nil, err
	}
	return &auths.LoginResponse{
		UserId: uint32(userinfo.Id),
		Token:  token,
	}, nil
}
