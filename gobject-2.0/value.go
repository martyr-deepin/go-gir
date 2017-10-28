package gobject

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <glib-object.h>
#include <stdlib.h>

static GValue *
_g_value_alloc()
{
	return (g_new0(GValue, 1));
}

static GValue *
_g_value_init(GType g_type)
{
	GValue          *value;

	value = g_new0(GValue, 1);
	return (g_value_init(value, g_type));
}

static gboolean
_g_is_value(GValue *val)
{
	return (G_IS_VALUE(val));
}

static GType
_g_value_type(GValue *val)
{
	return (G_VALUE_TYPE(val));
}

static GType
_g_value_fundamental(GType type)
{
	return (G_TYPE_FUNDAMENTAL(type));
}

static GObjectClass *
_g_object_get_class (GObject *object)
{
	return (G_OBJECT_GET_CLASS(object));
}
*/
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
	"github.com/electricface/go-auto-gir/util"
)

/*
 * GValue
 */

// ValueAlloc allocates a Value and sets a runtime finalizer to call
// g_value_unset() on the underlying GValue after leaving scope.
// ValueAlloc() returns a non-nil error if the allocation failed.

func ValueNew() Value {
	return wrapValue(C._g_value_alloc())
}

func (v Value) Free() {
	C.g_free(C.gpointer(v.native()))
}

// ValueInit is a wrapper around g_value_init() and allocates and
// initializes a new Value with the Type t.
func (v Value) Init(gType Type) {
	C.g_value_init(v.native(), C.GType(gType))
}

// ValueFromNative returns a type-asserted pointer to the Value.
//func ValueFromNative(l unsafe.Pointer) Value {
//TODO why it does not add finalizer to the value?
//return Value{(*C.GValue)(l)}
//}

func (v *Value) Unset() {
	C.g_value_unset(v.native())
}

// Type is a wrapper around the G_VALUE_HOLDS_GTYPE() macro and
// the g_value_get_gtype() function.  GetType() returns TYPE_INVALID if v
// does not hold a Type, or otherwise returns the Type of v.
func (v Value) Type() (actual Type, fundamental Type, err error) {
	if !util.Int2Bool(int(C._g_is_value(v.native()))) {
		return actual, fundamental, errors.New("invalid GValue")
	}
	cActual := C._g_value_type(v.native())
	cFundamental := C._g_value_fundamental(cActual)
	return Type(cActual), Type(cFundamental), nil
}

// GValue converts a Go type to a comparable GValue.  GValue()
// returns a non-nil error if the conversion was unsuccessful.
func ToGValue(v interface{}) (gvalue Value, err error) {
	gvalue = ValueNew()
	if v == nil {
		gvalue.Init(TYPE_POINTER)
		gvalue.SetPointer(unsafe.Pointer(nil))
		return gvalue, nil
	}

	switch e := v.(type) {
	case bool:
		gvalue.Init(TYPE_BOOLEAN)
		gvalue.SetBoolean(e)
		return gvalue, nil

	case int8:
		gvalue.Init(TYPE_CHAR)
		gvalue.SetSchar(e)
		return gvalue, nil

	case int64:
		gvalue.Init(TYPE_INT64)
		gvalue.SetInt64(e)
		return gvalue, nil

	case int:
		gvalue.Init(TYPE_INT)
		gvalue.SetInt(e)
		return gvalue, nil

	case uint8:
		gvalue.Init(TYPE_UCHAR)
		gvalue.SetUchar(e)
		return gvalue, nil

	case uint64:
		gvalue.Init(TYPE_UINT64)
		gvalue.SetUint64(e)
		return gvalue, nil

	case uint:
		gvalue.Init(TYPE_UINT)
		gvalue.SetUint(e)
		return gvalue, nil

	case float32:
		gvalue.Init(TYPE_FLOAT)
		gvalue.SetFloat(e)
		return gvalue, nil

	case float64:
		gvalue.Init(TYPE_DOUBLE)
		gvalue.SetDouble(e)
		return gvalue, nil

	case string:
		gvalue.Init(TYPE_STRING)
		gvalue.SetString(e)
		return gvalue, nil

	case Object:
		gvalue.Init(TYPE_OBJECT)
		gvalue.SetInstance(e.Ptr)
		return gvalue, nil

	default:
		/* Try this since above doesn't catch constants under other types */
		rval := reflect.ValueOf(v)
		switch rval.Kind() {
		case reflect.Int8:
			gvalue.Init(TYPE_CHAR)
			gvalue.SetSchar(int8(rval.Int()))
			return gvalue, nil

			//case reflect.Int16:
			//case reflect.Int32:

		case reflect.Int64:
			gvalue.Init(TYPE_INT64)
			gvalue.SetInt64(rval.Int())
			return gvalue, nil

		case reflect.Int:
			gvalue.Init(TYPE_INT)
			gvalue.SetInt(int(rval.Int()))
			return gvalue, nil

		case reflect.Uintptr, reflect.Ptr:
			gvalue.Init(TYPE_POINTER)
			gvalue.SetPointer(unsafe.Pointer(rval.Pointer()))
			return gvalue, nil
		}
	}

	return Value{}, errors.New("Type not implemented")
}

