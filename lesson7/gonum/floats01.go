package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}

	fmt.Printf("x:   %v\n", x)
	fmt.Printf("Min: %f\n", floats.Min(x))
	fmt.Printf("Max: %f\n", floats.Max(x))
	fmt.Printf("Sum: %f\n", floats.Sum(x))
}
