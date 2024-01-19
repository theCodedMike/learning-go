// 练习5.1：
// 修改findlinks1代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 在终端执行：
//  1. go build gopl.io/ch1_tutorial/1_5_fetching_a_url/fetch
//  2. go build gopl.io/ch5_functions/5_2_recursion/exercise_5_1
//  3. ./fetch https://golang.org | ./exercise_5_1
//     ./fetch https://baidu.com | ./exercise_5_1
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise_5_1: %v\n", err)
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

	return recurVisit(links, n.FirstChild)
}

func recurVisit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	links = visit(links, n)
	return recurVisit(links, n.NextSibling)
}