// GValueMarshaler is a marshal function to convert a GValue into an
// appropiate Go type.  The uintptr parameter is a *C.GValue.
type GValueMarshaler func(uintptr) (interface{}, error)

// TypeMarshaler represents an actual type and it's associated marshaler.
type TypeMarshaler struct {
	T Type
	F GValueMarshaler
}

// RegisterGValueMarshalers adds marshalers for several types to the
// internal marshalers map.  Once registered, calling GoValue on any
// Value witha registered type will return the data returned by the
// marshaler.
func RegisterGValueMarshalers(tm []TypeMarshaler) {
	gValueMarshalers.register(tm)
}

func RegisterGValueMarshaler(t Type, f GValueMarshaler) {
	gValueMarshalers.registerOne(t, f)
}

type marshalMap map[Type]GValueMarshaler

// gValueMarshalers is a map of Glib types to functions to marshal a
// GValue to a native Go type.
var gValueMarshalers = marshalMap{
	TYPE_INVALID:   marshalInvalid,
	TYPE_NONE:      marshalNone,
	TYPE_INTERFACE: marshalInterface,
	TYPE_CHAR:      marshalChar,
	TYPE_UCHAR:     marshalUchar,
	TYPE_BOOLEAN:   marshalBoolean,
	TYPE_INT:       marshalInt,
	TYPE_LONG:      marshalLong,
	TYPE_ENUM:      marshalEnum,
	TYPE_INT64:     marshalInt64,
	TYPE_UINT:      marshalUint,
	TYPE_ULONG:     marshalUlong,
	TYPE_FLAGS:     marshalFlags,
	TYPE_UINT64:    marshalUint64,
	TYPE_FLOAT:     marshalFloat,
	TYPE_DOUBLE:    marshalDouble,
	TYPE_STRING:    marshalString,
	TYPE_POINTER:   marshalPointer,
	TYPE_BOXED:     marshalBoxed,
	TYPE_OBJECT:    marshalObject,
	TYPE_VARIANT:   marshalVariant,
}

func (m marshalMap) registerOne(t Type, f GValueMarshaler) {
	m[t] = f
}

func (m marshalMap) register(tm []TypeMarshaler) {
	for i := range tm {
		m[tm[i].T] = tm[i].F
	}
}

func (m marshalMap) lookup(v Value) (GValueMarshaler, error) {
	actual, fundamental, err := v.Type()
	if err != nil {
		return nil, err
	}

	if f, ok := m[actual]; ok {
		return f, nil
	}
	if f, ok := m[fundamental]; ok {
		return f, nil
	}
	return nil, errors.New("missing marshaler for type")
}

func marshalInvalid(_ uintptr) (interface{}, error) {
	return nil, errors.New("invalid type")
}

func marshalNone(_ uintptr) (interface{}, error) {
	return nil, nil
}

func marshalInterface(_ uintptr) (interface{}, error) {
	return nil, errors.New("interface conversion not yet implemented")
}

func marshalChar(p uintptr) (interface{}, error) {
	c := C.g_value_get_schar((*C.GValue)(unsafe.Pointer(p)))
	return int8(c), nil
}

func marshalUchar(p uintptr) (interface{}, error) {
	c := C.g_value_get_uchar((*C.GValue)(unsafe.Pointer(p)))
	return uint8(c), nil
}

func marshalBoolean(p uintptr) (interface{}, error) {
	c := C.g_value_get_boolean((*C.GValue)(unsafe.Pointer(p)))
	return util.Int2Bool(int(c)), nil
}

