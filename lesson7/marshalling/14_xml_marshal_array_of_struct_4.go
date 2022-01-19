package main

import (
	"encoding/xml"
	"fmt"
)

// User struct represent one item to be marshalled into XML
type User struct {
	ID      uint32 `xml:"id"`
	Name    string `xml:"user_name"`
	Surname string `xml:"surname"`
}

// Users struct represents list of items to be marshalled into XML
type Users struct {
	XMLName xml.Name `xml:"users"`
	List    []User   `xml:"user"`
}

func main() {
	var users Users = Users{
		List: []User{
			User{
				ID:      1,
				Name:    "Pepek",
				Surname: "Vyskoč"},
			User{
				ID:      2,
				Name:    "Pepek",
				Surname: "Vyskoč"},
			User{
				ID:      3,
				Name:    "Josef",
				Surname: "Vyskočil"},
		},
	}

	usersAsXML, _ := xml.MarshalIndent(users, "", "    ")
	fmt.Println(string(usersAsXML))
}
