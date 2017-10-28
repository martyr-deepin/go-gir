package glib

/*
#cgo pkg-config: glib-2.0
#include <stdlib.h>
#include <glib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"

// Struct Source
type Source struct {
	Ptr unsafe.Pointer
}

func (v Source) native() *C.GSource {
	return (*C.GSource)(v.Ptr)
}
func wrapSource(p *C.GSource) Source {
	return Source{unsafe.Pointer(p)}
}
func WrapSource(p unsafe.Pointer) Source {
	return Source{p}
}

// Attach is a wrapper around g_source_attach().
func (source Source) Attach(context MainContext) uint {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource

	// Var for Go: context
	// Var for C: context0
	// Type for Go: MainContext
	// Type for C: *C.GMainContext
	ret0 := C.g_source_attach(source.native(), context.native())
	return uint(ret0)
}

// Destroy is a wrapper around g_source_destroy().
func (source Source) Destroy() {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	C.g_source_destroy(source.native())
}

// GetCanRecurse is a wrapper around g_source_get_can_recurse().
func (source Source) GetCanRecurse() bool {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_get_can_recurse(source.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetContext is a wrapper around g_source_get_context().
func (source Source) GetContext() MainContext {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_get_context(source.native())
	return wrapMainContext(ret0)
}

// GetId is a wrapper around g_source_get_id().
func (source Source) GetId() uint {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_get_id(source.native())
	return uint(ret0)
}

// GetName is a wrapper around g_source_get_name().
func (source Source) GetName() string {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_get_name(source.native())
	return C.GoString(ret0)
}

// GetPriority is a wrapper around g_source_get_priority().
func (source Source) GetPriority() int {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_get_priority(source.native())
	return int(ret0)
}

// GetTime is a wrapper around g_source_get_time().
func (source Source) GetTime() int64 {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_get_time(source.native())
	return int64(ret0)
}

// IsDestroyed is a wrapper around g_source_is_destroyed().
func (source Source) IsDestroyed() bool {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_is_destroyed(source.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_source_ref().
func (source Source) Ref() Source {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	ret0 := C.g_source_ref(source.native())
	return wrapSource(ret0)
}

// SetCanRecurse is a wrapper around g_source_set_can_recurse().
func (source Source) SetCanRecurse(can_recurse bool) {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource

	// Var for Go: can_recurse
	// Var for C: can_recurse0
	// Type for Go: bool
	// Type for C: C.gboolean
	C.g_source_set_can_recurse(source.native(), C.gboolean(util.Bool2Int(can_recurse)) /*go:.util*/)
}

// SetName is a wrapper around g_source_set_name().
func (source Source) SetName(name string) {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource

	// Var for Go: name
	// Var for C: name0
	// Type for Go: string
	// Type for C: *C.char
	name0 := C.CString(name)
	defer C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	C.g_source_set_name(source.native(), name0)
}

// SetPriority is a wrapper around g_source_set_priority().
func (source Source) SetPriority(priority int) {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource

	// Var for Go: priority
	// Var for C: priority0
	// Type for Go: int
	// Type for C: C.gint
	C.g_source_set_priority(source.native(), C.gint(priority))
}

// Unref is a wrapper around g_source_unref().
func (source Source) Unref() {

	// Var for Go: source
	// Var for C: source0
	// Type for Go: Source
	// Type for C: *C.GSource
	C.g_source_unref(source.native())
}

// SourceRemove is a wrapper around g_source_remove().
func SourceRemove(tag uint) bool {

	// Var for Go: tag
	// Var for C: tag0
	// Type for Go: uint
	// Type for C: C.guint
	ret0 := C.g_source_remove(C.guint(tag))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SourceSetNameById is a wrapper around g_source_set_name_by_id().
func SourceSetNameById(tag uint, name string) {

	// Var for Go: tag
	// Var for C: tag0
	// Type for Go: uint
	// Type for C: C.guint

	// Var for Go: name
	// Var for C: name0
	// Type for Go: string
	// Type for C: *C.char
	name0 := C.CString(name)
	defer C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	C.g_source_set_name_by_id(C.guint(tag), name0)
}
