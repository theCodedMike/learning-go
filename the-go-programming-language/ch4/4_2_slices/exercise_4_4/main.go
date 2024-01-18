// 练习4.4：
// 编写一个rotate函数，通过一次循环完成旋转
package main

import "fmt"

// 在终端执行：
//
//	go run ./ch4/4_2_slices/exercise_4_4/main.go
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("before rotate: ", nums) // [1 2 3 4 5]
	rotate(nums, 2)
	fmt.Println(" after rotate: ", nums) // [4 5 1 2 3]

	nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("before rotate: ", nums) // [1 2 3 4 5 6]
	rotate(nums, -2)
	fmt.Println(" after rotate: ", nums) // [3 4 5 6 1 2]
}

// rotate slice, if step > 0, rotate right; if step < 0, rotate left
func rotate(data []int, step int) {
	size := len(data)
	if size == 0 {
		return
	}
	step = step % size
	if step == 0 {
		return
	}

	if step > 0 {
		// 向右旋转
		for i := 0; i < step; i++ {
			temp := data[size-1]
			for j := size - 2; j >= 0; j-- {
				data[j+1] = data[j]
			}
			data[0] = temp
		}
	} else {
		// 向左旋转
		for i := 0; i > step; i-- {
			temp := data[0]
			for j := 1; j < size; j++ {
				data[j-1] = data[j]
			}
			data[size-1] = temp
		}
	}
}
