package main

import "fmt"

// Key is a new data type
type Key struct {
	id   uint32
	role string
}

// User is a new data type
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	m1 := make(map[Key]User)
	fmt.Println(m1)

	m1[Key{1, "root"}] = User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	m1[Key{2, "gopher"}] = User{
		id:      2,
		name:    "Rob",
		surname: "Pike"}

	fmt.Println(m1)

	delete(m1, Key{1, "root"})
	fmt.Println(m1)

	delete(m1, Key{1000, "somebody else"})
	fmt.Println(m1)
}
