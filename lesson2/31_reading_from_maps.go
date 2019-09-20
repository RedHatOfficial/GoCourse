package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	fmt.Println(m1)

	m1["one"] = "I"
	m1["two"] = "II"
	m1["three"] = "III"
	m1["four"] = "IV"
	m1["five"] = "V"

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
