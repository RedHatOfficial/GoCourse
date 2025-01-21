// Technologie WebAssembly a GopherJS
//
// - rozhraní mezi jazyky Go a JavaScript
// - zavolání callback funkce

package main

import (
	"fmt"
	"syscall/js"
)

func CallFunction(this js.Value, args []js.Value) any {
	// kontrola počtu předaných argumentů
	if len(args) != 1 {
		fmt.Printf("incorrect number of arguments %d, but exactly one is accepted\n", len(args))
		return nil
	}

	// kontrola typu předaného argumentu
	typ := args[0].Type()
	if typ != js.TypeFunction {
		fmt.Printf("Argument #0 has incorrect type %s\n", typ.String())
		return nil
	}

	function := args[0]

	function.Invoke("Called from Go!")

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	js.Global().Set("callFunction", js.FuncOf(CallFunction))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
