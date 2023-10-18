package main

import "C"

//export add
func add(x, y int) int {
	return x + y
}

func main() {}
