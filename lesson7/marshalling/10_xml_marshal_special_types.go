package main

import (
	"encoding/xml"
	"fmt"
	"math"
)

// Foobar is a data structure to be serialized into XML
type Foobar struct {
	XMLName xml.Name `xml:"foobar"`
	ID      uint32   `xml:"id"`
	X       float64  `xml:"x"`
	Y       float64  `xml:"y"`
	Z       float64  `xml:"z"`
	Next    *Foobar  `xml:"foobar"`
}

func main() {
	f := Foobar{
		ID:   42,
		X:    math.NaN(),
		Y:    math.Inf(1),
		Z:    math.Inf(-1),
		Next: nil}

	g := Foobar{
		ID:   43,
		X:    math.NaN(),
		Y:    math.Inf(1),
		Z:    math.Inf(-1),
		Next: &f}

	asXML, err := xml.MarshalIndent(g, "", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(asXML))
	}
}
