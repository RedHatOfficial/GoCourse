package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	var i int = 10
	for i < 11 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("Break the loop!")
		break
	}
}
