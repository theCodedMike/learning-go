// 练习5.4：
// 扩展函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 在终端执行：
//  1. go build gopl.io/ch1_tutorial/1_5_fetching_a_url/fetch
//  2. go build gopl.io/ch5_functions/5_2_recursion/exercise_5_4
//  3. ./fetch https://baidu.com | ./exercise_5_4
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise_5_4: %v\n", err)
		os.Exit(1)
	}

	for k, vals := range visit(nil, doc) {
		fmt.Printf("%s:\n", k)
		for i, val := range vals {
			fmt.Printf("  %2d: %s\n", i, val)
		}
	}
}

// visit appends to links each link found in n and returns the result.
func visit(container map[string][]string, n *html.Node) map[string][]string {
	if container == nil {
		container = make(map[string][]string)
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" {
			links := container["links"]
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
			container["links"] = links

		} else if n.Data == "img" {
			imgs := container["images"]
			for _, a := range n.Attr {
				if a.Key == "src" {
					imgs = append(imgs, a.Val)
				}
			}
			container["images"] = imgs
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		container = visit(container, c)
	}

	return container
}
