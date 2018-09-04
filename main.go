package main

import "C"
import (
	"fmt"
	"unsafe"

	"github.com/pankona/cgo-server-sample/server"
)

func main() {
}

var cb unsafe.Pointer

//export SetCallback
func SetCallback(funcPointerFromC unsafe.Pointer) {
	cb = funcPointerFromC
	fmt.Printf("SetCallback called. cb = %v\n", cb)
}

//export Run
func Run() error {
	return server.Run(callback)
}

func callback() string {
	fmt.Println("Reached to callback in Go!")
	var out = make([]C.char, 256)
	callCFunc(cb, &out[0])

	s := C.GoString(&out[0])
	fmt.Printf("buffer written in C is: %s\n", s)

	return s
}
