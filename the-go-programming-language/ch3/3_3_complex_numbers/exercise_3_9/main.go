// 练习3.9：
// 编写一个web服务器，用于给客户端生成分形的图像。运行客户端并使用HTTP参数来指定x、y和zoom参数。
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

// 在终端执行：
//  1. go run ./ch3/3_3_complex_numbers/exercise_3_9/main.go
//  2. 打开浏览器，在地址栏输入`localhost:8080/pic?x=1024&y=1024&zoom=2`
func main() {
	http.HandleFunc("/pic", picHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func picHandler(w http.ResponseWriter, r *http.Request) {
	width, height, zoom := parseURL(r)
	xMin, yMin, xMax, yMax := -zoom, -zoom, zoom, zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(yMax-yMin) + float64(yMin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xMax-xMin) + float64(xMin)
			z := complex(x, y)
			// Image point (px, py) represents complex value Z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		log.Fatal("encode error: ", err)
	}
}

func parseURL(r *http.Request) (width int, height int, zoom int) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	x := r.Form.Get("x")
	y := r.Form.Get("y")
	z := r.Form.Get("zoom")

	width, errX := strconv.Atoi(x)
	if errX != nil {
		log.Print("parse x error: ", errX)
	}
	height, errY := strconv.Atoi(y)
	if errY != nil {
		log.Print("parse y error: ", errY)
	}
	zoom, errZ := strconv.Atoi(z)
	if errZ != nil {
		log.Print("parse zoom error: ", errZ)
	}

	return width, height, zoom
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
