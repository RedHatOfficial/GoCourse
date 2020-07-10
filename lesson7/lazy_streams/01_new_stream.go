package main

import (
	"fmt"
	"github.com/wesovilabs/koazee/stream"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Printf("input:   %v\n", values)

	stream1 := stream.New(values)
	fmt.Printf("stream:   %v\n", stream1.Out().Val())

	fmt.Println()

	words := []string{"one", "two", "three", "four", "five"}
	fmt.Printf("words:    %v\n", words)

	stream2 := stream.New(words)
	fmt.Printf("stream:   %v\n", stream2.Out().Val())

	fmt.Println()

	anything := []interface{}{"one", 1, 1.0, 1 + 1i}
	fmt.Printf("anything: %v\n", anything)

	stream3 := stream.New(anything)
	fmt.Printf("stream:   %v\n", stream3.Out().Val())
}
