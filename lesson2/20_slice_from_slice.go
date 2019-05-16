package main

import "fmt"

func main() {
	var a [10]int

	slice1 := a[4:9]
	slice2 := slice1[3:]

	fmt.Printf("Array:            %v\n", a)
	fmt.Printf("Array length:     %d\n\n", len(a))

	fmt.Printf("Slice 1:          %v\n", slice1)
	fmt.Printf("Slice 1 length:   %d\n", len(slice1))
	fmt.Printf("Slice 1 capacity: %d\n\n", cap(slice1))

	fmt.Printf("Slice 2:          %v\n", slice2)
	fmt.Printf("Slice 2 length:   %d\n", len(slice2))
	fmt.Printf("Slice 2 capacity: %d\n\n", cap(slice2))

	// cont

	slice2[0] = 99
	slice2[1] = 99

	fmt.Printf("Array:            %v\n", a)
	fmt.Printf("Slice 1:          %v\n", slice1)
	fmt.Printf("Slice 2:          %v\n", slice2)
}
