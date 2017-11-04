package gobject

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"

// Struct Value
type Value struct {
	Ptr unsafe.Pointer
}

func (v Value) native() *C.GValue {
	return (*C.GValue)(v.Ptr)
}
func wrapValue(p *C.GValue) Value {
	return Value{unsafe.Pointer(p)}
}
func WrapValue(p unsafe.Pointer) Value {
	return Value{p}
}

// GetPointer is a wrapper around g_value_get_pointer().
func (value Value) GetPointer() unsafe.Pointer {
	ret0 := C.g_value_get_pointer(value.native())
	return unsafe.Pointer(ret0)
}

// GetString is a wrapper around g_value_get_string().
func (value Value) GetString() string {
	ret0 := C.g_value_get_string(value.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetBoolean is a wrapper around g_value_set_boolean().
func (value Value) SetBoolean(v_boolean bool) {
	C.g_value_set_boolean(value.native(), C.gboolean(util.Bool2Int(v_boolean)) /*go:.util*/)
}

// SetDouble is a wrapper around g_value_set_double().
func (value Value) SetDouble(v_double float64) {
	C.g_value_set_double(value.native(), C.gdouble(v_double))
}

// SetFloat is a wrapper around g_value_set_float().
func (value Value) SetFloat(v_float float32) {
	C.g_value_set_float(value.native(), C.gfloat(v_float))
}

// SetInstance is a wrapper around g_value_set_instance().
func (value Value) SetInstance(instance unsafe.Pointer) {
	C.g_value_set_instance(value.native(), C.gpointer(instance))
}

// SetInt is a wrapper around g_value_set_int().
func (value Value) SetInt(v_int int) {
	C.g_value_set_int(value.native(), C.gint(v_int))
}

// SetInt64 is a wrapper around g_value_set_int64().
func (value Value) SetInt64(v_int64 int64) {
	C.g_value_set_int64(value.native(), C.gint64(v_int64))
}

// SetPointer is a wrapper around g_value_set_pointer().
func (value Value) SetPointer(v_pointer unsafe.Pointer) {
	C.g_value_set_pointer(value.native(), C.gpointer(v_pointer))
}

// SetSchar is a wrapper around g_value_set_schar().
func (value Value) SetSchar(v_char int8) {
	C.g_value_set_schar(value.native(), C.gint8(v_char))
}

// SetString is a wrapper around g_value_set_string().
func (value Value) SetString(v_string string) {
	v_string0 := (*C.gchar)(C.CString(v_string))
	C.g_value_set_string(value.native(), v_string0)
	C.free(unsafe.Pointer(v_string0)) /*ch:<stdlib.h>*/
}

// SetUchar is a wrapper around g_value_set_uchar().
func (value Value) SetUchar(v_uchar byte) {
	C.g_value_set_uchar(value.native(), C.guchar(v_uchar))
}

// SetUint is a wrapper around g_value_set_uint().
func (value Value) SetUint(v_uint uint) {
	C.g_value_set_uint(value.native(), C.guint(v_uint))
}

// SetUint64 is a wrapper around g_value_set_uint64().
func (value Value) SetUint64(v_uint64 uint64) {
	C.g_value_set_uint64(value.native(), C.guint64(v_uint64))
}

// Unset is a wrapper around g_value_unset().
func (value Value) Unset() {
	C.g_value_unset(value.native())
}

// Object Object
type Object struct {
	Ptr unsafe.Pointer
}

func (v Object) native() *C.GObject {
	return (*C.GObject)(v.Ptr)
}
func wrapObject(p *C.GObject) (v Object) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapObject(p unsafe.Pointer) (v Object) {
	v.Ptr = p
	return
}

// ForceFloating is a wrapper around g_object_force_floating().
func (object Object) ForceFloating() {
	C.g_object_force_floating(object.native())
}

// IsFloating is a wrapper around g_object_is_floating().
func (object Object) IsFloating() bool {
	ret0 := C.g_object_is_floating(C.gpointer(object.native()))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_object_ref().
func (object Object) Ref() Object {
	ret0 := C.g_object_ref(C.gpointer(object.native()))
	return WrapObject(unsafe.Pointer(ret0))
}

// RefSink is a wrapper around g_object_ref_sink().
func (object Object) RefSink() Object {
	ret0 := C.g_object_ref_sink(C.gpointer(object.native()))
	return WrapObject(unsafe.Pointer(ret0))
}

// Unref is a wrapper around g_object_unref().
func (object Object) Unref() {
	C.g_object_unref(C.gpointer(object.native()))
}

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
	C.g_closure_invalidate(closure.native())
}

// Ref is a wrapper around g_closure_ref().
func (closure Closure) Ref() Closure {
	ret0 := C.g_closure_ref(closure.native())
	return wrapClosure(ret0)
}

// Sink is a wrapper around g_closure_sink().
func (closure Closure) Sink() {
	C.g_closure_sink(closure.native())
}

// Unref is a wrapper around g_closure_unref().
func (closure Closure) Unref() {
	C.g_closure_unref(closure.native())
}

type BindingFlags int

const (
	BindingFlagsDefault       BindingFlags = 0
	BindingFlagsBidirectional              = 1
	BindingFlagsSyncCreate                 = 2
	BindingFlagsInvertBoolean              = 4
)

type ConnectFlags int

const (
	ConnectFlagsAfter   ConnectFlags = 1
	ConnectFlagsSwapped              = 2
)

type ParamFlags int

const (
	ParamFlagsReadable       ParamFlags = 1
	ParamFlagsWritable                  = 2
	ParamFlagsReadwrite                 = 3
	ParamFlagsConstruct                 = 4
	ParamFlagsConstructOnly             = 8
	ParamFlagsLaxValidation             = 16
	ParamFlagsStaticName                = 32
	ParamFlagsPrivate                   = 32
	ParamFlagsStaticNick                = 64
	ParamFlagsStaticBlurb               = 128
	ParamFlagsExplicitNotify            = 1073741824
	ParamFlagsDeprecated                = 2147483648
)

type SignalFlags int

const (
	SignalFlagsRunFirst    SignalFlags = 1
	SignalFlagsRunLast                 = 2
	SignalFlagsRunCleanup              = 4
	SignalFlagsNoRecurse               = 8
	SignalFlagsDetailed                = 16
	SignalFlagsAction                  = 32
	SignalFlagsNoHooks                 = 64
	SignalFlagsMustCollect             = 128
	SignalFlagsDeprecated              = 256
)

type SignalMatchType int

const (
	SignalMatchTypeId        SignalMatchType = 1
	SignalMatchTypeDetail                    = 2
	SignalMatchTypeClosure                   = 4
	SignalMatchTypeFunc                      = 8
	SignalMatchTypeData                      = 16
	SignalMatchTypeUnblocked                 = 32
)

type TypeDebugFlags int

const (
	TypeDebugFlagsNone          TypeDebugFlags = 0
	TypeDebugFlagsObjects                      = 1
	TypeDebugFlagsSignals                      = 2
	TypeDebugFlagsInstanceCount                = 4
	TypeDebugFlagsMask                         = 7
)

type TypeFlags int

const (
	TypeFlagsAbstract      TypeFlags = 16
	TypeFlagsValueAbstract           = 32
)

type TypeFundamentalFlags int

const (
	TypeFundamentalFlagsClassed        TypeFundamentalFlags = 1
	TypeFundamentalFlagsInstantiatable                      = 2
	TypeFundamentalFlagsDerivable                           = 4
	TypeFundamentalFlagsDeepDerivable                       = 8
)
