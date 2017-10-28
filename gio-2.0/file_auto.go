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
import "github.com/electricface/go-auto-gir/util"

// Interface File
type File struct {
	Ptr unsafe.Pointer
}

func (v File) native() *C.GFile {
	return (*C.GFile)(v.Ptr)
}
func wrapFile(p *C.GFile) File {
	return File{unsafe.Pointer(p)}
}
func WrapFile(p unsafe.Pointer) File {
	return File{p}
}

// GetPath is a wrapper around g_file_get_path().
func (file File) GetPath() string {

	// Var for Go: file
	// Var for C: file0
	// Type for Go: File
	// Type for C: *C.GFile
	ret0 := C.g_file_get_path(file.native())
	defer C.g_free(C.gpointer(ret0))
	return C.GoString(ret0)
}

// Replace is a wrapper around g_file_replace().
func (file File) Replace(etag string, make_backup bool, flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {

	// Var for Go: file
	// Var for C: file0
	// Type for Go: File
	// Type for C: *C.GFile

	// Var for Go: etag
	// Var for C: etag0
	// Type for Go: string
	// Type for C: *C.char
	etag0 := C.CString(etag)
	defer C.free(unsafe.Pointer(etag0)) /*ch:<stdlib.h>*/

	// Var for Go: make_backup
	// Var for C: make_backup0
	// Type for Go: bool
	// Type for C: C.gboolean

	// Var for Go: flags
	// Var for C: flags0
	// Type for Go: FileCreateFlags
	// Type for C: C.GFileCreateFlags

	// Var for Go: cancellable
	// Var for C: cancellable0
	// Type for Go: Cancellable
	// Type for C: *C.GCancellable
	var err glib.Error
	ret0 := C.g_file_replace(file.native(), etag0, C.gboolean(util.Bool2Int(make_backup)) /*go:.util*/, C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}
