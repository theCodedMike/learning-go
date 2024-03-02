package exercise_2_4

import "testing"

// 在终端执行：
//
// go test -bench=BenchmarkPopCount ./ch2_program_structure/2_6_packages_and_files/exercise_2_4
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(10)
	}
}
