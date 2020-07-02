package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Id      uint32   `xml:"id"`
	Name    string   `xml:"user_name"`
	Surname string   `xml:"surname"`
}

type Users struct {
	XMLName xml.Name `xml:"users"`
	List    []User
}

func main() {
	var users Users = Users{
		List: []User{
			User{
				Id:      1,
				Name:    "Pepek",
				Surname: "Vyskoč"},
			User{
				Id:      2,
				Name:    "Pepek",
				Surname: "Vyskoč"},
			User{
				Id:      3,
				Name:    "Josef",
				Surname: "Vyskočil"},
		},
	}

	usersAsXML, _ := xml.MarshalIndent(users, "", "    ")
	fmt.Println(string(usersAsXML))
}
