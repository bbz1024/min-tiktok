package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/comment/internal/logic"
	"min-tiktok/api/comment/internal/svc"
	"min-tiktok/api/comment/internal/types"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}

		l := logic.NewCommentListLogic(r.Context(), svcCtx)
		resp, _ := l.CommentList(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
