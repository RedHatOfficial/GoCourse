package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	d := mat.NewDiagDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(d))

	dim1, dim2 := d.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
