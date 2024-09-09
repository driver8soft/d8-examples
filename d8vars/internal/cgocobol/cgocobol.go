package d8cgo

/*
#cgo CFLAGS: -I/opt/homebrew/Cellar/gnucobol/3.2/include
#cgo LDFLAGS: -L/opt/homebrew/Cellar/gnucobol/3.2/lib -lcob
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <libcob.h>
static void* allocArgv(int argc) {
    return malloc(sizeof(char *) * argc);
}
*/
import "C"
import (
	"errors"
	"log"
	"unsafe"
)

func CobolCall(pgm string, commarea string) error {

	// Initilize gnucobol
	C.cob_init(C.int(0), nil)
	log.Println("gnucobol initialized")

	// Get program name
	c_prog := C.CString(pgm)
	defer C.free(unsafe.Pointer(c_prog))

	// Prepare args (argc, argv)
	c_argc := C.int(1)
	c_argv := (*[0xfff]*C.char)(C.allocArgv(c_argc))
	defer C.free(unsafe.Pointer(c_argv))
	c_argv[0] = C.CString(commarea)

	// Check COBOL program
	n := C.cob_resolve(c_prog)
	if n == nil {
		err := errors.New("program not found =" + pgm)
		return err
	}

	// Call COBOL program
	log.Println("program started")
	ret := C.cob_call(c_prog, c_argc, (*unsafe.Pointer)(unsafe.Pointer(c_argv)))
	log.Printf("program return-code %v", ret)

	// D8 stop gnucobol
	C.cob_tidy()
	log.Println("gnucobol end execution")

	return nil

}
