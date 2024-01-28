package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// 在终端执行：
//
//	go run ./ch5_functions/5_10_recover/title3/main.go https://baidu.com https://gopl.io
func main() {
	for _, url := range os.Args[1:] {
		title, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch %q fialed: %v\n\n", url, err)
			continue
		}

		fmt.Printf("title of %q is %q\n\n", url, title)
	}
}

func fetch(url string) (title string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http get failed: %v\n", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status not ok: %v\n", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parse failed: %v\n", err)
	}

	return soleTitle(doc)
}

// soleTitle returns the text of the first non-empty title element in doc,
// and an error if there was not exactly one.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case bailout{}: /// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic: carry on panicking
		}
	}()

	// Bail out of recursion if we find more than one nonempty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}

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
