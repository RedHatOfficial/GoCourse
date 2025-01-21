// Technologie WebAssembly a GopherJS
//
// - rozhraní mezi jazyky Go a JavaScript
// - kontrola počtu argumentů předaných funkci PrintArrayLength
// - kontrola typu argumentů předaných funkci PrintArrayLength
// - výpočet délky pole

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintArrayLength(this js.Value, args []js.Value) any {
	// kontrola počtu předaných argumentů
	if len(args) != 1 {
		fmt.Printf("incorrect number of arguments %d, but exactly one is accepted\n", len(args))
		return nil
	}

	// kontrola typu předaného argumentu
	typ := args[0].Type()
	if typ != js.TypeObject {
		fmt.Printf("Argument #0 has incorrect type %s\n", typ.String())
		return nil
	}

	// získat atribut s délkou pole
	array := args[0]
	length := array.Length()

	// zobrazit zprávu
	fmt.Printf("Array length = %d\n", length)

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce printArrayLength tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printArrayLength", js.FuncOf(PrintArrayLength))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
