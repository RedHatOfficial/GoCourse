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

	p1, _ := stream.Contains(User{3, "Josef", "Vyskočil"})
	fmt.Printf("contains? %v\n", p1)

	p2, _ := stream.Contains(User{4, "Josef", "Vyskočil"})
	fmt.Printf("contains? %v\n", p2)
}
