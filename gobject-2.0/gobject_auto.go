package gobject

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"
import "github.com/electricface/go-auto-gir/glib-2.0"

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

// BindProperty is a wrapper around g_object_bind_property().
func (source Object) BindProperty(source_property string, target Object, target_property string, flags BindingFlags) Binding {
	source_property0 := (*C.gchar)(C.CString(source_property))
	target_property0 := (*C.gchar)(C.CString(target_property))
	ret0 := C.g_object_bind_property(C.gpointer(source.native()), source_property0, C.gpointer(target.native()), target_property0, C.GBindingFlags(flags))
	C.free(unsafe.Pointer(source_property0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(target_property0)) /*ch:<stdlib.h>*/
	return wrapBinding(ret0)
}

// BindPropertyWithClosures is a wrapper around g_object_bind_property_with_closures().
func (source Object) BindPropertyWithClosures(source_property string, target Object, target_property string, flags BindingFlags, transform_to Closure, transform_from Closure) Binding {
	source_property0 := (*C.gchar)(C.CString(source_property))
	target_property0 := (*C.gchar)(C.CString(target_property))
	ret0 := C.g_object_bind_property_with_closures(C.gpointer(source.native()), source_property0, C.gpointer(target.native()), target_property0, C.GBindingFlags(flags), transform_to.native(), transform_from.native())
	C.free(unsafe.Pointer(source_property0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(target_property0)) /*ch:<stdlib.h>*/
	return wrapBinding(ret0)
}

// ForceFloating is a wrapper around g_object_force_floating().
func (object Object) ForceFloating() {
	C.g_object_force_floating(object.native())
}

// FreezeNotify is a wrapper around g_object_freeze_notify().
func (object Object) FreezeNotify() {
	C.g_object_freeze_notify(object.native())
}

// GetData is a wrapper around g_object_get_data().
func (object Object) GetData(key string) unsafe.Pointer {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_object_get_data(object.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return unsafe.Pointer(ret0)
}

// GetProperty is a wrapper around g_object_get_property().
func (object Object) GetProperty(property_name string, value Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.g_object_get_property(object.native(), property_name0, value.native())
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// GetQdata is a wrapper around g_object_get_qdata().
func (object Object) GetQdata(quark /*gir:GLib*/ glib.Quark) unsafe.Pointer {
	ret0 := C.g_object_get_qdata(object.native(), C.GQuark(quark))
	return unsafe.Pointer(ret0)
}

// IsFloating is a wrapper around g_object_is_floating().
func (object Object) IsFloating() bool {
	ret0 := C.g_object_is_floating(C.gpointer(object.native()))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Notify is a wrapper around g_object_notify().
func (object Object) Notify(property_name string) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.g_object_notify(object.native(), property_name0)
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// NotifyByPspec is a wrapper around g_object_notify_by_pspec().
func (object Object) NotifyByPspec(pspec ParamSpec) {
	C.g_object_notify_by_pspec(object.native(), pspec.native())
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

// RunDispose is a wrapper around g_object_run_dispose().
func (object Object) RunDispose() {
	C.g_object_run_dispose(object.native())
}

// SetData is a wrapper around g_object_set_data().
func (object Object) SetData(key string, data unsafe.Pointer) {
	key0 := (*C.gchar)(C.CString(key))
	C.g_object_set_data(object.native(), key0, C.gpointer(data))
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// SetProperty is a wrapper around g_object_set_property().
func (object Object) SetProperty(property_name string, value Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.g_object_set_property(object.native(), property_name0, value.native())
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// SetQdata is a wrapper around g_object_set_qdata().
func (object Object) SetQdata(quark /*gir:GLib*/ glib.Quark, data unsafe.Pointer) {
	C.g_object_set_qdata(object.native(), C.GQuark(quark), C.gpointer(data))
}

// StealData is a wrapper around g_object_steal_data().
func (object Object) StealData(key string) unsafe.Pointer {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_object_steal_data(object.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return unsafe.Pointer(ret0)
}

// StealQdata is a wrapper around g_object_steal_qdata().
func (object Object) StealQdata(quark /*gir:GLib*/ glib.Quark) unsafe.Pointer {
	ret0 := C.g_object_steal_qdata(object.native(), C.GQuark(quark))
	return unsafe.Pointer(ret0)
}

// ThawNotify is a wrapper around g_object_thaw_notify().
func (object Object) ThawNotify() {
	C.g_object_thaw_notify(object.native())
}

// Unref is a wrapper around g_object_unref().
func (object Object) Unref() {
	C.g_object_unref(C.gpointer(object.native()))
}

// WatchClosure is a wrapper around g_object_watch_closure().
func (object Object) WatchClosure(closure Closure) {
	C.g_object_watch_closure(object.native(), closure.native())
}

// ObjectInterfaceFindProperty is a wrapper around g_object_interface_find_property().
func ObjectInterfaceFindProperty(g_iface TypeInterface, property_name string) ParamSpec {
	property_name0 := (*C.gchar)(C.CString(property_name))
	ret0 := C.g_object_interface_find_property(C.gpointer(g_iface.native()), property_name0)
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
	return wrapParamSpec(ret0)
}

// ObjectInterfaceInstallProperty is a wrapper around g_object_interface_install_property().
func ObjectInterfaceInstallProperty(g_iface TypeInterface, pspec ParamSpec) {
	C.g_object_interface_install_property(C.gpointer(g_iface.native()), pspec.native())
}

// ObjectInterfaceListProperties is a wrapper around g_object_interface_list_properties().
func ObjectInterfaceListProperties(g_iface TypeInterface) []ParamSpec {
	var n_properties_p0 C.guint
	ret0 := C.g_object_interface_list_properties(C.gpointer(g_iface.native()), &n_properties_p0)
	var ret0Slice []*C.GParamSpec
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(n_properties_p0)) /*go:.util*/
	ret := make([]ParamSpec, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = wrapParamSpec(elem)
	}
	C.g_free(C.gpointer(ret0))
	return ret
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

// Object Binding
type Binding struct {
	Object
}

func (v Binding) native() *C.GBinding {
	return (*C.GBinding)(v.Ptr)
}
func wrapBinding(p *C.GBinding) (v Binding) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBinding(p unsafe.Pointer) (v Binding) {
	v.Ptr = p
	return
}

// GetFlags is a wrapper around g_binding_get_flags().
func (binding Binding) GetFlags() BindingFlags {
	ret0 := C.g_binding_get_flags(binding.native())
	return BindingFlags(ret0)
}

// GetSource is a wrapper around g_binding_get_source().
func (binding Binding) GetSource() Object {
	ret0 := C.g_binding_get_source(binding.native())
	return wrapObject(ret0)
}

// GetSourceProperty is a wrapper around g_binding_get_source_property().
func (binding Binding) GetSourceProperty() string {
	ret0 := C.g_binding_get_source_property(binding.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTarget is a wrapper around g_binding_get_target().
func (binding Binding) GetTarget() Object {
	ret0 := C.g_binding_get_target(binding.native())
	return wrapObject(ret0)
}

// GetTargetProperty is a wrapper around g_binding_get_target_property().
func (binding Binding) GetTargetProperty() string {
	ret0 := C.g_binding_get_target_property(binding.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Unbind is a wrapper around g_binding_unbind().
func (binding Binding) Unbind() {
	C.g_binding_unbind(binding.native())
}

// Object ParamSpec
type ParamSpec struct {
	Ptr unsafe.Pointer
}

func (v ParamSpec) native() *C.GParamSpec {
	return (*C.GParamSpec)(v.Ptr)
}
func wrapParamSpec(p *C.GParamSpec) (v ParamSpec) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapParamSpec(p unsafe.Pointer) (v ParamSpec) {
	v.Ptr = p
	return
}

// GetBlurb is a wrapper around g_param_spec_get_blurb().
func (pspec ParamSpec) GetBlurb() string {
	ret0 := C.g_param_spec_get_blurb(pspec.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDefaultValue is a wrapper around g_param_spec_get_default_value().
func (pspec ParamSpec) GetDefaultValue() Value {
	ret0 := C.g_param_spec_get_default_value(pspec.native())
	return wrapValue(ret0)
}

// GetName is a wrapper around g_param_spec_get_name().
func (pspec ParamSpec) GetName() string {
	ret0 := C.g_param_spec_get_name(pspec.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNameQuark is a wrapper around g_param_spec_get_name_quark().
func (pspec ParamSpec) GetNameQuark() /*gir:GLib*/ glib.Quark {
	ret0 := C.g_param_spec_get_name_quark(pspec.native())
	return /*gir:GLib*/ glib.Quark(ret0)
}

// GetNick is a wrapper around g_param_spec_get_nick().
func (pspec ParamSpec) GetNick() string {
	ret0 := C.g_param_spec_get_nick(pspec.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetQdata is a wrapper around g_param_spec_get_qdata().
func (pspec ParamSpec) GetQdata(quark /*gir:GLib*/ glib.Quark) unsafe.Pointer {
	ret0 := C.g_param_spec_get_qdata(pspec.native(), C.GQuark(quark))
	return unsafe.Pointer(ret0)
}

// GetRedirectTarget is a wrapper around g_param_spec_get_redirect_target().
func (pspec ParamSpec) GetRedirectTarget() ParamSpec {
	ret0 := C.g_param_spec_get_redirect_target(pspec.native())
	return wrapParamSpec(ret0)
}

// Ref is a wrapper around g_param_spec_ref().
func (pspec ParamSpec) Ref() ParamSpec {
	ret0 := C.g_param_spec_ref(pspec.native())
	return wrapParamSpec(ret0)
}

// RefSink is a wrapper around g_param_spec_ref_sink().
func (pspec ParamSpec) RefSink() ParamSpec {
	ret0 := C.g_param_spec_ref_sink(pspec.native())
	return wrapParamSpec(ret0)
}

// SetQdata is a wrapper around g_param_spec_set_qdata().
func (pspec ParamSpec) SetQdata(quark /*gir:GLib*/ glib.Quark, data unsafe.Pointer) {
	C.g_param_spec_set_qdata(pspec.native(), C.GQuark(quark), C.gpointer(data))
}

// Sink is a wrapper around g_param_spec_sink().
func (pspec ParamSpec) Sink() {
	C.g_param_spec_sink(pspec.native())
}

// StealQdata is a wrapper around g_param_spec_steal_qdata().
func (pspec ParamSpec) StealQdata(quark /*gir:GLib*/ glib.Quark) unsafe.Pointer {
	ret0 := C.g_param_spec_steal_qdata(pspec.native(), C.GQuark(quark))
	return unsafe.Pointer(ret0)
}

// Unref is a wrapper around g_param_spec_unref().
func (pspec ParamSpec) Unref() {
	C.g_param_spec_unref(pspec.native())
}

// Struct TypeInterface
type TypeInterface struct {
	Ptr unsafe.Pointer
}

func (v TypeInterface) native() *C.GTypeInterface {
	return (*C.GTypeInterface)(v.Ptr)
}
func wrapTypeInterface(p *C.GTypeInterface) TypeInterface {
	return TypeInterface{unsafe.Pointer(p)}
}
func WrapTypeInterface(p unsafe.Pointer) TypeInterface {
	return TypeInterface{p}
}

// PeekParent is a wrapper around g_type_interface_peek_parent().
func (g_iface TypeInterface) PeekParent() TypeInterface {
	ret0 := C.g_type_interface_peek_parent(C.gpointer(g_iface.native()))
	return WrapTypeInterface(unsafe.Pointer(ret0))
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
