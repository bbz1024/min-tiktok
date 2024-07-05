package handler

import (
	"fmt"
	"min-tiktok/common/consts/code"
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
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.ParamError, code.ParamErrorMsg))
			return
		}
		fmt.Println(req)
		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, _ := l.GetUserInfo(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
