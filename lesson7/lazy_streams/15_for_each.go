package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func printInt(x int) {
	fmt.Printf("%d\t", x)
}

func main() {
	values1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("input #1: %v\n", values1)

	stream1 := koazee.StreamOf(values1)
	stream1.ForEach(printInt)

	fmt.Println()

	values2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("input #2: %v\n", values2)

	stream2 := koazee.StreamOf(values2)
	stream2.ForEach(printInt).Do()
}
