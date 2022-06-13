// Strong type check applies there!

package main

import "fmt"

func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	printValue[int]("www.root.cz")
	printValue[[]string]('*')
	printValue[string](42)
	printValue[int](3.14)
	printValue[byte](1 + 2i)
	printValue[[]byte]([]int{1, 2, 3})
}