func marshalInt(p uintptr) (interface{}, error) {
	c := C.g_value_get_int((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalLong(p uintptr) (interface{}, error) {
	c := C.g_value_get_long((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalEnum(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalInt64(p uintptr) (interface{}, error) {
	c := C.g_value_get_int64((*C.GValue)(unsafe.Pointer(p)))
	return int64(c), nil
}

func marshalUint(p uintptr) (interface{}, error) {
	c := C.g_value_get_uint((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalUlong(p uintptr) (interface{}, error) {
	c := C.g_value_get_ulong((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalUint64(p uintptr) (interface{}, error) {
	c := C.g_value_get_uint64((*C.GValue)(unsafe.Pointer(p)))
	return uint64(c), nil
}

func marshalFloat(p uintptr) (interface{}, error) {
	c := C.g_value_get_float((*C.GValue)(unsafe.Pointer(p)))
	return float32(c), nil
}

func marshalDouble(p uintptr) (interface{}, error) {
	c := C.g_value_get_double((*C.GValue)(unsafe.Pointer(p)))
	return float64(c), nil
}

func marshalString(p uintptr) (interface{}, error) {
	c := C.g_value_get_string((*C.GValue)(unsafe.Pointer(p)))
	return C.GoString((*C.char)(c)), nil
}

func marshalBoxed(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return uintptr(unsafe.Pointer(c)), nil
}

func marshalPointer(p uintptr) (interface{}, error) {
	c := C.g_value_get_pointer((*C.GValue)(unsafe.Pointer(p)))
	return unsafe.Pointer(c), nil
}

func marshalObject(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return WrapObject(unsafe.Pointer(c)), nil
}

func marshalVariant(p uintptr) (interface{}, error) {
	// TODO:
	c := C.g_value_get_variant((*C.GValue)(unsafe.Pointer(p)))
	_ = c
	//return wrapVariant(c), nil
	return nil, nil
}

// GoValue converts a Value to comparable Go type.  GoValue()
// returns a non-nil error if the conversion was unsuccessful.  The
// returned interface{} must be type asserted as the actual Go
// representation of the Value.
//
// This function is a wrapper around the many g_value_get_*()
// functions, depending on the type of the Value.
func (v Value) GoValue() (interface{}, error) {
	f, err := gValueMarshalers.lookup(v)
	if err != nil {
		return nil, err
	}

	//No need to add finalizer because it is already done by ValueAlloc and ValueInit
	rv, err := f(uintptr(unsafe.Pointer(v.native())))
	return rv, err
}

var errTypeConvert = errors.New("type convert failed")

func (v Value) Set(iVal interface{}) error {
	cType := C._g_value_type(v.native())
	gType := Type(cType)
	switch gType {
	case TYPE_CHAR:
		val, ok := iVal.(int8)
		if !ok {
			return errTypeConvert
		}
		v.SetSchar(val)

	case TYPE_UCHAR:
		val, ok := iVal.(uint8)
		if !ok {
			return errTypeConvert
		}
		v.SetUchar(val)

	case TYPE_BOOLEAN:
		val, ok := iVal.(bool)
		if !ok {
			return errTypeConvert
		}
		v.SetBoolean(val)

	case TYPE_INT:
		val, ok := iVal.(int)
		if !ok {
			return errTypeConvert
		}
		v.SetInt(val)

	case TYPE_UINT:
		val, ok := iVal.(uint)
		if !ok {
			return errTypeConvert
		}
		v.SetUint(val)

	case TYPE_INT64:
		val, ok := iVal.(int64)
		if !ok {
			return errTypeConvert
		}
		v.SetInt64(val)

	case TYPE_UINT64:
		val, ok := iVal.(uint64)
		if !ok {
			return errTypeConvert
		}
		v.SetUint64(val)

	case TYPE_FLOAT:
		val, ok := iVal.(float32)
		if !ok {
			return errTypeConvert
		}
		v.SetFloat(val)

	case TYPE_DOUBLE:
		val, ok := iVal.(float64)
		if !ok {
			return errTypeConvert
		}
		v.SetDouble(val)

	case TYPE_STRING:
		val, ok := iVal.(string)
		if !ok {
			return errTypeConvert
		}
		v.SetString(val)

	case TYPE_OBJECT:
		val, ok := iVal.(Object)
		if !ok {
			return errTypeConvert
		}
		v.SetInstance(val.Ptr)

	default:
		return errTypeConvert
	}

	return nil
}
