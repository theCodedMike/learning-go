package main

import (
	"fmt"
)

// 在终端执行：
//
//  1. go build ./ch6_methods/6_1_method_declarations/geometry
//  2. ./geometry
func main() {
	p := Point{X: 1, Y: 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q)) // 5
	fmt.Println(p.Distance(q))  // 5

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // 12
}
