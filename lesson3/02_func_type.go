package main

import "fmt"

func function1(x, y int) int {
	return x + y
}

func function2(x, y int) int {
	return x * y
}

func main() {
	var a func(int, int) int
	fmt.Println(a)

	a = function1
	fmt.Println(a)
	fmt.Println(a(10, 20))

	a = function2
	fmt.Println(a)
	fmt.Println(a(10, 20))
}
