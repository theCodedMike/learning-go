// 练习6.1：
// 为bit数组实现下面的这些方法
// ```
// func (*IntSet) Len() int      // return the number of elements
// func (*IntSet) Remove(x int)  // remove x from the set
// func (*IntSet) Clear()        // remove all elements from the set
// func (*IntSet) Copy() *IntSet // return a copy of the set
// ```
package main

import (
	"fmt"
	"gopl.io/ch6_methods/6_5_example_bit_vector_type/intset"
)

type IntSet = intset.IntSet

// 在终端执行：
//
//	go run ./ch6_methods/6_5_example_bit_vector_type/exercise_6_1/main.go
func main() {
	var x IntSet
	x.Add(1)
	fmt.Printf("add 1: %v\n", x.String()) // {1}, len: 1
	x.Add(10)
	fmt.Printf("add 10: %v\n", x.String()) // {1 10}, len: 2
	x.Add(-10)
	fmt.Printf("add -10: %v\n", x.String()) // {1 10}, len: 2
	x.Add(192)
	fmt.Printf("add 192: %v\n", x.String()) // {1 10 192}, len: 3
	x.Add(3)
	fmt.Printf("add 3: %v\n", x.String()) // {1 3 10 192}, len: 4
	fmt.Printf("now len: %v\n", x.Len())  // 4
	x.Add(1024)
	x.Add(111)
	x.Add(10)
	fmt.Printf("add 1024,111,10: %v\n", x.String()) // {1 3 10 111 192 1024}, len: 6
	fmt.Printf("now len: %v\n\n", x.Len())          // 6

	fmt.Printf("has 10: %v\n", x.Has(10))   // true
	fmt.Printf("has -10: %v\n", x.Has(-10)) // false
	fmt.Printf("has 20: %v\n\n", x.Has(20)) // false

	x.Remove(-10)
	fmt.Printf("remove -10: %v\n", x.String()) // {1 3 10 111 192 1024}, len: 6
	fmt.Printf("has -10: %v\n", x.Has(-10))    // false
	x.Remove(10)
	fmt.Printf("remove 10: %v\n", x.String()) // {1 3 111 192 1024}, len: 5
	fmt.Printf("has 10: %v\n", x.Has(10))     // false
	x.Remove(1)
	fmt.Printf("remove 1: %v\n\n", x.String()) // {3 111 192 1024}, len: 4

	cp := x.Copy()
	cp.Add(333)
	fmt.Printf("cp: %v\n", cp.String())  // {3 111 192 333 1024}, len: 5
	fmt.Printf(" x: %v\n\n", x.String()) // {3 111 192 1024}, len: 4

	x.Clear()
	fmt.Printf("clear: %v\n", x.String())     // {}, len: 0
	fmt.Printf("has 111: %v\n\n", x.Has(111)) // false

	x.Add(101)
	fmt.Printf("add 101: %v\n", x.String()) // {101}, len: 1
	x.Add(102)
	fmt.Printf("add 102: %v\n", x.String()) // {101 102}, len: 2
	x.Add(103)
	fmt.Printf("add 103: %v\n", x.String())   // {101 102 103}, len: 3
	fmt.Printf("has 103: %v\n", x.Has(103))   // true
	fmt.Printf("has 104: %v\n\n", x.Has(104)) // false

	x.Remove(104)
	fmt.Printf("remove 104: %v\n", x.String()) // {101 102 103}, len: 3
	x.Remove(101)
	fmt.Printf("remove 101: %v\n", x.String()) // {102 103}, len: 2
	fmt.Printf("has 101: %v\n", x.Has(101))    // false
}
