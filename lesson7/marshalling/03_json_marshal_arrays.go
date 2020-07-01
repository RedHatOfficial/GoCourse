package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a1 [10]byte
	var a2 [10]int32
	a3 := [10]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}
	a4 := []string{"www", "root", "cz"}
	a5 := []interface{}{1, "root", 3.1415, true, []int{1, 2, 3, 4}}
	matice := [4][3]float32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{0, -1, 0},
	}

	a1_json, _ := json.Marshal(a1)
	fmt.Println(string(a1_json))

	a2_json, _ := json.Marshal(a2)
	fmt.Println(string(a2_json))

	a3_json, _ := json.Marshal(a3)
	fmt.Println(string(a3_json))

	a4_json, _ := json.Marshal(a4)
	fmt.Println(string(a4_json))

	a5_json, _ := json.Marshal(a5)
	fmt.Println(string(a5_json))

	matice_json, _ := json.Marshal(matice)
	fmt.Println(string(matice_json))
}
