// Technologie GopherJS
//
// - manipulace s DOMem přímo z jazyka Go
// - změna atributu vybraného elementu

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("started")

	// získání objektu typu "window" (z pohledu JavaScriptu)
	window := js.Global()

	// přečtení instance objektu "document"
	document := window.Get("document")

	// získání reference na element s ID="header" umístěného
	// na HTML stránce
	element := document.Call("getElementById", "header")

	// změna atributu elementu
	// (text uvnitř značky)
	element.Set("innerHTML", "foobar")

	fmt.Println("finished")
}
