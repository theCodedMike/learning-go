package word

// 在终端执行：
//
// go test ./ch11_testing/11_2_test_functions/word1  测试这个包下的所有测试函数
// go test -v ./ch11_testing/11_2_test_functions/word1  -v 可以输出每个测试函数的运行时间
// go test -run=TestFrenchPalindrome ./ch11_testing/11_2_test_functions/word1 单独测试`TestFrenchPalindrome`这个函数

import "testing"

// 以下都是黑盒测试
func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}

}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
