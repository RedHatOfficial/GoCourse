// Technologie GopherJS
//
// - kresleni do canvasu

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// rozměry canvasu
	const CanvasWidth = 256
	const CanvasHeight = 256

	fmt.Println("started")

	// získání objektu typu "window" (z pohledu JavaScriptu)
	window := js.Global()
	fmt.Println("window", window)

	// přečtení instance objektu "document"
	document := window.Get("document")
	fmt.Println("document", document)

	// změna nadpisu
	element := document.Call("getElementById", "header")
	fmt.Println("element", element)
	element.Set("innerHTML", "2D canvas")

	// vytvoření elementu typu "canvas"
	canvas := document.Call("createElement", "canvas")
	canvas.Set("height", CanvasWidth)
	canvas.Set("width", CanvasHeight)

	// vložení canvasu na HTML stránku
	document.Get("body").Call("appendChild", canvas)

	// vykreslení grafických objektů na canvas
	context2d := canvas.Call("getContext", "2d")

	// obdélník
	context2d.Set("fillStyle", "#c0c0c0")
	context2d.Call("fillRect", 0, 0, CanvasWidth, CanvasHeight)

	// obdélník zobrazený uvnitř prvního obdélníku
	context2d.Set("fillStyle", "yellow")
	context2d.Call("fillRect", 10, 10, CanvasWidth-20, CanvasHeight-20)

	fmt.Println("finished")
}
