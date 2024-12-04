package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	ID      uint32 `xml:"id"`
	Name    string `xml:"user_name"`
	Surname string `xml:"surname"`
}

func main() {
	user := User{
		ID:      1,
		Name:    "John",
		Surname: "Doe"}

	userAsXML, _ := xml.MarshalIndent(user, "", "    ")
	fmt.Println(string(userAsXML))
}
