package gobject

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
#include <stdlib.h>
*/
import "C"
import "github.com/electricface/go-auto-gir/glib-2.0"
import "github.com/electricface/go-auto-gir/util"
import "unsafe"

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
func (v Value) IsNil() bool {
	return v.Ptr == nil
}
func IWrapValue(p unsafe.Pointer) interface{} {
	return WrapValue(p)
}

// Copy is a wrapper around g_value_copy().
func (src_value Value) Copy(dest_value Value) {
	C.g_value_copy(src_value.native(), dest_value.native())
}

// DupBoxed is a wrapper around g_value_dup_boxed().
func (value Value) DupBoxed() unsafe.Pointer {
	ret0 := C.g_value_dup_boxed(value.native())
	return unsafe.Pointer(ret0)
}

// DupObject is a wrapper around g_value_dup_object().
func (value Value) DupObject() Object {
	ret0 := C.g_value_dup_object(value.native())
	return WrapObject(unsafe.Pointer(ret0))
}

// DupParam is a wrapper around g_value_dup_param().
func (value Value) DupParam() ParamSpec {
	ret0 := C.g_value_dup_param(value.native())
	return wrapParamSpec(ret0)
}

// DupString is a wrapper around g_value_dup_string().
func (value Value) DupString() string {
	ret0 := C.g_value_dup_string(value.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// DupVariant is a wrapper around g_value_dup_variant().
func (value Value) DupVariant() glib.Variant {
	ret0 := C.g_value_dup_variant(value.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// FitsPointer is a wrapper around g_value_fits_pointer().
func (value Value) FitsPointer() bool {
	ret0 := C.g_value_fits_pointer(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetBoolean is a wrapper around g_value_get_boolean().
func (value Value) GetBoolean() bool {
	ret0 := C.g_value_get_boolean(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetBoxed is a wrapper around g_value_get_boxed().
func (value Value) GetBoxed() unsafe.Pointer {
	ret0 := C.g_value_get_boxed(value.native())
	return unsafe.Pointer(ret0)
}

// GetDouble is a wrapper around g_value_get_double().
func (value Value) GetDouble() float64 {
	ret0 := C.g_value_get_double(value.native())
	return float64(ret0)
}

// GetEnum is a wrapper around g_value_get_enum().
func (value Value) GetEnum() int {
	ret0 := C.g_value_get_enum(value.native())
	return int(ret0)
}

// GetFlags is a wrapper around g_value_get_flags().
func (value Value) GetFlags() uint {
	ret0 := C.g_value_get_flags(value.native())
	return uint(ret0)
}

// GetFloat is a wrapper around g_value_get_float().
func (value Value) GetFloat() float32 {
	ret0 := C.g_value_get_float(value.native())
	return float32(ret0)
}

// GetInt is a wrapper around g_value_get_int().
func (value Value) GetInt() int {
	ret0 := C.g_value_get_int(value.native())
	return int(ret0)
}

// GetInt64 is a wrapper around g_value_get_int64().
func (value Value) GetInt64() int64 {
	ret0 := C.g_value_get_int64(value.native())
	return int64(ret0)
}

// GetLong is a wrapper around g_value_get_long().
func (value Value) GetLong() int {
	ret0 := C.g_value_get_long(value.native())
	return int(ret0)
}

// GetObject is a wrapper around g_value_get_object().
func (value Value) GetObject() Object {
	ret0 := C.g_value_get_object(value.native())
	return WrapObject(unsafe.Pointer(ret0))
}

// GetParam is a wrapper around g_value_get_param().
func (value Value) GetParam() ParamSpec {
	ret0 := C.g_value_get_param(value.native())
	return wrapParamSpec(ret0)
}

// GetPointer is a wrapper around g_value_get_pointer().
func (value Value) GetPointer() unsafe.Pointer {
	ret0 := C.g_value_get_pointer(value.native())
	return unsafe.Pointer(ret0)
}

// GetSchar is a wrapper around g_value_get_schar().
func (value Value) GetSchar() int8 {
	ret0 := C.g_value_get_schar(value.native())
	return int8(ret0)
}

// GetString is a wrapper around g_value_get_string().
func (value Value) GetString() string {
	ret0 := C.g_value_get_string(value.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUchar is a wrapper around g_value_get_uchar().
func (value Value) GetUchar() byte {
	ret0 := C.g_value_get_uchar(value.native())
	return byte(ret0)
}

// GetUint is a wrapper around g_value_get_uint().
func (value Value) GetUint() uint {
	ret0 := C.g_value_get_uint(value.native())
	return uint(ret0)
}

// GetUint64 is a wrapper around g_value_get_uint64().
func (value Value) GetUint64() uint64 {
	ret0 := C.g_value_get_uint64(value.native())
	return uint64(ret0)
}

// GetUlong is a wrapper around g_value_get_ulong().
func (value Value) GetUlong() uint {
	ret0 := C.g_value_get_ulong(value.native())
	return uint(ret0)
}

// GetVariant is a wrapper around g_value_get_variant().
func (value Value) GetVariant() glib.Variant {
	ret0 := C.g_value_get_variant(value.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// InitFromInstance is a wrapper around g_value_init_from_instance().
func (value Value) InitFromInstance(instance TypeInstance) {
	C.g_value_init_from_instance(value.native(), C.gpointer(instance.native()))
}

// PeekPointer is a wrapper around g_value_peek_pointer().
func (value Value) PeekPointer() unsafe.Pointer {
	ret0 := C.g_value_peek_pointer(value.native())
	return unsafe.Pointer(ret0)
}

// Reset is a wrapper around g_value_reset().
func (value Value) Reset() Value {
	ret0 := C.g_value_reset(value.native())
	return wrapValue(ret0)
}

// SetBoolean is a wrapper around g_value_set_boolean().
func (value Value) SetBoolean(v_boolean bool) {
	C.g_value_set_boolean(value.native(), C.gboolean(util.Bool2Int(v_boolean)) /*go:.util*/)
}

// SetBoxed is a wrapper around g_value_set_boxed().
func (value Value) SetBoxed(v_boxed unsafe.Pointer) {
	C.g_value_set_boxed(value.native(), C.gconstpointer(v_boxed))
}

// SetDouble is a wrapper around g_value_set_double().
func (value Value) SetDouble(v_double float64) {
	C.g_value_set_double(value.native(), C.gdouble(v_double))
}

// SetEnum is a wrapper around g_value_set_enum().
func (value Value) SetEnum(v_enum int) {
	C.g_value_set_enum(value.native(), C.gint(v_enum))
}

// SetFlags is a wrapper around g_value_set_flags().
func (value Value) SetFlags(v_flags uint) {
	C.g_value_set_flags(value.native(), C.guint(v_flags))
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

// SetLong is a wrapper around g_value_set_long().
func (value Value) SetLong(v_long int) {
	C.g_value_set_long(value.native(), C.glong(v_long))
}

// SetObject is a wrapper around g_value_set_object().
func (value Value) SetObject(v_object Object) {
	C.g_value_set_object(value.native(), C.gpointer(v_object.native()))
}

// SetParam is a wrapper around g_value_set_param().
func (value Value) SetParam(param ParamSpec) {
	C.g_value_set_param(value.native(), param.native())
}

// SetPointer is a wrapper around g_value_set_pointer().
func (value Value) SetPointer(v_pointer unsafe.Pointer) {
	C.g_value_set_pointer(value.native(), C.gpointer(v_pointer))
}

// SetSchar is a wrapper around g_value_set_schar().
func (value Value) SetSchar(v_char int8) {
	C.g_value_set_schar(value.native(), C.gint8(v_char))
}

// SetStaticBoxed is a wrapper around g_value_set_static_boxed().
func (value Value) SetStaticBoxed(v_boxed unsafe.Pointer) {
	C.g_value_set_static_boxed(value.native(), C.gconstpointer(v_boxed))
}

// SetStaticString is a wrapper around g_value_set_static_string().
func (value Value) SetStaticString(v_string string) {
	v_string0 := (*C.gchar)(C.CString(v_string))
	C.g_value_set_static_string(value.native(), v_string0)
	C.free(unsafe.Pointer(v_string0)) /*ch:<stdlib.h>*/
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

// SetUlong is a wrapper around g_value_set_ulong().
func (value Value) SetUlong(v_ulong uint) {
	C.g_value_set_ulong(value.native(), C.gulong(v_ulong))
}

// SetVariant is a wrapper around g_value_set_variant().
func (value Value) SetVariant(variant glib.Variant) {
	C.g_value_set_variant(value.native(), (*C.GVariant)(variant.Ptr))
}

// TakeBoxed is a wrapper around g_value_take_boxed().
func (value Value) TakeBoxed(v_boxed unsafe.Pointer) {
	C.g_value_take_boxed(value.native(), C.gconstpointer(v_boxed))
}

// TakeObject is a wrapper around g_value_take_object().
func (value Value) TakeObject(v_object unsafe.Pointer) {
	C.g_value_take_object(value.native(), C.gpointer(v_object))
}

// TakeParam is a wrapper around g_value_take_param().
func (value Value) TakeParam(param ParamSpec) {
	C.g_value_take_param(value.native(), param.native())
}

// TakeString is a wrapper around g_value_take_string().
func (value Value) TakeString(v_string string) {
	v_string0 := (*C.gchar)(C.CString(v_string))
	C.g_value_take_string(value.native(), v_string0)
	C.free(unsafe.Pointer(v_string0)) /*ch:<stdlib.h>*/
}

// TakeVariant is a wrapper around g_value_take_variant().
func (value Value) TakeVariant(variant glib.Variant) {
	C.g_value_take_variant(value.native(), (*C.GVariant)(variant.Ptr))
}

// Transform is a wrapper around g_value_transform().
func (src_value Value) Transform(dest_value Value) bool {
	ret0 := C.g_value_transform(src_value.native(), dest_value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
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
func wrapObject(p *C.GObject) Object {
	return Object{unsafe.Pointer(p)}
}
func WrapObject(p unsafe.Pointer) Object {
	return Object{p}
}
func (v Object) IsNil() bool {
	return v.Ptr == nil
}
func IWrapObject(p unsafe.Pointer) interface{} {
	return WrapObject(p)
}
func (v Object) GetType() Type {
	return Type(C.g_object_get_type())
}
func (v Object) GetGValueGetter() GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapObject(unsafe.Pointer(ptr)), nil
	}
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
func (v Closure) IsNil() bool {
	return v.Ptr == nil
}
func IWrapClosure(p unsafe.Pointer) interface{} {
	return WrapClosure(p)
}

// ClosureNewObject is a wrapper around g_closure_new_object().
func ClosureNewObject(sizeof_closure uint, object Object) Closure {
	ret0 := C.g_closure_new_object(C.guint(sizeof_closure), object.native())
	return wrapClosure(ret0)
}

// ClosureNewSimple is a wrapper around g_closure_new_simple().
func ClosureNewSimple(sizeof_closure uint, data unsafe.Pointer) Closure {
	ret0 := C.g_closure_new_simple(C.guint(sizeof_closure), C.gpointer(data))
	return wrapClosure(ret0)
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
func (v Binding) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBinding(p unsafe.Pointer) interface{} {
	return WrapBinding(p)
}
func (v Binding) GetType() Type {
	return Type(C.g_binding_get_type())
}
func (v Binding) GetGValueGetter() GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBinding(unsafe.Pointer(ptr)), nil
	}
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
func wrapParamSpec(p *C.GParamSpec) ParamSpec {
	return ParamSpec{unsafe.Pointer(p)}
}
func WrapParamSpec(p unsafe.Pointer) ParamSpec {
	return ParamSpec{p}
}
func (v ParamSpec) IsNil() bool {
	return v.Ptr == nil
}
func IWrapParamSpec(p unsafe.Pointer) interface{} {
	return WrapParamSpec(p)
}
func (v ParamSpec) GetGValueGetter() GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapParamSpec(unsafe.Pointer(ptr)), nil
	}
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
func (v TypeInterface) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTypeInterface(p unsafe.Pointer) interface{} {
	return WrapTypeInterface(p)
}

// PeekParent is a wrapper around g_type_interface_peek_parent().
func (g_iface TypeInterface) PeekParent() TypeInterface {
	ret0 := C.g_type_interface_peek_parent(C.gpointer(g_iface.native()))
	return WrapTypeInterface(unsafe.Pointer(ret0))
}

// Struct TypeInstance
type TypeInstance struct {
	Ptr unsafe.Pointer
}

func (v TypeInstance) native() *C.GTypeInstance {
	return (*C.GTypeInstance)(v.Ptr)
}
func wrapTypeInstance(p *C.GTypeInstance) TypeInstance {
	return TypeInstance{unsafe.Pointer(p)}
}
func WrapTypeInstance(p unsafe.Pointer) TypeInstance {
	return TypeInstance{p}
}
func (v TypeInstance) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTypeInstance(p unsafe.Pointer) interface{} {
	return WrapTypeInstance(p)
}

// Struct ValueArray
type ValueArray struct {
	Ptr unsafe.Pointer
}

func (v ValueArray) native() *C.GValueArray {
	return (*C.GValueArray)(v.Ptr)
}
func wrapValueArray(p *C.GValueArray) ValueArray {
	return ValueArray{unsafe.Pointer(p)}
}
func WrapValueArray(p unsafe.Pointer) ValueArray {
	return ValueArray{p}
}
func (v ValueArray) IsNil() bool {
	return v.Ptr == nil
}
func IWrapValueArray(p unsafe.Pointer) interface{} {
	return WrapValueArray(p)
}

type Type uint
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
