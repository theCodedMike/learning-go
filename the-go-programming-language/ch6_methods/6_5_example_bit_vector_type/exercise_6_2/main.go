// 练习6.2:
// 定义一个变参方法`(*IntSet).AddAll(...int)`，这个方法可以添加一组IntSet，比如s.AddAll(1, 2, 3)。
package main

import (
	"fmt"
	"gopl.io/ch6_methods/6_5_example_bit_vector_type/intset"
)

type IntSet = intset.IntSet

// 在终端执行：
//
//	go run ./ch6_methods/6_5_example_bit_vector_type/exercise_6_2/main.go
func main() {
	var x IntSet
	fmt.Printf("x: %v\n", x.String()) // {}, len: 0
	x.AddAll(1, 2, 3)
	fmt.Printf("add 1,2,3: %v\n\n", x.String()) // {1 2 3}, len: 3

	x.Remove(2)
	fmt.Printf("remove 2: %v\n", x.String()) // {1 3}, len: 2
}
