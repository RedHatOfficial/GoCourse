package main

import "fmt"

func f() int {
	return 5
}

func main() {
	var value int = 6
	switch i := 5; i {
	case value:
		fmt.Println(i, "= value")
	case f():
		fmt.Println(i, "= f()")
	case 4 + 1:
		fmt.Println(i, "= 4 + 1")
	}
}
