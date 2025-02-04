// Technologie GopherJS
//
// Program typu "Hello, world!"
// Použití funkce println

package main

import (
	"syscall/js"
)

func main() {
	println("Hello, world!")
	js.Global().Call("alert", "Hello, world!")
}
