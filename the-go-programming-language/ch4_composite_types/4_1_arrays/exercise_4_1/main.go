// 练习4.1：
// 编写一个函数，计算2个SHA256哈希吗中不同bit的数目。（参考2.6.2节的PopCount函数。）
package main

import (
	"crypto/sha256"
	"fmt"
)

// 在终端执行：
//
//	go run ./ch4_composite_types/4_1_arrays/exercise_4_1/main.go
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(countDiffBit(&c1, &c2)) // 125

	c1 = [32]byte{} // c1数组清零
	for i := 0; i < 32; i++ {
		c2[i] = 255
	}
	fmt.Println(countDiffBit(&c1, &c2)) // 256

	c2 = [32]byte{}                     // c2数组清零
	fmt.Println(countDiffBit(&c1, &c2)) // 0

	c2[0] = 255
	c2[1] = 255
	fmt.Println(countDiffBit(&c1, &c2)) // 16
}

func countDiffBit(a *[32]byte, b *[32]byte) uint {
	var diff uint

	for i := 0; i < 32; i++ {
		diff += countBit1(a[i] ^ b[i])
	}

	return diff
}

func countBit1(p byte) uint {
	var count uint

	for p != 0 {
		if p&1 == 1 {
			count++
		}
		p >>= 1
	}

	return count
}
