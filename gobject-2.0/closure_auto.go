package gobject

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
*/
import "C"
import "unsafe"

// Struct Closure
type Closure struct {
	Ptr unsafe.Pointer
}

func (v Closure) native() *C.GClosure {
	return (*C.GClosure)(v.Ptr)
}
func wrapClosure(p *C.GClosure) Closure {
	return Closure{unsafe.Pointer(p)}
}
func WrapClosure(p unsafe.Pointer) Closure {
	return Closure{p}
}

// Invalidate is a wrapper around g_closure_invalidate().
func (closure Closure) Invalidate() {

	// Var for Go: closure
	// Var for C: closure0
	// Type for Go: Closure
	// Type for C: *C.GClosure
	C.g_closure_invalidate(closure.native())
}

// Ref is a wrapper around g_closure_ref().
func (closure Closure) Ref() Closure {

	// Var for Go: closure
	// Var for C: closure0
	// Type for Go: Closure
	// Type for C: *C.GClosure
	ret0 := C.g_closure_ref(closure.native())
	return wrapClosure(ret0)
}

// Sink is a wrapper around g_closure_sink().
func (closure Closure) Sink() {

	// Var for Go: closure
	// Var for C: closure0
	// Type for Go: Closure
	// Type for C: *C.GClosure
	C.g_closure_sink(closure.native())
}

// Unref is a wrapper around g_closure_unref().
func (closure Closure) Unref() {

	// Var for Go: closure
	// Var for C: closure0
	// Type for Go: Closure
	// Type for C: *C.GClosure
	C.g_closure_unref(closure.native())
}
