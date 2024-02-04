// 练习7.12：
// 修改/list的handler，让它把输出打印成一个HTML的表格而不是文本。html/template包可能对你有帮助。
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var dbListTemplate = template.Must(template.New("dbListTemplate").Parse(`
<table>
<h1>{{.Total}} items:</h1>
<tr style='text-align: left'>
  <th>#</th>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range .Items}}
<tr>
  <td>{{.Id}}</td>
  <td>{{.Item}}</td>
  <td>{{.Price}}</td>
</tr>
{{end}}
</table>
`))

// 在终端执行：
//
//  1. 启动服务: go run ./ch7_interfaces/7_7_the_http_handler_interface/exercise_7_12/main.go
//  2. 在浏览器地址栏输入:
//     localhost:8000/list
//     localhost:8000/insert?item=computer&price=5520
//     localhost:8000/insert?item=glass&price=350
//     localhost:8000/update?item=glass&price=454
//     localhost:8000/list
//     localhost:8000/delete?item=computer
//     localhost:8000/list
func main() {
	db := database{"shoes": 50, "socks": 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/insert", db.insert)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var lock sync.RWMutex

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

type DBItem struct {
	Id    int
	Item  string
	Price dollars
}

type DBResult struct {
	Total int
	Items []*DBItem
}

func (db database) list(w http.ResponseWriter, _ *http.Request) {
	var dbItems []*DBItem

	idx := 0
	for item, price := range db {
		idx++
		one := new(DBItem)
		one.Id = idx
		one.Item = item
		one.Price = price
		dbItems = append(dbItems, one)
	}

	var dbResult = DBResult{Total: idx, Items: dbItems}

	if err := dbListTemplate.Execute(w, dbResult); err != nil {
		log.Fatal(err)
	}
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
