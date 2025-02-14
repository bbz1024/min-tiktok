package errorcode

import "github.com/zeromicro/go-zero/core/logx"

func init() {
	logx.Info("errorcode init finished")
}
func ccc() {
	a := 6 / 0
	if a != 1 {
		logx.Error("error")
	}
}
