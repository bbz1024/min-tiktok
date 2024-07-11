package thread

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
	"testing"
	"time"
)

func TestThread(t *testing.T) {
	// x
	runner := threading.NewTaskRunner(10)
	start := time.Now()
	for i := 0; i < 10; i++ {
		runner.Schedule(func() {
			//time.Sleep(time.Second)
		})
	}
	runner.Wait()
	fmt.Println(start.Sub(time.Now()))
}
func TestThread2(t *testing.T) {
	// x
	start := time.Now()
	c := threading.NewStableRunner(func(i int) error {
		//time.Sleep(time.Second)
		fmt.Println(i)
		return nil
	})
	for i := 0; i < 10; i++ {
		c.Push(i)
	}
	c.Wait()

	fmt.Println(start.Sub(time.Now()))
}
func TestThread3(t *testing.T) {
	// x
	c := ""
	commentCnt, err := strconv.ParseInt(c, 10, 64)
	fmt.Println(commentCnt)
	fmt.Println(err)
}
