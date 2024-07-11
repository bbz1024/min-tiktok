package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zhenghaoz/gorse/client"
	"io"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"min-tiktok/common/cryptx"
	"min-tiktok/common/uid"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/models/user"
	"min-tiktok/services/auths/auths"
	"min-tiktok/services/auths/internal/svc"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *auths.RegisterRequest) (*auths.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		// user not exist
		now := time.Now()
		if errors.Is(err, user.ErrNotFound) {
			u := &user.Users{
				Username:  in.Username,
				Password:  cryptx.PasswordEncrypt(in.Password),
				Createdat: now,
				Updatedat: now,
			}
			// signature
			resp, err := http.Get(l.svcCtx.Config.UserInfo.SignatureUrl)
			if err != nil || resp.StatusCode != http.StatusOK {
				u.Signature = sql.NullString{String: u.Username, Valid: true}
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				u.Signature = sql.NullString{String: u.Username, Valid: true}
			}
			u.Signature = sql.NullString{String: string(body), Valid: true}
			u.Avatar = sql.NullString{
				String: fmt.Sprintf("%s/%s.png", l.svcCtx.Config.UserInfo.AvatarUrl, url.QueryEscape(u.Username)),
				Valid:  true,
			}
			u.Backgroundimage = sql.NullString{String: l.svcCtx.Config.UserInfo.BackImageUrl, Valid: true}
			userinfo, err := l.svcCtx.UserModel.Insert(l.ctx, u)
			if err != nil {
				return nil, err
			}
			userID, err := userinfo.LastInsertId()
			if err != nil {
				return nil, err
			}
			// gen token & set
			// 生成token，使用session（这里简单的使用uuid）,存储在redis
			token := uid.GenUid(l.ctx, int(userID))
			key := fmt.Sprintf(keys.UserTokenKey, token)
			if err := l.svcCtx.Rdb.SetCtx(l.ctx, key, strconv.FormatUint(uint64(userID), 10)); err != nil {
				logx.Errorw("set token error: %s", logx.Field("err", err))
				return nil, err
			}
			if _, err := l.svcCtx.GorseClient.InsertUser(l.ctx, client.User{
				UserId:  fmt.Sprintf("%d", userID),
				Comment: in.Username,
			}); err != nil {
				logx.Errorw("insert user error: %s", logx.Field("err", err))
				return nil, err
			}
			return &auths.RegisterResponse{
				UserId: uint32(userID),
				Token:  token,
			}, nil
		}
		// other error
		return nil, err
	}
	// user exist
	return &auths.RegisterResponse{
		StatusCode: code.UserExistedCode,
		StatusMsg:  code.UserExistedMsg,
	}, nil
}
