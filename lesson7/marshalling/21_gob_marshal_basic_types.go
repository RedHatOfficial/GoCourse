package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

func encodeAndDecode(msg string, value interface{}) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(value)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("%s value encoded into %d bytes\n", msg, len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)
	}
}

func main() {
	var b bool = true
	encodeAndDecode("Boolean", b)

	var x uint8 = 42
	encodeAndDecode("Uint8", x)

	var y uint16 = 42
	encodeAndDecode("Uint16", y)

	var z uint32 = 42
	encodeAndDecode("Uint32", z)
}
