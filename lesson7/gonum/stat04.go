package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math/rand"
)

func main() {
	ys := [100]float64{}
	for i := 0; i < len(ys); i++ {
		ys[i] = 50.0 - float64(i) + 2.0*rand.Float64() - 1.0
	}
	fmt.Printf("Mean:      %f\n", stat.Mean(ys[:], nil))
	fmt.Printf("Std.dev:   %f\n", stat.StdDev(ys[:], nil))
	fmt.Printf("Std.error: %f\n", stat.StdErr(stat.StdDev(ys[:], nil), 100))
}
