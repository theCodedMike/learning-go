// 练习3.13：
// 编写KB、MB的常量声明，然后扩展到YB。
package main

import "fmt"

const (
	KiB = 1024 << (10 * iota) // 1024
	MiB                       // 1048576
	GiB                       // 1073741824
	TiB                       // 1099511627776             (exceeds 1 << 32)
	PiB                       // 1125899906842624
	EiB                       // 1152921504606846976
	ZiB                       // 1180591620717411303424    (exceeds 1 << 64)
	YiB                       // 1208925819614629174706176
)

// 在终端执行：
//
//	go run ./ch3/3_6_constants/exercise_3_13/main.go
func main() {
	fmt.Printf("KiB %d, type %T\n", KiB, KiB)
	fmt.Printf("MiB %d, type %T\n", MiB, MiB)
	fmt.Printf("GiB %d, type %T\n", GiB, GiB)
	fmt.Printf("TiB %d, type %T\n", TiB, TiB)
	fmt.Printf("PiB %d, type %T\n", PiB, PiB)
	fmt.Printf("EiB %d, type %T\n", EiB, EiB)
	//fmt.Printf("ZiB %d, type %T\n", ZiB, ZiB) // 无法表示ZiB
	//fmt.Printf("YiB %d, type %T\n", YiB, YiB) // 无法表示YiB

	// KiB 1024, type int
	// MiB 1048576, type int
	// GiB 1073741824, type int
	// TiB 1099511627776, type int
	// PiB 1125899906842624, type int
	// EiB 1152921504606846976, type int
}
