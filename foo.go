package main
/*
#cgo CFLAGS: -I /usr/include 
#cgo LDFLAGS: -L /usr/lib64 -lslurm 
#include <slurm/spank.h>
#include <stdlib.h>
#include <stdio.h>

static void my_slurm_info(char* s) {
    slurm_info(s);
}
static void my_slurm_error(char* s) {
    slurm_error(s);
}
static void my_slurm_verbose(char* s) {
    slurm_verbose(s);
}
static void my_slurm_debug(char* s) {
    slurm_debug(s);
}
static void my_slurm_debug2(char* s) {
    slurm_debug2(s);
}
static void my_slurm_debug3(char* s) {
    slurm_debug3(s);
}
static void my_slurm_spank_log(char* s) {
    slurm_spank_log(s);
}
// this is to reduce the pesky warning with Wall: function defined not used blablabla
static void dummy() {
    my_slurm_info("");
    my_slurm_error("");
    my_slurm_verbose("");
    my_slurm_debug("");
    my_slurm_debug2("");
    my_slurm_debug3("");
    my_slurm_spank_log("");
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
// Wrappers ------------------------------------------------------------------------------------------------
// These functions get defined only because Go cannot call directly any variadic C function
// From: https://golang.org/cmd/cgo/#hdr-Go_references_to_C
//  Calling variadic C functions is not supported. It is possible to circumvent this by using a 
//   C function wrapper. Other than that, we can call any SPANK piece with C.piece_name like 
//   C.struct_stat. or C.ESPANK_SUCCESS
// ---------------------------------------------------------------------------------------------------------
func info(s string) {
    cs := C.CString(" go: " + s)
    C.my_slurm_info(cs)
    C.free(unsafe.Pointer(cs))
}
func error(s string) {
    cs := C.CString(" go: " + s)
    C.my_slurm_error(cs)
    C.free(unsafe.Pointer(cs))
}
func verbose(s string) {
    cs := C.CString(" go: " + s)
    C.my_slurm_verbose(cs)
    C.free(unsafe.Pointer(cs))
}
func debug(s string) {
    cs := C.CString(" go: " + s)
    C.my_slurm_debug(cs)
    C.free(unsafe.Pointer(cs))
}
func debug2(s string) {
    cs := C.CString(" go: " + s)
    C.my_slurm_debug(cs)
    C.free(unsafe.Pointer(cs))
}
func debug3(s string) {
    cs := C.CString(" go: " + s)
    C.my_slurm_debug(cs)
    C.free(unsafe.Pointer(cs))
}
func spank_log(s string) {
    cs := C.CString(s)
    C.my_slurm_spank_log(cs)
    C.free(unsafe.Pointer(cs))
}

// https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
// This is ugly... but well, for demonstration purposes, it can work I guess :grimace: 
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

// ---------------------------------------------------------------------------------------------------------


// SPANK functions -----------------------------------------------------------------------------------------

//export Spank_init
func Spank_init() (C.int) {
    var v int = 9
    verbose("Starting ["+trace()+"]")    
    s := fmt.Sprintf("%s%d", "Will return to C the value v=", v)
    spank_log(s)
    verbose("Finishing ["+trace()+"]")    
    return C.int(v)
}

//export Spank_slurmd_exit
func Spank_slurmd_exit() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_job_epilog
func Spank_job_epilog() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_exit
func Spank_exit() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_task_exit
func Spank_task_exit() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_task_post_fork
// func Spank_task_post_fork() (C.int) {
//     verbose("Starting ["+trace()+"]")
//     verbose("Finishing ["+trace()+"]")    
//     return C.int(C.ESPANK_SUCCESS)
// }

//export Spank_task_init
func Spank_task_init() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_task_init_privileged
func Spank_task_init_privileged() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_user_init
func Spank_user_init() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_local_user_init
func Spank_local_user_init() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_init_post_opt
func Spank_init_post_opt() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

//export Spank_job_prolog
func Spank_job_prolog() (C.int) {
    verbose("Starting ["+trace()+"]")
    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

// ---------------------------------------------------------------------------------------------------------

// Main is needed to be able to compile this with Go
func main() {}
