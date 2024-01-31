// Byte counter demonstrates an implementation of io.Writer that counts bytes.
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// 在终端执行：
//
//	go run ./ch7_interfaces/7_1_interfaces_as_contracts/bytecounter/main.go
func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12
}
