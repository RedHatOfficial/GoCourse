package main

import (
	"fmt"
	"math"
)

type LineType struct {
	x1, y1 float64
	x2, y2 float64
}

func (line LineType) length() float64 {
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

func main() {
	line1 := LineType{x1: 0, y1: 0, x2: 100, y2: 100}
	fmt.Println(line1)
	lineLength := line1.length()
	fmt.Println(lineLength)
}
