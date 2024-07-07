package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/feed/internal/logic"
	"min-tiktok/api/feed/internal/svc"
	"min-tiktok/api/feed/internal/types"
)

func ListVideosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListVideosReq
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}

		l := logic.NewListVideosLogic(r.Context(), svcCtx)
		resp, _ := l.ListVideos(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
