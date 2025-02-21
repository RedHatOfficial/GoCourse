package main

import (
	"fmt"
	"reflect"
)

func test_get_type(x any) {
	value := reflect.ValueOf(x)
	typ := value.Type()
	fmt.Println("type is: ", typ)
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

	var nil5 []int = nil
	test_get_type(nil5)
}
