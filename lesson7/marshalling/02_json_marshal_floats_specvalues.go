// https://stackoverflow.com/questions/1423081/json-left-out-infinity-and-nan-json-status-in-ecmascript

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

	jsonOutput, _ = json.Marshal(a)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(b)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(c)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(d)
	fmt.Println(string(jsonOutput))

	jsonOutput, _ = json.Marshal(e)
	fmt.Println(string(jsonOutput))
}
