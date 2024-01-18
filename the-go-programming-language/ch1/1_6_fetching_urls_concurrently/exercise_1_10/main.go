// 练习1.10：
// 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行2遍请求，查看2次时间是否有较大的差别，并且每次取到的响应内容
// 是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 在终端执行：
//  1. go build ./ch1/1_6_fetching_urls_concurrently/exercise_1_10
//  2. ./exercise_1_10 2 https://www.baidu.com
func main() {
	start := time.Now()
	// 获取执行的次数这一参数
	timesStr := os.Args[1]
	if len(timesStr) == 0 {
		fmt.Fprintf(os.Stderr, "第1个参数(表示执行的次数)为空")
		os.Exit(1)
	}
	times, err := strconv.ParseInt(timesStr, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "解析第1个参数为整数时(%s)出错", timesStr)
		os.Exit(1)
	}

	// 获取url这一参数
	url := os.Args[2]
	if len(url) == 0 {
		fmt.Fprintf(os.Stderr, "第2个参数(表示URL)为空")
		os.Exit(1)
	}

	// start a goroutine
	chn := make(chan string)
	for i := int64(1); i <= times; i++ {
		go fetch(url, i, chn)
	}
	// receive from channel chn
	for i := int64(0); i < times; i++ {
		fmt.Println(<-chn)
	}

	fmt.Printf("总耗时：%.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, times int64, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%d: %v", times, err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("%d: while reading %s: %v", times, url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("第%d次: %.2fs %7d %s", times, secs, nbytes, url)
}
