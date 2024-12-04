package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user" json:"-"`
	ID      uint32   `xml:"id" json:"user_id"`
	Name    string   `xml:"name>first" json:"user_name"`
	Surname string   `xml:"name>last" json:"surname"`
}

func main() {
	user := User{
		ID:      1,
		Name:    "John",
		Surname: "Doe"}

	userAsXML, _ := xml.MarshalIndent(user, "", "    ")
	fmt.Println(string(userAsXML))

	fmt.Println()

	userAsJSON, _ := json.Marshal(user)
	fmt.Println(string(userAsJSON))
}
