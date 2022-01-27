package main

import (
	"encoding/xml"
	"fmt"
)

// User struct represent one item to be marshalled into XML
type User struct {
	XMLName xml.Name `xml:"user"`
	ID      uint32   `xml:"id"`
	Name    string   `xml:"user_name"`
	Surname string   `xml:"surname"`
}

func main() {
	var users = [3]User{
		{
			ID:      1,
			Name:    "Pepek",
			Surname: "Vyskoč"},
		{
			ID:      2,
			Name:    "Pepek",
			Surname: "Vyskoč"},
		{
			ID:      3,
			Name:    "Josef",
			Surname: "Vyskočil"},
	}

	usersAsXML, _ := xml.MarshalIndent(users, "", "    ")
	fmt.Println(string(usersAsXML))
}
