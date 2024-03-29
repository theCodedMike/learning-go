// Package geometry defines simple types for plane geometry.
// package geometry
package main

import (
	"math"
)

type Point struct {
	X, Y float64
}

// Distance : traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance : same thing, but as a method of the Point type
// 这里的参数p，叫做方法的接收器（receiver）
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
