package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/comment/comment"

	"min-tiktok/api/comment/internal/svc"
	"min-tiktok/api/comment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionRequest) (resp *types.CommentActionResponse, err error) {
	res, err := l.svcCtx.CommentRpc.ActionComment(l.ctx, &comment.ActionCommentRequest{
		ActorId:     req.ActorID,
		VideoId:     req.VideoID,
		ActionType:  comment.ActionCommentType(req.ActionType),
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	})
	resp = new(types.CommentActionResponse)
	if err != nil {
		resp.StatusMsg = code.ServerErrorMsg
		resp.StatusCode = code.ServerError
		l.Errorw("call rpc CommentRpc.ActionComment error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusMsg = res.StatusMsg
		resp.StatusCode = res.StatusCode
		return
	}
	return &types.CommentActionResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
		Comment: types.Comment{
			ID: res.Comment.Id,
			User: types.User{
				ID:              res.Comment.User.Id,
				Name:            res.Comment.User.Name,
				FollowCount:     res.Comment.User.FollowCount,
				FollowerCount:   res.Comment.User.FollowerCount,
				IsFollow:        res.Comment.User.IsFollow,
				Avatar:          res.Comment.User.Avatar,
				BackgroundImage: res.Comment.User.BackgroundImage,
				Signature:       res.Comment.User.Signature,
				TotalFavorited:  res.Comment.User.TotalFavorited,
				WorkCount:       res.Comment.User.WorkCount,
				FavoriteCount:   res.Comment.User.FavoriteCount,
			},
			Content:    res.Comment.Content,
			CreateDate: res.Comment.CreateDate,
		},
	}, nil

}
