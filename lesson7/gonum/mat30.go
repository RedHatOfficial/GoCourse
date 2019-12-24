package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	t1 := mat.NewTriDense(3, mat.Upper, nil)
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t1))

	dim1, dim2 := t1.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)

	fmt.Println()

	t2 := mat.NewTriDense(3, mat.Lower, nil)
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t2))

	dim1, dim2 = t2.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
