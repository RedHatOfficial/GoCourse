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
	var d mat.Dense
	var e mat.Dense

	printMatrix(&c, "c")
	printMatrix(&d, "d")
	printMatrix(&e, "e")

	m1 := mat.NewDense(3, 4, nil)
	m2 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})

	printMatrix(m1, "m1")
	printMatrix(m2, "m2")

	m3 := m2.T()

	printMatrix(m3, "m3")

	c.Add(m2, m2)
	printMatrix(&c, "c")

	d.Mul(m2, m3)
	printMatrix(&d, "d")

	e.MulElem(m3, m3)
	printMatrix(&e, "e")
}
