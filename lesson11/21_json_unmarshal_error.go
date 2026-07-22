package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      uint32 `json:"user_id"`
	Name    string `json:"user_name"`
	Surname string `json:"surname"`
}

func main() {
	userAsJSON := []byte("this is not proper JSON")

	var user2 User
	fmt.Println(user2)

	err := json.Unmarshal(userAsJSON, &user2)
	fmt.Println(err)
	fmt.Println(user2)
}
