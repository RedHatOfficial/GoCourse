package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      uint32 `json:"user_id"`
	Name    string `json:"user_name"`
	surname string `json:"surname"`
}

func main() {
	user := User{
		ID:      1,
		Name:    "John",
		surname: "Doe"}

	fmt.Println(user)
	userAsJSON, _ := json.MarshalIndent(user, "", "    ")

	var user2 User
	fmt.Println(user2)

	err := json.Unmarshal(userAsJSON, &user2)
	fmt.Println(err)
	fmt.Println(user2)
}
