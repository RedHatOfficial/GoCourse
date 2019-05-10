package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	user1 := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	fmt.Println(user1)

	user1.id = 2
	user1.name = "Steve"
	user1.surname = "Ballmer"

	fmt.Println(user1)
}
