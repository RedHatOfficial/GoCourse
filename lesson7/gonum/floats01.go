package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}

	fmt.Printf("x:    %v\n", x)
	fmt.Printf("Min:  %5.0f\n", floats.Min(x))
	fmt.Printf("Max:  %5.0f\n", floats.Max(x))
	fmt.Printf("Sum:  %5.0f\n", floats.Sum(x))
	fmt.Printf("Prod: %5.0f\n", floats.Prod(x))
}
