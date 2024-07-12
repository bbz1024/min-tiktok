package logic

import (
	"context"
	"database/sql"
	"fmt"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/consts/keys"
	"min-tiktok/common/consts/variable"
	"min-tiktok/services/feedback/feedback"
	"min-tiktok/services/user/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	comment2 "min-tiktok/models/comment"
	"min-tiktok/services/comment/comment"
	"min-tiktok/services/comment/internal/svc"
)

type ActionCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionCommentLogic {
	return &ActionCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionCommentLogic) ActionComment(in *comment.ActionCommentRequest) (*comment.ActionCommentResponse, error) {
	var resp = new(comment.ActionCommentResponse)
	now := time.Now().UTC()
	key := fmt.Sprintf(keys.VideoInfoKey, in.VideoId)
	switch in.ActionType {
	case comment.ActionCommentType_ActionDelete:
		if err := l.svcCtx.CommentModel.Delete(l.ctx, uint64(in.CommentId)); err != nil {
			l.Errorw("delete comment error", logx.Field("err", err))
			return nil, err
		}
		resp.Comment = new(comment.Comments)
		resp.Comment.User = new(comment.UserInfo)
		// desc video comment count
		if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, key, keys.VideoCommentCount, -1); err != nil {
			return nil, err
		}
	case comment.ActionCommentType_ActionCreate:
		row, err := l.svcCtx.CommentModel.Insert(l.ctx, &comment2.Comment{
			Userid:    uint64(in.ActorId),
			Content:   sql.NullString{String: in.CommentText, Valid: true},
			Videoid:   uint64(in.VideoId),
			Createdat: now,
			Updatedat: now,
		})

		if err != nil {
			l.Errorw("insert comment error", logx.Field("err", err))
			return nil, err
		}
		id, err := row.LastInsertId()
		if err != nil {
			l.Errorw("get last insert id error", logx.Field("err", err))
			return nil, err
		}
		// incr video comment count
		if _, err := l.svcCtx.Rdb.HincrbyCtx(l.ctx, key, keys.VideoCommentCount, 1); err != nil {
			return nil, err
		}
		resp.Comment = &comment.Comments{
			Id:         uint32(id),
			CreateDate: now.Format("2006-01-02 15:04"),
			Content:    in.CommentText,
		}
		res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserRequest{
			UserId: in.ActorId,
		})
		if err != nil || res.StatusCode != code.OK {
			l.Errorw("get user info error", logx.Field("err", err))
			return nil, err
		}
		//  feedback
		feedres, err := l.svcCtx.FeedbackRpc.Feedback(l.ctx, &feedback.FeedbackRequest{
			UserId:   in.ActorId,
			VideoIds: []uint32{in.VideoId},
			Type:     variable.CommentFeedBack,
		})
		if err != nil || feedres.StatusCode != code.OK {
			l.Errorw("feedback error", logx.Field("err", err))
		}
		resp.Comment.User = &comment.UserInfo{
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
	}
	return resp, nil
}
