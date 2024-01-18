// 练习4.3：
// 重写reverse函数，使用数组指针代替slice。
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch4/4_2_slices/exercise_4_3/main.go
func main() {
	// array
	nums1 := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&nums1)
	fmt.Println(nums1) // [5 4 3 2 1 0]

	nums2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	reverse((*[6]int)(nums2[:6]))
	fmt.Println(nums2) // [6 5 4 3 2 1 7]

	// slice
	s1 := []int{0, 1, 2, 3, 4, 5}
	reverse((*[6]int)(s1))
	fmt.Println(s1) // [5 4 3 2 1 0]

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	reverse((*[6]int)(s2[1:7]))
	fmt.Println(s2) // [1 7 6 5 4 3 2 8]

	s3 := []int{2, 4, 6, 8}
	// panic: runtime error: cannot convert slice with length 4 to array or pointer to array with length 6
	reverse((*[6]int)(s3))
	fmt.Println(s3)

}

func reverse(data *[6]int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
