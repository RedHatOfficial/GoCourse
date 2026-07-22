package main

import "fmt"

// User is a new data type
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	m1 := make(map[string]User)
	fmt.Println(m1)

	m1["1st"] = User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	m1["2nd"] = User{
		id:      2,
		name:    "Rob",
		surname: "Pike"}

	fmt.Println(m1)
}
