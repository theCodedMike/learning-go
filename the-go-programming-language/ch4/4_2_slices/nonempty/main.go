// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch4/4_2_slices/nonempty/main.go
func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)           // ["one" "three" "three"]
}

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
