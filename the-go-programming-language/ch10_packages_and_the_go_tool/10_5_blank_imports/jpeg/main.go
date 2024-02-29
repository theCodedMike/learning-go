// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

// 在终端执行：
//
//  1. go build ./ch3_basic_data_types/3_3_complex_numbers/mandelbrot
//  2. go build ./ch10_packages_and_the_go_tool/10_5_blank_imports/jpeg
//  3. ./mandelbrot | ./jpeg >out2.jpg
func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{
		Quality: 95,
	})
}
