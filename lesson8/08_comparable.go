// The good old "compare" function defined for one data type.

package main

import "fmt"

func compare(x, y int) bool {
	return x < y
}

func main() {
	fmt.Println(compare(1, 2))
}
