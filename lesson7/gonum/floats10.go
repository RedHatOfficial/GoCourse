package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{1, 2, 3, 4, 5, 6}
	z := []float64{1.0, 2.2, 2.8, 4.2, 4.8, 6.0}

	fmt.Printf("x:   %v\n", x)
	fmt.Printf("y:   %v\n", y)
	fmt.Printf("z:   %v\n", z)

	fmt.Printf("x~=y (±0,10)?: %t\n", floats.EqualApprox(x, y, 0.1))
	fmt.Printf("x~=z (±0,05)?: %t\n", floats.EqualApprox(x, z, 0.05))
	fmt.Printf("x~=z (±0,09)?: %t\n", floats.EqualApprox(x, z, 0.09))
	fmt.Printf("x~=z (±0,10)?: %t\n", floats.EqualApprox(x, z, 0.10))
	fmt.Printf("x~=z (±0,11)?: %t\n", floats.EqualApprox(x, z, 0.11))
}
