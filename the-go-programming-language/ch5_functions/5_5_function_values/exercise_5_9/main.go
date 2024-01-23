// 练习5.9：
// 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
package main

import (
	"bytes"
	"fmt"
)

// 在终端执行：
//
//	go run ./ch5_functions/5_5_function_values/exercise_5_9/main.go
func main() {
	s := "siffoowlifokfoo"
	var f = func(sub string) string {
		return "FUCK"
	}
	fmt.Println(expand(s, f))
}
func expand(s string, f func(string) string) string {
	if len(s) < 3 {
		return s
	}
	var bf bytes.Buffer
	for i, size := 0, len(s); i < size; {
		if i+3 <= size && s[i:i+3] == "foo" {
			bf.WriteString(f(""))
			i += 3
		} else {
			bf.WriteString(s[i : i+1])
			i++
		}
	}
	return bf.String()
}
