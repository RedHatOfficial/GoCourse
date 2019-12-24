package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v := mat.NewSymDense(3, nil)
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(v))

	cap1, cap2 := v.Caps()
	fmt.Printf("Capacity:   %dx%d\n", cap1, cap2)

	dim1, dim2 := v.Dims()
	fmt.Printf("Dimensions: %dx%d\n", dim1, dim2)
}
