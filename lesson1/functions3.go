package main

import "fmt"

func swap(a, b int) (x int, y int) {
	y = a
	x = b
	return
}

func main() {
	var a int = 3
	b := 4
	fmt.Println("a: ", a, ", b: ", b)
	a, b = swap(a, b)
	fmt.Println("a: ", a, ", b: ", b)
}
