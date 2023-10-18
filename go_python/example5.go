package main

import "C"

import "fmt"

//export hello
func hello(name string) int {
	fmt.Printf("Hello %s\n", name)
	return len(name)
}

func main() {}
