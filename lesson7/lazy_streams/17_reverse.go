package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("input: %v\n", values)

	stream1 := koazee.StreamOf(values)
	fmt.Printf("stream1:  %v\n", stream1.Out().Val())

	stream2 := stream1.Reverse().Do()
	fmt.Printf("stream2:  %v\n", stream2.Out().Val())

	stream3 := stream2.Reverse().Do()
	fmt.Printf("stream3:  %v\n", stream3.Out().Val())
}
