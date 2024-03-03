// Package unsafeptr demonstrates basic use of unsafe.Pointer.
package main

import (
	"fmt"
	"unsafe"
)

// 在终端执行：
//
// go run ./ch13_low_level_programming/13_2_unsafe_pointer/unsafeptr/main.go
func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// 和 pb := &x.b等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	// NOTE: subtly incorrect!
	//tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	//pb := (*int16)(unsafe.Pointer(tmp))
	//*pb = 42

	fmt.Println(x.b) // 42
}
