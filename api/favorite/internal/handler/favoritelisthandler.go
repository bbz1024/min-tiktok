package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/favorite/internal/logic"
	"min-tiktok/api/favorite/internal/svc"
	"min-tiktok/api/favorite/internal/types"
)

func FavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteListLogic(r.Context(), svcCtx)
		resp, _ := l.FavoriteList(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
