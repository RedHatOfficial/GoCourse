package main

import "fmt"

func main() {
	var matrix [4][3]float32
	fmt.Printf("Matrix:    %v\n", matrix)

	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[j]); i++ {
			matrix[j][i] = 1.0 * float32(i+1) * float32(j+1)
		}
	}
	fmt.Printf("Matrix:    %v\n", matrix)
}
