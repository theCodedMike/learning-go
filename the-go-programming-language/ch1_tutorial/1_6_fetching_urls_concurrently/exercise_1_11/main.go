// 练习1.11：
// 在fetchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里排名靠前的。如果一个网站没有响应，程序将采取怎样的行为？
// （8.9小节描述了在这种情况下的应对机制）
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
//  1. go build ./ch1_tutorial/1_6_fetching_urls_concurrently/exercise_1_11
//  2. ./exercise_1_11 4 https://www.baidu.com https://cn.bing.com https://www.google.com
//     谷歌在国内是无法访问的；这里提供了3个网站，每个网址重复访问4次。
func main() {
	start := time.Now()
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

	ch := make(chan string)
	for _, url := range os.Args[2:] {
		for i := int64(1); i <= times; i++ {
			go fetch(url, i, ch)
		}
	}

	for range os.Args[2:] {
		for i := int64(1); i <= times; i++ {
			fmt.Println(<-ch)
		}
	}

	fmt.Printf("总耗时: %.2f(秒)\n", time.Since(start).Seconds())
}

// 执行次数那个循环不能放在fetch方法内，如果放在fetch方法内，main函数会提前结束(不会等待)。
func fetch(url string, i int64, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("=> %s, 第%d次: %v", url, i, err)
		return
	}

	size, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("=> %s, 第%d次: copy error, %v", url, i, err)
		return
	}

	ch <- fmt.Sprintf("[%s, 第%d次: 耗时%.2f(秒)，获取%d(字节)]", url, i, time.Since(start).Seconds(), size)
}

// [https://cn.bing.com, 第1次: 耗时0.48(秒)，获取9028(字节)]
// [https://cn.bing.com, 第3次: 耗时0.49(秒)，获取9028(字节)]
// [https://cn.bing.com, 第4次: 耗时0.49(秒)，获取9028(字节)]
// [https://cn.bing.com, 第2次: 耗时0.50(秒)，获取9028(字节)]
// [https://www.baidu.com, 第2次: 耗时0.68(秒)，获取2443(字节)]
// [https://www.baidu.com, 第1次: 耗时0.69(秒)，获取2443(字节)]
// [https://www.baidu.com, 第4次: 耗时0.69(秒)，获取2443(字节)]
// [https://www.baidu.com, 第3次: 耗时0.69(秒)，获取2443(字节)]
// => https://www.google.com, 第3次: Get "https://www.google.com": dial tcp [2a03:2880:f136:83:face:b00c:0:25de]:443: i/o timeout
// => https://www.google.com, 第1次: Get "https://www.google.com": dial tcp [2a03:2880:f136:83:face:b00c:0:25de]:443: i/o timeout
// => https://www.google.com, 第2次: Get "https://www.google.com": dial tcp [2a03:2880:f136:83:face:b00c:0:25de]:443: i/o timeout
// => https://www.google.com, 第4次: Get "https://www.google.com": dial tcp [2a03:2880:f136:83:face:b00c:0:25de]:443: i/o timeout
// 总耗时: 30.00(秒)
