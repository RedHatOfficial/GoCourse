package main

import "fmt"

func main() {
	if j := 10; j < 9 {
		fmt.Println(j, "< 9")
	} else {
		fmt.Println(j, "< 9 is false")
	}
}
