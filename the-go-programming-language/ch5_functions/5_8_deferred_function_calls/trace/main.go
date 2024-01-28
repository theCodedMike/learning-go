// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"fmt"
	"log"
	"time"
)

const delay = 1

// 在终端执行：
//
//	go run ./ch5_functions/5_8_deferred_function_calls/trace/main.go
func main() {
	bigSlowOperation()
	// 2024/01/27 12:53:44 enter bigSlowOperation
	// 2024/01/27 12:53:54 exit bigSlowOperation (10.000253383s)

	fmt.Println(double(10))
	// double(10) = 20
	// 20

	fmt.Println(triple(4))
	// double(4) = 8
	// 12
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work
	time.Sleep(delay * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// defer语句中的函数会在return语句更新返回值变量后再执行
func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

// 被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值
func triple(x int) (result int) {
	defer func() {
		result += x
	}()
	return double(x)
}
