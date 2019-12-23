package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func printMatrix(m mat.Matrix, name string) {
	fmt.Printf(" *** %s ***\n", name)
	fmt.Println(mat.Formatted(m))
	fmt.Println()

}

func main() {
	var c mat.Dense

	m2 := mat.NewDense(3, 3, []float64{5, 0, 0, 0, 2, 0, 0, 0, 1})

	printMatrix(m2, "m2")

	c.Inverse(m2)
	printMatrix(&c, "c")

	var d mat.Dense
	d.Mul(&c, m2)
	printMatrix(&d, "d")
}
