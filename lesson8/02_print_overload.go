// Go does not permit overriding functions i.e.
// having two or more functions with the same name
// in one module/package.

// It avoids ugly "name mangling" C++ problem

package main

import "fmt"

func printValue(value string) {
	fmt.Println(value)
}

func printValue(value int) {
	fmt.Println(value)
}

func main() {
	printValue("www.root.cz")
}
