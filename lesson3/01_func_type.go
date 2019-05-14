package main

import "fmt"

func function1() {
	fmt.Println("function1")
}

func function2() {
	fmt.Println("function2")
}

func main() {
	var a func()
	fmt.Println(a)

	a = function1
	fmt.Println(a)
	a()

	a = function2
	fmt.Println(a)
	a()
}
