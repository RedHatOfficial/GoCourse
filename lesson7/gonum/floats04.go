package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{10, 20, 30, 40, 50, 60}
	z := make([]float64, len(x))

	fmt.Printf("x:   %v\n", x)
	fmt.Printf("y:   %v\n\n", y)

	floats.AddTo(z, x, y)
	fmt.Printf("x+y: %v\n", z)

	floats.SubTo(z, x, y)
	fmt.Printf("x-y: %v\n", z)

	floats.MulTo(z, x, y)
	fmt.Printf("x*y: %v\n", z)

	floats.DivTo(z, x, y)
	fmt.Printf("x/y: %v\n", z)
}
