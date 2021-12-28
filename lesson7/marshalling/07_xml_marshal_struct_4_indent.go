package main

import (
	"encoding/xml"
	"fmt"
)

// User1 struct represent one item to be marshalled into XML
type User1 struct {
	XMLName xml.Name `xml:"user"`
	id      uint32   `xml:"id"`
	name    string   `xml:"user_name"`
	surname string   `xml:"surname"`
}

// User2 struct represent one item to be marshalled into XML
type User2 struct {
	XMLName xml.Name `xml:"user"`
	Id      uint32   `xml:"id"`
	Name    string   `xml:"user_name"`
	Surname string   `xml:"surname"`
}

func main() {
	user1 := User1{
		id:      1,
		name:    "Pepek",
		surname: "Vyskoč"}

	user2 := User2{
		Id:      1,
		Name:    "Pepek",
		Surname: "Vyskoč"}

	user1asXML, _ := xml.MarshalIndent(user1, "", "    ")
	fmt.Println(string(user1asXML))

	fmt.Println()

	user2asXML, _ := xml.MarshalIndent(user2, "", "    ")
	fmt.Println(string(user2asXML))

	fmt.Println()

	user2asXML, _ = xml.MarshalIndent(user2, "", "\t")
	fmt.Println(string(user2asXML))

	fmt.Println()

	user2asXML, _ = xml.MarshalIndent(user2, "\t", "\t")
	fmt.Println(string(user2asXML))
}
