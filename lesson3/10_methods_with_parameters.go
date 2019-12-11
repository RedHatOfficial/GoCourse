package main

import "fmt"

type LineType struct {
	x1, y1 float64
	x2, y2 float64
}

func (line LineType) translate(dx, dy float64) {
	fmt.Printf("Translating line %v by %f %f\n", line, dx, dy)
	line.x1 += dx
	line.y1 += dy
	line.x2 += dx
	line.y2 += dy
}

func main() {
	line1 := LineType{x1: 0, y1: 0, x2: 100, y2: 100}
	fmt.Println(line1)
	line1.translate(5, 5)
	fmt.Println(line1)
}
