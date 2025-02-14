package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err.Error())
	}

	// 文件操作
	fmt.Println("File opened")

	// 忘记关闭文件
}
