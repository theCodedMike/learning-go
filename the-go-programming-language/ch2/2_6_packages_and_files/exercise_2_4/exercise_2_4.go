// Package exercise_2_4
//
// 练习2.4：
// 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。
package exercise_2_4

// PopCount 返回一个数字(二进制形式)中1的个数
//
// 每次查看最右边的比特位是0或1
func PopCount(x uint64) int {
	count := 0

	for x != 0 {
		if (x & 1) == 1 {
			count++
		}
		x >>= 1
	}

	return count
}
