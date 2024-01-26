// 练习5.17：
// 编写多参数版本的ElementByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。下面给出了2个例子。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// 在终端执行：
//
//	go run ./ch5_functions/5_7_variadic_functions/exercise_5_17/main.go https://baidu.com
func main() {
	url := os.Args[1]

	imgs := fetch(url, "img")
	for i, img := range imgs {
		fmt.Printf("<img> %d: %v\n", i+1, img)
	}
	fmt.Printf("%s has %d <img> tags\n\n", url, len(imgs))

	hs := fetch(url, "h1", "h2", "h3", "h4")
	for i, h := range hs {
		fmt.Printf("<h> %d: %v\n", i+1, h)
	}
	fmt.Printf("%s has %d <h> tags\n", url, len(hs))
}

func fetch(url string, name ...string) []*html.Node {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get failed: %v\n", err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("status is not OK: %s\n", resp.Status)
		return nil
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("html parse failed: %v\n", err)
		return nil
	}

	return elementsByTagName(doc, name...)
}

func elementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) == 0 {
		fmt.Println("标签名为空！")
		return nil
	}

	tags := make(map[string]bool)
	for _, tag := range name {
		tags[tag] = true
	}

	return traversal(doc, tags, nil)
}

func traversal(n *html.Node, tags map[string]bool, res []*html.Node) []*html.Node {
	if n.Type == html.ElementNode && tags[n.Data] {
		res = append(res, n)
	}

	for i := n.FirstChild; i != nil; i = i.NextSibling {
		res = traversal(i, tags, res)
	}

	return res
}
