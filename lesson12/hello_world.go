// Technologie GopherJS
//
// Program typu "Hello, world!"
// Použití funkce fmt.Println

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, world!")
	js.Global().Call("alert", "Hello, world!")
}
