package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/comment/internal/logic"
	"min-tiktok/api/comment/internal/svc"
	"min-tiktok/api/comment/internal/types"
)

func CommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentActionRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)

			return
		}
		l := logic.NewCommentActionLogic(r.Context(), svcCtx)
		resp, _ := l.CommentAction(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
