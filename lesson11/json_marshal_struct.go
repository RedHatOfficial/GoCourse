package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	user := User{
		ID:      1,
		Name:    "John",
		Surname: "Doe"}

	userAsJSON, _ := json.Marshal(user)
	fmt.Println(string(userAsJSON))
}
