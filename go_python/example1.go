package main

import "C"
import "fmt"

//export hello
func hello() {
	fmt.Println("Hello, world!")
}

func main() {}
