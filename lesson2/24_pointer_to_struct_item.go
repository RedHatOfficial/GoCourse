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
	fmt.Println("------------------")

	var p_n *string
	var p_s *string
	p_n = &u.name
	p_s = &u.surname

	fmt.Println(p_n)
	fmt.Println(p_s)
	fmt.Println(*p_n)
	fmt.Println(*p_s)
	fmt.Println("------------------")

	(*p_n) = "Steve"
	(*p_s) = "Ballmer"
	fmt.Println(*p_n)
	fmt.Println(*p_s)
}
