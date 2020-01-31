package main

import (
	"fmt"
	"math"
)

// OpenShape represents an "open" 2D shape, such as line, arc, spline etc.
type OpenShape interface {
	length() float64
}

// ClosedShape represents a "closed" 2D shape, such as square, circle, and ellipse
type ClosedShape interface {
	area() float64
}

func length(shape OpenShape) float64 {
	return shape.length()
}

func area(shape ClosedShape) float64 {
	return shape.area()
}

// Line represents a line in 2D plane
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

type Circle struct {
	x, y   float64
	radius float64
}

type Ellipse struct {
	x, y float64
	a, b float64
}

type Rectangle struct {
	x, y          float64
	width, height float64
}

func (line Line) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (ellipse Ellipse) area() float64 {
	return math.Pi * ellipse.a * ellipse.b
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}
	fmt.Println("Line")
	fmt.Println(line1)
	fmt.Println(length(line1))
	fmt.Println(line1.length())
	fmt.Println()

	fmt.Println("Rectangle")
	r := Rectangle{x: 0, y: 0, width: 100, height: 100}
	fmt.Println(r)
	fmt.Println(area(r))
	fmt.Println(r.area())
	fmt.Println()

	// cont

	fmt.Println("Circle")
	c := Circle{x: 0, y: 0, radius: 100}
	fmt.Println(c)
	fmt.Println(area(c))
	fmt.Println(c.area())
	fmt.Println()

	fmt.Println("Ellipse")
	e := Ellipse{x: 0, y: 0, a: 100, b: 50}
	fmt.Println(e)
	fmt.Println(area(e))
	fmt.Println(e.area())
	fmt.Println()
}
