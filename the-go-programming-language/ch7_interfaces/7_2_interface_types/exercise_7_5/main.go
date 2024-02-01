// 练习7.5：
// io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。
// 实现这个LimitReader函数：
package main

import (
	"io"
)

// 在终端执行：
//
//	go run ./ch7_interfaces/7_2_interface_types/exercise_7_5/main.go
func main() {
	LimitReader(nil, 10)
}

func LimitReader(r io.Reader, n int64) io.Reader {
	//todo!
	read, err := r.Read([]byte("hello"))
	if err != nil {
		return nil
	}
	if int64(read) > n {
		return nil
	}
	return nil
}
