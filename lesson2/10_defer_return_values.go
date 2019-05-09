package main

import "fmt"

func function1() (i int) {
	i = 1
	return
}

func function2() (i int) {
	defer func() { i = 2 }()
	return 1
}

func function3() (i int) {
	defer func() { i += 2 }()
	return 1
}

func main() {
	fmt.Printf("Return value of function1: %d\n", function1())
	fmt.Printf("Return value of function2: %d\n", function2())
	fmt.Printf("Return value of function3: %d\n", function3())
}
