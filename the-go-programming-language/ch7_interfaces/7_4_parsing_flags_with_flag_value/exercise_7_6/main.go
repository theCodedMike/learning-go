// 练习7.6：
// 使tempFlag支持开尔文温度。
package main

import (
	"flag"
	"fmt"
	"gopl.io/ch7_interfaces/7_4_parsing_flags_with_flag_value/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 37.5, "the temperature")

// 在终端执行：
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/exercise_7_6/main.go
//
//	go run ./ch7_interfaces/7_4_parsing_flags_with_flag_value/tempflag/main.go -temp 273.15K
func main() {
	flag.Parse()
	fmt.Println(*temp)
}
