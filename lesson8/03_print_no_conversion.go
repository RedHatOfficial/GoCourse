// Go also does not permit automatic type conversion
// (which is IMHO one of the best thing of Go)

package main

import "fmt"

type Value string

func printValue(value Value) {
	fmt.Println(value)
}

func main() {
	v := "www.root.cz" // string
	printValue(v)
}
