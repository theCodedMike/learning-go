// Package exercise_2_1
//
// 练习2.1：
// 向tempconv包添加类型、常量和函数用来处理Celvin绝对温度的转换，Kelvin绝对零度是-273.15°C，kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的
//
// T(K) = 273.15 + t(°C)
package exercise_2_1

import "fmt"

type Kelvin float64

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
