// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"fmt"
	"os"
)

// 在终端执行：
//  1. go build ./ch3_basic_data_types/3_5_strings/comma
//  2. ./comma 1 12 123 1234 123456 12345678
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

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
