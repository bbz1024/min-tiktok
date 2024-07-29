package limiter

import (
	"github.com/zeromicro/go-zero/core/syncx"
	"sync"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	l := syncx.NewLimit(10)
	var wg sync.WaitGroup
	wg.Add(50)

	for i := 0; i < 50; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			if l.TryBorrow() {
				time.Sleep(time.Millisecond * 10)
				t.Log("request success")
				l.Return()
				return
			} else {
				t.Log("request fail")
				return
			}
		}()
	}
	wg.Wait()
}
