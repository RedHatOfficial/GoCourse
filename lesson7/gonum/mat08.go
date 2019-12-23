package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	var c mat.Dense
	var d mat.Dense
	var e mat.Dense

	fmt.Println(" *** c ***")
	fmt.Println(mat.Formatted(&c))
	fmt.Println()

	fmt.Println(" *** d ***")
	fmt.Println(mat.Formatted(&d))
	fmt.Println()

	fmt.Println(" *** e ***")
	fmt.Println(mat.Formatted(&d))
	fmt.Println()

	m1 := mat.NewDense(3, 4, nil)
	m2 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})

	fmt.Println(" *** m1 ***")
	fmt.Println(mat.Formatted(m1))
	fmt.Println()

	fmt.Println("  *** m2 ***")
	fmt.Println(mat.Formatted(m2))
	fmt.Println()

	m3 := m2.T()

	fmt.Println("  *** m3 ***")
	fmt.Println(mat.Formatted(m3))
	fmt.Println()

	c.Add(m2, m2)
	fmt.Println("   *** c ***")
	fmt.Println(mat.Formatted(&c))
	fmt.Println()

	d.Mul(m2, m3)
	fmt.Println("   *** d ***")
	fmt.Println(mat.Formatted(&d))
	fmt.Println()

	e.MulElem(m3, m3)
	fmt.Println("   *** e ***")
	fmt.Println(mat.Formatted(&e))
	fmt.Println()
}
