package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/feed/feed"
	"min-tiktok/services/feed/internal/svc"
	"min-tiktok/services/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideosByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosByUserIDLogic {
	return &ListVideosByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListVideosByUserID query by user_id 获取某个用户的视频列表
func (l *ListVideosByUserIDLogic) ListVideosByUserID(in *feed.ListVideosByUserIDRequest) (*feed.ListVideosByUserIDResponse, error) {
	videoList, err := l.svcCtx.VideoModel.ListVideoByUserId(l.ctx, int64(in.UserId))
	if err != nil {
		logx.Errorw("query video list by user_id error", logx.Field("err", err))
		return &feed.ListVideosByUserIDResponse{
			StatusCode: code.ServerError,
			StatusMsg:  code.ServerErrorMsg,
		}, err
	}
	runner := threading.NewTaskRunner(10)
	video := make([]*feed.Video, 0, len(videoList))
	for _, v := range videoList {
		videoInfo := &feed.Video{
			Id:       uint32(v.Id),
			PlayUrl:  v.Playurl,
			CoverUrl: v.Coverurl,
			Title:    v.Title,
		}
		//runner.Schedule(func() {
		res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userclient.UserRequest{
			UserId:  in.UserId,
			ActorId: in.ActorId,
		})
		if err != nil {
			logx.Errorw("call rpc UserRpc.GetUserInfo error ", logx.Field("err", err))
			return &feed.ListVideosByUserIDResponse{
				StatusCode: code.ServerError,
				StatusMsg:  code.ServerErrorMsg,
			}, err
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
		video = append(video, videoInfo)
	}
	runner.Wait()
	return &feed.ListVideosByUserIDResponse{
		VideoList: video,
	}, nil

}
