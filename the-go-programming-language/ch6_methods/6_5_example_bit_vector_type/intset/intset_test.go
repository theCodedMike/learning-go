package intset

import (
	"fmt"
	"testing"
)

func TestIntSet_One(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // {1 9 144}

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // {9 42}

	x.UnionWith(&y)
	fmt.Println(x.String())           // {1 9 42 144}
	fmt.Println(x.Has(9), x.Has(123)) // true false
}

func TestIntSet_Two(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // {1 9 42 144}
	fmt.Println(x.String()) // {1 9 42 144} // 编译器会隐式地在x前插入&操作符
	fmt.Println(x)          // {[4398046511618 0 65536]}
}
