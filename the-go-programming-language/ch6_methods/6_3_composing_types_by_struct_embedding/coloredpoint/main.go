// Colored point demonstrates struct embedding.
package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (c ColoredPoint) Distance(q Point) float64 {
	return c.Point.Distance(q)
}

func (c *ColoredPoint) ScaleBy(factor float64) {
	c.Point.ScaleBy(factor)
}

// 在终端执行：
//
//	go run ./ch6_methods/6_3_composing_types_by_struct_embedding/coloredpoint/main.go
func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	fmt.Println(p.Distance(q.Point)) // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // 10
}
