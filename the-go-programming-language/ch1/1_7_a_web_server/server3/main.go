// handler echoes the HTTP request.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// 在终端执行：
//
//  1. go build ./ch1/1_7_a_web_server/server3
//
//  2. ./server3
//
//     或者
//
//  3. go run ./ch1/1_7_a_web_server/server3/main.go
//
//  3. 打开浏览器，在地址栏输入"localhost:8000/" 或 "localhost:8000/hello" 或 "localhost:8000/?cycles=20&name=limei"
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// This field(指Form) is only available after ParseForm is called.
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
