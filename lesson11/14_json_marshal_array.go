package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := [10]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	aJSON, _ := json.MarshalIndent(a, "", "    ")
	fmt.Println(string(aJSON))
}
