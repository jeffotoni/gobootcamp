package main

// #include <stdio.h>
// #include <stdlib.h>
//
// static void print(char *s) {
// 		printf("%s\n", s);
// }
import "C"
import "unsafe"

func main() {
	cs := C.CString("Golang Devopsfest...")
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}
