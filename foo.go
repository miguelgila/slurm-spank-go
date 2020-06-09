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
    "runtime"
    "regexp"
    "strings"
)

//export PrintInt
func PrintInt(x int) {
    fmt.Println(x)
}

// https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
// This is ugly... but well, for demonstration purposes, can work I guess :grimace: 
func trace() (string){
    pc := make([]uintptr, 10)  // at least 1 entry needed
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    s := fmt.Sprint("%s", f.Name())
    re, err := regexp.Compile(`[^\w]`)
	if err != nil {
        return ""
	}
    str1 := re.ReplaceAllString(s, ".")
    str2 := strings.Replace(str1, ".smain.", "", -1)
	return str2
}

func verbose(s string) {
    cs := C.CString(" go: " + s)
    C.my_slurm_verbose(cs)
    C.free(unsafe.Pointer(cs))
}

//export Spank_init
func Spank_init() (C.int) {
    var v int = 9
    verbose("Starting ["+trace()+"]")    
    s := fmt.Sprintf("%s%d", "Will return to C the value v=", v)
    verbose(s)
    verbose("Finishing ["+trace()+"]")    
    return C.int(v)
}

//export Spank_slurmd_exit
func Spank_slurmd_exit() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_job_epilog
func Spank_job_epilog() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_exit
func Spank_exit() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_task_exit
func Spank_task_exit() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_task_post_fork
// func Spank_task_post_fork() (C.int) {
//     verbose("Starting ["+trace()+"]")
//     verbose("Finishing ["+trace()+"]")    
//     return C.int(0)
// }

//export Spank_task_init
func Spank_task_init() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_task_init_privileged
func Spank_task_init_privileged() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_user_init
func Spank_user_init() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_local_user_init
func Spank_local_user_init() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_init_post_opt
func Spank_init_post_opt() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

//export Spank_job_prolog
func Spank_job_prolog() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(0)
}

func main() {}
