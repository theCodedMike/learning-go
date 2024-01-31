// 练习6.4:
// 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。
package main

import (
	"fmt"
	"gopl.io/ch6_methods/6_5_example_bit_vector_type/intset"
)

type IntSet = intset.IntSet

// 在终端执行：
//
//	go run ./ch6_methods/6_5_example_bit_vector_type/exercise_6_4/main.go
func main() {
	var x IntSet
	x.AddAll(19, 100, 38701, -278, 9, 10)

	fmt.Println(x.String()) // {9 10 19 100 38701}, len: 5
	for i, elem := range x.Elems() {
		fmt.Printf("%v: %v\n", i+1, elem)
	}
	// 1: 9
	// 2: 10
	// 3: 19
	// 4: 100
	// 5: 38701
}
