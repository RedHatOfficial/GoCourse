// Generic "add" function for three distinct data types.

package main

import "fmt"

type numeric interface {
	int | float64 | complex128
}

func add[T numeric](x T, y T) T {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.5, 2.6))
	fmt.Println(add(1i, 2+4i))
}
