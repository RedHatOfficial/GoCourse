package main

// #include "cgo-independent.h"
// #include <stdlib.h>
import "C"
import "fmt"
import "unsafe"

func main() {
	fmt.Println("Hello from Go.")
	hello := C.CString("Hello from C.\n")
	err := C.hello_from_c(hello)
	if err == C.eof {
		fmt.Println("Printing in C failed with EOF.")
	}
	C.free(unsafe.Pointer(hello))
}
