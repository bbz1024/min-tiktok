package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/user/internal/logic"
	"min-tiktok/api/user/internal/svc"
	"min-tiktok/api/user/internal/types"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, _ := l.GetUserInfo(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
