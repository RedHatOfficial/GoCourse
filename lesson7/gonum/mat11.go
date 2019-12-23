package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v := mat.NewVecDense(10, nil)
	fmt.Printf("Value:\n%v\n", mat.Formatted(v))
	fmt.Printf("Length:     %d\n", v.Len())
	fmt.Printf("Capacity:   %d\n", v.Cap())
	dim1, dim2 := v.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
