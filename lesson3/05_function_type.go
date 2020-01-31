package main

import (
	"fmt"
)

// BynaryOp represents any function that takes two int parameters and returns int as a result
type BinaryOp func(x, y int) int

func applyBinaryOp(a int, b int, operation BinaryOp) int {
	return operation(a, b)
}

func testBinaryOps(a int, b int) {
	f1 := func(a int, b int) int { return a + b }
	f2 := func(a int, b int) int { return a - b }
	f3 := func(a int, b int) int { return a * b }
	f4 := func(a int, b int) int { return a / b }
	fmt.Printf("%d + %d = %d\n", a, b, applyBinaryOp(a, b, f1))
	fmt.Printf("%d - %d = %d\n", a, b, applyBinaryOp(a, b, f2))
	fmt.Printf("%d * %d = %d\n", a, b, applyBinaryOp(a, b, f3))
	fmt.Printf("%d / %d = %d\n", a, b, applyBinaryOp(a, b, f4))
}

func main() {
	testBinaryOps(1, 2)
}
