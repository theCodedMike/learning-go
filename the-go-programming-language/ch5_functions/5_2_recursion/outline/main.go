// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 在终端执行：
//  1. go build gopl.io/ch1_tutorial/1_5_fetching_a_url/fetch
//  2. go build gopl.io/ch5_functions/5_2_recursion/outline
//  3. ./fetch https://golang.org | ./outline
//     ./fetch https://baidu.com | ./outline
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
