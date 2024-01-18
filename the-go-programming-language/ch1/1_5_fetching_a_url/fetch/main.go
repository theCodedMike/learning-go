// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 在终端执行：
//  1. go build ./ch1/1_5_fetching_a_url/fetch
//  2. ./fetch http://gopl.io
//  3. ./fetch http://bad.gopl.io
//     或者
//  4. go run ./ch1/1_5_fetching_a_url/fetch/main.go http://gopl.io
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}

	// 部分输出如下：
	// $ ./fetch http://bad.gopl.io
	// fetch: Get "http://bad.gopl.io": dial tcp: lookup bad.gopl.io: no such host
}
