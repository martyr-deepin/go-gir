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
	"unsafe"

	"fmt"
	"github.com/electricface/go-auto-gir/glib-2.0"
	"github.com/electricface/go-auto-gir/util"
	"sync"
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

var gvalueGetters = struct {
	m map[Type]GValueGetter
	sync.Mutex
}{
	m: make(map[Type]GValueGetter),
}

type GValueGetter func(unsafe.Pointer) (interface{}, error)

func registerGValueGetter(typ Type, getter GValueGetter) {
	gvalueGetters.Lock()
	gvalueGetters.m[typ] = getter
	gvalueGetters.Unlock()
}

func (v Value) Get() (interface{}, error) {
	actualType, fundamentalType, err := v.Type()
	if err != nil {
		return nil, err
	}
	fmt.Println("actual Type:", actualType)
	fmt.Println("fund Type:", fundamentalType)

	val, err := v.get(actualType)
	if err == nil {
		// good
		return val, nil
	} else if err == errTypeUnknown {
		// fallback to fundamental type
		return v.get(fundamentalType)
	}
	return nil, err
}

var errTypeUnknown = errors.New("unknown type")

func (v Value) get(typ Type) (interface{}, error) {
	switch typ {
	case TYPE_INVALID:
		return nil, errors.New("invalid type")
	case TYPE_NONE:
		return nil, nil
	case TYPE_INTERFACE:
		return nil, errors.New("interface conversion not yet implemented")
	case TYPE_CHAR:
		ret := C.g_value_get_schar(v.native())
		return int(ret), nil
	case TYPE_UCHAR:
		ret := C.g_value_get_uchar(v.native())
		return byte(ret), nil
	case TYPE_BOOLEAN:
		ret := C.g_value_get_boolean(v.native())
		return util.Int2Bool(int(ret)), nil
	case TYPE_INT:
		ret := C.g_value_get_int(v.native())
		return int(ret), nil
	case TYPE_LONG:
		ret := C.g_value_get_long(v.native())
		return int(ret), nil
	case TYPE_ENUM:
		ret := C.g_value_get_enum(v.native())
		return int(ret), nil
	case TYPE_INT64:
		ret := C.g_value_get_int64(v.native())
		return int64(ret), nil
	case TYPE_UINT:
		ret := C.g_value_get_uint(v.native())
		return uint(ret), nil
	case TYPE_FLAGS:
		ret := C.g_value_get_flags(v.native())
		return uint(ret), nil
	case TYPE_UINT64:
		ret := C.g_value_get_uint64(v.native())
		return uint64(ret), nil
	case TYPE_FLOAT:
		ret := C.g_value_get_float(v.native())
		return float32(ret), nil
	case TYPE_DOUBLE:
		ret := C.g_value_get_double(v.native())
		return float64(ret), nil
	case TYPE_STRING:
		ret := C.g_value_get_string(v.native())
		return C.GoString((*C.char)(ret)), nil
	case TYPE_POINTER:
		ret := C.g_value_get_pointer(v.native())
		return unsafe.Pointer(ret), nil
	case TYPE_BOXED:
		ret := C.g_value_get_boxed(v.native())
		return unsafe.Pointer(ret), nil
	case TYPE_OBJECT:
		ret := C.g_value_get_object(v.native())
		return WrapObject(unsafe.Pointer(ret)), nil
	case TYPE_VARIANT:
		ret := C.g_value_get_variant(v.native())
		return glib.WrapVariant(unsafe.Pointer(ret)), nil
	}

	gvalueGetters.Lock()
	getter, ok := gvalueGetters.m[typ]
	gvalueGetters.Unlock()
	if !ok {
		return nil, errTypeUnknown
	}
	return getter(v.Ptr)
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
