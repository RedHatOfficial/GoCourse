package main

import "fmt"

func function(x int) {
	fmt.Printf("Defer %d\n", x)
}

func classify(x int) string {
	defer function(x)
	switch x {
	case 0:
		return "zero"
	case 2, 4, 6, 8:
		return "even number"
	case 1, 3, 5, 7, 9:
		return "odd number"
	default:
		return "?"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		println(x, classify(x))
	}
}
