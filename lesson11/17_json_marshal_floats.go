package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a float64 = 0.0
	var b float64 = 1e10

	var jsonOutput []byte

	jsonOutput, _ = json.Marshal(a)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(b)
	fmt.Println(string(jsonOutput))
}
