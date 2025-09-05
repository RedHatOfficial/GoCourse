package main

import (
	"encoding/xml"
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

	userAsXML, _ := xml.Marshal(user)
	fmt.Println(string(userAsXML))
}
