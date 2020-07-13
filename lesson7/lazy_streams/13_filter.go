package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func evenValue(x int) bool {
	return x%2 == 0
}

func divisibleByThree(x int) bool {
	return x%3 == 0
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("input: %v\n", values)

	stream1 := koazee.StreamOf(values)
	stream2 := stream1.Filter(evenValue).Do()
	stream3 := stream1.Filter(divisibleByThree).Do()

	fmt.Printf("stream1:  %v\n", stream1.Out().Val())
	fmt.Printf("stream2:  %v\n", stream2.Out().Val())
	fmt.Printf("stream3:  %v\n", stream3.Out().Val())
}
