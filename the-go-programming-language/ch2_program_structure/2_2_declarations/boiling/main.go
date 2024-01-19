// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

// 在终端执行：
//
//	go run ./ch2_program_structure/2_2_declarations/boiling/main.go
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	// boiling point = 212°F or 100°C
}
