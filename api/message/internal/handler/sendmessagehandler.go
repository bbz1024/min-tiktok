package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/message/internal/logic"
	"min-tiktok/api/message/internal/svc"
	"min-tiktok/api/message/internal/types"
)

func SendMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageActionReq
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		l := logic.NewSendMessageLogic(r.Context(), svcCtx)
		resp, _ := l.SendMessage(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
