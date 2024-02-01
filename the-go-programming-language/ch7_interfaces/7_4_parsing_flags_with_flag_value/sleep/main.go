// The sleep program sleeps for a specified period of time.
package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

// 在终端执行：
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/sleep/main.go
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/sleep/main.go -period 10s
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/sleep/main.go -period 1.5h10m30s
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
