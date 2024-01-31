// 练习7.1：
// 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常地有用。
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Counter struct {
	words int
	lines int
}

func (c *Counter) IncreaseLines() {
	c.lines++
}

func (c *Counter) Write(p []byte) (n int, err error) {
	c.words += len(bytes.Fields(p))
	return 0, nil
}

func (c *Counter) String() string {
	return fmt.Sprintf("Counter{words: %v, lines: %v}\n", c.words, c.lines)
}

// 在终端执行：
//
//	go run ./ch7_interfaces/7_1_interfaces_as_contracts/exercise_7_1/main.go
func main() {
	var counter Counter
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		content := scan.Bytes()
		counter.IncreaseLines()
		_, _ = counter.Write(content)
	}

	fmt.Println(&counter)
}
