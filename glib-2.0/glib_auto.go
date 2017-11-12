package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import "github.com/electricface/go-auto-gir/util"
import "unsafe"

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
func (v Variant) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVariant(p unsafe.Pointer) interface{} {
	return WrapVariant(p)
}

// VariantNewArray is a wrapper around g_variant_new_array().
func VariantNewArray(child_type VariantType, children []Variant) Variant {
	children0 := make([]*C.GVariant, len(children))
	for idx, elemG := range children {
		children0[idx] = elemG.native()
	}
	var children0Ptr **C.GVariant
	if len(children0) > 0 {
		children0Ptr = &children0[0]
	}
	ret0 := C.g_variant_new_array(child_type.native(), children0Ptr, C.gsize(len(children)))
	return wrapVariant(ret0)
}

// VariantNewBoolean is a wrapper around g_variant_new_boolean().
func VariantNewBoolean(value bool) Variant {
	ret0 := C.g_variant_new_boolean(C.gboolean(util.Bool2Int(value)) /*go:.util*/)
	return wrapVariant(ret0)
}

// VariantNewByte is a wrapper around g_variant_new_byte().
func VariantNewByte(value byte) Variant {
	ret0 := C.g_variant_new_byte(C.guchar(value))
	return wrapVariant(ret0)
}

// VariantNewBytestring is a wrapper around g_variant_new_bytestring().
func VariantNewBytestring(string []byte) Variant {
	string0 := make([]C.gchar, len(string))
	for idx, elemG := range string {
		string0[idx] = C.gchar(elemG)
	}
	var string0Ptr *C.gchar
	if len(string0) > 0 {
		string0Ptr = &string0[0]
	}
	ret0 := C.g_variant_new_bytestring(string0Ptr)
	return wrapVariant(ret0)
}

// VariantNewBytestringArray is a wrapper around g_variant_new_bytestring_array().
func VariantNewBytestringArray(strv []string) Variant {
	strv0 := make([]*C.gchar, len(strv))
	for idx, elemG := range strv {
		elem := (*C.gchar)(C.CString(elemG))
		strv0[idx] = elem
	}
	var strv0Ptr **C.gchar
	if len(strv0) > 0 {
		strv0Ptr = &strv0[0]
	}
	ret0 := C.g_variant_new_bytestring_array(strv0Ptr, C.gssize(len(strv)))
	for _, elem := range strv0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return wrapVariant(ret0)
}

// VariantNewDictEntry is a wrapper around g_variant_new_dict_entry().
func VariantNewDictEntry(key Variant, value Variant) Variant {
	ret0 := C.g_variant_new_dict_entry(key.native(), value.native())
	return wrapVariant(ret0)
}

// VariantNewDouble is a wrapper around g_variant_new_double().
func VariantNewDouble(value float64) Variant {
	ret0 := C.g_variant_new_double(C.gdouble(value))
	return wrapVariant(ret0)
}

// VariantNewFixedArray is a wrapper around g_variant_new_fixed_array().
func VariantNewFixedArray(element_type VariantType, elements unsafe.Pointer, n_elements uint, element_size uint) Variant {
	ret0 := C.g_variant_new_fixed_array(element_type.native(), C.gconstpointer(elements), C.gsize(n_elements), C.gsize(element_size))
	return wrapVariant(ret0)
}

// VariantNewFromBytes is a wrapper around g_variant_new_from_bytes().
func VariantNewFromBytes(type_ VariantType, bytes Bytes, trusted bool) Variant {
	ret0 := C.g_variant_new_from_bytes(type_.native(), bytes.native(), C.gboolean(util.Bool2Int(trusted)) /*go:.util*/)
	return wrapVariant(ret0)
}

// VariantNewHandle is a wrapper around g_variant_new_handle().
func VariantNewHandle(value int32) Variant {
	ret0 := C.g_variant_new_handle(C.gint32(value))
	return wrapVariant(ret0)
}

// VariantNewInt16 is a wrapper around g_variant_new_int16().
func VariantNewInt16(value int16) Variant {
	ret0 := C.g_variant_new_int16(C.gint16(value))
	return wrapVariant(ret0)
}

// VariantNewInt32 is a wrapper around g_variant_new_int32().
func VariantNewInt32(value int32) Variant {
	ret0 := C.g_variant_new_int32(C.gint32(value))
	return wrapVariant(ret0)
}

// VariantNewInt64 is a wrapper around g_variant_new_int64().
func VariantNewInt64(value int64) Variant {
	ret0 := C.g_variant_new_int64(C.gint64(value))
	return wrapVariant(ret0)
}

// VariantNewMaybe is a wrapper around g_variant_new_maybe().
func VariantNewMaybe(child_type VariantType, child Variant) Variant {
	ret0 := C.g_variant_new_maybe(child_type.native(), child.native())
	return wrapVariant(ret0)
}

// VariantNewObjectPath is a wrapper around g_variant_new_object_path().
func VariantNewObjectPath(object_path string) Variant {
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_variant_new_object_path(object_path0)
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return wrapVariant(ret0)
}

// VariantNewObjv is a wrapper around g_variant_new_objv().
func VariantNewObjv(strv []string) Variant {
	strv0 := make([]*C.gchar, len(strv))
	for idx, elemG := range strv {
		elem := (*C.gchar)(C.CString(elemG))
		strv0[idx] = elem
	}
	var strv0Ptr **C.gchar
	if len(strv0) > 0 {
		strv0Ptr = &strv0[0]
	}
	ret0 := C.g_variant_new_objv(strv0Ptr, C.gssize(len(strv)))
	for _, elem := range strv0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return wrapVariant(ret0)
}

// VariantNewSignature is a wrapper around g_variant_new_signature().
func VariantNewSignature(signature string) Variant {
	signature0 := (*C.gchar)(C.CString(signature))
	ret0 := C.g_variant_new_signature(signature0)
	C.free(unsafe.Pointer(signature0)) /*ch:<stdlib.h>*/
	return wrapVariant(ret0)
}

// VariantNewString is a wrapper around g_variant_new_string().
func VariantNewString(string string) Variant {
	string0 := (*C.gchar)(C.CString(string))
	ret0 := C.g_variant_new_string(string0)
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	return wrapVariant(ret0)
}

// VariantNewStrv is a wrapper around g_variant_new_strv().
func VariantNewStrv(strv []string) Variant {
	strv0 := make([]*C.gchar, len(strv))
	for idx, elemG := range strv {
		elem := (*C.gchar)(C.CString(elemG))
		strv0[idx] = elem
	}
	var strv0Ptr **C.gchar
	if len(strv0) > 0 {
		strv0Ptr = &strv0[0]
	}
	ret0 := C.g_variant_new_strv(strv0Ptr, C.gssize(len(strv)))
	for _, elem := range strv0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return wrapVariant(ret0)
}

// VariantNewTakeString is a wrapper around g_variant_new_take_string().
func VariantNewTakeString(string string) Variant {
	string0 := (*C.gchar)(C.CString(string))
	ret0 := C.g_variant_new_take_string(string0)
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	return wrapVariant(ret0)
}

// VariantNewTuple is a wrapper around g_variant_new_tuple().
func VariantNewTuple(children []Variant) Variant {
	children0 := make([]*C.GVariant, len(children))
	for idx, elemG := range children {
		children0[idx] = elemG.native()
	}
	var children0Ptr **C.GVariant
	if len(children0) > 0 {
		children0Ptr = &children0[0]
	}
	ret0 := C.g_variant_new_tuple(children0Ptr, C.gsize(len(children)))
	return wrapVariant(ret0)
}

// VariantNewUint16 is a wrapper around g_variant_new_uint16().
func VariantNewUint16(value uint16) Variant {
	ret0 := C.g_variant_new_uint16(C.guint16(value))
	return wrapVariant(ret0)
}

// VariantNewUint32 is a wrapper around g_variant_new_uint32().
func VariantNewUint32(value uint32) Variant {
	ret0 := C.g_variant_new_uint32(C.guint32(value))
	return wrapVariant(ret0)
}

// VariantNewUint64 is a wrapper around g_variant_new_uint64().
func VariantNewUint64(value uint64) Variant {
	ret0 := C.g_variant_new_uint64(C.guint64(value))
	return wrapVariant(ret0)
}

// VariantNewVariant is a wrapper around g_variant_new_variant().
func VariantNewVariant(value Variant) Variant {
	ret0 := C.g_variant_new_variant(value.native())
	return wrapVariant(ret0)
}

// Byteswap is a wrapper around g_variant_byteswap().
func (value Variant) Byteswap() Variant {
	ret0 := C.g_variant_byteswap(value.native())
	return wrapVariant(ret0)
}

// CheckFormatString is a wrapper around g_variant_check_format_string().
func (value Variant) CheckFormatString(format_string string, copy_only bool) bool {
	format_string0 := (*C.gchar)(C.CString(format_string))
	ret0 := C.g_variant_check_format_string(value.native(), format_string0, C.gboolean(util.Bool2Int(copy_only)) /*go:.util*/)
	C.free(unsafe.Pointer(format_string0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))        /*go:.util*/
}

// Classify is a wrapper around g_variant_classify().
func (value Variant) Classify() VariantClass {
	ret0 := C.g_variant_classify(value.native())
	return VariantClass(ret0)
}

// Compare is a wrapper around g_variant_compare().
func (one Variant) Compare(two Variant) int {
	ret0 := C.g_variant_compare(C.gconstpointer(one.native()), C.gconstpointer(two.native()))
	return int(ret0)
}

