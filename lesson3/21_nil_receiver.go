package main

import (
	"fmt"
	"math"
)

type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func (line *Line) length() float64 {
	if line == nil {
		return 0
	}
	return math.Hypot(line.x1-line.x2, line.y1-line.y2)
}

func main() {
	var line *Line
	fmt.Printf("(%v, %T)\n", line, line)
	fmt.Println(line.length())
}
