package main

import (
	"fmt"
	"math"
)

// OpenShape represents an "open" 2D shape, such as line, arc, spline etc.
type OpenShape interface {
	length() float64
}

// Line represents a line in 2D plane
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func length(shape OpenShape) float64 {
	return shape.length()
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}

	fmt.Println(line1)

	lineLength := length(line1)
	fmt.Println(lineLength)
}
