// Technologie GopherJS
//
// - manipulace s DOMem přímo z jazyka Go
// - přidání nových elementů do HTML stránky

package main

import (
	"fmt"
	"syscall/js"
)

// rekurzivní výpočet faktoriálu
func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}

func main() {
	fmt.Println("started")

	// získání objektu typu "window" (z pohledu JavaScriptu)
	window := js.Global()

	// přečtení instance objektu "document"
	document := window.Get("document")

	// získání reference na element s ID="header" umístěného
	// na HTML stránce
	element := document.Call("getElementById", "header")

	// změna atributu elementu
	// (text uvnitř značky)
	element.Set("innerHTML", "Factorial table")

	// konstrukce tabulky faktoriálů
	for n := int64(0); n <= 10; n++ {
		// výpočet faktoriálu
		f := Factorial(n)
		message := fmt.Sprintf("%2d! = %d", n, f)

		// vytvoření nového elementu
		pre := document.Call("createElement", "pre")
		pre.Set("innerHTML", message)

		// přidání elementu do HTML stránky
		document.Get("body").Call("appendChild", pre)
	}

	fmt.Println("finished")
}
