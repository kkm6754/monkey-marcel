package main

import (
	"fmt"
	"unsafe"
)

func main() {
	f := 1.0
	i := *(*uint64)(unsafe.Pointer(&f))

	fmt.Println(i)
}