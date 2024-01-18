package main

import (
	"fmt"
	"os"
	"strings"
)

// 在终端执行：
//  1. go build ./ch1/1_2_command_line_arguments/echo3
//  2. ./echo3 hello i love you
//     或者
//  3. go run ./ch1/1_2_command_line_arguments/echo3/main.go hello i love you
func main() {
	fmt.Println(strings.Join(os.Args[1:], ","))
}
