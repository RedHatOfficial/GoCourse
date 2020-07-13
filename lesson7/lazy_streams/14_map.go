package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func doubleValue(x int) int {
	return x * 2
}

func negate(x int) int {
	return -x
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Printf("input: %v\n", values)

	stream1 := koazee.StreamOf(values)
	stream2 := stream1.Map(doubleValue).Do()
	stream3 := stream1.Map(doubleValue).Map(negate).Do()

	fmt.Printf("stream1:  %v\n", stream1.Out().Val())
	fmt.Printf("stream2:  %v\n", stream2.Out().Val())
	fmt.Printf("stream3:  %v\n", stream3.Out().Val())
}
