package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"
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
		err := os.WriteFile("1.bson", bsonOutput, 0o644)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("And stored into file")
		}
	}
}
