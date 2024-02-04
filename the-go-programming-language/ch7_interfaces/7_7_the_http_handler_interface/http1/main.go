// Http1 is a rudimentary e-commerce server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// 在终端执行：
//
//  1. 启动服务: go run ./ch7_interfaces/7_7_the_http_handler_interface/http1/main.go
//  2. 在浏览器地址栏输入: localhost:8000
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

func (db database) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	for item, price := range db {
		_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func ListenAndServe(address string, h Handler) error {
	server := &http.Server{Addr: address, Handler: h}
	return server.ListenAndServe()
}
