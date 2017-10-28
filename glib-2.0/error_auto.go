package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
*/
import "C"
import "unsafe"

// Struct Error
type Error struct {
	Ptr unsafe.Pointer
}

func (v Error) native() *C.GError {
	return (*C.GError)(v.Ptr)
}
func wrapError(p *C.GError) Error {
	return Error{unsafe.Pointer(p)}
}
func WrapError(p unsafe.Pointer) Error {
	return Error{p}
}

// Free is a wrapper around g_error_free().
func (error Error) Free() {

	// Var for Go: error
	// Var for C: error0
	// Type for Go: Error
	// Type for C: *C.GError
	C.g_error_free(error.native())
}
