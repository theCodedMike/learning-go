// 练习7.4：
// strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其他值）。实现一个简单版本的NewReader，
// 用它来构造一个接收字符串输入的HTML解析器（§5.2）
package main

import (
	"fmt"
	"strings"
)

// 在终端执行：
//
//	go run ./ch7_interfaces/7_2_interface_types/exercise_7_4/main.go
func main() {
	reader := strings.NewReader("hello")
	for i, size := 0, reader.Len(); i < size; i++ {
		fmt.Println(reader.ReadByte())
	}
	// 104 <nil>
	// 101 <nil>
	// 108 <nil>
	// 108 <nil>
	// 111 <nil>
}
