// 练习3.11：
// 完善comma函数，以支持浮点数处理和一个可选的正负号处理。
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// 在终端执行：
//  1. go build ./ch3_basic_data_types/3_5_strings/exercise_3_11
//  2. ./exercise_3_11 1 12 +123 -1234 -123456 12345678 3.14 31.4 314.15926 3141.5926 +31415.92657 -31415926.57321 -0.5536 -.35239 +3927332.3242
func main() {
	for _, s := range os.Args[1:] {
		fmt.Printf("%20s : %20s\n", s, comma(s))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	var intPartSize int

	if dot := strings.IndexByte(s, '.'); dot != -1 {
		intPartSize = dot
	} else {
		intPartSize = len(s)
	}

	if s[0] == '-' || s[0] == '+' {
		intPartSize--
	}

	rem := intPartSize % 3
	var idx = -1
	var meetDot = false
	for _, ch := range s {
		if ch == '.' {
			meetDot = true
		}
		if ch != '-' && ch != '+' && !meetDot {
			idx++
			if idx > 0 && idx%3 == rem {
				buf.WriteByte(',')
			}
		}
		buf.WriteRune(ch)
	}

	return buf.String()
}
