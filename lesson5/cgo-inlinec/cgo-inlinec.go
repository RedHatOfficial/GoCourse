package main

// #include <stdio.h>
//
//const int eof = EOF;
//int hello_from_c(_GoString_ stringGo){
//  int err;
//  for(int i=0;i<_GoStringLen(stringGo);i++){
//    err = putchar(*(_GoStringPtr(stringGo)+i));
//    if(err==EOF) break;
//  }
//  return err;
//}
import "C"
import "fmt"

func main() {
	fmt.Println("Hello from Go.")
	err := C.hello_from_c("Hello from C.\n")
	if err == C.eof {
		fmt.Println("Printing in C failed with EOF.")
	}
}
