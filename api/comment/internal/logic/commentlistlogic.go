package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/comment/comment"

	"min-tiktok/api/comment/internal/svc"
	"min-tiktok/api/comment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "api")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	res, err := l.svcCtx.CommentRpc.ListComment(l.ctx, &comment.ListCommentRequest{
		VideoId: req.VideoID,
		ActorId: req.ActorID,
	})

	resp = new(types.CommentListResponse)
	if err != nil {
		resp.StatusMsg = code.ServerErrorMsg
		resp.StatusCode = code.ServerError
		l.Errorw("call rpc CommentRpc.ListComment error ", logx.Field("err", err))
		return
	}

	if res.StatusCode != code.OK {
		resp.StatusMsg = res.StatusMsg
		resp.StatusCode = res.StatusCode
		return
	}
	for _, v := range res.CommentList {
		resp.CommentList = append(resp.CommentList, &types.Comment{
			ID: v.Id,
			User: types.User{
				ID:              v.User.Id,
				Name:            v.User.Name,
				FollowCount:     v.User.FollowCount,
				FollowerCount:   v.User.FollowerCount,
				IsFollow:        v.User.IsFollow,
				Avatar:          v.User.Avatar,
				BackgroundImage: v.User.BackgroundImage,
				Signature:       v.User.Signature,
				TotalFavorited:  v.User.TotalFavorited,
				WorkCount:       v.User.WorkCount,
				FavoriteCount:   v.User.FavoriteCount,
			},
			Content:    v.Content,
			CreateDate: v.CreateDate,
		})
	}
	return
}
