package gorse

import (
	"context"
	"github.com/zhenghaoz/gorse/client"
	"testing"
)

func TestConn(t *testing.T) {
	//http://127.0.0.1:8087
	cli := client.NewGorseClient("http://192.168.241.128:8088", "5105502fc46a411c896aa5b50c31e951")
	if _, err := cli.InsertItem(context.TODO(), client.Item{
		ItemId:  "1115",
		Comment: "hello",
	}); err != nil {
		t.Error(err)
		return
	}
}
