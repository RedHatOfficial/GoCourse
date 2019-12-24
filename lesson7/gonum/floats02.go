package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{1, 5, 3, 5, 5, 0}

	fmt.Printf("x:   %v\n", x)
	fmt.Printf("y:   %v\n\n", y)
	fmt.Printf("Min index for x: %d\n", floats.MinIdx(x))
	fmt.Printf("Max index for x: %d\n\n", floats.MaxIdx(x))
	fmt.Printf("Min index for y: %d\n", floats.MinIdx(y))
	fmt.Printf("Max index for y: %d\n\n", floats.MaxIdx(y))
}
