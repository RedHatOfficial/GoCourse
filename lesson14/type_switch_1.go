package main

import "fmt"

func test_type(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("Integer value:", v)
	case bool:
		fmt.Println("Boolean value:", v)
	case string:
		fmt.Println("String value:", v)
	default:
		fmt.Println("Unsupported value")
	}
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
