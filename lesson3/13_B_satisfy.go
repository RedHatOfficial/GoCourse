package main

import "fmt"

// Adder is an interface with one method signature
type Adder interface {
	add(int, int) int
}

// AdderImpl is a user-defined data types that satisfy Added interface
type AdderImpl struct {
}

func (a AdderImpl) add(x int, y int) int {
	return x + y
}

func main() {
	a := AdderImpl{}
	fmt.Printf("Result is %d\n", a.add(1, 2))
}
