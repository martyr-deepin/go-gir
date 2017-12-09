package gdkpixbuf

/*
#cgo pkg-config: gdk-pixbuf-2.0
#include <gdk-pixbuf/gdk-pixbuf.h>
#include <stdlib.h>
*/
import "C"
import "github.com/linuxdeepin/go-gir/gio-2.0"
import "github.com/linuxdeepin/go-gir/glib-2.0"
import "github.com/linuxdeepin/go-gir/gobject-2.0"
import "github.com/linuxdeepin/go-gir/util"
import "unsafe"

// Struct PixbufFormat
type PixbufFormat struct {
	Ptr unsafe.Pointer
}

func (v PixbufFormat) native() *C.GdkPixbufFormat {
	return (*C.GdkPixbufFormat)(v.Ptr)
}
func wrapPixbufFormat(p *C.GdkPixbufFormat) PixbufFormat {
	return PixbufFormat{Ptr: unsafe.Pointer(p)}
}
func WrapPixbufFormat(p unsafe.Pointer) PixbufFormat {
	return PixbufFormat{Ptr: p}
}
func (v PixbufFormat) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPixbufFormat(p unsafe.Pointer) interface{} {
	return WrapPixbufFormat(p)
}

// Copy is a wrapper around gdk_pixbuf_format_copy().
func (format PixbufFormat) Copy() PixbufFormat {
	ret0 := C.gdk_pixbuf_format_copy(format.native())
	return wrapPixbufFormat(ret0)
}

// Free is a wrapper around gdk_pixbuf_format_free().
func (format PixbufFormat) Free() {
	C.gdk_pixbuf_format_free(format.native())
}

