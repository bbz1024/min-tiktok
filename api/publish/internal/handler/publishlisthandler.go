package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/publish/internal/logic"
	"min-tiktok/api/publish/internal/svc"
	"min-tiktok/api/publish/internal/types"
)

func PublishListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		l := logic.NewPublishListLogic(r.Context(), svcCtx)
		resp, _ := l.PublishList(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
