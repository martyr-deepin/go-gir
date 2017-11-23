package cairo

/*
#cgo pkg-config: cairo-gobject
#include <cairo-gobject.h>
*/
import "C"
import "unsafe"

// Struct Context
type Context struct {
	Ptr unsafe.Pointer
}

func (v Context) native() *C.cairo_t {
	return (*C.cairo_t)(v.Ptr)
}
func wrapContext(p *C.cairo_t) Context {
	return Context{unsafe.Pointer(p)}
}
func WrapContext(p unsafe.Pointer) Context {
	return Context{p}
}
func (v Context) IsNil() bool {
	return v.Ptr == nil
}
func IWrapContext(p unsafe.Pointer) interface{} {
	return WrapContext(p)
}

// Struct Surface
type Surface struct {
	Ptr unsafe.Pointer
}

func (v Surface) native() *C.cairo_surface_t {
	return (*C.cairo_surface_t)(v.Ptr)
}
func wrapSurface(p *C.cairo_surface_t) Surface {
	return Surface{unsafe.Pointer(p)}
}
func WrapSurface(p unsafe.Pointer) Surface {
	return Surface{p}
}
func (v Surface) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSurface(p unsafe.Pointer) interface{} {
	return WrapSurface(p)
}

// Struct Matrix
type Matrix struct {
	Ptr unsafe.Pointer
}

func (v Matrix) native() *C.cairo_matrix_t {
	return (*C.cairo_matrix_t)(v.Ptr)
}
func wrapMatrix(p *C.cairo_matrix_t) Matrix {
	return Matrix{unsafe.Pointer(p)}
}
func WrapMatrix(p unsafe.Pointer) Matrix {
	return Matrix{p}
}
func (v Matrix) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMatrix(p unsafe.Pointer) interface{} {
	return WrapMatrix(p)
}

// Struct Pattern
type Pattern struct {
	Ptr unsafe.Pointer
}

func (v Pattern) native() *C.cairo_pattern_t {
	return (*C.cairo_pattern_t)(v.Ptr)
}
func wrapPattern(p *C.cairo_pattern_t) Pattern {
	return Pattern{unsafe.Pointer(p)}
}
func WrapPattern(p unsafe.Pointer) Pattern {
	return Pattern{p}
}
func (v Pattern) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPattern(p unsafe.Pointer) interface{} {
	return WrapPattern(p)
}

// Struct Region
type Region struct {
	Ptr unsafe.Pointer
}

func (v Region) native() *C.cairo_region_t {
	return (*C.cairo_region_t)(v.Ptr)
}
func wrapRegion(p *C.cairo_region_t) Region {
	return Region{unsafe.Pointer(p)}
}
func WrapRegion(p unsafe.Pointer) Region {
	return Region{p}
}
func (v Region) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRegion(p unsafe.Pointer) interface{} {
	return WrapRegion(p)
}

// Struct FontOptions
type FontOptions struct {
	Ptr unsafe.Pointer
}

func (v FontOptions) native() *C.cairo_font_options_t {
	return (*C.cairo_font_options_t)(v.Ptr)
}
func wrapFontOptions(p *C.cairo_font_options_t) FontOptions {
	return FontOptions{unsafe.Pointer(p)}
}
func WrapFontOptions(p unsafe.Pointer) FontOptions {
	return FontOptions{p}
}
func (v FontOptions) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontOptions(p unsafe.Pointer) interface{} {
	return WrapFontOptions(p)
}

// Struct FontType
type FontType struct {
	Ptr unsafe.Pointer
}

func (v FontType) native() *C.cairo_font_type_t {
	return (*C.cairo_font_type_t)(v.Ptr)
}
func wrapFontType(p *C.cairo_font_type_t) FontType {
	return FontType{unsafe.Pointer(p)}
}
func WrapFontType(p unsafe.Pointer) FontType {
	return FontType{p}
}
func (v FontType) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontType(p unsafe.Pointer) interface{} {
	return WrapFontType(p)
}

// Struct FontFace
type FontFace struct {
	Ptr unsafe.Pointer
}

func (v FontFace) native() *C.cairo_font_face_t {
	return (*C.cairo_font_face_t)(v.Ptr)
}
func wrapFontFace(p *C.cairo_font_face_t) FontFace {
	return FontFace{unsafe.Pointer(p)}
}
func WrapFontFace(p unsafe.Pointer) FontFace {
	return FontFace{p}
}
func (v FontFace) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontFace(p unsafe.Pointer) interface{} {
	return WrapFontFace(p)
}

// Struct ScaledFont
type ScaledFont struct {
	Ptr unsafe.Pointer
}

func (v ScaledFont) native() *C.cairo_scaled_font_t {
	return (*C.cairo_scaled_font_t)(v.Ptr)
}
func wrapScaledFont(p *C.cairo_scaled_font_t) ScaledFont {
	return ScaledFont{unsafe.Pointer(p)}
}
func WrapScaledFont(p unsafe.Pointer) ScaledFont {
	return ScaledFont{p}
}
func (v ScaledFont) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScaledFont(p unsafe.Pointer) interface{} {
	return WrapScaledFont(p)
}

// Struct Path
type Path struct {
	Ptr unsafe.Pointer
}

func (v Path) native() *C.cairo_path_t {
	return (*C.cairo_path_t)(v.Ptr)
}
func wrapPath(p *C.cairo_path_t) Path {
	return Path{unsafe.Pointer(p)}
}
func WrapPath(p unsafe.Pointer) Path {
	return Path{p}
}
func (v Path) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPath(p unsafe.Pointer) interface{} {
	return WrapPath(p)
}

// Struct RectangleInt
type RectangleInt struct {
	Ptr unsafe.Pointer
}

func (v RectangleInt) native() *C.cairo_rectangle_int_t {
	return (*C.cairo_rectangle_int_t)(v.Ptr)
}
func wrapRectangleInt(p *C.cairo_rectangle_int_t) RectangleInt {
	return RectangleInt{unsafe.Pointer(p)}
}
func WrapRectangleInt(p unsafe.Pointer) RectangleInt {
	return RectangleInt{p}
}
func (v RectangleInt) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRectangleInt(p unsafe.Pointer) interface{} {
	return WrapRectangleInt(p)
}

type Content int

const (
	ContentColor      Content = 4096
	ContentAlpha              = 8192
	ContentColorAlpha         = 12288
)
