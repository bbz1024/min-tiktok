package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/relation/internal/logic"
	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"
)

func RelationFollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFollowListRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)

			return
		}

		l := logic.NewRelationFollowListLogic(r.Context(), svcCtx)
		resp, _ := l.RelationFollowList(&req)

		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
