// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResp struct {
	Token      string `json:"token"`
	UserID     uint32 `json:"user_id"`
	StatusCode uint32 `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type RegisterReq struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterResp struct {
	Token      string `json:"token"`
	UserID     uint32 `json:"user_id"`
	StatusCode uint32 `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
