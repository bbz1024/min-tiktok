package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	"min-tiktok/common/consts/code"
	"net/http"
)

func WithMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				httpx.Error(w, errors.New(code.ServerError, code.ServerErrorMsg))
				return
			}
		}()
		next(w, r)
	}
}
