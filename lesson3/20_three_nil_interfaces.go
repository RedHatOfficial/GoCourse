package main

import "fmt"

type T int

func main() {
	var p1 *int = nil
	var p2 *string = nil

	var i1 interface{} = nil
        var i2 interface{} = p1
        var i3 interface{} = p2

	fmt.Printf("%T\t%v\n", i1, i1)
	fmt.Printf("%T\t%v\n", i2, i2)
	fmt.Printf("%T\t%v\n", i3, i3)

	fmt.Println()
	fmt.Printf("%v\n", i1==i2)
	fmt.Printf("%v\n", i1==i3)
	fmt.Printf("%v\n", i2==i3)
}
