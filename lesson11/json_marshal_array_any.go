package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := []interface{}{1, "Red Hat", 3.1415, true, []int{1, 2, 3, 4}}

	aJSON, _ := json.MarshalIndent(a, "", "    ")
	fmt.Println(string(aJSON))
}
