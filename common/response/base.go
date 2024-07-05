package response

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
