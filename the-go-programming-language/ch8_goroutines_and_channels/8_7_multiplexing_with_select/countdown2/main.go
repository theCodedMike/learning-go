// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"os"
	"time"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_7_multiplexing_with_select/countdown2/main.go
func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	go func() {
		for countdown := 10; countdown > 0; countdown-- {
			fmt.Println(countdown)
			time.Sleep(1 * time.Second)
		}
	}()
	select {
	case <-time.After(10 * time.Second): // Do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
