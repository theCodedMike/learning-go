// 练习7.11：
// 增加额外的handler让客户端可以创建、读取、更新和删除数据库记录。例如一个形如`/update?item=socks&price=6`的请求会更新库存清单里一个货品
// 的价格并且当这个货品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// 在终端执行：
//
//  1. 启动服务: go run ./ch7_interfaces/7_7_the_http_handler_interface/exercise_7_11/main.go
//  2. 在浏览器地址栏输入:
//     localhost:8000
//     localhost:8000/list
//     localhost:8000/insert?item=computer&price=5520
//     localhost:8000/insert?item=glass&price=350
//     localhost:8000/update?item=glass&price=454
//     localhost:8000/delete?item=computer
func main() {
	db := database{"shoes": 50, "socks": 5}

	http.HandleFunc("/", db.list)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/insert", db.insert)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var lock sync.RWMutex

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

func (db database) insert(w http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	args := req.URL.Query()
	item := strings.TrimSpace(args.Get("item"))
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "param item is invalid.\n")
		return
	}

	price := strings.TrimSpace(args.Get("price"))
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "param price is invalid.\n")
		return
	}

	f64Price, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "param price is invalid: %v\n", price)
		return
	}

	key := strings.ToLower(item)
	old, exist := db[key]
	db[key] = dollars(float32(f64Price))

	w.WriteHeader(http.StatusOK)
	if exist {
		_, _ = fmt.Fprintf(w, "successfully update, %s: %s\n", key, old)
	} else {
		_, _ = fmt.Fprintf(w, "successfully insert, %s: %s\n", key, db[key])
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	args := req.URL.Query()
	item := strings.TrimSpace(args.Get("item"))
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "param item is invalid.\n")
		return
	}

	key := strings.ToLower(item)
	old, exist := db[key]
	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "no such item: %v.\n", key)
		return
	}

	price := strings.TrimSpace(args.Get("price"))
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "param price is invalid.\n")
		return
	}

	f64Price, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "param price is invalid: %v\n", price)
		return
	}

	db[key] = dollars(float32(f64Price))
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "successfully update, %s: %s\n", key, old)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	args := req.URL.Query()
	item := strings.TrimSpace(args.Get("item"))
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "param item is invalid.\n")
		return
	}

	key := strings.ToLower(item)
	old, exist := db[key]
	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "no such item: %v.\n", key)
		return
	}

	delete(db, key)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "successfully delete, %s: %s\n", key, old)
}
