package main

import "fmt"

func main() {
	m1 := make(map[int]string)
	fmt.Println(m1)

	m1[0] = "one"
	m1[1] = "two"
	m1[2] = "three"
	m1[3] = "four"

	fmt.Println(m1)
}
