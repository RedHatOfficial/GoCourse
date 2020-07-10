package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 1}
	fmt.Printf("input: %v\n", values)

	stream := koazee.StreamOf(values)

	i1, _ := stream.IndexOf(1)
	fmt.Printf("index of 1: %v\n", i1)

	i2, _ := stream.LastIndexOf(1)
	fmt.Printf("last index of 1: %v\n", i2)

	i3, _ := stream.IndexOf(42)
	fmt.Printf("index of 42: %v\n", i3)

	fmt.Println()

	words := []string{"one", "two", "three", "four", "five"}
	fmt.Printf("words: %v\n", words)

	stream = koazee.StreamOf(words)

	i4, _ := stream.IndexOf("one")
	fmt.Printf("index of 'one': %v\n", i4)

	i5, _ := stream.LastIndexOf("two")
	fmt.Printf("last index of 'one': %v\n", i5)

	i6, e6 := stream.IndexOf("foobar")
	fmt.Printf("index of 'foobar': %v\n", i6)
	fmt.Printf("error: %v\n", e6)

	fmt.Println()

	i7, e7 := stream.IndexOf(42)
	fmt.Printf("index of 42: %v\n", i7)
	fmt.Printf("error: %v\n", e7)
}
