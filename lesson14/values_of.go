package main

import (
	"fmt"
	"reflect"
)

func test_get_type(x any) {
	value := reflect.ValueOf(x)
	fmt.Println(x)
	fmt.Println(value)
	fmt.Println()
}

func main() {
	x := 42
	test_get_type(x)

	y := true
	test_get_type(y)

	z := "foobar"
	test_get_type(z)

	w := 1 + 2i
	test_get_type(w)
}
