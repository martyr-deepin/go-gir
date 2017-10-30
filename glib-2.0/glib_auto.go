package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"

// Struct Variant
type Variant struct {
	Ptr unsafe.Pointer
}

func (v Variant) native() *C.GVariant {
	return (*C.GVariant)(v.Ptr)
}
func wrapVariant(p *C.GVariant) Variant {
	return Variant{unsafe.Pointer(p)}
}
func WrapVariant(p unsafe.Pointer) Variant {
	return Variant{p}
}

// VariantNewBoolean is a wrapper around g_variant_new_boolean().
func VariantNewBoolean(value bool) Variant {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: bool
	// Type for C: C.gboolean
	ret0 := C.g_variant_new_boolean(C.gboolean(util.Bool2Int(value)) /*go:.util*/)
	return wrapVariant(ret0)
}

// GetType is a wrapper around g_variant_get_type().
func (value Variant) GetType() VariantType {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Variant
	// Type for C: *C.GVariant
	ret0 := C.g_variant_get_type(value.native())
	return wrapVariantType(ret0)
}

// IsFloating is a wrapper around g_variant_is_floating().
func (value Variant) IsFloating() bool {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Variant
	// Type for C: *C.GVariant
	ret0 := C.g_variant_is_floating(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_variant_ref().
func (value Variant) Ref() Variant {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Variant
	// Type for C: *C.GVariant
	ret0 := C.g_variant_ref(value.native())
	return wrapVariant(ret0)
}

// RefSink is a wrapper around g_variant_ref_sink().
func (value Variant) RefSink() Variant {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Variant
	// Type for C: *C.GVariant
	ret0 := C.g_variant_ref_sink(value.native())
	return wrapVariant(ret0)
}

// Unref is a wrapper around g_variant_unref().
func (value Variant) Unref() {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Variant
	// Type for C: *C.GVariant
	C.g_variant_unref(value.native())
}

// Struct VariantType
type VariantType struct {
	Ptr unsafe.Pointer
}

func (v VariantType) native() *C.GVariantType {
	return (*C.GVariantType)(v.Ptr)
}
func wrapVariantType(p *C.GVariantType) VariantType {
	return VariantType{unsafe.Pointer(p)}
}
func WrapVariantType(p unsafe.Pointer) VariantType {
	return VariantType{p}
}

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
