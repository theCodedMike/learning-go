// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"gopl.io/ch2_program_structure/2_6_packages_and_files/tempconv"
	"os"
	"strconv"
)

// 在终端执行：
//  1. go build ./ch2_program_structure/2_6_packages_and_files/cf
//  2. ./cf 100 200
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

	// $ ./cf 100 200
	// 100°F = 37.77777777777778°C, 100°C = 212°F
	// 200°F = 93.33333333333333°C, 200°C = 392°F
}
