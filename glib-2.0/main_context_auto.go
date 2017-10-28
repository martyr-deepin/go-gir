package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"

// Struct MainContext
type MainContext struct {
	Ptr unsafe.Pointer
}

func (v MainContext) native() *C.GMainContext {
	return (*C.GMainContext)(v.Ptr)
}
func wrapMainContext(p *C.GMainContext) MainContext {
	return MainContext{unsafe.Pointer(p)}
}
func WrapMainContext(p unsafe.Pointer) MainContext {
	return MainContext{p}
}

// MainContextNew is a wrapper around g_main_context_new().
func MainContextNew() MainContext {
	ret0 := C.g_main_context_new()
	return wrapMainContext(ret0)
}

// FindSourceById is a wrapper around g_main_context_find_source_by_id().
func (context MainContext) FindSourceById(source_id uint) Source {

	// Var for Go: context
	// Var for C: context0
	// Type for Go: MainContext
	// Type for C: *C.GMainContext

	// Var for Go: source_id
	// Var for C: source_id0
	// Type for Go: uint
	// Type for C: C.guint
	ret0 := C.g_main_context_find_source_by_id(context.native(), C.guint(source_id))
	return wrapSource(ret0)
}

// Iteration is a wrapper around g_main_context_iteration().
func (context MainContext) Iteration(may_block bool) bool {

	// Var for Go: context
	// Var for C: context0
	// Type for Go: MainContext
	// Type for C: *C.GMainContext

	// Var for Go: may_block
	// Var for C: may_block0
	// Type for Go: bool
	// Type for C: C.gboolean
	ret0 := C.g_main_context_iteration(context.native(), C.gboolean(util.Bool2Int(may_block)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Pending is a wrapper around g_main_context_pending().
func (context MainContext) Pending() bool {

	// Var for Go: context
	// Var for C: context0
	// Type for Go: MainContext
	// Type for C: *C.GMainContext
	ret0 := C.g_main_context_pending(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_main_context_ref().
func (context MainContext) Ref() MainContext {

	// Var for Go: context
	// Var for C: context0
	// Type for Go: MainContext
	// Type for C: *C.GMainContext
	ret0 := C.g_main_context_ref(context.native())
	return wrapMainContext(ret0)
}

// Unref is a wrapper around g_main_context_unref().
func (context MainContext) Unref() {

	// Var for Go: context
	// Var for C: context0
	// Type for Go: MainContext
	// Type for C: *C.GMainContext
	C.g_main_context_unref(context.native())
}

// MainContextDefault is a wrapper around g_main_context_default().
func MainContextDefault() MainContext {
	ret0 := C.g_main_context_default()
	return wrapMainContext(ret0)
}
