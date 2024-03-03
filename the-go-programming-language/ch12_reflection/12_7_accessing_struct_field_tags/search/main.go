package main

import (
	"fmt"
	"gopl.io/ch12_reflection/12_7_accessing_struct_field_tags/params"
	"log"
	"net/http"
)

// 在终端执行：
//
//  1. go run ./ch12_reflection/12_7_accessing_struct_field_tags/search/main.go
//  2. 在浏览器地址栏输入：
//     http://localhost:12345/search
//     http://localhost:12345/search?l=golang&l=programming&max=100
//     http://localhost:12345/search?x=true&l=golang&l=programming
//     http://localhost:12345/search?q=hello&x=123
//     http://localhost:12345/search?q=hello&max=lots
func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	_, _ = fmt.Fprintf(resp, "Search: %+v\n", data)
}
