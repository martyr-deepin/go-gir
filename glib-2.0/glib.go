package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import (
	"time"
	"unsafe"
)

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

func (v List) InsertBefore(sibling List, data unsafe.Pointer) List {
	list := C.g_list_insert_before(v.native(), sibling.native(), C.gpointer(data))
	return wrapList(list, v.DataWrap)
}

type SList struct {
	Ptr unsafe.Pointer
	DataWrap DataWrapFunc
}

func WrapSList(p unsafe.Pointer, dataWrap DataWrapFunc) SList {
	return SList{Ptr: p, DataWrap: dataWrap}
}

func wrapSList(p *C.GSList, dataWrap DataWrapFunc) SList {
	return SList{Ptr: unsafe.Pointer(p), DataWrap: dataWrap}
}

func (v SList) native() *C.GSList {
	return (*C.GSList)(v.Ptr)
}

func (v SList) Append(data unsafe.Pointer) SList {
	list := C.g_slist_append(v.native(), C.gpointer(data))
	return wrapSList(list, v.DataWrap)
}

func (v SList) Prepend(data unsafe.Pointer) SList {
	list := C.g_slist_prepend(v.native(), C.gpointer(data))
	return wrapSList(list, v.DataWrap)
}

func (v SList) Insert(data unsafe.Pointer, position int) SList {
	list := C.g_slist_insert(v.native(), C.gpointer(data), C.gint(position))
	return wrapSList(list, v.DataWrap)
}

// TimeVal
func TimeValNew() TimeVal {
	ptr := C.g_malloc0(C.gsize(C.sizeof_GTimeVal))
	return TimeVal{
		Ptr: unsafe.Pointer(ptr),
	}
}

func (v TimeVal) Free() {
	C.g_free(C.gpointer(v.Ptr))
}

// field tv_sec
func (v TimeVal) Seconds() int64 {
	native := v.native()
	return int64(native.tv_sec)
}

func (v TimeVal) SetSeconds(secs int64) {
	native := v.native()
	native.tv_sec = C.glong(secs)
}

// field tv_usec
func (v TimeVal) Microseconds() int64 {
	native := v.native()
	return int64(native.tv_usec)
}

func (v TimeVal) SetMicroseconds(usecs int64) {
	native := v.native()
	native.tv_usec = C.glong(usecs)
}

func (v TimeVal) GoValue() time.Time {
	native := v.native()
	return time.Unix(int64(native.tv_sec), int64(native.tv_usec)*1000)
}
