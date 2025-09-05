package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/ugorji/go/codec"
)

// Item represents value stored in binary tree
type Item int

// Node represents one node in binary tree
type Node struct {
	Value Item
	Left  *Node
	Right *Node
}

// BinaryTree is a root node of binary tree
type BinaryTree struct {
	Root *Node
}

// Insert method inserts new value into binary tree
func (bt *BinaryTree) Insert(value Item) {
	node := &Node{value, nil, nil}
	if bt.Root == nil {
		bt.Root = node
	} else {
		insertNode(bt.Root, node)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

func printTree(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printTree(node.Left, level)
		fmt.Printf(format+"%d\n", node.Value)
		printTree(node.Right, level)
	}
}

func encodeBinaryTreeIntoJSON(bt BinaryTree) ([]byte, error) {
	jsonOutput, err := json.Marshal(bt)

	if err != nil {
		return jsonOutput, err
	}
	return jsonOutput, nil
}

func encodeBinaryTreeIntoIndentedJSON(bt BinaryTree) ([]byte, error) {
	jsonOutput, err := json.MarshalIndent(bt, "", "    ")

	if err != nil {
		return jsonOutput, err
	}
	return jsonOutput, nil
}

func encodeBinaryTreeIntoXML(bt BinaryTree) ([]byte, error) {
	xmlOutput, err := xml.Marshal(bt)

	if err != nil {
		return xmlOutput, err
	}
	return xmlOutput, nil
}

func encodeBinaryTreeIntoIndentedXML(bt BinaryTree) ([]byte, error) {
	xmlOutput, err := xml.MarshalIndent(bt, "", "    ")

	if err != nil {
		return xmlOutput, err
	}
	return xmlOutput, nil
}

func encodeBinaryTreeIntoGob(bt BinaryTree) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(bt)
	if err != nil {
		return buffer.Bytes(), err
	}
	return buffer.Bytes(), nil
}

func encodeBinaryTreeIntoMsgPack(bt BinaryTree) ([]byte, error) {
	var buffer bytes.Buffer

	// handler
	var handler codec.MsgpackHandle

	// objekt realizující zakódování dat
	encoder := codec.NewEncoder(&buffer, &handler)

	// zakódování dat
	err := encoder.Encode(bt)
	if err != nil {
		return buffer.Bytes(), err
	}
	return buffer.Bytes(), nil
}

func saveBinaryTree(encodedTree []byte, filename string) {
	err := os.WriteFile(filename, encodedTree, 0o644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Stored into file", filename)
	}
}

func constructTree(bt *BinaryTree, min, max int) {
	middle := (min + max) / 2
	if min < middle && middle < max {
		fmt.Println(middle)
		bt.Insert(Item(middle))
		constructTree(bt, min, middle)
		constructTree(bt, middle, max)
	}
}

func printBufferInfo(buffer []byte) {
	fmt.Println("\nBuffer with encoded tree: ", len(buffer))
}

func main() {
	var bt BinaryTree
	constructTree(&bt, 0, 256)

	printTree(bt.Root, 0)

	encodedTree, err := encodeBinaryTreeIntoXML(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedTree)
	saveBinaryTree(encodedTree, "tree1.xml")

	encodedTree, err = encodeBinaryTreeIntoIndentedXML(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedTree)
	saveBinaryTree(encodedTree, "tree2.xml")

	encodedTree, err = encodeBinaryTreeIntoJSON(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedTree)
	saveBinaryTree(encodedTree, "tree1.json")

	encodedTree, err = encodeBinaryTreeIntoIndentedJSON(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedTree)
	saveBinaryTree(encodedTree, "tree2.json")

	encodedTree, err = encodeBinaryTreeIntoGob(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedTree)
	saveBinaryTree(encodedTree, "tree1.gob")

	encodedTree, err = encodeBinaryTreeIntoMsgPack(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedTree)
	saveBinaryTree(encodedTree, "/tmp/tree1.bin")

}
