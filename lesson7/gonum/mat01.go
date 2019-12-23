package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	zero := mat.NewDense(5, 6, nil)
	fmt.Printf("%v\n", zero)
}
