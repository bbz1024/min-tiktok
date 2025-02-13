package main

import (
	"fmt"
	"math"
)

func main() {
	var maxInt32 int32 = math.MaxInt32
	fmt.Println("Max int32:", maxInt32)
	fmt.Println(maxInt32 + 1)

	var a float64 = 0.1
	var b float64 = 0.2
	fmt.Println("0.1 + 0.2 =", a+b)

	var ptr *int
	fmt.Println(*ptr) // 运行时错误

	arr := [3]int{1, 2, 3}
	fmt.Println(arr[3]) // 编译错误或运行时错误

	infiniteRecursion()
}

func infiniteRecursion() {
	infiniteRecursion()
}
