// 练习4.2：
// 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
)

// 在终端执行：
//  1. go build ./ch4_composite_types/4_1_arrays/exercise_4_2
//  2. ./exercise_4_2 x      (这里会按256格式编码)
//     ./exercise_4_2 X      (这里会按256格式编码)
//     ./exercise_4_2 256 x  (这里会按256格式编码)
//     ./exercise_4_2 384 x  (这里会按384格式编码)
//     ./exercise_4_2 512 x  (这里会按512格式编码)
//     ./exercise_4_2 50  x  (这里会按256格式编码)
//     ./exercise_4_2 102 x  (这里会按256格式编码)
//     ./exercise_4_2 102 x y z (这里会按256格式编码)
func main() {
	if l := len(os.Args); l <= 1 {
		fmt.Fprintf(os.Stderr, "参数个数不足")
		os.Exit(1)
	}

	from := 2
	bits, err := strconv.Atoi(os.Args[1])
	if err != nil {
		from = 1
		bits = 256
	}

	for _, s := range os.Args[from:] {
		data := []byte(s)
		if bits == 384 {
			fmt.Printf("SHA384(%q) = %x\n", s, sha512.Sum384(data))
		} else if bits == 512 {
			fmt.Printf("SHA512(%q) = %x\n", s, sha512.Sum512(data))
		} else {
			fmt.Printf("SHA256(%q) = %x\n", s, sha256.Sum256(data))
		}
	}

	// SHA256("x") = 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// SHA256("X") = 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015

}
