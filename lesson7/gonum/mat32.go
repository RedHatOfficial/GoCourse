package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	t1 := mat.NewTriDense(3, mat.Upper, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	transT1 := t1.T()
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t1))
	fmt.Printf("Trans:\n%v\n\n", mat.Formatted(transT1))

	fmt.Println()

	t2 := mat.NewTriDense(3, mat.Lower, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	transT2 := t2.T()
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(t2))
	fmt.Printf("Trans:\n%v\n\n", mat.Formatted(transT2))
}
