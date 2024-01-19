// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// 在终端执行：
//
//  1. go build ./ch1_tutorial/1_7_a_web_server/server1
//
//  2. ./server1
//
//     或者
//
//  3. go run ./ch1_tutorial/1_7_a_web_server/server1/main.go
//
//  4. 打开浏览器，在地址栏输入"localhost:8000/"，此时页面会显示"URL.Path = "/""
//     在地址栏输入"localhost:8000/login"，此时页面会显示"URL.Path = "/login""
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
