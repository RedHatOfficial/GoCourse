package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func printInt(x int) {
	fmt.Printf("%d\n", x)
}

func doubleValue(x int) int {
	return x * 2
}

func negate(x int) int {
	return -x
}

func evenValue(x int) bool {
	return x%2 == 0
}

func divisibleBy3(x int) bool {
	return x%3 == 0
}

func main() {
	values1 := []int{1, 2, 3}
	fmt.Printf("input #1: %v\n", values1)

	stream1 := koazee.StreamOf(values1).
		Filter(evenValue).
		Filter(divisibleBy3).
		Map(negate).
		Map(doubleValue)
	stream1.ForEach(printInt).Do()
}
