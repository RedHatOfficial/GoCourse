package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
)

// User struct represent one item to be marshalled into BSON
type User struct {
	ID      uint32
	Name    string
	Surname string
}

func main() {
	user := User{
		1,
		"Pepek",
		"Vyskoč"}

	var bsonOutput []byte

	bsonOutput, err := bson.Marshal(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Encoded into %d bytes\n", len(bsonOutput))
		err := os.WriteFile("1.bson", bsonOutput, 0644)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("And stored into file")
		}
	}
}
