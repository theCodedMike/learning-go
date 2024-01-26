// 练习5.15：
// 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接受1个参数的版本。
package main

import (
	"fmt"
	"math"
)

// 在终端执行：
//
//	go run ./ch5_functions/5_7_variadic_functions/exercise_5_15/main.go
func main() {
	fmt.Println(max())                   // 9223372036854775807
	fmt.Println(max(1))                  // 1
	fmt.Println(max(1, -5))              // 1
	fmt.Println(max(19, 1999, -378, 93)) // 1999

	fmt.Println(min())                   // -9223372036854775808
	fmt.Println(min(1))                  // 1
	fmt.Println(min(1, -5))              // -5
	fmt.Println(min(19, 1999, -378, 93)) // -378
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return math.MaxInt
	}
	maxVal := math.MinInt

	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return math.MinInt
	}
	minVal := math.MaxInt

	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}

	return minVal
}
