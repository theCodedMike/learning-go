// The urlvalues command demonstrates a map type with methods.
package main

import (
	"fmt"
	"net/url"
)

// 在终端执行：
//
//	go run ./ch6_methods/6_2_methods_with_a_pointer_receiver/urlvalues/main.go
func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // en
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // 1
	fmt.Println(m["item"])     // [1 2]

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map
}
