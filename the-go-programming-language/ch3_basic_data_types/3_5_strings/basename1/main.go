// Basename1 reads file names from stdin and prints the base name of each one.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 在终端执行：
//
//	go run ./ch3_basic_data_types/3_5_strings/basename1/main.go
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
