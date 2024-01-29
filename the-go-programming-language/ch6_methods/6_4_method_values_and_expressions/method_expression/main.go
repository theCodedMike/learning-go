package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) String() string {
	return fmt.Sprintf("{x: %.2v, y: %.2v}", p.X, p.Y)
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

// 在终端执行：
//
//	go run ./ch6_methods/6_4_method_values_and_expressions/method_expression/main.go
func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // 5
	fmt.Printf("%T\n", distance) // func(main.Point, main.Point) float64

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // {x: 2, y: 4}
	fmt.Printf("%T\n", scale) // func(*main.Point, float64)

	line := []Point{{0, 0}, {2, 2}, {8, 10}}
	for i, point := range line {
		fmt.Printf("%d: %v\n", i+1, point)
	}
	//1: {x: 0, y: 0}
	//2: {x: 2, y: 2}
	//3: {x: 8, y: 10}
	Path(line).TranslateBy(Point{-1, -2}, true)

	for i, point := range line {
		fmt.Printf("%d: %v\n", i+1, point)
	}
	//1: {x: -1, y: -2}
	//2: {x: 1, y: 0}
	//3: {x: 7, y: 8}
}
