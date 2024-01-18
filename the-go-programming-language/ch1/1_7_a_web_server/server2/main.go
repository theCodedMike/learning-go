// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// 在终端执行：
//
//  1. go build ./ch1/1_7_a_web_server/server2
//
//  2. ./server2
//
//     或者
//
//  3. go run ./ch1/1_7_a_web_server/server2/main.go
//
//  4. 打开浏览器，在地址栏输入"localhost:8000/"，此时页面会显示"URL.Path = "/""
//     在地址栏输入"localhost:8000/login"，此时页面会显示"URL.Path = "/login""
//     在地址栏输入"localhost:8000/count"，此时页面会显示"Count 5""
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
