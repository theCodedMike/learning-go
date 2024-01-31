// 练习7.3：
// 为在gopl.io/ch4_composite_types/4_4_structs/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const indent = 4

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func (t *TreeNode) String() string {
	var buf bytes.Buffer
	depth := 0

	var writeBuf = func(val string) {
		if depth > 1 {
			if depth > 2 {
				_, _ = fmt.Fprintf(&buf, "%s", strings.Repeat(" ", (depth-2)*indent))
			}
			buf.WriteString("|—— ")
		}
		_, _ = fmt.Fprintf(&buf, "%v\n", val)
	}

	var dfs func(n *TreeNode, parentIsLeaf bool)
	dfs = func(n *TreeNode, parentIsLeaf bool) {
		depth++
		if n == nil {
			if !parentIsLeaf {
				writeBuf("nil")
			}
		} else {
			writeBuf(strconv.Itoa(n.value))

			parentIsLeaf = n.left == nil && n.right == nil
			dfs(n.left, parentIsLeaf)
			dfs(n.right, parentIsLeaf)
		}
		depth--
	}

	dfs(t, false)
	return buf.String()
}

func (t *TreeNode) RecurPreOrder() []int {
	var vals []int
	var preOrder func(n *TreeNode, vals *[]int)
	preOrder = func(n *TreeNode, vals *[]int) {
		if n != nil {
			*vals = append(*vals, n.value)
			preOrder(n.left, vals)
			preOrder(n.right, vals)
		}
	}

	preOrder(t, &vals)
	return vals
}

func (t *TreeNode) IterPreOrder() []int {
	var vals []int
	if t != nil {
		var stack []*TreeNode
		stack = append(stack, t)
		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			vals = append(vals, curr.value)
			if curr.right != nil {
				stack = append(stack, curr.right)
			}
			if curr.left != nil {
				stack = append(stack, curr.left)
			}
		}
	}
	return vals
}

func (t *TreeNode) RecurInOrder() []int {
	var vals []int
	var inOrder func(n *TreeNode, vals *[]int)
	inOrder = func(n *TreeNode, vals *[]int) {
		if n != nil {
			inOrder(n.left, vals)
			*vals = append(*vals, n.value)
			inOrder(n.right, vals)
		}
	}

	inOrder(t, &vals)
	return vals
}

func (t *TreeNode) IterInOrder() []int {
	var vals []int
	if t != nil {
		var stack []*TreeNode
		var curr = t

		for len(stack) > 0 || curr != nil {
			for curr != nil {
				stack = append(stack, curr)
				curr = curr.left
			}

			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vals = append(vals, pop.value)
			curr = pop.right
		}
	}
	return vals
}

func (t *TreeNode) RecurPostOrder() []int {
	var vals []int
	var postOrder func(n *TreeNode, vals *[]int)
	postOrder = func(n *TreeNode, vals *[]int) {
		if n != nil {
			postOrder(n.left, vals)
			postOrder(n.right, vals)
			*vals = append(*vals, n.value)
		}
	}

	postOrder(t, &vals)
	return vals
}

func (t *TreeNode) IterPostOrder() []int {
	var vals []int
	if t != nil {
		var stack []any
		stack = append(stack, t)

		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch curr.(type) {
			case *TreeNode:
				{
					node := curr.(*TreeNode)
					stack = append(stack, node.value)
					if node.right != nil {
						stack = append(stack, node.right)
					}
					if node.left != nil {
						stack = append(stack, node.left)
					}
				}
			case int:
				{
					vals = append(vals, curr.(int))
				}
			}
		}
	}
	return vals
}

func (t *TreeNode) LevelOrder() []int {
	var vals []int
	if t != nil {
		var queue []*TreeNode
		queue = append(queue, t)

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			vals = append(vals, curr.value)
			if curr.left != nil {
				queue = append(queue, curr.left)
			}
			if curr.right != nil {
				queue = append(queue, curr.right)
			}
		}
	}
	return vals
}

func Build(elems ...int) *TreeNode {
	var root *TreeNode
	for _, elem := range elems {
		root = add(root, elem)
	}
	return root
}

func add(t *TreeNode, elem int) *TreeNode {
	if t == nil {
		t = new(TreeNode)
		t.value = elem
		return t
	}

	if elem < t.value {
		t.left = add(t.left, elem)
	} else {
		t.right = add(t.right, elem)
	}

	return t
}

// 在终端执行：
//
//	go run ./ch7_interfaces/7_1_interfaces_as_contracts/exercise_7_3/main.go
func main() {
	//     10
	//    /  \
	//   4    16
	//  / \    \
	// 1   9    20
	//    /      \
	//   7        50
	//           /  \
	//          30   100
	root := Build(10, 4, 16, 1, 9, 20, 7, 50, 30, 100)
	fmt.Printf("递归先序: %v\n", root.RecurPreOrder())  // [10 4 1 9 7 16 20 50 30 100]
	fmt.Printf("迭代先序: %v\n", root.IterPreOrder())   // [10 4 1 9 7 16 20 50 30 100]
	fmt.Printf("递归中序: %v\n", root.RecurInOrder())   // [1 4 7 9 10 16 20 30 50 100]
	fmt.Printf("迭代中序: %v\n", root.IterInOrder())    // [1 4 7 9 10 16 20 30 50 100]
	fmt.Printf("递归后序: %v\n", root.RecurPostOrder()) // [1 7 9 4 30 100 50 20 16 10]
	fmt.Printf("迭代后序: %v\n", root.IterPostOrder())  // [1 7 9 4 30 100 50 20 16 10]
	fmt.Printf("层序遍历: %v\n", root.LevelOrder())     // [10 4 16 1 9 20 7 50 30 100]
	//10
	//|—— 4
	//    |—— 1
	//    |—— 9
	//        |—— 7
	//        |—— nil
	//|—— 16
	//    |—— nil
	//    |—— 20
	//        |—— nil
	//        |—— 50
	//            |—— 30
	//            |—— 100
	fmt.Println(root)
}
