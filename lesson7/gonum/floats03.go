package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}

	fmt.Printf("x:        %v\n", x)

	floats.Reverse(x)
	fmt.Printf("reversed: %v\n", x)

	floats.Scale(10, x)
	fmt.Printf("scaled:   %v\n", x)
}
