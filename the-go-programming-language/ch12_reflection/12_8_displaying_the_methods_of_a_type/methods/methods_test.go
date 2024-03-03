package methods_test

// 在终端执行：
//
// go test -v ./ch12_reflection/12_8_displaying_the_methods_of_a_type/methods

import (
	"gopl.io/ch12_reflection/12_8_displaying_the_methods_of_a_type/methods"
	"strings"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	methods.Print(time.Hour)
	// Output:
	// type time.Duration
	// func (time.Duration) Abs() time.Duration
	// func (time.Duration) Hours() float64
	// func (time.Duration) Microseconds() int64
	// func (time.Duration) Milliseconds() int64
	// func (time.Duration) Minutes() float64
	// func (time.Duration) Nanoseconds() int64
	// func (time.Duration) Round(time.Duration) time.Duration
	// func (time.Duration) Seconds() float64
	// func (time.Duration) String() string
	// func (time.Duration) Truncate(time.Duration) time.Duration

	methods.Print(new(strings.Replacer))
	// Output:
	// type *strings.Replacer
	// func (*strings.Replacer) Replace(string) string
	// func (*strings.Replacer) WriteString(io.Writer, string) (int, error)

}
