// 练习6.3:
// ` (*IntSet).UnionWith`会用`|`操作符计算2个集合的并集，我们再为IntSet实现另外几个函数：
// IntersectWith（交集：元素在A集合B集合均出现）
// DifferenceWith（差集：元素出现在A集合，未出现在B集合）
// SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。
package main

import (
	"fmt"
	"gopl.io/ch6_methods/6_5_example_bit_vector_type/intset"
)

type IntSet = intset.IntSet

// 在终端执行：
//
//	go run ./ch6_methods/6_5_example_bit_vector_type/exercise_6_3/main.go
func main() {
	var x IntSet
	x.AddAll(1, 2, 3)
	var y IntSet
	y.AddAll(3, 4, 5)

	fmt.Println("============ Test IntersectWith ============")
	fmt.Printf("x: %v\n", x.String()) // x: {1 2 3}, len: 3
	fmt.Printf("y: %v\n", y.String()) // y: {3 4 5}, len: 3
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {3}, len: 1

	x.AddAll(4, 5, 100, -11, 33, 2, 3, 3)
	fmt.Printf("x: %v\n", x.String()) // x: {2 3 4 5 33 100}, len: 6
	fmt.Printf("y: %v\n", y.String()) // y: {3 4 5}, len: 3
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {3 4 5}, len: 3

	fmt.Printf("x: %v\n", x.String()) // x: {3 4 5}, len: 3
	fmt.Printf("y: %v\n", y.String()) // y: {3 4 5}, len: 3
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {3 4 5}, len: 3

	y.Clear()
	fmt.Printf("x: %v\n", x.String()) // x: {3 4 5}, len: 3
	fmt.Printf("y: %v\n", y.String()) // y: {}, len: 0
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {}, len: 0

	y.AddAll(10, 11, 10, 12)
	fmt.Printf("x: %v\n", x.String()) // x: {}, len: 0
	fmt.Printf("y: %v\n", y.String()) // y: {10 11 12}, len: 3
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {}, len: 0

	x.AddAll(100, 1000, 10038, 183023)
	fmt.Printf("x: %v\n", x.String()) // x: {100 1000 10038 183023}, len: 4
	fmt.Printf("y: %v\n", y.String()) // y: {10 11 12}, len: 3
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {}, len: 0

	x.AddAll(10, 13, 12, 17, 2)
	fmt.Printf("x: %v\n", x.String()) // x: {2 10 12 13 17}, len: 5
	fmt.Printf("y: %v\n", y.String()) // y: {10 11 12}, len: 3
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {10 12}, len: 2

	x.AddAll(65, 83, 190)
	y.AddAll(83, 190, 164, 5299)
	fmt.Printf("x: %v\n", x.String()) // x: {10 12 65 83 190}, len: 5
	fmt.Printf("y: %v\n", y.String()) // y: {10 11 12 83 164 190 5299}, len: 7
	x.IntersectWith(&y)
	fmt.Printf("after IntersectWith y, x: %v\n\n", x.String()) // x: {10 12 83 190}, len: 4

	fmt.Println("============ Test DifferenceWith ============")
	x.Add(11)
	y.Remove(11)
	fmt.Printf("x: %v\n", x.String()) // x: {10 11 12 83 190}, len: 5
	fmt.Printf("y: %v\n", y.String()) // y: {10 12 83 164 190 5299}, len: 6
	x.DifferenceWith(&y)
	fmt.Printf("after DifferenceWith y, x: %v\n\n", x.String()) // x: {11 164 5299}, len: 3

	x.AddAll(89, 190, 1000, 10, 12)
	fmt.Printf("x: %v\n", x.String()) // x: {10 11 12 89 190 1000}, len: 6
	fmt.Printf("y: %v\n", y.String()) // y: {10 12}, len: 2
	x.DifferenceWith(&y)
	fmt.Printf("after DifferenceWith y, x: %v\n\n", x.String()) // x: {11 89 190 1000}, len: 4

	y.AddAll(89, 190, 1000)
	fmt.Printf("x: %v\n", x.String()) // x: {11 89 190 1000}, len: 4
	fmt.Printf("y: %v\n", y.String()) // y: {10 12 89, 190, 1000}, len: 5
	x.DifferenceWith(&y)
	fmt.Printf("after DifferenceWith y, x: %v\n\n", x.String()) // x: {10 11 12}, len: 3
}
