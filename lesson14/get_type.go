package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Println("type is: ", t)
}
