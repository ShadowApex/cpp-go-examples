// package name: libhelloplugin
package main

import "C"
import "fmt"

//export SayHello
func SayHello(text string) {
	fmt.Println("Hello from Go:", text)
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
