package main
/*
#cgo pkg-config: slurm
//#c__go CFLAGS: -pthread -fPIC -I /usr/include
//#c__go LDFLAGS: -L /usr/lib64 -lslurm 

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

// This doesn't work, and I don't know why. Apparently I can't use any function in spank.h 
//  that are not defined with 'extern' in them. I'm guessing it has to do with some LDFLAGS
//  or some cgo directives... but well, right now it just doesn't wanna work
// TODO Add the ability to call spank functions, without it, using Go in a spank plugin is 
//  very limited!
//int my_spank_remote(spank_t spank)
//{
//    return spank_remote(spank);
//}

*/
import "C"

// TODO: Test simplified makefile as per 
// https://golang.org/cmd/cgo/
// When the Go tool sees that one or more Go files use the special import "C", it will look for other non-Go files in the directory and compile them as part of the Go package. Any .c, .s, .S or .sx files will be compiled with the C compiler. Any .cc, .cpp, or .cxx files will be compiled with the C++ compiler. Any .f, .F, .for or .f90 files will be compiled with the fortran compiler. Any .h, .hh, .hpp, or .hxx files will not be compiled separately, but, if these header files are changed, the package (including its non-Go source files) will be recompiled. Note that changes to files in other directories do not cause the package to be recompiled, so all non-Go source code for the package should be stored in the package directory, not in subdirectories. The default C and C++ compilers may be changed by the CC and CXX environment variables, respectively; those environment variables may include command line options.

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
func Spank_init(sp C.spank_t) (C.int) {
    var v int = 9
    verbose("Starting ["+trace()+"]")    
    // This won't work, look above
    //var a = int(C.my_spank_remote(sp))
    //fmt.printf("a=%d", a)
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
func Spank_job_prolog(spank C.spank_t) (C.int) {
    verbose("Starting ["+trace()+"]")

    verbose("Finishing ["+trace()+"]")    
    return C.int(C.ESPANK_SUCCESS)
}

// ---------------------------------------------------------------------------------------------------------

// Main is needed to be able to compile this with Go
func main() {}
