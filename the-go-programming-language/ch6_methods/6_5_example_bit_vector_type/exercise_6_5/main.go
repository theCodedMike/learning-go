// 练习6.5:
// 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
// 修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，
// 这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
package main

import (
	"fmt"
	"gopl.io/ch6_methods/6_5_example_bit_vector_type/intset"
)

type UintSet = intset.UintSet

// 在终端执行：
//
//	go run ./ch6_methods/6_5_example_bit_vector_type/exercise_6_5/main.go
func main() {
	var x UintSet
	fmt.Println(x.Has(10)) // false

	x.Add(1902)
	//x.Add(-10) // can't add, panic
	x.Add(189)
	x.AddAll(10, 15, 78, 90, 140, 189, 150)
	fmt.Println(x)                    // {10 15 78 90 140 150 189 1902}, len: 8
	fmt.Println(x.Len())              // 8
	fmt.Println(x.Has(15), x.Has(11)) // true false
	fmt.Println()

	x.Remove(78)
	x.Remove(12)
	for i, elem := range x.Elems() {
		fmt.Printf("%v: %v\n", i+1, elem)
	}
	//1: 10
	//2: 15
	//3: 90
	//4: 140
	//5: 150
	//6: 189
	//7: 1902
	fmt.Println(x.Has(78), x.Has(12)) // false false
	fmt.Println(x.Len())              // 7
	fmt.Println()

	// todo!
}
