package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	user := User{
		1,
		"John",
		"Doe"}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Encoded size: %d bytes\n", buffer.Len())

	decoder := gob.NewDecoder(&buffer)

	var user2 User

	err = decoder.Decode(&user2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user2)
}
