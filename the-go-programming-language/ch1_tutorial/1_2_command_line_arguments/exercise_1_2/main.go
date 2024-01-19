// 练习1.2：
// 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)

// 在终端执行：
//  1. go build ./ch1_tutorial/1_2_command_line_arguments/exercise_1_2
//  2. ./exercise_1_2 hello i "love you"
//     或者
//  3. go run ./ch1_tutorial/1_2_command_line_arguments/exercise_1_1/main.go hello i "love you"
func main() {
	fmt.Println("idx    arg")
	for idx, arg := range os.Args[1:] {
		fmt.Printf("%2d:   %s\n", idx, arg)
	}

	// 输出如下：
	// idx    arg
	//  0:   hello
	//  1:   i
	//  2:   love you
}
