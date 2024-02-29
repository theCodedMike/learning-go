// 练习9.2：
// 重写2.6.2节中的PopCount的例子，使用sync.Once，只在第一次需要用到的时候进行初始化。
package main

import (
	"fmt"
	"sync"
)

// 在终端执行：
//
//	go run ./ch9_concurrency_with_shared_variables/9_5_lazy_initialization_sync_once/exercise_9_2/main.go
func main() {
	for i := uint64(0); i <= 10; i++ {
		fmt.Println(i, ":", PopCount(i))
	}
	// 0 : 0
	// 1 : 1
	// 2 : 1
	// 3 : 2
	// 4 : 1
	// 5 : 2
	// 6 : 2
	// 7 : 3
	// 8 : 1
	// 9 : 2
	// 10 : 2
}

var loadOnce sync.Once

// pc[i] is the population count of i.
var pc [256]byte

func load() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	loadOnce.Do(load)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
