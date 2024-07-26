package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	user2 "min-tiktok/models/user"
	"min-tiktok/services/user/internal/svc"
	"min-tiktok/services/user/user"
	"strconv"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserRequest) (*user.UserResponse, error) {
	// get userinfo by mapreduce opt
	// must init because of the mapreduce not order
	var res = &user.UserResponse{
		StatusCode: code.OK,
		StatusMsg:  code.OkMsg,
	}
	localKey := fmt.Sprintf(keys.UserLocalCacheKey, in.UserId)
	// get from cache / mysql
	var userInfo = new(user.UserInfo)
	err := mr.Finish(func() error {
		localUserInfo, err := l.svcCtx.LocalCache.Take(localKey, func() (any, error) {
			// from cache or db
			l.Infow("get user info from cache", logx.Field("userId", in.UserId))
			return l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.UserId))
		})
		if err != nil {
			return err
		}
		v := localUserInfo.(*user2.Users)
		userInfo.Id = uint32(v.Id)
		userInfo.Name = v.Username
		userInfo.Avatar = v.Avatar.String
		userInfo.BackgroundImage = v.Backgroundimage.String
		userInfo.Signature = v.Signature.String
		return nil

	}, func() error {
		// from cache
		key := fmt.Sprintf(keys.UserInfoKey, in.UserId)
		info, err := l.svcCtx.Rdb.HgetallCtx(l.ctx, key)
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}
		followCount, _ := strconv.ParseInt(info[keys.FollowCount], 10, 32)
		followerCount, _ := strconv.ParseInt(info[keys.FollowerCount], 10, 32)
		totalFavorite, _ := strconv.ParseInt(info[keys.TotalFavorite], 10, 32)
		workCount, _ := strconv.ParseInt(info[keys.WorkCount], 10, 32)
		favoriteCount, _ := strconv.ParseInt(info[keys.FavoriteCount], 10, 32)

		userInfo.FollowCount = uint32(followCount)
		userInfo.FollowerCount = uint32(followerCount)
		userInfo.TotalFavorited = uint32(totalFavorite)
		userInfo.WorkCount = uint32(workCount)
		userInfo.FavoriteCount = uint32(favoriteCount)
		// authed user
		if in.ActorId != 0 && in.ActorId != in.UserId {
			// check exist in user follow list
			userFollowKey := fmt.Sprintf(keys.UserFollow, in.ActorId)
			exist, err := l.svcCtx.Rdb.SismemberCtx(l.ctx, userFollowKey, in.UserId)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					userInfo.IsFollow = false
					return nil
				}
				l.Errorw("get user is follow error", logx.Field("err", err))
				return err
			}
			userInfo.IsFollow = exist
		}
		return nil
	})
	// exist error
	if err != nil {
		if errors.Is(err, user2.ErrNotFound) {
			res.StatusCode = code.UserNotFoundCode
			res.StatusMsg = code.UserNotFoundMsg
			return res, nil
		}
		logx.Errorw("get user info error", logx.Field("err", err))
		return nil, err
	}
	res.User = userInfo
	return res, nil
}
