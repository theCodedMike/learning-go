// Package exercise_2_3
//
// 练习2.3：
// 重写PopCount函数，用一个循环代替单一的表达式。比较2个版本的性能。（11.4节会展示如何系统地比较2种不同实现的性能）
package exercise_2_3

// PopCount 返回一个数字(二进制形式)中1的个数
//
// 求余法
func PopCount(x uint64) int {
	count := 0

	for x != 0 {
		if (x % 2) == 1 {
			count++
		}
		x /= 2
	}

	return count
}
