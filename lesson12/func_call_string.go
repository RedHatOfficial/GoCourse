// Technologie WebAssembly a GopherJS
//
// - rozhraní mezi jazyky Go a JavaScript
// - kontrola počtu argumentů předaných funkci PrintMessage
// - kontrola typu argumentů předaných funkci PrintMessage
// - provedení konverze na nativní typy jazyka Go

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintMessage(this js.Value, args []js.Value) any {
	// kontrola počtu předaných argumentů
	if len(args) != 1 {
		fmt.Printf("incorrect number of arguments %d, but exactly one is accepted\n", len(args))
		return nil
	}

	// kontrola typu předaného argumentu
	typ := args[0].Type()
	if typ != js.TypeString {
		fmt.Printf("Argument #0 has incorrect type %s\n", typ.String())
		return nil
	}

	// provést konverzi
	message := args[0].String()

	// zobrazit zprávu
	fmt.Println(message)

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce PrintMessage tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printMessage", js.FuncOf(PrintMessage))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
