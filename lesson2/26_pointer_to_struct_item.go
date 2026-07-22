package main

import "fmt"

// User is a new data type
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

	var pName *string
	var pSurname *string
	pName = &u.name
	pSurname = &u.surname

	fmt.Println(pName)
	fmt.Println(pSurname)
	fmt.Println(*pName)
	fmt.Println(*pSurname)
	fmt.Println("------------------")

	(*pName) = "Rob"
	(*pSurname) = "Pike"
	fmt.Println(*pName)
	fmt.Println(*pSurname)
}
