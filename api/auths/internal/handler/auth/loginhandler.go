package auth

import (
	"min-tiktok/common/consts/code"
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/auths/internal/logic/auth"
	"min-tiktok/api/auths/internal/svc"
	"min-tiktok/api/auths/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		// param valid
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.ParamError, code.ParamErrorMsg))
			return
		}
		// bloom filter
		if !svcCtx.UserFilter.TestString(req.UserName) {
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.UserNotFoundCode, code.UserNotFoundMsg))
			return
		}
		// logic
		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, _ := l.Login(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
