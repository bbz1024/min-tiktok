package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relation"
	"min-tiktok/services/user/userclient"
	"strconv"
)

// 定义一个通用的逻辑处理方法，用于获取用户列表
func fetchUserList(ctx context.Context,
	key string, actorID uint32,
	rdb *redis.Redis, userRpc userclient.User) ([]*relation.UserInfo, error) {
	set, err := rdb.SmembersCtx(ctx, key)
	if err != nil {
		return nil, err
	}
	userList := make([]*relation.UserInfo, 0, len(set))

	runner := threading.NewTaskRunner(10)
	var err2 error
	for _, followId := range set {
		userId, err := strconv.ParseInt(followId, 10, 64)
		if err != nil {
			logx.Errorw("strconv.ParseInt failed", logx.Field("err", err))
			return nil, err
		}
		runner.Schedule(func() {
			// 这里假设in和response都是具有ActorId字段的结构体
			res, err := userRpc.GetUserInfo(ctx, &userclient.UserRequest{
				ActorId: actorID,
				UserId:  uint32(userId),
			})
			if err != nil {
				logx.Errorw("call rpc UserRpc.GetUserInfo failed ", logx.Field("err", err))
				err2 = err
				return
			}
			if res.StatusCode != code.OK {
				logx.Errorw("call rpc UserRpc.GetUserInfo failed", logx.Field("err", res.StatusMsg))
				err2 = err
				return
			}
			userList = append(userList, &relation.UserInfo{
				Id:              res.User.Id,
				Name:            res.User.Name,
				FollowCount:     res.User.FollowCount,
				FollowerCount:   res.User.FollowerCount,
				Avatar:          res.User.Avatar,
				BackgroundImage: res.User.BackgroundImage,
				Signature:       res.User.Signature,
				TotalFavorited:  res.User.TotalFavorited,
				WorkCount:       res.User.WorkCount,
				FavoriteCount:   res.User.FavoriteCount,
				IsFollow:        res.User.IsFollow,
			})
		})
	}
	runner.Wait()
	if err2 != nil {
		logx.Errorw("call rpc UserRpc.GetUserInfo failed", logx.Field("err", err2))
		return nil, err2
	}
	return userList, nil
}
