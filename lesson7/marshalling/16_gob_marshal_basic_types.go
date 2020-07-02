package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

func main() {
	var a bool = true

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(a)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("Encoded into %d bytes\n", len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)
	}
}
