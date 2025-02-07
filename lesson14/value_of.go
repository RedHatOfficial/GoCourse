package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	v := reflect.ValueOf(x)
	fmt.Println(x)
	fmt.Println(v)
}
