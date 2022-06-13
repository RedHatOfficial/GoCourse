// Type parameters introduced to Go in version 1.18
// -> generic function!

// Now the function printValue accepts value of any type

// (in Scala - any is "top type", none is "bottom type")

package main

import "fmt"

func printValue[T any](value T) {
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