// DupBytestring is a wrapper around g_variant_dup_bytestring().
func (value Variant) DupBytestring() []byte {
	var length0 C.gsize
	ret0 := C.g_variant_dup_bytestring(value.native(), &length0)
	var ret0Slice []C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]byte, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = byte(elem)
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// DupBytestringArray is a wrapper around g_variant_dup_bytestring_array().
func (value Variant) DupBytestringArray() []string {
	var length0 C.gsize
	ret0 := C.g_variant_dup_bytestring_array(value.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// DupObjv is a wrapper around g_variant_dup_objv().
func (value Variant) DupObjv() []string {
	var length0 C.gsize
	ret0 := C.g_variant_dup_objv(value.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// DupString is a wrapper around g_variant_dup_string().
func (value Variant) DupString() (string, uint) {
	var length0 C.gsize
	ret0 := C.g_variant_dup_string(value.native(), &length0)
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret, uint(length0)
}

// DupStrv is a wrapper around g_variant_dup_strv().
func (value Variant) DupStrv() []string {
	var length0 C.gsize
	ret0 := C.g_variant_dup_strv(value.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// Equal is a wrapper around g_variant_equal().
func (one Variant) Equal(two Variant) bool {
	ret0 := C.g_variant_equal(C.gconstpointer(one.native()), C.gconstpointer(two.native()))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetBoolean is a wrapper around g_variant_get_boolean().
func (value Variant) GetBoolean() bool {
	ret0 := C.g_variant_get_boolean(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetByte is a wrapper around g_variant_get_byte().
func (value Variant) GetByte() byte {
	ret0 := C.g_variant_get_byte(value.native())
	return byte(ret0)
}

// GetBytestring is a wrapper around g_variant_get_bytestring().
func (value Variant) GetBytestring() []byte {
	ret0 := C.g_variant_get_bytestring(value.native())
	var ret0Slice []C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(C.gchar(0))) /*go:.util*/) /*go:.util*/
	ret := make([]byte, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = byte(elem)
	}
	return ret
}

// GetBytestringArray is a wrapper around g_variant_get_bytestring_array().
func (value Variant) GetBytestringArray() []string {
	var length0 C.gsize
	ret0 := C.g_variant_get_bytestring_array(value.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetChildValue is a wrapper around g_variant_get_child_value().
func (value Variant) GetChildValue(index_ uint) Variant {
	ret0 := C.g_variant_get_child_value(value.native(), C.gsize(index_))
	return wrapVariant(ret0)
}

// GetData is a wrapper around g_variant_get_data().
func (value Variant) GetData() unsafe.Pointer {
	ret0 := C.g_variant_get_data(value.native())
	return unsafe.Pointer(ret0)
}

// GetDataAsBytes is a wrapper around g_variant_get_data_as_bytes().
func (value Variant) GetDataAsBytes() Bytes {
	ret0 := C.g_variant_get_data_as_bytes(value.native())
	return wrapBytes(ret0)
}

// GetDouble is a wrapper around g_variant_get_double().
func (value Variant) GetDouble() float64 {
	ret0 := C.g_variant_get_double(value.native())
	return float64(ret0)
}

// GetHandle is a wrapper around g_variant_get_handle().
func (value Variant) GetHandle() int32 {
	ret0 := C.g_variant_get_handle(value.native())
	return int32(ret0)
}

// GetInt16 is a wrapper around g_variant_get_int16().
func (value Variant) GetInt16() int16 {
	ret0 := C.g_variant_get_int16(value.native())
	return int16(ret0)
}

// GetInt32 is a wrapper around g_variant_get_int32().
func (value Variant) GetInt32() int32 {
	ret0 := C.g_variant_get_int32(value.native())
	return int32(ret0)
}

// GetInt64 is a wrapper around g_variant_get_int64().
func (value Variant) GetInt64() int64 {
	ret0 := C.g_variant_get_int64(value.native())
	return int64(ret0)
}

// GetMaybe is a wrapper around g_variant_get_maybe().
func (value Variant) GetMaybe() Variant {
	ret0 := C.g_variant_get_maybe(value.native())
	return wrapVariant(ret0)
}

// GetNormalForm is a wrapper around g_variant_get_normal_form().
func (value Variant) GetNormalForm() Variant {
	ret0 := C.g_variant_get_normal_form(value.native())
	return wrapVariant(ret0)
}

// GetObjv is a wrapper around g_variant_get_objv().
func (value Variant) GetObjv() []string {
	var length0 C.gsize
	ret0 := C.g_variant_get_objv(value.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetSize is a wrapper around g_variant_get_size().
func (value Variant) GetSize() uint {
	ret0 := C.g_variant_get_size(value.native())
	return uint(ret0)
}

// GetStrv is a wrapper around g_variant_get_strv().
func (value Variant) GetStrv() []string {
	var length0 C.gsize
	ret0 := C.g_variant_get_strv(value.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetType is a wrapper around g_variant_get_type().
func (value Variant) GetType() VariantType {
	ret0 := C.g_variant_get_type(value.native())
	return wrapVariantType(ret0)
}

// GetTypeString is a wrapper around g_variant_get_type_string().
func (value Variant) GetTypeString() string {
	ret0 := C.g_variant_get_type_string(value.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUint16 is a wrapper around g_variant_get_uint16().
func (value Variant) GetUint16() uint16 {
	ret0 := C.g_variant_get_uint16(value.native())
	return uint16(ret0)
}

// GetUint32 is a wrapper around g_variant_get_uint32().
func (value Variant) GetUint32() uint32 {
	ret0 := C.g_variant_get_uint32(value.native())
	return uint32(ret0)
}

// GetUint64 is a wrapper around g_variant_get_uint64().
func (value Variant) GetUint64() uint64 {
	ret0 := C.g_variant_get_uint64(value.native())
	return uint64(ret0)
}

// GetVariant is a wrapper around g_variant_get_variant().
func (value Variant) GetVariant() Variant {
	ret0 := C.g_variant_get_variant(value.native())
	return wrapVariant(ret0)
}

// Hash is a wrapper around g_variant_hash().
func (value Variant) Hash() uint {
	ret0 := C.g_variant_hash(C.gconstpointer(value.native()))
	return uint(ret0)
}

// IsContainer is a wrapper around g_variant_is_container().
func (value Variant) IsContainer() bool {
	ret0 := C.g_variant_is_container(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsFloating is a wrapper around g_variant_is_floating().
func (value Variant) IsFloating() bool {
	ret0 := C.g_variant_is_floating(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsNormalForm is a wrapper around g_variant_is_normal_form().
func (value Variant) IsNormalForm() bool {
	ret0 := C.g_variant_is_normal_form(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsOfType is a wrapper around g_variant_is_of_type().
func (value Variant) IsOfType(type_ VariantType) bool {
	ret0 := C.g_variant_is_of_type(value.native(), type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterNew is a wrapper around g_variant_iter_new().
func (value Variant) IterNew() VariantIter {
	ret0 := C.g_variant_iter_new(value.native())
	return wrapVariantIter(ret0)
}

// LookupValue is a wrapper around g_variant_lookup_value().
func (dictionary Variant) LookupValue(key string, expected_type VariantType) Variant {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_variant_lookup_value(dictionary.native(), key0, expected_type.native())
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return wrapVariant(ret0)
}

// NChildren is a wrapper around g_variant_n_children().
func (value Variant) NChildren() uint {
	ret0 := C.g_variant_n_children(value.native())
	return uint(ret0)
}

// Print is a wrapper around g_variant_print().
func (value Variant) Print(type_annotate bool) string {
	ret0 := C.g_variant_print(value.native(), C.gboolean(util.Bool2Int(type_annotate)) /*go:.util*/)
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Ref is a wrapper around g_variant_ref().
func (value Variant) Ref() Variant {
	ret0 := C.g_variant_ref(value.native())
	return wrapVariant(ret0)
}

// RefSink is a wrapper around g_variant_ref_sink().
func (value Variant) RefSink() Variant {
	ret0 := C.g_variant_ref_sink(value.native())
	return wrapVariant(ret0)
}

// Store is a wrapper around g_variant_store().
func (value Variant) Store(data unsafe.Pointer) {
	C.g_variant_store(value.native(), C.gpointer(data))
}

// TakeRef is a wrapper around g_variant_take_ref().
func (value Variant) TakeRef() Variant {
	ret0 := C.g_variant_take_ref(value.native())
	return wrapVariant(ret0)
}

// Unref is a wrapper around g_variant_unref().
func (value Variant) Unref() {
	C.g_variant_unref(value.native())
}

// VariantIsObjectPath is a wrapper around g_variant_is_object_path().
func VariantIsObjectPath(string string) bool {
	string0 := (*C.gchar)(C.CString(string))
	ret0 := C.g_variant_is_object_path(string0)
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// VariantIsSignature is a wrapper around g_variant_is_signature().
func VariantIsSignature(string string) bool {
	string0 := (*C.gchar)(C.CString(string))
	ret0 := C.g_variant_is_signature(string0)
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// VariantParseErrorPrintContext is a wrapper around g_variant_parse_error_print_context().
func VariantParseErrorPrintContext(error Error, source_str string) string {
	source_str0 := (*C.gchar)(C.CString(source_str))
	ret0 := C.g_variant_parse_error_print_context(error.native(), source_str0)
	C.free(unsafe.Pointer(source_str0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// VariantParseErrorQuark is a wrapper around g_variant_parse_error_quark().
func VariantParseErrorQuark() Quark {
	ret0 := C.g_variant_parse_error_quark()
	return Quark(ret0)
}

// Struct Bytes
type Bytes struct {
	Ptr unsafe.Pointer
}

func (v Bytes) native() *C.GBytes {
	return (*C.GBytes)(v.Ptr)
}
func wrapBytes(p *C.GBytes) Bytes {
	return Bytes{unsafe.Pointer(p)}
}
func WrapBytes(p unsafe.Pointer) Bytes {
	return Bytes{p}
}
func (v Bytes) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBytes(p unsafe.Pointer) interface{} {
	return WrapBytes(p)
}

// GetSize is a wrapper around g_bytes_get_size().
func (bytes Bytes) GetSize() uint {
	ret0 := C.g_bytes_get_size(bytes.native())
	return uint(ret0)
}

// NewFromBytes is a wrapper around g_bytes_new_from_bytes().
func (bytes Bytes) NewFromBytes(offset uint, length uint) Bytes {
	ret0 := C.g_bytes_new_from_bytes(bytes.native(), C.gsize(offset), C.gsize(length))
	return wrapBytes(ret0)
}

// Ref is a wrapper around g_bytes_ref().
func (bytes Bytes) Ref() Bytes {
	ret0 := C.g_bytes_ref(bytes.native())
	return wrapBytes(ret0)
}

// Unref is a wrapper around g_bytes_unref().
func (bytes Bytes) Unref() {
	C.g_bytes_unref(bytes.native())
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
func (v VariantType) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVariantType(p unsafe.Pointer) interface{} {
	return WrapVariantType(p)
}

// VariantTypeNew is a wrapper around g_variant_type_new().
func VariantTypeNew(type_string string) VariantType {
	type_string0 := (*C.gchar)(C.CString(type_string))
	ret0 := C.g_variant_type_new(type_string0)
	C.free(unsafe.Pointer(type_string0)) /*ch:<stdlib.h>*/
	return wrapVariantType(ret0)
}

// VariantTypeNewArray is a wrapper around g_variant_type_new_array().
func VariantTypeNewArray(element VariantType) VariantType {
	ret0 := C.g_variant_type_new_array(element.native())
	return wrapVariantType(ret0)
}

// VariantTypeNewDictEntry is a wrapper around g_variant_type_new_dict_entry().
func VariantTypeNewDictEntry(key VariantType, value VariantType) VariantType {
	ret0 := C.g_variant_type_new_dict_entry(key.native(), value.native())
	return wrapVariantType(ret0)
}

// VariantTypeNewMaybe is a wrapper around g_variant_type_new_maybe().
func VariantTypeNewMaybe(element VariantType) VariantType {
	ret0 := C.g_variant_type_new_maybe(element.native())
	return wrapVariantType(ret0)
}

// VariantTypeNewTuple is a wrapper around g_variant_type_new_tuple().
func VariantTypeNewTuple(items []VariantType) VariantType {
	items0 := make([]*C.GVariantType, len(items))
	for idx, elemG := range items {
		items0[idx] = elemG.native()
	}
	var items0Ptr **C.GVariantType
	if len(items0) > 0 {
		items0Ptr = &items0[0]
	}
	ret0 := C.g_variant_type_new_tuple(items0Ptr, C.gint(len(items)))
	return wrapVariantType(ret0)
}

// Copy is a wrapper around g_variant_type_copy().
func (type_ VariantType) Copy() VariantType {
	ret0 := C.g_variant_type_copy(type_.native())
	return wrapVariantType(ret0)
}

// DupString is a wrapper around g_variant_type_dup_string().
func (type_ VariantType) DupString() string {
	ret0 := C.g_variant_type_dup_string(type_.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Element is a wrapper around g_variant_type_element().
func (type_ VariantType) Element() VariantType {
	ret0 := C.g_variant_type_element(type_.native())
	return wrapVariantType(ret0)
}

// Equal is a wrapper around g_variant_type_equal().
func (type1 VariantType) Equal(type2 VariantType) bool {
	ret0 := C.g_variant_type_equal(C.gconstpointer(type1.native()), C.gconstpointer(type2.native()))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// First is a wrapper around g_variant_type_first().
func (type_ VariantType) First() VariantType {
	ret0 := C.g_variant_type_first(type_.native())
	return wrapVariantType(ret0)
}

// Free is a wrapper around g_variant_type_free().
func (type_ VariantType) Free() {
	C.g_variant_type_free(type_.native())
}

// GetStringLength is a wrapper around g_variant_type_get_string_length().
func (type_ VariantType) GetStringLength() uint {
	ret0 := C.g_variant_type_get_string_length(type_.native())
	return uint(ret0)
}

// Hash is a wrapper around g_variant_type_hash().
func (type_ VariantType) Hash() uint {
	ret0 := C.g_variant_type_hash(C.gconstpointer(type_.native()))
	return uint(ret0)
}

// IsArray is a wrapper around g_variant_type_is_array().
func (type_ VariantType) IsArray() bool {
	ret0 := C.g_variant_type_is_array(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsBasic is a wrapper around g_variant_type_is_basic().
func (type_ VariantType) IsBasic() bool {
	ret0 := C.g_variant_type_is_basic(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsContainer is a wrapper around g_variant_type_is_container().
func (type_ VariantType) IsContainer() bool {
	ret0 := C.g_variant_type_is_container(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsDefinite is a wrapper around g_variant_type_is_definite().
func (type_ VariantType) IsDefinite() bool {
	ret0 := C.g_variant_type_is_definite(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsDictEntry is a wrapper around g_variant_type_is_dict_entry().
func (type_ VariantType) IsDictEntry() bool {
	ret0 := C.g_variant_type_is_dict_entry(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsMaybe is a wrapper around g_variant_type_is_maybe().
func (type_ VariantType) IsMaybe() bool {
	ret0 := C.g_variant_type_is_maybe(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsSubtypeOf is a wrapper around g_variant_type_is_subtype_of().
func (type_ VariantType) IsSubtypeOf(supertype VariantType) bool {
	ret0 := C.g_variant_type_is_subtype_of(type_.native(), supertype.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsTuple is a wrapper around g_variant_type_is_tuple().
func (type_ VariantType) IsTuple() bool {
	ret0 := C.g_variant_type_is_tuple(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsVariant is a wrapper around g_variant_type_is_variant().
func (type_ VariantType) IsVariant() bool {
	ret0 := C.g_variant_type_is_variant(type_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Key is a wrapper around g_variant_type_key().
func (type_ VariantType) Key() VariantType {
	ret0 := C.g_variant_type_key(type_.native())
	return wrapVariantType(ret0)
}

// NItems is a wrapper around g_variant_type_n_items().
func (type_ VariantType) NItems() uint {
	ret0 := C.g_variant_type_n_items(type_.native())
	return uint(ret0)
}

// Next is a wrapper around g_variant_type_next().
func (type_ VariantType) Next() VariantType {
	ret0 := C.g_variant_type_next(type_.native())
	return wrapVariantType(ret0)
}

// PeekString is a wrapper around g_variant_type_peek_string().
func (type_ VariantType) PeekString() string {
	ret0 := C.g_variant_type_peek_string(type_.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Value is a wrapper around g_variant_type_value().
func (type_ VariantType) Value() VariantType {
	ret0 := C.g_variant_type_value(type_.native())
	return wrapVariantType(ret0)
}

// VariantTypeChecked_ is a wrapper around g_variant_type_checked_().
func VariantTypeChecked_(arg0 string) VariantType {
	arg00 := (*C.gchar)(C.CString(arg0))
	ret0 := C.g_variant_type_checked_(arg00)
	C.free(unsafe.Pointer(arg00)) /*ch:<stdlib.h>*/
	return wrapVariantType(ret0)
}

// VariantTypeStringIsValid is a wrapper around g_variant_type_string_is_valid().
func VariantTypeStringIsValid(type_string string) bool {
	type_string0 := (*C.gchar)(C.CString(type_string))
	ret0 := C.g_variant_type_string_is_valid(type_string0)
	C.free(unsafe.Pointer(type_string0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// VariantTypeStringScan is a wrapper around g_variant_type_string_scan().
func VariantTypeStringScan(string string, limit string) (bool, string) {
	string0 := (*C.gchar)(C.CString(string))
	limit0 := (*C.gchar)(C.CString(limit))
	var endptr0 *C.gchar
	ret0 := C.g_variant_type_string_scan(string0, limit0, &endptr0)
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(limit0))  /*ch:<stdlib.h>*/
	endptr := C.GoString((*C.char)(endptr0))
	defer C.g_free(C.gpointer(endptr0))
	return util.Int2Bool(int(ret0)) /*go:.util*/, endptr
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
func (v Error) IsNil() bool {
	return v.Ptr == nil
}
func IWrapError(p unsafe.Pointer) interface{} {
	return WrapError(p)
}

// Free is a wrapper around g_error_free().
func (error Error) Free() {
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
func (v MainLoop) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMainLoop(p unsafe.Pointer) interface{} {
	return WrapMainLoop(p)
}

// MainLoopNew is a wrapper around g_main_loop_new().
func MainLoopNew(context MainContext, is_running bool) MainLoop {
	ret0 := C.g_main_loop_new(context.native(), C.gboolean(util.Bool2Int(is_running)) /*go:.util*/)
	return wrapMainLoop(ret0)
}

// GetContext is a wrapper around g_main_loop_get_context().
func (loop MainLoop) GetContext() MainContext {
	ret0 := C.g_main_loop_get_context(loop.native())
	return wrapMainContext(ret0)
}

// IsRunning is a wrapper around g_main_loop_is_running().
func (loop MainLoop) IsRunning() bool {
	ret0 := C.g_main_loop_is_running(loop.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Quit is a wrapper around g_main_loop_quit().
func (loop MainLoop) Quit() {
	C.g_main_loop_quit(loop.native())
}

// Ref is a wrapper around g_main_loop_ref().
func (loop MainLoop) Ref() MainLoop {
	ret0 := C.g_main_loop_ref(loop.native())
	return wrapMainLoop(ret0)
}

// Run is a wrapper around g_main_loop_run().
func (loop MainLoop) Run() {
	C.g_main_loop_run(loop.native())
}

// Unref is a wrapper around g_main_loop_unref().
func (loop MainLoop) Unref() {
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
func (v MainContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMainContext(p unsafe.Pointer) interface{} {
	return WrapMainContext(p)
}

// MainContextNew is a wrapper around g_main_context_new().
func MainContextNew() MainContext {
	ret0 := C.g_main_context_new()
	return wrapMainContext(ret0)
}

// Acquire is a wrapper around g_main_context_acquire().
func (context MainContext) Acquire() bool {
	ret0 := C.g_main_context_acquire(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// AddPoll is a wrapper around g_main_context_add_poll().
func (context MainContext) AddPoll(fd PollFD, priority int) {
	C.g_main_context_add_poll(context.native(), fd.native(), C.gint(priority))
}

// Dispatch is a wrapper around g_main_context_dispatch().
func (context MainContext) Dispatch() {
	C.g_main_context_dispatch(context.native())
}

// FindSourceById is a wrapper around g_main_context_find_source_by_id().
func (context MainContext) FindSourceById(source_id uint) Source {
	ret0 := C.g_main_context_find_source_by_id(context.native(), C.guint(source_id))
	return wrapSource(ret0)
}

// FindSourceByUserData is a wrapper around g_main_context_find_source_by_user_data().
func (context MainContext) FindSourceByUserData(user_data unsafe.Pointer) Source {
	ret0 := C.g_main_context_find_source_by_user_data(context.native(), C.gpointer(user_data))
	return wrapSource(ret0)
}

// IsOwner is a wrapper around g_main_context_is_owner().
func (context MainContext) IsOwner() bool {
	ret0 := C.g_main_context_is_owner(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Iteration is a wrapper around g_main_context_iteration().
func (context MainContext) Iteration(may_block bool) bool {
	ret0 := C.g_main_context_iteration(context.native(), C.gboolean(util.Bool2Int(may_block)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Pending is a wrapper around g_main_context_pending().
func (context MainContext) Pending() bool {
	ret0 := C.g_main_context_pending(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PopThreadDefault is a wrapper around g_main_context_pop_thread_default().
func (context MainContext) PopThreadDefault() {
	C.g_main_context_pop_thread_default(context.native())
}

// PushThreadDefault is a wrapper around g_main_context_push_thread_default().
func (context MainContext) PushThreadDefault() {
	C.g_main_context_push_thread_default(context.native())
}

// Ref is a wrapper around g_main_context_ref().
func (context MainContext) Ref() MainContext {
	ret0 := C.g_main_context_ref(context.native())
	return wrapMainContext(ret0)
}

// Release is a wrapper around g_main_context_release().
func (context MainContext) Release() {
	C.g_main_context_release(context.native())
}

// RemovePoll is a wrapper around g_main_context_remove_poll().
func (context MainContext) RemovePoll(fd PollFD) {
	C.g_main_context_remove_poll(context.native(), fd.native())
}

// Unref is a wrapper around g_main_context_unref().
func (context MainContext) Unref() {
	C.g_main_context_unref(context.native())
}

// Wakeup is a wrapper around g_main_context_wakeup().
func (context MainContext) Wakeup() {
	C.g_main_context_wakeup(context.native())
}

// MainContextDefault is a wrapper around g_main_context_default().
func MainContextDefault() MainContext {
	ret0 := C.g_main_context_default()
	return wrapMainContext(ret0)
}

// MainContextGetThreadDefault is a wrapper around g_main_context_get_thread_default().
func MainContextGetThreadDefault() MainContext {
	ret0 := C.g_main_context_get_thread_default()
	return wrapMainContext(ret0)
}

// MainContextRefThreadDefault is a wrapper around g_main_context_ref_thread_default().
func MainContextRefThreadDefault() MainContext {
	ret0 := C.g_main_context_ref_thread_default()
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
func (v Source) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSource(p unsafe.Pointer) interface{} {
	return WrapSource(p)
}

// AddChildSource is a wrapper around g_source_add_child_source().
func (source Source) AddChildSource(child_source Source) {
	C.g_source_add_child_source(source.native(), child_source.native())
}

// AddPoll is a wrapper around g_source_add_poll().
func (source Source) AddPoll(fd PollFD) {
	C.g_source_add_poll(source.native(), fd.native())
}

// AddUnixFd is a wrapper around g_source_add_unix_fd().
func (source Source) AddUnixFd(fd int, events IOCondition) unsafe.Pointer {
	ret0 := C.g_source_add_unix_fd(source.native(), C.gint(fd), C.GIOCondition(events))
	return unsafe.Pointer(ret0)
}

// Attach is a wrapper around g_source_attach().
func (source Source) Attach(context MainContext) uint {
	ret0 := C.g_source_attach(source.native(), context.native())
	return uint(ret0)
}

// Destroy is a wrapper around g_source_destroy().
func (source Source) Destroy() {
	C.g_source_destroy(source.native())
}

// GetCanRecurse is a wrapper around g_source_get_can_recurse().
func (source Source) GetCanRecurse() bool {
	ret0 := C.g_source_get_can_recurse(source.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetContext is a wrapper around g_source_get_context().
func (source Source) GetContext() MainContext {
	ret0 := C.g_source_get_context(source.native())
	return wrapMainContext(ret0)
}

// GetId is a wrapper around g_source_get_id().
func (source Source) GetId() uint {
	ret0 := C.g_source_get_id(source.native())
	return uint(ret0)
}

// GetName is a wrapper around g_source_get_name().
func (source Source) GetName() string {
	ret0 := C.g_source_get_name(source.native())
	ret := C.GoString(ret0)
	return ret
}

// GetPriority is a wrapper around g_source_get_priority().
func (source Source) GetPriority() int {
	ret0 := C.g_source_get_priority(source.native())
	return int(ret0)
}

// GetReadyTime is a wrapper around g_source_get_ready_time().
func (source Source) GetReadyTime() int64 {
	ret0 := C.g_source_get_ready_time(source.native())
	return int64(ret0)
}

// GetTime is a wrapper around g_source_get_time().
func (source Source) GetTime() int64 {
	ret0 := C.g_source_get_time(source.native())
	return int64(ret0)
}

// IsDestroyed is a wrapper around g_source_is_destroyed().
func (source Source) IsDestroyed() bool {
	ret0 := C.g_source_is_destroyed(source.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ModifyUnixFd is a wrapper around g_source_modify_unix_fd().
func (source Source) ModifyUnixFd(tag unsafe.Pointer, new_events IOCondition) {
	C.g_source_modify_unix_fd(source.native(), C.gpointer(tag), C.GIOCondition(new_events))
}

// QueryUnixFd is a wrapper around g_source_query_unix_fd().
func (source Source) QueryUnixFd(tag unsafe.Pointer) IOCondition {
	ret0 := C.g_source_query_unix_fd(source.native(), C.gpointer(tag))
	return IOCondition(ret0)
}

// Ref is a wrapper around g_source_ref().
func (source Source) Ref() Source {
	ret0 := C.g_source_ref(source.native())
	return wrapSource(ret0)
}

// RemoveChildSource is a wrapper around g_source_remove_child_source().
func (source Source) RemoveChildSource(child_source Source) {
	C.g_source_remove_child_source(source.native(), child_source.native())
}

// RemovePoll is a wrapper around g_source_remove_poll().
func (source Source) RemovePoll(fd PollFD) {
	C.g_source_remove_poll(source.native(), fd.native())
}

// RemoveUnixFd is a wrapper around g_source_remove_unix_fd().
func (source Source) RemoveUnixFd(tag unsafe.Pointer) {
	C.g_source_remove_unix_fd(source.native(), C.gpointer(tag))
}

// SetCanRecurse is a wrapper around g_source_set_can_recurse().
func (source Source) SetCanRecurse(can_recurse bool) {
	C.g_source_set_can_recurse(source.native(), C.gboolean(util.Bool2Int(can_recurse)) /*go:.util*/)
}

// SetName is a wrapper around g_source_set_name().
func (source Source) SetName(name string) {
	name0 := C.CString(name)
	C.g_source_set_name(source.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetPriority is a wrapper around g_source_set_priority().
func (source Source) SetPriority(priority int) {
	C.g_source_set_priority(source.native(), C.gint(priority))
}

// SetReadyTime is a wrapper around g_source_set_ready_time().
func (source Source) SetReadyTime(ready_time int64) {
	C.g_source_set_ready_time(source.native(), C.gint64(ready_time))
}

// Unref is a wrapper around g_source_unref().
func (source Source) Unref() {
	C.g_source_unref(source.native())
}

// SourceRemove is a wrapper around g_source_remove().
func SourceRemove(tag uint) bool {
	ret0 := C.g_source_remove(C.guint(tag))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SourceRemoveByUserData is a wrapper around g_source_remove_by_user_data().
func SourceRemoveByUserData(user_data unsafe.Pointer) bool {
	ret0 := C.g_source_remove_by_user_data(C.gpointer(user_data))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SourceSetNameById is a wrapper around g_source_set_name_by_id().
func SourceSetNameById(tag uint, name string) {
	name0 := C.CString(name)
	C.g_source_set_name_by_id(C.guint(tag), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// Struct VariantIter
type VariantIter struct {
	Ptr unsafe.Pointer
}

func (v VariantIter) native() *C.GVariantIter {
	return (*C.GVariantIter)(v.Ptr)
}
func wrapVariantIter(p *C.GVariantIter) VariantIter {
	return VariantIter{unsafe.Pointer(p)}
}
func WrapVariantIter(p unsafe.Pointer) VariantIter {
	return VariantIter{p}
}
func (v VariantIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVariantIter(p unsafe.Pointer) interface{} {
	return WrapVariantIter(p)
}

// Copy is a wrapper around g_variant_iter_copy().
func (iter VariantIter) Copy() VariantIter {
	ret0 := C.g_variant_iter_copy(iter.native())
	return wrapVariantIter(ret0)
}

// Free is a wrapper around g_variant_iter_free().
func (iter VariantIter) Free() {
	C.g_variant_iter_free(iter.native())
}

// Init is a wrapper around g_variant_iter_init().
func (iter VariantIter) Init(value Variant) uint {
	ret0 := C.g_variant_iter_init(iter.native(), value.native())
	return uint(ret0)
}

// NChildren is a wrapper around g_variant_iter_n_children().
func (iter VariantIter) NChildren() uint {
	ret0 := C.g_variant_iter_n_children(iter.native())
	return uint(ret0)
}

// NextValue is a wrapper around g_variant_iter_next_value().
func (iter VariantIter) NextValue() Variant {
	ret0 := C.g_variant_iter_next_value(iter.native())
	return wrapVariant(ret0)
}

// Struct PollFD
type PollFD struct {
	Ptr unsafe.Pointer
}

func (v PollFD) native() *C.GPollFD {
	return (*C.GPollFD)(v.Ptr)
}
func wrapPollFD(p *C.GPollFD) PollFD {
	return PollFD{unsafe.Pointer(p)}
}
func WrapPollFD(p unsafe.Pointer) PollFD {
	return PollFD{p}
}
func (v PollFD) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPollFD(p unsafe.Pointer) interface{} {
	return WrapPollFD(p)
}

// Struct TimeVal
type TimeVal struct {
	Ptr unsafe.Pointer
}

func (v TimeVal) native() *C.GTimeVal {
	return (*C.GTimeVal)(v.Ptr)
}
func wrapTimeVal(p *C.GTimeVal) TimeVal {
	return TimeVal{unsafe.Pointer(p)}
}
func WrapTimeVal(p unsafe.Pointer) TimeVal {
	return TimeVal{p}
}
func (v TimeVal) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTimeVal(p unsafe.Pointer) interface{} {
	return WrapTimeVal(p)
}

// Add is a wrapper around g_time_val_add().
func (time_ TimeVal) Add(microseconds int) {
	C.g_time_val_add(time_.native(), C.glong(microseconds))
}

// ToIso8601 is a wrapper around g_time_val_to_iso8601().
func (time_ TimeVal) ToIso8601() string {
	ret0 := C.g_time_val_to_iso8601(time_.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Struct KeyFile
type KeyFile struct {
	Ptr unsafe.Pointer
}

func (v KeyFile) native() *C.GKeyFile {
	return (*C.GKeyFile)(v.Ptr)
}
func wrapKeyFile(p *C.GKeyFile) KeyFile {
	return KeyFile{unsafe.Pointer(p)}
}
func WrapKeyFile(p unsafe.Pointer) KeyFile {
	return KeyFile{p}
}
func (v KeyFile) IsNil() bool {
	return v.Ptr == nil
}
func IWrapKeyFile(p unsafe.Pointer) interface{} {
	return WrapKeyFile(p)
}

// KeyFileNew is a wrapper around g_key_file_new().
func KeyFileNew() KeyFile {
	ret0 := C.g_key_file_new()
	return wrapKeyFile(ret0)
}

// Free is a wrapper around g_key_file_free().
func (key_file KeyFile) Free() {
	C.g_key_file_free(key_file.native())
}

// GetBoolean is a wrapper around g_key_file_get_boolean().
func (key_file KeyFile) GetBoolean(group_name string, key string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_boolean(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetBooleanList is a wrapper around g_key_file_get_boolean_list().
func (key_file KeyFile) GetBooleanList(group_name string, key string) ([]bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var length0 C.gsize
	var err Error
	ret0 := C.g_key_file_get_boolean_list(key_file.native(), group_name0, key0, &length0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	var ret0Slice []C.gboolean
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]bool, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = util.Int2Bool(int(elem)) /*go:.util*/
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GetComment is a wrapper around g_key_file_get_comment().
func (key_file KeyFile) GetComment(group_name string, key string) (string, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_comment(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", err.GoValue()
	}
	return ret, nil
}

// GetDouble is a wrapper around g_key_file_get_double().
func (key_file KeyFile) GetDouble(group_name string, key string) (float64, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_double(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0.0, err.GoValue()
	}
	return float64(ret0), nil
}

// GetDoubleList is a wrapper around g_key_file_get_double_list().
func (key_file KeyFile) GetDoubleList(group_name string, key string) ([]float64, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var length0 C.gsize
	var err Error
	ret0 := C.g_key_file_get_double_list(key_file.native(), group_name0, key0, &length0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	var ret0Slice []C.gdouble
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]float64, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = float64(elem)
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GetGroups is a wrapper around g_key_file_get_groups().
func (key_file KeyFile) GetGroups() ([]string, uint) {
	var length0 C.gsize
	ret0 := C.g_key_file_get_groups(key_file.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret, uint(length0)
}

// GetInt64 is a wrapper around g_key_file_get_int64().
func (key_file KeyFile) GetInt64(group_name string, key string) (int64, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_int64(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int64(ret0), nil
}

// GetInteger is a wrapper around g_key_file_get_integer().
func (key_file KeyFile) GetInteger(group_name string, key string) (int, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_integer(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// GetIntegerList is a wrapper around g_key_file_get_integer_list().
func (key_file KeyFile) GetIntegerList(group_name string, key string) ([]int, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var length0 C.gsize
	var err Error
	ret0 := C.g_key_file_get_integer_list(key_file.native(), group_name0, key0, &length0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	var ret0Slice []C.gint
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]int, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = int(elem)
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GetKeys is a wrapper around g_key_file_get_keys().
func (key_file KeyFile) GetKeys(group_name string) ([]string, uint, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	var length0 C.gsize
	var err Error
	ret0 := C.g_key_file_get_keys(key_file.native(), group_name0, &length0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// GetLocaleString is a wrapper around g_key_file_get_locale_string().
func (key_file KeyFile) GetLocaleString(group_name string, key string, locale string) (string, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	locale0 := (*C.gchar)(C.CString(locale))
	var err Error
	ret0 := C.g_key_file_get_locale_string(key_file.native(), group_name0, key0, locale0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(locale0))     /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", err.GoValue()
	}
	return ret, nil
}

// GetLocaleStringList is a wrapper around g_key_file_get_locale_string_list().
func (key_file KeyFile) GetLocaleStringList(group_name string, key string, locale string) ([]string, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	locale0 := (*C.gchar)(C.CString(locale))
	var length0 C.gsize
	var err Error
	ret0 := C.g_key_file_get_locale_string_list(key_file.native(), group_name0, key0, locale0, &length0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(locale0))     /*ch:<stdlib.h>*/
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GetStartGroup is a wrapper around g_key_file_get_start_group().
func (key_file KeyFile) GetStartGroup() string {
	ret0 := C.g_key_file_get_start_group(key_file.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetString is a wrapper around g_key_file_get_string().
func (key_file KeyFile) GetString(group_name string, key string) (string, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_string(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", err.GoValue()
	}
	return ret, nil
}

// GetStringList is a wrapper around g_key_file_get_string_list().
func (key_file KeyFile) GetStringList(group_name string, key string) ([]string, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var length0 C.gsize
	var err Error
	ret0 := C.g_key_file_get_string_list(key_file.native(), group_name0, key0, &length0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GetUint64 is a wrapper around g_key_file_get_uint64().
func (key_file KeyFile) GetUint64(group_name string, key string) (uint64, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_uint64(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint64(ret0), nil
}

// GetValue is a wrapper around g_key_file_get_value().
func (key_file KeyFile) GetValue(group_name string, key string) (string, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_get_value(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", err.GoValue()
	}
	return ret, nil
}

// HasGroup is a wrapper around g_key_file_has_group().
func (key_file KeyFile) HasGroup(group_name string) bool {
	group_name0 := (*C.gchar)(C.CString(group_name))
	ret0 := C.g_key_file_has_group(key_file.native(), group_name0)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// HasKey is a wrapper around g_key_file_has_key().
func (key_file KeyFile) HasKey(group_name string, key string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_has_key(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadFromBytes is a wrapper around g_key_file_load_from_bytes().
func (key_file KeyFile) LoadFromBytes(bytes Bytes, flags KeyFileFlags) (bool, error) {
	var err Error
	ret0 := C.g_key_file_load_from_bytes(key_file.native(), bytes.native(), C.GKeyFileFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadFromData is a wrapper around g_key_file_load_from_data().
func (key_file KeyFile) LoadFromData(data string, length uint, flags KeyFileFlags) (bool, error) {
	data0 := (*C.gchar)(C.CString(data))
	var err Error
	ret0 := C.g_key_file_load_from_data(key_file.native(), data0, C.gsize(length), C.GKeyFileFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(data0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadFromDataDirs is a wrapper around g_key_file_load_from_data_dirs().
func (key_file KeyFile) LoadFromDataDirs(file string, flags KeyFileFlags) (bool, string, error) {
	file0 := (*C.gchar)(C.CString(file))
	var full_path0 *C.gchar
	var err Error
	ret0 := C.g_key_file_load_from_data_dirs(key_file.native(), file0, &full_path0, C.GKeyFileFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file0)) /*ch:<stdlib.h>*/
	full_path := C.GoString((*C.char)(full_path0))
	defer C.g_free(C.gpointer(full_path0))
	if err.Ptr != nil {
		defer err.Free()
		return false, "", err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, full_path, nil
}

// LoadFromDirs is a wrapper around g_key_file_load_from_dirs().
func (key_file KeyFile) LoadFromDirs(file string, search_dirs []string, flags KeyFileFlags) (bool, string, error) {
	file0 := (*C.gchar)(C.CString(file))
	search_dirs0 := make([]*C.gchar, len(search_dirs))
	for idx, elemG := range search_dirs {
		elem := (*C.gchar)(C.CString(elemG))
		search_dirs0[idx] = elem
	}
	var search_dirs0Ptr **C.gchar
	if len(search_dirs0) > 0 {
		search_dirs0Ptr = &search_dirs0[0]
	}
	var full_path0 *C.gchar
	var err Error
	ret0 := C.g_key_file_load_from_dirs(key_file.native(), file0, search_dirs0Ptr, &full_path0, C.GKeyFileFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file0)) /*ch:<stdlib.h>*/
	for _, elem := range search_dirs0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	full_path := C.GoString((*C.char)(full_path0))
	defer C.g_free(C.gpointer(full_path0))
	if err.Ptr != nil {
		defer err.Free()
		return false, "", err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, full_path, nil
}

// LoadFromFile is a wrapper around g_key_file_load_from_file().
func (key_file KeyFile) LoadFromFile(file string, flags KeyFileFlags) (bool, error) {
	file0 := (*C.gchar)(C.CString(file))
	var err Error
	ret0 := C.g_key_file_load_from_file(key_file.native(), file0, C.GKeyFileFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Ref is a wrapper around g_key_file_ref().
func (key_file KeyFile) Ref() KeyFile {
	ret0 := C.g_key_file_ref(key_file.native())
	return wrapKeyFile(ret0)
}

// RemoveComment is a wrapper around g_key_file_remove_comment().
func (key_file KeyFile) RemoveComment(group_name string, key string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_remove_comment(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// RemoveGroup is a wrapper around g_key_file_remove_group().
func (key_file KeyFile) RemoveGroup(group_name string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	var err Error
	ret0 := C.g_key_file_remove_group(key_file.native(), group_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// RemoveKey is a wrapper around g_key_file_remove_key().
func (key_file KeyFile) RemoveKey(group_name string, key string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	var err Error
	ret0 := C.g_key_file_remove_key(key_file.native(), group_name0, key0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SaveToFile is a wrapper around g_key_file_save_to_file().
func (key_file KeyFile) SaveToFile(filename string) (bool, error) {
	filename0 := (*C.gchar)(C.CString(filename))
	var err Error
	ret0 := C.g_key_file_save_to_file(key_file.native(), filename0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetBoolean is a wrapper around g_key_file_set_boolean().
func (key_file KeyFile) SetBoolean(group_name string, key string, value bool) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	C.g_key_file_set_boolean(key_file.native(), group_name0, key0, C.gboolean(util.Bool2Int(value)) /*go:.util*/)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
}

// SetComment is a wrapper around g_key_file_set_comment().
func (key_file KeyFile) SetComment(group_name string, key string, comment string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	comment0 := (*C.gchar)(C.CString(comment))
	var err Error
	ret0 := C.g_key_file_set_comment(key_file.native(), group_name0, key0, comment0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(comment0))    /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetDouble is a wrapper around g_key_file_set_double().
func (key_file KeyFile) SetDouble(group_name string, key string, value float64) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	C.g_key_file_set_double(key_file.native(), group_name0, key0, C.gdouble(value))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
}

// SetInt64 is a wrapper around g_key_file_set_int64().
func (key_file KeyFile) SetInt64(group_name string, key string, value int64) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	C.g_key_file_set_int64(key_file.native(), group_name0, key0, C.gint64(value))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
}

// SetInteger is a wrapper around g_key_file_set_integer().
func (key_file KeyFile) SetInteger(group_name string, key string, value int) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	C.g_key_file_set_integer(key_file.native(), group_name0, key0, C.gint(value))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
}

// SetListSeparator is a wrapper around g_key_file_set_list_separator().
func (key_file KeyFile) SetListSeparator(separator int8) {
	C.g_key_file_set_list_separator(key_file.native(), C.gchar(separator))
}

// SetLocaleString is a wrapper around g_key_file_set_locale_string().
func (key_file KeyFile) SetLocaleString(group_name string, key string, locale string, string string) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	locale0 := (*C.gchar)(C.CString(locale))
	string0 := (*C.gchar)(C.CString(string))
	C.g_key_file_set_locale_string(key_file.native(), group_name0, key0, locale0, string0)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(locale0))     /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(string0))     /*ch:<stdlib.h>*/
}

// SetString is a wrapper around g_key_file_set_string().
func (key_file KeyFile) SetString(group_name string, key string, string string) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	string0 := (*C.gchar)(C.CString(string))
	C.g_key_file_set_string(key_file.native(), group_name0, key0, string0)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(string0))     /*ch:<stdlib.h>*/
}

// SetUint64 is a wrapper around g_key_file_set_uint64().
func (key_file KeyFile) SetUint64(group_name string, key string, value uint64) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	C.g_key_file_set_uint64(key_file.native(), group_name0, key0, C.guint64(value))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
}

// SetValue is a wrapper around g_key_file_set_value().
func (key_file KeyFile) SetValue(group_name string, key string, value string) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	key0 := (*C.gchar)(C.CString(key))
	value0 := (*C.gchar)(C.CString(value))
	C.g_key_file_set_value(key_file.native(), group_name0, key0, value0)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))      /*ch:<stdlib.h>*/
}

// ToData is a wrapper around g_key_file_to_data().
func (key_file KeyFile) ToData() (string, uint, error) {
	var length0 C.gsize
	var err Error
	ret0 := C.g_key_file_to_data(key_file.native(), &length0, (**C.GError)(unsafe.Pointer(&err)))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// Unref is a wrapper around g_key_file_unref().
func (key_file KeyFile) Unref() {
	C.g_key_file_unref(key_file.native())
}

// KeyFileErrorQuark is a wrapper around g_key_file_error_quark().
func KeyFileErrorQuark() Quark {
	ret0 := C.g_key_file_error_quark()
	return Quark(ret0)
}

// Struct TimeZone
type TimeZone struct {
	Ptr unsafe.Pointer
}

func (v TimeZone) native() *C.GTimeZone {
	return (*C.GTimeZone)(v.Ptr)
}
func wrapTimeZone(p *C.GTimeZone) TimeZone {
	return TimeZone{unsafe.Pointer(p)}
}
func WrapTimeZone(p unsafe.Pointer) TimeZone {
	return TimeZone{p}
}
func (v TimeZone) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTimeZone(p unsafe.Pointer) interface{} {
	return WrapTimeZone(p)
}

// TimeZoneNew is a wrapper around g_time_zone_new().
func TimeZoneNew(identifier string) TimeZone {
	identifier0 := (*C.gchar)(C.CString(identifier))
	ret0 := C.g_time_zone_new(identifier0)
	C.free(unsafe.Pointer(identifier0)) /*ch:<stdlib.h>*/
	return wrapTimeZone(ret0)
}

// TimeZoneNewLocal is a wrapper around g_time_zone_new_local().
func TimeZoneNewLocal() TimeZone {
	ret0 := C.g_time_zone_new_local()
	return wrapTimeZone(ret0)
}

// TimeZoneNewUtc is a wrapper around g_time_zone_new_utc().
func TimeZoneNewUtc() TimeZone {
	ret0 := C.g_time_zone_new_utc()
	return wrapTimeZone(ret0)
}

// FindInterval is a wrapper around g_time_zone_find_interval().
func (tz TimeZone) FindInterval(type_ TimeType, time_ int64) int {
	ret0 := C.g_time_zone_find_interval(tz.native(), C.GTimeType(type_), C.gint64(time_))
	return int(ret0)
}

// GetAbbreviation is a wrapper around g_time_zone_get_abbreviation().
func (tz TimeZone) GetAbbreviation(interval int) string {
	ret0 := C.g_time_zone_get_abbreviation(tz.native(), C.gint(interval))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetOffset is a wrapper around g_time_zone_get_offset().
func (tz TimeZone) GetOffset(interval int) int32 {
	ret0 := C.g_time_zone_get_offset(tz.native(), C.gint(interval))
	return int32(ret0)
}

// IsDst is a wrapper around g_time_zone_is_dst().
func (tz TimeZone) IsDst(interval int) bool {
	ret0 := C.g_time_zone_is_dst(tz.native(), C.gint(interval))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_time_zone_ref().
func (tz TimeZone) Ref() TimeZone {
	ret0 := C.g_time_zone_ref(tz.native())
	return wrapTimeZone(ret0)
}

// Unref is a wrapper around g_time_zone_unref().
func (tz TimeZone) Unref() {
	C.g_time_zone_unref(tz.native())
}

// Struct DateTime
type DateTime struct {
	Ptr unsafe.Pointer
}

func (v DateTime) native() *C.GDateTime {
	return (*C.GDateTime)(v.Ptr)
}
func wrapDateTime(p *C.GDateTime) DateTime {
	return DateTime{unsafe.Pointer(p)}
}
func WrapDateTime(p unsafe.Pointer) DateTime {
	return DateTime{p}
}
func (v DateTime) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDateTime(p unsafe.Pointer) interface{} {
	return WrapDateTime(p)
}

// DateTimeNew is a wrapper around g_date_time_new().
func DateTimeNew(tz TimeZone, year int, month int, day int, hour int, minute int, seconds float64) DateTime {
	ret0 := C.g_date_time_new(tz.native(), C.gint(year), C.gint(month), C.gint(day), C.gint(hour), C.gint(minute), C.gdouble(seconds))
	return wrapDateTime(ret0)
}

// DateTimeNewFromTimevalLocal is a wrapper around g_date_time_new_from_timeval_local().
func DateTimeNewFromTimevalLocal(tv TimeVal) DateTime {
	ret0 := C.g_date_time_new_from_timeval_local(tv.native())
	return wrapDateTime(ret0)
}

// DateTimeNewFromTimevalUtc is a wrapper around g_date_time_new_from_timeval_utc().
func DateTimeNewFromTimevalUtc(tv TimeVal) DateTime {
	ret0 := C.g_date_time_new_from_timeval_utc(tv.native())
	return wrapDateTime(ret0)
}

// DateTimeNewFromUnixLocal is a wrapper around g_date_time_new_from_unix_local().
func DateTimeNewFromUnixLocal(t int64) DateTime {
	ret0 := C.g_date_time_new_from_unix_local(C.gint64(t))
	return wrapDateTime(ret0)
}

// DateTimeNewFromUnixUtc is a wrapper around g_date_time_new_from_unix_utc().
func DateTimeNewFromUnixUtc(t int64) DateTime {
	ret0 := C.g_date_time_new_from_unix_utc(C.gint64(t))
	return wrapDateTime(ret0)
}

// DateTimeNewLocal is a wrapper around g_date_time_new_local().
func DateTimeNewLocal(year int, month int, day int, hour int, minute int, seconds float64) DateTime {
	ret0 := C.g_date_time_new_local(C.gint(year), C.gint(month), C.gint(day), C.gint(hour), C.gint(minute), C.gdouble(seconds))
	return wrapDateTime(ret0)
}

// DateTimeNewNow is a wrapper around g_date_time_new_now().
func DateTimeNewNow(tz TimeZone) DateTime {
	ret0 := C.g_date_time_new_now(tz.native())
	return wrapDateTime(ret0)
}

// DateTimeNewNowLocal is a wrapper around g_date_time_new_now_local().
func DateTimeNewNowLocal() DateTime {
	ret0 := C.g_date_time_new_now_local()
	return wrapDateTime(ret0)
}

// DateTimeNewNowUtc is a wrapper around g_date_time_new_now_utc().
func DateTimeNewNowUtc() DateTime {
	ret0 := C.g_date_time_new_now_utc()
	return wrapDateTime(ret0)
}

// DateTimeNewUtc is a wrapper around g_date_time_new_utc().
func DateTimeNewUtc(year int, month int, day int, hour int, minute int, seconds float64) DateTime {
	ret0 := C.g_date_time_new_utc(C.gint(year), C.gint(month), C.gint(day), C.gint(hour), C.gint(minute), C.gdouble(seconds))
	return wrapDateTime(ret0)
}

// AddDays is a wrapper around g_date_time_add_days().
func (datetime DateTime) AddDays(days int) DateTime {
	ret0 := C.g_date_time_add_days(datetime.native(), C.gint(days))
	return wrapDateTime(ret0)
}

// AddFull is a wrapper around g_date_time_add_full().
func (datetime DateTime) AddFull(years int, months int, days int, hours int, minutes int, seconds float64) DateTime {
	ret0 := C.g_date_time_add_full(datetime.native(), C.gint(years), C.gint(months), C.gint(days), C.gint(hours), C.gint(minutes), C.gdouble(seconds))
	return wrapDateTime(ret0)
}

// AddHours is a wrapper around g_date_time_add_hours().
func (datetime DateTime) AddHours(hours int) DateTime {
	ret0 := C.g_date_time_add_hours(datetime.native(), C.gint(hours))
	return wrapDateTime(ret0)
}

// AddMinutes is a wrapper around g_date_time_add_minutes().
func (datetime DateTime) AddMinutes(minutes int) DateTime {
	ret0 := C.g_date_time_add_minutes(datetime.native(), C.gint(minutes))
	return wrapDateTime(ret0)
}

// AddMonths is a wrapper around g_date_time_add_months().
func (datetime DateTime) AddMonths(months int) DateTime {
	ret0 := C.g_date_time_add_months(datetime.native(), C.gint(months))
	return wrapDateTime(ret0)
}

// AddSeconds is a wrapper around g_date_time_add_seconds().
func (datetime DateTime) AddSeconds(seconds float64) DateTime {
	ret0 := C.g_date_time_add_seconds(datetime.native(), C.gdouble(seconds))
	return wrapDateTime(ret0)
}

// AddWeeks is a wrapper around g_date_time_add_weeks().
func (datetime DateTime) AddWeeks(weeks int) DateTime {
	ret0 := C.g_date_time_add_weeks(datetime.native(), C.gint(weeks))
	return wrapDateTime(ret0)
}

// AddYears is a wrapper around g_date_time_add_years().
func (datetime DateTime) AddYears(years int) DateTime {
	ret0 := C.g_date_time_add_years(datetime.native(), C.gint(years))
	return wrapDateTime(ret0)
}

// Format is a wrapper around g_date_time_format().
func (datetime DateTime) Format(format string) string {
	format0 := (*C.gchar)(C.CString(format))
	ret0 := C.g_date_time_format(datetime.native(), format0)
	C.free(unsafe.Pointer(format0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetDayOfMonth is a wrapper around g_date_time_get_day_of_month().
func (datetime DateTime) GetDayOfMonth() int {
	ret0 := C.g_date_time_get_day_of_month(datetime.native())
	return int(ret0)
}

// GetDayOfWeek is a wrapper around g_date_time_get_day_of_week().
func (datetime DateTime) GetDayOfWeek() int {
	ret0 := C.g_date_time_get_day_of_week(datetime.native())
	return int(ret0)
}

// GetDayOfYear is a wrapper around g_date_time_get_day_of_year().
func (datetime DateTime) GetDayOfYear() int {
	ret0 := C.g_date_time_get_day_of_year(datetime.native())
	return int(ret0)
}

// GetHour is a wrapper around g_date_time_get_hour().
func (datetime DateTime) GetHour() int {
	ret0 := C.g_date_time_get_hour(datetime.native())
	return int(ret0)
}

// GetMicrosecond is a wrapper around g_date_time_get_microsecond().
func (datetime DateTime) GetMicrosecond() int {
	ret0 := C.g_date_time_get_microsecond(datetime.native())
	return int(ret0)
}

// GetMinute is a wrapper around g_date_time_get_minute().
func (datetime DateTime) GetMinute() int {
	ret0 := C.g_date_time_get_minute(datetime.native())
	return int(ret0)
}

// GetMonth is a wrapper around g_date_time_get_month().
func (datetime DateTime) GetMonth() int {
	ret0 := C.g_date_time_get_month(datetime.native())
	return int(ret0)
}

// GetSecond is a wrapper around g_date_time_get_second().
func (datetime DateTime) GetSecond() int {
	ret0 := C.g_date_time_get_second(datetime.native())
	return int(ret0)
}

// GetSeconds is a wrapper around g_date_time_get_seconds().
func (datetime DateTime) GetSeconds() float64 {
	ret0 := C.g_date_time_get_seconds(datetime.native())
	return float64(ret0)
}

// GetTimezoneAbbreviation is a wrapper around g_date_time_get_timezone_abbreviation().
func (datetime DateTime) GetTimezoneAbbreviation() string {
	ret0 := C.g_date_time_get_timezone_abbreviation(datetime.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWeekNumberingYear is a wrapper around g_date_time_get_week_numbering_year().
func (datetime DateTime) GetWeekNumberingYear() int {
	ret0 := C.g_date_time_get_week_numbering_year(datetime.native())
	return int(ret0)
}

// GetWeekOfYear is a wrapper around g_date_time_get_week_of_year().
func (datetime DateTime) GetWeekOfYear() int {
	ret0 := C.g_date_time_get_week_of_year(datetime.native())
	return int(ret0)
}

// GetYear is a wrapper around g_date_time_get_year().
func (datetime DateTime) GetYear() int {
	ret0 := C.g_date_time_get_year(datetime.native())
	return int(ret0)
}

// GetYmd is a wrapper around g_date_time_get_ymd().
func (datetime DateTime) GetYmd() (int, int, int) {
	var year0 C.gint
	var month0 C.gint
	var day0 C.gint
	C.g_date_time_get_ymd(datetime.native(), &year0, &month0, &day0)
	return int(year0), int(month0), int(day0)
}

// IsDaylightSavings is a wrapper around g_date_time_is_daylight_savings().
func (datetime DateTime) IsDaylightSavings() bool {
	ret0 := C.g_date_time_is_daylight_savings(datetime.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_date_time_ref().
func (datetime DateTime) Ref() DateTime {
	ret0 := C.g_date_time_ref(datetime.native())
	return wrapDateTime(ret0)
}

// ToLocal is a wrapper around g_date_time_to_local().
func (datetime DateTime) ToLocal() DateTime {
	ret0 := C.g_date_time_to_local(datetime.native())
	return wrapDateTime(ret0)
}

// ToTimeval is a wrapper around g_date_time_to_timeval().
func (datetime DateTime) ToTimeval(tv TimeVal) bool {
	ret0 := C.g_date_time_to_timeval(datetime.native(), tv.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ToTimezone is a wrapper around g_date_time_to_timezone().
func (datetime DateTime) ToTimezone(tz TimeZone) DateTime {
	ret0 := C.g_date_time_to_timezone(datetime.native(), tz.native())
	return wrapDateTime(ret0)
}

// ToUnix is a wrapper around g_date_time_to_unix().
func (datetime DateTime) ToUnix() int64 {
	ret0 := C.g_date_time_to_unix(datetime.native())
	return int64(ret0)
}

// ToUtc is a wrapper around g_date_time_to_utc().
func (datetime DateTime) ToUtc() DateTime {
	ret0 := C.g_date_time_to_utc(datetime.native())
	return wrapDateTime(ret0)
}

// Unref is a wrapper around g_date_time_unref().
func (datetime DateTime) Unref() {
	C.g_date_time_unref(datetime.native())
}

// DateTimeCompare is a wrapper around g_date_time_compare().
func DateTimeCompare(dt1 unsafe.Pointer, dt2 unsafe.Pointer) int {
	ret0 := C.g_date_time_compare(C.gconstpointer(dt1), C.gconstpointer(dt2))
	return int(ret0)
}

// DateTimeEqual is a wrapper around g_date_time_equal().
func DateTimeEqual(dt1 unsafe.Pointer, dt2 unsafe.Pointer) bool {
	ret0 := C.g_date_time_equal(C.gconstpointer(dt1), C.gconstpointer(dt2))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// DateTimeHash is a wrapper around g_date_time_hash().
func DateTimeHash(datetime unsafe.Pointer) uint {
	ret0 := C.g_date_time_hash(C.gconstpointer(datetime))
	return uint(ret0)
}

// Struct OptionGroup
type OptionGroup struct {
	Ptr unsafe.Pointer
}

func (v OptionGroup) native() *C.GOptionGroup {
	return (*C.GOptionGroup)(v.Ptr)
}
func wrapOptionGroup(p *C.GOptionGroup) OptionGroup {
	return OptionGroup{unsafe.Pointer(p)}
}
func WrapOptionGroup(p unsafe.Pointer) OptionGroup {
	return OptionGroup{p}
}
func (v OptionGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOptionGroup(p unsafe.Pointer) interface{} {
	return WrapOptionGroup(p)
}

// AddEntries is a wrapper around g_option_group_add_entries().
func (group OptionGroup) AddEntries(entries OptionEntry) {
	C.g_option_group_add_entries(group.native(), entries.native())
}

// Ref is a wrapper around g_option_group_ref().
func (group OptionGroup) Ref() OptionGroup {
	ret0 := C.g_option_group_ref(group.native())
	return wrapOptionGroup(ret0)
}

// SetTranslationDomain is a wrapper around g_option_group_set_translation_domain().
func (group OptionGroup) SetTranslationDomain(domain string) {
	domain0 := (*C.gchar)(C.CString(domain))
	C.g_option_group_set_translation_domain(group.native(), domain0)
	C.free(unsafe.Pointer(domain0)) /*ch:<stdlib.h>*/
}

// Unref is a wrapper around g_option_group_unref().
func (group OptionGroup) Unref() {
	C.g_option_group_unref(group.native())
}

// Struct OptionEntry
type OptionEntry struct {
	Ptr unsafe.Pointer
}

func (v OptionEntry) native() *C.GOptionEntry {
	return (*C.GOptionEntry)(v.Ptr)
}
func wrapOptionEntry(p *C.GOptionEntry) OptionEntry {
	return OptionEntry{unsafe.Pointer(p)}
}
func WrapOptionEntry(p unsafe.Pointer) OptionEntry {
	return OptionEntry{p}
}
func (v OptionEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOptionEntry(p unsafe.Pointer) interface{} {
	return WrapOptionEntry(p)
}

// Struct VariantBuilder
type VariantBuilder struct {
	Ptr unsafe.Pointer
}

func (v VariantBuilder) native() *C.GVariantBuilder {
	return (*C.GVariantBuilder)(v.Ptr)
}
func wrapVariantBuilder(p *C.GVariantBuilder) VariantBuilder {
	return VariantBuilder{unsafe.Pointer(p)}
}
func WrapVariantBuilder(p unsafe.Pointer) VariantBuilder {
	return VariantBuilder{p}
}
func (v VariantBuilder) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVariantBuilder(p unsafe.Pointer) interface{} {
	return WrapVariantBuilder(p)
}

// VariantBuilderNew is a wrapper around g_variant_builder_new().
func VariantBuilderNew(type_ VariantType) VariantBuilder {
	ret0 := C.g_variant_builder_new(type_.native())
	return wrapVariantBuilder(ret0)
}

// AddValue is a wrapper around g_variant_builder_add_value().
func (builder VariantBuilder) AddValue(value Variant) {
	C.g_variant_builder_add_value(builder.native(), value.native())
}

// Clear is a wrapper around g_variant_builder_clear().
func (builder VariantBuilder) Clear() {
	C.g_variant_builder_clear(builder.native())
}

// Close is a wrapper around g_variant_builder_close().
func (builder VariantBuilder) Close() {
	C.g_variant_builder_close(builder.native())
}

// End is a wrapper around g_variant_builder_end().
func (builder VariantBuilder) End() Variant {
	ret0 := C.g_variant_builder_end(builder.native())
	return wrapVariant(ret0)
}

// Init is a wrapper around g_variant_builder_init().
func (builder VariantBuilder) Init(type_ VariantType) {
	C.g_variant_builder_init(builder.native(), type_.native())
}

// Open is a wrapper around g_variant_builder_open().
func (builder VariantBuilder) Open(type_ VariantType) {
	C.g_variant_builder_open(builder.native(), type_.native())
}

// Ref is a wrapper around g_variant_builder_ref().
func (builder VariantBuilder) Ref() VariantBuilder {
	ret0 := C.g_variant_builder_ref(builder.native())
	return wrapVariantBuilder(ret0)
}

// Unref is a wrapper around g_variant_builder_unref().
func (builder VariantBuilder) Unref() {
	C.g_variant_builder_unref(builder.native())
}

// IdleSourceNew is a wrapper around g_idle_source_new().
func IdleSourceNew() Source {
	ret0 := C.g_idle_source_new()
	return wrapSource(ret0)
}

// TimeoutSourceNew is a wrapper around g_timeout_source_new().
func TimeoutSourceNew(interval uint) Source {
	ret0 := C.g_timeout_source_new(C.guint(interval))
	return wrapSource(ret0)
}

// TimeoutSourceNewSeconds is a wrapper around g_timeout_source_new_seconds().
func TimeoutSourceNewSeconds(interval uint) Source {
	ret0 := C.g_timeout_source_new_seconds(C.guint(interval))
	return wrapSource(ret0)
}

type DateDay uint8
type DateYear uint16
type MutexLocker unsafe.Pointer
type Pid int
type Quark uint32
type Strv unsafe.Pointer
type Time int32
type TimeSpan int64
type Type uint
type BookmarkFileError int

const (
	BookmarkFileErrorInvalidUri       BookmarkFileError = 0
	BookmarkFileErrorInvalidValue                       = 1
	BookmarkFileErrorAppNotRegistered                   = 2
	BookmarkFileErrorUriNotFound                        = 3
	BookmarkFileErrorRead                               = 4
	BookmarkFileErrorUnknownEncoding                    = 5
	BookmarkFileErrorWrite                              = 6
	BookmarkFileErrorFileNotFound                       = 7
)

type ChecksumType int

const (
	ChecksumTypeMd5    ChecksumType = 0
	ChecksumTypeSha1                = 1
	ChecksumTypeSha256              = 2
	ChecksumTypeSha512              = 3
)

type ConvertError int

const (
	ConvertErrorNoConversion    ConvertError = 0
	ConvertErrorIllegalSequence              = 1
	ConvertErrorFailed                       = 2
	ConvertErrorPartialInput                 = 3
	ConvertErrorBadUri                       = 4
	ConvertErrorNotAbsolutePath              = 5
	ConvertErrorNoMemory                     = 6
)

type DateDMY int

const (
	DateDMYDay   DateDMY = 0
	DateDMYMonth         = 1
	DateDMYYear          = 2
)

type DateMonth int

const (
	DateMonthBadMonth  DateMonth = 0
	DateMonthJanuary             = 1
	DateMonthFebruary            = 2
	DateMonthMarch               = 3
	DateMonthApril               = 4
	DateMonthMay                 = 5
	DateMonthJune                = 6
	DateMonthJuly                = 7
	DateMonthAugust              = 8
	DateMonthSeptember           = 9
	DateMonthOctober             = 10
	DateMonthNovember            = 11
	DateMonthDecember            = 12
)

type DateWeekday int

const (
	DateWeekdayBadWeekday DateWeekday = 0
	DateWeekdayMonday                 = 1
	DateWeekdayTuesday                = 2
	DateWeekdayWednesday              = 3
	DateWeekdayThursday               = 4
	DateWeekdayFriday                 = 5
	DateWeekdaySaturday               = 6
	DateWeekdaySunday                 = 7
)

type ErrorType int

const (
	ErrorTypeUnknown           ErrorType = 0
	ErrorTypeUnexpEof                    = 1
	ErrorTypeUnexpEofInString            = 2
	ErrorTypeUnexpEofInComment           = 3
	ErrorTypeNonDigitInConst             = 4
	ErrorTypeDigitRadix                  = 5
	ErrorTypeFloatRadix                  = 6
	ErrorTypeFloatMalformed              = 7
)

type FileError int

const (
	FileErrorExist       FileError = 0
	FileErrorIsdir                 = 1
	FileErrorAcces                 = 2
	FileErrorNametoolong           = 3
	FileErrorNoent                 = 4
	FileErrorNotdir                = 5
	FileErrorNxio                  = 6
	FileErrorNodev                 = 7
	FileErrorRofs                  = 8
	FileErrorTxtbsy                = 9
	FileErrorFault                 = 10
	FileErrorLoop                  = 11
	FileErrorNospc                 = 12
	FileErrorNomem                 = 13
	FileErrorMfile                 = 14
	FileErrorNfile                 = 15
	FileErrorBadf                  = 16
	FileErrorInval                 = 17
	FileErrorPipe                  = 18
	FileErrorAgain                 = 19
	FileErrorIntr                  = 20
	FileErrorIo                    = 21
	FileErrorPerm                  = 22
	FileErrorNosys                 = 23
	FileErrorFailed                = 24
)

type IOChannelError int

const (
	IOChannelErrorFbig     IOChannelError = 0
	IOChannelErrorInval                   = 1
	IOChannelErrorIo                      = 2
	IOChannelErrorIsdir                   = 3
	IOChannelErrorNospc                   = 4
	IOChannelErrorNxio                    = 5
	IOChannelErrorOverflow                = 6
	IOChannelErrorPipe                    = 7
	IOChannelErrorFailed                  = 8
)

type IOError int

const (
	IOErrorNone    IOError = 0
	IOErrorAgain           = 1
	IOErrorInval           = 2
	IOErrorUnknown         = 3
)

type IOStatus int

const (
	IOStatusError  IOStatus = 0
	IOStatusNormal          = 1
	IOStatusEof             = 2
	IOStatusAgain           = 3
)

type KeyFileError int

const (
	KeyFileErrorUnknownEncoding KeyFileError = 0
	KeyFileErrorParse                        = 1
	KeyFileErrorNotFound                     = 2
	KeyFileErrorKeyNotFound                  = 3
	KeyFileErrorGroupNotFound                = 4
	KeyFileErrorInvalidValue                 = 5
)

type LogWriterOutput int

const (
	LogWriterOutputHandled   LogWriterOutput = 1
	LogWriterOutputUnhandled                 = 0
)

type MarkupError int

const (
	MarkupErrorBadUtf8          MarkupError = 0
	MarkupErrorEmpty                        = 1
	MarkupErrorParse                        = 2
	MarkupErrorUnknownElement               = 3
	MarkupErrorUnknownAttribute             = 4
	MarkupErrorInvalidContent               = 5
	MarkupErrorMissingAttribute             = 6
)

type NormalizeMode int

const (
	NormalizeModeDefault        NormalizeMode = 0
	NormalizeModeNfd                          = 0
	NormalizeModeDefaultCompose               = 1
	NormalizeModeNfc                          = 1
	NormalizeModeAll                          = 2
	NormalizeModeNfkd                         = 2
	NormalizeModeAllCompose                   = 3
	NormalizeModeNfkc                         = 3
)

type OnceStatus int

const (
	OnceStatusNotcalled OnceStatus = 0
	OnceStatusProgress             = 1
	OnceStatusReady                = 2
)

type OptionArg int

const (
	OptionArgNone          OptionArg = 0
	OptionArgString                  = 1
	OptionArgInt                     = 2
	OptionArgCallback                = 3
	OptionArgFilename                = 4
	OptionArgStringArray             = 5
	OptionArgFilenameArray           = 6
	OptionArgDouble                  = 7
	OptionArgInt64                   = 8
)

type OptionError int

const (
	OptionErrorUnknownOption OptionError = 0
	OptionErrorBadValue                  = 1
	OptionErrorFailed                    = 2
)

type RegexError int

const (
	RegexErrorCompile                                  RegexError = 0
	RegexErrorOptimize                                            = 1
	RegexErrorReplace                                             = 2
	RegexErrorMatch                                               = 3
	RegexErrorInternal                                            = 4
	RegexErrorStrayBackslash                                      = 101
	RegexErrorMissingControlChar                                  = 102
	RegexErrorUnrecognizedEscape                                  = 103
	RegexErrorQuantifiersOutOfOrder                               = 104
	RegexErrorQuantifierTooBig                                    = 105
	RegexErrorUnterminatedCharacterClass                          = 106
	RegexErrorInvalidEscapeInCharacterClass                       = 107
	RegexErrorRangeOutOfOrder                                     = 108
	RegexErrorNothingToRepeat                                     = 109
	RegexErrorUnrecognizedCharacter                               = 112
	RegexErrorPosixNamedClassOutsideClass                         = 113
	RegexErrorUnmatchedParenthesis                                = 114
	RegexErrorInexistentSubpatternReference                       = 115
	RegexErrorUnterminatedComment                                 = 118
	RegexErrorExpressionTooLarge                                  = 120
	RegexErrorMemoryError                                         = 121
	RegexErrorVariableLengthLookbehind                            = 125
	RegexErrorMalformedCondition                                  = 126
	RegexErrorTooManyConditionalBranches                          = 127
	RegexErrorAssertionExpected                                   = 128
	RegexErrorUnknownPosixClassName                               = 130
	RegexErrorPosixCollatingElementsNotSupported                  = 131
	RegexErrorHexCodeTooLarge                                     = 134
	RegexErrorInvalidCondition                                    = 135
	RegexErrorSingleByteMatchInLookbehind                         = 136
	RegexErrorInfiniteLoop                                        = 140
	RegexErrorMissingSubpatternNameTerminator                     = 142
	RegexErrorDuplicateSubpatternName                             = 143
	RegexErrorMalformedProperty                                   = 146
	RegexErrorUnknownProperty                                     = 147
	RegexErrorSubpatternNameTooLong                               = 148
	RegexErrorTooManySubpatterns                                  = 149
	RegexErrorInvalidOctalValue                                   = 151
	RegexErrorTooManyBranchesInDefine                             = 154
	RegexErrorDefineRepetion                                      = 155
	RegexErrorInconsistentNewlineOptions                          = 156
	RegexErrorMissingBackReference                                = 157
	RegexErrorInvalidRelativeReference                            = 158
	RegexErrorBacktrackingControlVerbArgumentForbidden            = 159
	RegexErrorUnknownBacktrackingControlVerb                      = 160
	RegexErrorNumberTooBig                                        = 161
	RegexErrorMissingSubpatternName                               = 162
	RegexErrorMissingDigit                                        = 163
	RegexErrorInvalidDataCharacter                                = 164
	RegexErrorExtraSubpatternName                                 = 165
	RegexErrorBacktrackingControlVerbArgumentRequired             = 166
	RegexErrorInvalidControlChar                                  = 168
	RegexErrorMissingName                                         = 169
	RegexErrorNotSupportedInClass                                 = 171
	RegexErrorTooManyForwardReferences                            = 172
	RegexErrorNameTooLong                                         = 175
	RegexErrorCharacterValueTooLarge                              = 176
)

type SeekType int

const (
	SeekTypeCur SeekType = 0
	SeekTypeSet          = 1
	SeekTypeEnd          = 2
)

type ShellError int

const (
	ShellErrorBadQuoting  ShellError = 0
	ShellErrorEmptyString            = 1
	ShellErrorFailed                 = 2
)

type SliceConfig int

const (
	SliceConfigAlwaysMalloc      SliceConfig = 1
	SliceConfigBypassMagazines               = 2
	SliceConfigWorkingSetMsecs               = 3
	SliceConfigColorIncrement                = 4
	SliceConfigChunkSizes                    = 5
	SliceConfigContentionCounter             = 6
)

type SpawnError int

const (
	SpawnErrorFork        SpawnError = 0
	SpawnErrorRead                   = 1
	SpawnErrorChdir                  = 2
	SpawnErrorAcces                  = 3
	SpawnErrorPerm                   = 4
	SpawnErrorTooBig                 = 5
	SpawnError2big                   = 5
	SpawnErrorNoexec                 = 6
	SpawnErrorNametoolong            = 7
	SpawnErrorNoent                  = 8
	SpawnErrorNomem                  = 9
	SpawnErrorNotdir                 = 10
	SpawnErrorLoop                   = 11
	SpawnErrorTxtbusy                = 12
	SpawnErrorIo                     = 13
	SpawnErrorNfile                  = 14
	SpawnErrorMfile                  = 15
	SpawnErrorInval                  = 16
	SpawnErrorIsdir                  = 17
	SpawnErrorLibbad                 = 18
	SpawnErrorFailed                 = 19
)

type TestFileType int

const (
	TestFileTypeDist  TestFileType = 0
	TestFileTypeBuilt              = 1
)

type TestLogType int

const (
	TestLogTypeNone        TestLogType = 0
	TestLogTypeError                   = 1
	TestLogTypeStartBinary             = 2
	TestLogTypeListCase                = 3
	TestLogTypeSkipCase                = 4
	TestLogTypeStartCase               = 5
	TestLogTypeStopCase                = 6
	TestLogTypeMinResult               = 7
	TestLogTypeMaxResult               = 8
	TestLogTypeMessage                 = 9
	TestLogTypeStartSuite              = 10
	TestLogTypeStopSuite               = 11
)

type ThreadError int

const (
	ThreadErrorThreadErrorAgain ThreadError = 0
)

type TimeType int

const (
	TimeTypeStandard  TimeType = 0
	TimeTypeDaylight           = 1
	TimeTypeUniversal          = 2
)

type TokenType int

const (
	TokenTypeEof            TokenType = 0
	TokenTypeLeftParen                = 40
	TokenTypeRightParen               = 41
	TokenTypeLeftCurly                = 123
	TokenTypeRightCurly               = 125
	TokenTypeLeftBrace                = 91
	TokenTypeRightBrace               = 93
	TokenTypeEqualSign                = 61
	TokenTypeComma                    = 44
	TokenTypeNone                     = 256
	TokenTypeError                    = 257
	TokenTypeChar                     = 258
	TokenTypeBinary                   = 259
	TokenTypeOctal                    = 260
	TokenTypeInt                      = 261
	TokenTypeHex                      = 262
	TokenTypeFloat                    = 263
	TokenTypeString                   = 264
	TokenTypeSymbol                   = 265
	TokenTypeIdentifier               = 266
	TokenTypeIdentifierNull           = 267
	TokenTypeCommentSingle            = 268
	TokenTypeCommentMulti             = 269
)

type TraverseType int

const (
	TraverseTypeInOrder    TraverseType = 0
	TraverseTypePreOrder                = 1
	TraverseTypePostOrder               = 2
	TraverseTypeLevelOrder              = 3
)

type UnicodeBreakType int

const (
	UnicodeBreakTypeMandatory                  UnicodeBreakType = 0
	UnicodeBreakTypeCarriageReturn                              = 1
	UnicodeBreakTypeLineFeed                                    = 2
	UnicodeBreakTypeCombiningMark                               = 3
	UnicodeBreakTypeSurrogate                                   = 4
	UnicodeBreakTypeZeroWidthSpace                              = 5
	UnicodeBreakTypeInseparable                                 = 6
	UnicodeBreakTypeNonBreakingGlue                             = 7
	UnicodeBreakTypeContingent                                  = 8
	UnicodeBreakTypeSpace                                       = 9
	UnicodeBreakTypeAfter                                       = 10
	UnicodeBreakTypeBefore                                      = 11
	UnicodeBreakTypeBeforeAndAfter                              = 12
	UnicodeBreakTypeHyphen                                      = 13
	UnicodeBreakTypeNonStarter                                  = 14
	UnicodeBreakTypeOpenPunctuation                             = 15
	UnicodeBreakTypeClosePunctuation                            = 16
	UnicodeBreakTypeQuotation                                   = 17
	UnicodeBreakTypeExclamation                                 = 18
	UnicodeBreakTypeIdeographic                                 = 19
	UnicodeBreakTypeNumeric                                     = 20
	UnicodeBreakTypeInfixSeparator                              = 21
	UnicodeBreakTypeSymbol                                      = 22
	UnicodeBreakTypeAlphabetic                                  = 23
	UnicodeBreakTypePrefix                                      = 24
	UnicodeBreakTypePostfix                                     = 25
	UnicodeBreakTypeComplexContext                              = 26
	UnicodeBreakTypeAmbiguous                                   = 27
	UnicodeBreakTypeUnknown                                     = 28
	UnicodeBreakTypeNextLine                                    = 29
	UnicodeBreakTypeWordJoiner                                  = 30
	UnicodeBreakTypeHangulLJamo                                 = 31
	UnicodeBreakTypeHangulVJamo                                 = 32
	UnicodeBreakTypeHangulTJamo                                 = 33
	UnicodeBreakTypeHangulLvSyllable                            = 34
	UnicodeBreakTypeHangulLvtSyllable                           = 35
	UnicodeBreakTypeCloseParanthesis                            = 36
	UnicodeBreakTypeConditionalJapaneseStarter                  = 37
	UnicodeBreakTypeHebrewLetter                                = 38
	UnicodeBreakTypeRegionalIndicator                           = 39
	UnicodeBreakTypeEmojiBase                                   = 40
	UnicodeBreakTypeEmojiModifier                               = 41
	UnicodeBreakTypeZeroWidthJoiner                             = 42
)

type UnicodeScript int

const (
	UnicodeScriptInvalidCode           UnicodeScript = -1
	UnicodeScriptCommon                              = 0
	UnicodeScriptInherited                           = 1
	UnicodeScriptArabic                              = 2
	UnicodeScriptArmenian                            = 3
	UnicodeScriptBengali                             = 4
	UnicodeScriptBopomofo                            = 5
	UnicodeScriptCherokee                            = 6
	UnicodeScriptCoptic                              = 7
	UnicodeScriptCyrillic                            = 8
	UnicodeScriptDeseret                             = 9
	UnicodeScriptDevanagari                          = 10
	UnicodeScriptEthiopic                            = 11
	UnicodeScriptGeorgian                            = 12
	UnicodeScriptGothic                              = 13
	UnicodeScriptGreek                               = 14
	UnicodeScriptGujarati                            = 15
	UnicodeScriptGurmukhi                            = 16
	UnicodeScriptHan                                 = 17
	UnicodeScriptHangul                              = 18
	UnicodeScriptHebrew                              = 19
	UnicodeScriptHiragana                            = 20
	UnicodeScriptKannada                             = 21
	UnicodeScriptKatakana                            = 22
	UnicodeScriptKhmer                               = 23
	UnicodeScriptLao                                 = 24
	UnicodeScriptLatin                               = 25
	UnicodeScriptMalayalam                           = 26
	UnicodeScriptMongolian                           = 27
	UnicodeScriptMyanmar                             = 28
	UnicodeScriptOgham                               = 29
	UnicodeScriptOldItalic                           = 30
	UnicodeScriptOriya                               = 31
	UnicodeScriptRunic                               = 32
	UnicodeScriptSinhala                             = 33
	UnicodeScriptSyriac                              = 34
	UnicodeScriptTamil                               = 35
	UnicodeScriptTelugu                              = 36
	UnicodeScriptThaana                              = 37
	UnicodeScriptThai                                = 38
	UnicodeScriptTibetan                             = 39
	UnicodeScriptCanadianAboriginal                  = 40
	UnicodeScriptYi                                  = 41
	UnicodeScriptTagalog                             = 42
	UnicodeScriptHanunoo                             = 43
	UnicodeScriptBuhid                               = 44
	UnicodeScriptTagbanwa                            = 45
	UnicodeScriptBraille                             = 46
	UnicodeScriptCypriot                             = 47
	UnicodeScriptLimbu                               = 48
	UnicodeScriptOsmanya                             = 49
	UnicodeScriptShavian                             = 50
	UnicodeScriptLinearB                             = 51
	UnicodeScriptTaiLe                               = 52
	UnicodeScriptUgaritic                            = 53
	UnicodeScriptNewTaiLue                           = 54
	UnicodeScriptBuginese                            = 55
	UnicodeScriptGlagolitic                          = 56
	UnicodeScriptTifinagh                            = 57
	UnicodeScriptSylotiNagri                         = 58
	UnicodeScriptOldPersian                          = 59
	UnicodeScriptKharoshthi                          = 60
	UnicodeScriptUnknown                             = 61
	UnicodeScriptBalinese                            = 62
	UnicodeScriptCuneiform                           = 63
	UnicodeScriptPhoenician                          = 64
	UnicodeScriptPhagsPa                             = 65
	UnicodeScriptNko                                 = 66
	UnicodeScriptKayahLi                             = 67
	UnicodeScriptLepcha                              = 68
	UnicodeScriptRejang                              = 69
	UnicodeScriptSundanese                           = 70
	UnicodeScriptSaurashtra                          = 71
	UnicodeScriptCham                                = 72
	UnicodeScriptOlChiki                             = 73
	UnicodeScriptVai                                 = 74
	UnicodeScriptCarian                              = 75
	UnicodeScriptLycian                              = 76
	UnicodeScriptLydian                              = 77
	UnicodeScriptAvestan                             = 78
	UnicodeScriptBamum                               = 79
	UnicodeScriptEgyptianHieroglyphs                 = 80
	UnicodeScriptImperialAramaic                     = 81
	UnicodeScriptInscriptionalPahlavi                = 82
	UnicodeScriptInscriptionalParthian               = 83
	UnicodeScriptJavanese                            = 84
	UnicodeScriptKaithi                              = 85
	UnicodeScriptLisu                                = 86
	UnicodeScriptMeeteiMayek                         = 87
	UnicodeScriptOldSouthArabian                     = 88
	UnicodeScriptOldTurkic                           = 89
	UnicodeScriptSamaritan                           = 90
	UnicodeScriptTaiTham                             = 91
	UnicodeScriptTaiViet                             = 92
	UnicodeScriptBatak                               = 93
	UnicodeScriptBrahmi                              = 94
	UnicodeScriptMandaic                             = 95
	UnicodeScriptChakma                              = 96
	UnicodeScriptMeroiticCursive                     = 97
	UnicodeScriptMeroiticHieroglyphs                 = 98
	UnicodeScriptMiao                                = 99
	UnicodeScriptSharada                             = 100
	UnicodeScriptSoraSompeng                         = 101
	UnicodeScriptTakri                               = 102
	UnicodeScriptBassaVah                            = 103
	UnicodeScriptCaucasianAlbanian                   = 104
	UnicodeScriptDuployan                            = 105
	UnicodeScriptElbasan                             = 106
	UnicodeScriptGrantha                             = 107
	UnicodeScriptKhojki                              = 108
	UnicodeScriptKhudawadi                           = 109
	UnicodeScriptLinearA                             = 110
	UnicodeScriptMahajani                            = 111
	UnicodeScriptManichaean                          = 112
	UnicodeScriptMendeKikakui                        = 113
	UnicodeScriptModi                                = 114
	UnicodeScriptMro                                 = 115
	UnicodeScriptNabataean                           = 116
	UnicodeScriptOldNorthArabian                     = 117
	UnicodeScriptOldPermic                           = 118
	UnicodeScriptPahawhHmong                         = 119
	UnicodeScriptPalmyrene                           = 120
	UnicodeScriptPauCinHau                           = 121
	UnicodeScriptPsalterPahlavi                      = 122
	UnicodeScriptSiddham                             = 123
	UnicodeScriptTirhuta                             = 124
	UnicodeScriptWarangCiti                          = 125
	UnicodeScriptAhom                                = 126
	UnicodeScriptAnatolianHieroglyphs                = 127
	UnicodeScriptHatran                              = 128
	UnicodeScriptMultani                             = 129
	UnicodeScriptOldHungarian                        = 130
	UnicodeScriptSignwriting                         = 131
	UnicodeScriptAdlam                               = 132
	UnicodeScriptBhaiksuki                           = 133
	UnicodeScriptMarchen                             = 134
	UnicodeScriptNewa                                = 135
	UnicodeScriptOsage                               = 136
	UnicodeScriptTangut                              = 137
)

type UnicodeType int

const (
	UnicodeTypeControl            UnicodeType = 0
	UnicodeTypeFormat                         = 1
	UnicodeTypeUnassigned                     = 2
	UnicodeTypePrivateUse                     = 3
	UnicodeTypeSurrogate                      = 4
	UnicodeTypeLowercaseLetter                = 5
	UnicodeTypeModifierLetter                 = 6
	UnicodeTypeOtherLetter                    = 7
	UnicodeTypeTitlecaseLetter                = 8
	UnicodeTypeUppercaseLetter                = 9
	UnicodeTypeSpacingMark                    = 10
	UnicodeTypeEnclosingMark                  = 11
	UnicodeTypeNonSpacingMark                 = 12
	UnicodeTypeDecimalNumber                  = 13
	UnicodeTypeLetterNumber                   = 14
	UnicodeTypeOtherNumber                    = 15
	UnicodeTypeConnectPunctuation             = 16
	UnicodeTypeDashPunctuation                = 17
	UnicodeTypeClosePunctuation               = 18
	UnicodeTypeFinalPunctuation               = 19
	UnicodeTypeInitialPunctuation             = 20
	UnicodeTypeOtherPunctuation               = 21
	UnicodeTypeOpenPunctuation                = 22
	UnicodeTypeCurrencySymbol                 = 23
	UnicodeTypeModifierSymbol                 = 24
	UnicodeTypeMathSymbol                     = 25
	UnicodeTypeOtherSymbol                    = 26
	UnicodeTypeLineSeparator                  = 27
	UnicodeTypeParagraphSeparator             = 28
	UnicodeTypeSpaceSeparator                 = 29
)

type UserDirectory int

const (
	UserDirectoryDirectoryDesktop     UserDirectory = 0
	UserDirectoryDirectoryDocuments                 = 1
	UserDirectoryDirectoryDownload                  = 2
	UserDirectoryDirectoryMusic                     = 3
	UserDirectoryDirectoryPictures                  = 4
	UserDirectoryDirectoryPublicShare               = 5
	UserDirectoryDirectoryTemplates                 = 6
	UserDirectoryDirectoryVideos                    = 7
	UserDirectoryNDirectories                       = 8
)

type VariantClass int

const (
	VariantClassBoolean    VariantClass = 98
	VariantClassByte                    = 121
	VariantClassInt16                   = 110
	VariantClassUint16                  = 113
	VariantClassInt32                   = 105
	VariantClassUint32                  = 117
	VariantClassInt64                   = 120
	VariantClassUint64                  = 116
	VariantClassHandle                  = 104
	VariantClassDouble                  = 100
	VariantClassString                  = 115
	VariantClassObjectPath              = 111
	VariantClassSignature               = 103
	VariantClassVariant                 = 118
	VariantClassMaybe                   = 109
	VariantClassArray                   = 97
	VariantClassTuple                   = 40
	VariantClassDictEntry               = 123
)

type VariantParseError int

const (
	VariantParseErrorFailed                     VariantParseError = 0
	VariantParseErrorBasicTypeExpected                            = 1
	VariantParseErrorCannotInferType                              = 2
	VariantParseErrorDefiniteTypeExpected                         = 3
	VariantParseErrorInputNotAtEnd                                = 4
	VariantParseErrorInvalidCharacter                             = 5
	VariantParseErrorInvalidFormatString                          = 6
	VariantParseErrorInvalidObjectPath                            = 7
	VariantParseErrorInvalidSignature                             = 8
	VariantParseErrorInvalidTypeString                            = 9
	VariantParseErrorNoCommonType                                 = 10
	VariantParseErrorNumberOutOfRange                             = 11
	VariantParseErrorNumberTooBig                                 = 12
	VariantParseErrorTypeError                                    = 13
	VariantParseErrorUnexpectedToken                              = 14
	VariantParseErrorUnknownKeyword                               = 15
	VariantParseErrorUnterminatedStringConstant                   = 16
	VariantParseErrorValueExpected                                = 17
)

type AsciiType int

const (
	AsciiTypeAlnum  AsciiType = 1
	AsciiTypeAlpha            = 2
	AsciiTypeCntrl            = 4
	AsciiTypeDigit            = 8
	AsciiTypeGraph            = 16
	AsciiTypeLower            = 32
	AsciiTypePrint            = 64
	AsciiTypePunct            = 128
	AsciiTypeSpace            = 256
	AsciiTypeUpper            = 512
	AsciiTypeXdigit           = 1024
)

type FileTest int

const (
	FileTestIsRegular    FileTest = 1
	FileTestIsSymlink             = 2
	FileTestIsDir                 = 4
	FileTestIsExecutable          = 8
	FileTestExists                = 16
)

type FormatSizeFlags int

const (
	FormatSizeFlagsDefault    FormatSizeFlags = 0
	FormatSizeFlagsLongFormat                 = 1
	FormatSizeFlagsIecUnits                   = 2
)

type HookFlagMask int

const (
	HookFlagMaskActive HookFlagMask = 1
	HookFlagMaskInCall              = 2
	HookFlagMaskMask                = 15
)

type IOCondition int

const (
	IOConditionIn   IOCondition = 1
	IOConditionOut              = 4
	IOConditionPri              = 2
	IOConditionErr              = 8
	IOConditionHup              = 16
	IOConditionNval             = 32
)

type IOFlags int

const (
	IOFlagsAppend      IOFlags = 1
	IOFlagsNonblock            = 2
	IOFlagsIsReadable          = 4
	IOFlagsIsWritable          = 8
	IOFlagsIsWriteable         = 8
	IOFlagsIsSeekable          = 16
	IOFlagsMask                = 31
	IOFlagsGetMask             = 31
	IOFlagsSetMask             = 3
)

type KeyFileFlags int

const (
	KeyFileFlagsNone             KeyFileFlags = 0
	KeyFileFlagsKeepComments                  = 1
	KeyFileFlagsKeepTranslations              = 2
)

type LogLevelFlags int

const (
	LogLevelFlagsFlagRecursion LogLevelFlags = 1
	LogLevelFlagsFlagFatal                   = 2
	LogLevelFlagsLevelError                  = 4
	LogLevelFlagsLevelCritical               = 8
	LogLevelFlagsLevelWarning                = 16
	LogLevelFlagsLevelMessage                = 32
	LogLevelFlagsLevelInfo                   = 64
	LogLevelFlagsLevelDebug                  = 128
	LogLevelFlagsLevelMask                   = -4
)

type MarkupCollectType int

const (
	MarkupCollectTypeInvalid  MarkupCollectType = 0
	MarkupCollectTypeString                     = 1
	MarkupCollectTypeStrdup                     = 2
	MarkupCollectTypeBoolean                    = 3
	MarkupCollectTypeTristate                   = 4
	MarkupCollectTypeOptional                   = 65536
)

type MarkupParseFlags int

const (
	MarkupParseFlagsDoNotUseThisUnsupportedFlag MarkupParseFlags = 1
	MarkupParseFlagsTreatCdataAsText                             = 2
	MarkupParseFlagsPrefixErrorPosition                          = 4
	MarkupParseFlagsIgnoreQualified                              = 8
)

type OptionFlags int

const (
	OptionFlagsNone        OptionFlags = 0
	OptionFlagsHidden                  = 1
	OptionFlagsInMain                  = 2
	OptionFlagsReverse                 = 4
	OptionFlagsNoArg                   = 8
	OptionFlagsFilename                = 16
	OptionFlagsOptionalArg             = 32
	OptionFlagsNoalias                 = 64
)

type RegexCompileFlags int

const (
	RegexCompileFlagsCaseless         RegexCompileFlags = 1
	RegexCompileFlagsMultiline                          = 2
	RegexCompileFlagsDotall                             = 4
	RegexCompileFlagsExtended                           = 8
	RegexCompileFlagsAnchored                           = 16
	RegexCompileFlagsDollarEndonly                      = 32
	RegexCompileFlagsUngreedy                           = 512
	RegexCompileFlagsRaw                                = 2048
	RegexCompileFlagsNoAutoCapture                      = 4096
	RegexCompileFlagsOptimize                           = 8192
	RegexCompileFlagsFirstline                          = 262144
	RegexCompileFlagsDupnames                           = 524288
	RegexCompileFlagsNewlineCr                          = 1048576
	RegexCompileFlagsNewlineLf                          = 2097152
	RegexCompileFlagsNewlineCrlf                        = 3145728
	RegexCompileFlagsNewlineAnycrlf                     = 5242880
	RegexCompileFlagsBsrAnycrlf                         = 8388608
	RegexCompileFlagsJavascriptCompat                   = 33554432
)

type RegexMatchFlags int

const (
	RegexMatchFlagsAnchored        RegexMatchFlags = 16
	RegexMatchFlagsNotbol                          = 128
	RegexMatchFlagsNoteol                          = 256
	RegexMatchFlagsNotempty                        = 1024
	RegexMatchFlagsPartial                         = 32768
	RegexMatchFlagsNewlineCr                       = 1048576
	RegexMatchFlagsNewlineLf                       = 2097152
	RegexMatchFlagsNewlineCrlf                     = 3145728
	RegexMatchFlagsNewlineAny                      = 4194304
	RegexMatchFlagsNewlineAnycrlf                  = 5242880
	RegexMatchFlagsBsrAnycrlf                      = 8388608
	RegexMatchFlagsBsrAny                          = 16777216
	RegexMatchFlagsPartialSoft                     = 32768
	RegexMatchFlagsPartialHard                     = 134217728
	RegexMatchFlagsNotemptyAtstart                 = 268435456
)

type SpawnFlags int

const (
	SpawnFlagsDefault              SpawnFlags = 0
	SpawnFlagsLeaveDescriptorsOpen            = 1
	SpawnFlagsDoNotReapChild                  = 2
	SpawnFlagsSearchPath                      = 4
	SpawnFlagsStdoutToDevNull                 = 8
	SpawnFlagsStderrToDevNull                 = 16
	SpawnFlagsChildInheritsStdin              = 32
	SpawnFlagsFileAndArgvZero                 = 64
	SpawnFlagsSearchPathFromEnvp              = 128
	SpawnFlagsCloexecPipes                    = 256
)

type TestSubprocessFlags int

const (
	TestSubprocessFlagsStdin  TestSubprocessFlags = 1
	TestSubprocessFlagsStdout                     = 2
	TestSubprocessFlagsStderr                     = 4
)

type TestTrapFlags int

const (
	TestTrapFlagsSilenceStdout TestTrapFlags = 128
	TestTrapFlagsSilenceStderr               = 256
	TestTrapFlagsInheritStdin                = 512
)

type TraverseFlags int

const (
	TraverseFlagsLeaves    TraverseFlags = 1
	TraverseFlagsNonLeaves               = 2
	TraverseFlagsAll                     = 3
	TraverseFlagsMask                    = 3
	TraverseFlagsLeafs                   = 1
	TraverseFlagsNonLeafs                = 2
)
