package main

import (
	"fmt"
	"reflect"
)

func test_get_type(x any) {
	value := reflect.ValueOf(x)
	fmt.Printf("value %v of type %T\n", x, x)
	fmt.Printf("value %v of type %T\n", value, value)
	fmt.Println()
}

type user struct {
	name    string
	surname string
}

func main() {
	var nil1 *int = nil
	test_get_type(nil1)

	var nil2 *bool = nil
	test_get_type(nil2)

	var nil3 *string = nil
	test_get_type(nil3)

	var nil4 *user = nil
	test_get_type(nil4)

	var nil5 interface{} = nil
	test_get_type(nil5)

	var nil6 []int = nil
	test_get_type(nil6)
}
