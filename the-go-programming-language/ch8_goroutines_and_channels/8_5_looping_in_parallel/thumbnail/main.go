package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 在终端执行：
//
//  1. go build ./ch8_goroutines_and_channels/8_5_looping_in_parallel/thumbnail
//  2. ./thumbnail
//  3. ./http_status_code.jpg
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Println(thumb)
	}

	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
