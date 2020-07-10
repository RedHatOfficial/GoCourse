package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Printf("input: %v\n", values)

	stream := koazee.StreamOf(values)

	p1, _ := stream.Contains(2)
	fmt.Printf("contains 2? %v\n", p1)

	p2, _ := stream.Contains(42)
	fmt.Printf("contains 42? %v\n", p2)

	fmt.Println()

	words := []string{"one", "two", "three", "four", "five"}
	fmt.Printf("words: %v\n", words)

	stream = koazee.StreamOf(words)

	p3, _ := stream.Contains("one")
	fmt.Printf("contains 'one'? %v\n", p3)

	p4, e1 := stream.Contains("ten")
	fmt.Printf("contains 'ten'? %v\n", p4)
	fmt.Printf("error %v\n", e1)

	p4, e2 := stream.Contains(42)
	fmt.Printf("contains 42? %v\n", p4)
	fmt.Printf("error %v\n", e2)
}
