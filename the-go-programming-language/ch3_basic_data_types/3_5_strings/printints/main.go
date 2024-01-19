// Printints demonstrates the use of bytes.Buffer to format a string.
package main

import (
	"bytes"
	"fmt"
)

// 在终端执行：
//
//	go run ./ch3_basic_data_types/3_5_strings/printints/main.go
func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
	// [1, 2, 3]
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	return buf.String()
}
