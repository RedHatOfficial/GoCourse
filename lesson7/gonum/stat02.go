package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	values := []float64{1, 2, 3, 4, 5}
	fmt.Printf("Geometric mean: %f\n", stat.GeometricMean(values, nil))

	weights := []float64{1, 1, 1, 1, 1}
	fmt.Printf("Weighted geometric mean #1: %f\n", stat.GeometricMean(values, weights))

	weights2 := []float64{1, 1, 1, 1, 6}
	fmt.Printf("Weighted geometric mean #2: %f\n", stat.GeometricMean(values, weights2))
}
