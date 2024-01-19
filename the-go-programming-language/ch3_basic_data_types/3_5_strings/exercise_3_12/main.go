// 练习3.12：
// 编写一个函数，判断2个字符串是否是相互打乱的，即它们有着相同的字符，但是对应不同的顺序
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch3_basic_data_types/3_5_strings/exercise_3_12/main.go
func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println(isAnagram("", ""))               // true
	fmt.Println(isAnagram("bbb", "bb"))          // false
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counter := make(map[rune]int)

	// count
	for _, ch := range s1 {
		counter[ch]++
	}

	// compare
	for _, ch := range s2 {
		if count, ok := counter[ch]; !ok || count == 0 {
			return false
		}
		counter[ch]--
	}

	return true
}
