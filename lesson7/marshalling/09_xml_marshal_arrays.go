package main

import (
	"encoding/xml"
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

	a1asXML, _ := xml.Marshal(a1)
	fmt.Println(string(a1asXML))

	a2asXML, _ := xml.Marshal(a2)
	fmt.Println(string(a2asXML))

	a3asXML, _ := xml.Marshal(a3)
	fmt.Println(string(a3asXML))

	a4asXML, _ := xml.Marshal(a4)
	fmt.Println(string(a4asXML))

	a5asXML, _ := xml.Marshal(a5)
	fmt.Println(string(a5asXML))

	maticeasXML, _ := xml.Marshal(matice)
	fmt.Println(string(maticeasXML))
}
