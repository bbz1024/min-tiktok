package response

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/common/consts/code"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func NewResponse(statusCode int, statusMsg string) *Response {
	return &Response{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
func NewParamError(ctx context.Context, w http.ResponseWriter, err error) {
	logx.Infow("params invalid", logx.Field("err", err))
	httpx.OkJsonCtx(ctx, w, NewResponse(code.ParamError, code.ParamErrorMsg))
}
func NewBackError(ctx context.Context, w http.ResponseWriter, err error) {
	logx.Errorw("back error", logx.Field("err", err))
	httpx.OkJsonCtx(ctx, w, NewResponse(code.ServerError, code.ServerErrorMsg))
	
}
