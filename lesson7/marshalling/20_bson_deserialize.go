package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

type User struct {
	Id      uint32
	Name    string
	Surname string
}

func main() {
	var user User

	bsonInput, err := ioutil.ReadFile("1.bson")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read %d bytes\n", len(bsonInput))

	err = bson.Unmarshal(bsonInput, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Deserialized value")
	fmt.Println(user)
}
