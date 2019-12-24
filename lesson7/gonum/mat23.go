package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v := mat.NewSymDense(3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	t := v.T()

	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t))

	dim1, dim2 := t.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
