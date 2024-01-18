// 练习4.5：
// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch4/4_2_slices/exercise_4_5/main.go
func main() {
	s1 := []string{"1", "2", "2", "2", "3", "3"}
	fmt.Println(removeDuplicates(s1)) // [1 2 3]
	fmt.Println(s1)                   // [1 2 3 2 3 3]

	s1 = []string{"2", "2", "2", "2", "3", "3", "3"}
	fmt.Println(removeDuplicates(s1)) // [2 3]
	fmt.Println(s1)                   // [2 3 2 2 3 3 3]

	s1 = []string{"1", "2", "3"}
	fmt.Println(removeDuplicates(s1)) // [1 2 3]
	fmt.Println(s1)                   // [1 2 3]
}

// 移除相邻重复的元素，使每个元素只出现一次
func removeDuplicates(s []string) []string {
	if s == nil || len(s) == 0 {
		return s
	}

	diff := 0
	for i, size := 1, len(s); i < size; i++ {
		if s[i] != s[diff] {
			diff++
			if diff != i {
				s[diff] = s[i]
			}
		}
	}

	return s[:diff+1]
}
