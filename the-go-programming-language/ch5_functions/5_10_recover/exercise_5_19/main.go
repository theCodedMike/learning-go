// 练习5.19：
// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch5_functions/5_10_recover/exercise_5_19/main.go
func main() {
	fmt.Println(noReturn())
}

func noReturn() (result string) {
	defer func() {
		switch p := recover(); p {
		case "":
			result = "no panic"
		case "panic":
			result = "a panic occurred!"
		default:
			panic("panic in recover!")
		}
	}()

	panic("panic")
}
