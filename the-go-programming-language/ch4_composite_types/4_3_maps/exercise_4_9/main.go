// 练习4.9：
// 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第1次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词
// 而不是按行输入。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// 在终端执行：
//
//	go run ./ch4_composite_types/4_3_maps/exercise_4_9/main.go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	wordFreq := make(map[string]int)

	for scanner.Scan() {
		word := strings.TrimFunc(scanner.Text(), func(r rune) bool {
			return unicode.IsPunct(r)
		})
		if len(word) > 0 {
			wordFreq[word]++
		}
	}

	for word, count := range wordFreq {
		fmt.Printf("%12s: %d\n", word, count)
	}
}
