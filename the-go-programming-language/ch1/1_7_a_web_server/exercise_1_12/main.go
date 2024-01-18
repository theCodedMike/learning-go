// 练习1.12：
// 修改Lissajour服务，从URL读取变量，比如你可以访问"http://localhost:8000/?cycles=20"这个URL，这样访问可以将程序里的cycles默认的5
// 修改为20。字符串转换为数字可以调用strconv.Atoi函数。你可以在godoc里查看strconv.Atoi的详细说明。
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var (
	palette = []color.Color{color.White, color.Black, color.RGBA{
		R: 160,
		G: 32,
		B: 240,
		A: 100,
	}}
	cycles = 5 // number of complete x oscillator revolutions
	lock   sync.Mutex
)

const (
	whiteIdx  = 0 // first color in palette
	blackIdx  = 1 // second color in palette
	purpleIdx = 2 // third color in palette

	cyclesFiled = "cycles"
)

// 在终端执行：
//
//  1. go build ./ch1/1_7_a_web_server/exercise_1_12
//
//  2. ./exercise_1_12
//
//     或者
//
//  3. go run ./ch1/1_7_a_web_server/exercise_1_12/main.go
//
//  4. 打开浏览器，在地址栏输入"localhost:8080/" 或 "localhost:8080/?cycles=20" 或 "localhost:8080/?name=limei"
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if r.Form.Has(cyclesFiled) {
		val := r.Form.Get(cyclesFiled)
		if len(val) != 0 {
			newCycles, err := strconv.Atoi(val)
			if err != nil {
				log.Print(err)
			} else {
				cycles = newCycles
			}
		}
	}

	// 这里需要设置Content-Type，否则图像不显示
	w.Header().Set("Content-Type", "image/gif")
	lissajous(w)
	lock.Unlock()
}

func lissajous(out io.Writer) {
	const (
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
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), purpleIdx)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	if err := gif.EncodeAll(out, &anim); err != nil {
		log.Print(err)
	}
}
