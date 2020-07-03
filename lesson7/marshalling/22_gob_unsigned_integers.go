package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

func encodeAndDecodeUint(value uint64) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(value)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("%20d value encoded into %d bytes: ", value, len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)
	}
}

func main() {
	var value uint64 = 1
	for i := 0; i < 64; i++ {
		encodeAndDecodeUint(value)
		value <<= 1
	}
}
