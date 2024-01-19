// 练习2.2：
// 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，
// 长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
package main

import (
	"fmt"
	"os"
	"strconv"
)

// 在终端执行：
//  1. go build ./ch2_program_structure/2_6_packages_and_files/exercise_2_2
//  2. ./exercise_2_2
func main() {
	if len(os.Args) > 1 {
		// 从命令行读取参数
		for _, arg := range os.Args[1:] {
			val, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Parse Error: %v\n", err)
				os.Exit(1)
			}
			printVal(val)
		}
	} else {
		// 从标准输入读取参数
		var input float64
		fmt.Print("请输入数值(Ctrl + D to quit)：")
		for {
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Scan Error: %v\n", err)
				os.Exit(1)
			}
			printVal(input)
		}
	}
}

func printVal(input float64) {
	foot, metre := Foot(input), Metre(input)
	fmt.Printf("%s = %s, %s = %s\n", foot, FToM(foot), metre, MToF(metre))
	pound, kg := Pound(input), Kilogram(input)
	fmt.Printf("%s = %s, %s = %s\n", pound, PToKg(pound), kg, KgToP(kg))
}
