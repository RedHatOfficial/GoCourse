package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Printf("input:    %v\n", values)

	stream := koazee.StreamOf(values)
	fmt.Printf("stream:   %v\n", stream.Out().Val())

	fmt.Println()

	words := []string{"one", "two", "three", "four", "five"}
	fmt.Printf("words:    %v\n", words)

	stream = koazee.StreamOf(words)
	fmt.Printf("stream:   %v\n", stream.Out().Val())

	fmt.Println()

	anything := []interface{}{"one", 1, 1.0, 1 + 1i}
	fmt.Printf("anything: %v\n", anything)

	stream3 := koazee.StreamOf(anything)
	fmt.Printf("stream:   %v\n", stream3.Out().Val())
}