// GetDescription is a wrapper around gdk_pixbuf_format_get_description().
func (format PixbufFormat) GetDescription() string {
	ret0 := C.gdk_pixbuf_format_get_description(format.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetExtensions is a wrapper around gdk_pixbuf_format_get_extensions().
func (format PixbufFormat) GetExtensions() []string {
	ret0 := C.gdk_pixbuf_format_get_extensions(format.native())
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
	return ret
}

// GetLicense is a wrapper around gdk_pixbuf_format_get_license().
func (format PixbufFormat) GetLicense() string {
	ret0 := C.gdk_pixbuf_format_get_license(format.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetMimeTypes is a wrapper around gdk_pixbuf_format_get_mime_types().
func (format PixbufFormat) GetMimeTypes() []string {
	ret0 := C.gdk_pixbuf_format_get_mime_types(format.native())
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
	return ret
}

// GetName is a wrapper around gdk_pixbuf_format_get_name().
func (format PixbufFormat) GetName() string {
	ret0 := C.gdk_pixbuf_format_get_name(format.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// IsDisabled is a wrapper around gdk_pixbuf_format_is_disabled().
func (format PixbufFormat) IsDisabled() bool {
	ret0 := C.gdk_pixbuf_format_is_disabled(format.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsSaveOptionSupported is a wrapper around gdk_pixbuf_format_is_save_option_supported().
func (format PixbufFormat) IsSaveOptionSupported(option_key string) bool {
	option_key0 := (*C.gchar)(C.CString(option_key))
	ret0 := C.gdk_pixbuf_format_is_save_option_supported(format.native(), option_key0)
	C.free(unsafe.Pointer(option_key0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// IsScalable is a wrapper around gdk_pixbuf_format_is_scalable().
func (format PixbufFormat) IsScalable() bool {
	ret0 := C.gdk_pixbuf_format_is_scalable(format.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsWritable is a wrapper around gdk_pixbuf_format_is_writable().
func (format PixbufFormat) IsWritable() bool {
	ret0 := C.gdk_pixbuf_format_is_writable(format.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetDisabled is a wrapper around gdk_pixbuf_format_set_disabled().
func (format PixbufFormat) SetDisabled(disabled bool) {
	C.gdk_pixbuf_format_set_disabled(format.native(), C.gboolean(util.Bool2Int(disabled)) /*go:.util*/)
}

// Object Pixbuf
type Pixbuf struct {
	gio.IconIface
	gio.LoadableIconIface
	gobject.Object
}

func (v Pixbuf) native() *C.GdkPixbuf {
	return (*C.GdkPixbuf)(v.Ptr)
}
func wrapPixbuf(p *C.GdkPixbuf) (v Pixbuf) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPixbuf(p unsafe.Pointer) (v Pixbuf) {
	v.Ptr = p
	return
}
func (v Pixbuf) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPixbuf(p unsafe.Pointer) interface{} {
	return WrapPixbuf(p)
}
func (v Pixbuf) GetType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_get_type())
}
func (v Pixbuf) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPixbuf(unsafe.Pointer(ptr)), nil
	}
}
func (v Pixbuf) Icon() gio.Icon {
	return gio.WrapIcon(v.Ptr)
}
func (v Pixbuf) LoadableIcon() gio.LoadableIcon {
	return gio.WrapLoadableIcon(v.Ptr)
}

// PixbufNew is a wrapper around gdk_pixbuf_new().
func PixbufNew(colorspace Colorspace, has_alpha bool, bits_per_sample int, width int, height int) Pixbuf {
	ret0 := C.gdk_pixbuf_new(C.GdkColorspace(colorspace), C.gboolean(util.Bool2Int(has_alpha)) /*go:.util*/, C.int(bits_per_sample), C.int(width), C.int(height))
	return wrapPixbuf(ret0)
}

// PixbufNewFromFile is a wrapper around gdk_pixbuf_new_from_file().
func PixbufNewFromFile(filename string) (Pixbuf, error) {
	filename0 := C.CString(filename)
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_file(filename0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromFileAtScale is a wrapper around gdk_pixbuf_new_from_file_at_scale().
func PixbufNewFromFileAtScale(filename string, width int, height int, preserve_aspect_ratio bool) (Pixbuf, error) {
	filename0 := C.CString(filename)
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_file_at_scale(filename0, C.int(width), C.int(height), C.gboolean(util.Bool2Int(preserve_aspect_ratio)) /*go:.util*/, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromFileAtSize is a wrapper around gdk_pixbuf_new_from_file_at_size().
func PixbufNewFromFileAtSize(filename string, width int, height int) (Pixbuf, error) {
	filename0 := C.CString(filename)
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_file_at_size(filename0, C.int(width), C.int(height), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromResource is a wrapper around gdk_pixbuf_new_from_resource().
func PixbufNewFromResource(resource_path string) (Pixbuf, error) {
	resource_path0 := C.CString(resource_path)
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_resource(resource_path0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromResourceAtScale is a wrapper around gdk_pixbuf_new_from_resource_at_scale().
func PixbufNewFromResourceAtScale(resource_path string, width int, height int, preserve_aspect_ratio bool) (Pixbuf, error) {
	resource_path0 := C.CString(resource_path)
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_resource_at_scale(resource_path0, C.int(width), C.int(height), C.gboolean(util.Bool2Int(preserve_aspect_ratio)) /*go:.util*/, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromStream is a wrapper around gdk_pixbuf_new_from_stream().
func PixbufNewFromStream(stream gio.InputStream, cancellable gio.Cancellable) (Pixbuf, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_stream((*C.GInputStream)(stream.Ptr), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromStreamAtScale is a wrapper around gdk_pixbuf_new_from_stream_at_scale().
func PixbufNewFromStreamAtScale(stream gio.InputStream, width int, height int, preserve_aspect_ratio bool, cancellable gio.Cancellable) (Pixbuf, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_stream_at_scale((*C.GInputStream)(stream.Ptr), C.gint(width), C.gint(height), C.gboolean(util.Bool2Int(preserve_aspect_ratio)) /*go:.util*/, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromStreamFinish is a wrapper around gdk_pixbuf_new_from_stream_finish().
func PixbufNewFromStreamFinish(async_result gio.AsyncResult) (Pixbuf, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_new_from_stream_finish((*C.GAsyncResult)(async_result.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Pixbuf{}, err.GoValue()
	}
	return wrapPixbuf(ret0), nil
}

// PixbufNewFromXpmData is a wrapper around gdk_pixbuf_new_from_xpm_data().
func PixbufNewFromXpmData(data []string) Pixbuf {
	data0 := make([]*C.char, len(data))
	for idx, elemG := range data {
		elem := C.CString(elemG)
		data0[idx] = elem
	}
	var data0Ptr **C.char
	if len(data0) > 0 {
		data0Ptr = &data0[0]
	}
	ret0 := C.gdk_pixbuf_new_from_xpm_data(data0Ptr)
	for _, elem := range data0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return wrapPixbuf(ret0)
}

// AddAlpha is a wrapper around gdk_pixbuf_add_alpha().
func (pixbuf Pixbuf) AddAlpha(substitute_color bool, r byte, g byte, b byte) Pixbuf {
	ret0 := C.gdk_pixbuf_add_alpha(pixbuf.native(), C.gboolean(util.Bool2Int(substitute_color)) /*go:.util*/, C.guchar(r), C.guchar(g), C.guchar(b))
	return wrapPixbuf(ret0)
}

// ApplyEmbeddedOrientation is a wrapper around gdk_pixbuf_apply_embedded_orientation().
func (src Pixbuf) ApplyEmbeddedOrientation() Pixbuf {
	ret0 := C.gdk_pixbuf_apply_embedded_orientation(src.native())
	return wrapPixbuf(ret0)
}

// Composite is a wrapper around gdk_pixbuf_composite().
func (src Pixbuf) Composite(dest Pixbuf, dest_x int, dest_y int, dest_width int, dest_height int, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type InterpType, overall_alpha int) {
	C.gdk_pixbuf_composite(src.native(), dest.native(), C.int(dest_x), C.int(dest_y), C.int(dest_width), C.int(dest_height), C.double(offset_x), C.double(offset_y), C.double(scale_x), C.double(scale_y), C.GdkInterpType(interp_type), C.int(overall_alpha))
}

// CompositeColor is a wrapper around gdk_pixbuf_composite_color().
func (src Pixbuf) CompositeColor(dest Pixbuf, dest_x int, dest_y int, dest_width int, dest_height int, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type InterpType, overall_alpha int, check_x int, check_y int, check_size int, color1 uint32, color2 uint32) {
	C.gdk_pixbuf_composite_color(src.native(), dest.native(), C.int(dest_x), C.int(dest_y), C.int(dest_width), C.int(dest_height), C.double(offset_x), C.double(offset_y), C.double(scale_x), C.double(scale_y), C.GdkInterpType(interp_type), C.int(overall_alpha), C.int(check_x), C.int(check_y), C.int(check_size), C.guint32(color1), C.guint32(color2))
}

// CompositeColorSimple is a wrapper around gdk_pixbuf_composite_color_simple().
func (src Pixbuf) CompositeColorSimple(dest_width int, dest_height int, interp_type InterpType, overall_alpha int, check_size int, color1 uint32, color2 uint32) Pixbuf {
	ret0 := C.gdk_pixbuf_composite_color_simple(src.native(), C.int(dest_width), C.int(dest_height), C.GdkInterpType(interp_type), C.int(overall_alpha), C.int(check_size), C.guint32(color1), C.guint32(color2))
	return wrapPixbuf(ret0)
}

// Copy is a wrapper around gdk_pixbuf_copy().
func (pixbuf Pixbuf) Copy() Pixbuf {
	ret0 := C.gdk_pixbuf_copy(pixbuf.native())
	return wrapPixbuf(ret0)
}

// CopyArea is a wrapper around gdk_pixbuf_copy_area().
func (src_pixbuf Pixbuf) CopyArea(src_x int, src_y int, width int, height int, dest_pixbuf Pixbuf, dest_x int, dest_y int) {
	C.gdk_pixbuf_copy_area(src_pixbuf.native(), C.int(src_x), C.int(src_y), C.int(width), C.int(height), dest_pixbuf.native(), C.int(dest_x), C.int(dest_y))
}

// CopyOptions is a wrapper around gdk_pixbuf_copy_options().
func (src_pixbuf Pixbuf) CopyOptions(dest_pixbuf Pixbuf) bool {
	ret0 := C.gdk_pixbuf_copy_options(src_pixbuf.native(), dest_pixbuf.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Fill is a wrapper around gdk_pixbuf_fill().
func (pixbuf Pixbuf) Fill(pixel uint32) {
	C.gdk_pixbuf_fill(pixbuf.native(), C.guint32(pixel))
}

// Flip is a wrapper around gdk_pixbuf_flip().
func (src Pixbuf) Flip(horizontal bool) Pixbuf {
	ret0 := C.gdk_pixbuf_flip(src.native(), C.gboolean(util.Bool2Int(horizontal)) /*go:.util*/)
	return wrapPixbuf(ret0)
}

// GetBitsPerSample is a wrapper around gdk_pixbuf_get_bits_per_sample().
func (pixbuf Pixbuf) GetBitsPerSample() int {
	ret0 := C.gdk_pixbuf_get_bits_per_sample(pixbuf.native())
	return int(ret0)
}

// GetByteLength is a wrapper around gdk_pixbuf_get_byte_length().
func (pixbuf Pixbuf) GetByteLength() uint {
	ret0 := C.gdk_pixbuf_get_byte_length(pixbuf.native())
	return uint(ret0)
}

// GetColorspace is a wrapper around gdk_pixbuf_get_colorspace().
func (pixbuf Pixbuf) GetColorspace() Colorspace {
	ret0 := C.gdk_pixbuf_get_colorspace(pixbuf.native())
	return Colorspace(ret0)
}

// GetHasAlpha is a wrapper around gdk_pixbuf_get_has_alpha().
func (pixbuf Pixbuf) GetHasAlpha() bool {
	ret0 := C.gdk_pixbuf_get_has_alpha(pixbuf.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHeight is a wrapper around gdk_pixbuf_get_height().
func (pixbuf Pixbuf) GetHeight() int {
	ret0 := C.gdk_pixbuf_get_height(pixbuf.native())
	return int(ret0)
}

// GetNChannels is a wrapper around gdk_pixbuf_get_n_channels().
func (pixbuf Pixbuf) GetNChannels() int {
	ret0 := C.gdk_pixbuf_get_n_channels(pixbuf.native())
	return int(ret0)
}

// GetOption is a wrapper around gdk_pixbuf_get_option().
func (pixbuf Pixbuf) GetOption(key string) string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gdk_pixbuf_get_option(pixbuf.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// gdk_pixbuf_get_pixels shadowed by get_pixels_with_length

// GetPixels is a wrapper around gdk_pixbuf_get_pixels_with_length().
func (pixbuf Pixbuf) GetPixels() []byte {
	var length0 C.guint
	ret0 := C.gdk_pixbuf_get_pixels_with_length(pixbuf.native(), &length0)
	var ret0Slice []C.guchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]byte, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = byte(elem)
	}
	return ret
}

// GetRowstride is a wrapper around gdk_pixbuf_get_rowstride().
func (pixbuf Pixbuf) GetRowstride() int {
	ret0 := C.gdk_pixbuf_get_rowstride(pixbuf.native())
	return int(ret0)
}

// GetWidth is a wrapper around gdk_pixbuf_get_width().
func (pixbuf Pixbuf) GetWidth() int {
	ret0 := C.gdk_pixbuf_get_width(pixbuf.native())
	return int(ret0)
}

// NewSubpixbuf is a wrapper around gdk_pixbuf_new_subpixbuf().
func (src_pixbuf Pixbuf) NewSubpixbuf(src_x int, src_y int, width int, height int) Pixbuf {
	ret0 := C.gdk_pixbuf_new_subpixbuf(src_pixbuf.native(), C.int(src_x), C.int(src_y), C.int(width), C.int(height))
	return wrapPixbuf(ret0)
}

// ReadPixelBytes is a wrapper around gdk_pixbuf_read_pixel_bytes().
func (pixbuf Pixbuf) ReadPixelBytes() glib.Bytes {
	ret0 := C.gdk_pixbuf_read_pixel_bytes(pixbuf.native())
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// RemoveOption is a wrapper around gdk_pixbuf_remove_option().
func (pixbuf Pixbuf) RemoveOption(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gdk_pixbuf_remove_option(pixbuf.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func (src Pixbuf) RotateSimple(angle PixbufRotation) Pixbuf {
	ret0 := C.gdk_pixbuf_rotate_simple(src.native(), C.GdkPixbufRotation(angle))
	return wrapPixbuf(ret0)
}

// SaturateAndPixelate is a wrapper around gdk_pixbuf_saturate_and_pixelate().
func (src Pixbuf) SaturateAndPixelate(dest Pixbuf, saturation float32, pixelate bool) {
	C.gdk_pixbuf_saturate_and_pixelate(src.native(), dest.native(), C.gfloat(saturation), C.gboolean(util.Bool2Int(pixelate)) /*go:.util*/)
}

// SaveToStreamv is a wrapper around gdk_pixbuf_save_to_streamv().
func (pixbuf Pixbuf) SaveToStreamv(stream gio.OutputStream, type_ string, option_keys []string, option_values []string, cancellable gio.Cancellable) (bool, error) {
	type0 := C.CString(type_)
	option_keys0 := make([]*C.char, len(option_keys))
	for idx, elemG := range option_keys {
		elem := C.CString(elemG)
		option_keys0[idx] = elem
	}
	var option_keys0Ptr **C.char
	if len(option_keys0) > 0 {
		option_keys0Ptr = &option_keys0[0]
	}
	option_values0 := make([]*C.char, len(option_values))
	for idx, elemG := range option_values {
		elem := C.CString(elemG)
		option_values0[idx] = elem
	}
	var option_values0Ptr **C.char
	if len(option_values0) > 0 {
		option_values0Ptr = &option_values0[0]
	}
	var err glib.Error
	ret0 := C.gdk_pixbuf_save_to_streamv(pixbuf.native(), (*C.GOutputStream)(stream.Ptr), type0, option_keys0Ptr, option_values0Ptr, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(type0)) /*ch:<stdlib.h>*/
	for _, elem := range option_keys0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	for _, elem := range option_values0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Savev is a wrapper around gdk_pixbuf_savev().
func (pixbuf Pixbuf) Savev(filename string, type_ string, option_keys []string, option_values []string) (bool, error) {
	filename0 := C.CString(filename)
	type0 := C.CString(type_)
	option_keys0 := make([]*C.char, len(option_keys))
	for idx, elemG := range option_keys {
		elem := C.CString(elemG)
		option_keys0[idx] = elem
	}
	var option_keys0Ptr **C.char
	if len(option_keys0) > 0 {
		option_keys0Ptr = &option_keys0[0]
	}
	option_values0 := make([]*C.char, len(option_values))
	for idx, elemG := range option_values {
		elem := C.CString(elemG)
		option_values0[idx] = elem
	}
	var option_values0Ptr **C.char
	if len(option_values0) > 0 {
		option_values0Ptr = &option_values0[0]
	}
	var err glib.Error
	ret0 := C.gdk_pixbuf_savev(pixbuf.native(), filename0, type0, option_keys0Ptr, option_values0Ptr, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(type0))     /*ch:<stdlib.h>*/
	for _, elem := range option_keys0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	for _, elem := range option_values0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Scale is a wrapper around gdk_pixbuf_scale().
func (src Pixbuf) Scale(dest Pixbuf, dest_x int, dest_y int, dest_width int, dest_height int, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type InterpType) {
	C.gdk_pixbuf_scale(src.native(), dest.native(), C.int(dest_x), C.int(dest_y), C.int(dest_width), C.int(dest_height), C.double(offset_x), C.double(offset_y), C.double(scale_x), C.double(scale_y), C.GdkInterpType(interp_type))
}

// ScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func (src Pixbuf) ScaleSimple(dest_width int, dest_height int, interp_type InterpType) Pixbuf {
	ret0 := C.gdk_pixbuf_scale_simple(src.native(), C.int(dest_width), C.int(dest_height), C.GdkInterpType(interp_type))
	return wrapPixbuf(ret0)
}

// SetOption is a wrapper around gdk_pixbuf_set_option().
func (pixbuf Pixbuf) SetOption(key string, value string) bool {
	key0 := (*C.gchar)(C.CString(key))
	value0 := (*C.gchar)(C.CString(value))
	ret0 := C.gdk_pixbuf_set_option(pixbuf.native(), key0, value0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))  /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PixbufGetFileInfo is a wrapper around gdk_pixbuf_get_file_info().
func PixbufGetFileInfo(filename string) (PixbufFormat, int, int) {
	filename0 := (*C.gchar)(C.CString(filename))
	var width0 C.gint
	var height0 C.gint
	ret0 := C.gdk_pixbuf_get_file_info(filename0, &width0, &height0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	return wrapPixbufFormat(ret0), int(width0), int(height0)
}

// PixbufGetFileInfoFinish is a wrapper around gdk_pixbuf_get_file_info_finish().
func PixbufGetFileInfoFinish(async_result gio.AsyncResult) (PixbufFormat, int, int, error) {
	var width0 C.gint
	var height0 C.gint
	var err glib.Error
	ret0 := C.gdk_pixbuf_get_file_info_finish((*C.GAsyncResult)(async_result.Ptr), &width0, &height0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return PixbufFormat{}, 0, 0, err.GoValue()
	}
	return wrapPixbufFormat(ret0), int(width0), int(height0), nil
}

// PixbufSaveToStreamFinish is a wrapper around gdk_pixbuf_save_to_stream_finish().
func PixbufSaveToStreamFinish(async_result gio.AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_save_to_stream_finish((*C.GAsyncResult)(async_result.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object PixbufAnimation
type PixbufAnimation struct {
	gobject.Object
}

func (v PixbufAnimation) native() *C.GdkPixbufAnimation {
	return (*C.GdkPixbufAnimation)(v.Ptr)
}
func wrapPixbufAnimation(p *C.GdkPixbufAnimation) (v PixbufAnimation) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPixbufAnimation(p unsafe.Pointer) (v PixbufAnimation) {
	v.Ptr = p
	return
}
func (v PixbufAnimation) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPixbufAnimation(p unsafe.Pointer) interface{} {
	return WrapPixbufAnimation(p)
}
func (v PixbufAnimation) GetType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_animation_get_type())
}
func (v PixbufAnimation) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPixbufAnimation(unsafe.Pointer(ptr)), nil
	}
}

// PixbufAnimationNewFromFile is a wrapper around gdk_pixbuf_animation_new_from_file().
func PixbufAnimationNewFromFile(filename string) (PixbufAnimation, error) {
	filename0 := C.CString(filename)
	var err glib.Error
	ret0 := C.gdk_pixbuf_animation_new_from_file(filename0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PixbufAnimation{}, err.GoValue()
	}
	return wrapPixbufAnimation(ret0), nil
}

// PixbufAnimationNewFromResource is a wrapper around gdk_pixbuf_animation_new_from_resource().
func PixbufAnimationNewFromResource(resource_path string) (PixbufAnimation, error) {
	resource_path0 := C.CString(resource_path)
	var err glib.Error
	ret0 := C.gdk_pixbuf_animation_new_from_resource(resource_path0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PixbufAnimation{}, err.GoValue()
	}
	return wrapPixbufAnimation(ret0), nil
}

// PixbufAnimationNewFromStream is a wrapper around gdk_pixbuf_animation_new_from_stream().
func PixbufAnimationNewFromStream(stream gio.InputStream, cancellable gio.Cancellable) (PixbufAnimation, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_animation_new_from_stream((*C.GInputStream)(stream.Ptr), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return PixbufAnimation{}, err.GoValue()
	}
	return wrapPixbufAnimation(ret0), nil
}

// PixbufAnimationNewFromStreamFinish is a wrapper around gdk_pixbuf_animation_new_from_stream_finish().
func PixbufAnimationNewFromStreamFinish(async_result gio.AsyncResult) (PixbufAnimation, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_animation_new_from_stream_finish((*C.GAsyncResult)(async_result.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return PixbufAnimation{}, err.GoValue()
	}
	return wrapPixbufAnimation(ret0), nil
}

// GetHeight is a wrapper around gdk_pixbuf_animation_get_height().
func (animation PixbufAnimation) GetHeight() int {
	ret0 := C.gdk_pixbuf_animation_get_height(animation.native())
	return int(ret0)
}

// GetIter is a wrapper around gdk_pixbuf_animation_get_iter().
func (animation PixbufAnimation) GetIter(start_time glib.TimeVal) PixbufAnimationIter {
	ret0 := C.gdk_pixbuf_animation_get_iter(animation.native(), (*C.GTimeVal)(start_time.Ptr))
	return wrapPixbufAnimationIter(ret0)
}

// GetStaticImage is a wrapper around gdk_pixbuf_animation_get_static_image().
func (animation PixbufAnimation) GetStaticImage() Pixbuf {
	ret0 := C.gdk_pixbuf_animation_get_static_image(animation.native())
	return wrapPixbuf(ret0)
}

// GetWidth is a wrapper around gdk_pixbuf_animation_get_width().
func (animation PixbufAnimation) GetWidth() int {
	ret0 := C.gdk_pixbuf_animation_get_width(animation.native())
	return int(ret0)
}

// IsStaticImage is a wrapper around gdk_pixbuf_animation_is_static_image().
func (animation PixbufAnimation) IsStaticImage() bool {
	ret0 := C.gdk_pixbuf_animation_is_static_image(animation.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object PixbufAnimationIter
type PixbufAnimationIter struct {
	gobject.Object
}

func (v PixbufAnimationIter) native() *C.GdkPixbufAnimationIter {
	return (*C.GdkPixbufAnimationIter)(v.Ptr)
}
func wrapPixbufAnimationIter(p *C.GdkPixbufAnimationIter) (v PixbufAnimationIter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPixbufAnimationIter(p unsafe.Pointer) (v PixbufAnimationIter) {
	v.Ptr = p
	return
}
func (v PixbufAnimationIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPixbufAnimationIter(p unsafe.Pointer) interface{} {
	return WrapPixbufAnimationIter(p)
}
func (v PixbufAnimationIter) GetType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_animation_iter_get_type())
}
func (v PixbufAnimationIter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPixbufAnimationIter(unsafe.Pointer(ptr)), nil
	}
}

// Advance is a wrapper around gdk_pixbuf_animation_iter_advance().
func (iter PixbufAnimationIter) Advance(current_time glib.TimeVal) bool {
	ret0 := C.gdk_pixbuf_animation_iter_advance(iter.native(), (*C.GTimeVal)(current_time.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDelayTime is a wrapper around gdk_pixbuf_animation_iter_get_delay_time().
func (iter PixbufAnimationIter) GetDelayTime() int {
	ret0 := C.gdk_pixbuf_animation_iter_get_delay_time(iter.native())
	return int(ret0)
}

// GetPixbuf is a wrapper around gdk_pixbuf_animation_iter_get_pixbuf().
func (iter PixbufAnimationIter) GetPixbuf() Pixbuf {
	ret0 := C.gdk_pixbuf_animation_iter_get_pixbuf(iter.native())
	return wrapPixbuf(ret0)
}

// OnCurrentlyLoadingFrame is a wrapper around gdk_pixbuf_animation_iter_on_currently_loading_frame().
func (iter PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {
	ret0 := C.gdk_pixbuf_animation_iter_on_currently_loading_frame(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object PixbufLoader
type PixbufLoader struct {
	gobject.Object
}

func (v PixbufLoader) native() *C.GdkPixbufLoader {
	return (*C.GdkPixbufLoader)(v.Ptr)
}
func wrapPixbufLoader(p *C.GdkPixbufLoader) (v PixbufLoader) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPixbufLoader(p unsafe.Pointer) (v PixbufLoader) {
	v.Ptr = p
	return
}
func (v PixbufLoader) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPixbufLoader(p unsafe.Pointer) interface{} {
	return WrapPixbufLoader(p)
}
func (v PixbufLoader) GetType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_loader_get_type())
}
func (v PixbufLoader) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPixbufLoader(unsafe.Pointer(ptr)), nil
	}
}

// PixbufLoaderNew is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() PixbufLoader {
	ret0 := C.gdk_pixbuf_loader_new()
	return wrapPixbufLoader(ret0)
}

// PixbufLoaderNewWithMimeType is a wrapper around gdk_pixbuf_loader_new_with_mime_type().
func PixbufLoaderNewWithMimeType(mime_type string) (PixbufLoader, error) {
	mime_type0 := C.CString(mime_type)
	var err glib.Error
	ret0 := C.gdk_pixbuf_loader_new_with_mime_type(mime_type0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(mime_type0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PixbufLoader{}, err.GoValue()
	}
	return wrapPixbufLoader(ret0), nil
}

// PixbufLoaderNewWithType is a wrapper around gdk_pixbuf_loader_new_with_type().
func PixbufLoaderNewWithType(image_type string) (PixbufLoader, error) {
	image_type0 := C.CString(image_type)
	var err glib.Error
	ret0 := C.gdk_pixbuf_loader_new_with_type(image_type0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(image_type0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PixbufLoader{}, err.GoValue()
	}
	return wrapPixbufLoader(ret0), nil
}

// Close is a wrapper around gdk_pixbuf_loader_close().
func (loader PixbufLoader) Close() (bool, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_loader_close(loader.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetAnimation is a wrapper around gdk_pixbuf_loader_get_animation().
func (loader PixbufLoader) GetAnimation() PixbufAnimation {
	ret0 := C.gdk_pixbuf_loader_get_animation(loader.native())
	return wrapPixbufAnimation(ret0)
}

// GetFormat is a wrapper around gdk_pixbuf_loader_get_format().
func (loader PixbufLoader) GetFormat() PixbufFormat {
	ret0 := C.gdk_pixbuf_loader_get_format(loader.native())
	return wrapPixbufFormat(ret0)
}

// GetPixbuf is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (loader PixbufLoader) GetPixbuf() Pixbuf {
	ret0 := C.gdk_pixbuf_loader_get_pixbuf(loader.native())
	return wrapPixbuf(ret0)
}

// SetSize is a wrapper around gdk_pixbuf_loader_set_size().
func (loader PixbufLoader) SetSize(width int, height int) {
	C.gdk_pixbuf_loader_set_size(loader.native(), C.int(width), C.int(height))
}

// Write is a wrapper around gdk_pixbuf_loader_write().
func (loader PixbufLoader) Write(buf []byte) (bool, error) {
	buf0 := make([]C.guchar, len(buf))
	for idx, elemG := range buf {
		buf0[idx] = C.guchar(elemG)
	}
	var buf0Ptr *C.guchar
	if len(buf0) > 0 {
		buf0Ptr = &buf0[0]
	}
	var err glib.Error
	ret0 := C.gdk_pixbuf_loader_write(loader.native(), buf0Ptr, C.gsize(len(buf)), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// WriteBytes is a wrapper around gdk_pixbuf_loader_write_bytes().
func (loader PixbufLoader) WriteBytes(buffer glib.Bytes) (bool, error) {
	var err glib.Error
	ret0 := C.gdk_pixbuf_loader_write_bytes(loader.native(), (*C.GBytes)(buffer.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object PixbufSimpleAnim
type PixbufSimpleAnim struct {
	PixbufAnimation
}

func (v PixbufSimpleAnim) native() *C.GdkPixbufSimpleAnim {
	return (*C.GdkPixbufSimpleAnim)(v.Ptr)
}
func wrapPixbufSimpleAnim(p *C.GdkPixbufSimpleAnim) (v PixbufSimpleAnim) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPixbufSimpleAnim(p unsafe.Pointer) (v PixbufSimpleAnim) {
	v.Ptr = p
	return
}
func (v PixbufSimpleAnim) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPixbufSimpleAnim(p unsafe.Pointer) interface{} {
	return WrapPixbufSimpleAnim(p)
}
func (v PixbufSimpleAnim) GetType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_simple_anim_get_type())
}
func (v PixbufSimpleAnim) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPixbufSimpleAnim(unsafe.Pointer(ptr)), nil
	}
}

// PixbufSimpleAnimNew is a wrapper around gdk_pixbuf_simple_anim_new().
func PixbufSimpleAnimNew(width int, height int, rate float32) PixbufSimpleAnim {
	ret0 := C.gdk_pixbuf_simple_anim_new(C.gint(width), C.gint(height), C.gfloat(rate))
	return wrapPixbufSimpleAnim(ret0)
}

// AddFrame is a wrapper around gdk_pixbuf_simple_anim_add_frame().
func (animation PixbufSimpleAnim) AddFrame(pixbuf Pixbuf) {
	C.gdk_pixbuf_simple_anim_add_frame(animation.native(), pixbuf.native())
}

// GetLoop is a wrapper around gdk_pixbuf_simple_anim_get_loop().
func (animation PixbufSimpleAnim) GetLoop() bool {
	ret0 := C.gdk_pixbuf_simple_anim_get_loop(animation.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetLoop is a wrapper around gdk_pixbuf_simple_anim_set_loop().
func (animation PixbufSimpleAnim) SetLoop(loop bool) {
	C.gdk_pixbuf_simple_anim_set_loop(animation.native(), C.gboolean(util.Bool2Int(loop)) /*go:.util*/)
}

type Colorspace int

const (
	ColorspaceRgb Colorspace = 0
)

type InterpType int

const (
	InterpTypeNearest  InterpType = 0
	InterpTypeTiles               = 1
	InterpTypeBilinear            = 2
	InterpTypeHyper               = 3
)

type PixbufAlphaMode int

const (
	PixbufAlphaModeBilevel PixbufAlphaMode = 0
	PixbufAlphaModeFull                    = 1
)

type PixbufError int

const (
	PixbufErrorCorruptImage         PixbufError = 0
	PixbufErrorInsufficientMemory               = 1
	PixbufErrorBadOption                        = 2
	PixbufErrorUnknownType                      = 3
	PixbufErrorUnsupportedOperation             = 4
	PixbufErrorFailed                           = 5
	PixbufErrorIncompleteAnimation              = 6
)

type PixbufRotation int

const (
	PixbufRotationNone             PixbufRotation = 0
	PixbufRotationCounterclockwise                = 90
	PixbufRotationUpsidedown                      = 180
	PixbufRotationClockwise                       = 270
)

type PixdataDumpType int

const (
	PixdataDumpTypePixdataStream PixdataDumpType = 0
	PixdataDumpTypePixdataStruct                 = 1
	PixdataDumpTypeMacros                        = 2
	PixdataDumpTypeGtypes                        = 0
	PixdataDumpTypeCtypes                        = 256
	PixdataDumpTypeStatic                        = 512
	PixdataDumpTypeConst                         = 1024
	PixdataDumpTypeRleDecoder                    = 65536
)

type PixdataType int

const (
	PixdataTypeColorTypeRgb    PixdataType = 1
	PixdataTypeColorTypeRgba               = 2
	PixdataTypeColorTypeMask               = 255
	PixdataTypeSampleWidth8                = 65536
	PixdataTypeSampleWidthMask             = 983040
	PixdataTypeEncodingRaw                 = 16777216
	PixdataTypeEncodingRle                 = 33554432
	PixdataTypeEncodingMask                = 251658240
)
