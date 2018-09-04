package main

/*
#include <stdio.h>

typedef int (*callback)(char *in_arg);

int
callFunc(void *in_cb, char *in_arg) {
	printf("callFunc in C!\n");
	callback cb = (callback)in_cb;
	return cb(in_arg);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func callCFunc(cb unsafe.Pointer, out *C.char) error {
	fmt.Println("Reached to callCFunc in Go!")
	ret := C.callFunc(cb, out)
	if ret != 0 {
		// TODO: return error
	}
	return nil
}
