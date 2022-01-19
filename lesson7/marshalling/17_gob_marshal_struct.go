package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

// User struct represent one item to be marshalled into GOB
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

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(user)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("Encoded into %d bytes\n", len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)
	}
}
