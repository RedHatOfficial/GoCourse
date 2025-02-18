package main

import "fmt"

func test_type(value any) {
	int_val, ok := value.(int)
	if ok {
		fmt.Println("Integer value:", int_val)
		return
	}

	bool_val, ok := value.(bool)
	if ok {
		fmt.Println("Boolean value:", bool_val)
		return
	}

	string_val, ok := value.(bool)
	if ok {
		fmt.Println("String value:", string_val)
		return
	}

	fmt.Println("Unsupported value")
}

func main() {
	x := 42
	test_type(x)

	y := true
	test_type(y)

	z := "foobar"
	test_type(z)

	w := 1 + 2i
	test_type(w)
}
