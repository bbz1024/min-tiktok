package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/relation/internal/logic"
	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"
)

func RelationActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RealtionActionReuqest

		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}

		l := logic.NewRelationActionLogic(r.Context(), svcCtx)
		resp, _ := l.RelationAction(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
