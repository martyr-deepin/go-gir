package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <stdlib.h>
*/
import "C"

// GetString is a wrapper around g_variant_get_string().
func (variant Variant) GetString() string {
	var length0 C.gsize
	ret0 := C.g_variant_get_string(variant.native(), &length0)
	ret := C.GoStringN((*C.char)(ret0), C.int(length0))
	return ret
}
