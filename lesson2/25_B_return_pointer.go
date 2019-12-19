package main

import "fmt"

func getPointer() *int {
	var i int = 42
	return &i
}

func main() {
	p := getPointer()
	fmt.Printf("%v\n", p)
	fmt.Printf("%d\n", *p)
}
