// 练习1.7：
// 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替换例子中的ioutil.ReadAll来拷贝响应结构体到
// os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 在终端执行：
//  1. go build ./ch1_tutorial/1_5_fetching_a_url/exercise_1_7
//  2. ./exercise_1_7 http://gopl.io
//  3. ./exercise_1_7 https://bad.gopl.io
//     或者
//  4. go run ./ch1_tutorial/1_5_fetching_a_url/exercise_1_7/main.go http://gopl.io
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy error: %s, %v\n", url, err)
			os.Exit(1)
		}
	}
}
