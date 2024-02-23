// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"gopl.io/ch5_functions/5_6_anonymous_functions/links"
	"log"
	"os"
)

// 在终端执行：
//
//	go run ./ch8_goroutines_and_channels/8_6_example_concurrent_web_crawler/crawl2/main.go https://baidu.com
func main() {
	workList := make(chan []string)
	var n int // number of pending sends to workList

	// Start with the command-line arguments.
	n++
	go func() {
		workList <- os.Args[1:]
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}
