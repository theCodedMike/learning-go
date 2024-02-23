// Crawl1 crawls web links starting with the command-line arguments.
//
// This version quickly exhausts available file descriptors
// due to excessive concurrent calls to links.Extract.
//
// Also, it never terminates because the worklist is never closed.
package main

import (
	"fmt"
	"gopl.io/ch5_functions/5_6_anonymous_functions/links"
	"log"
	"os"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_6_example_concurrent_web_crawler/crawl1/main.go https://baidu.com
func main() {
	workList := make(chan []string)

	// Start with the command-line arguments.
	go func() {
		workList <- os.Args[1:]
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
