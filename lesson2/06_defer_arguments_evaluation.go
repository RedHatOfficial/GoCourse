package main

import "fmt"

func function(a []int) {
	fmt.Printf("Defer %v\n", a)
}

func main() {
	var x = []int{1, 2, 3}
	fmt.Printf("Current x value = %v\n", x)
	defer function(x)

	x[0] = 0
	fmt.Printf("Current x value = %v\n", x)
	defer function(x)

	x[1] = 0
	fmt.Printf("Current x value = %v\n", x)
	defer function(x)

	x[2] = 0
	fmt.Printf("Current x value = %v\n", x)
}
