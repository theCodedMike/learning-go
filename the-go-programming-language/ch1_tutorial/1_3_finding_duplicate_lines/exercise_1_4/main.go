// 练习1.4：
// 修改dup2，出现重复的行时打印文件名称。
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 在终端执行：
//  1. go build ./ch1_tutorial/1_3_finding_duplicate_lines/exercise_1_4
//  2. ./exercise_1_4 ./test.txt
//     或者
//  3. go run ./ch1_tutorial/1_3_finding_duplicate_lines/exercise_1_4/main.go ./test.txt
func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

		fmt.Println("\n---出现次数大于1的结果如下---")
		for key, val := range counts {
			if val > 1 {
				fmt.Printf("%s: %15s\t%d\n", filename, key, val)
			}
		}
	}

}
