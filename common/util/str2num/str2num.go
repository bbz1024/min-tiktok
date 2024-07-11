package str2num

import (
	"strconv"
)

func Str2Num(s string) (int, error) {
	if s != "" {
		commentCnt, err := strconv.ParseInt(s, 10, 64)
		if err != nil {

			return 0, err
		}
		return int(commentCnt), nil
	}
	return 0, nil
}
