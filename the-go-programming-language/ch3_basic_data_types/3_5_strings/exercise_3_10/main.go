// 练习3.10：
// 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
package main

import (
	"bytes"
	"fmt"
	"os"
)

// 在终端执行：
//  1. go build ./ch3_basic_data_types/3_5_strings/exercise_3_10
//  2. ./exercise_3_10 1 12 123 1234 123456 12345678
func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}
	// 1
	// 12
	// 123
	// 1,234
	// 123,456
	// 12,345,678
}

func comma(s string) string {
	var buf bytes.Buffer
	rem := len(s) % 3

	for i, c := range s {
		if i > 0 && i%3 == rem {
			buf.WriteByte(',')
		}
		buf.WriteRune(c)
	}

	return buf.String()
}
