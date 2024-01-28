// Defer1 demonstrates a deferred call being invoked during a panic.
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch5_functions/5_9_panic/defer1/main.go
func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

// f(3)
// f(2)
// f(1)
// defer 1
// defer 2
// defer 3
// panic: runtime error: integer divide by zero
//
// goroutine 1 [running]:
// main.f(0x4b4568?)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer1/main.go:13 +0x118
// main.f(0x1)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer1/main.go:15 +0xfe
// main.f(0x2)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer1/main.go:15 +0xfe
// main.f(0x3)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer1/main.go:15 +0xfe
// main.main()
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer1/main.go:9 +0x18
// exit status 2
