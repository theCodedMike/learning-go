// 练习4.6：
// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
package main

import (
	"fmt"
	"unicode"
)

// 在终端执行：
//
//	go run ./ch4/4_2_slices/exercise_4_6/main.go
func main() {
	s := "hello  world    haha"
	res := removeDupSpaces([]byte(s))
	fmt.Println(string(res)) // hello world haha

	s = "   h  e l l o  ,  w o r    l d."
	res = removeDupSpaces([]byte(s))
	fmt.Println(string(res)) //  h e l l o , w o r l d.
}

func removeDupSpaces(data []byte) []byte {
	if data == nil || len(data) == 0 {
		return data
	}

	k := -1
	space := 0
	for i, size := 0, len(data); i < size; i++ {
		if unicode.IsSpace(rune(data[i])) {
			space++
			if space == 1 {
				k++
			}
		} else {
			space = 0
			k++
		}

		if k != i {
			data[k] = data[i]
		}
	}

	return data[:k+1]
}
