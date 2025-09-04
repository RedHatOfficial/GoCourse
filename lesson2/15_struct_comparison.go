package main

import "fmt"

// User is a new data type
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
	user2.name = "Rob"
	user2.surname = "Pike"

	user3 := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	fmt.Println(user1 == user2)
	fmt.Println(user1 == user3)
}
