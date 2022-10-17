package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var a int = 3
	b := 4
	fmt.Println("a: ", a, ", b: ", b)
	x, y := swap(a, b)
	fmt.Println("x: ", x, ", y: ", y)
}
