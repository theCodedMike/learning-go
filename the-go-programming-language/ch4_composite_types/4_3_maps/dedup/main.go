package main

// Dedup prints only one instance of each line; duplicates are removed.
import (
	"bufio"
	"fmt"
	"os"
)

// 在终端执行：
//
//	go run ./ch4_composite_types/4_3_maps/dedup/main.go
func main() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
