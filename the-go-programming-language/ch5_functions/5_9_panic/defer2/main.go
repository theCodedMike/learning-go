// Defer2 demonstrates a deferred call to runtime.Stack during a panic.
package main

import (
	"fmt"
	"os"
	"runtime"
)

// 在终端执行：
//
//	go run ./ch5_functions/5_9_panic/defer2/main.go
func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
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
// goroutine 1 [running]:
// main.printStack()
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:19 +0x2e
// panic({0x487240?, 0x51c750?})
//         /usr/local/go/src/runtime/panic.go:914 +0x21f
// main.f(0x4b4638?)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:24 +0x118
// main.f(0x1)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:26 +0xfe
// main.f(0x2)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:26 +0xfe
// main.f(0x3)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:26 +0xfe
// main.main()
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:14 +0x35
// panic: runtime error: integer divide by zero
//
// goroutine 1 [running]:
// main.f(0x4b4638?)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:24 +0x118
// main.f(0x1)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:26 +0xfe
// main.f(0x2)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:26 +0xfe
// main.f(0x3)
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:26 +0xfe
// main.main()
//         /home/Mike/Code/Go/learning-go/the-go-programming-language/ch5_functions/5_9_panic/defer2/main.go:14 +0x35
// exit status 2
