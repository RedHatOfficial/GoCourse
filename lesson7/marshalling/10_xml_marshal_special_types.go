package main

import (
	"encoding/xml"
	"fmt"
	"math"
)

type Foobar struct {
	XMLName xml.Name `xml:"foobar"`
	Id      uint32   `xml:"id"`
	X       float64  `xml:"x"`
	Y       float64  `xml:"y"`
	Z       float64  `xml:"z"`
	Next    *Foobar  `xml:"foobar"`
}

func main() {
	f := Foobar{
		Id:   42,
		X:    math.NaN(),
		Y:    math.Inf(1),
		Z:    math.Inf(-1),
		Next: nil}

	g := Foobar{
		Id:   43,
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
