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

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue
	ret0 := C.g_value_get_pointer(value.native())
	return unsafe.Pointer(ret0)
}

// GetString is a wrapper around g_value_get_string().
func (value Value) GetString() string {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue
	ret0 := C.g_value_get_string(value.native())
	ret := C.GoString((*C.char)(ret0))
	defer C.g_free(C.gpointer(ret0))
	return ret
}

// SetBoolean is a wrapper around g_value_set_boolean().
func (value Value) SetBoolean(v_boolean bool) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_boolean
	// Var for C: v_boolean0
	// Type for Go: bool
	// Type for C: C.gboolean
	C.g_value_set_boolean(value.native(), C.gboolean(util.Bool2Int(v_boolean)) /*go:.util*/)
}

// SetDouble is a wrapper around g_value_set_double().
func (value Value) SetDouble(v_double float64) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_double
	// Var for C: v_double0
	// Type for Go: float64
	// Type for C: C.gdouble
	C.g_value_set_double(value.native(), C.gdouble(v_double))
}

// SetFloat is a wrapper around g_value_set_float().
func (value Value) SetFloat(v_float float32) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_float
	// Var for C: v_float0
	// Type for Go: float32
	// Type for C: C.gfloat
	C.g_value_set_float(value.native(), C.gfloat(v_float))
}

// SetInstance is a wrapper around g_value_set_instance().
func (value Value) SetInstance(instance unsafe.Pointer) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: instance
	// Var for C: instance0
	// Type for Go: unsafe.Pointer
	// Type for C: C.gpointer
	C.g_value_set_instance(value.native(), C.gpointer(instance))
}

// SetInt is a wrapper around g_value_set_int().
func (value Value) SetInt(v_int int) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_int
	// Var for C: v_int0
	// Type for Go: int
	// Type for C: C.gint
	C.g_value_set_int(value.native(), C.gint(v_int))
}

// SetInt64 is a wrapper around g_value_set_int64().
func (value Value) SetInt64(v_int64 int64) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_int64
	// Var for C: v_int640
	// Type for Go: int64
	// Type for C: C.gint64
	C.g_value_set_int64(value.native(), C.gint64(v_int64))
}

// SetPointer is a wrapper around g_value_set_pointer().
func (value Value) SetPointer(v_pointer unsafe.Pointer) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_pointer
	// Var for C: v_pointer0
	// Type for Go: unsafe.Pointer
	// Type for C: C.gpointer
	C.g_value_set_pointer(value.native(), C.gpointer(v_pointer))
}

// SetSchar is a wrapper around g_value_set_schar().
func (value Value) SetSchar(v_char int8) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_char
	// Var for C: v_char0
	// Type for Go: int8
	// Type for C: C.gint8
	C.g_value_set_schar(value.native(), C.gint8(v_char))
}

// SetString is a wrapper around g_value_set_string().
func (value Value) SetString(v_string string) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_string
	// Var for C: v_string0
	// Type for Go: string
	// Type for C: *C.gchar
	v_string0 := (*C.gchar)(C.CString(v_string))
	defer C.free(unsafe.Pointer(v_string0)) /*ch:<stdlib.h>*/
	C.g_value_set_string(value.native(), v_string0)
}

// SetUchar is a wrapper around g_value_set_uchar().
func (value Value) SetUchar(v_uchar byte) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_uchar
	// Var for C: v_uchar0
	// Type for Go: byte
	// Type for C: C.guchar
	C.g_value_set_uchar(value.native(), C.guchar(v_uchar))
}

// SetUint is a wrapper around g_value_set_uint().
func (value Value) SetUint(v_uint uint) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_uint
	// Var for C: v_uint0
	// Type for Go: uint
	// Type for C: C.guint
	C.g_value_set_uint(value.native(), C.guint(v_uint))
}

// SetUint64 is a wrapper around g_value_set_uint64().
func (value Value) SetUint64(v_uint64 uint64) {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue

	// Var for Go: v_uint64
	// Var for C: v_uint640
	// Type for Go: uint64
	// Type for C: C.guint64
	C.g_value_set_uint64(value.native(), C.guint64(v_uint64))
}

// Unset is a wrapper around g_value_unset().
func (value Value) Unset() {

	// Var for Go: value
	// Var for C: value0
	// Type for Go: Value
	// Type for C: *C.GValue
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

type BindingFlags C.GBindingFlags

const (
	BindingFlagsDefault       BindingFlags = 0
	BindingFlagsBidirectional              = 1
	BindingFlagsSyncCreate                 = 2
	BindingFlagsInvertBoolean              = 4
)

type ConnectFlags C.GConnectFlags

const (
	ConnectFlagsAfter   ConnectFlags = 1
	ConnectFlagsSwapped              = 2
)

type ParamFlags C.GParamFlags

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

type SignalFlags C.GSignalFlags

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

type SignalMatchType C.GSignalMatchType

const (
	SignalMatchTypeId        SignalMatchType = 1
	SignalMatchTypeDetail                    = 2
	SignalMatchTypeClosure                   = 4
	SignalMatchTypeFunc                      = 8
	SignalMatchTypeData                      = 16
	SignalMatchTypeUnblocked                 = 32
)

type TypeDebugFlags C.GTypeDebugFlags

const (
	TypeDebugFlagsNone          TypeDebugFlags = 0
	TypeDebugFlagsObjects                      = 1
	TypeDebugFlagsSignals                      = 2
	TypeDebugFlagsInstanceCount                = 4
	TypeDebugFlagsMask                         = 7
)

type TypeFlags C.GTypeFlags

const (
	TypeFlagsAbstract      TypeFlags = 16
	TypeFlagsValueAbstract           = 32
)

type TypeFundamentalFlags C.GTypeFundamentalFlags

const (
	TypeFundamentalFlagsClassed        TypeFundamentalFlags = 1
	TypeFundamentalFlagsInstantiatable                      = 2
	TypeFundamentalFlagsDerivable                           = 4
	TypeFundamentalFlagsDeepDerivable                       = 8
)
