package breaker

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/breaker"
	"math/rand"
	"testing"
	"time"
)

type mockError struct {
	status int
}

func (e mockError) Error() string {
	return fmt.Sprintf("HTTP STATUS: %d", e.status)
}

func mockRequest() error {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(100)
	if num%4 == 0 {
		return nil
	} else if num%5 == 0 {
		return mockError{status: 500}
	}
	return errors.New("dummy")
}
func TestBreaker(t *testing.T) {
	b := breaker.NewBreaker(breaker.WithName("test"))

	//b := breaker.NewBreaker(breaker.WithName("test"))
	for i := 0; i < 1000; i++ {
		err := b.DoWithAcceptable(func() error {
			r := rand.Intn(10)
			if r%2 == 0 {
				fmt.Println("success")
				return nil
			} else {
				return errors.New("dummy")
			}
		}, func(err error) bool {
			if err == nil {
				return false
			}
			fmt.Println(err)
			return true
		})
		fmt.Println(err, 11)
	}
}
