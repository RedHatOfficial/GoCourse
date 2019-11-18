package main

type I interface{}

type T struct{}

func (T) M() {
	println("Nothing happened...")
}

func main() {
	var i I = T{}
	println(i)

	var t T = T{}
	t.M()
}
