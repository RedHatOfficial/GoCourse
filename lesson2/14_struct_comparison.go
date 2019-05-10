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

	var user2 User
	user2.id = 1
	user2.name = "Steve"
	user2.surname = "Ballmer"

	user3 := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	fmt.Println(user1 == user2)
	fmt.Println(user1 == user3)
}
