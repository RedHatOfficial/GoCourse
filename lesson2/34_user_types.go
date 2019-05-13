package main

import "fmt"

type Id uint32
type Name string
type Surname string

func main() {
	var i Id
	i = 0
	fmt.Println(i)

	var n Name
	var s Surname

	n = "Jan"
	s = "Novák"

	fmt.Println(n)
	fmt.Println(s)
}
