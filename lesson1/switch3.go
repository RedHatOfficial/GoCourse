package main

import "fmt"

func main() {
	i := 5
	switch {
	case i < 5:
		fmt.Println(i, "< 5")
	case i > 5:
		fmt.Println(i, "> 5")
	default:
		fmt.Println(i, "is not less or more than 5")
	}
}
