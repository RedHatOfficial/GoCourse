package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

type Item int

type Node struct {
	Value Item
	Left  *Node
	Right *Node
}

func encodeStructureIntoBSON(s Node) ([]byte, error) {
	bsonOutput, err := bson.Marshal(s)

	if err != nil {
		return bsonOutput, err
	} else {
		return bsonOutput, nil
	}
}

func encodeStructureIntoJSON(s Node) ([]byte, error) {
	jsonOutput, err := json.Marshal(s)

	if err != nil {
		return jsonOutput, err
	} else {
		return jsonOutput, nil
	}
}

func encodeStructureIntoIndentedJSON(s Node) ([]byte, error) {
	jsonOutput, err := json.MarshalIndent(s, "", "    ")

	if err != nil {
		return jsonOutput, err
	} else {
		return jsonOutput, nil
	}
}

func encodeStructureIntoXML(s Node) ([]byte, error) {
	xmlOutput, err := xml.Marshal(s)

	if err != nil {
		return xmlOutput, err
	} else {
		return xmlOutput, nil
	}
}

func encodeStructureIntoIndentedXML(s Node) ([]byte, error) {
	xmlOutput, err := xml.MarshalIndent(s, "", "    ")

	if err != nil {
		return xmlOutput, err
	} else {
		return xmlOutput, nil
	}
}

func encodeStructureIntoGob(s Node) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(s)
	if err != nil {
		return buffer.Bytes(), err
	} else {
		return buffer.Bytes(), nil
	}
}

func save(encodedStructure []byte, filename string) {
	err := ioutil.WriteFile(filename, encodedStructure, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Stored into file", filename)
	}
}

func printBufferInfo(buffer []byte) {
	fmt.Println("\nBuffer with encoded structure: ", len(buffer))
}

func main() {
	var a, b, c Node

	a.Value = 1
	b.Value = 2
	c.Value = 3
	a.Left = &b
	a.Right = &c
	b.Left = nil
	b.Right = &c
	c.Left = nil
	c.Right = nil

	encodedStructure, err := encodeStructureIntoXML(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedStructure)
	save(encodedStructure, "structure1.xml")

	encodedStructure, err = encodeStructureIntoIndentedXML(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedStructure)
	save(encodedStructure, "structure2.xml")

	encodedStructure, err = encodeStructureIntoJSON(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedStructure)
	save(encodedStructure, "structure1.json")

	encodedStructure, err = encodeStructureIntoIndentedJSON(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedStructure)
	save(encodedStructure, "structure2.json")

	encodedStructure, err = encodeStructureIntoBSON(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedStructure)
	save(encodedStructure, "structure1.bson")

	encodedStructure, err = encodeStructureIntoGob(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedStructure)
	save(encodedStructure, "structure1.gob")
}
