// 练习5.8：
// 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNode的遍历。使用修改后的代码编写ElementById函数，根据用户输入的id
// 查找第1个拥有该id元素的HTML元素，找到后停止遍历。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

const Indent int = 4

// 在终端执行：
//
//	go run ./ch5_functions/5_5_function_values/exercise_5_8/main.go  https://baidu.com
func main() {
	if len(os.Args) <= 1 {
		log.Printf("缺少url等参数")
		os.Exit(1)
	}

	url := os.Args[1]
	id := ""
	if len(os.Args) > 2 {
		id = os.Args[2]
	}

	find, err := outline(url, id)
	if err != nil {
		log.Printf("err: %v\n", err)
		os.Exit(1)
	}

	if find {
		fmt.Printf("\n\n---Find it---\n")
	} else {
		fmt.Printf("\n\n---No such a attribute---\n")
	}
}

func outline(url string, id string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return false, err
	}

	return ElementByID(doc, id) != nil, nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre func(n *html.Node, id string) bool, post func(n *html.Node)) (bool, *html.Node) {
	if pre != nil {
		if find := pre(n, id); find {
			return find, n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if find, node := forEachNode(c, id, pre, post); find {
			return find, node
		}
	}
	if post != nil {
		post(n)
	}
	return false, nil
}

var depth int

func startElement(n *html.Node, id string) bool {
	findId := false
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*Indent, "", n.Data)

		for _, attr := range n.Attr {
			if attr.Key == id {
				findId = true
			}
			fmt.Printf(" %s=%s", attr.Key, attr.Val)
		}

		fmt.Printf(">")
		if isBiTag(n.Data) {
			fmt.Println()
		}

		depth++
	}
	return findId
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if isBiTag(n.Data) {
			fmt.Printf("%*s</%s>", depth*Indent, "", n.Data)
		}
		fmt.Println()
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	_, node := forEachNode(doc, id, startElement, endElement)
	return node
}

func isBiTag(tag string) bool {
	switch tag {
	case "meta", "link", "img", "input", "param", "hr", "br":
		return false
	default:
		return true
	}
}
