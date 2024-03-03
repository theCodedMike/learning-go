// Bzipper reads input, bzip2-compresses it, and writes it out.
package main

import (
	"gopl.io/ch13_low_level_programming/13_4_calling_c_code_with_cgo/bzip"
	"io"
	"log"
	"os"
)

// 在终端执行：
//
//  1. 构建：go build ./ch13_low_level_programming/13_4_calling_c_code_with_cgo/bzipper
//  2. 压缩
//     wc -c < /usr/share/dict/MainEnglishDictionary_ProbWL.txt
//     10894740
//     sha256sum < /usr/share/dict/MainEnglishDictionary_ProbWL.txt
//     63b970db6b2d76efe192f66b33ba8b55f02d1809ad88ae6cc18dfd1291989a85  -
//     ./bzipper < /usr/share/dict/MainEnglishDictionary_ProbWL.txt | wc -c
//     3759817
//     ./bzipper < /usr/share/dict/MainEnglishDictionary_ProbWL.txt | bunzip2 | sha256sum
//     63b970db6b2d76efe192f66b33ba8b55f02d1809ad88ae6cc18dfd1291989a85  -
func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
