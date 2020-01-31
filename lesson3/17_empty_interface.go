package main

// I represents a new interface type (in this case empty interface)
type I interface{}

// T represents an user-defined data type
type T struct{}

// M is method defined for a data type T
func (T) M() {
	println("Nothing happened...")
}

func main() {
	var i I = T{}
	println(i)

	var t T = T{}
	t.M()
}
