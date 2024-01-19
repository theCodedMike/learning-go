// 练习4.8：
// 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// 在终端执行：
//
//	go run ./ch4_composite_types/4_3_maps/exercise_4_8/main.go
func main() {
	counter := make(map[rune]int)
	in := bufio.NewReader(os.Stdin)
	numbers := 0
	chars := 0

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise4.8: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsLetter(r) {
			chars++
			counter[r]++
		}
		if unicode.IsNumber(r) {
			numbers++
			counter[r]++
		}
	}

	fmt.Println()
	for r, i := range counter {
		fmt.Printf("%q: %d\n", r, i)
	}
	fmt.Println("  Chars: ", chars)
	fmt.Println("Numbers: ", numbers)
}
