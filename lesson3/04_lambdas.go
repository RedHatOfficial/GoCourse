package main

import (
	"fmt"
)

func testBinaryOps(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, func(a int, b int) int { return a + b }(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, func(a int, b int) int { return a - b }(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, func(a int, b int) int { return a * b }(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, func(a int, b int) int { return a / b }(a, b))
}

func main() {
	testBinaryOps(1, 2)
}
