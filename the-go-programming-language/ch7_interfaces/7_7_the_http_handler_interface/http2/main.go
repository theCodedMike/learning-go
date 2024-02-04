// Http2 is an e-commerce server with /list and /price endpoints.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// 在终端执行：
//
//  1. 启动服务: go run ./ch7_interfaces/7_7_the_http_handler_interface/http2/main.go
//  2. 在浏览器地址栏输入:
//     localhost:8000
//     localhost:8000/list
//     localhost:8000/list?item=socks
//     localhost:8000/price
//     localhost:8000/price?item=socks
//     localhost:8000/price?item=glass
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(ListenAndServe("localhost:8000", db))
}

// Handler
//
// This is a custom Handler interface
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		_, _ = fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func ListenAndServe(address string, h Handler) error {
	server := &http.Server{Addr: address, Handler: h}
	return server.ListenAndServe()
}
