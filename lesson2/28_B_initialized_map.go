package main

import "fmt"

func main() {
	var m1 map[string]int = make(map[string]int)
	fmt.Println(m1)

	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3
	m1["four"] = 4

	fmt.Println(m1)
}
