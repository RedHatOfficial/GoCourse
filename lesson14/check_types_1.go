package main

import (
	"fmt"
	"reflect"
)

func test_check_type(x any) {
	value := reflect.ValueOf(x)
	typ := value.Type()
	fmt.Printf("value %v of type %s\n", value, typ)
}

type user struct {
	id   int
	name string
}

func main() {
	x1 := 42
	test_check_type(x1)

	x2 := 3.14
	test_check_type(x2)

	x4 := 1 + 2i
	test_check_type(x4)

	x5 := true
	test_check_type(x5)

	x6 := '?'
	test_check_type(x6)

	x7 := "foobar"
	test_check_type(x7)

	x8 := [...]int{1, 2, 3}
	test_check_type(x8)

	x9 := []int{1, 2, 3}
	test_check_type(x9)

	x10 := make(map[int]string)
	test_check_type(x10)

	x11 := user{}
	test_check_type(x11)

	x12 := func(x int) int { return x + 1 }
	test_check_type(x12)
}
