package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/comment/comment"
	"min-tiktok/services/comment/internal/svc"
	"min-tiktok/services/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentLogic {
	return &ListCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCommentLogic) ListComment(in *comment.ListCommentRequest) (*comment.ListCommentResponse, error) {
	//l.svcCtx.CommentModel.F
	commentList, err := l.svcCtx.CommentModel.GetCommentList(l.ctx)
	if err != nil {
		return nil, err
	}
	var resp = new(comment.ListCommentResponse)
	var comments = make([]*comment.Comments, len(commentList))
	var err2 error
	var runner = threading.NewTaskRunner(10)
	for i, cmt := range commentList {
		order := i
		runner.Schedule(func() {
			c := &comment.Comments{
				Content:    cmt.Content.String,
				CreateDate: cmt.Createdat.Format("01-02"),
				Id:         uint32(cmt.Id),
			}
			res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserRequest{
				UserId:  uint32(cmt.Userid),
				ActorId: in.ActorId,
			})
			if err != nil {
				err2 = err
				l.Errorw("get user info error", logx.Field("err", err))
				return
			}
			if res.StatusCode != code.OK {
				resp.StatusCode = uint32(res.StatusCode)
				resp.StatusMsg = res.StatusMsg
				return
			}
			c.User = &comment.UserInfo{
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
			comments[order] = c
		})

	}
	runner.Wait()
	if err2 != nil {
		return nil, err2
	}
	resp.CommentList = comments
	return resp, nil
}
