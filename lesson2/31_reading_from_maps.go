package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Println(m1)

	m1["one"] = 0
	m1["two"] = 1
	m1["three"] = 2

	fmt.Println(m1)

	// cont

	value, exist := m1["two"]

	if exist {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	value, exist = m1["thousand"]

	if exist {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}
}
