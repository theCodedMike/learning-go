// 练习1.1：
// 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
package main

import (
	"fmt"
	"os"
)

// 在终端执行：
//  1. go build ./ch1_tutorial/1_2_command_line_arguments/exercise_1_1
//  2. ./exercise_1_1 hello i "love you"
//     或者
//  3. go run ./ch1_tutorial/1_2_command_line_arguments/exercise_1_1/main.go hello i love you
func main() {
	// for _, arg := range os.Args { // 这样也可以
	for _, arg := range os.Args[0:] {
		fmt.Println(arg)
	}

	// 输出结果如下：
	// ./exercise_1_1
	// hello
	// i
	// love you
}
