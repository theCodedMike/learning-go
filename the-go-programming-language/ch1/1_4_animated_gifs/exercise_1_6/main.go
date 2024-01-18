// 练习1.6：
// 修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex的第3个参数，看看显示结果吧
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var newPalette = []color.Color{
	color.White, // 白色
	color.RGBA{ // #00FF00 绿色
		G: 255,
		A: 100, // 100表示完全不透明，0表示完全透明
	}, color.RGBA{ // #00B2EE 深蓝色2
		G: 178,
		B: 238,
		A: 100,
	}}

const (
	whiteIdx = 0 // first color in newPalette
	greenIdx = 1 // second color in newPalette
	blueIdx  = 2 // third color in newPalette
)

// 在终端执行：
//  1. go build ./ch1/1_4_animated_gifs/exercise_1_6
//  2. ./exercise_1_6 >out.gif
//     或者
//  3. go run ./ch1/1_4_animated_gifs/exercise_1_6/main.go >out.gif
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	myLissajous(os.Stdout)
}

func myLissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, newPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blueIdx)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
