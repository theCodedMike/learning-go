// 练习3.3：
// 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色(#0000ff)
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	rows          = 3
	cols          = 8
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var tmpData [rows][cols]float64

// 在终端执行：
//
//	go run ./ch3_basic_data_types/3_2_floating_point_numbers/exercise_3_3/main.go >out.svg
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			setColor([]float64{ax, ay, bx, by, cx, cy, dx, dy})
		}
	}
	// 处理最后一行数据
	setColor([]float64{})

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

func setColor(row []float64) {
	for i, v := range row {
		tmpData[rows-1][i] = v
	}

	if tmpData[1][0] != 0 {
		var stroke = ""
		if isPeak() {
			stroke = "stroke=\"red\""
		} else if isTrough() {
			stroke = "stroke=\"blue\""
		}
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' %s/>\n",
			tmpData[1][0], tmpData[1][1], tmpData[1][2], tmpData[1][3], tmpData[1][4], tmpData[1][5], tmpData[1][6], tmpData[1][7], stroke)
	}

	// 后2行数据上移1行，即row只能插入在第3行
	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tmpData[i-1][j] = tmpData[i][j]
			if i == rows-1 {
				tmpData[i][j] = 0
			}
		}
	}
}

// 判断是否是波峰
// 这个算法不对，不清楚如何判断什么时候是波峰
func isPeak() bool {
	if tmpData[0][0] == 0 || tmpData[rows-1][0] == 0 {
		return false
	}
	var isPeak = true
	for i := 0; i < 8; i++ {
		if !(tmpData[1][i] >= tmpData[0][i] && tmpData[1][i] >= tmpData[2][i]) {
			isPeak = false
			break
		}
	}
	return isPeak
}

// 判断是否是谷底
// 这个算法不对，不清楚如何判断什么时候是波谷
func isTrough() bool {
	if tmpData[0][0] == 0 || tmpData[rows-1][0] == 0 {
		return false
	}
	var isTrough = true
	for i := 0; i < 8; i++ {
		if !(tmpData[1][i] <= tmpData[0][i] && tmpData[1][i] <= tmpData[2][i]) {
			isTrough = false
			break
		}
	}
	return isTrough
}
