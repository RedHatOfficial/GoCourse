package main

import (
	"fmt"
	"math"
)

type OpenShape interface {
	length() float64
}

type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

type Point struct {
	x, y float64
}

func (point Point) length() float64 {
	return 0
}

func do_something(shape OpenShape) {
	s := shape.(Point)
	fmt.Println(s.x, s.y)
}

func main() {
	p := Point{x: 10, y: 20}
	do_something(p)

	l := Line{x1: 0, y1: 0, x2: 10, y2: 20}
	do_something(l)
}
