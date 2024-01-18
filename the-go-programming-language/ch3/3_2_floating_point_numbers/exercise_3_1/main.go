// 练习3.1：
// 如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素（虽然许多SVG渲染器会妥善处理这类问题）。修改程序跳过无效的多边形。
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyRange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyScale       = width / 2 / xyRange // pixels per x or y unit
	zScale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// 在终端执行：
//
//	go run ./ch3/3_2_floating_point_numbers/exercise_3_1/main.go >out.svg
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, isInvalid := corner(i+1, j)
			if isInvalid {
				continue
			}
			bx, by, isInvalid := corner(i, j)
			if isInvalid {
				continue
			}
			cx, cy, isInvalid := corner(i, j+1)
			if isInvalid {
				continue
			}
			dx, dy, isInvalid := corner(i+1, j+1)
			if isInvalid {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale

	return sx, sy, isNaNOrInf(sx) || isNaNOrInf(sy)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

func isNaNOrInf(v float64) bool {
	return math.IsNaN(v) || math.IsInf(v, 0)
}
