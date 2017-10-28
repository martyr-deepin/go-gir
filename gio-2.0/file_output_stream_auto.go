package gio

/*
#cgo pkg-config: gio-2.0 gio-unix-2.0
#include <stdlib.h>
#include <gio/gdesktopappinfo.h>
#include <gio/gfiledescriptorbased.h>
#include <gio/gio.h>
#include <gio/gunixconnection.h>
#include <gio/gunixcredentialsmessage.h>
#include <gio/gunixfdlist.h>
#include <gio/gunixfdmessage.h>
#include <gio/gunixinputstream.h>
#include <gio/gunixmounts.h>
#include <gio/gunixoutputstream.h>
#include <gio/gunixsocketaddress.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/glib-2.0"

// Object FileOutputStream
type FileOutputStream struct {
	OutputStream
}

func (v FileOutputStream) native() *C.GFileOutputStream {
	return (*C.GFileOutputStream)(v.Ptr)
}
func wrapFileOutputStream(p *C.GFileOutputStream) (v FileOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileOutputStream(p unsafe.Pointer) (v FileOutputStream) {
	v.Ptr = p
	return
}

// QueryInfo is a wrapper around g_file_output_stream_query_info().
func (stream FileOutputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {

	// Var for Go: stream
	// Var for C: stream0
	// Type for Go: FileOutputStream
	// Type for C: *C.GFileOutputStream

	// Var for Go: attributes
	// Var for C: attributes0
	// Type for Go: string
	// Type for C: *C.char
	attributes0 := C.CString(attributes)
	defer C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/

	// Var for Go: cancellable
	// Var for C: cancellable0
	// Type for Go: Cancellable
	// Type for C: *C.GCancellable
	var err glib.Error
	ret0 := C.g_file_output_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}
