// 练习1.8：
// 修改fetch这个例子，如果输入的url参数没有http://前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const HttpPrefix = "http://"

// 在终端执行：
//  1. go build ./ch1_tutorial/1_5_fetching_a_url/exercise_1_8
//  2. ./exercise_1_8 gopl.io
//  3. ./exercise_1_8 bad.gopl.io
//     或者
//  4. go run ./ch1_tutorial/1_5_fetching_a_url/exercise_1_8/main.go http://gopl.io
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

		all, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "readall error: %s, %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", all)
	}
}
