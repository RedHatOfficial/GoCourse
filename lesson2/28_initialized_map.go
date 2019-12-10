package main

import "fmt"

func main() {
	m1 := make(map[int]string)
	fmt.Println(m1)

	m1[1] = "one"
	m1[2] = "two"
	m1[3] = "three"
	m1[4] = "four"

	fmt.Println(m1)
}
