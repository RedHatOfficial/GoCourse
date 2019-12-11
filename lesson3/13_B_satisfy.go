package main

import "fmt"

type Adder interface {
	add(int, int) int
}

type AdderImpl struct {
}

func (a AdderImpl) add(x int, y int) int {
	return x + y
}

func main() {
	a := AdderImpl{}
	fmt.Printf("Result is %d\n", a.add(1, 2))
}
