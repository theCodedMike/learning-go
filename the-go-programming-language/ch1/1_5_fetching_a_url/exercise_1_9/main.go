// 练习1.9：
// 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const HttpPrefix = "http://"

// 在终端执行：
//  1. go build ./ch1/1_5_fetching_a_url/exercise_1_9
//  2. ./exercise_1_9 gopl.io
//  3. ./exercise_1_9 bad.gopl.io
//     或者
//  4. go run ./ch1/1_5_fetching_a_url/exercise_1_9/main.go http://gopl.io
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, HttpPrefix) {
			url = HttpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("fetch %s, Status Code: %d\n", url, resp.StatusCode)
	}
}
