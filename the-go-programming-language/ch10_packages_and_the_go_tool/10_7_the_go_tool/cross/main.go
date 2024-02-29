package main

import (
	"fmt"
	"runtime"
)

// 在终端执行：
// go run ./ch10_packages_and_the_go_tool/10_7_the_go_tool/cross/main.go
// output: linux amd64
//
// 交叉编译时：
// GOOS=darwin GOARCH=amd64 go run ./ch10_packages_and_the_go_tool/10_7_the_go_tool/cross/main.go
// output: darwin amd64
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
