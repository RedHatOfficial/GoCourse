package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math/rand"
)

const values = 10

func main() {
	xs := [values]float64{}
	ys := [values]float64{}

	for i := 0; i < len(xs); i++ {
		xs[i] = float64(i)
	}

	for i := 0; i < len(ys); i++ {
		ys[i] = 50.0 - float64(i) + 2.0*rand.Float64() - 1.0
	}

	offset, slope := stat.LinearRegression(xs[:], ys[:], nil, false)
	fmt.Printf("Offset:  %f\n", offset)
	fmt.Printf("Slope:   %f\n", slope)
}
