package main

import "fmt"

func main() {
	var i int = 42

	fmt.Println(i)

	var p_i *int
	fmt.Println(p_i)

	p_i = &i

	fmt.Println(p_i)
	fmt.Println(*p_i)

	*p_i++

	fmt.Println(i)
	fmt.Println(*p_i)
}
