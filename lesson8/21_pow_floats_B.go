package main

import (
	"fmt"
	"math"
)

func pow32(x, y float32) float32 {
	return float32(math.Pow(float64(x), float64(y)))
}

func pow64(x, y float64) float64 {
	return math.Pow(x, y)
}

func main() {
	for x := float32(0.0); x < 5.0; x += 0.5 {
		result := pow32(x, 2)
		fmt.Printf("%T %v\n", result, result)
	}

	fmt.Println()

	for x := 0.0; x < 5.0; x += 0.5 {
		result := pow64(x, 2)
		fmt.Printf("%T %v\n", result, result)
	}
}
