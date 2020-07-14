package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func max(x, y int) int {
	fmt.Printf("%d %d\n", x, y)
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	values1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("input #1: %v\n", values1)

	maxVal := koazee.StreamOf(values1).Reduce(max)
	fmt.Printf("max value: %d\n", maxVal.Int())
}
