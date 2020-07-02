package main

import (
	"fmt"
	"github.com/fxamacker/cbor/v2"
)

func main() {
	var a bool = true

	var jsonOutput []byte

	cborOutput, _ = cbor.Marshal(a)
	fmt.Println(string(cborOutput))
}
