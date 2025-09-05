package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"new-user"`
	ID      uint32   `xml:"id"`
	Name    string   `xml:"name>first"`
	Surname string   `xml:"name>last"`
}

func main() {
	user := User{
		ID:      1,
		Name:    "John",
		Surname: "Doe"}

	userAsXML, _ := xml.MarshalIndent(user, "", "    ")
	fmt.Println(string(userAsXML))
}
