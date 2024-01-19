// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 在终端执行：
//  1. go build gopl.io/ch1_tutorial/1_5_fetching_a_url/fetch
//  2. go build gopl.io/ch5_functions/5_2_recursion/findlinks1
//  3. ./fetch https://golang.org | ./findlinks1
//     ./fetch https://baidu.com | ./findlinks1
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

// package html
//
// type Node struct {
//     Type                    NodeType
//     Data                    string
//     Attr                    []Attribute
//     FirstChild, NextSibling *Node
// }
//
// type NodeType int32
//
// const (
//     ErrorNode NodeType = iota
//     TextNode
//     DocumentNode
//     ElementNode
//     CommentNode
//     DoctypeNode
// )
//
// type Attribute struct {
//     Key, Val string
// }
//
// func Parse(r io.Reader) (*Node, error)
