// 练习4.12：
// 流行的web漫画服务xkcd也提供了JSON接口。例如，一个https://xkcd.com/571/info.0.json请求将返回一个很多人喜爱的571编号的详细描述。
// 下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	StoreFileName = "./ch4/issues/exercise_4_12/comic%d.json"
	FetchURL      = "https://xkcd.com/%d/info.0.json"
)

type XKCDInfo struct {
	Num              int
	Year, Month, Day string
	Link             string
	News             string
	Title            string
	SafeTitle        string `json:"safe_title"`
	Transcript       string
	Alt              string
	Img              string
}

// 在终端执行：
//  1. go build ./ch4_composite_types/4_5_json/exercise_4_12/main.go
//  2. ./exercise_4_12 571
//     ./exercise_4_12 571 animal
//     ./exercise_4_12 150
//     ./exercise_4_12 150 sleep
func main() {
	params := os.Args[1:]
	id, err := strconv.Atoi(params[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "未输入漫画编号")
		os.Exit(1)
	}

	filename := fmt.Sprintf(StoreFileName, id)
	exist, file, err := checkFileExist(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to check file: %v", err)
		os.Exit(1)
	}

	var transcript string
	if !exist {
		// 本地没有，则调用api查询
		fmt.Printf("%q not exist, need to get from network.\n", filename)
		url := fmt.Sprintf(FetchURL, id)
		fmt.Printf("Now fetch %q from Internet...", url)

		// 查询
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get: %v", err)
			os.Exit(1)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Fprintf(os.Stderr, "Status not ok: %v", resp.Status)
			os.Exit(1)
		}

		// json -> struct
		var data XKCDInfo
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			resp.Body.Close()
			fmt.Fprintf(os.Stderr, "Failed to decode: %v", err)
			os.Exit(1)
		}

		// 美化json
		content, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to marshalIndent: %s", err)
			os.Exit(1)
		}

		// 将美化后的json存在本地
		if _, err := file.Write(content); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write: %v", err)
			os.Exit(1)
		}

		transcript = data.Transcript
	} else {
		// 本地有缓存，重复利用
		fmt.Printf("%q exist, wo can use it.\n", filename)

		// 转变为struct
		var data XKCDInfo
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to decode: %v", err)
			os.Exit(1)
		}

		transcript = data.Transcript
	}

	fmt.Printf("\n\n%s\n\n", transcript)
	// 判断transcript字段是否有匹配的关键词
	if len(params) > 1 {
		term := params[1]
		if strings.Contains(transcript, term) {
			fmt.Printf("comic%d.json contains %q\n", id, term)
		} else {
			fmt.Printf("comic%d.json not contains %q\n", id, term)
		}
	}
}

func checkFileExist(filename string) (bool, *os.File, error) {
	if _, err := os.Stat(filename); err != nil {
		if !os.IsNotExist(err) {
			if err := os.Remove(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to remove: %v", err)
			}
		}
		file, err := os.Create(filename)
		return false, file, err
	} else {
		file, err := os.Open(filename)
		return true, file, err
	}
}
