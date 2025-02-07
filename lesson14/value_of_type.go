package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	v := reflect.ValueOf(x)
	fmt.Printf("value %v of type %T\n", x, x)
	fmt.Printf("value %v of type %T\n", v, v)
}
