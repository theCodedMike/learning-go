package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 在终端执行：
//  1. go build ./ch1/1_3_finding_duplicate_lines/dup3
//  2. ./dup3 ./test.txt
//     或者
//  3. go run ./ch1/1_3_finding_duplicate_lines/dup3/main.go ./test.txt
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
				fmt.Printf("%d\t%s\n", val, key)
			}
		}
	}

}
