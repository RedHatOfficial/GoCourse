package main

import "fmt"

func main() {
	var a1 [10]int

	a2 := a1

	fmt.Printf("Array 1: %v\n", a1)
	fmt.Printf("Array 2: %v\n", a2)

	for i := 0; i < len(a1); i++ {
		a1[i] = i * 2
	}

	fmt.Println("-------------------------------")

	fmt.Printf("Array 1: %v\n", a1)
	fmt.Printf("Array 2: %v\n", a2)
}
