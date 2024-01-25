// 练习5.14：
// 使用breadthFirst遍历其他数据结构。比如，topoSort例子中的课程依赖关系（有向图）、个人计算机的文件层次结构（树）；
// 你所在城市的公交或地铁线路（无向图）。
package main

import (
	"container/list"
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// 在终端执行：
//
//	go run ./ch5_functions/5_6_anonymous_functions/exercise_5_14/main.go
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	queue := list.New()
	queue.PushBack(keys)
	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).([]string)
		for _, course := range front {
			if !seen[course] {
				seen[course] = true
				order = append(order, course)
				queue.PushBack(m[course])
			}
		}
	}

	return order
}
