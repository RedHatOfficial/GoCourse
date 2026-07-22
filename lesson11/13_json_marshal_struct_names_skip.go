package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      uint32 `json:"-"`
	Name    string `json:"user_name"`
	Surname string `json:"surname"`
}

func main() {
	user := User{
		ID:      1,
		Name:    "John",
		Surname: "Doe"}

	userAsJSON, _ := json.MarshalIndent(user, "", "    ")
	fmt.Println(string(userAsJSON))
}
