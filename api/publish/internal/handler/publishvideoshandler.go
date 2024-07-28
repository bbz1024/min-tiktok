package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/syncx"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"min-tiktok/api/publish/internal/logic"
	"min-tiktok/api/publish/internal/svc"
	"min-tiktok/api/publish/internal/types"
)

const MaxSize = 1 << 20 * 100 // 100M
const MaxLimit = 10

var limiter = syncx.NewLimit(MaxLimit)

func PublishVideosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.TryBorrow() {
			logx.Infow("limit trigger")
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.TooManyRequestCode, code.TooManyRequestMsg))
			return
		}
		defer func() {
			if err := limiter.Return(); err != nil {
				httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.TooManyRequestCode, code.TooManyRequestMsg))
				logx.Errorw("limiter return error", logx.Field("err", err))
			}
		}()

		var req types.PublishActionReq
		if err := httpx.ParseForm(r, &req); err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		// get file
		data, header, err := r.FormFile("data")
		if err != nil {
			response.NewParamError(r.Context(), w, err)
			return
		}
		// valid size
		if header.Size > MaxSize {
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.VideoOverSizeCode, code.VideoOverSizeMsg))
			return
		}
		// read
		var b = make([]byte, header.Size)
		if _, err := data.Read(b); err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.ServerError, code.ServerErrorMsg))
			return
		}
		// valid file type

		detectedContentType := http.DetectContentType(b)
		if detectedContentType != "video/mp4" {
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.VideoTypeCode, code.VideoTypeMsg))
			return
		}
		// close stream
		if err := data.Close(); err != nil {
			httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.ServerError, code.ServerErrorMsg))
			return
		}
		l := logic.NewPublishVideosLogic(r.Context(), svcCtx)
		req.Data = b
		resp, _ := l.PublishVideos(&req)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
