package main

import "fmt"

type user struct {
	id   uint32
	name string
}

type registeredUser struct {
	user
	email string
}

func main() {
	registeredUser := registeredUser{
		user: user{
			id:   1,
			name: "Linus",
		},
		email: "linus@torvalds.com",
	}
	fmt.Println(registeredUser.name, registeredUser.email)
}
