// Technologie WebAssembly
//
// - rozhraní mezi jazyky Go a JavaScript

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintHello(this js.Value, args []js.Value) any {
	// získání objektu typu "window" (z pohledu JavaScriptu)
	window := js.Global()

	// přečtení instance objektu "document"
	document := window.Get("document")

	// získání reference na element s ID="header" umístěného
	// na HTML stránce
	element := document.Call("getElementById", "header")

	// změna atributu elementu
	// (text uvnitř značky)
	element.Set("innerHTML", "Hello from Go")

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce PrintHello tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printHello", js.FuncOf(PrintHello))

	// realizace nekonečného čekání
	<-c

	fmt.Println("finished")
}
