// 练习5.12：
// gopl.io/ch5/outline2(5.5节)的startelement和endElement共用了全局变量depth，将它们修改为匿名函数，使其共享outline中的局部变量。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

const indent = 4

// 在终端执行：
//
//	go run ./ch5_functions/5_6_anonymous_functions/exercise_5_12/main.go https://baidu.com
func main() {
	for _, url := range os.Args[1:] {
		if outline(url) != nil {
			continue
		}
	}
}
func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var depth int

	startElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*indent, "", node.Data)
			depth++
		}
	}

	endElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*indent, "", node.Data)
		}
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
