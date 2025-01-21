// Technologie WebAssembly a GopherJS
//
// - rozhraní mezi jazyky Go a JavaScript
// - kontrola počtu argumentů předaných funkci arraySum
// - kontrola typu argumentů předaných funkci arraySum
// - výpočet délky pole
// - výpočet součtu prvků pole

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func ArraySum(this js.Value, args []js.Value) any {
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

	// projít prvky pole a vypočítat sumu
	sum := 0

	for i := 0; i < length; i++ {
		item := array.Index(i)
		typ := item.Type()
		if typ != js.TypeNumber {
			fmt.Printf("Item #%d has incorrect type %s\n", i, typ.String())
			return nil
		}
		value := item.Int()
		sum += value

	}
	fmt.Printf("Sum = %d\n", sum)
	fmt.Println()

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce arraySum tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("arraySum", js.FuncOf(ArraySum))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
