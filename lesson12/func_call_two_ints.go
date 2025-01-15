// Technologie WebAssembly
//
// - rozhraní mezi jazyky Go a JavaScript
// - kontrola počtu argumentů předaných funkci PrintSum
// - kontrola typu argumentů předaných funkci PrintSum
// - provedení konverze na nativní typy jazyka Go

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintSum(this js.Value, args []js.Value) any {
	// kontrola počtu předaných argumentů
	if len(args) != 2 {
		fmt.Printf("incorrect number of arguments %d, but just two are accepted\n", len(args))
		return nil
	}

	// kontrola typu předaných argumentů
	// (pozor - původní syntaxe zápisu smyčky)
	for index := 0; index < 2; index++ {
		typ := args[index].Type()
		if typ != js.TypeNumber {
			fmt.Printf("Argument #%d has incorrect type %s\n", index, typ.String())
			return nil
		}
	}

	// počet i typ argumentů je korektní
	// lze tedy provést jejich konverzi
	x := args[0].Int()
	y := args[1].Int()

	// vypočítat výsledek
	z := x + y

	// zobrazit výsledek
	fmt.Printf("%d + %d = %d\n", x, y, z)

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce PrintSum tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printSum", js.FuncOf(PrintSum))

	// realizace nekonečného čekání
	<-c

	fmt.Println("finished")
}
