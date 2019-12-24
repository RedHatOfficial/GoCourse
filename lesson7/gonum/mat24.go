package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(m))

	fmt.Printf("Min: %f\n", mat.Min(m))
	fmt.Printf("Max: %f\n", mat.Max(m))
	fmt.Printf("Sum: %f\n", mat.Sum(m))
	fmt.Printf("Det: %f\n", mat.Det(m))
}
