package main

/*
#cgo CXXFLAGS: -I.
#cgo LDFLAGS: -lstdc++ -L ./ -lsayhello
#include "sayhello.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("vim-go")
	C.sayhello(C.CString("WHAT"))
}
