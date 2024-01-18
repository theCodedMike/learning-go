// 练习1.3：
// 左实验测试潜在低效的版本和使用了strings.Join的版本的运行时间差异。
// （1.6节讲解了部分`time`包，11.4节讲解了如何写标准测试程序，以得到系统性的性能评测。）
package main

import "fmt"

// 在终端执行：
//  1. go build ./ch1/1_2_command_line_arguments/exercise_1_3
//  2. ./exercise_1_3
//     或者
//  3. go run ./ch1/1_2_command_line_arguments/exercise_1_3/main.go
func main() {
	fmt.Println("hello, exercise 1.3")
}
