package main

import (
	"fmt"
	"reflect"
)

func into_type(value reflect.Value, typ reflect.Type) {
	switch typ.Name() {
	case "int":
		fallthrough
	case "int32":
		realValue := value.Int()
		fmt.Println("is an integer with value", realValue)
	case "float64":
		realValue := value.Float()
		fmt.Println("is a float64 with value", realValue)
	case "complex128":
		realValue := value.Complex()
		fmt.Println("is a complex128 with value", realValue)
	case "bool":
		realValue := value.Bool()
		fmt.Println("is a boolean with value", realValue)
	case "string":
		realValue := value.String()
		fmt.Println("is a string with value", realValue)
	case "user":
		realValue := value.String()
		fmt.Println("is an user with value", realValue)
	default:
		fmt.Println("is ***unknown***")
	}
}

func test_check_type(x any) {
	value := reflect.ValueOf(x)
	typ := value.Type()
	fmt.Printf("value %v ", value)
	into_type(value, typ)
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

	//var x13 interface{} = nil
	//test_check_type(x13)
}
