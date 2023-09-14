package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"os"
)

// Item represents value stored in binary tree node
type Item int

// Node represents one node of binary tree
type Node struct {
	Value Item
	Left  *Node
	Right *Node
}

// BinaryTree represents root node of binary tree
type BinaryTree struct {
	Root *Node
}

// Insert methods inserts new node into binary tree
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

		err = os.WriteFile("tree1.gob", content, 0o644)
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
