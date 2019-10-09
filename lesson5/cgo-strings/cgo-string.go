package main

// #include <stdlib.h>
// #include <string.h>
import "C"
import "unsafe"
import "fmt"

func main() {

	gostring1 := "The quick brown fox jumps over the lazy dog"
	gostring2 := "quick"

	//memory management
	cstring1 := C.CString(gostring1)
	defer C.free(unsafe.Pointer(cstring1))
	cstring2 := C.CString(gostring2)
	defer C.free(unsafe.Pointer(cstring2))

	cstring3 := C.strstr(cstring1, cstring2)
	//defer C.free(unsafe.Pointer(cstring3))
	//memory management end

	fmt.Printf("%v\n", gostring1)
	fmt.Printf("%v: %v\n", cstring1, C.GoString(cstring1))
	fmt.Printf("%v\n", gostring2)
	fmt.Printf("%v: %v\n", cstring2, C.GoString(cstring2))
	fmt.Printf("%v: %v\n", cstring3, C.GoString(cstring3))

}
