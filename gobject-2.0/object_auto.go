package gobject

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"

// Object Object
type Object struct {
	Ptr unsafe.Pointer
}

func (v Object) native() *C.GObject {
	return (*C.GObject)(v.Ptr)
}
func wrapObject(p *C.GObject) Object {
	return Object{unsafe.Pointer(p)}
}
func WrapObject(p unsafe.Pointer) Object {
	return Object{p}
}

// ForceFloating is a wrapper around g_object_force_floating().
func (object Object) ForceFloating() {

	// Var for Go: object
	// Var for C: object0
	// Type for Go: Object
	// Type for C: *C.GObject
	C.g_object_force_floating(object.native())
}

// IsFloating is a wrapper around g_object_is_floating().
func (object Object) IsFloating() bool {

	// Var for Go: object
	// Var for C: object0
	// Type for Go: Object
	// Type for C: C.gpointer
	ret0 := C.g_object_is_floating(C.gpointer(object.native()))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_object_ref().
func (object Object) Ref() Object {

	// Var for Go: object
	// Var for C: object0
	// Type for Go: Object
	// Type for C: C.gpointer
	ret0 := C.g_object_ref(C.gpointer(object.native()))
	return WrapObject(unsafe.Pointer(ret0))
}

// RefSink is a wrapper around g_object_ref_sink().
func (object Object) RefSink() Object {

	// Var for Go: object
	// Var for C: object0
	// Type for Go: Object
	// Type for C: C.gpointer
	ret0 := C.g_object_ref_sink(C.gpointer(object.native()))
	return WrapObject(unsafe.Pointer(ret0))
}

// Unref is a wrapper around g_object_unref().
func (object Object) Unref() {

	// Var for Go: object
	// Var for C: object0
	// Type for Go: Object
	// Type for C: C.gpointer
	C.g_object_unref(C.gpointer(object.native()))
}
