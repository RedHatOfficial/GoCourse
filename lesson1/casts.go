package main

import "fmt"

func main() {
	a := float64(5)
	p := (*int)(nil)
	s := "hello"
	fmt.Printf("a: %f, p: %p, s: %s\n", a, p, s)
}
