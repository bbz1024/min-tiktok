package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"min-tiktok/common/store/qiniu"
	"min-tiktok/models/video"
	"min-tiktok/services/publish/internal/svc"
	"min-tiktok/services/publish/publish"
	"time"
)

type ActionVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionVideoLogic {
	return &ActionVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionVideoLogic) ActionVideo(in *publish.ActionVideoReq) (*publish.ActionVideoResp, error) {

	// store video to local
	/*
		videoName := fmt.Sprintf("%s.mp4", uuid.New().String())
		videoCoverName := fmt.Sprintf("%s.jpg", uuid.New().String())
		videoPath := path.Join(l.svcCtx.Config.StorePath, videoName)
		videoCoverPath := path.Join(l.svcCtx.Config.StorePath, videoCoverName)

		if err := os.MkdirAll(l.svcCtx.Config.StorePath, 0755); err != nil { // 0755 for read/write/execute permissions
			return nil, err
		}
		file, err := os.Create(videoPath)
		if err != nil {
			return nil, err
		}
		buf := bufio.NewWriter(file)
		if _, err := buf.Write(in.Data); err != nil {
			return nil, err
		}
		if err := buf.Flush(); err != nil {
			fmt.Println(err)
		}
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
		if _, err = capture.GetSnapshot(videoPath, videoCoverPath, 1); err != nil {
			return nil, err
		}

	*/
	//return &publish.ActionVideoResp{}, nil
	fileKey, err := qiniu.UploadToQiNiu(l.ctx,
		l.svcCtx.Config.QiNiu.AccessKey,
		l.svcCtx.Config.QiNiu.SecretKey,
		in.Data, "mp4",
		l.svcCtx.Config.QiNiu.Bucket,
	)
	if err != nil {
		logx.Errorw("upload video to qiniu ", logx.Field("err", err))
		return &publish.ActionVideoResp{
			StatusCode: code.ServerError,
			StatusMsg:  code.ServerErrorMsg,
		}, err
	}
	url := fmt.Sprintf("%s/%s", l.svcCtx.Config.QiNiu.VideoDomain, fileKey)
	// video insert to db
	now := time.Now()
	res, err := l.svcCtx.VideoModel.Insert(l.ctx, &video.Video{
		Userid:    uint64(in.ActorId),
		Playurl:   url,
		Coverurl:  url,
		Title:     in.Title,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		logx.Errorw("insert video to db ", logx.Field("err", err))
		return &publish.ActionVideoResp{
			StatusCode: code.ServerError,
			StatusMsg:  code.ServerErrorMsg,
		}, err
	}
	videoID, _ := res.LastInsertId()
	// user work count incr
	key := fmt.Sprintf(keys.UserInfoKey, in.ActorId)
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, key, keys.WorkCount, 1); err != nil && !errors.Is(err, redis.Nil) {
		logx.Errorw("incr user work count ", logx.Field("err", err))
		return &publish.ActionVideoResp{
			StatusCode: code.ServerError,
			StatusMsg:  code.ServerErrorMsg,
		}, err
	}
	// add video id to user video list
	key = fmt.Sprintf(keys.UserWorkKey, in.ActorId)
	if _, err := l.svcCtx.Rdb.SaddCtx(l.ctx, key, videoID); err != nil {
		logx.Errorw("add video id to user video list ", logx.Field("err", err))
		return &publish.ActionVideoResp{
			StatusCode: code.ServerError,
			StatusMsg:  code.ServerErrorMsg,
		}, err
	}
	return &publish.ActionVideoResp{}, nil
}
