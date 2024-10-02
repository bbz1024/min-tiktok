package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/keys"
	"min-tiktok/common/store/qiniu"
	"min-tiktok/models/video"
	"min-tiktok/services/publish/internal/mq"
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
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "service")),
	}
}

// ActionVideo 投稿
func (l *ActionVideoLogic) ActionVideo(in *publish.ActionVideoReq) (*publish.ActionVideoResp, error) {

	// 1. upload  video to qiniu
	fileKey, err := qiniu.UploadToQiNiu(l.ctx,
		l.svcCtx.Config.QiNiu.AccessKey,
		l.svcCtx.Config.QiNiu.SecretKey,
		in.Data, "mp4",
		l.svcCtx.Config.QiNiu.Bucket,
	)
	if err != nil {
		l.Errorw("upload video to qiniu ", logx.Field("err", err))
		return nil, err
	}

	// 2. video insert to db
	url := fmt.Sprintf("http://%s/%s", l.svcCtx.Config.QiNiu.VideoDomain, fileKey)
	now := time.Now().UTC()
	res, err := l.svcCtx.VideoModel.Insert(l.ctx, &video.Video{
		Userid:    uint64(in.ActorId),
		Playurl:   url,
		Coverurl:  url,
		Title:     in.Title,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		l.Errorw("insert video to db ", logx.Field("err", err))
		return nil, err
	}

	// 3. user work count incr
	videoID, _ := res.LastInsertId()
	key := fmt.Sprintf(keys.UserInfoKey, in.ActorId)
	if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, key, keys.WorkCount, 1); err != nil && !errors.Is(err, redis.Nil) {
		l.Errorw("incr user work count ", logx.Field("err", err))
		return nil, err
	}

	// 4. add video id to user video list
	key = fmt.Sprintf(keys.UserWorkKey, in.ActorId)
	if _, err := l.svcCtx.Rdb.SaddCtx(l.ctx, key, videoID); err != nil {
		l.Errorw("add video id to user video list ", logx.Field("err", err))
		return nil, err
	}

	// 5. get video author
	videoInfoKey := fmt.Sprintf(keys.VideoInfoKey, videoID)
	if err := l.svcCtx.Rdb.HsetCtx(l.ctx, videoInfoKey, keys.VideoAuthorID, fmt.Sprintf("%d", in.ActorId)); err != nil {
		l.Errorw("set video author id ", logx.Field("err", err))
		return nil, err
	}

	// -------------------- async --------------------
	// 6. put mq  extract video summery by gpt
	if err := mq.GetExtractVideoText().Product(mq.ExtractVideoTextReq{VideoID: uint32(videoID)}); err != nil {
		l.Errorw("extract video summery by gpt with queue ", logx.Field("err", err))
		return nil, err
	}
	l.Infow("upload video success", logx.Field("video_id", videoID), logx.Field("video_url", url))
	return &publish.ActionVideoResp{}, nil
}
