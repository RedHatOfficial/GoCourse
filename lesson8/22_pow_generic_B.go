package main

import (
	"fmt"
	"math"
)

type floats interface {
	float32 | float64
}

func pow[T floats](x T, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func main() {
	for x := float32(0.0); x < 5.0; x += 0.5 {
		result := pow(x, 2)
		fmt.Printf("%T %v\n", result, result)
	}

	fmt.Println()

	for x := 0.0; x < 5.0; x += 0.5 {
		result := pow(x, 2)
		fmt.Printf("%T %v\n", result, result)
	}
}
