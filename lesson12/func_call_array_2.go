// Technologie WebAssembly a GopherJS
//
// - rozhraní mezi jazyky Go a JavaScript
// - kontrola počtu argumentů předaných funkci PrintArray
// - kontrola typu argumentů předaných funkci PrintArray
// - výpočet délky pole
// - výpis prvků pole

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintArray(this js.Value, args []js.Value) any {
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

	// projít prvky pole
	for i := 0; i < length; i++ {
		item := array.Index(i)
		fmt.Printf("Item #%d has type = %s and value %v\n",
			i,
			item.Type(),
			item)
	}

	fmt.Println()

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce printArray tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printArray", js.FuncOf(PrintArray))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
