// 练习5.2：
// 编写一个函数，统计HTML树中同类元素（p、div、span等）的个数。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 在终端执行：
//  1. go build gopl.io/ch1_tutorial/1_5_fetching_a_url/fetch
//  2. go build gopl.io/ch5_functions/5_2_recursion/exercise_5_2
//  3. ./fetch https://baidu.com | ./exercise_5_2
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise_5_2: %v\n", err)
		os.Exit(1)
	}

	for k, v := range visit(nil, doc) {
		fmt.Printf("%10s: %2d个\n", k, v)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(counter map[string]uint, n *html.Node) map[string]uint {
	if counter == nil {
		counter = make(map[string]uint)
	}

	if n.Type == html.ElementNode {
		counter["<"+n.Data+">"]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(counter, c)
	}

	return counter
}
