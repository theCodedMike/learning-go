// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import (
	"fmt"
	"time"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_4_channels/pipeline2/main.go
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			fmt.Printf("Counter send x(%5d) to naturals...\n", x)
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			fmt.Printf("Squarer send x(%5d) to squares...\n", x*x)
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Printf("Printer recv x(%5d) from squares...\n\n", x)
	}
}
