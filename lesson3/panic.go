package main

import "fmt"

func bar() {
	var a1 []byte
	fmt.Println(a1[1000])
}

func foo() {
	bar()
}

func main() {
	foo()
}
