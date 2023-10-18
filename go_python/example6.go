package main

import "C"

import "fmt"

//export hello
func hello(name *C.char) int {
	goName := C.GoString(name)
	fmt.Printf("Hello %s\n", goName)
	return len(goName)
}

func main() {}
