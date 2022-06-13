// It is possible to specify which variant of generic function to use
// during function call

// (not needed IMHO, but interesting)

package main

import "fmt"

func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	printValue[string]("www.root.cz")
	printValue[rune]('*')
	printValue[int](42)
	printValue[float32](3.14)
	printValue[complex64](1 + 2i)
	printValue[[]int]([]int{1, 2, 3})
}
