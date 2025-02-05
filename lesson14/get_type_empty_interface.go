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
	var nil1 interface{} = nil
	test_get_type(nil1)
}
