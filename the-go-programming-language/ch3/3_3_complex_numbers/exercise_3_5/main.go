// 练习3.5：
// 实现一个彩色的Mandelbrot图像，使用image.NewRGBA创建图像，使用color.RGBA或color.YCbCr生成颜色。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// 在终端执行：
//
//	go run ./ch3/3_3_complex_numbers/exercise_3_5/main.go >out.png
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value Z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				R: n,
				G: 255 - contrast*n,
				B: 255 - n,
				A: 100,
			}
		}
	}

	// #9370DB
	return color.RGBA{
		R: 147,
		G: 112,
		B: 219,
		A: 90,
	}
}
