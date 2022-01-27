package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

// User data type contains all attributes about an user
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users = []User{
		{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}
	fmt.Println(users)

	fmt.Printf("input:  %v\n", users)

	stream := koazee.StreamOf(users)
	fmt.Printf("stream: %v\n", stream.Out().Val())
}
