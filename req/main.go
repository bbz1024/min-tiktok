package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func get() {
	url := "http://nginx.kubernetes-devops.cn:31672/douyin/user/login/?username=123456&password=123456"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "PostmanRuntime-ApipostRuntime/1.1.0")
	req.Header.Add("Connection", "keep-alive")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			get()
		}()
	}
	wg.Wait()
	fmt.Println("done")

}
