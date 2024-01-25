package main

import "fmt"

// 在终端执行：
//
//	go run ./ch5_functions/5_6_anonymous_functions/squares/main.go
func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // 25
}

// squares 返回一个匿名函数，该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
