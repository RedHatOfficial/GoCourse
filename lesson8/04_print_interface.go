// "Generics" implemented in runtime.

// Not very efficient nor elegant, but doable in a way

package main

import "fmt"

func printValue(value interface{}) {
	fmt.Println(value)
}

func main() {
	printValue("www.root.cz")
	printValue('*')
	printValue(42)
	printValue(3.14)
	printValue(1 + 2i)
	printValue([]int{1, 2, 3})
}
