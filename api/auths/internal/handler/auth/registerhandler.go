package auth

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/auths/internal/logic/auth"
	"min-tiktok/api/auths/internal/svc"
	"min-tiktok/api/auths/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		// param valid
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		// logic
		l := auth.NewRegisterLogic(r.Context(), svcCtx)
		resp, _ := l.Register(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
