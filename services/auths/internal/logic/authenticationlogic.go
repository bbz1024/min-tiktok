package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"strconv"

	"min-tiktok/services/auths/auths"
	"min-tiktok/services/auths/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthenticationLogic) Authentication(in *auths.AuthsRequest) (*auths.AuthsResponse, error) {
	if in.Token == "" {
		return &auths.AuthsResponse{
			StatusCode: code.AuthErrorCode,
			StatusMsg:  code.AuthErrorMsg,
		}, nil
	}

	var id uint64
	key := fmt.Sprintf(keys.UserTokenKey, in.Token)
	userId, err := l.svcCtx.Rdb.GetCtx(l.ctx, key)
	if userId == "" {
		goto authErr
	}
	if err != nil {
		if errors.Is(err, redis.Nil) {
			goto authErr
		}
		return nil, err
	}
	id, err = strconv.ParseUint(userId, 10, 32)
	if err == nil {
		return &auths.AuthsResponse{
			UserId: uint32(id),
		}, nil
	}
authErr:
	return &auths.AuthsResponse{
		StatusCode: code.AuthErrorCode,
		StatusMsg:  code.AuthErrorMsg,
	}, nil
}
