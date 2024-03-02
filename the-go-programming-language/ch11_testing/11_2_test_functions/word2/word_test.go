package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 在终端执行：
//
// go test ./ch11_testing/11_2_test_functions/word2
//
// go test -coverprofile=c.out ./ch11_testing/11_2_test_functions/word2 // 可以测试覆盖率，这里会输出一个c.out文件
// go tool cover -html=c.out // 以html打开c.out文件
//
// go test -bench=BenchmarkIsPalindrome ./ch11_testing/11_2_test_functions/word2 // 执行基准测试
// go test -bench=BenchmarkIsPalindromeOptimized ./ch11_testing/11_2_test_functions/word2 // 执行基准测试

// 以下都是黑盒测试
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

// BenchmarkIsPalindrome is a benchmark function
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// BenchmarkIsPalindromeOptimized is a benchmark function
func BenchmarkIsPalindromeOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeOptimized("A man, a plan, a canal: Panama")
	}
}

// ExampleIsPalindrome is a example function
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}
