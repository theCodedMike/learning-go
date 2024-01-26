// 练习5.16：
// 编写多参数版本的strings.Join。
package main

import (
	"bytes"
	"fmt"
)

// 在终端执行：
//
//	go run ./ch5_functions/5_7_variadic_functions/exercise_5_16/main.go
func main() {
	fmt.Println(join(","))                                         //
	fmt.Println(join(",", "a"))                                    // a
	fmt.Println(join("+", "a", "zd", "-2"))                        // a+zd+-2
	fmt.Println(join("***", "i", "love", "u", "it's", "a", "lie")) // i***love***u***it's***a***lie

	chars := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(join("|", chars...)) // a|b|c|d|e|f
}

func join(sep string, vals ...string) string {
	if len(vals) == 0 {
		return ""
	}

	var s bytes.Buffer
	for i, val := range vals {
		if i > 0 {
			s.WriteString(sep)
		}
		s.WriteString(val)
	}

	return s.String()
}
