package main

import "fmt"

func main() {
	var a [10]int

	slice := a[:]

	fmt.Printf("Array before modification: %v\n", a)
	fmt.Printf("Slice before modification: %v\n", slice)
	fmt.Println()

	for i := 0; i < len(a); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Array after modification:  %v\n", a)
	fmt.Printf("Slice after modification:  %v\n", slice)
	fmt.Println()

	for i := 0; i < len(slice); i++ {
		slice[i] = 42
	}

	fmt.Printf("Array after modification:  %v\n", a)
	fmt.Printf("Slice after modification:  %v\n", slice)
}
