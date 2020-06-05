package main
/*
#cgo CFLAGS: -I /usr/include 
#cgo LDFLAGS: -L /usr/lib64 -lslurm 
#include <slurm/spank.h>
#include <stdlib.h>
#include <stdio.h>
static void my_slurm_verbose(char* s) {
    slurm_verbose(s);
}
*/
import "C"

import (
    "fmt"
    "unsafe"
)

//export PrintInt
func PrintInt(x int) {
    fmt.Println(x)
}

//export Spank_init
func Spank_init() (C.int) {
    var v int = 9
    fmt.Println("bar: Go>Spank_init() starting")
    
    cs := C.CString("bar: Go>Spank_init() starting __CString__ using slurm_verbose")
    C.my_slurm_verbose(cs)
    C.free(unsafe.Pointer(cs))

    fmt.Println("bar: Go>Spank_init() v=%d", v)
    fmt.Println("bar: Go>Spank_init() end")
    return C.int(v)
}

func main() {}
