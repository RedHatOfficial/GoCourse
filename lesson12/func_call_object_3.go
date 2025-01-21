// Technologie WebAssembly a GopherJS
//
// - rozhraní mezi jazyky Go a JavaScript
// - kontrola počtu argumentů předaných funkci PrintObject
// - kontrola typu argumentů předaných funkci PrintObject

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintObject(this js.Value, args []js.Value) any {
	// kontrola počtu předaných argumentů
	if len(args) != 1 {
		fmt.Printf("incorrect number of arguments %d, but exactly one is accepted\n", len(args))
		return nil
	}

	object := args[0]

	// kontrola typu předaného argumentu
	typ := object.Type()
	if typ != js.TypeObject {
		fmt.Printf("Argument #0 has incorrect type %s\n", typ.String())
		return nil
	}

	firstName := object.Get("firstName").String()
	lastName := object.Get("lastName").String()
	age := object.Get("age").Int()
	eyeColor := object.Get("eyeColor").String()

	// zobrazit atributy objektu
	fmt.Printf("First name: %s (%T)\n", firstName, firstName)
	fmt.Printf("Last name:  %s (%T)\n", lastName, lastName)
	fmt.Printf("Age:        %d (%T)\n", age, age)
	fmt.Printf("Eye color:  %s (%T)\n", eyeColor, eyeColor)

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce PrintObject tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printObject", js.FuncOf(PrintObject))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
