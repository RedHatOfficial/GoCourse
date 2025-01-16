// Technologie WebAssembly
//
// - rozhraní mezi jazyky Go a JavaScript
// - funkce PrintHello neočekává resp. nezpracovává žadné argumenty

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintHello(this js.Value, args []js.Value) any {
	fmt.Println("function PrintHello called")

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
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
