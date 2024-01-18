// Package exercise_2_5
//
// 练习2.5：
// 表达式`x&(x-1)`用于将x的最低一个非零bit位清零。使用这个算法重写PopCount函数，然后比较性能。
package exercise_2_5

// PopCount 返回一个数字(二进制形式)中1的个数
func PopCount(x uint64) int {
	count := 0

	for x != 0 {
		y := x - 1
		if (x & y) == y {
			count++
		}
		x >>= 1
	}

	return count
}
