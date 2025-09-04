package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i < 11; i++ {
		slice = append(slice, i)
		fmt.Printf("Slice:          %v\n", slice)
		fmt.Printf("Slice length:   %d\n", len(slice))
		fmt.Printf("Slice capacity: %d\n\n", cap(slice))
	}
}
