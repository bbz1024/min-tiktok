package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/favorite/internal/logic"
	"min-tiktok/api/favorite/internal/svc"
	"min-tiktok/api/favorite/internal/types"
)

func FavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		l := logic.NewFavoriteActionLogic(r.Context(), svcCtx)
		resp, _ := l.FavoriteAction(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
