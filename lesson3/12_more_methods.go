package main

import "fmt"

// Line represents a line in 2D plane
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func (line *Line) translate(dx, dy float64) {
	fmt.Printf("Translating line %v by %f %f\n", *line, dx, dy)
}

func (line *Line) rotate(angle float64) {
	fmt.Printf("Rotating line %v by %f\n", *line, angle)
}

func main() {
	line1 := Line{x1: 0, y1: 0, x2: 100, y2: 100}
	fmt.Println(line1)
	line1.translate(5, 5)
	fmt.Println(line1)
	line1.rotate(90)
	fmt.Println(line1)
}
