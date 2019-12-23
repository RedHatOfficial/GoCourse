package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v := mat.NewVecDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("Value:\n%v\n", mat.Formatted(v))
	fmt.Printf("Length:     %d\n", v.Len())
	fmt.Printf("Capacity:   %d\n", v.Cap())
	dim1, dim2 := v.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
