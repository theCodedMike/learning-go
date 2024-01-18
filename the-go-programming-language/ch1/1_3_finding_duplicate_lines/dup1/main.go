// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//  1. fmt.Printf函数的格式化输出：
//
// / %d          十进制整数
// / %x, %o, %b  十六进制，八进制，二进制整数。
// / %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
// / %t          布尔：true或false
// / %c          字符（rune） (Unicode码点)
// / %s          字符串
// / %q          带双引号的字符串"abc"或带单引号的字符'c'
// / %v          变量的自然形式（natural format）
// / %T          变量的类型
// / %%          字面上的百分号标志（无操作数）
//
// 2. 按`Ctrl + D`结束输入
//
// 3. 在终端执行：
//  1. go build ./ch1/1_3_finding_duplicate_lines/dup1
//  2. ./dup1
//     或者
//  3. go run ./ch1/1_3_finding_duplicate_lines/dup1/main.go
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println("\n---出现次数大于1的结果如下---")

	// NOTE: ignoring potential errors from input.Err()
	for key, val := range counts {
		if val > 1 {
			fmt.Printf("%d\t%s\n", val, key)
		}
	}
}
