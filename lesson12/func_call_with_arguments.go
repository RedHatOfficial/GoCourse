// Technologie WebAssembly
//
// - rozhraní mezi jazyky Go a JavaScript
// - zobrazení počtu argumentů předaných funkci PrintArguments
// - zobrazení obsahu argumentů předaných funkci PrintArguments

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintArguments(this js.Value, args []js.Value) any {
	fmt.Println("function PrintArguments called")
	fmt.Printf("# of parameters: %d\n", len(args))

	for i, arg := range args {
		fmt.Printf("parameter # %d: %v\n", i, arg)
	}

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce PrintArguments tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printArguments", js.FuncOf(PrintArguments))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
