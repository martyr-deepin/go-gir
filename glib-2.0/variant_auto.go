package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
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
