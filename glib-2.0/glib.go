package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// GetString is a wrapper around g_variant_get_string().
func (variant Variant) GetString() string {
	var length0 C.gsize
	ret0 := C.g_variant_get_string(variant.native(), &length0)
	ret := C.GoStringN((*C.char)(ret0), C.int(length0))
	return ret
}

type DataWrapFunc func(unsafe.Pointer) interface{}

type List struct {
	Ptr      unsafe.Pointer
	DataWrap DataWrapFunc
}

func (v List) native() *C.GList {
	return (*C.GList)(v.Ptr)
}

func WrapList(p unsafe.Pointer, dataWrap DataWrapFunc) List {
	return List{Ptr: p, DataWrap: dataWrap}
}

func wrapList(p *C.GList, dataWrap DataWrapFunc) List {
	return List{Ptr: unsafe.Pointer(p), DataWrap: dataWrap}
}

func (v List) ForEach(fn func(item interface{})) {
	dataWrap := v.DataWrap
	for l := v.native(); l != nil; l = l.next {
		fn(dataWrap(unsafe.Pointer(l.data)))
	}
}

func (v List) FullFree(fn func(item interface{})) {
	v.ForEach(fn)
	v.Free()
}

func (v List) Next() List {
	native := v.native()
	return wrapList(native.next, v.DataWrap)
}

func (v List) Previous() List {
	native := v.native()
	return wrapList(native.prev, v.DataWrap)
}

func (v List) Free() {
	C.g_list_free(v.native())
}

func (v List) Data() interface{} {
	native := v.native()
	return v.DataWrap(unsafe.Pointer(native.data))
}

func (v List) Length() uint {
	return uint(C.g_list_length(v.native()))
}

func (v List) NthData(n uint) interface{} {
	data := C.g_list_nth_data(v.native(), C.guint(n))
	return v.DataWrap(unsafe.Pointer(data))
}

func (v List) Nth(n uint) List {
	list := C.g_list_nth(v.native(), C.guint(n))
	return wrapList(list, v.DataWrap)
}

func (v List) Append(data unsafe.Pointer) List {
	list := C.g_list_append(v.native(), C.gpointer(data))
	return wrapList(list, v.DataWrap)
}

func (v List) Prepend(data unsafe.Pointer) List {
	list := C.g_list_prepend(v.native(), C.gpointer(data))
	return wrapList(list, v.DataWrap)
}

func (v List) Insert(data unsafe.Pointer, position int) List {
	list := C.g_list_insert(v.native(), C.gpointer(data), C.gint(position))
	return wrapList(list, v.DataWrap)
}
