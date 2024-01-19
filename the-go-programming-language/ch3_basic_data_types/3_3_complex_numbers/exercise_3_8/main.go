// 练习3.8：
// 通过提高精度来生成更多级别的分形。使用4种不同精度类型的数字实现相同的分形：complex64、complex128、big.Float和big.Rat。（后面2种类型在
// math/big包声明。Float是有限精度的浮点数；Rat是无限精度的有理数。）它们的性能和内存使用对比如何？当渲染图可见时缩放的级别是多少？
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
//	go run ./ch3_basic_data_types/3_3_complex_numbers/exercise_3_8/main.go >out.png
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value Z.
			img.Set(px, py, mandelbrot(complex128(z)))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}

	return color.Black
}
