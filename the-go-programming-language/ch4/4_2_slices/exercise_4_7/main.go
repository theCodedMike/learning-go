// 练习4.7：
// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch4/4_2_slices/exercise_4_7/main.go
func main() {
	s := "hello"
	bytes := []byte(s)
	reverse(bytes)
	fmt.Println(string(bytes)) // olleh
}

func reverse(b []byte) {
	if b == nil || len(b) == 0 {
		return
	}

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
