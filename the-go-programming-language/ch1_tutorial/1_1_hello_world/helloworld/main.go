// first program
package main

import "fmt"

// 在终端执行：
//
// 一、执行方式1
//
//  1. go build ./ch1_tutorial/1_1_hello_world/helloworld
//
//  2. ./helloworld
//
// 二、执行方式2（推荐。如果需要带参数则更推荐执行方式1）
//
//	go run ./ch1_tutorial/1_1_hello_world/helloworld/main.go
func main() {
	fmt.Println("hello world!")
}
