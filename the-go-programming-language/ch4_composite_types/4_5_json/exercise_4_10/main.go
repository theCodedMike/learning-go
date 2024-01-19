// 练习4.10：
// 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年的。
package main

import (
	"fmt"
	"gopl.io/ch4_composite_types/4_5_json/github"
	"log"
	"os"
	"time"
)

const (
	OneMonthHours = 30 * 24
	OneYearHours  = OneMonthHours * 12
)

// 在终端执行：
//  1. go build ./ch4_composite_types/4_5_json/exercise_4_10
//  2. ./exercise_4_10 repo:golang/go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		dateTime, cls := classifyByCreatedTime(item.CreatedAt)
		fmt.Printf("#%-5d, %9.9s, %s, %s, %.55s\n", item.Number, item.User.Login, dateTime, cls, item.Title)
	}
}

func classifyByCreatedTime(createTime time.Time) (dateTime string, classification string) {
	dateTime = createTime.Format(time.DateTime)

	dur := time.Since(createTime).Hours()
	classification = ""
	if dur < OneMonthHours {
		classification = "不到一个月"
	} else if dur < OneYearHours {
		classification = "不到一年"
	} else {
		classification = "超过一年"
	}

	return dateTime, classification
}
