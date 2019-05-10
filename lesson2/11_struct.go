package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var user1 User

	fmt.Println(user1)

	user1.id = 1
	user1.name = "Linus"
	user1.surname = "Torvalds"

	fmt.Println(user1)
}
