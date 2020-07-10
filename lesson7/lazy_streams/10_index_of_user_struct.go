package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users = []User{
		User{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}
	fmt.Println(users)

	stream := koazee.StreamOf(users)

	i1, _ := stream.IndexOf(User{3, "Josef", "Vyskočil"})
	fmt.Printf("index of #1: %v\n", i1)

	i2, _ := stream.LastIndexOf(User{3, "Josef", "Vyskočil"})
	fmt.Printf("last index of #1: %v\n", i2)

	i3, _ := stream.IndexOf(User{})
	fmt.Printf("index of #2: %v\n", i3)
}
