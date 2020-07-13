package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("input: %v\n", values)

	stream := koazee.StreamOf(values)
	fmt.Printf("stream:   %v\n", stream.Out().Val())

	stream2 := stream.Take(3, 6)
	fmt.Printf("stream2:  %v\n", stream2.Out().Val())

	stream3 := stream.Take(3, 4)
	fmt.Printf("stream3:  %v\n", stream3.Out().Val())

	stream4 := stream.Take(3, 3)
	fmt.Printf("stream4:  %v\n", stream4.Out().Val())

	stream5 := stream.Take(4, 3)
	fmt.Printf("stream5:  %v\n", stream5.Out().Val())

	stream6 := stream.Take(4, 100)
	fmt.Printf("stream6:  %v\n", stream6.Out().Val())
}
