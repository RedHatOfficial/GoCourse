package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	var a float64 = -0.0
	var b float64 = math.NaN()
	var c float64 = -math.NaN()
	var d float64 = math.Inf(1)
	var e float64 = math.Inf(-1)

	var jsonOutput []byte
	var err error

	jsonOutput, err = json.Marshal(a)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(b)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(c)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(d)
	fmt.Println(err, string(jsonOutput))

	jsonOutput, err = json.Marshal(e)
	fmt.Println(err, string(jsonOutput))
}
