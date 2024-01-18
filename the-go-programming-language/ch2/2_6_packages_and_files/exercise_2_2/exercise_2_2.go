// 练习2.2：
// 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，
// 长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
package main

import "fmt"

type Foot float64     // 英尺
type Metre float64    // 米
type Pound float64    // 磅
type Kilogram float64 // 千克(公斤)

const (
	FootRateMetre = 3.2808399  // 1米 = 3.2808399英尺
	PoundRateKg   = 2.20462262 // 1千克 = 2.20462262磅
)

func (f Foot) String() string {
	return fmt.Sprintf("%g(英尺)", f)
}

func (m Metre) String() string {
	return fmt.Sprintf("%g(米)", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g(磅)", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g(千克)", k)
}

// FToM converts a Foot(英尺) to Metre(米).
func FToM(f Foot) Metre {
	return Metre(f / FootRateMetre)
}

// MToF converts a Metre(米) to Foot(英尺).
func MToF(m Metre) Foot {
	return Foot(m * FootRateMetre)
}

// PToKg converts a Pound(磅) to Kilogram(千克).
func PToKg(p Pound) Kilogram {
	return Kilogram(p / PoundRateKg)
}

// KgToP converts a Kilogram(千克) to Pound(磅).
func KgToP(kg Kilogram) Pound {
	return Pound(kg * PoundRateKg)
}
