package main

import "fmt"

type ID uint32
type Name string
type Surname string

func registerUser(id ID, name Name, surname Surname) {
	fmt.Printf("Registering: %d %s %s", id, name, surname)
}

func main() {
	var i ID = 1
	var n Name = "Jan"
	var s Surname = "Novák"

	registerUser(i, n, s)
}
