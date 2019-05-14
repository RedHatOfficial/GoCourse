package main

import "fmt"

func funkce1(x int, y int) int {
	return x + y
}

func funkce2(x int, y int) int {
	return x * y
}

func main() {
	var a func(int, int) int
	fmt.Println(a)

	a = funkce1
	fmt.Println(a)
	fmt.Println(a(10, 20))

	a = funkce2
	fmt.Println(a)
	fmt.Println(a(10, 20))
}
