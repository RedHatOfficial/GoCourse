package main

import "C"

import "fmt"

//export add
func add(x, y int64) int64 {
	result := x + y
	fmt.Printf("Called add(%d, %d) with result %d\n", x, y, result)
	return result
}

func main() {}
