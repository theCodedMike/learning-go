// 练习5.7：
// 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（<a href='...'>）。
// 使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。编写测试，验证程序输出的格式是否正确。（详见第11章）
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

// 在终端执行：
//
//	go run ./ch5_functions/5_5_function_values/exercise_5_7/main.go  https://baidu.com
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*4, "", n.Data)
		for _, attr := range n.Attr {
			key := attr.Key
			val := attr.Val
			if key == "class" {
				if len(strings.Fields(val)) > 1 {
					val = "\"" + val + "\""
				}
			}
			fmt.Printf(" %s=%s", key, val)
		}
		fmt.Print(">")
		if _, newLine := printNewLine(n); newLine {
			fmt.Println()
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("<!--%s-->\n", n.Data)
	} else if n.Type == html.DoctypeNode {
		fmt.Printf("<!DOCTYPE %s>\n", n.Data)
	} else {
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			if n.Parent != nil && (n.Parent.Data == "script" || n.Parent.Data == "style") {
				fmt.Println(s)
			} else if n.Parent != nil && (n.Parent.Data == "a" || n.Parent.Data == "title" || n.Parent.Data == "h1") {
				fmt.Print(s)
			} else {
				fmt.Printf("%*s%s\n", depth*4, "", s)
			}
		}
	}

	if n.Type != html.CommentNode {
		depth++
	}
}

func endElement(n *html.Node) {
	depth--
	if n.Type == html.ElementNode {
		if isBiTag, newLine := printNewLine(n); isBiTag {
			if newLine {
				fmt.Printf("%*s", depth*4, "")
			}
			fmt.Printf("</%s>", n.Data)
		}
		fmt.Println()
	}
}

// 判断是否是单标签
func printNewLine(n *html.Node) (isBiTag bool, newLine bool) {
	switch n.Data {
	case "meta", "link", "img", "input", "param", "hr", "br":
		return false, false
	case "a", "title", "h1":
		return true, false
	default:
		return true, true
	}
}
