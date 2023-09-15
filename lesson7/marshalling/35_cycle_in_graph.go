package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"
)

// Item represents value to be stored in linked list
type Item int

// Node represents one node in linked list
type Node struct {
	Value Item
	Next  *Node
}

func encodeStructureIntoBSON(s Node) ([]byte, error) {
	bsonOutput, err := bson.Marshal(s)

	if err != nil {
		return bsonOutput, err
	}
	return bsonOutput, nil
}

func encodeStructureIntoJSON(s Node) ([]byte, error) {
	jsonOutput, err := json.Marshal(s)

	if err != nil {
		return jsonOutput, err
	}
	return jsonOutput, nil
}

func encodeStructureIntoIndentedJSON(s Node) ([]byte, error) {
	jsonOutput, err := json.MarshalIndent(s, "", "    ")

	if err != nil {
		return jsonOutput, err
	}
	return jsonOutput, nil
}

func encodeStructureIntoXML(s Node) ([]byte, error) {
	xmlOutput, err := xml.Marshal(s)

	if err != nil {
		return xmlOutput, err
	}
	return xmlOutput, nil
}

func encodeStructureIntoIndentedXML(s Node) ([]byte, error) {
	xmlOutput, err := xml.MarshalIndent(s, "", "    ")

	if err != nil {
		return xmlOutput, err
	}
	return xmlOutput, nil
}

func encodeStructureIntoGob(s Node) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(s)
	if err != nil {
		return buffer.Bytes(), err
	}
	return buffer.Bytes(), nil
}

func save(encodedStructure []byte, filename string) {
	err := os.WriteFile(filename, encodedStructure, 0o644)
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
	var a, b Node

	a.Value = 1
	b.Value = 2
	a.Next = &b
	b.Next = &a

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
