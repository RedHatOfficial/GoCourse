package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	big := mat.NewDense(100, 100, nil)
	for i := 0; i < 100; i++ {
		big.Set(i, i, 1)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(mat.Formatted(big, mat.Prefix(" "), mat.Excerpt(i)))
		fmt.Println()
	}
}
