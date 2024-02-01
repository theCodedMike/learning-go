// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"
	"gopl.io/ch7_interfaces/7_4_parsing_flags_with_flag_value/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

// 在终端执行：
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/tempflag/main.go
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/tempflag/main.go -temp -18C
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/tempflag/main.go -temp 212°F
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/tempflag/main.go -temp 273.15K
func main() {
	flag.Parse()
	fmt.Println(*temp)
}
