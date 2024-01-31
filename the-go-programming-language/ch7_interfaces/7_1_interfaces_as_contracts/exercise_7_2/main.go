// 练习7.2：
// 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示
// 新的写入字节数的int64类型指针。
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type WordsCounter int

func (w *WordsCounter) Write(p []byte) (int, error) {
	*w += WordsCounter(len(bytes.Fields(p)))
	return len(p) * 8, nil
}

type Counter struct {
	words io.Writer
	size  int64
}

func (c *Counter) Write(p []byte) (int, error) {
	write, err := c.words.Write(p)
	if err != nil {
		return 0, err
	}
	c.size += int64(write)
	return 0, nil
}

func (c *Counter) String() string {
	return fmt.Sprintf("Counter{words: %v, size: %v}\n", int(*c.words.(*WordsCounter)), c.size)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := Counter{words: w, size: 0}
	return &counter, &counter.size
}

// 在终端执行：
//
//	go run ./ch7_interfaces/7_1_interfaces_as_contracts/exercise_7_2/main.go
func main() {
	var wCount WordsCounter
	writer, i := CountingWriter(&wCount)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ct := scanner.Bytes()
		_, err := writer.Write(ct)
		if err != nil {
			continue
		}
	}

	fmt.Println(writer, *i)
}
