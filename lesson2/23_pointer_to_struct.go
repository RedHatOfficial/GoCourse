package main

import "fmt"

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var u User

	u.id = 1
	u.name = "Linus"
	u.surname = "Torvalds"

	fmt.Println(u)

	var p_u *User
	p_u = &u

	fmt.Println(p_u)
	fmt.Println(*p_u)

	(*p_u).id = 10000
	fmt.Println(*p_u)

	p_u.id = 20000
	fmt.Println(*p_u)
}
