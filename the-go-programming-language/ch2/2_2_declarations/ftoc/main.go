// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch2/2_2_declarations/ftoc/main.go
func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	// 32°F = 0°C
	// 212°F = 100°C
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
