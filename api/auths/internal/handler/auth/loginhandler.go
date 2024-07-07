package auth

import (
	"github.com/zeromicro/go-zero/core/logx"
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
			response.NewParamError(r.Context(), w, err)
			return
		}
		// bloom filter
		if !svcCtx.UserFilter.TestString(req.UserName) {
			logx.WithContext(r.Context()).Infow("user not found", logx.Field("userName", req.UserName))
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.UserNotFoundCode, code.UserNotFoundMsg))
			return
		}
		// logic
		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, _ := l.Login(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
