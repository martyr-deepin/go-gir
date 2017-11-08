package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type ErrorGo struct {
	Domain  Quark
	Code    int
	Message string
}

func (err ErrorGo) Error() string {
	cDomainStr := C.g_quark_to_string(C.GQuark(err.Domain))
	return fmt.Sprintf("GError{Domain: %s(%d), Code: %d, Message: %q}",
		C.GoString((*C.char)(cDomainStr)), err.Domain, err.Code, err.Message)
}

func (v Error) String() string {
	return fmt.Sprintf("%#v", v.native())
}

func (v Error) GoValue() *ErrorGo {
	if v.Ptr == nil {
		return nil
	}
	var err = v.native()
	return &ErrorGo{
		Domain:  Quark(err.domain),
		Code:    int(err.code),
		Message: C.GoString((*C.char)(err.message)),
	}
}

func ErrorNewLiteral(domain Quark, code int, message string) Error {
	cMsg := C.CString(message)
	defer C.free(unsafe.Pointer(cMsg))
	err := C.g_error_new_literal(C.GQuark(domain), C.gint(code), (*C.gchar)(cMsg))
	return wrapError(err)
}

func ErrorNew(domain Quark, code int, format string, args ...interface{}) Error {
	msg := fmt.Sprintf(format, args...)
	return ErrorNewLiteral(domain, code, msg)
}
