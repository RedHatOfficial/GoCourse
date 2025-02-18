package main

import "fmt"

type Interface1 interface {
	foo()
}

type Interface2 interface {
	bar()
}

type Type1 struct {
	name string
}

func (t Type1) foo() {
}

type Type2 struct {
	name string
}

func (t Type2) bar() {
}

type Type3 struct {
	name string
}

func (t Type3) foo() {
}

func (t Type3) bar() {
}

func test_type(value any) {
	interface1_val, ok := value.(Interface1)
	if ok {
		fmt.Println("Interface1 value:", interface1_val)
	}

	interface2_val, ok := value.(Interface2)
	if ok {
		fmt.Println("Interface2 value:", interface2_val)
	}

	type1_val, ok := value.(Type1)
	if ok {
		fmt.Println("Type1 value:", type1_val)
	}

	type2_val, ok := value.(Type2)
	if ok {
		fmt.Println("Type2 value:", type2_val)
	}

	type3_val, ok := value.(Type3)
	if ok {
		fmt.Println("Type3 value:", type3_val)
	}
}

func main() {
	x := Type1{"x"}
	test_type(x)

	y := Type2{"y"}
	test_type(y)

	z := Type3{"z"}
	test_type(z)
}
