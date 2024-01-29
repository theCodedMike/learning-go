package main

import (
	"fmt"
	"math"
	"time"
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

type Rocket struct {
	/* ... */
}

func (r *Rocket) Launch() {
	fmt.Println("this is rocket launch method...")
	/* ... */
}

// 在终端执行：
//
//	go run ./ch6_methods/6_4_method_values_and_expressions/method_value/main.go
func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance       // method value
	fmt.Printf("%T\n", distanceFromP) // func(main.Point) float64
	fmt.Println(distanceFromP(q))     // 5

	var origin Point
	fmt.Println(distanceFromP(origin)) // 2.23606797749979

	scaleP := p.ScaleBy        // method value
	fmt.Printf("%T\n", scaleP) // func(float64)
	fmt.Println(p)             // {x: 1, y: 2}
	scaleP(2)
	fmt.Println(p) // {x: 2, y: 4}
	scaleP(3)
	fmt.Println(p) // {x: 6, y: 12}
	scaleP(10)
	fmt.Println(p) // {x: 60, y: 120}

	r := new(Rocket)
	time.AfterFunc(3*time.Second, r.Launch) // 方法可以作为匿名函数
}
