package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
}

func printUser(u User) {
	fmt.Println(u)
}

func main() {
	user1 := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds"}

	printUser(user1)

	var user2 User
	user2.id = 1
	user2.name = "Rob"
	user2.surname = "Pike"

	printUser(user2)
}
