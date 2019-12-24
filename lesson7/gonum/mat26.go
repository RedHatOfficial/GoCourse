package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("Value:\n%v\n\n", mat.Formatted(m))

	for column := 0; column < 3; column++ {
		fmt.Printf("Column #%d:\n%v\n\n", column, mat.Col(nil, column, m))
	}

	for row := 0; row < 3; row++ {
		fmt.Printf("Row #%d:\n%v\n\n", row, mat.Row(nil, row, m))
	}
}
