package main

import "fmt"

func main() {
	var i int = 5

	if i < 6 {
		fmt.Println(i, "< 6")
	}

	if j := 10; j < 11 {
		fmt.Println(j, "< 11")
	}
}
