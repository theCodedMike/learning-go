// 练习5.5：
// 实现countWordsAndImages。（参考练习4.9如何分词）
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
//	go run ./ch5_functions/5_3_multiple_return_values/exercise_5_5/main.go  https://baidu.com
func main() {
	for _, url := range os.Args[1:] {
		counter, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Url: %s, Err: %v\n", url, err)
			continue
		}

		for k, v := range counter {
			fmt.Printf("%s: %d个\n", k, len(counter[k]))
			for _, s := range v {
				fmt.Printf("  %s\n", s)
			}
			fmt.Println()
		}
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (counter map[string][]string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("status not ok")
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %s", err)
	}

	return countWordsAndImages(doc, nil), nil
}
func countWordsAndImages(n *html.Node, counter map[string][]string) map[string][]string {
	if n == nil {
		return nil
	}
	if counter == nil {
		counter = make(map[string][]string)
	}

	if n.Type == html.TextNode {
		for _, word := range strings.Fields(n.Data) {
			if len(word) != 0 {
				counter["Words"] = append(counter["Words"], word)
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				counter["Images"] = append(counter["Images"], attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		counter = countWordsAndImages(c, counter)
	}

	return counter
}
