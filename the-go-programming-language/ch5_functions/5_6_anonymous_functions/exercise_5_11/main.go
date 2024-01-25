// 练习5.11：
// 现在线性代数的老师把微积分设置为前置课程。完善topoSort，使其能检测有向图中的环。
package main

import "fmt"

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
	//"linear algebra":        {"calculus"},
}

// 在终端执行：
//
//	go run ./ch5_functions/5_6_anonymous_functions/exercise_5_11/main.go
func main() {
	fmt.Println(hasCycle(prereqs))
}

func hasCycle(m map[string][]string) bool {
	seen := make(map[string]bool)
	var visitAll func(string, map[string]bool) bool

	visitAll = func(item string, path map[string]bool) bool {
		// 如果当前节点已被访问过，则返回false
		if seen[item] {
			return false
		}
		// 将当前节点标记为已访问
		seen[item] = true
		path[item] = true
		// 遍历当前节点的所有相邻节点
		for _, next := range m[item] {
			// 如果相邻节点已在访问路径上，说明存在环，返回true
			if path[next] || visitAll(next, path) {
				return true
			}
		}
		// 将当前节点从访问路径中移除
		delete(path, item)
		return false
	}

	// 遍历map中的每个节点，并递归调用 visitAll
	for key := range m {
		path := make(map[string]bool)
		if visitAll(key, path) {
			return true
		}
	}

	// 如果所有节点都遍历完，则说明图中不存在环
	return false
}
