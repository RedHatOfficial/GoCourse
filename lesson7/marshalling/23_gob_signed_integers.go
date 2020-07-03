package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

func encodeAndDecodeInt(value int64) {
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
	var value int64 = -1
	for i := 0; i < 63; i++ {
		encodeAndDecodeInt(value)
		value <<= 1
	}
}
