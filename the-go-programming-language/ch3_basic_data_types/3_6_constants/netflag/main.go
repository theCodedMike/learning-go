// Netflag demonstrates an integer type used as a bit field.
package main

import (
	"fmt"
	"net"
)

// 在终端执行：
//
//	go run ./ch3_basic_data_types/3_6_constants/netflag/main.go
func main() {
	v := net.FlagMulticast | net.FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // 10001 true

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // 10000 false

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // 10010 false
	fmt.Printf("%b %t\n", v, IsCast(v)) // 10010 true
}

func IsUp(v net.Flags) bool {
	return v&net.FlagUp == net.FlagUp
}

func TurnDown(v *net.Flags) {
	*v &^= net.FlagUp
}

func SetBroadcast(v *net.Flags) {
	*v |= net.FlagBroadcast
}

func IsCast(v net.Flags) bool {
	return v&(net.FlagBroadcast|net.FlagMulticast) != 0
}
