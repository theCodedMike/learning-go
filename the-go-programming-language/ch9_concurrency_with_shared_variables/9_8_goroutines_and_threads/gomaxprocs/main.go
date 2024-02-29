package main

import "fmt"

// 在终端执行：
//
// GOMAXPROCS=1 go run ./ch9_concurrency_with_shared_variables/9_8_goroutines_and_threads/gomaxprocs/main.go
// GOMAXPROCS=2 go run ./ch9_concurrency_with_shared_variables/9_8_goroutines_and_threads/gomaxprocs/main.go
func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
