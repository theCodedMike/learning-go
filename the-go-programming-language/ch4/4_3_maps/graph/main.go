// Graph shows how to use a map of maps to represent a directed graph.
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch4/4_3_maps/graph/main.go
func main() {
	addEdge("a", "b")
	addEdge("c", "d")
	addEdge("a", "d")
	addEdge("d", "a")
	fmt.Println(hasEdge("a", "b")) // true
	fmt.Println(hasEdge("c", "d")) // true
	fmt.Println(hasEdge("a", "d")) // true
	fmt.Println(hasEdge("d", "a")) // true
	fmt.Println(hasEdge("x", "b")) // false
	fmt.Println(hasEdge("c", "d")) // true
	fmt.Println(hasEdge("x", "d")) // false
	fmt.Println(hasEdge("d", "x")) // false
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
