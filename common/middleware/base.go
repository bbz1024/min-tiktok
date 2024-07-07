package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/common/consts/code"
	"min-tiktok/common/response"
	"min-tiktok/services/auths/auths"
	"min-tiktok/services/auths/authsclient"
	"net/http"
	"slices"
	"strconv"
	"sync"
)

var WhitePath = []string{
	"/douyin/user/login",
	"/douyin/user/register",
	"/douyin/comment/list",
	"/douyin/publish/list",
	"/douyin/favorite/list",
}

// OptionPath 可能携带token的path
var OptionPath = []string{
	"/douyin/feed/",
	"/douyin/user/",
	"/douyin/relation/follow/list/",
	"/douyin/relation/follower/list/",
}
var once sync.Once
var authRpc authsclient.Auths

func WrapperAuthMiddleware(rpcConf zrpc.RpcClientConf) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// white path
			if slices.Contains(WhitePath, r.URL.Path) {
				//
				next(w, r)
				return
			}
			var token = r.PostFormValue("token")
			if r.Method == http.MethodGet {
				token = r.FormValue("token")
			}
			// optional token
			if token == "" && slices.Contains(OptionPath, r.URL.Path) {
				next(w, r)
				return
			}
			// init rpc
			once.Do(func() {
				authRpc = authsclient.NewAuths(zrpc.MustNewClient(rpcConf))
			})
			// auth
			res, err := authRpc.Authentication(r.Context(), &auths.AuthsRequest{
				Token: token,
			})

			// back err
			if err != nil {
				httpx.OkJsonCtx(r.Context(), w, response.NewResponse(code.ServerError, code.ServerErrorMsg))
				return
			}
			// auth failed
			if res.StatusCode != 0 {
				httpx.OkJsonCtx(r.Context(), w, response.NewResponse(int(res.StatusCode), res.StatusMsg))
				return
			}
			// with actor_id
			r.Form.Set("actor_id", strconv.Itoa(int(res.UserId)))
			next(w, r)
		}
	}
}
