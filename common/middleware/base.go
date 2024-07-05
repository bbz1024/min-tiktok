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

var once sync.Once
var authRpc authsclient.Auths

func WrapperAuthMiddleware(rpcConf zrpc.RpcClientConf) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// white path
			if slices.Index(WhitePath, r.URL.Path) != -1 {
				next(w, r)
				return
			}
			//fmt.Println(r.RequestURI) == r.URL.Path
			// get query param
			var token = r.PostFormValue("token")
			if r.Method == http.MethodGet {
				token = r.FormValue("token")
			}

			if token == "" {
				http.Error(w, "token is empty", http.StatusUnauthorized)
				return
			}

			// init rpc
			once.Do(func() {
				authRpc = authsclient.NewAuths(zrpc.MustNewClient(rpcConf))
			})
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
			/*
				actorID := res.UserId
				// write actorID to context  => r.Context().Value(keys.ActorID)
				ctx := context.WithValue(r.Context(), keys.ActorID, actorID)
				r = r.WithContext(ctx)
			*/
			r.Form.Set("actor_id", strconv.Itoa(int(res.UserId)))
			next(w, r)
		}
	}
}
