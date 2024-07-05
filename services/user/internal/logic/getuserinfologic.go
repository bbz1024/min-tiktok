package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/user/internal/svc"
	"min-tiktok/services/user/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
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
	var res = new(user.UserResponse)
	// get userinfo by mapreduce opt
	err := mr.Finish(func() error {
		// from db
		userinfo, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.ActorId))
		if err != nil {
			return err
		}
		res.User = &user.UserInfo{
			Id:              uint32(userinfo.Id),
			Name:            userinfo.Username,
			Avatar:          userinfo.Avatar.String,
			BackgroundImage: userinfo.Backgroundimage.String,
			Signature:       userinfo.Signature.String,
		}
		return nil

	}, func() error {
		// from cache
		key := fmt.Sprintf(keys.UserInfoKey, in.ActorId)
		info, err := l.svcCtx.Rdb.HgetallCtx(l.ctx, key)
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}
		if len(info) > 0 {
			/*
				FollowCount   = "follow_count"   // 关注数量
				FollowerCount = "follower_count" // 粉丝数量
				TotalFavorite = "total_favorite" // 获赞数量
				WorkCount     = "work_count"     // 作品数量
				FavoriteCount = "favorite_count" // 点赞数量
			*/
			followCount, _ := strconv.ParseInt(info[keys.FollowCount], 10, 32)
			followerCount, _ := strconv.ParseInt(info[keys.FollowerCount], 10, 32)
			totalFavorite, _ := strconv.ParseInt(info[keys.TotalFavorite], 10, 32)
			workCount, _ := strconv.ParseInt(info[keys.WorkCount], 10, 32)
			favoriteCount, _ := strconv.ParseInt(info[keys.FavoriteCount], 10, 32)
			res.User.FollowCount = uint32(followCount)
			res.User.FollowerCount = uint32(followerCount)
			res.User.TotalFavorited = uint32(totalFavorite)
			res.User.WorkCount = uint32(workCount)
			res.User.FavoriteCount = uint32(favoriteCount)
			return nil
		}
		return nil
	})
	// exist error
	if err != nil {
		return nil, err
	}
	return res, nil
}
