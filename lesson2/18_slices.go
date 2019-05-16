package main

import "fmt"

func main() {
	var a1 [100]byte
	var a2 [100]int32

	fmt.Printf("Array 1 length:   %d\n", len(a1))
	fmt.Printf("Array 2 length:   %d\n", len(a2))
	fmt.Println()

	var slice0 []byte = a1[:]
	fmt.Printf("Slice 0 length:   %d\n", len(slice0))
	fmt.Printf("Slice 0 capacity: %d\n", cap(slice0))
	fmt.Println()

	var slice1 []byte = a1[10:20]
	fmt.Printf("Slice 1 length:   %d\n", len(slice1))
	fmt.Printf("Slice 1 capacity: %d\n", cap(slice1))
	fmt.Println()

	// cont

	var slice2 = a1[20:30]
	fmt.Printf("Slice 2 length:   %d\n", len(slice2))
	fmt.Printf("Slice 2 capacity: %d\n", cap(slice2))
	fmt.Println()

	slice3 := a1[30:40]
	fmt.Printf("Slice 3 length:   %d\n", len(slice3))
	fmt.Printf("Slice 3 capacity: %d\n", cap(slice3))
}
