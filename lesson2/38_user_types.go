package main

import "fmt"

// ID is user-defined data type
type ID uint32

// Name is user-defined data type
type Name string

// Surnam is user-defined data type
type Surname string

func main() {
	var i ID
	i = 0
	fmt.Println(i)

	var n Name
	var s Surname
	n = "Jan"
	s = "Novák"
	fmt.Println(n)
	fmt.Println(s)
}
