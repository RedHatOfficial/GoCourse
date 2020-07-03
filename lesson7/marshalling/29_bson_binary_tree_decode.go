package main

import (
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

type BinaryTree struct {
	Root *Node
}

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

func encodeBinaryTree(bt BinaryTree) ([]byte, error) {
	bsonOutput, err := bson.Marshal(bt)

	if err != nil {
		return bsonOutput, err
	} else {
		return bsonOutput, nil
	}
}

func decodeBinaryTree(encodedTree []byte) (BinaryTree, error) {
	var newTree BinaryTree
	err := bson.Unmarshal(encodedTree, &newTree)
	return newTree, err
}

func saveBinaryTree(encodedTree []byte, filename string) {
	err := ioutil.WriteFile(filename, encodedTree, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Stored into file\n")
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

	printTree(bt.Root, 0)

	encodedTree, err := encodeBinaryTree(bt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nBuffer with encoded tree: ", len(encodedTree), "\n")

	saveBinaryTree(encodedTree, "tree.bson")

	decodedTree, err := decodeBinaryTree(encodedTree)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		printTree(decodedTree.Root, 0)
	}
}
