package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	user1 := User{
		1,
		"Linus",
		"Torvalds"}

	fmt.Println(user1)

	user1.id = 2
	user1.name = "Rob"
	user1.surname = "Pike"

	fmt.Println(user1)
}
