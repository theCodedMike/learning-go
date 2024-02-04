// Http3a is an e-commerce server that registers the /list and /price
// endpoints by calling (*http.ServeMux).HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// 在终端执行：
//
//  1. 启动服务: go run ./ch7_interfaces/7_7_the_http_handler_interface/http3a/main.go
//  2. 在浏览器地址栏输入:
//     localhost:8000
//     localhost:8000/list
//     localhost:8000/list?item=socks
//     localhost:8000/price
//     localhost:8000/price?item=socks
//     localhost:8000/price?item=glass
func main() {
	db := database{"shoes": 50, "socks": 5}

	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, _ *http.Request) {
	for item, price := range db {
		_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
}
