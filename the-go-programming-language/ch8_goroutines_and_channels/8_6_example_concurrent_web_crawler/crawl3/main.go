// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
package main

import (
	"fmt"
	"gopl.io/ch5_functions/5_6_anonymous_functions/links"
	"log"
	"os"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_6_example_concurrent_web_crawler/crawl3/main.go https://baidu.com
func main() {
	workList := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to workList
	go func() {
		workList <- os.Args[1:]
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	// The main goroutine de-duplicates workList items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
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
