// 练习5.3：
// 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素，因为这些元素对浏览者是不可见的。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

// 在终端执行：
//  1. go build gopl.io/ch1_tutorial/1_5_fetching_a_url/fetch
//  2. go build gopl.io/ch5_functions/5_2_recursion/exercise_5_3
//  3. ./fetch https://baidu.com | ./exercise_5_3
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise_5_3: %v\n", err)
		os.Exit(1)
	}

	printText(doc)
}

func printText(n *html.Node) {
	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) != 0 {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(c)
	}
}
