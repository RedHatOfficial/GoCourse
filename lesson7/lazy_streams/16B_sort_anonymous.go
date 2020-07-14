package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	values := []int{100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1}
	fmt.Printf("input: %v\n", values)

	stream1 := koazee.StreamOf(values)
	fmt.Printf("stream1:  %v\n", stream1.Out().Val())

	stream1.Sort(func(x, y int) int { return x - y }).Do()
	fmt.Printf("stream1:  %v\n", stream1.Out().Val())

	stream1.Sort(func(x, y int) int { return y - x }).Do()
	fmt.Printf("stream1:  %v\n", stream1.Out().Val())
}
