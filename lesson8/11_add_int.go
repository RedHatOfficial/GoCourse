// Non generic "add" function.

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
