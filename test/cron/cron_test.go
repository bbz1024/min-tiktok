package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	schedule := cron.New()
	// 秒 分 时  日  月  周
	err := schedule.AddFunc("*/1 * * * *", func() {
		fmt.Println("222")
	})
	if err != nil {
		fmt.Println(err)
	}
	schedule.Start()
	time.Sleep(time.Second * 60)
}
