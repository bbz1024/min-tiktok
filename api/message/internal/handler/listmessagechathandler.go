package handler

import (
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/message/internal/logic"
	"min-tiktok/api/message/internal/svc"
	"min-tiktok/api/message/internal/types"
)

func ListMessageChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageChatReq
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		l := logic.NewListMessageChatLogic(r.Context(), svcCtx)
		resp, _ := l.ListMessageChat(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
