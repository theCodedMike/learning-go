// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

//  1. 声明一个变量有好几种方式，下面这些都等价
//     s := ""（只能用在函数内部，而不能用于包变量）
//     var s string （依赖于字符串的默认初始化零值机制，被初始化为""）
//     var s = ""
//     var s string = ""
//
// 2. 在终端执行：
//  1. go build ./ch1/1_2_command_line_arguments/echo2
//  2. ./echo2 hello i love you
//     或者
//  3. go run ./ch1/1_2_command_line_arguments/echo2/main.go hello i love you
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // 这里的_省略的是索引
		fmt.Println(arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
