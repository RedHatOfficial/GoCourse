package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

type Item int

type Node struct {
	value Item
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) Insert(value Item) {
	node := &Node{value, nil, nil}
	if bt.root == nil {
		bt.root = node
	} else {
		insertNode(bt.root, node)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func encodeAndDecodeBinaryTree(bt BinaryTree) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(bt)
	if err != nil {
		fmt.Println(err)
	} else {
		content := buffer.Bytes()
		fmt.Printf("Binary tree encoded into %d bytes: ", len(content))
		encoded := hex.EncodeToString(content)
		fmt.Println(encoded)

		err = ioutil.WriteFile("tree1.gob", content, 0644)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("And stored into file")
		}
	}
}

func main() {
	var bt BinaryTree
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(1)
	bt.Insert(4)
	bt.Insert(6)
	bt.Insert(8)
	bt.Insert(9)
	bt.Insert(10)
	bt.Insert(0)

	encodeAndDecodeBinaryTree(bt)
}
