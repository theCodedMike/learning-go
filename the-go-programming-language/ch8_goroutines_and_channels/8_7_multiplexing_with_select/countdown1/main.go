// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"time"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_7_multiplexing_with_select/countdown1/main.go
func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
