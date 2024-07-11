package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"min-tiktok/common/util/str2num"
	"min-tiktok/models/video"
	"min-tiktok/services/feed/feed"
	"min-tiktok/services/user/userclient"
)

// FetchVideoDetails fetches detailed information for a list of videos.
func FetchVideoDetails(ctx context.Context, videoList []*video.Video, actorId uint32, userRpc userclient.User, rdb *redis.Redis) ([]*feed.Video, error) {
	runner := threading.NewTaskRunner(10)
	videos := make([]*feed.Video, len(videoList))
	var err2 error
	for i, v := range videoList {
		order := i
		runner.Schedule(func() {
			videoInfo := &feed.Video{
				Id:       uint32(v.Id),
				PlayUrl:  v.Playurl,
				CoverUrl: v.Coverurl,
				Title:    v.Title,
			}
			// get video info from redis
			videoKey := fmt.Sprintf(keys.VideoInfoKey, v.Id)
			info, err := rdb.HgetallCtx(ctx, videoKey)
			if err != nil {
				err2 = err
				logx.Errorw("get video info from redis error", logx.Field("err", err))
				return
			}
			// get favorite and comment count from redis
			favoriteCntStr := info[keys.VideoFavoriteCount]
			commentCntStr := info[keys.VideoCommentCount]
			var favoriteCnt, commentCnt int
			var isFavorite bool

			if favoriteCnt, err = str2num.Str2Num(favoriteCntStr); err != nil {
				err2 = err
				logx.Errorw("get video info from redis error", logx.Field("err", err))
				return
			}
			if commentCnt, err = str2num.Str2Num(commentCntStr); err != nil {
				err2 = err
				logx.Errorw("get video info from redis error", logx.Field("err", err))
				return
			}
			videoInfo.FavoriteCount = uint32(favoriteCnt)
			videoInfo.CommentCount = uint32(commentCnt)

			// is favorite
			if actorId != 0 {
				key := fmt.Sprintf(keys.UserFavoriteKey, actorId)
				if isFavorite, err = rdb.SismemberCtx(ctx, key, uint32(v.Id)); err != nil {
					err2 = err
					logx.Errorw("get video info from redis error", logx.Field("err", err))
					return
				}
				videoInfo.IsFavorite = isFavorite
			}

			res, err := userRpc.GetUserInfo(ctx, &userclient.UserRequest{
				UserId:  uint32(v.Userid),
				ActorId: actorId,
			})
			if err != nil {
				err2 = err
				logx.Errorw("call rpc UserRpc.GetUserInfo error ", logx.Field("err", err))
				return
			}
			if res.StatusCode != code.OK {
				logx.Errorw("call rpc UserRpc.GetUserInfo error ", logx.Field("err", res.StatusMsg))
				return
			}
			videoInfo.Author = &feed.UserInfo{
				Id:              res.User.Id,
				Name:            res.User.Name,
				FollowCount:     res.User.FollowCount,
				FollowerCount:   res.User.FollowerCount,
				IsFollow:        res.User.IsFollow,
				Avatar:          res.User.Avatar,
				BackgroundImage: res.User.BackgroundImage,
				Signature:       res.User.Signature,
				TotalFavorited:  res.User.TotalFavorited,
				WorkCount:       res.User.WorkCount,
				FavoriteCount:   res.User.FavoriteCount,
			}
			videos[order] = videoInfo
		})
	}

	runner.Wait()

	if err2 != nil {
		return nil, err2
	}

	return videos, nil
}
