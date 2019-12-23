package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	v2 := mat.NewVecDense(5, []float64{1, 0, 2, 0, 3})
	v := mat.NewVecDense(5, nil)

	v.ScaleVec(10.0, v2)

	fmt.Printf("v2:\n%v\n\n", mat.Formatted(v2))
	fmt.Printf("v:\n%v\n\n", mat.Formatted(v))
}
