package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"

// Struct MainLoop
type MainLoop struct {
	Ptr unsafe.Pointer
}

func (v MainLoop) native() *C.GMainLoop {
	return (*C.GMainLoop)(v.Ptr)
}
func wrapMainLoop(p *C.GMainLoop) MainLoop {
	return MainLoop{unsafe.Pointer(p)}
}
func WrapMainLoop(p unsafe.Pointer) MainLoop {
	return MainLoop{p}
}

// MainLoopNew is a wrapper around g_main_loop_new().
func MainLoopNew(context MainContext, is_running bool) MainLoop {

	// Var for Go: context
	// Var for C: context0
	// Type for Go: MainContext
	// Type for C: *C.GMainContext

	// Var for Go: is_running
	// Var for C: is_running0
	// Type for Go: bool
	// Type for C: C.gboolean
	ret0 := C.g_main_loop_new(context.native(), C.gboolean(util.Bool2Int(is_running)) /*go:.util*/)
	return wrapMainLoop(ret0)
}

// GetContext is a wrapper around g_main_loop_get_context().
func (loop MainLoop) GetContext() MainContext {

	// Var for Go: loop
	// Var for C: loop0
	// Type for Go: MainLoop
	// Type for C: *C.GMainLoop
	ret0 := C.g_main_loop_get_context(loop.native())
	return wrapMainContext(ret0)
}

// IsRunning is a wrapper around g_main_loop_is_running().
func (loop MainLoop) IsRunning() bool {

	// Var for Go: loop
	// Var for C: loop0
	// Type for Go: MainLoop
	// Type for C: *C.GMainLoop
	ret0 := C.g_main_loop_is_running(loop.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Quit is a wrapper around g_main_loop_quit().
func (loop MainLoop) Quit() {

	// Var for Go: loop
	// Var for C: loop0
	// Type for Go: MainLoop
	// Type for C: *C.GMainLoop
	C.g_main_loop_quit(loop.native())
}

// Ref is a wrapper around g_main_loop_ref().
func (loop MainLoop) Ref() MainLoop {

	// Var for Go: loop
	// Var for C: loop0
	// Type for Go: MainLoop
	// Type for C: *C.GMainLoop
	ret0 := C.g_main_loop_ref(loop.native())
	return wrapMainLoop(ret0)
}

// Run is a wrapper around g_main_loop_run().
func (loop MainLoop) Run() {

	// Var for Go: loop
	// Var for C: loop0
	// Type for Go: MainLoop
	// Type for C: *C.GMainLoop
	C.g_main_loop_run(loop.native())
}

// Unref is a wrapper around g_main_loop_unref().
func (loop MainLoop) Unref() {

	// Var for Go: loop
	// Var for C: loop0
	// Type for Go: MainLoop
	// Type for C: *C.GMainLoop
	C.g_main_loop_unref(loop.native())
}
