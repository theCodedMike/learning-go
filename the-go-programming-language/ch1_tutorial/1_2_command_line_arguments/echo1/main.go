// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

// 1. Go语言只有for循环这一种循环语句
//
//	for initialization; condition; post {
//	    // zero or more statements
//	}
//
//	for condition {
//	   // a traditional "while" loop
//	}
//
//	for {
//	   // a traditional infinite loop
//	}
//
// 2. 在Go语言中，只有i++/i--是合法的，++i/--i是非法的
//
// 3. 符号:=是短变量声明（short variable declaration）的一部分
//
// 4. 在终端执行：
//  1. go build ./ch1_tutorial/1_2_command_line_arguments/echo1
//  2. ./echo1 hello i love you （注：这里相当于传递了4个参数，分别为hello、i、love和you）
//     或者
//  3. go run ./ch1_tutorial/1_2_command_line_arguments/echo1/main.go hello i love you
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
