// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import (
	"fmt"
	"time"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_4_channels/pipeline1/main.go
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			fmt.Printf("Counter send x(%d) to naturals...\n", x)
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			fmt.Printf("Squarer send x(%d) to squares...\n", x*x)
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Printf("Printer recv x(%d) from squares...\n", <-squares)
	}
}
