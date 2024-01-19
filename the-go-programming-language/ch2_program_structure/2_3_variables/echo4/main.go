// Echo4 prints its command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// 在终端执行：
//  1. go build ./ch2_program_structure/2_3_variables/echo4
//  2. ./echo4 a bc def
//     ./echo4 -s / a bc def
//     ./echo4 -n a bc def
//     ./echo4 -help
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
