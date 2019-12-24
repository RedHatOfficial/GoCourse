package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	t1 := mat.NewTriDense(3, mat.Upper, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t1))

	dim1, dim2 := t1.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)

	fmt.Println()

	t2 := mat.NewTriDense(3, mat.Lower, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t2))

	dim1, dim2 = t2.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
