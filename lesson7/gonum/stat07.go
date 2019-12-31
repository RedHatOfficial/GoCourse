package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	ys := []float64{-2, 0, 2, 4, 6}
	offset, slope := stat.LinearRegression(xs, ys, nil, false)
	fmt.Printf("Offset:  %f\n", offset)
	fmt.Printf("Slope:   %f\n", slope)
}
