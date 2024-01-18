// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 在终端执行：
//  1. go build ./ch1/1_3_finding_duplicate_lines/dup2
//  2. 输入`./dup2`可以测试在终端的输入
//  3. 输入`./dup2 ./test.txt`可以测试本项目中的文件
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil { // 如果err的值不是nil，说明打开文件时出错了
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	fmt.Println("\n---出现次数大于1的结果如下---")
	for key, val := range counts {
		if val > 1 {
			fmt.Printf("%d\t%s\n", val, key)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
