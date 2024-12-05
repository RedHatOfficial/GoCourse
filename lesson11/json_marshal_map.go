package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{})
	m["foo"] = 10
	m["bar"] = "bar"
	m["baz"] = true
	m["xyzzy"] = []int{1, 2, 3}

	aJSON, _ := json.MarshalIndent(m, "", "    ")
	fmt.Println(string(aJSON))
}
