package gtk

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk-a11y.h>
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>
*/
import "C"
import "github.com/linuxdeepin/go-gir/atk-1.0"
import "github.com/linuxdeepin/go-gir/cairo-1.0"
import "github.com/linuxdeepin/go-gir/gdk-3.0"
import "github.com/linuxdeepin/go-gir/gdkpixbuf-2.0"
import "github.com/linuxdeepin/go-gir/gio-2.0"
import "github.com/linuxdeepin/go-gir/glib-2.0"
import "github.com/linuxdeepin/go-gir/gobject-2.0"
import "github.com/linuxdeepin/go-gir/pango-1.0"
import "github.com/linuxdeepin/go-gir/util"
import "unsafe"

// Struct AccelGroupEntry
type AccelGroupEntry struct {
	Ptr unsafe.Pointer
}

func (v AccelGroupEntry) native() *C.GtkAccelGroupEntry {
	return (*C.GtkAccelGroupEntry)(v.Ptr)
}
func wrapAccelGroupEntry(p *C.GtkAccelGroupEntry) AccelGroupEntry {
	return AccelGroupEntry{unsafe.Pointer(p)}
}
func WrapAccelGroupEntry(p unsafe.Pointer) AccelGroupEntry {
	return AccelGroupEntry{p}
}
func (v AccelGroupEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAccelGroupEntry(p unsafe.Pointer) interface{} {
	return WrapAccelGroupEntry(p)
}

// Struct AccelKey
type AccelKey struct {
	Ptr unsafe.Pointer
}

func (v AccelKey) native() *C.GtkAccelKey {
	return (*C.GtkAccelKey)(v.Ptr)
}
func wrapAccelKey(p *C.GtkAccelKey) AccelKey {
	return AccelKey{unsafe.Pointer(p)}
}
func WrapAccelKey(p unsafe.Pointer) AccelKey {
	return AccelKey{p}
}
func (v AccelKey) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAccelKey(p unsafe.Pointer) interface{} {
	return WrapAccelKey(p)
}

// Struct ActionEntry
type ActionEntry struct {
	Ptr unsafe.Pointer
}

func (v ActionEntry) native() *C.GtkActionEntry {
	return (*C.GtkActionEntry)(v.Ptr)
}
func wrapActionEntry(p *C.GtkActionEntry) ActionEntry {
	return ActionEntry{unsafe.Pointer(p)}
}
func WrapActionEntry(p unsafe.Pointer) ActionEntry {
	return ActionEntry{p}
}
func (v ActionEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapActionEntry(p unsafe.Pointer) interface{} {
	return WrapActionEntry(p)
}

// Struct BindingArg
type BindingArg struct {
	Ptr unsafe.Pointer
}

func (v BindingArg) native() *C.GtkBindingArg {
	return (*C.GtkBindingArg)(v.Ptr)
}
func wrapBindingArg(p *C.GtkBindingArg) BindingArg {
	return BindingArg{unsafe.Pointer(p)}
}
func WrapBindingArg(p unsafe.Pointer) BindingArg {
	return BindingArg{p}
}
func (v BindingArg) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBindingArg(p unsafe.Pointer) interface{} {
	return WrapBindingArg(p)
}

// Struct BindingEntry
type BindingEntry struct {
	Ptr unsafe.Pointer
}

func (v BindingEntry) native() *C.GtkBindingEntry {
	return (*C.GtkBindingEntry)(v.Ptr)
}
func wrapBindingEntry(p *C.GtkBindingEntry) BindingEntry {
	return BindingEntry{unsafe.Pointer(p)}
}
func WrapBindingEntry(p unsafe.Pointer) BindingEntry {
	return BindingEntry{p}
}
func (v BindingEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBindingEntry(p unsafe.Pointer) interface{} {
	return WrapBindingEntry(p)
}

// BindingEntryAddSignalFromString is a wrapper around gtk_binding_entry_add_signal_from_string().
func BindingEntryAddSignalFromString(binding_set BindingSet, signal_desc string) /*gir:GLib*/ glib.TokenType {
	signal_desc0 := (*C.gchar)(C.CString(signal_desc))
	ret0 := C.gtk_binding_entry_add_signal_from_string(binding_set.native(), signal_desc0)
	C.free(unsafe.Pointer(signal_desc0)) /*ch:<stdlib.h>*/
	return /*gir:GLib*/ glib.TokenType(ret0)
}

// Struct BindingSet
type BindingSet struct {
	Ptr unsafe.Pointer
}

func (v BindingSet) native() *C.GtkBindingSet {
	return (*C.GtkBindingSet)(v.Ptr)
}
func wrapBindingSet(p *C.GtkBindingSet) BindingSet {
	return BindingSet{unsafe.Pointer(p)}
}
func WrapBindingSet(p unsafe.Pointer) BindingSet {
	return BindingSet{p}
}
func (v BindingSet) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBindingSet(p unsafe.Pointer) interface{} {
	return WrapBindingSet(p)
}

// BindingSetByClass is a wrapper around gtk_binding_set_by_class().
func BindingSetByClass(object_class unsafe.Pointer) BindingSet {
	ret0 := C.gtk_binding_set_by_class(C.gpointer(object_class))
	return wrapBindingSet(ret0)
}

// BindingSetFind is a wrapper around gtk_binding_set_find().
func BindingSetFind(set_name string) BindingSet {
	set_name0 := (*C.gchar)(C.CString(set_name))
	ret0 := C.gtk_binding_set_find(set_name0)
	C.free(unsafe.Pointer(set_name0)) /*ch:<stdlib.h>*/
	return wrapBindingSet(ret0)
}

// BindingSetNew is a wrapper around gtk_binding_set_new().
func BindingSetNew(set_name string) BindingSet {
	set_name0 := (*C.gchar)(C.CString(set_name))
	ret0 := C.gtk_binding_set_new(set_name0)
	C.free(unsafe.Pointer(set_name0)) /*ch:<stdlib.h>*/
	return wrapBindingSet(ret0)
}

// Struct BindingSignal
type BindingSignal struct {
	Ptr unsafe.Pointer
}

func (v BindingSignal) native() *C.GtkBindingSignal {
	return (*C.GtkBindingSignal)(v.Ptr)
}
func wrapBindingSignal(p *C.GtkBindingSignal) BindingSignal {
	return BindingSignal{unsafe.Pointer(p)}
}
func WrapBindingSignal(p unsafe.Pointer) BindingSignal {
	return BindingSignal{p}
}
func (v BindingSignal) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBindingSignal(p unsafe.Pointer) interface{} {
	return WrapBindingSignal(p)
}

// Struct Border
type Border struct {
	Ptr unsafe.Pointer
}

func (v Border) native() *C.GtkBorder {
	return (*C.GtkBorder)(v.Ptr)
}
func wrapBorder(p *C.GtkBorder) Border {
	return Border{unsafe.Pointer(p)}
}
func WrapBorder(p unsafe.Pointer) Border {
	return Border{p}
}
func (v Border) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBorder(p unsafe.Pointer) interface{} {
	return WrapBorder(p)
}

// BorderNew is a wrapper around gtk_border_new().
func BorderNew() Border {
	ret0 := C.gtk_border_new()
	return wrapBorder(ret0)
}

// Copy is a wrapper around gtk_border_copy().
func (border_ Border) Copy() Border {
	ret0 := C.gtk_border_copy(border_.native())
	return wrapBorder(ret0)
}

// Free is a wrapper around gtk_border_free().
func (border_ Border) Free() {
	C.gtk_border_free(border_.native())
}

// Struct CssSection
type CssSection struct {
	Ptr unsafe.Pointer
}

func (v CssSection) native() *C.GtkCssSection {
	return (*C.GtkCssSection)(v.Ptr)
}
func wrapCssSection(p *C.GtkCssSection) CssSection {
	return CssSection{unsafe.Pointer(p)}
}
func WrapCssSection(p unsafe.Pointer) CssSection {
	return CssSection{p}
}
func (v CssSection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCssSection(p unsafe.Pointer) interface{} {
	return WrapCssSection(p)
}

// GetEndLine is a wrapper around gtk_css_section_get_end_line().
func (section CssSection) GetEndLine() uint {
	ret0 := C.gtk_css_section_get_end_line(section.native())
	return uint(ret0)
}

// GetEndPosition is a wrapper around gtk_css_section_get_end_position().
func (section CssSection) GetEndPosition() uint {
	ret0 := C.gtk_css_section_get_end_position(section.native())
	return uint(ret0)
}

// GetFile is a wrapper around gtk_css_section_get_file().
func (section CssSection) GetFile() gio.File {
	ret0 := C.gtk_css_section_get_file(section.native())
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetParent is a wrapper around gtk_css_section_get_parent().
func (section CssSection) GetParent() CssSection {
	ret0 := C.gtk_css_section_get_parent(section.native())
	return wrapCssSection(ret0)
}

// GetSectionType is a wrapper around gtk_css_section_get_section_type().
func (section CssSection) GetSectionType() CssSectionType {
	ret0 := C.gtk_css_section_get_section_type(section.native())
	return CssSectionType(ret0)
}

// GetStartLine is a wrapper around gtk_css_section_get_start_line().
func (section CssSection) GetStartLine() uint {
	ret0 := C.gtk_css_section_get_start_line(section.native())
	return uint(ret0)
}

// GetStartPosition is a wrapper around gtk_css_section_get_start_position().
func (section CssSection) GetStartPosition() uint {
	ret0 := C.gtk_css_section_get_start_position(section.native())
	return uint(ret0)
}

// Ref is a wrapper around gtk_css_section_ref().
func (section CssSection) Ref() CssSection {
	ret0 := C.gtk_css_section_ref(section.native())
	return wrapCssSection(ret0)
}

// Unref is a wrapper around gtk_css_section_unref().
func (section CssSection) Unref() {
	C.gtk_css_section_unref(section.native())
}

// Struct FileFilterInfo
type FileFilterInfo struct {
	Ptr unsafe.Pointer
}

func (v FileFilterInfo) native() *C.GtkFileFilterInfo {
	return (*C.GtkFileFilterInfo)(v.Ptr)
}
func wrapFileFilterInfo(p *C.GtkFileFilterInfo) FileFilterInfo {
	return FileFilterInfo{unsafe.Pointer(p)}
}
func WrapFileFilterInfo(p unsafe.Pointer) FileFilterInfo {
	return FileFilterInfo{p}
}
func (v FileFilterInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileFilterInfo(p unsafe.Pointer) interface{} {
	return WrapFileFilterInfo(p)
}

// Struct FixedChild
type FixedChild struct {
	Ptr unsafe.Pointer
}

func (v FixedChild) native() *C.GtkFixedChild {
	return (*C.GtkFixedChild)(v.Ptr)
}
func wrapFixedChild(p *C.GtkFixedChild) FixedChild {
	return FixedChild{unsafe.Pointer(p)}
}
func WrapFixedChild(p unsafe.Pointer) FixedChild {
	return FixedChild{p}
}
func (v FixedChild) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFixedChild(p unsafe.Pointer) interface{} {
	return WrapFixedChild(p)
}

// Struct Gradient
type Gradient struct {
	Ptr unsafe.Pointer
}

func (v Gradient) native() *C.GtkGradient {
	return (*C.GtkGradient)(v.Ptr)
}
func wrapGradient(p *C.GtkGradient) Gradient {
	return Gradient{unsafe.Pointer(p)}
}
func WrapGradient(p unsafe.Pointer) Gradient {
	return Gradient{p}
}
func (v Gradient) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGradient(p unsafe.Pointer) interface{} {
	return WrapGradient(p)
}

// Struct IMContextInfo
type IMContextInfo struct {
	Ptr unsafe.Pointer
}

func (v IMContextInfo) native() *C.GtkIMContextInfo {
	return (*C.GtkIMContextInfo)(v.Ptr)
}
func wrapIMContextInfo(p *C.GtkIMContextInfo) IMContextInfo {
	return IMContextInfo{unsafe.Pointer(p)}
}
func WrapIMContextInfo(p unsafe.Pointer) IMContextInfo {
	return IMContextInfo{p}
}
func (v IMContextInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIMContextInfo(p unsafe.Pointer) interface{} {
	return WrapIMContextInfo(p)
}

// Struct IconSet
type IconSet struct {
	Ptr unsafe.Pointer
}

func (v IconSet) native() *C.GtkIconSet {
	return (*C.GtkIconSet)(v.Ptr)
}
func wrapIconSet(p *C.GtkIconSet) IconSet {
	return IconSet{unsafe.Pointer(p)}
}
func WrapIconSet(p unsafe.Pointer) IconSet {
	return IconSet{p}
}
func (v IconSet) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIconSet(p unsafe.Pointer) interface{} {
	return WrapIconSet(p)
}

// Struct IconSource
type IconSource struct {
	Ptr unsafe.Pointer
}

func (v IconSource) native() *C.GtkIconSource {
	return (*C.GtkIconSource)(v.Ptr)
}
func wrapIconSource(p *C.GtkIconSource) IconSource {
	return IconSource{unsafe.Pointer(p)}
}
func WrapIconSource(p unsafe.Pointer) IconSource {
	return IconSource{p}
}
func (v IconSource) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIconSource(p unsafe.Pointer) interface{} {
	return WrapIconSource(p)
}

// Struct PadActionEntry
type PadActionEntry struct {
	Ptr unsafe.Pointer
}

func (v PadActionEntry) native() *C.GtkPadActionEntry {
	return (*C.GtkPadActionEntry)(v.Ptr)
}
func wrapPadActionEntry(p *C.GtkPadActionEntry) PadActionEntry {
	return PadActionEntry{unsafe.Pointer(p)}
}
func WrapPadActionEntry(p unsafe.Pointer) PadActionEntry {
	return PadActionEntry{p}
}
func (v PadActionEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPadActionEntry(p unsafe.Pointer) interface{} {
	return WrapPadActionEntry(p)
}

// Struct PageRange
type PageRange struct {
	Ptr unsafe.Pointer
}

func (v PageRange) native() *C.GtkPageRange {
	return (*C.GtkPageRange)(v.Ptr)
}
func wrapPageRange(p *C.GtkPageRange) PageRange {
	return PageRange{unsafe.Pointer(p)}
}
func WrapPageRange(p unsafe.Pointer) PageRange {
	return PageRange{p}
}
func (v PageRange) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPageRange(p unsafe.Pointer) interface{} {
	return WrapPageRange(p)
}

// Struct PaperSize
type PaperSize struct {
	Ptr unsafe.Pointer
}

func (v PaperSize) native() *C.GtkPaperSize {
	return (*C.GtkPaperSize)(v.Ptr)
}
func wrapPaperSize(p *C.GtkPaperSize) PaperSize {
	return PaperSize{unsafe.Pointer(p)}
}
func WrapPaperSize(p unsafe.Pointer) PaperSize {
	return PaperSize{p}
}
func (v PaperSize) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPaperSize(p unsafe.Pointer) interface{} {
	return WrapPaperSize(p)
}

// PaperSizeNew is a wrapper around gtk_paper_size_new().
func PaperSizeNew(name string) PaperSize {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_paper_size_new(name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapPaperSize(ret0)
}

// PaperSizeNewCustom is a wrapper around gtk_paper_size_new_custom().
func PaperSizeNewCustom(name string, display_name string, width float64, height float64, unit Unit) PaperSize {
	name0 := (*C.gchar)(C.CString(name))
	display_name0 := (*C.gchar)(C.CString(display_name))
	ret0 := C.gtk_paper_size_new_custom(name0, display_name0, C.gdouble(width), C.gdouble(height), C.GtkUnit(unit))
	C.free(unsafe.Pointer(name0))         /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
	return wrapPaperSize(ret0)
}

// PaperSizeNewFromGvariant is a wrapper around gtk_paper_size_new_from_gvariant().
func PaperSizeNewFromGvariant(variant glib.Variant) PaperSize {
	ret0 := C.gtk_paper_size_new_from_gvariant((*C.GVariant)(variant.Ptr))
	return wrapPaperSize(ret0)
}

// PaperSizeNewFromIpp is a wrapper around gtk_paper_size_new_from_ipp().
func PaperSizeNewFromIpp(ipp_name string, width float64, height float64) PaperSize {
	ipp_name0 := (*C.gchar)(C.CString(ipp_name))
	ret0 := C.gtk_paper_size_new_from_ipp(ipp_name0, C.gdouble(width), C.gdouble(height))
	C.free(unsafe.Pointer(ipp_name0)) /*ch:<stdlib.h>*/
	return wrapPaperSize(ret0)
}

// PaperSizeNewFromKeyFile is a wrapper around gtk_paper_size_new_from_key_file().
func PaperSizeNewFromKeyFile(key_file glib.KeyFile, group_name string) (PaperSize, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	var err glib.Error
	ret0 := C.gtk_paper_size_new_from_key_file((*C.GKeyFile)(key_file.Ptr), group_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PaperSize{}, err.GoValue()
	}
	return wrapPaperSize(ret0), nil
}

// PaperSizeNewFromPpd is a wrapper around gtk_paper_size_new_from_ppd().
func PaperSizeNewFromPpd(ppd_name string, ppd_display_name string, width float64, height float64) PaperSize {
	ppd_name0 := (*C.gchar)(C.CString(ppd_name))
	ppd_display_name0 := (*C.gchar)(C.CString(ppd_display_name))
	ret0 := C.gtk_paper_size_new_from_ppd(ppd_name0, ppd_display_name0, C.gdouble(width), C.gdouble(height))
	C.free(unsafe.Pointer(ppd_name0))         /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(ppd_display_name0)) /*ch:<stdlib.h>*/
	return wrapPaperSize(ret0)
}

// Copy is a wrapper around gtk_paper_size_copy().
func (other PaperSize) Copy() PaperSize {
	ret0 := C.gtk_paper_size_copy(other.native())
	return wrapPaperSize(ret0)
}

// Free is a wrapper around gtk_paper_size_free().
func (size PaperSize) Free() {
	C.gtk_paper_size_free(size.native())
}

// GetDefaultBottomMargin is a wrapper around gtk_paper_size_get_default_bottom_margin().
func (size PaperSize) GetDefaultBottomMargin(unit Unit) float64 {
	ret0 := C.gtk_paper_size_get_default_bottom_margin(size.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetDefaultLeftMargin is a wrapper around gtk_paper_size_get_default_left_margin().
func (size PaperSize) GetDefaultLeftMargin(unit Unit) float64 {
	ret0 := C.gtk_paper_size_get_default_left_margin(size.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetDefaultRightMargin is a wrapper around gtk_paper_size_get_default_right_margin().
func (size PaperSize) GetDefaultRightMargin(unit Unit) float64 {
	ret0 := C.gtk_paper_size_get_default_right_margin(size.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetDefaultTopMargin is a wrapper around gtk_paper_size_get_default_top_margin().
func (size PaperSize) GetDefaultTopMargin(unit Unit) float64 {
	ret0 := C.gtk_paper_size_get_default_top_margin(size.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetDisplayName is a wrapper around gtk_paper_size_get_display_name().
func (size PaperSize) GetDisplayName() string {
	ret0 := C.gtk_paper_size_get_display_name(size.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetHeight is a wrapper around gtk_paper_size_get_height().
func (size PaperSize) GetHeight(unit Unit) float64 {
	ret0 := C.gtk_paper_size_get_height(size.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetName is a wrapper around gtk_paper_size_get_name().
func (size PaperSize) GetName() string {
	ret0 := C.gtk_paper_size_get_name(size.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPpdName is a wrapper around gtk_paper_size_get_ppd_name().
func (size PaperSize) GetPpdName() string {
	ret0 := C.gtk_paper_size_get_ppd_name(size.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWidth is a wrapper around gtk_paper_size_get_width().
func (size PaperSize) GetWidth(unit Unit) float64 {
	ret0 := C.gtk_paper_size_get_width(size.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// IsCustom is a wrapper around gtk_paper_size_is_custom().
func (size PaperSize) IsCustom() bool {
	ret0 := C.gtk_paper_size_is_custom(size.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsEqual is a wrapper around gtk_paper_size_is_equal().
func (size1 PaperSize) IsEqual(size2 PaperSize) bool {
	ret0 := C.gtk_paper_size_is_equal(size1.native(), size2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsIpp is a wrapper around gtk_paper_size_is_ipp().
func (size PaperSize) IsIpp() bool {
	ret0 := C.gtk_paper_size_is_ipp(size.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetSize is a wrapper around gtk_paper_size_set_size().
func (size PaperSize) SetSize(width float64, height float64, unit Unit) {
	C.gtk_paper_size_set_size(size.native(), C.gdouble(width), C.gdouble(height), C.GtkUnit(unit))
}

// ToGvariant is a wrapper around gtk_paper_size_to_gvariant().
func (paper_size PaperSize) ToGvariant() glib.Variant {
	ret0 := C.gtk_paper_size_to_gvariant(paper_size.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// ToKeyFile is a wrapper around gtk_paper_size_to_key_file().
func (size PaperSize) ToKeyFile(key_file glib.KeyFile, group_name string) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	C.gtk_paper_size_to_key_file(size.native(), (*C.GKeyFile)(key_file.Ptr), group_name0)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
}

// PaperSizeGetDefault is a wrapper around gtk_paper_size_get_default().
func PaperSizeGetDefault() string {
	ret0 := C.gtk_paper_size_get_default()
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// PaperSizeGetPaperSizes is a wrapper around gtk_paper_size_get_paper_sizes().
func PaperSizeGetPaperSizes(include_custom bool) glib.List {
	ret0 := C.gtk_paper_size_get_paper_sizes(C.gboolean(util.Bool2Int(include_custom)) /*go:.util*/)
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapPaperSize(p) }) /*gir:GLib*/
}

// Struct RadioActionEntry
type RadioActionEntry struct {
	Ptr unsafe.Pointer
}

func (v RadioActionEntry) native() *C.GtkRadioActionEntry {
	return (*C.GtkRadioActionEntry)(v.Ptr)
}
func wrapRadioActionEntry(p *C.GtkRadioActionEntry) RadioActionEntry {
	return RadioActionEntry{unsafe.Pointer(p)}
}
func WrapRadioActionEntry(p unsafe.Pointer) RadioActionEntry {
	return RadioActionEntry{p}
}
func (v RadioActionEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRadioActionEntry(p unsafe.Pointer) interface{} {
	return WrapRadioActionEntry(p)
}

// Struct RcProperty
type RcProperty struct {
	Ptr unsafe.Pointer
}

func (v RcProperty) native() *C.GtkRcProperty {
	return (*C.GtkRcProperty)(v.Ptr)
}
func wrapRcProperty(p *C.GtkRcProperty) RcProperty {
	return RcProperty{unsafe.Pointer(p)}
}
func WrapRcProperty(p unsafe.Pointer) RcProperty {
	return RcProperty{p}
}
func (v RcProperty) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRcProperty(p unsafe.Pointer) interface{} {
	return WrapRcProperty(p)
}

// Struct RecentData
type RecentData struct {
	Ptr unsafe.Pointer
}

func (v RecentData) native() *C.GtkRecentData {
	return (*C.GtkRecentData)(v.Ptr)
}
func wrapRecentData(p *C.GtkRecentData) RecentData {
	return RecentData{unsafe.Pointer(p)}
}
func WrapRecentData(p unsafe.Pointer) RecentData {
	return RecentData{p}
}
func (v RecentData) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentData(p unsafe.Pointer) interface{} {
	return WrapRecentData(p)
}

// Struct RecentFilterInfo
type RecentFilterInfo struct {
	Ptr unsafe.Pointer
}

func (v RecentFilterInfo) native() *C.GtkRecentFilterInfo {
	return (*C.GtkRecentFilterInfo)(v.Ptr)
}
func wrapRecentFilterInfo(p *C.GtkRecentFilterInfo) RecentFilterInfo {
	return RecentFilterInfo{unsafe.Pointer(p)}
}
func WrapRecentFilterInfo(p unsafe.Pointer) RecentFilterInfo {
	return RecentFilterInfo{p}
}
func (v RecentFilterInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentFilterInfo(p unsafe.Pointer) interface{} {
	return WrapRecentFilterInfo(p)
}

// Struct RecentInfo
type RecentInfo struct {
	Ptr unsafe.Pointer
}

func (v RecentInfo) native() *C.GtkRecentInfo {
	return (*C.GtkRecentInfo)(v.Ptr)
}
func wrapRecentInfo(p *C.GtkRecentInfo) RecentInfo {
	return RecentInfo{unsafe.Pointer(p)}
}
func WrapRecentInfo(p unsafe.Pointer) RecentInfo {
	return RecentInfo{p}
}
func (v RecentInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentInfo(p unsafe.Pointer) interface{} {
	return WrapRecentInfo(p)
}

// CreateAppInfo is a wrapper around gtk_recent_info_create_app_info().
func (info RecentInfo) CreateAppInfo(app_name string) (gio.AppInfo, error) {
	app_name0 := (*C.gchar)(C.CString(app_name))
	var err glib.Error
	ret0 := C.gtk_recent_info_create_app_info(info.native(), app_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(app_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return gio.AppInfo{}, err.GoValue()
	}
	return gio.WrapAppInfo(unsafe.Pointer(ret0)) /*gir:Gio*/, nil
}

// Exists is a wrapper around gtk_recent_info_exists().
func (info RecentInfo) Exists() bool {
	ret0 := C.gtk_recent_info_exists(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetAge is a wrapper around gtk_recent_info_get_age().
func (info RecentInfo) GetAge() int {
	ret0 := C.gtk_recent_info_get_age(info.native())
	return int(ret0)
}

// GetApplications is a wrapper around gtk_recent_info_get_applications().
func (info RecentInfo) GetApplications() []string {
	var length0 C.gsize
	ret0 := C.gtk_recent_info_get_applications(info.native(), &length0)
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

// GetDescription is a wrapper around gtk_recent_info_get_description().
func (info RecentInfo) GetDescription() string {
	ret0 := C.gtk_recent_info_get_description(info.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDisplayName is a wrapper around gtk_recent_info_get_display_name().
func (info RecentInfo) GetDisplayName() string {
	ret0 := C.gtk_recent_info_get_display_name(info.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetGicon is a wrapper around gtk_recent_info_get_gicon().
func (info RecentInfo) GetGicon() gio.Icon {
	ret0 := C.gtk_recent_info_get_gicon(info.native())
	return gio.WrapIcon(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetGroups is a wrapper around gtk_recent_info_get_groups().
func (info RecentInfo) GetGroups() []string {
	var length0 C.gsize
	ret0 := C.gtk_recent_info_get_groups(info.native(), &length0)
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

// GetIcon is a wrapper around gtk_recent_info_get_icon().
func (info RecentInfo) GetIcon(size int) gdkpixbuf.Pixbuf {
	ret0 := C.gtk_recent_info_get_icon(info.native(), C.gint(size))
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetMimeType is a wrapper around gtk_recent_info_get_mime_type().
func (info RecentInfo) GetMimeType() string {
	ret0 := C.gtk_recent_info_get_mime_type(info.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPrivateHint is a wrapper around gtk_recent_info_get_private_hint().
func (info RecentInfo) GetPrivateHint() bool {
	ret0 := C.gtk_recent_info_get_private_hint(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShortName is a wrapper around gtk_recent_info_get_short_name().
func (info RecentInfo) GetShortName() string {
	ret0 := C.gtk_recent_info_get_short_name(info.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUri is a wrapper around gtk_recent_info_get_uri().
func (info RecentInfo) GetUri() string {
	ret0 := C.gtk_recent_info_get_uri(info.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUriDisplay is a wrapper around gtk_recent_info_get_uri_display().
func (info RecentInfo) GetUriDisplay() string {
	ret0 := C.gtk_recent_info_get_uri_display(info.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// HasApplication is a wrapper around gtk_recent_info_has_application().
func (info RecentInfo) HasApplication(app_name string) bool {
	app_name0 := (*C.gchar)(C.CString(app_name))
	ret0 := C.gtk_recent_info_has_application(info.native(), app_name0)
	C.free(unsafe.Pointer(app_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))   /*go:.util*/
}

// HasGroup is a wrapper around gtk_recent_info_has_group().
func (info RecentInfo) HasGroup(group_name string) bool {
	group_name0 := (*C.gchar)(C.CString(group_name))
	ret0 := C.gtk_recent_info_has_group(info.native(), group_name0)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// IsLocal is a wrapper around gtk_recent_info_is_local().
func (info RecentInfo) IsLocal() bool {
	ret0 := C.gtk_recent_info_is_local(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LastApplication is a wrapper around gtk_recent_info_last_application().
func (info RecentInfo) LastApplication() string {
	ret0 := C.gtk_recent_info_last_application(info.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Match is a wrapper around gtk_recent_info_match().
func (info_a RecentInfo) Match(info_b RecentInfo) bool {
	ret0 := C.gtk_recent_info_match(info_a.native(), info_b.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around gtk_recent_info_ref().
func (info RecentInfo) Ref() RecentInfo {
	ret0 := C.gtk_recent_info_ref(info.native())
	return wrapRecentInfo(ret0)
}

// Unref is a wrapper around gtk_recent_info_unref().
func (info RecentInfo) Unref() {
	C.gtk_recent_info_unref(info.native())
}

// Struct RequestedSize
type RequestedSize struct {
	Ptr unsafe.Pointer
}

func (v RequestedSize) native() *C.GtkRequestedSize {
	return (*C.GtkRequestedSize)(v.Ptr)
}
func wrapRequestedSize(p *C.GtkRequestedSize) RequestedSize {
	return RequestedSize{unsafe.Pointer(p)}
}
func WrapRequestedSize(p unsafe.Pointer) RequestedSize {
	return RequestedSize{p}
}
func (v RequestedSize) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRequestedSize(p unsafe.Pointer) interface{} {
	return WrapRequestedSize(p)
}

// Struct Requisition
type Requisition struct {
	Ptr unsafe.Pointer
}

func (v Requisition) native() *C.GtkRequisition {
	return (*C.GtkRequisition)(v.Ptr)
}
func wrapRequisition(p *C.GtkRequisition) Requisition {
	return Requisition{unsafe.Pointer(p)}
}
func WrapRequisition(p unsafe.Pointer) Requisition {
	return Requisition{p}
}
func (v Requisition) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRequisition(p unsafe.Pointer) interface{} {
	return WrapRequisition(p)
}

// RequisitionNew is a wrapper around gtk_requisition_new().
func RequisitionNew() Requisition {
	ret0 := C.gtk_requisition_new()
	return wrapRequisition(ret0)
}

// Copy is a wrapper around gtk_requisition_copy().
func (requisition Requisition) Copy() Requisition {
	ret0 := C.gtk_requisition_copy(requisition.native())
	return wrapRequisition(ret0)
}

// Free is a wrapper around gtk_requisition_free().
func (requisition Requisition) Free() {
	C.gtk_requisition_free(requisition.native())
}

// Struct SelectionData
type SelectionData struct {
	Ptr unsafe.Pointer
}

func (v SelectionData) native() *C.GtkSelectionData {
	return (*C.GtkSelectionData)(v.Ptr)
}
func wrapSelectionData(p *C.GtkSelectionData) SelectionData {
	return SelectionData{unsafe.Pointer(p)}
}
func WrapSelectionData(p unsafe.Pointer) SelectionData {
	return SelectionData{p}
}
func (v SelectionData) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSelectionData(p unsafe.Pointer) interface{} {
	return WrapSelectionData(p)
}

// Copy is a wrapper around gtk_selection_data_copy().
func (data SelectionData) Copy() SelectionData {
	ret0 := C.gtk_selection_data_copy(data.native())
	return wrapSelectionData(ret0)
}

// Free is a wrapper around gtk_selection_data_free().
func (data SelectionData) Free() {
	C.gtk_selection_data_free(data.native())
}

// gtk_selection_data_get_data shadowed by get_data_with_length

// GetData is a wrapper around gtk_selection_data_get_data_with_length().
func (selection_data SelectionData) GetData() []byte {
	var length0 C.gint
	ret0 := C.gtk_selection_data_get_data_with_length(selection_data.native(), &length0)
	var ret0Slice []C.guchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]byte, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = byte(elem)
	}
	return ret
}

// GetFormat is a wrapper around gtk_selection_data_get_format().
func (selection_data SelectionData) GetFormat() int {
	ret0 := C.gtk_selection_data_get_format(selection_data.native())
	return int(ret0)
}

// GetLength is a wrapper around gtk_selection_data_get_length().
func (selection_data SelectionData) GetLength() int {
	ret0 := C.gtk_selection_data_get_length(selection_data.native())
	return int(ret0)
}

// GetPixbuf is a wrapper around gtk_selection_data_get_pixbuf().
func (selection_data SelectionData) GetPixbuf() gdkpixbuf.Pixbuf {
	ret0 := C.gtk_selection_data_get_pixbuf(selection_data.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetUris is a wrapper around gtk_selection_data_get_uris().
func (selection_data SelectionData) GetUris() []string {
	ret0 := C.gtk_selection_data_get_uris(selection_data.native())
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

// SetPixbuf is a wrapper around gtk_selection_data_set_pixbuf().
func (selection_data SelectionData) SetPixbuf(pixbuf gdkpixbuf.Pixbuf) bool {
	ret0 := C.gtk_selection_data_set_pixbuf(selection_data.native(), (*C.GdkPixbuf)(pixbuf.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetText is a wrapper around gtk_selection_data_set_text().
func (selection_data SelectionData) SetText(str string, len int) bool {
	str0 := (*C.gchar)(C.CString(str))
	ret0 := C.gtk_selection_data_set_text(selection_data.native(), str0, C.gint(len))
	C.free(unsafe.Pointer(str0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUris is a wrapper around gtk_selection_data_set_uris().
func (selection_data SelectionData) SetUris(uris []string) bool {
	uris0 := make([]*C.gchar, len(uris))
	for idx, elemG := range uris {
		elem := (*C.gchar)(C.CString(elemG))
		uris0[idx] = elem
	}
	var uris0Ptr **C.gchar
	if len(uris0) > 0 {
		uris0Ptr = &uris0[0]
	}
	ret0 := C.gtk_selection_data_set_uris(selection_data.native(), uris0Ptr)
	for _, elem := range uris0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TargetsIncludeImage is a wrapper around gtk_selection_data_targets_include_image().
func (selection_data SelectionData) TargetsIncludeImage(writable bool) bool {
	ret0 := C.gtk_selection_data_targets_include_image(selection_data.native(), C.gboolean(util.Bool2Int(writable)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TargetsIncludeRichText is a wrapper around gtk_selection_data_targets_include_rich_text().
func (selection_data SelectionData) TargetsIncludeRichText(buffer TextBuffer) bool {
	ret0 := C.gtk_selection_data_targets_include_rich_text(selection_data.native(), buffer.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TargetsIncludeText is a wrapper around gtk_selection_data_targets_include_text().
func (selection_data SelectionData) TargetsIncludeText() bool {
	ret0 := C.gtk_selection_data_targets_include_text(selection_data.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TargetsIncludeUri is a wrapper around gtk_selection_data_targets_include_uri().
func (selection_data SelectionData) TargetsIncludeUri() bool {
	ret0 := C.gtk_selection_data_targets_include_uri(selection_data.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object TextBuffer
type TextBuffer struct {
	gobject.Object
}

func (v TextBuffer) native() *C.GtkTextBuffer {
	return (*C.GtkTextBuffer)(v.Ptr)
}
func wrapTextBuffer(p *C.GtkTextBuffer) (v TextBuffer) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextBuffer(p unsafe.Pointer) (v TextBuffer) {
	v.Ptr = p
	return
}
func (v TextBuffer) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextBuffer(p unsafe.Pointer) interface{} {
	return WrapTextBuffer(p)
}
func (v TextBuffer) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_buffer_get_type())
}
func (v TextBuffer) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextBuffer(unsafe.Pointer(ptr)), nil
	}
}

// TextBufferNew is a wrapper around gtk_text_buffer_new().
func TextBufferNew(table TextTagTable) TextBuffer {
	ret0 := C.gtk_text_buffer_new(table.native())
	return wrapTextBuffer(ret0)
}

// AddMark is a wrapper around gtk_text_buffer_add_mark().
func (buffer TextBuffer) AddMark(mark TextMark, where TextIter) {
	C.gtk_text_buffer_add_mark(buffer.native(), mark.native(), where.native())
}

// AddSelectionClipboard is a wrapper around gtk_text_buffer_add_selection_clipboard().
func (buffer TextBuffer) AddSelectionClipboard(clipboard Clipboard) {
	C.gtk_text_buffer_add_selection_clipboard(buffer.native(), clipboard.native())
}

// ApplyTag is a wrapper around gtk_text_buffer_apply_tag().
func (buffer TextBuffer) ApplyTag(tag TextTag, start TextIter, end TextIter) {
	C.gtk_text_buffer_apply_tag(buffer.native(), tag.native(), start.native(), end.native())
}

// ApplyTagByName is a wrapper around gtk_text_buffer_apply_tag_by_name().
func (buffer TextBuffer) ApplyTagByName(name string, start TextIter, end TextIter) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_text_buffer_apply_tag_by_name(buffer.native(), name0, start.native(), end.native())
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// Backspace is a wrapper around gtk_text_buffer_backspace().
func (buffer TextBuffer) Backspace(iter TextIter, interactive bool, default_editable bool) bool {
	ret0 := C.gtk_text_buffer_backspace(buffer.native(), iter.native(), C.gboolean(util.Bool2Int(interactive)) /*go:.util*/, C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BeginUserAction is a wrapper around gtk_text_buffer_begin_user_action().
func (buffer TextBuffer) BeginUserAction() {
	C.gtk_text_buffer_begin_user_action(buffer.native())
}

// CopyClipboard is a wrapper around gtk_text_buffer_copy_clipboard().
func (buffer TextBuffer) CopyClipboard(clipboard Clipboard) {
	C.gtk_text_buffer_copy_clipboard(buffer.native(), clipboard.native())
}

// CreateChildAnchor is a wrapper around gtk_text_buffer_create_child_anchor().
func (buffer TextBuffer) CreateChildAnchor(iter TextIter) TextChildAnchor {
	ret0 := C.gtk_text_buffer_create_child_anchor(buffer.native(), iter.native())
	return wrapTextChildAnchor(ret0)
}

// CreateMark is a wrapper around gtk_text_buffer_create_mark().
func (buffer TextBuffer) CreateMark(mark_name string, where TextIter, left_gravity bool) TextMark {
	mark_name0 := (*C.gchar)(C.CString(mark_name))
	ret0 := C.gtk_text_buffer_create_mark(buffer.native(), mark_name0, where.native(), C.gboolean(util.Bool2Int(left_gravity)) /*go:.util*/)
	C.free(unsafe.Pointer(mark_name0)) /*ch:<stdlib.h>*/
	return wrapTextMark(ret0)
}

// CutClipboard is a wrapper around gtk_text_buffer_cut_clipboard().
func (buffer TextBuffer) CutClipboard(clipboard Clipboard, default_editable bool) {
	C.gtk_text_buffer_cut_clipboard(buffer.native(), clipboard.native(), C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
}

// Delete is a wrapper around gtk_text_buffer_delete().
func (buffer TextBuffer) Delete(start TextIter, end TextIter) {
	C.gtk_text_buffer_delete(buffer.native(), start.native(), end.native())
}

// DeleteInteractive is a wrapper around gtk_text_buffer_delete_interactive().
func (buffer TextBuffer) DeleteInteractive(start_iter TextIter, end_iter TextIter, default_editable bool) bool {
	ret0 := C.gtk_text_buffer_delete_interactive(buffer.native(), start_iter.native(), end_iter.native(), C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// DeleteMark is a wrapper around gtk_text_buffer_delete_mark().
func (buffer TextBuffer) DeleteMark(mark TextMark) {
	C.gtk_text_buffer_delete_mark(buffer.native(), mark.native())
}

// DeleteMarkByName is a wrapper around gtk_text_buffer_delete_mark_by_name().
func (buffer TextBuffer) DeleteMarkByName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_text_buffer_delete_mark_by_name(buffer.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// DeleteSelection is a wrapper around gtk_text_buffer_delete_selection().
func (buffer TextBuffer) DeleteSelection(interactive bool, default_editable bool) bool {
	ret0 := C.gtk_text_buffer_delete_selection(buffer.native(), C.gboolean(util.Bool2Int(interactive)) /*go:.util*/, C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EndUserAction is a wrapper around gtk_text_buffer_end_user_action().
func (buffer TextBuffer) EndUserAction() {
	C.gtk_text_buffer_end_user_action(buffer.native())
}

// GetCharCount is a wrapper around gtk_text_buffer_get_char_count().
func (buffer TextBuffer) GetCharCount() int {
	ret0 := C.gtk_text_buffer_get_char_count(buffer.native())
	return int(ret0)
}

// GetCopyTargetList is a wrapper around gtk_text_buffer_get_copy_target_list().
func (buffer TextBuffer) GetCopyTargetList() TargetList {
	ret0 := C.gtk_text_buffer_get_copy_target_list(buffer.native())
	return wrapTargetList(ret0)
}

// GetHasSelection is a wrapper around gtk_text_buffer_get_has_selection().
func (buffer TextBuffer) GetHasSelection() bool {
	ret0 := C.gtk_text_buffer_get_has_selection(buffer.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInsert is a wrapper around gtk_text_buffer_get_insert().
func (buffer TextBuffer) GetInsert() TextMark {
	ret0 := C.gtk_text_buffer_get_insert(buffer.native())
	return wrapTextMark(ret0)
}

// GetLineCount is a wrapper around gtk_text_buffer_get_line_count().
func (buffer TextBuffer) GetLineCount() int {
	ret0 := C.gtk_text_buffer_get_line_count(buffer.native())
	return int(ret0)
}

// GetMark is a wrapper around gtk_text_buffer_get_mark().
func (buffer TextBuffer) GetMark(name string) TextMark {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_text_buffer_get_mark(buffer.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapTextMark(ret0)
}

// GetModified is a wrapper around gtk_text_buffer_get_modified().
func (buffer TextBuffer) GetModified() bool {
	ret0 := C.gtk_text_buffer_get_modified(buffer.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPasteTargetList is a wrapper around gtk_text_buffer_get_paste_target_list().
func (buffer TextBuffer) GetPasteTargetList() TargetList {
	ret0 := C.gtk_text_buffer_get_paste_target_list(buffer.native())
	return wrapTargetList(ret0)
}

// GetSelectionBound is a wrapper around gtk_text_buffer_get_selection_bound().
func (buffer TextBuffer) GetSelectionBound() TextMark {
	ret0 := C.gtk_text_buffer_get_selection_bound(buffer.native())
	return wrapTextMark(ret0)
}

// GetSlice is a wrapper around gtk_text_buffer_get_slice().
func (buffer TextBuffer) GetSlice(start TextIter, end TextIter, include_hidden_chars bool) string {
	ret0 := C.gtk_text_buffer_get_slice(buffer.native(), start.native(), end.native(), C.gboolean(util.Bool2Int(include_hidden_chars)) /*go:.util*/)
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetTagTable is a wrapper around gtk_text_buffer_get_tag_table().
func (buffer TextBuffer) GetTagTable() TextTagTable {
	ret0 := C.gtk_text_buffer_get_tag_table(buffer.native())
	return wrapTextTagTable(ret0)
}

// GetText is a wrapper around gtk_text_buffer_get_text().
func (buffer TextBuffer) GetText(start TextIter, end TextIter, include_hidden_chars bool) string {
	ret0 := C.gtk_text_buffer_get_text(buffer.native(), start.native(), end.native(), C.gboolean(util.Bool2Int(include_hidden_chars)) /*go:.util*/)
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Insert is a wrapper around gtk_text_buffer_insert().
func (buffer TextBuffer) Insert(iter TextIter, text string, len int) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_text_buffer_insert(buffer.native(), iter.native(), text0, C.gint(len))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// InsertAtCursor is a wrapper around gtk_text_buffer_insert_at_cursor().
func (buffer TextBuffer) InsertAtCursor(text string, len int) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_text_buffer_insert_at_cursor(buffer.native(), text0, C.gint(len))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// InsertChildAnchor is a wrapper around gtk_text_buffer_insert_child_anchor().
func (buffer TextBuffer) InsertChildAnchor(iter TextIter, anchor TextChildAnchor) {
	C.gtk_text_buffer_insert_child_anchor(buffer.native(), iter.native(), anchor.native())
}

// InsertInteractive is a wrapper around gtk_text_buffer_insert_interactive().
func (buffer TextBuffer) InsertInteractive(iter TextIter, text string, len int, default_editable bool) bool {
	text0 := (*C.gchar)(C.CString(text))
	ret0 := C.gtk_text_buffer_insert_interactive(buffer.native(), iter.native(), text0, C.gint(len), C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
	C.free(unsafe.Pointer(text0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InsertInteractiveAtCursor is a wrapper around gtk_text_buffer_insert_interactive_at_cursor().
func (buffer TextBuffer) InsertInteractiveAtCursor(text string, len int, default_editable bool) bool {
	text0 := (*C.gchar)(C.CString(text))
	ret0 := C.gtk_text_buffer_insert_interactive_at_cursor(buffer.native(), text0, C.gint(len), C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
	C.free(unsafe.Pointer(text0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InsertMarkup is a wrapper around gtk_text_buffer_insert_markup().
func (buffer TextBuffer) InsertMarkup(iter TextIter, markup string, len int) {
	markup0 := (*C.gchar)(C.CString(markup))
	C.gtk_text_buffer_insert_markup(buffer.native(), iter.native(), markup0, C.gint(len))
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// InsertRange is a wrapper around gtk_text_buffer_insert_range().
func (buffer TextBuffer) InsertRange(iter TextIter, start TextIter, end TextIter) {
	C.gtk_text_buffer_insert_range(buffer.native(), iter.native(), start.native(), end.native())
}

// InsertRangeInteractive is a wrapper around gtk_text_buffer_insert_range_interactive().
func (buffer TextBuffer) InsertRangeInteractive(iter TextIter, start TextIter, end TextIter, default_editable bool) bool {
	ret0 := C.gtk_text_buffer_insert_range_interactive(buffer.native(), iter.native(), start.native(), end.native(), C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MoveMark is a wrapper around gtk_text_buffer_move_mark().
func (buffer TextBuffer) MoveMark(mark TextMark, where TextIter) {
	C.gtk_text_buffer_move_mark(buffer.native(), mark.native(), where.native())
}

// MoveMarkByName is a wrapper around gtk_text_buffer_move_mark_by_name().
func (buffer TextBuffer) MoveMarkByName(name string, where TextIter) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_text_buffer_move_mark_by_name(buffer.native(), name0, where.native())
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// PasteClipboard is a wrapper around gtk_text_buffer_paste_clipboard().
func (buffer TextBuffer) PasteClipboard(clipboard Clipboard, override_location TextIter, default_editable bool) {
	C.gtk_text_buffer_paste_clipboard(buffer.native(), clipboard.native(), override_location.native(), C.gboolean(util.Bool2Int(default_editable)) /*go:.util*/)
}

// PlaceCursor is a wrapper around gtk_text_buffer_place_cursor().
func (buffer TextBuffer) PlaceCursor(where TextIter) {
	C.gtk_text_buffer_place_cursor(buffer.native(), where.native())
}

// RemoveAllTags is a wrapper around gtk_text_buffer_remove_all_tags().
func (buffer TextBuffer) RemoveAllTags(start TextIter, end TextIter) {
	C.gtk_text_buffer_remove_all_tags(buffer.native(), start.native(), end.native())
}

// RemoveSelectionClipboard is a wrapper around gtk_text_buffer_remove_selection_clipboard().
func (buffer TextBuffer) RemoveSelectionClipboard(clipboard Clipboard) {
	C.gtk_text_buffer_remove_selection_clipboard(buffer.native(), clipboard.native())
}

// RemoveTag is a wrapper around gtk_text_buffer_remove_tag().
func (buffer TextBuffer) RemoveTag(tag TextTag, start TextIter, end TextIter) {
	C.gtk_text_buffer_remove_tag(buffer.native(), tag.native(), start.native(), end.native())
}

// RemoveTagByName is a wrapper around gtk_text_buffer_remove_tag_by_name().
func (buffer TextBuffer) RemoveTagByName(name string, start TextIter, end TextIter) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_text_buffer_remove_tag_by_name(buffer.native(), name0, start.native(), end.native())
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SelectRange is a wrapper around gtk_text_buffer_select_range().
func (buffer TextBuffer) SelectRange(ins TextIter, bound TextIter) {
	C.gtk_text_buffer_select_range(buffer.native(), ins.native(), bound.native())
}

// SetModified is a wrapper around gtk_text_buffer_set_modified().
func (buffer TextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(buffer.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetText is a wrapper around gtk_text_buffer_set_text().
func (buffer TextBuffer) SetText(text string, len int) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_text_buffer_set_text(buffer.native(), text0, C.gint(len))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// Object TextTagTable
type TextTagTable struct {
	gobject.Object
}

func (v TextTagTable) native() *C.GtkTextTagTable {
	return (*C.GtkTextTagTable)(v.Ptr)
}
func wrapTextTagTable(p *C.GtkTextTagTable) (v TextTagTable) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextTagTable(p unsafe.Pointer) (v TextTagTable) {
	v.Ptr = p
	return
}
func (v TextTagTable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextTagTable(p unsafe.Pointer) interface{} {
	return WrapTextTagTable(p)
}
func (v TextTagTable) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_tag_table_get_type())
}
func (v TextTagTable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextTagTable(unsafe.Pointer(ptr)), nil
	}
}
func (v TextTagTable) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// TextTagTableNew is a wrapper around gtk_text_tag_table_new().
func TextTagTableNew() TextTagTable {
	ret0 := C.gtk_text_tag_table_new()
	return wrapTextTagTable(ret0)
}

// Add is a wrapper around gtk_text_tag_table_add().
func (table TextTagTable) Add(tag TextTag) bool {
	ret0 := C.gtk_text_tag_table_add(table.native(), tag.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSize is a wrapper around gtk_text_tag_table_get_size().
func (table TextTagTable) GetSize() int {
	ret0 := C.gtk_text_tag_table_get_size(table.native())
	return int(ret0)
}

// Lookup is a wrapper around gtk_text_tag_table_lookup().
func (table TextTagTable) Lookup(name string) TextTag {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_text_tag_table_lookup(table.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapTextTag(ret0)
}

// Remove is a wrapper around gtk_text_tag_table_remove().
func (table TextTagTable) Remove(tag TextTag) {
	C.gtk_text_tag_table_remove(table.native(), tag.native())
}

// Interface Buildable
type Buildable struct {
	Ptr unsafe.Pointer
}

func (v Buildable) native() *C.GtkBuildable {
	return (*C.GtkBuildable)(v.Ptr)
}
func wrapBuildable(p *C.GtkBuildable) Buildable {
	return Buildable{unsafe.Pointer(p)}
}
func WrapBuildable(p unsafe.Pointer) Buildable {
	return Buildable{p}
}
func (v Buildable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBuildable(p unsafe.Pointer) interface{} {
	return WrapBuildable(p)
}
func (v Buildable) GetType() gobject.Type {
	return gobject.Type(C.gtk_buildable_get_type())
}
func (v Buildable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBuildable(unsafe.Pointer(ptr)), nil
	}
}

// AddChild is a wrapper around gtk_buildable_add_child().
func (buildable Buildable) AddChild(builder Builder, child gobject.Object, type_ string) {
	type0 := (*C.gchar)(C.CString(type_))
	C.gtk_buildable_add_child(buildable.native(), builder.native(), (*C.GObject)(child.Ptr), type0)
	C.free(unsafe.Pointer(type0)) /*ch:<stdlib.h>*/
}

// ConstructChild is a wrapper around gtk_buildable_construct_child().
func (buildable Buildable) ConstructChild(builder Builder, name string) gobject.Object {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_buildable_construct_child(buildable.native(), builder.native(), name0)
	C.free(unsafe.Pointer(name0))                   /*ch:<stdlib.h>*/
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// CustomFinished is a wrapper around gtk_buildable_custom_finished().
func (buildable Buildable) CustomFinished(builder Builder, child gobject.Object, tagname string, data unsafe.Pointer) {
	tagname0 := (*C.gchar)(C.CString(tagname))
	C.gtk_buildable_custom_finished(buildable.native(), builder.native(), (*C.GObject)(child.Ptr), tagname0, C.gpointer(data))
	C.free(unsafe.Pointer(tagname0)) /*ch:<stdlib.h>*/
}

// GetInternalChild is a wrapper around gtk_buildable_get_internal_child().
func (buildable Buildable) GetInternalChild(builder Builder, childname string) gobject.Object {
	childname0 := (*C.gchar)(C.CString(childname))
	ret0 := C.gtk_buildable_get_internal_child(buildable.native(), builder.native(), childname0)
	C.free(unsafe.Pointer(childname0))              /*ch:<stdlib.h>*/
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetName is a wrapper around gtk_buildable_get_name().
func (buildable Buildable) GetName() string {
	ret0 := C.gtk_buildable_get_name(buildable.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// ParserFinished is a wrapper around gtk_buildable_parser_finished().
func (buildable Buildable) ParserFinished(builder Builder) {
	C.gtk_buildable_parser_finished(buildable.native(), builder.native())
}

// SetBuildableProperty is a wrapper around gtk_buildable_set_buildable_property().
func (buildable Buildable) SetBuildableProperty(builder Builder, name string, value gobject.Value) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_buildable_set_buildable_property(buildable.native(), builder.native(), name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetName is a wrapper around gtk_buildable_set_name().
func (buildable Buildable) SetName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_buildable_set_name(buildable.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// Object Builder
type Builder struct {
	gobject.Object
}

func (v Builder) native() *C.GtkBuilder {
	return (*C.GtkBuilder)(v.Ptr)
}
func wrapBuilder(p *C.GtkBuilder) (v Builder) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBuilder(p unsafe.Pointer) (v Builder) {
	v.Ptr = p
	return
}
func (v Builder) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBuilder(p unsafe.Pointer) interface{} {
	return WrapBuilder(p)
}
func (v Builder) GetType() gobject.Type {
	return gobject.Type(C.gtk_builder_get_type())
}
func (v Builder) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBuilder(unsafe.Pointer(ptr)), nil
	}
}

// BuilderNew is a wrapper around gtk_builder_new().
func BuilderNew() Builder {
	ret0 := C.gtk_builder_new()
	return wrapBuilder(ret0)
}

// BuilderNewFromFile is a wrapper around gtk_builder_new_from_file().
func BuilderNewFromFile(filename string) Builder {
	filename0 := (*C.gchar)(C.CString(filename))
	ret0 := C.gtk_builder_new_from_file(filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	return wrapBuilder(ret0)
}

// BuilderNewFromResource is a wrapper around gtk_builder_new_from_resource().
func BuilderNewFromResource(resource_path string) Builder {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	ret0 := C.gtk_builder_new_from_resource(resource_path0)
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
	return wrapBuilder(ret0)
}

// BuilderNewFromString is a wrapper around gtk_builder_new_from_string().
func BuilderNewFromString(string string, length int) Builder {
	string0 := (*C.gchar)(C.CString(string))
	ret0 := C.gtk_builder_new_from_string(string0, C.gssize(length))
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	return wrapBuilder(ret0)
}

// AddFromFile is a wrapper around gtk_builder_add_from_file().
func (builder Builder) AddFromFile(filename string) (uint, error) {
	filename0 := (*C.gchar)(C.CString(filename))
	var err glib.Error
	ret0 := C.gtk_builder_add_from_file(builder.native(), filename0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// AddFromResource is a wrapper around gtk_builder_add_from_resource().
func (builder Builder) AddFromResource(resource_path string) (uint, error) {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	var err glib.Error
	ret0 := C.gtk_builder_add_from_resource(builder.native(), resource_path0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// AddFromString is a wrapper around gtk_builder_add_from_string().
func (builder Builder) AddFromString(buffer string, length uint) (uint, error) {
	buffer0 := (*C.gchar)(C.CString(buffer))
	var err glib.Error
	ret0 := C.gtk_builder_add_from_string(builder.native(), buffer0, C.gsize(length), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(buffer0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// AddObjectsFromFile is a wrapper around gtk_builder_add_objects_from_file().
func (builder Builder) AddObjectsFromFile(filename string, object_ids []string) (uint, error) {
	filename0 := (*C.gchar)(C.CString(filename))
	object_ids0 := make([]*C.gchar, len(object_ids))
	for idx, elemG := range object_ids {
		elem := (*C.gchar)(C.CString(elemG))
		object_ids0[idx] = elem
	}
	var object_ids0Ptr **C.gchar
	if len(object_ids0) > 0 {
		object_ids0Ptr = &object_ids0[0]
	}
	var err glib.Error
	ret0 := C.gtk_builder_add_objects_from_file(builder.native(), filename0, object_ids0Ptr, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	for _, elem := range object_ids0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// AddObjectsFromResource is a wrapper around gtk_builder_add_objects_from_resource().
func (builder Builder) AddObjectsFromResource(resource_path string, object_ids []string) (uint, error) {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	object_ids0 := make([]*C.gchar, len(object_ids))
	for idx, elemG := range object_ids {
		elem := (*C.gchar)(C.CString(elemG))
		object_ids0[idx] = elem
	}
	var object_ids0Ptr **C.gchar
	if len(object_ids0) > 0 {
		object_ids0Ptr = &object_ids0[0]
	}
	var err glib.Error
	ret0 := C.gtk_builder_add_objects_from_resource(builder.native(), resource_path0, object_ids0Ptr, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
	for _, elem := range object_ids0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// AddObjectsFromString is a wrapper around gtk_builder_add_objects_from_string().
func (builder Builder) AddObjectsFromString(buffer string, length uint, object_ids []string) (uint, error) {
	buffer0 := (*C.gchar)(C.CString(buffer))
	object_ids0 := make([]*C.gchar, len(object_ids))
	for idx, elemG := range object_ids {
		elem := (*C.gchar)(C.CString(elemG))
		object_ids0[idx] = elem
	}
	var object_ids0Ptr **C.gchar
	if len(object_ids0) > 0 {
		object_ids0Ptr = &object_ids0[0]
	}
	var err glib.Error
	ret0 := C.gtk_builder_add_objects_from_string(builder.native(), buffer0, C.gsize(length), object_ids0Ptr, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(buffer0)) /*ch:<stdlib.h>*/
	for _, elem := range object_ids0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// ConnectSignals is a wrapper around gtk_builder_connect_signals().
func (builder Builder) ConnectSignals(user_data unsafe.Pointer) {
	C.gtk_builder_connect_signals(builder.native(), C.gpointer(user_data))
}

// ExposeObject is a wrapper around gtk_builder_expose_object().
func (builder Builder) ExposeObject(name string, object gobject.Object) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_builder_expose_object(builder.native(), name0, (*C.GObject)(object.Ptr))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// ExtendWithTemplate is a wrapper around gtk_builder_extend_with_template().
func (builder Builder) ExtendWithTemplate(widget Widget, template_type /*Gir:GObject*/ gobject.Type, buffer string, length uint) (uint, error) {
	buffer0 := (*C.gchar)(C.CString(buffer))
	var err glib.Error
	ret0 := C.gtk_builder_extend_with_template(builder.native(), widget.native(), C.GType(template_type), buffer0, C.gsize(length), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(buffer0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// GetApplication is a wrapper around gtk_builder_get_application().
func (builder Builder) GetApplication() Application {
	ret0 := C.gtk_builder_get_application(builder.native())
	return wrapApplication(ret0)
}

// GetObject is a wrapper around gtk_builder_get_object().
func (builder Builder) GetObject(name string) gobject.Object {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_builder_get_object(builder.native(), name0)
	C.free(unsafe.Pointer(name0))                   /*ch:<stdlib.h>*/
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetTranslationDomain is a wrapper around gtk_builder_get_translation_domain().
func (builder Builder) GetTranslationDomain() string {
	ret0 := C.gtk_builder_get_translation_domain(builder.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTypeFromName is a wrapper around gtk_builder_get_type_from_name().
func (builder Builder) GetTypeFromName(type_name string) /*Gir:GObject*/ gobject.Type {
	type_name0 := C.CString(type_name)
	ret0 := C.gtk_builder_get_type_from_name(builder.native(), type_name0)
	C.free(unsafe.Pointer(type_name0)) /*ch:<stdlib.h>*/
	return /*Gir:GObject*/ gobject.Type(ret0)
}

// SetTranslationDomain is a wrapper around gtk_builder_set_translation_domain().
func (builder Builder) SetTranslationDomain(domain string) {
	domain0 := (*C.gchar)(C.CString(domain))
	C.gtk_builder_set_translation_domain(builder.native(), domain0)
	C.free(unsafe.Pointer(domain0)) /*ch:<stdlib.h>*/
}

// Object Widget
type Widget struct {
	gobject.InitiallyUnowned
}

func (v Widget) native() *C.GtkWidget {
	return (*C.GtkWidget)(v.Ptr)
}
func wrapWidget(p *C.GtkWidget) (v Widget) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapWidget(p unsafe.Pointer) (v Widget) {
	v.Ptr = p
	return
}
func (v Widget) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWidget(p unsafe.Pointer) interface{} {
	return WrapWidget(p)
}
func (v Widget) GetType() gobject.Type {
	return gobject.Type(C.gtk_widget_get_type())
}
func (v Widget) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapWidget(unsafe.Pointer(ptr)), nil
	}
}
func (v Widget) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Widget) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Activate is a wrapper around gtk_widget_activate().
func (widget Widget) Activate() bool {
	ret0 := C.gtk_widget_activate(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (widget Widget) AddEvents(events int) {
	C.gtk_widget_add_events(widget.native(), C.gint(events))
}

// AddMnemonicLabel is a wrapper around gtk_widget_add_mnemonic_label().
func (widget Widget) AddMnemonicLabel(label Widget) {
	C.gtk_widget_add_mnemonic_label(widget.native(), label.native())
}

// CanActivateAccel is a wrapper around gtk_widget_can_activate_accel().
func (widget Widget) CanActivateAccel(signal_id uint) bool {
	ret0 := C.gtk_widget_can_activate_accel(widget.native(), C.guint(signal_id))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ChildFocus is a wrapper around gtk_widget_child_focus().
func (widget Widget) ChildFocus(direction DirectionType) bool {
	ret0 := C.gtk_widget_child_focus(widget.native(), C.GtkDirectionType(direction))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ChildNotify is a wrapper around gtk_widget_child_notify().
func (widget Widget) ChildNotify(child_property string) {
	child_property0 := (*C.gchar)(C.CString(child_property))
	C.gtk_widget_child_notify(widget.native(), child_property0)
	C.free(unsafe.Pointer(child_property0)) /*ch:<stdlib.h>*/
}

// ComputeExpand is a wrapper around gtk_widget_compute_expand().
func (widget Widget) ComputeExpand(orientation Orientation) bool {
	ret0 := C.gtk_widget_compute_expand(widget.native(), C.GtkOrientation(orientation))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CreatePangoContext is a wrapper around gtk_widget_create_pango_context().
func (widget Widget) CreatePangoContext() pango.Context {
	ret0 := C.gtk_widget_create_pango_context(widget.native())
	return pango.WrapContext(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// CreatePangoLayout is a wrapper around gtk_widget_create_pango_layout().
func (widget Widget) CreatePangoLayout(text string) pango.Layout {
	text0 := (*C.gchar)(C.CString(text))
	ret0 := C.gtk_widget_create_pango_layout(widget.native(), text0)
	C.free(unsafe.Pointer(text0))                 /*ch:<stdlib.h>*/
	return pango.WrapLayout(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// Destroy is a wrapper around gtk_widget_destroy().
func (widget Widget) Destroy() {
	C.gtk_widget_destroy(widget.native())
}

// DragCheckThreshold is a wrapper around gtk_drag_check_threshold().
func (widget Widget) DragCheckThreshold(start_x int, start_y int, current_x int, current_y int) bool {
	ret0 := C.gtk_drag_check_threshold(widget.native(), C.gint(start_x), C.gint(start_y), C.gint(current_x), C.gint(current_y))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// DragDestAddImageTargets is a wrapper around gtk_drag_dest_add_image_targets().
func (widget Widget) DragDestAddImageTargets() {
	C.gtk_drag_dest_add_image_targets(widget.native())
}

// DragDestAddTextTargets is a wrapper around gtk_drag_dest_add_text_targets().
func (widget Widget) DragDestAddTextTargets() {
	C.gtk_drag_dest_add_text_targets(widget.native())
}

// DragDestAddUriTargets is a wrapper around gtk_drag_dest_add_uri_targets().
func (widget Widget) DragDestAddUriTargets() {
	C.gtk_drag_dest_add_uri_targets(widget.native())
}

// DragDestGetTargetList is a wrapper around gtk_drag_dest_get_target_list().
func (widget Widget) DragDestGetTargetList() TargetList {
	ret0 := C.gtk_drag_dest_get_target_list(widget.native())
	return wrapTargetList(ret0)
}

// DragDestGetTrackMotion is a wrapper around gtk_drag_dest_get_track_motion().
func (widget Widget) DragDestGetTrackMotion() bool {
	ret0 := C.gtk_drag_dest_get_track_motion(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// DragDestSetTargetList is a wrapper around gtk_drag_dest_set_target_list().
func (widget Widget) DragDestSetTargetList(target_list TargetList) {
	C.gtk_drag_dest_set_target_list(widget.native(), target_list.native())
}

// DragDestSetTrackMotion is a wrapper around gtk_drag_dest_set_track_motion().
func (widget Widget) DragDestSetTrackMotion(track_motion bool) {
	C.gtk_drag_dest_set_track_motion(widget.native(), C.gboolean(util.Bool2Int(track_motion)) /*go:.util*/)
}

// DragDestUnset is a wrapper around gtk_drag_dest_unset().
func (widget Widget) DragDestUnset() {
	C.gtk_drag_dest_unset(widget.native())
}

// DragHighlight is a wrapper around gtk_drag_highlight().
func (widget Widget) DragHighlight() {
	C.gtk_drag_highlight(widget.native())
}

// DragSourceAddImageTargets is a wrapper around gtk_drag_source_add_image_targets().
func (widget Widget) DragSourceAddImageTargets() {
	C.gtk_drag_source_add_image_targets(widget.native())
}

// DragSourceAddTextTargets is a wrapper around gtk_drag_source_add_text_targets().
func (widget Widget) DragSourceAddTextTargets() {
	C.gtk_drag_source_add_text_targets(widget.native())
}

// DragSourceAddUriTargets is a wrapper around gtk_drag_source_add_uri_targets().
func (widget Widget) DragSourceAddUriTargets() {
	C.gtk_drag_source_add_uri_targets(widget.native())
}

// DragSourceGetTargetList is a wrapper around gtk_drag_source_get_target_list().
func (widget Widget) DragSourceGetTargetList() TargetList {
	ret0 := C.gtk_drag_source_get_target_list(widget.native())
	return wrapTargetList(ret0)
}

// DragSourceSetIconGicon is a wrapper around gtk_drag_source_set_icon_gicon().
func (widget Widget) DragSourceSetIconGicon(icon gio.Icon) {
	C.gtk_drag_source_set_icon_gicon(widget.native(), (*C.GIcon)(icon.Ptr))
}

// DragSourceSetIconName is a wrapper around gtk_drag_source_set_icon_name().
func (widget Widget) DragSourceSetIconName(icon_name string) {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	C.gtk_drag_source_set_icon_name(widget.native(), icon_name0)
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
}

// DragSourceSetTargetList is a wrapper around gtk_drag_source_set_target_list().
func (widget Widget) DragSourceSetTargetList(target_list TargetList) {
	C.gtk_drag_source_set_target_list(widget.native(), target_list.native())
}

// DragSourceUnset is a wrapper around gtk_drag_source_unset().
func (widget Widget) DragSourceUnset() {
	C.gtk_drag_source_unset(widget.native())
}

// DragUnhighlight is a wrapper around gtk_drag_unhighlight().
func (widget Widget) DragUnhighlight() {
	C.gtk_drag_unhighlight(widget.native())
}

// ErrorBell is a wrapper around gtk_widget_error_bell().
func (widget Widget) ErrorBell() {
	C.gtk_widget_error_bell(widget.native())
}

// FreezeChildNotify is a wrapper around gtk_widget_freeze_child_notify().
func (widget Widget) FreezeChildNotify() {
	C.gtk_widget_freeze_child_notify(widget.native())
}

// GetAccessible is a wrapper around gtk_widget_get_accessible().
func (widget Widget) GetAccessible() atk.Object {
	ret0 := C.gtk_widget_get_accessible(widget.native())
	return atk.WrapObject(unsafe.Pointer(ret0)) /*gir:Atk*/
}

// GetActionGroup is a wrapper around gtk_widget_get_action_group().
func (widget Widget) GetActionGroup(prefix string) gio.ActionGroup {
	prefix0 := (*C.gchar)(C.CString(prefix))
	ret0 := C.gtk_widget_get_action_group(widget.native(), prefix0)
	C.free(unsafe.Pointer(prefix0))                  /*ch:<stdlib.h>*/
	return gio.WrapActionGroup(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetAllocatedBaseline is a wrapper around gtk_widget_get_allocated_baseline().
func (widget Widget) GetAllocatedBaseline() int {
	ret0 := C.gtk_widget_get_allocated_baseline(widget.native())
	return int(ret0)
}

// GetAllocatedHeight is a wrapper around gtk_widget_get_allocated_height().
func (widget Widget) GetAllocatedHeight() int {
	ret0 := C.gtk_widget_get_allocated_height(widget.native())
	return int(ret0)
}

// GetAllocatedWidth is a wrapper around gtk_widget_get_allocated_width().
func (widget Widget) GetAllocatedWidth() int {
	ret0 := C.gtk_widget_get_allocated_width(widget.native())
	return int(ret0)
}

// GetAncestor is a wrapper around gtk_widget_get_ancestor().
func (widget Widget) GetAncestor(widget_type /*Gir:GObject*/ gobject.Type) Widget {
	ret0 := C.gtk_widget_get_ancestor(widget.native(), C.GType(widget_type))
	return wrapWidget(ret0)
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (widget Widget) GetAppPaintable() bool {
	ret0 := C.gtk_widget_get_app_paintable(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCanDefault is a wrapper around gtk_widget_get_can_default().
func (widget Widget) GetCanDefault() bool {
	ret0 := C.gtk_widget_get_can_default(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (widget Widget) GetCanFocus() bool {
	ret0 := C.gtk_widget_get_can_focus(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChildVisible is a wrapper around gtk_widget_get_child_visible().
func (widget Widget) GetChildVisible() bool {
	ret0 := C.gtk_widget_get_child_visible(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDirection is a wrapper around gtk_widget_get_direction().
func (widget Widget) GetDirection() TextDirection {
	ret0 := C.gtk_widget_get_direction(widget.native())
	return TextDirection(ret0)
}

// GetDoubleBuffered is a wrapper around gtk_widget_get_double_buffered().
func (widget Widget) GetDoubleBuffered() bool {
	ret0 := C.gtk_widget_get_double_buffered(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (widget Widget) GetEvents() int {
	ret0 := C.gtk_widget_get_events(widget.native())
	return int(ret0)
}

// GetFocusOnClick is a wrapper around gtk_widget_get_focus_on_click().
func (widget Widget) GetFocusOnClick() bool {
	ret0 := C.gtk_widget_get_focus_on_click(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFontMap is a wrapper around gtk_widget_get_font_map().
func (widget Widget) GetFontMap() pango.FontMap {
	ret0 := C.gtk_widget_get_font_map(widget.native())
	return pango.WrapFontMap(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetFontOptions is a wrapper around gtk_widget_get_font_options().
func (widget Widget) GetFontOptions() cairo.FontOptions {
	ret0 := C.gtk_widget_get_font_options(widget.native())
	return cairo.WrapFontOptions(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetHalign is a wrapper around gtk_widget_get_halign().
func (widget Widget) GetHalign() Align {
	ret0 := C.gtk_widget_get_halign(widget.native())
	return Align(ret0)
}

// GetHasTooltip is a wrapper around gtk_widget_get_has_tooltip().
func (widget Widget) GetHasTooltip() bool {
	ret0 := C.gtk_widget_get_has_tooltip(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (widget Widget) GetHasWindow() bool {
	ret0 := C.gtk_widget_get_has_window(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHexpand is a wrapper around gtk_widget_get_hexpand().
func (widget Widget) GetHexpand() bool {
	ret0 := C.gtk_widget_get_hexpand(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHexpandSet is a wrapper around gtk_widget_get_hexpand_set().
func (widget Widget) GetHexpandSet() bool {
	ret0 := C.gtk_widget_get_hexpand_set(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMapped is a wrapper around gtk_widget_get_mapped().
func (widget Widget) GetMapped() bool {
	ret0 := C.gtk_widget_get_mapped(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (widget Widget) GetMarginBottom() int {
	ret0 := C.gtk_widget_get_margin_bottom(widget.native())
	return int(ret0)
}

// GetMarginEnd is a wrapper around gtk_widget_get_margin_end().
func (widget Widget) GetMarginEnd() int {
	ret0 := C.gtk_widget_get_margin_end(widget.native())
	return int(ret0)
}

// GetMarginStart is a wrapper around gtk_widget_get_margin_start().
func (widget Widget) GetMarginStart() int {
	ret0 := C.gtk_widget_get_margin_start(widget.native())
	return int(ret0)
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (widget Widget) GetMarginTop() int {
	ret0 := C.gtk_widget_get_margin_top(widget.native())
	return int(ret0)
}

// GetName is a wrapper around gtk_widget_get_name().
func (widget Widget) GetName() string {
	ret0 := C.gtk_widget_get_name(widget.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (widget Widget) GetNoShowAll() bool {
	ret0 := C.gtk_widget_get_no_show_all(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetOpacity is a wrapper around gtk_widget_get_opacity().
func (widget Widget) GetOpacity() float64 {
	ret0 := C.gtk_widget_get_opacity(widget.native())
	return float64(ret0)
}

// GetPangoContext is a wrapper around gtk_widget_get_pango_context().
func (widget Widget) GetPangoContext() pango.Context {
	ret0 := C.gtk_widget_get_pango_context(widget.native())
	return pango.WrapContext(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (widget Widget) GetParent() Widget {
	ret0 := C.gtk_widget_get_parent(widget.native())
	return wrapWidget(ret0)
}

// GetParentWindow is a wrapper around gtk_widget_get_parent_window().
func (widget Widget) GetParentWindow() gdk.Window {
	ret0 := C.gtk_widget_get_parent_window(widget.native())
	return gdk.WrapWindow(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetPath is a wrapper around gtk_widget_get_path().
func (widget Widget) GetPath() WidgetPath {
	ret0 := C.gtk_widget_get_path(widget.native())
	return wrapWidgetPath(ret0)
}

// GetPreferredHeight is a wrapper around gtk_widget_get_preferred_height().
func (widget Widget) GetPreferredHeight() (int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	C.gtk_widget_get_preferred_height(widget.native(), &minimum_height0, &natural_height0)
	return int(minimum_height0), int(natural_height0)
}

// GetPreferredHeightAndBaselineForWidth is a wrapper around gtk_widget_get_preferred_height_and_baseline_for_width().
func (widget Widget) GetPreferredHeightAndBaselineForWidth(width int) (int, int, int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	var minimum_baseline0 C.gint
	var natural_baseline0 C.gint
	C.gtk_widget_get_preferred_height_and_baseline_for_width(widget.native(), C.gint(width), &minimum_height0, &natural_height0, &minimum_baseline0, &natural_baseline0)
	return int(minimum_height0), int(natural_height0), int(minimum_baseline0), int(natural_baseline0)
}

// GetPreferredHeightForWidth is a wrapper around gtk_widget_get_preferred_height_for_width().
func (widget Widget) GetPreferredHeightForWidth(width int) (int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	C.gtk_widget_get_preferred_height_for_width(widget.native(), C.gint(width), &minimum_height0, &natural_height0)
	return int(minimum_height0), int(natural_height0)
}

// GetPreferredWidth is a wrapper around gtk_widget_get_preferred_width().
func (widget Widget) GetPreferredWidth() (int, int) {
	var minimum_width0 C.gint
	var natural_width0 C.gint
	C.gtk_widget_get_preferred_width(widget.native(), &minimum_width0, &natural_width0)
	return int(minimum_width0), int(natural_width0)
}

// GetPreferredWidthForHeight is a wrapper around gtk_widget_get_preferred_width_for_height().
func (widget Widget) GetPreferredWidthForHeight(height int) (int, int) {
	var minimum_width0 C.gint
	var natural_width0 C.gint
	C.gtk_widget_get_preferred_width_for_height(widget.native(), C.gint(height), &minimum_width0, &natural_width0)
	return int(minimum_width0), int(natural_width0)
}

// GetRealized is a wrapper around gtk_widget_get_realized().
func (widget Widget) GetRealized() bool {
	ret0 := C.gtk_widget_get_realized(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetReceivesDefault is a wrapper around gtk_widget_get_receives_default().
func (widget Widget) GetReceivesDefault() bool {
	ret0 := C.gtk_widget_get_receives_default(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRequestMode is a wrapper around gtk_widget_get_request_mode().
func (widget Widget) GetRequestMode() SizeRequestMode {
	ret0 := C.gtk_widget_get_request_mode(widget.native())
	return SizeRequestMode(ret0)
}

// GetScaleFactor is a wrapper around gtk_widget_get_scale_factor().
func (widget Widget) GetScaleFactor() int {
	ret0 := C.gtk_widget_get_scale_factor(widget.native())
	return int(ret0)
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (widget Widget) GetSensitive() bool {
	ret0 := C.gtk_widget_get_sensitive(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSettings is a wrapper around gtk_widget_get_settings().
func (widget Widget) GetSettings() Settings {
	ret0 := C.gtk_widget_get_settings(widget.native())
	return wrapSettings(ret0)
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (widget Widget) GetSizeRequest() (int, int) {
	var width0 C.gint
	var height0 C.gint
	C.gtk_widget_get_size_request(widget.native(), &width0, &height0)
	return int(width0), int(height0)
}

// GetStateFlags is a wrapper around gtk_widget_get_state_flags().
func (widget Widget) GetStateFlags() StateFlags {
	ret0 := C.gtk_widget_get_state_flags(widget.native())
	return StateFlags(ret0)
}

// GetStyleContext is a wrapper around gtk_widget_get_style_context().
func (widget Widget) GetStyleContext() StyleContext {
	ret0 := C.gtk_widget_get_style_context(widget.native())
	return wrapStyleContext(ret0)
}

// GetSupportMultidevice is a wrapper around gtk_widget_get_support_multidevice().
func (widget Widget) GetSupportMultidevice() bool {
	ret0 := C.gtk_widget_get_support_multidevice(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTemplateChild is a wrapper around gtk_widget_get_template_child().
func (widget Widget) GetTemplateChild(widget_type /*Gir:GObject*/ gobject.Type, name string) gobject.Object {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_widget_get_template_child(widget.native(), C.GType(widget_type), name0)
	C.free(unsafe.Pointer(name0))                   /*ch:<stdlib.h>*/
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetTooltipMarkup is a wrapper around gtk_widget_get_tooltip_markup().
func (widget Widget) GetTooltipMarkup() string {
	ret0 := C.gtk_widget_get_tooltip_markup(widget.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
func (widget Widget) GetTooltipText() string {
	ret0 := C.gtk_widget_get_tooltip_text(widget.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetTooltipWindow is a wrapper around gtk_widget_get_tooltip_window().
func (widget Widget) GetTooltipWindow() Window {
	ret0 := C.gtk_widget_get_tooltip_window(widget.native())
	return wrapWindow(ret0)
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (widget Widget) GetToplevel() Widget {
	ret0 := C.gtk_widget_get_toplevel(widget.native())
	return wrapWidget(ret0)
}

// GetValign is a wrapper around gtk_widget_get_valign().
func (widget Widget) GetValign() Align {
	ret0 := C.gtk_widget_get_valign(widget.native())
	return Align(ret0)
}

// GetValignWithBaseline is a wrapper around gtk_widget_get_valign_with_baseline().
func (widget Widget) GetValignWithBaseline() Align {
	ret0 := C.gtk_widget_get_valign_with_baseline(widget.native())
	return Align(ret0)
}

// GetVexpand is a wrapper around gtk_widget_get_vexpand().
func (widget Widget) GetVexpand() bool {
	ret0 := C.gtk_widget_get_vexpand(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVexpandSet is a wrapper around gtk_widget_get_vexpand_set().
func (widget Widget) GetVexpandSet() bool {
	ret0 := C.gtk_widget_get_vexpand_set(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (widget Widget) GetVisible() bool {
	ret0 := C.gtk_widget_get_visible(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVisual is a wrapper around gtk_widget_get_visual().
func (widget Widget) GetVisual() gdk.Visual {
	ret0 := C.gtk_widget_get_visual(widget.native())
	return gdk.WrapVisual(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GrabAdd is a wrapper around gtk_grab_add().
func (widget Widget) GrabAdd() {
	C.gtk_grab_add(widget.native())
}

// GrabDefault is a wrapper around gtk_widget_grab_default().
func (widget Widget) GrabDefault() {
	C.gtk_widget_grab_default(widget.native())
}

// GrabFocus is a wrapper around gtk_widget_grab_focus().
func (widget Widget) GrabFocus() {
	C.gtk_widget_grab_focus(widget.native())
}

// GrabRemove is a wrapper around gtk_grab_remove().
func (widget Widget) GrabRemove() {
	C.gtk_grab_remove(widget.native())
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (widget Widget) HasDefault() bool {
	ret0 := C.gtk_widget_has_default(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (widget Widget) HasFocus() bool {
	ret0 := C.gtk_widget_has_focus(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (widget Widget) HasGrab() bool {
	ret0 := C.gtk_widget_has_grab(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasScreen is a wrapper around gtk_widget_has_screen().
func (widget Widget) HasScreen() bool {
	ret0 := C.gtk_widget_has_screen(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (widget Widget) HasVisibleFocus() bool {
	ret0 := C.gtk_widget_has_visible_focus(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Hide is a wrapper around gtk_widget_hide().
func (widget Widget) Hide() {
	C.gtk_widget_hide(widget.native())
}

// HideOnDelete is a wrapper around gtk_widget_hide_on_delete().
func (widget Widget) HideOnDelete() bool {
	ret0 := C.gtk_widget_hide_on_delete(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (widget Widget) InDestruction() bool {
	ret0 := C.gtk_widget_in_destruction(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InitTemplate is a wrapper around gtk_widget_init_template().
func (widget Widget) InitTemplate() {
	C.gtk_widget_init_template(widget.native())
}

// InsertActionGroup is a wrapper around gtk_widget_insert_action_group().
func (widget Widget) InsertActionGroup(name string, group gio.ActionGroup) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_widget_insert_action_group(widget.native(), name0, (*C.GActionGroup)(group.Ptr))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// IsAncestor is a wrapper around gtk_widget_is_ancestor().
func (widget Widget) IsAncestor(ancestor Widget) bool {
	ret0 := C.gtk_widget_is_ancestor(widget.native(), ancestor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (widget Widget) IsDrawable() bool {
	ret0 := C.gtk_widget_is_drawable(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsFocus is a wrapper around gtk_widget_is_focus().
func (widget Widget) IsFocus() bool {
	ret0 := C.gtk_widget_is_focus(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (widget Widget) IsSensitive() bool {
	ret0 := C.gtk_widget_is_sensitive(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (widget Widget) IsToplevel() bool {
	ret0 := C.gtk_widget_is_toplevel(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsVisible is a wrapper around gtk_widget_is_visible().
func (widget Widget) IsVisible() bool {
	ret0 := C.gtk_widget_is_visible(widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// KeynavFailed is a wrapper around gtk_widget_keynav_failed().
func (widget Widget) KeynavFailed(direction DirectionType) bool {
	ret0 := C.gtk_widget_keynav_failed(widget.native(), C.GtkDirectionType(direction))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ListActionPrefixes is a wrapper around gtk_widget_list_action_prefixes().
func (widget Widget) ListActionPrefixes() []string {
	ret0 := C.gtk_widget_list_action_prefixes(widget.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// ListMnemonicLabels is a wrapper around gtk_widget_list_mnemonic_labels().
func (widget Widget) ListMnemonicLabels() glib.List {
	ret0 := C.gtk_widget_list_mnemonic_labels(widget.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWidget(p) }) /*gir:GLib*/
}

// Map is a wrapper around gtk_widget_map().
func (widget Widget) Map() {
	C.gtk_widget_map(widget.native())
}

// MnemonicActivate is a wrapper around gtk_widget_mnemonic_activate().
func (widget Widget) MnemonicActivate(group_cycling bool) bool {
	ret0 := C.gtk_widget_mnemonic_activate(widget.native(), C.gboolean(util.Bool2Int(group_cycling)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// QueueAllocate is a wrapper around gtk_widget_queue_allocate().
func (widget Widget) QueueAllocate() {
	C.gtk_widget_queue_allocate(widget.native())
}

// QueueComputeExpand is a wrapper around gtk_widget_queue_compute_expand().
func (widget Widget) QueueComputeExpand() {
	C.gtk_widget_queue_compute_expand(widget.native())
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (widget Widget) QueueDraw() {
	C.gtk_widget_queue_draw(widget.native())
}

// QueueDrawArea is a wrapper around gtk_widget_queue_draw_area().
func (widget Widget) QueueDrawArea(x int, y int, width int, height int) {
	C.gtk_widget_queue_draw_area(widget.native(), C.gint(x), C.gint(y), C.gint(width), C.gint(height))
}

// QueueResize is a wrapper around gtk_widget_queue_resize().
func (widget Widget) QueueResize() {
	C.gtk_widget_queue_resize(widget.native())
}

// QueueResizeNoRedraw is a wrapper around gtk_widget_queue_resize_no_redraw().
func (widget Widget) QueueResizeNoRedraw() {
	C.gtk_widget_queue_resize_no_redraw(widget.native())
}

// Realize is a wrapper around gtk_widget_realize().
func (widget Widget) Realize() {
	C.gtk_widget_realize(widget.native())
}

// RemoveMnemonicLabel is a wrapper around gtk_widget_remove_mnemonic_label().
func (widget Widget) RemoveMnemonicLabel(label Widget) {
	C.gtk_widget_remove_mnemonic_label(widget.native(), label.native())
}

// RemoveTickCallback is a wrapper around gtk_widget_remove_tick_callback().
func (widget Widget) RemoveTickCallback(id uint) {
	C.gtk_widget_remove_tick_callback(widget.native(), C.guint(id))
}

// ResetStyle is a wrapper around gtk_widget_reset_style().
func (widget Widget) ResetStyle() {
	C.gtk_widget_reset_style(widget.native())
}

// SetAccelPath is a wrapper around gtk_widget_set_accel_path().
func (widget Widget) SetAccelPath(accel_path string, accel_group AccelGroup) {
	accel_path0 := (*C.gchar)(C.CString(accel_path))
	C.gtk_widget_set_accel_path(widget.native(), accel_path0, accel_group.native())
	C.free(unsafe.Pointer(accel_path0)) /*ch:<stdlib.h>*/
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (widget Widget) SetAppPaintable(app_paintable bool) {
	C.gtk_widget_set_app_paintable(widget.native(), C.gboolean(util.Bool2Int(app_paintable)) /*go:.util*/)
}

// SetCanDefault is a wrapper around gtk_widget_set_can_default().
func (widget Widget) SetCanDefault(can_default bool) {
	C.gtk_widget_set_can_default(widget.native(), C.gboolean(util.Bool2Int(can_default)) /*go:.util*/)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (widget Widget) SetCanFocus(can_focus bool) {
	C.gtk_widget_set_can_focus(widget.native(), C.gboolean(util.Bool2Int(can_focus)) /*go:.util*/)
}

// SetChildVisible is a wrapper around gtk_widget_set_child_visible().
func (widget Widget) SetChildVisible(is_visible bool) {
	C.gtk_widget_set_child_visible(widget.native(), C.gboolean(util.Bool2Int(is_visible)) /*go:.util*/)
}

// SetDirection is a wrapper around gtk_widget_set_direction().
func (widget Widget) SetDirection(dir TextDirection) {
	C.gtk_widget_set_direction(widget.native(), C.GtkTextDirection(dir))
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (widget Widget) SetEvents(events int) {
	C.gtk_widget_set_events(widget.native(), C.gint(events))
}

// SetFocusOnClick is a wrapper around gtk_widget_set_focus_on_click().
func (widget Widget) SetFocusOnClick(focus_on_click bool) {
	C.gtk_widget_set_focus_on_click(widget.native(), C.gboolean(util.Bool2Int(focus_on_click)) /*go:.util*/)
}

// SetHalign is a wrapper around gtk_widget_set_halign().
func (widget Widget) SetHalign(align Align) {
	C.gtk_widget_set_halign(widget.native(), C.GtkAlign(align))
}

// SetHasTooltip is a wrapper around gtk_widget_set_has_tooltip().
func (widget Widget) SetHasTooltip(has_tooltip bool) {
	C.gtk_widget_set_has_tooltip(widget.native(), C.gboolean(util.Bool2Int(has_tooltip)) /*go:.util*/)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (widget Widget) SetHasWindow(has_window bool) {
	C.gtk_widget_set_has_window(widget.native(), C.gboolean(util.Bool2Int(has_window)) /*go:.util*/)
}

// SetHexpand is a wrapper around gtk_widget_set_hexpand().
func (widget Widget) SetHexpand(expand bool) {
	C.gtk_widget_set_hexpand(widget.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// SetHexpandSet is a wrapper around gtk_widget_set_hexpand_set().
func (widget Widget) SetHexpandSet(set bool) {
	C.gtk_widget_set_hexpand_set(widget.native(), C.gboolean(util.Bool2Int(set)) /*go:.util*/)
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (widget Widget) SetMapped(mapped bool) {
	C.gtk_widget_set_mapped(widget.native(), C.gboolean(util.Bool2Int(mapped)) /*go:.util*/)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (widget Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(widget.native(), C.gint(margin))
}

// SetMarginEnd is a wrapper around gtk_widget_set_margin_end().
func (widget Widget) SetMarginEnd(margin int) {
	C.gtk_widget_set_margin_end(widget.native(), C.gint(margin))
}

// SetMarginStart is a wrapper around gtk_widget_set_margin_start().
func (widget Widget) SetMarginStart(margin int) {
	C.gtk_widget_set_margin_start(widget.native(), C.gint(margin))
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (widget Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(widget.native(), C.gint(margin))
}

// SetName is a wrapper around gtk_widget_set_name().
func (widget Widget) SetName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_widget_set_name(widget.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (widget Widget) SetNoShowAll(no_show_all bool) {
	C.gtk_widget_set_no_show_all(widget.native(), C.gboolean(util.Bool2Int(no_show_all)) /*go:.util*/)
}

// SetOpacity is a wrapper around gtk_widget_set_opacity().
func (widget Widget) SetOpacity(opacity float64) {
	C.gtk_widget_set_opacity(widget.native(), C.double(opacity))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (widget Widget) SetParent(parent Widget) {
	C.gtk_widget_set_parent(widget.native(), parent.native())
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (widget Widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(widget.native(), C.gboolean(util.Bool2Int(realized)) /*go:.util*/)
}

// SetReceivesDefault is a wrapper around gtk_widget_set_receives_default().
func (widget Widget) SetReceivesDefault(receives_default bool) {
	C.gtk_widget_set_receives_default(widget.native(), C.gboolean(util.Bool2Int(receives_default)) /*go:.util*/)
}

// SetRedrawOnAllocate is a wrapper around gtk_widget_set_redraw_on_allocate().
func (widget Widget) SetRedrawOnAllocate(redraw_on_allocate bool) {
	C.gtk_widget_set_redraw_on_allocate(widget.native(), C.gboolean(util.Bool2Int(redraw_on_allocate)) /*go:.util*/)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (widget Widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(widget.native(), C.gboolean(util.Bool2Int(sensitive)) /*go:.util*/)
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (widget Widget) SetSizeRequest(width int, height int) {
	C.gtk_widget_set_size_request(widget.native(), C.gint(width), C.gint(height))
}

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (widget Widget) SetStateFlags(flags StateFlags, clear bool) {
	C.gtk_widget_set_state_flags(widget.native(), C.GtkStateFlags(flags), C.gboolean(util.Bool2Int(clear)) /*go:.util*/)
}

// SetSupportMultidevice is a wrapper around gtk_widget_set_support_multidevice().
func (widget Widget) SetSupportMultidevice(support_multidevice bool) {
	C.gtk_widget_set_support_multidevice(widget.native(), C.gboolean(util.Bool2Int(support_multidevice)) /*go:.util*/)
}

// SetTooltipMarkup is a wrapper around gtk_widget_set_tooltip_markup().
func (widget Widget) SetTooltipMarkup(markup string) {
	markup0 := (*C.gchar)(C.CString(markup))
	C.gtk_widget_set_tooltip_markup(widget.native(), markup0)
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (widget Widget) SetTooltipText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_widget_set_tooltip_text(widget.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetTooltipWindow is a wrapper around gtk_widget_set_tooltip_window().
func (widget Widget) SetTooltipWindow(custom_window Window) {
	C.gtk_widget_set_tooltip_window(widget.native(), custom_window.native())
}

// SetValign is a wrapper around gtk_widget_set_valign().
func (widget Widget) SetValign(align Align) {
	C.gtk_widget_set_valign(widget.native(), C.GtkAlign(align))
}

// SetVexpand is a wrapper around gtk_widget_set_vexpand().
func (widget Widget) SetVexpand(expand bool) {
	C.gtk_widget_set_vexpand(widget.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// SetVexpandSet is a wrapper around gtk_widget_set_vexpand_set().
func (widget Widget) SetVexpandSet(set bool) {
	C.gtk_widget_set_vexpand_set(widget.native(), C.gboolean(util.Bool2Int(set)) /*go:.util*/)
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (widget Widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(widget.native(), C.gboolean(util.Bool2Int(visible)) /*go:.util*/)
}

// Show is a wrapper around gtk_widget_show().
func (widget Widget) Show() {
	C.gtk_widget_show(widget.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (widget Widget) ShowAll() {
	C.gtk_widget_show_all(widget.native())
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (widget Widget) ShowNow() {
	C.gtk_widget_show_now(widget.native())
}

// StyleGetProperty is a wrapper around gtk_widget_style_get_property().
func (widget Widget) StyleGetProperty(property_name string, value gobject.Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.gtk_widget_style_get_property(widget.native(), property_name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// ThawChildNotify is a wrapper around gtk_widget_thaw_child_notify().
func (widget Widget) ThawChildNotify() {
	C.gtk_widget_thaw_child_notify(widget.native())
}

// TranslateCoordinates is a wrapper around gtk_widget_translate_coordinates().
func (src_widget Widget) TranslateCoordinates(dest_widget Widget, src_x int, src_y int) (bool, int, int) {
	var dest_x0 C.gint
	var dest_y0 C.gint
	ret0 := C.gtk_widget_translate_coordinates(src_widget.native(), dest_widget.native(), C.gint(src_x), C.gint(src_y), &dest_x0, &dest_y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(dest_x0), int(dest_y0)
}

// TriggerTooltipQuery is a wrapper around gtk_widget_trigger_tooltip_query().
func (widget Widget) TriggerTooltipQuery() {
	C.gtk_widget_trigger_tooltip_query(widget.native())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (widget Widget) Unmap() {
	C.gtk_widget_unmap(widget.native())
}

// Unparent is a wrapper around gtk_widget_unparent().
func (widget Widget) Unparent() {
	C.gtk_widget_unparent(widget.native())
}

// Unrealize is a wrapper around gtk_widget_unrealize().
func (widget Widget) Unrealize() {
	C.gtk_widget_unrealize(widget.native())
}

// UnsetStateFlags is a wrapper around gtk_widget_unset_state_flags().
func (widget Widget) UnsetStateFlags(flags StateFlags) {
	C.gtk_widget_unset_state_flags(widget.native(), C.GtkStateFlags(flags))
}

// WidgetGetDefaultDirection is a wrapper around gtk_widget_get_default_direction().
func WidgetGetDefaultDirection() TextDirection {
	ret0 := C.gtk_widget_get_default_direction()
	return TextDirection(ret0)
}

// WidgetSetDefaultDirection is a wrapper around gtk_widget_set_default_direction().
func WidgetSetDefaultDirection(dir TextDirection) {
	C.gtk_widget_set_default_direction(C.GtkTextDirection(dir))
}

// Struct SettingsValue
type SettingsValue struct {
	Ptr unsafe.Pointer
}

func (v SettingsValue) native() *C.GtkSettingsValue {
	return (*C.GtkSettingsValue)(v.Ptr)
}
func wrapSettingsValue(p *C.GtkSettingsValue) SettingsValue {
	return SettingsValue{unsafe.Pointer(p)}
}
func WrapSettingsValue(p unsafe.Pointer) SettingsValue {
	return SettingsValue{p}
}
func (v SettingsValue) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSettingsValue(p unsafe.Pointer) interface{} {
	return WrapSettingsValue(p)
}

// Struct StockItem
type StockItem struct {
	Ptr unsafe.Pointer
}

func (v StockItem) native() *C.GtkStockItem {
	return (*C.GtkStockItem)(v.Ptr)
}
func wrapStockItem(p *C.GtkStockItem) StockItem {
	return StockItem{unsafe.Pointer(p)}
}
func WrapStockItem(p unsafe.Pointer) StockItem {
	return StockItem{p}
}
func (v StockItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStockItem(p unsafe.Pointer) interface{} {
	return WrapStockItem(p)
}

// Struct SymbolicColor
type SymbolicColor struct {
	Ptr unsafe.Pointer
}

func (v SymbolicColor) native() *C.GtkSymbolicColor {
	return (*C.GtkSymbolicColor)(v.Ptr)
}
func wrapSymbolicColor(p *C.GtkSymbolicColor) SymbolicColor {
	return SymbolicColor{unsafe.Pointer(p)}
}
func WrapSymbolicColor(p unsafe.Pointer) SymbolicColor {
	return SymbolicColor{p}
}
func (v SymbolicColor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSymbolicColor(p unsafe.Pointer) interface{} {
	return WrapSymbolicColor(p)
}

// Struct TableChild
type TableChild struct {
	Ptr unsafe.Pointer
}

func (v TableChild) native() *C.GtkTableChild {
	return (*C.GtkTableChild)(v.Ptr)
}
func wrapTableChild(p *C.GtkTableChild) TableChild {
	return TableChild{unsafe.Pointer(p)}
}
func WrapTableChild(p unsafe.Pointer) TableChild {
	return TableChild{p}
}
func (v TableChild) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTableChild(p unsafe.Pointer) interface{} {
	return WrapTableChild(p)
}

// Struct TableRowCol
type TableRowCol struct {
	Ptr unsafe.Pointer
}

func (v TableRowCol) native() *C.GtkTableRowCol {
	return (*C.GtkTableRowCol)(v.Ptr)
}
func wrapTableRowCol(p *C.GtkTableRowCol) TableRowCol {
	return TableRowCol{unsafe.Pointer(p)}
}
func WrapTableRowCol(p unsafe.Pointer) TableRowCol {
	return TableRowCol{p}
}
func (v TableRowCol) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTableRowCol(p unsafe.Pointer) interface{} {
	return WrapTableRowCol(p)
}

// Struct TargetEntry
type TargetEntry struct {
	Ptr unsafe.Pointer
}

func (v TargetEntry) native() *C.GtkTargetEntry {
	return (*C.GtkTargetEntry)(v.Ptr)
}
func wrapTargetEntry(p *C.GtkTargetEntry) TargetEntry {
	return TargetEntry{unsafe.Pointer(p)}
}
func WrapTargetEntry(p unsafe.Pointer) TargetEntry {
	return TargetEntry{p}
}
func (v TargetEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTargetEntry(p unsafe.Pointer) interface{} {
	return WrapTargetEntry(p)
}

// TargetEntryNew is a wrapper around gtk_target_entry_new().
func TargetEntryNew(target string, flags uint, info uint) TargetEntry {
	target0 := (*C.gchar)(C.CString(target))
	ret0 := C.gtk_target_entry_new(target0, C.guint(flags), C.guint(info))
	C.free(unsafe.Pointer(target0)) /*ch:<stdlib.h>*/
	return wrapTargetEntry(ret0)
}

// Copy is a wrapper around gtk_target_entry_copy().
func (data TargetEntry) Copy() TargetEntry {
	ret0 := C.gtk_target_entry_copy(data.native())
	return wrapTargetEntry(ret0)
}

// Free is a wrapper around gtk_target_entry_free().
func (data TargetEntry) Free() {
	C.gtk_target_entry_free(data.native())
}

// Struct TargetList
type TargetList struct {
	Ptr unsafe.Pointer
}

func (v TargetList) native() *C.GtkTargetList {
	return (*C.GtkTargetList)(v.Ptr)
}
func wrapTargetList(p *C.GtkTargetList) TargetList {
	return TargetList{unsafe.Pointer(p)}
}
func WrapTargetList(p unsafe.Pointer) TargetList {
	return TargetList{p}
}
func (v TargetList) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTargetList(p unsafe.Pointer) interface{} {
	return WrapTargetList(p)
}

// AddImageTargets is a wrapper around gtk_target_list_add_image_targets().
func (list TargetList) AddImageTargets(info uint, writable bool) {
	C.gtk_target_list_add_image_targets(list.native(), C.guint(info), C.gboolean(util.Bool2Int(writable)) /*go:.util*/)
}

// AddRichTextTargets is a wrapper around gtk_target_list_add_rich_text_targets().
func (list TargetList) AddRichTextTargets(info uint, deserializable bool, buffer TextBuffer) {
	C.gtk_target_list_add_rich_text_targets(list.native(), C.guint(info), C.gboolean(util.Bool2Int(deserializable)) /*go:.util*/, buffer.native())
}

// AddTextTargets is a wrapper around gtk_target_list_add_text_targets().
func (list TargetList) AddTextTargets(info uint) {
	C.gtk_target_list_add_text_targets(list.native(), C.guint(info))
}

// AddUriTargets is a wrapper around gtk_target_list_add_uri_targets().
func (list TargetList) AddUriTargets(info uint) {
	C.gtk_target_list_add_uri_targets(list.native(), C.guint(info))
}

// Ref is a wrapper around gtk_target_list_ref().
func (list TargetList) Ref() TargetList {
	ret0 := C.gtk_target_list_ref(list.native())
	return wrapTargetList(ret0)
}

// Unref is a wrapper around gtk_target_list_unref().
func (list TargetList) Unref() {
	C.gtk_target_list_unref(list.native())
}

// Struct TargetPair
type TargetPair struct {
	Ptr unsafe.Pointer
}

func (v TargetPair) native() *C.GtkTargetPair {
	return (*C.GtkTargetPair)(v.Ptr)
}
func wrapTargetPair(p *C.GtkTargetPair) TargetPair {
	return TargetPair{unsafe.Pointer(p)}
}
func WrapTargetPair(p unsafe.Pointer) TargetPair {
	return TargetPair{p}
}
func (v TargetPair) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTargetPair(p unsafe.Pointer) interface{} {
	return WrapTargetPair(p)
}

// Struct TextAppearance
type TextAppearance struct {
	Ptr unsafe.Pointer
}

func (v TextAppearance) native() *C.GtkTextAppearance {
	return (*C.GtkTextAppearance)(v.Ptr)
}
func wrapTextAppearance(p *C.GtkTextAppearance) TextAppearance {
	return TextAppearance{unsafe.Pointer(p)}
}
func WrapTextAppearance(p unsafe.Pointer) TextAppearance {
	return TextAppearance{p}
}
func (v TextAppearance) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextAppearance(p unsafe.Pointer) interface{} {
	return WrapTextAppearance(p)
}

// Struct TextAttributes
type TextAttributes struct {
	Ptr unsafe.Pointer
}

func (v TextAttributes) native() *C.GtkTextAttributes {
	return (*C.GtkTextAttributes)(v.Ptr)
}
func wrapTextAttributes(p *C.GtkTextAttributes) TextAttributes {
	return TextAttributes{unsafe.Pointer(p)}
}
func WrapTextAttributes(p unsafe.Pointer) TextAttributes {
	return TextAttributes{p}
}
func (v TextAttributes) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextAttributes(p unsafe.Pointer) interface{} {
	return WrapTextAttributes(p)
}

// TextAttributesNew is a wrapper around gtk_text_attributes_new().
func TextAttributesNew() TextAttributes {
	ret0 := C.gtk_text_attributes_new()
	return wrapTextAttributes(ret0)
}

// Copy is a wrapper around gtk_text_attributes_copy().
func (src TextAttributes) Copy() TextAttributes {
	ret0 := C.gtk_text_attributes_copy(src.native())
	return wrapTextAttributes(ret0)
}

// CopyValues is a wrapper around gtk_text_attributes_copy_values().
func (src TextAttributes) CopyValues(dest TextAttributes) {
	C.gtk_text_attributes_copy_values(src.native(), dest.native())
}

// Ref is a wrapper around gtk_text_attributes_ref().
func (values TextAttributes) Ref() TextAttributes {
	ret0 := C.gtk_text_attributes_ref(values.native())
	return wrapTextAttributes(ret0)
}

// Unref is a wrapper around gtk_text_attributes_unref().
func (values TextAttributes) Unref() {
	C.gtk_text_attributes_unref(values.native())
}

// Struct TextIter
type TextIter struct {
	Ptr unsafe.Pointer
}

func (v TextIter) native() *C.GtkTextIter {
	return (*C.GtkTextIter)(v.Ptr)
}
func wrapTextIter(p *C.GtkTextIter) TextIter {
	return TextIter{unsafe.Pointer(p)}
}
func WrapTextIter(p unsafe.Pointer) TextIter {
	return TextIter{p}
}
func (v TextIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextIter(p unsafe.Pointer) interface{} {
	return WrapTextIter(p)
}

// Assign is a wrapper around gtk_text_iter_assign().
func (iter TextIter) Assign(other TextIter) {
	C.gtk_text_iter_assign(iter.native(), other.native())
}

// BackwardChar is a wrapper around gtk_text_iter_backward_char().
func (iter TextIter) BackwardChar() bool {
	ret0 := C.gtk_text_iter_backward_char(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardChars is a wrapper around gtk_text_iter_backward_chars().
func (iter TextIter) BackwardChars(count int) bool {
	ret0 := C.gtk_text_iter_backward_chars(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardCursorPosition is a wrapper around gtk_text_iter_backward_cursor_position().
func (iter TextIter) BackwardCursorPosition() bool {
	ret0 := C.gtk_text_iter_backward_cursor_position(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardCursorPositions is a wrapper around gtk_text_iter_backward_cursor_positions().
func (iter TextIter) BackwardCursorPositions(count int) bool {
	ret0 := C.gtk_text_iter_backward_cursor_positions(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardLine is a wrapper around gtk_text_iter_backward_line().
func (iter TextIter) BackwardLine() bool {
	ret0 := C.gtk_text_iter_backward_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardLines is a wrapper around gtk_text_iter_backward_lines().
func (iter TextIter) BackwardLines(count int) bool {
	ret0 := C.gtk_text_iter_backward_lines(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardSentenceStart is a wrapper around gtk_text_iter_backward_sentence_start().
func (iter TextIter) BackwardSentenceStart() bool {
	ret0 := C.gtk_text_iter_backward_sentence_start(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardSentenceStarts is a wrapper around gtk_text_iter_backward_sentence_starts().
func (iter TextIter) BackwardSentenceStarts(count int) bool {
	ret0 := C.gtk_text_iter_backward_sentence_starts(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardToTagToggle is a wrapper around gtk_text_iter_backward_to_tag_toggle().
func (iter TextIter) BackwardToTagToggle(tag TextTag) bool {
	ret0 := C.gtk_text_iter_backward_to_tag_toggle(iter.native(), tag.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardVisibleCursorPosition is a wrapper around gtk_text_iter_backward_visible_cursor_position().
func (iter TextIter) BackwardVisibleCursorPosition() bool {
	ret0 := C.gtk_text_iter_backward_visible_cursor_position(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardVisibleCursorPositions is a wrapper around gtk_text_iter_backward_visible_cursor_positions().
func (iter TextIter) BackwardVisibleCursorPositions(count int) bool {
	ret0 := C.gtk_text_iter_backward_visible_cursor_positions(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardVisibleLine is a wrapper around gtk_text_iter_backward_visible_line().
func (iter TextIter) BackwardVisibleLine() bool {
	ret0 := C.gtk_text_iter_backward_visible_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardVisibleLines is a wrapper around gtk_text_iter_backward_visible_lines().
func (iter TextIter) BackwardVisibleLines(count int) bool {
	ret0 := C.gtk_text_iter_backward_visible_lines(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardVisibleWordStart is a wrapper around gtk_text_iter_backward_visible_word_start().
func (iter TextIter) BackwardVisibleWordStart() bool {
	ret0 := C.gtk_text_iter_backward_visible_word_start(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardVisibleWordStarts is a wrapper around gtk_text_iter_backward_visible_word_starts().
func (iter TextIter) BackwardVisibleWordStarts(count int) bool {
	ret0 := C.gtk_text_iter_backward_visible_word_starts(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardWordStart is a wrapper around gtk_text_iter_backward_word_start().
func (iter TextIter) BackwardWordStart() bool {
	ret0 := C.gtk_text_iter_backward_word_start(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardWordStarts is a wrapper around gtk_text_iter_backward_word_starts().
func (iter TextIter) BackwardWordStarts(count int) bool {
	ret0 := C.gtk_text_iter_backward_word_starts(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanInsert is a wrapper around gtk_text_iter_can_insert().
func (iter TextIter) CanInsert(default_editability bool) bool {
	ret0 := C.gtk_text_iter_can_insert(iter.native(), C.gboolean(util.Bool2Int(default_editability)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Compare is a wrapper around gtk_text_iter_compare().
func (lhs TextIter) Compare(rhs TextIter) int {
	ret0 := C.gtk_text_iter_compare(lhs.native(), rhs.native())
	return int(ret0)
}

// Copy is a wrapper around gtk_text_iter_copy().
func (iter TextIter) Copy() TextIter {
	ret0 := C.gtk_text_iter_copy(iter.native())
	return wrapTextIter(ret0)
}

// Editable is a wrapper around gtk_text_iter_editable().
func (iter TextIter) Editable(default_setting bool) bool {
	ret0 := C.gtk_text_iter_editable(iter.native(), C.gboolean(util.Bool2Int(default_setting)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EndsLine is a wrapper around gtk_text_iter_ends_line().
func (iter TextIter) EndsLine() bool {
	ret0 := C.gtk_text_iter_ends_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EndsSentence is a wrapper around gtk_text_iter_ends_sentence().
func (iter TextIter) EndsSentence() bool {
	ret0 := C.gtk_text_iter_ends_sentence(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EndsTag is a wrapper around gtk_text_iter_ends_tag().
func (iter TextIter) EndsTag(tag TextTag) bool {
	ret0 := C.gtk_text_iter_ends_tag(iter.native(), tag.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EndsWord is a wrapper around gtk_text_iter_ends_word().
func (iter TextIter) EndsWord() bool {
	ret0 := C.gtk_text_iter_ends_word(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Equal is a wrapper around gtk_text_iter_equal().
func (lhs TextIter) Equal(rhs TextIter) bool {
	ret0 := C.gtk_text_iter_equal(lhs.native(), rhs.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardChar is a wrapper around gtk_text_iter_forward_char().
func (iter TextIter) ForwardChar() bool {
	ret0 := C.gtk_text_iter_forward_char(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardChars is a wrapper around gtk_text_iter_forward_chars().
func (iter TextIter) ForwardChars(count int) bool {
	ret0 := C.gtk_text_iter_forward_chars(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardCursorPosition is a wrapper around gtk_text_iter_forward_cursor_position().
func (iter TextIter) ForwardCursorPosition() bool {
	ret0 := C.gtk_text_iter_forward_cursor_position(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardCursorPositions is a wrapper around gtk_text_iter_forward_cursor_positions().
func (iter TextIter) ForwardCursorPositions(count int) bool {
	ret0 := C.gtk_text_iter_forward_cursor_positions(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardLine is a wrapper around gtk_text_iter_forward_line().
func (iter TextIter) ForwardLine() bool {
	ret0 := C.gtk_text_iter_forward_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardLines is a wrapper around gtk_text_iter_forward_lines().
func (iter TextIter) ForwardLines(count int) bool {
	ret0 := C.gtk_text_iter_forward_lines(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardSentenceEnd is a wrapper around gtk_text_iter_forward_sentence_end().
func (iter TextIter) ForwardSentenceEnd() bool {
	ret0 := C.gtk_text_iter_forward_sentence_end(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardSentenceEnds is a wrapper around gtk_text_iter_forward_sentence_ends().
func (iter TextIter) ForwardSentenceEnds(count int) bool {
	ret0 := C.gtk_text_iter_forward_sentence_ends(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardToEnd is a wrapper around gtk_text_iter_forward_to_end().
func (iter TextIter) ForwardToEnd() {
	C.gtk_text_iter_forward_to_end(iter.native())
}

// ForwardToLineEnd is a wrapper around gtk_text_iter_forward_to_line_end().
func (iter TextIter) ForwardToLineEnd() bool {
	ret0 := C.gtk_text_iter_forward_to_line_end(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardToTagToggle is a wrapper around gtk_text_iter_forward_to_tag_toggle().
func (iter TextIter) ForwardToTagToggle(tag TextTag) bool {
	ret0 := C.gtk_text_iter_forward_to_tag_toggle(iter.native(), tag.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardVisibleCursorPosition is a wrapper around gtk_text_iter_forward_visible_cursor_position().
func (iter TextIter) ForwardVisibleCursorPosition() bool {
	ret0 := C.gtk_text_iter_forward_visible_cursor_position(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardVisibleCursorPositions is a wrapper around gtk_text_iter_forward_visible_cursor_positions().
func (iter TextIter) ForwardVisibleCursorPositions(count int) bool {
	ret0 := C.gtk_text_iter_forward_visible_cursor_positions(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardVisibleLine is a wrapper around gtk_text_iter_forward_visible_line().
func (iter TextIter) ForwardVisibleLine() bool {
	ret0 := C.gtk_text_iter_forward_visible_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardVisibleLines is a wrapper around gtk_text_iter_forward_visible_lines().
func (iter TextIter) ForwardVisibleLines(count int) bool {
	ret0 := C.gtk_text_iter_forward_visible_lines(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardVisibleWordEnd is a wrapper around gtk_text_iter_forward_visible_word_end().
func (iter TextIter) ForwardVisibleWordEnd() bool {
	ret0 := C.gtk_text_iter_forward_visible_word_end(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardVisibleWordEnds is a wrapper around gtk_text_iter_forward_visible_word_ends().
func (iter TextIter) ForwardVisibleWordEnds(count int) bool {
	ret0 := C.gtk_text_iter_forward_visible_word_ends(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardWordEnd is a wrapper around gtk_text_iter_forward_word_end().
func (iter TextIter) ForwardWordEnd() bool {
	ret0 := C.gtk_text_iter_forward_word_end(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardWordEnds is a wrapper around gtk_text_iter_forward_word_ends().
func (iter TextIter) ForwardWordEnds(count int) bool {
	ret0 := C.gtk_text_iter_forward_word_ends(iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Free is a wrapper around gtk_text_iter_free().
func (iter TextIter) Free() {
	C.gtk_text_iter_free(iter.native())
}

// GetBuffer is a wrapper around gtk_text_iter_get_buffer().
func (iter TextIter) GetBuffer() TextBuffer {
	ret0 := C.gtk_text_iter_get_buffer(iter.native())
	return wrapTextBuffer(ret0)
}

// GetBytesInLine is a wrapper around gtk_text_iter_get_bytes_in_line().
func (iter TextIter) GetBytesInLine() int {
	ret0 := C.gtk_text_iter_get_bytes_in_line(iter.native())
	return int(ret0)
}

// GetCharsInLine is a wrapper around gtk_text_iter_get_chars_in_line().
func (iter TextIter) GetCharsInLine() int {
	ret0 := C.gtk_text_iter_get_chars_in_line(iter.native())
	return int(ret0)
}

// GetChildAnchor is a wrapper around gtk_text_iter_get_child_anchor().
func (iter TextIter) GetChildAnchor() TextChildAnchor {
	ret0 := C.gtk_text_iter_get_child_anchor(iter.native())
	return wrapTextChildAnchor(ret0)
}

// GetLanguage is a wrapper around gtk_text_iter_get_language().
func (iter TextIter) GetLanguage() pango.Language {
	ret0 := C.gtk_text_iter_get_language(iter.native())
	return pango.WrapLanguage(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetLine is a wrapper around gtk_text_iter_get_line().
func (iter TextIter) GetLine() int {
	ret0 := C.gtk_text_iter_get_line(iter.native())
	return int(ret0)
}

// GetLineIndex is a wrapper around gtk_text_iter_get_line_index().
func (iter TextIter) GetLineIndex() int {
	ret0 := C.gtk_text_iter_get_line_index(iter.native())
	return int(ret0)
}

// GetLineOffset is a wrapper around gtk_text_iter_get_line_offset().
func (iter TextIter) GetLineOffset() int {
	ret0 := C.gtk_text_iter_get_line_offset(iter.native())
	return int(ret0)
}

// GetOffset is a wrapper around gtk_text_iter_get_offset().
func (iter TextIter) GetOffset() int {
	ret0 := C.gtk_text_iter_get_offset(iter.native())
	return int(ret0)
}

// GetPixbuf is a wrapper around gtk_text_iter_get_pixbuf().
func (iter TextIter) GetPixbuf() gdkpixbuf.Pixbuf {
	ret0 := C.gtk_text_iter_get_pixbuf(iter.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetSlice is a wrapper around gtk_text_iter_get_slice().
func (start TextIter) GetSlice(end TextIter) string {
	ret0 := C.gtk_text_iter_get_slice(start.native(), end.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetText is a wrapper around gtk_text_iter_get_text().
func (start TextIter) GetText(end TextIter) string {
	ret0 := C.gtk_text_iter_get_text(start.native(), end.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetVisibleLineIndex is a wrapper around gtk_text_iter_get_visible_line_index().
func (iter TextIter) GetVisibleLineIndex() int {
	ret0 := C.gtk_text_iter_get_visible_line_index(iter.native())
	return int(ret0)
}

// GetVisibleLineOffset is a wrapper around gtk_text_iter_get_visible_line_offset().
func (iter TextIter) GetVisibleLineOffset() int {
	ret0 := C.gtk_text_iter_get_visible_line_offset(iter.native())
	return int(ret0)
}

// GetVisibleSlice is a wrapper around gtk_text_iter_get_visible_slice().
func (start TextIter) GetVisibleSlice(end TextIter) string {
	ret0 := C.gtk_text_iter_get_visible_slice(start.native(), end.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetVisibleText is a wrapper around gtk_text_iter_get_visible_text().
func (start TextIter) GetVisibleText(end TextIter) string {
	ret0 := C.gtk_text_iter_get_visible_text(start.native(), end.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// HasTag is a wrapper around gtk_text_iter_has_tag().
func (iter TextIter) HasTag(tag TextTag) bool {
	ret0 := C.gtk_text_iter_has_tag(iter.native(), tag.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InRange is a wrapper around gtk_text_iter_in_range().
func (iter TextIter) InRange(start TextIter, end TextIter) bool {
	ret0 := C.gtk_text_iter_in_range(iter.native(), start.native(), end.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InsideSentence is a wrapper around gtk_text_iter_inside_sentence().
func (iter TextIter) InsideSentence() bool {
	ret0 := C.gtk_text_iter_inside_sentence(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InsideWord is a wrapper around gtk_text_iter_inside_word().
func (iter TextIter) InsideWord() bool {
	ret0 := C.gtk_text_iter_inside_word(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsCursorPosition is a wrapper around gtk_text_iter_is_cursor_position().
func (iter TextIter) IsCursorPosition() bool {
	ret0 := C.gtk_text_iter_is_cursor_position(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsEnd is a wrapper around gtk_text_iter_is_end().
func (iter TextIter) IsEnd() bool {
	ret0 := C.gtk_text_iter_is_end(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsStart is a wrapper around gtk_text_iter_is_start().
func (iter TextIter) IsStart() bool {
	ret0 := C.gtk_text_iter_is_start(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Order is a wrapper around gtk_text_iter_order().
func (first TextIter) Order(second TextIter) {
	C.gtk_text_iter_order(first.native(), second.native())
}

// SetLine is a wrapper around gtk_text_iter_set_line().
func (iter TextIter) SetLine(line_number int) {
	C.gtk_text_iter_set_line(iter.native(), C.gint(line_number))
}

// SetLineIndex is a wrapper around gtk_text_iter_set_line_index().
func (iter TextIter) SetLineIndex(byte_on_line int) {
	C.gtk_text_iter_set_line_index(iter.native(), C.gint(byte_on_line))
}

// SetLineOffset is a wrapper around gtk_text_iter_set_line_offset().
func (iter TextIter) SetLineOffset(char_on_line int) {
	C.gtk_text_iter_set_line_offset(iter.native(), C.gint(char_on_line))
}

// SetOffset is a wrapper around gtk_text_iter_set_offset().
func (iter TextIter) SetOffset(char_offset int) {
	C.gtk_text_iter_set_offset(iter.native(), C.gint(char_offset))
}

// SetVisibleLineIndex is a wrapper around gtk_text_iter_set_visible_line_index().
func (iter TextIter) SetVisibleLineIndex(byte_on_line int) {
	C.gtk_text_iter_set_visible_line_index(iter.native(), C.gint(byte_on_line))
}

// SetVisibleLineOffset is a wrapper around gtk_text_iter_set_visible_line_offset().
func (iter TextIter) SetVisibleLineOffset(char_on_line int) {
	C.gtk_text_iter_set_visible_line_offset(iter.native(), C.gint(char_on_line))
}

// StartsLine is a wrapper around gtk_text_iter_starts_line().
func (iter TextIter) StartsLine() bool {
	ret0 := C.gtk_text_iter_starts_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// StartsSentence is a wrapper around gtk_text_iter_starts_sentence().
func (iter TextIter) StartsSentence() bool {
	ret0 := C.gtk_text_iter_starts_sentence(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// StartsTag is a wrapper around gtk_text_iter_starts_tag().
func (iter TextIter) StartsTag(tag TextTag) bool {
	ret0 := C.gtk_text_iter_starts_tag(iter.native(), tag.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// StartsWord is a wrapper around gtk_text_iter_starts_word().
func (iter TextIter) StartsWord() bool {
	ret0 := C.gtk_text_iter_starts_word(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TogglesTag is a wrapper around gtk_text_iter_toggles_tag().
func (iter TextIter) TogglesTag(tag TextTag) bool {
	ret0 := C.gtk_text_iter_toggles_tag(iter.native(), tag.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object TextTag
type TextTag struct {
	gobject.Object
}

func (v TextTag) native() *C.GtkTextTag {
	return (*C.GtkTextTag)(v.Ptr)
}
func wrapTextTag(p *C.GtkTextTag) (v TextTag) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextTag(p unsafe.Pointer) (v TextTag) {
	v.Ptr = p
	return
}
func (v TextTag) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextTag(p unsafe.Pointer) interface{} {
	return WrapTextTag(p)
}
func (v TextTag) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_tag_get_type())
}
func (v TextTag) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextTag(unsafe.Pointer(ptr)), nil
	}
}

// TextTagNew is a wrapper around gtk_text_tag_new().
func TextTagNew(name string) TextTag {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_text_tag_new(name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapTextTag(ret0)
}

// Changed is a wrapper around gtk_text_tag_changed().
func (tag TextTag) Changed(size_changed bool) {
	C.gtk_text_tag_changed(tag.native(), C.gboolean(util.Bool2Int(size_changed)) /*go:.util*/)
}

// GetPriority is a wrapper around gtk_text_tag_get_priority().
func (tag TextTag) GetPriority() int {
	ret0 := C.gtk_text_tag_get_priority(tag.native())
	return int(ret0)
}

// SetPriority is a wrapper around gtk_text_tag_set_priority().
func (tag TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(tag.native(), C.gint(priority))
}

// Object TextChildAnchor
type TextChildAnchor struct {
	gobject.Object
}

func (v TextChildAnchor) native() *C.GtkTextChildAnchor {
	return (*C.GtkTextChildAnchor)(v.Ptr)
}
func wrapTextChildAnchor(p *C.GtkTextChildAnchor) (v TextChildAnchor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextChildAnchor(p unsafe.Pointer) (v TextChildAnchor) {
	v.Ptr = p
	return
}
func (v TextChildAnchor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextChildAnchor(p unsafe.Pointer) interface{} {
	return WrapTextChildAnchor(p)
}
func (v TextChildAnchor) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_child_anchor_get_type())
}
func (v TextChildAnchor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextChildAnchor(unsafe.Pointer(ptr)), nil
	}
}

// TextChildAnchorNew is a wrapper around gtk_text_child_anchor_new().
func TextChildAnchorNew() TextChildAnchor {
	ret0 := C.gtk_text_child_anchor_new()
	return wrapTextChildAnchor(ret0)
}

// GetDeleted is a wrapper around gtk_text_child_anchor_get_deleted().
func (anchor TextChildAnchor) GetDeleted() bool {
	ret0 := C.gtk_text_child_anchor_get_deleted(anchor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetWidgets is a wrapper around gtk_text_child_anchor_get_widgets().
func (anchor TextChildAnchor) GetWidgets() glib.List {
	ret0 := C.gtk_text_child_anchor_get_widgets(anchor.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWidget(p) }) /*gir:GLib*/
}

// Struct ToggleActionEntry
type ToggleActionEntry struct {
	Ptr unsafe.Pointer
}

func (v ToggleActionEntry) native() *C.GtkToggleActionEntry {
	return (*C.GtkToggleActionEntry)(v.Ptr)
}
func wrapToggleActionEntry(p *C.GtkToggleActionEntry) ToggleActionEntry {
	return ToggleActionEntry{unsafe.Pointer(p)}
}
func WrapToggleActionEntry(p unsafe.Pointer) ToggleActionEntry {
	return ToggleActionEntry{p}
}
func (v ToggleActionEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToggleActionEntry(p unsafe.Pointer) interface{} {
	return WrapToggleActionEntry(p)
}

// Struct TreeIter
type TreeIter struct {
	Ptr unsafe.Pointer
}

func (v TreeIter) native() *C.GtkTreeIter {
	return (*C.GtkTreeIter)(v.Ptr)
}
func wrapTreeIter(p *C.GtkTreeIter) TreeIter {
	return TreeIter{unsafe.Pointer(p)}
}
func WrapTreeIter(p unsafe.Pointer) TreeIter {
	return TreeIter{p}
}
func (v TreeIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeIter(p unsafe.Pointer) interface{} {
	return WrapTreeIter(p)
}

// Copy is a wrapper around gtk_tree_iter_copy().
func (iter TreeIter) Copy() TreeIter {
	ret0 := C.gtk_tree_iter_copy(iter.native())
	return wrapTreeIter(ret0)
}

// Free is a wrapper around gtk_tree_iter_free().
func (iter TreeIter) Free() {
	C.gtk_tree_iter_free(iter.native())
}

// Struct TreePath
type TreePath struct {
	Ptr unsafe.Pointer
}

func (v TreePath) native() *C.GtkTreePath {
	return (*C.GtkTreePath)(v.Ptr)
}
func wrapTreePath(p *C.GtkTreePath) TreePath {
	return TreePath{unsafe.Pointer(p)}
}
func WrapTreePath(p unsafe.Pointer) TreePath {
	return TreePath{p}
}
func (v TreePath) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreePath(p unsafe.Pointer) interface{} {
	return WrapTreePath(p)
}

// TreePathNew is a wrapper around gtk_tree_path_new().
func TreePathNew() TreePath {
	ret0 := C.gtk_tree_path_new()
	return wrapTreePath(ret0)
}

// TreePathNewFirst is a wrapper around gtk_tree_path_new_first().
func TreePathNewFirst() TreePath {
	ret0 := C.gtk_tree_path_new_first()
	return wrapTreePath(ret0)
}

// gtk_tree_path_new_from_indices shadowed by new_from_indicesv

// TreePathNewFromIndices is a wrapper around gtk_tree_path_new_from_indicesv().
func TreePathNewFromIndices(indices []int) TreePath {
	indices0 := make([]C.gint, len(indices))
	for idx, elemG := range indices {
		indices0[idx] = C.gint(elemG)
	}
	var indices0Ptr *C.gint
	if len(indices0) > 0 {
		indices0Ptr = &indices0[0]
	}
	ret0 := C.gtk_tree_path_new_from_indicesv(indices0Ptr, C.gsize(len(indices)))
	return wrapTreePath(ret0)
}

// TreePathNewFromString is a wrapper around gtk_tree_path_new_from_string().
func TreePathNewFromString(path string) TreePath {
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.gtk_tree_path_new_from_string(path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	return wrapTreePath(ret0)
}

// AppendIndex is a wrapper around gtk_tree_path_append_index().
func (path TreePath) AppendIndex(index_ int) {
	C.gtk_tree_path_append_index(path.native(), C.gint(index_))
}

// Compare is a wrapper around gtk_tree_path_compare().
func (a TreePath) Compare(b TreePath) int {
	ret0 := C.gtk_tree_path_compare(a.native(), b.native())
	return int(ret0)
}

// Copy is a wrapper around gtk_tree_path_copy().
func (path TreePath) Copy() TreePath {
	ret0 := C.gtk_tree_path_copy(path.native())
	return wrapTreePath(ret0)
}

// Down is a wrapper around gtk_tree_path_down().
func (path TreePath) Down() {
	C.gtk_tree_path_down(path.native())
}

// Free is a wrapper around gtk_tree_path_free().
func (path TreePath) Free() {
	C.gtk_tree_path_free(path.native())
}

// GetDepth is a wrapper around gtk_tree_path_get_depth().
func (path TreePath) GetDepth() int {
	ret0 := C.gtk_tree_path_get_depth(path.native())
	return int(ret0)
}

// gtk_tree_path_get_indices shadowed by get_indices_with_depth

// GetIndices is a wrapper around gtk_tree_path_get_indices_with_depth().
func (path TreePath) GetIndices() []int {
	var depth0 C.gint
	ret0 := C.gtk_tree_path_get_indices_with_depth(path.native(), &depth0)
	var ret0Slice []C.gint
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(depth0)) /*go:.util*/
	ret := make([]int, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = int(elem)
	}
	return ret
}

// IsAncestor is a wrapper around gtk_tree_path_is_ancestor().
func (path TreePath) IsAncestor(descendant TreePath) bool {
	ret0 := C.gtk_tree_path_is_ancestor(path.native(), descendant.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsDescendant is a wrapper around gtk_tree_path_is_descendant().
func (path TreePath) IsDescendant(ancestor TreePath) bool {
	ret0 := C.gtk_tree_path_is_descendant(path.native(), ancestor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Next is a wrapper around gtk_tree_path_next().
func (path TreePath) Next() {
	C.gtk_tree_path_next(path.native())
}

// PrependIndex is a wrapper around gtk_tree_path_prepend_index().
func (path TreePath) PrependIndex(index_ int) {
	C.gtk_tree_path_prepend_index(path.native(), C.gint(index_))
}

// Prev is a wrapper around gtk_tree_path_prev().
func (path TreePath) Prev() bool {
	ret0 := C.gtk_tree_path_prev(path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ToString is a wrapper around gtk_tree_path_to_string().
func (path TreePath) ToString() string {
	ret0 := C.gtk_tree_path_to_string(path.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Up is a wrapper around gtk_tree_path_up().
func (path TreePath) Up() bool {
	ret0 := C.gtk_tree_path_up(path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Struct TreeRowReference
type TreeRowReference struct {
	Ptr unsafe.Pointer
}

func (v TreeRowReference) native() *C.GtkTreeRowReference {
	return (*C.GtkTreeRowReference)(v.Ptr)
}
func wrapTreeRowReference(p *C.GtkTreeRowReference) TreeRowReference {
	return TreeRowReference{unsafe.Pointer(p)}
}
func WrapTreeRowReference(p unsafe.Pointer) TreeRowReference {
	return TreeRowReference{p}
}
func (v TreeRowReference) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeRowReference(p unsafe.Pointer) interface{} {
	return WrapTreeRowReference(p)
}

// TreeRowReferenceNew is a wrapper around gtk_tree_row_reference_new().
func TreeRowReferenceNew(model TreeModel, path TreePath) TreeRowReference {
	ret0 := C.gtk_tree_row_reference_new(model.native(), path.native())
	return wrapTreeRowReference(ret0)
}

// TreeRowReferenceNewProxy is a wrapper around gtk_tree_row_reference_new_proxy().
func TreeRowReferenceNewProxy(proxy gobject.Object, model TreeModel, path TreePath) TreeRowReference {
	ret0 := C.gtk_tree_row_reference_new_proxy((*C.GObject)(proxy.Ptr), model.native(), path.native())
	return wrapTreeRowReference(ret0)
}

// Copy is a wrapper around gtk_tree_row_reference_copy().
func (reference TreeRowReference) Copy() TreeRowReference {
	ret0 := C.gtk_tree_row_reference_copy(reference.native())
	return wrapTreeRowReference(ret0)
}

// Free is a wrapper around gtk_tree_row_reference_free().
func (reference TreeRowReference) Free() {
	C.gtk_tree_row_reference_free(reference.native())
}

// GetModel is a wrapper around gtk_tree_row_reference_get_model().
func (reference TreeRowReference) GetModel() TreeModel {
	ret0 := C.gtk_tree_row_reference_get_model(reference.native())
	return wrapTreeModel(ret0)
}

// GetPath is a wrapper around gtk_tree_row_reference_get_path().
func (reference TreeRowReference) GetPath() TreePath {
	ret0 := C.gtk_tree_row_reference_get_path(reference.native())
	return wrapTreePath(ret0)
}

// Valid is a wrapper around gtk_tree_row_reference_valid().
func (reference TreeRowReference) Valid() bool {
	ret0 := C.gtk_tree_row_reference_valid(reference.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TreeRowReferenceDeleted is a wrapper around gtk_tree_row_reference_deleted().
func TreeRowReferenceDeleted(proxy gobject.Object, path TreePath) {
	C.gtk_tree_row_reference_deleted((*C.GObject)(proxy.Ptr), path.native())
}

// TreeRowReferenceInserted is a wrapper around gtk_tree_row_reference_inserted().
func TreeRowReferenceInserted(proxy gobject.Object, path TreePath) {
	C.gtk_tree_row_reference_inserted((*C.GObject)(proxy.Ptr), path.native())
}

// TreeRowReferenceReordered is a wrapper around gtk_tree_row_reference_reordered().
func TreeRowReferenceReordered(proxy gobject.Object, path TreePath, iter TreeIter, new_order []int) {
	new_order0 := make([]C.gint, len(new_order))
	for idx, elemG := range new_order {
		new_order0[idx] = C.gint(elemG)
	}
	var new_order0Ptr *C.gint
	if len(new_order0) > 0 {
		new_order0Ptr = &new_order0[0]
	}
	C.gtk_tree_row_reference_reordered((*C.GObject)(proxy.Ptr), path.native(), iter.native(), new_order0Ptr)
}

// Interface TreeModel
type TreeModel struct {
	Ptr unsafe.Pointer
}

func (v TreeModel) native() *C.GtkTreeModel {
	return (*C.GtkTreeModel)(v.Ptr)
}
func wrapTreeModel(p *C.GtkTreeModel) TreeModel {
	return TreeModel{unsafe.Pointer(p)}
}
func WrapTreeModel(p unsafe.Pointer) TreeModel {
	return TreeModel{p}
}
func (v TreeModel) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeModel(p unsafe.Pointer) interface{} {
	return WrapTreeModel(p)
}
func (v TreeModel) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_model_get_type())
}
func (v TreeModel) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeModel(unsafe.Pointer(ptr)), nil
	}
}

// FilterNew is a wrapper around gtk_tree_model_filter_new().
func (child_model TreeModel) FilterNew(root TreePath) TreeModel {
	ret0 := C.gtk_tree_model_filter_new(child_model.native(), root.native())
	return wrapTreeModel(ret0)
}

// GetColumnType is a wrapper around gtk_tree_model_get_column_type().
func (tree_model TreeModel) GetColumnType(index_ int) /*Gir:GObject*/ gobject.Type {
	ret0 := C.gtk_tree_model_get_column_type(tree_model.native(), C.gint(index_))
	return /*Gir:GObject*/ gobject.Type(ret0)
}

// GetFlags is a wrapper around gtk_tree_model_get_flags().
func (tree_model TreeModel) GetFlags() TreeModelFlags {
	ret0 := C.gtk_tree_model_get_flags(tree_model.native())
	return TreeModelFlags(ret0)
}

// GetNColumns is a wrapper around gtk_tree_model_get_n_columns().
func (tree_model TreeModel) GetNColumns() int {
	ret0 := C.gtk_tree_model_get_n_columns(tree_model.native())
	return int(ret0)
}

// GetPath is a wrapper around gtk_tree_model_get_path().
func (tree_model TreeModel) GetPath(iter TreeIter) TreePath {
	ret0 := C.gtk_tree_model_get_path(tree_model.native(), iter.native())
	return wrapTreePath(ret0)
}

// GetStringFromIter is a wrapper around gtk_tree_model_get_string_from_iter().
func (tree_model TreeModel) GetStringFromIter(iter TreeIter) string {
	ret0 := C.gtk_tree_model_get_string_from_iter(tree_model.native(), iter.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// IterHasChild is a wrapper around gtk_tree_model_iter_has_child().
func (tree_model TreeModel) IterHasChild(iter TreeIter) bool {
	ret0 := C.gtk_tree_model_iter_has_child(tree_model.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterNChildren is a wrapper around gtk_tree_model_iter_n_children().
func (tree_model TreeModel) IterNChildren(iter TreeIter) int {
	ret0 := C.gtk_tree_model_iter_n_children(tree_model.native(), iter.native())
	return int(ret0)
}

// IterNext is a wrapper around gtk_tree_model_iter_next().
func (tree_model TreeModel) IterNext(iter TreeIter) bool {
	ret0 := C.gtk_tree_model_iter_next(tree_model.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterPrevious is a wrapper around gtk_tree_model_iter_previous().
func (tree_model TreeModel) IterPrevious(iter TreeIter) bool {
	ret0 := C.gtk_tree_model_iter_previous(tree_model.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RefNode is a wrapper around gtk_tree_model_ref_node().
func (tree_model TreeModel) RefNode(iter TreeIter) {
	C.gtk_tree_model_ref_node(tree_model.native(), iter.native())
}

// RowChanged is a wrapper around gtk_tree_model_row_changed().
func (tree_model TreeModel) RowChanged(path TreePath, iter TreeIter) {
	C.gtk_tree_model_row_changed(tree_model.native(), path.native(), iter.native())
}

// RowDeleted is a wrapper around gtk_tree_model_row_deleted().
func (tree_model TreeModel) RowDeleted(path TreePath) {
	C.gtk_tree_model_row_deleted(tree_model.native(), path.native())
}

// RowHasChildToggled is a wrapper around gtk_tree_model_row_has_child_toggled().
func (tree_model TreeModel) RowHasChildToggled(path TreePath, iter TreeIter) {
	C.gtk_tree_model_row_has_child_toggled(tree_model.native(), path.native(), iter.native())
}

// RowInserted is a wrapper around gtk_tree_model_row_inserted().
func (tree_model TreeModel) RowInserted(path TreePath, iter TreeIter) {
	C.gtk_tree_model_row_inserted(tree_model.native(), path.native(), iter.native())
}

// gtk_tree_model_rows_reordered shadowed by rows_reordered_with_length

// RowsReordered is a wrapper around gtk_tree_model_rows_reordered_with_length().
func (tree_model TreeModel) RowsReordered(path TreePath, iter TreeIter, new_order []int) {
	new_order0 := make([]C.gint, len(new_order))
	for idx, elemG := range new_order {
		new_order0[idx] = C.gint(elemG)
	}
	var new_order0Ptr *C.gint
	if len(new_order0) > 0 {
		new_order0Ptr = &new_order0[0]
	}
	C.gtk_tree_model_rows_reordered_with_length(tree_model.native(), path.native(), iter.native(), new_order0Ptr, C.gint(len(new_order)))
}

// SortNewWithModel is a wrapper around gtk_tree_model_sort_new_with_model().
func (child_model TreeModel) SortNewWithModel() TreeModel {
	ret0 := C.gtk_tree_model_sort_new_with_model(child_model.native())
	return wrapTreeModel(ret0)
}

// UnrefNode is a wrapper around gtk_tree_model_unref_node().
func (tree_model TreeModel) UnrefNode(iter TreeIter) {
	C.gtk_tree_model_unref_node(tree_model.native(), iter.native())
}

// Struct WidgetPath
type WidgetPath struct {
	Ptr unsafe.Pointer
}

func (v WidgetPath) native() *C.GtkWidgetPath {
	return (*C.GtkWidgetPath)(v.Ptr)
}
func wrapWidgetPath(p *C.GtkWidgetPath) WidgetPath {
	return WidgetPath{unsafe.Pointer(p)}
}
func WrapWidgetPath(p unsafe.Pointer) WidgetPath {
	return WidgetPath{p}
}
func (v WidgetPath) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWidgetPath(p unsafe.Pointer) interface{} {
	return WrapWidgetPath(p)
}

// WidgetPathNew is a wrapper around gtk_widget_path_new().
func WidgetPathNew() WidgetPath {
	ret0 := C.gtk_widget_path_new()
	return wrapWidgetPath(ret0)
}

// AppendForWidget is a wrapper around gtk_widget_path_append_for_widget().
func (path WidgetPath) AppendForWidget(widget Widget) int {
	ret0 := C.gtk_widget_path_append_for_widget(path.native(), widget.native())
	return int(ret0)
}

// AppendType is a wrapper around gtk_widget_path_append_type().
func (path WidgetPath) AppendType(type_ /*Gir:GObject*/ gobject.Type) int {
	ret0 := C.gtk_widget_path_append_type(path.native(), C.GType(type_))
	return int(ret0)
}

// AppendWithSiblings is a wrapper around gtk_widget_path_append_with_siblings().
func (path WidgetPath) AppendWithSiblings(siblings WidgetPath, sibling_index uint) int {
	ret0 := C.gtk_widget_path_append_with_siblings(path.native(), siblings.native(), C.guint(sibling_index))
	return int(ret0)
}

// Copy is a wrapper around gtk_widget_path_copy().
func (path WidgetPath) Copy() WidgetPath {
	ret0 := C.gtk_widget_path_copy(path.native())
	return wrapWidgetPath(ret0)
}

// Free is a wrapper around gtk_widget_path_free().
func (path WidgetPath) Free() {
	C.gtk_widget_path_free(path.native())
}

// GetObjectType is a wrapper around gtk_widget_path_get_object_type().
func (path WidgetPath) GetObjectType() /*Gir:GObject*/ gobject.Type {
	ret0 := C.gtk_widget_path_get_object_type(path.native())
	return /*Gir:GObject*/ gobject.Type(ret0)
}

// HasParent is a wrapper around gtk_widget_path_has_parent().
func (path WidgetPath) HasParent(type_ /*Gir:GObject*/ gobject.Type) bool {
	ret0 := C.gtk_widget_path_has_parent(path.native(), C.GType(type_))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsType is a wrapper around gtk_widget_path_is_type().
func (path WidgetPath) IsType(type_ /*Gir:GObject*/ gobject.Type) bool {
	ret0 := C.gtk_widget_path_is_type(path.native(), C.GType(type_))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterAddClass is a wrapper around gtk_widget_path_iter_add_class().
func (path WidgetPath) IterAddClass(pos int, name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_widget_path_iter_add_class(path.native(), C.gint(pos), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// IterClearClasses is a wrapper around gtk_widget_path_iter_clear_classes().
func (path WidgetPath) IterClearClasses(pos int) {
	C.gtk_widget_path_iter_clear_classes(path.native(), C.gint(pos))
}

// IterGetName is a wrapper around gtk_widget_path_iter_get_name().
func (path WidgetPath) IterGetName(pos int) string {
	ret0 := C.gtk_widget_path_iter_get_name(path.native(), C.gint(pos))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// IterGetObjectName is a wrapper around gtk_widget_path_iter_get_object_name().
func (path WidgetPath) IterGetObjectName(pos int) string {
	ret0 := C.gtk_widget_path_iter_get_object_name(path.native(), C.gint(pos))
	ret := C.GoString(ret0)
	return ret
}

// IterGetObjectType is a wrapper around gtk_widget_path_iter_get_object_type().
func (path WidgetPath) IterGetObjectType(pos int) /*Gir:GObject*/ gobject.Type {
	ret0 := C.gtk_widget_path_iter_get_object_type(path.native(), C.gint(pos))
	return /*Gir:GObject*/ gobject.Type(ret0)
}

// IterGetSiblingIndex is a wrapper around gtk_widget_path_iter_get_sibling_index().
func (path WidgetPath) IterGetSiblingIndex(pos int) uint {
	ret0 := C.gtk_widget_path_iter_get_sibling_index(path.native(), C.gint(pos))
	return uint(ret0)
}

// IterGetSiblings is a wrapper around gtk_widget_path_iter_get_siblings().
func (path WidgetPath) IterGetSiblings(pos int) WidgetPath {
	ret0 := C.gtk_widget_path_iter_get_siblings(path.native(), C.gint(pos))
	return wrapWidgetPath(ret0)
}

// IterGetState is a wrapper around gtk_widget_path_iter_get_state().
func (path WidgetPath) IterGetState(pos int) StateFlags {
	ret0 := C.gtk_widget_path_iter_get_state(path.native(), C.gint(pos))
	return StateFlags(ret0)
}

// IterHasClass is a wrapper around gtk_widget_path_iter_has_class().
func (path WidgetPath) IterHasClass(pos int, name string) bool {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_widget_path_iter_has_class(path.native(), C.gint(pos), name0)
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterHasName is a wrapper around gtk_widget_path_iter_has_name().
func (path WidgetPath) IterHasName(pos int, name string) bool {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_widget_path_iter_has_name(path.native(), C.gint(pos), name0)
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterHasQclass is a wrapper around gtk_widget_path_iter_has_qclass().
func (path WidgetPath) IterHasQclass(pos int, qname /*gir:GLib*/ glib.Quark) bool {
	ret0 := C.gtk_widget_path_iter_has_qclass(path.native(), C.gint(pos), C.GQuark(qname))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterHasQname is a wrapper around gtk_widget_path_iter_has_qname().
func (path WidgetPath) IterHasQname(pos int, qname /*gir:GLib*/ glib.Quark) bool {
	ret0 := C.gtk_widget_path_iter_has_qname(path.native(), C.gint(pos), C.GQuark(qname))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterRemoveClass is a wrapper around gtk_widget_path_iter_remove_class().
func (path WidgetPath) IterRemoveClass(pos int, name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_widget_path_iter_remove_class(path.native(), C.gint(pos), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// IterSetName is a wrapper around gtk_widget_path_iter_set_name().
func (path WidgetPath) IterSetName(pos int, name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_widget_path_iter_set_name(path.native(), C.gint(pos), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// IterSetObjectName is a wrapper around gtk_widget_path_iter_set_object_name().
func (path WidgetPath) IterSetObjectName(pos int, name string) {
	name0 := C.CString(name)
	C.gtk_widget_path_iter_set_object_name(path.native(), C.gint(pos), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// IterSetObjectType is a wrapper around gtk_widget_path_iter_set_object_type().
func (path WidgetPath) IterSetObjectType(pos int, type_ /*Gir:GObject*/ gobject.Type) {
	C.gtk_widget_path_iter_set_object_type(path.native(), C.gint(pos), C.GType(type_))
}

// IterSetState is a wrapper around gtk_widget_path_iter_set_state().
func (path WidgetPath) IterSetState(pos int, state StateFlags) {
	C.gtk_widget_path_iter_set_state(path.native(), C.gint(pos), C.GtkStateFlags(state))
}

// Length is a wrapper around gtk_widget_path_length().
func (path WidgetPath) Length() int {
	ret0 := C.gtk_widget_path_length(path.native())
	return int(ret0)
}

// PrependType is a wrapper around gtk_widget_path_prepend_type().
func (path WidgetPath) PrependType(type_ /*Gir:GObject*/ gobject.Type) {
	C.gtk_widget_path_prepend_type(path.native(), C.GType(type_))
}

// Ref is a wrapper around gtk_widget_path_ref().
func (path WidgetPath) Ref() WidgetPath {
	ret0 := C.gtk_widget_path_ref(path.native())
	return wrapWidgetPath(ret0)
}

// ToString is a wrapper around gtk_widget_path_to_string().
func (path WidgetPath) ToString() string {
	ret0 := C.gtk_widget_path_to_string(path.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// Unref is a wrapper around gtk_widget_path_unref().
func (path WidgetPath) Unref() {
	C.gtk_widget_path_unref(path.native())
}

// Interface Actionable
type Actionable struct {
	Ptr unsafe.Pointer
}

func (v Actionable) native() *C.GtkActionable {
	return (*C.GtkActionable)(v.Ptr)
}
func wrapActionable(p *C.GtkActionable) Actionable {
	return Actionable{unsafe.Pointer(p)}
}
func WrapActionable(p unsafe.Pointer) Actionable {
	return Actionable{p}
}
func (v Actionable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapActionable(p unsafe.Pointer) interface{} {
	return WrapActionable(p)
}
func (v Actionable) GetType() gobject.Type {
	return gobject.Type(C.gtk_actionable_get_type())
}
func (v Actionable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapActionable(unsafe.Pointer(ptr)), nil
	}
}

// GetActionName is a wrapper around gtk_actionable_get_action_name().
func (actionable Actionable) GetActionName() string {
	ret0 := C.gtk_actionable_get_action_name(actionable.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetActionTargetValue is a wrapper around gtk_actionable_get_action_target_value().
func (actionable Actionable) GetActionTargetValue() glib.Variant {
	ret0 := C.gtk_actionable_get_action_target_value(actionable.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// SetActionName is a wrapper around gtk_actionable_set_action_name().
func (actionable Actionable) SetActionName(action_name string) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.gtk_actionable_set_action_name(actionable.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// SetActionTargetValue is a wrapper around gtk_actionable_set_action_target_value().
func (actionable Actionable) SetActionTargetValue(target_value glib.Variant) {
	C.gtk_actionable_set_action_target_value(actionable.native(), (*C.GVariant)(target_value.Ptr))
}

// SetDetailedActionName is a wrapper around gtk_actionable_set_detailed_action_name().
func (actionable Actionable) SetDetailedActionName(detailed_action_name string) {
	detailed_action_name0 := (*C.gchar)(C.CString(detailed_action_name))
	C.gtk_actionable_set_detailed_action_name(actionable.native(), detailed_action_name0)
	C.free(unsafe.Pointer(detailed_action_name0)) /*ch:<stdlib.h>*/
}

// Interface AppChooser
type AppChooser struct {
	Ptr unsafe.Pointer
}

func (v AppChooser) native() *C.GtkAppChooser {
	return (*C.GtkAppChooser)(v.Ptr)
}
func wrapAppChooser(p *C.GtkAppChooser) AppChooser {
	return AppChooser{unsafe.Pointer(p)}
}
func WrapAppChooser(p unsafe.Pointer) AppChooser {
	return AppChooser{p}
}
func (v AppChooser) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppChooser(p unsafe.Pointer) interface{} {
	return WrapAppChooser(p)
}
func (v AppChooser) GetType() gobject.Type {
	return gobject.Type(C.gtk_app_chooser_get_type())
}
func (v AppChooser) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppChooser(unsafe.Pointer(ptr)), nil
	}
}

// GetAppInfo is a wrapper around gtk_app_chooser_get_app_info().
func (self AppChooser) GetAppInfo() gio.AppInfo {
	ret0 := C.gtk_app_chooser_get_app_info(self.native())
	return gio.WrapAppInfo(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetContentType is a wrapper around gtk_app_chooser_get_content_type().
func (self AppChooser) GetContentType() string {
	ret0 := C.gtk_app_chooser_get_content_type(self.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Refresh is a wrapper around gtk_app_chooser_refresh().
func (self AppChooser) Refresh() {
	C.gtk_app_chooser_refresh(self.native())
}

// Interface CellEditable
type CellEditable struct {
	Ptr unsafe.Pointer
}

func (v CellEditable) native() *C.GtkCellEditable {
	return (*C.GtkCellEditable)(v.Ptr)
}
func wrapCellEditable(p *C.GtkCellEditable) CellEditable {
	return CellEditable{unsafe.Pointer(p)}
}
func WrapCellEditable(p unsafe.Pointer) CellEditable {
	return CellEditable{p}
}
func (v CellEditable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellEditable(p unsafe.Pointer) interface{} {
	return WrapCellEditable(p)
}
func (v CellEditable) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_editable_get_type())
}
func (v CellEditable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellEditable(unsafe.Pointer(ptr)), nil
	}
}

// EditingDone is a wrapper around gtk_cell_editable_editing_done().
func (cell_editable CellEditable) EditingDone() {
	C.gtk_cell_editable_editing_done(cell_editable.native())
}

// RemoveWidget is a wrapper around gtk_cell_editable_remove_widget().
func (cell_editable CellEditable) RemoveWidget() {
	C.gtk_cell_editable_remove_widget(cell_editable.native())
}

// Interface CellLayout
type CellLayout struct {
	Ptr unsafe.Pointer
}

func (v CellLayout) native() *C.GtkCellLayout {
	return (*C.GtkCellLayout)(v.Ptr)
}
func wrapCellLayout(p *C.GtkCellLayout) CellLayout {
	return CellLayout{unsafe.Pointer(p)}
}
func WrapCellLayout(p unsafe.Pointer) CellLayout {
	return CellLayout{p}
}
func (v CellLayout) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellLayout(p unsafe.Pointer) interface{} {
	return WrapCellLayout(p)
}
func (v CellLayout) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_layout_get_type())
}
func (v CellLayout) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellLayout(unsafe.Pointer(ptr)), nil
	}
}

// AddAttribute is a wrapper around gtk_cell_layout_add_attribute().
func (cell_layout CellLayout) AddAttribute(cell CellRenderer, attribute string, column int) {
	attribute0 := (*C.gchar)(C.CString(attribute))
	C.gtk_cell_layout_add_attribute(cell_layout.native(), cell.native(), attribute0, C.gint(column))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// Clear is a wrapper around gtk_cell_layout_clear().
func (cell_layout CellLayout) Clear() {
	C.gtk_cell_layout_clear(cell_layout.native())
}

// ClearAttributes is a wrapper around gtk_cell_layout_clear_attributes().
func (cell_layout CellLayout) ClearAttributes(cell CellRenderer) {
	C.gtk_cell_layout_clear_attributes(cell_layout.native(), cell.native())
}

// GetArea is a wrapper around gtk_cell_layout_get_area().
func (cell_layout CellLayout) GetArea() CellArea {
	ret0 := C.gtk_cell_layout_get_area(cell_layout.native())
	return wrapCellArea(ret0)
}

// GetCells is a wrapper around gtk_cell_layout_get_cells().
func (cell_layout CellLayout) GetCells() glib.List {
	ret0 := C.gtk_cell_layout_get_cells(cell_layout.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapCellRenderer(p) }) /*gir:GLib*/
}

// PackEnd is a wrapper around gtk_cell_layout_pack_end().
func (cell_layout CellLayout) PackEnd(cell CellRenderer, expand bool) {
	C.gtk_cell_layout_pack_end(cell_layout.native(), cell.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// PackStart is a wrapper around gtk_cell_layout_pack_start().
func (cell_layout CellLayout) PackStart(cell CellRenderer, expand bool) {
	C.gtk_cell_layout_pack_start(cell_layout.native(), cell.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// Reorder is a wrapper around gtk_cell_layout_reorder().
func (cell_layout CellLayout) Reorder(cell CellRenderer, position int) {
	C.gtk_cell_layout_reorder(cell_layout.native(), cell.native(), C.gint(position))
}

// Object CellRenderer
type CellRenderer struct {
	gobject.InitiallyUnowned
}

func (v CellRenderer) native() *C.GtkCellRenderer {
	return (*C.GtkCellRenderer)(v.Ptr)
}
func wrapCellRenderer(p *C.GtkCellRenderer) (v CellRenderer) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRenderer(p unsafe.Pointer) (v CellRenderer) {
	v.Ptr = p
	return
}
func (v CellRenderer) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRenderer(p unsafe.Pointer) interface{} {
	return WrapCellRenderer(p)
}
func (v CellRenderer) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_get_type())
}
func (v CellRenderer) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRenderer(unsafe.Pointer(ptr)), nil
	}
}

// GetAlignment is a wrapper around gtk_cell_renderer_get_alignment().
func (cell CellRenderer) GetAlignment() (float32, float32) {
	var xalign0 C.gfloat
	var yalign0 C.gfloat
	C.gtk_cell_renderer_get_alignment(cell.native(), &xalign0, &yalign0)
	return float32(xalign0), float32(yalign0)
}

// GetFixedSize is a wrapper around gtk_cell_renderer_get_fixed_size().
func (cell CellRenderer) GetFixedSize() (int, int) {
	var width0 C.gint
	var height0 C.gint
	C.gtk_cell_renderer_get_fixed_size(cell.native(), &width0, &height0)
	return int(width0), int(height0)
}

// GetPadding is a wrapper around gtk_cell_renderer_get_padding().
func (cell CellRenderer) GetPadding() (int, int) {
	var xpad0 C.gint
	var ypad0 C.gint
	C.gtk_cell_renderer_get_padding(cell.native(), &xpad0, &ypad0)
	return int(xpad0), int(ypad0)
}

// GetPreferredHeight is a wrapper around gtk_cell_renderer_get_preferred_height().
func (cell CellRenderer) GetPreferredHeight(widget Widget) (int, int) {
	var minimum_size0 C.gint
	var natural_size0 C.gint
	C.gtk_cell_renderer_get_preferred_height(cell.native(), widget.native(), &minimum_size0, &natural_size0)
	return int(minimum_size0), int(natural_size0)
}

// GetPreferredHeightForWidth is a wrapper around gtk_cell_renderer_get_preferred_height_for_width().
func (cell CellRenderer) GetPreferredHeightForWidth(widget Widget, width int) (int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	C.gtk_cell_renderer_get_preferred_height_for_width(cell.native(), widget.native(), C.gint(width), &minimum_height0, &natural_height0)
	return int(minimum_height0), int(natural_height0)
}

// GetPreferredWidth is a wrapper around gtk_cell_renderer_get_preferred_width().
func (cell CellRenderer) GetPreferredWidth(widget Widget) (int, int) {
	var minimum_size0 C.gint
	var natural_size0 C.gint
	C.gtk_cell_renderer_get_preferred_width(cell.native(), widget.native(), &minimum_size0, &natural_size0)
	return int(minimum_size0), int(natural_size0)
}

// GetPreferredWidthForHeight is a wrapper around gtk_cell_renderer_get_preferred_width_for_height().
func (cell CellRenderer) GetPreferredWidthForHeight(widget Widget, height int) (int, int) {
	var minimum_width0 C.gint
	var natural_width0 C.gint
	C.gtk_cell_renderer_get_preferred_width_for_height(cell.native(), widget.native(), C.gint(height), &minimum_width0, &natural_width0)
	return int(minimum_width0), int(natural_width0)
}

// GetRequestMode is a wrapper around gtk_cell_renderer_get_request_mode().
func (cell CellRenderer) GetRequestMode() SizeRequestMode {
	ret0 := C.gtk_cell_renderer_get_request_mode(cell.native())
	return SizeRequestMode(ret0)
}

// GetSensitive is a wrapper around gtk_cell_renderer_get_sensitive().
func (cell CellRenderer) GetSensitive() bool {
	ret0 := C.gtk_cell_renderer_get_sensitive(cell.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetState is a wrapper around gtk_cell_renderer_get_state().
func (cell CellRenderer) GetState(widget Widget, cell_state CellRendererState) StateFlags {
	ret0 := C.gtk_cell_renderer_get_state(cell.native(), widget.native(), C.GtkCellRendererState(cell_state))
	return StateFlags(ret0)
}

// GetVisible is a wrapper around gtk_cell_renderer_get_visible().
func (cell CellRenderer) GetVisible() bool {
	ret0 := C.gtk_cell_renderer_get_visible(cell.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsActivatable is a wrapper around gtk_cell_renderer_is_activatable().
func (cell CellRenderer) IsActivatable() bool {
	ret0 := C.gtk_cell_renderer_is_activatable(cell.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetAlignment is a wrapper around gtk_cell_renderer_set_alignment().
func (cell CellRenderer) SetAlignment(xalign float32, yalign float32) {
	C.gtk_cell_renderer_set_alignment(cell.native(), C.gfloat(xalign), C.gfloat(yalign))
}

// SetFixedSize is a wrapper around gtk_cell_renderer_set_fixed_size().
func (cell CellRenderer) SetFixedSize(width int, height int) {
	C.gtk_cell_renderer_set_fixed_size(cell.native(), C.gint(width), C.gint(height))
}

// SetPadding is a wrapper around gtk_cell_renderer_set_padding().
func (cell CellRenderer) SetPadding(xpad int, ypad int) {
	C.gtk_cell_renderer_set_padding(cell.native(), C.gint(xpad), C.gint(ypad))
}

// SetSensitive is a wrapper around gtk_cell_renderer_set_sensitive().
func (cell CellRenderer) SetSensitive(sensitive bool) {
	C.gtk_cell_renderer_set_sensitive(cell.native(), C.gboolean(util.Bool2Int(sensitive)) /*go:.util*/)
}

// SetVisible is a wrapper around gtk_cell_renderer_set_visible().
func (cell CellRenderer) SetVisible(visible bool) {
	C.gtk_cell_renderer_set_visible(cell.native(), C.gboolean(util.Bool2Int(visible)) /*go:.util*/)
}

// StopEditing is a wrapper around gtk_cell_renderer_stop_editing().
func (cell CellRenderer) StopEditing(canceled bool) {
	C.gtk_cell_renderer_stop_editing(cell.native(), C.gboolean(util.Bool2Int(canceled)) /*go:.util*/)
}

// Object CellArea
type CellArea struct {
	gobject.InitiallyUnowned
}

func (v CellArea) native() *C.GtkCellArea {
	return (*C.GtkCellArea)(v.Ptr)
}
func wrapCellArea(p *C.GtkCellArea) (v CellArea) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellArea(p unsafe.Pointer) (v CellArea) {
	v.Ptr = p
	return
}
func (v CellArea) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellArea(p unsafe.Pointer) interface{} {
	return WrapCellArea(p)
}
func (v CellArea) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_area_get_type())
}
func (v CellArea) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellArea(unsafe.Pointer(ptr)), nil
	}
}
func (v CellArea) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v CellArea) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}

// Activate is a wrapper around gtk_cell_area_activate().
func (area CellArea) Activate(context CellAreaContext, widget Widget, cell_area gdk.Rectangle, flags CellRendererState, edit_only bool) bool {
	ret0 := C.gtk_cell_area_activate(area.native(), context.native(), widget.native(), (*C.GdkRectangle)(cell_area.Ptr), C.GtkCellRendererState(flags), C.gboolean(util.Bool2Int(edit_only)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Add is a wrapper around gtk_cell_area_add().
func (area CellArea) Add(renderer CellRenderer) {
	C.gtk_cell_area_add(area.native(), renderer.native())
}

// AddFocusSibling is a wrapper around gtk_cell_area_add_focus_sibling().
func (area CellArea) AddFocusSibling(renderer CellRenderer, sibling CellRenderer) {
	C.gtk_cell_area_add_focus_sibling(area.native(), renderer.native(), sibling.native())
}

// ApplyAttributes is a wrapper around gtk_cell_area_apply_attributes().
func (area CellArea) ApplyAttributes(tree_model TreeModel, iter TreeIter, is_expander bool, is_expanded bool) {
	C.gtk_cell_area_apply_attributes(area.native(), tree_model.native(), iter.native(), C.gboolean(util.Bool2Int(is_expander)) /*go:.util*/, C.gboolean(util.Bool2Int(is_expanded)) /*go:.util*/)
}

// AttributeConnect is a wrapper around gtk_cell_area_attribute_connect().
func (area CellArea) AttributeConnect(renderer CellRenderer, attribute string, column int) {
	attribute0 := (*C.gchar)(C.CString(attribute))
	C.gtk_cell_area_attribute_connect(area.native(), renderer.native(), attribute0, C.gint(column))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// AttributeDisconnect is a wrapper around gtk_cell_area_attribute_disconnect().
func (area CellArea) AttributeDisconnect(renderer CellRenderer, attribute string) {
	attribute0 := (*C.gchar)(C.CString(attribute))
	C.gtk_cell_area_attribute_disconnect(area.native(), renderer.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// AttributeGetColumn is a wrapper around gtk_cell_area_attribute_get_column().
func (area CellArea) AttributeGetColumn(renderer CellRenderer, attribute string) int {
	attribute0 := (*C.gchar)(C.CString(attribute))
	ret0 := C.gtk_cell_area_attribute_get_column(area.native(), renderer.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// CellGetProperty is a wrapper around gtk_cell_area_cell_get_property().
func (area CellArea) CellGetProperty(renderer CellRenderer, property_name string, value gobject.Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.gtk_cell_area_cell_get_property(area.native(), renderer.native(), property_name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// CellSetProperty is a wrapper around gtk_cell_area_cell_set_property().
func (area CellArea) CellSetProperty(renderer CellRenderer, property_name string, value gobject.Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.gtk_cell_area_cell_set_property(area.native(), renderer.native(), property_name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// CopyContext is a wrapper around gtk_cell_area_copy_context().
func (area CellArea) CopyContext(context CellAreaContext) CellAreaContext {
	ret0 := C.gtk_cell_area_copy_context(area.native(), context.native())
	return wrapCellAreaContext(ret0)
}

// CreateContext is a wrapper around gtk_cell_area_create_context().
func (area CellArea) CreateContext() CellAreaContext {
	ret0 := C.gtk_cell_area_create_context(area.native())
	return wrapCellAreaContext(ret0)
}

// Focus is a wrapper around gtk_cell_area_focus().
func (area CellArea) Focus(direction DirectionType) bool {
	ret0 := C.gtk_cell_area_focus(area.native(), C.GtkDirectionType(direction))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCurrentPathString is a wrapper around gtk_cell_area_get_current_path_string().
func (area CellArea) GetCurrentPathString() string {
	ret0 := C.gtk_cell_area_get_current_path_string(area.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetEditWidget is a wrapper around gtk_cell_area_get_edit_widget().
func (area CellArea) GetEditWidget() CellEditable {
	ret0 := C.gtk_cell_area_get_edit_widget(area.native())
	return wrapCellEditable(ret0)
}

// GetEditedCell is a wrapper around gtk_cell_area_get_edited_cell().
func (area CellArea) GetEditedCell() CellRenderer {
	ret0 := C.gtk_cell_area_get_edited_cell(area.native())
	return wrapCellRenderer(ret0)
}

// GetFocusCell is a wrapper around gtk_cell_area_get_focus_cell().
func (area CellArea) GetFocusCell() CellRenderer {
	ret0 := C.gtk_cell_area_get_focus_cell(area.native())
	return wrapCellRenderer(ret0)
}

// GetFocusFromSibling is a wrapper around gtk_cell_area_get_focus_from_sibling().
func (area CellArea) GetFocusFromSibling(renderer CellRenderer) CellRenderer {
	ret0 := C.gtk_cell_area_get_focus_from_sibling(area.native(), renderer.native())
	return wrapCellRenderer(ret0)
}

// GetFocusSiblings is a wrapper around gtk_cell_area_get_focus_siblings().
func (area CellArea) GetFocusSiblings(renderer CellRenderer) glib.List {
	ret0 := C.gtk_cell_area_get_focus_siblings(area.native(), renderer.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapCellRenderer(p) }) /*gir:GLib*/
}

// GetPreferredHeight is a wrapper around gtk_cell_area_get_preferred_height().
func (area CellArea) GetPreferredHeight(context CellAreaContext, widget Widget) (int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	C.gtk_cell_area_get_preferred_height(area.native(), context.native(), widget.native(), &minimum_height0, &natural_height0)
	return int(minimum_height0), int(natural_height0)
}

// GetPreferredHeightForWidth is a wrapper around gtk_cell_area_get_preferred_height_for_width().
func (area CellArea) GetPreferredHeightForWidth(context CellAreaContext, widget Widget, width int) (int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	C.gtk_cell_area_get_preferred_height_for_width(area.native(), context.native(), widget.native(), C.gint(width), &minimum_height0, &natural_height0)
	return int(minimum_height0), int(natural_height0)
}

// GetPreferredWidth is a wrapper around gtk_cell_area_get_preferred_width().
func (area CellArea) GetPreferredWidth(context CellAreaContext, widget Widget) (int, int) {
	var minimum_width0 C.gint
	var natural_width0 C.gint
	C.gtk_cell_area_get_preferred_width(area.native(), context.native(), widget.native(), &minimum_width0, &natural_width0)
	return int(minimum_width0), int(natural_width0)
}

// GetPreferredWidthForHeight is a wrapper around gtk_cell_area_get_preferred_width_for_height().
func (area CellArea) GetPreferredWidthForHeight(context CellAreaContext, widget Widget, height int) (int, int) {
	var minimum_width0 C.gint
	var natural_width0 C.gint
	C.gtk_cell_area_get_preferred_width_for_height(area.native(), context.native(), widget.native(), C.gint(height), &minimum_width0, &natural_width0)
	return int(minimum_width0), int(natural_width0)
}

// GetRequestMode is a wrapper around gtk_cell_area_get_request_mode().
func (area CellArea) GetRequestMode() SizeRequestMode {
	ret0 := C.gtk_cell_area_get_request_mode(area.native())
	return SizeRequestMode(ret0)
}

// HasRenderer is a wrapper around gtk_cell_area_has_renderer().
func (area CellArea) HasRenderer(renderer CellRenderer) bool {
	ret0 := C.gtk_cell_area_has_renderer(area.native(), renderer.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsActivatable is a wrapper around gtk_cell_area_is_activatable().
func (area CellArea) IsActivatable() bool {
	ret0 := C.gtk_cell_area_is_activatable(area.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsFocusSibling is a wrapper around gtk_cell_area_is_focus_sibling().
func (area CellArea) IsFocusSibling(renderer CellRenderer, sibling CellRenderer) bool {
	ret0 := C.gtk_cell_area_is_focus_sibling(area.native(), renderer.native(), sibling.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Remove is a wrapper around gtk_cell_area_remove().
func (area CellArea) Remove(renderer CellRenderer) {
	C.gtk_cell_area_remove(area.native(), renderer.native())
}

// RemoveFocusSibling is a wrapper around gtk_cell_area_remove_focus_sibling().
func (area CellArea) RemoveFocusSibling(renderer CellRenderer, sibling CellRenderer) {
	C.gtk_cell_area_remove_focus_sibling(area.native(), renderer.native(), sibling.native())
}

// RequestRenderer is a wrapper around gtk_cell_area_request_renderer().
func (area CellArea) RequestRenderer(renderer CellRenderer, orientation Orientation, widget Widget, for_size int) (int, int) {
	var minimum_size0 C.gint
	var natural_size0 C.gint
	C.gtk_cell_area_request_renderer(area.native(), renderer.native(), C.GtkOrientation(orientation), widget.native(), C.gint(for_size), &minimum_size0, &natural_size0)
	return int(minimum_size0), int(natural_size0)
}

// SetFocusCell is a wrapper around gtk_cell_area_set_focus_cell().
func (area CellArea) SetFocusCell(renderer CellRenderer) {
	C.gtk_cell_area_set_focus_cell(area.native(), renderer.native())
}

// StopEditing is a wrapper around gtk_cell_area_stop_editing().
func (area CellArea) StopEditing(canceled bool) {
	C.gtk_cell_area_stop_editing(area.native(), C.gboolean(util.Bool2Int(canceled)) /*go:.util*/)
}

// Object CellAreaContext
type CellAreaContext struct {
	gobject.Object
}

func (v CellAreaContext) native() *C.GtkCellAreaContext {
	return (*C.GtkCellAreaContext)(v.Ptr)
}
func wrapCellAreaContext(p *C.GtkCellAreaContext) (v CellAreaContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellAreaContext(p unsafe.Pointer) (v CellAreaContext) {
	v.Ptr = p
	return
}
func (v CellAreaContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellAreaContext(p unsafe.Pointer) interface{} {
	return WrapCellAreaContext(p)
}
func (v CellAreaContext) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_area_context_get_type())
}
func (v CellAreaContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellAreaContext(unsafe.Pointer(ptr)), nil
	}
}

// Allocate is a wrapper around gtk_cell_area_context_allocate().
func (context CellAreaContext) Allocate(width int, height int) {
	C.gtk_cell_area_context_allocate(context.native(), C.gint(width), C.gint(height))
}

// GetAllocation is a wrapper around gtk_cell_area_context_get_allocation().
func (context CellAreaContext) GetAllocation() (int, int) {
	var width0 C.gint
	var height0 C.gint
	C.gtk_cell_area_context_get_allocation(context.native(), &width0, &height0)
	return int(width0), int(height0)
}

// GetArea is a wrapper around gtk_cell_area_context_get_area().
func (context CellAreaContext) GetArea() CellArea {
	ret0 := C.gtk_cell_area_context_get_area(context.native())
	return wrapCellArea(ret0)
}

// GetPreferredHeight is a wrapper around gtk_cell_area_context_get_preferred_height().
func (context CellAreaContext) GetPreferredHeight() (int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	C.gtk_cell_area_context_get_preferred_height(context.native(), &minimum_height0, &natural_height0)
	return int(minimum_height0), int(natural_height0)
}

// GetPreferredHeightForWidth is a wrapper around gtk_cell_area_context_get_preferred_height_for_width().
func (context CellAreaContext) GetPreferredHeightForWidth(width int) (int, int) {
	var minimum_height0 C.gint
	var natural_height0 C.gint
	C.gtk_cell_area_context_get_preferred_height_for_width(context.native(), C.gint(width), &minimum_height0, &natural_height0)
	return int(minimum_height0), int(natural_height0)
}

// GetPreferredWidth is a wrapper around gtk_cell_area_context_get_preferred_width().
func (context CellAreaContext) GetPreferredWidth() (int, int) {
	var minimum_width0 C.gint
	var natural_width0 C.gint
	C.gtk_cell_area_context_get_preferred_width(context.native(), &minimum_width0, &natural_width0)
	return int(minimum_width0), int(natural_width0)
}

// GetPreferredWidthForHeight is a wrapper around gtk_cell_area_context_get_preferred_width_for_height().
func (context CellAreaContext) GetPreferredWidthForHeight(height int) (int, int) {
	var minimum_width0 C.gint
	var natural_width0 C.gint
	C.gtk_cell_area_context_get_preferred_width_for_height(context.native(), C.gint(height), &minimum_width0, &natural_width0)
	return int(minimum_width0), int(natural_width0)
}

// PushPreferredHeight is a wrapper around gtk_cell_area_context_push_preferred_height().
func (context CellAreaContext) PushPreferredHeight(minimum_height int, natural_height int) {
	C.gtk_cell_area_context_push_preferred_height(context.native(), C.gint(minimum_height), C.gint(natural_height))
}

// PushPreferredWidth is a wrapper around gtk_cell_area_context_push_preferred_width().
func (context CellAreaContext) PushPreferredWidth(minimum_width int, natural_width int) {
	C.gtk_cell_area_context_push_preferred_width(context.native(), C.gint(minimum_width), C.gint(natural_width))
}

// Reset is a wrapper around gtk_cell_area_context_reset().
func (context CellAreaContext) Reset() {
	C.gtk_cell_area_context_reset(context.native())
}

// Interface ColorChooser
type ColorChooser struct {
	Ptr unsafe.Pointer
}

func (v ColorChooser) native() *C.GtkColorChooser {
	return (*C.GtkColorChooser)(v.Ptr)
}
func wrapColorChooser(p *C.GtkColorChooser) ColorChooser {
	return ColorChooser{unsafe.Pointer(p)}
}
func WrapColorChooser(p unsafe.Pointer) ColorChooser {
	return ColorChooser{p}
}
func (v ColorChooser) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColorChooser(p unsafe.Pointer) interface{} {
	return WrapColorChooser(p)
}
func (v ColorChooser) GetType() gobject.Type {
	return gobject.Type(C.gtk_color_chooser_get_type())
}
func (v ColorChooser) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapColorChooser(unsafe.Pointer(ptr)), nil
	}
}

// GetUseAlpha is a wrapper around gtk_color_chooser_get_use_alpha().
func (chooser ColorChooser) GetUseAlpha() bool {
	ret0 := C.gtk_color_chooser_get_use_alpha(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUseAlpha is a wrapper around gtk_color_chooser_set_use_alpha().
func (chooser ColorChooser) SetUseAlpha(use_alpha bool) {
	C.gtk_color_chooser_set_use_alpha(chooser.native(), C.gboolean(util.Bool2Int(use_alpha)) /*go:.util*/)
}

// Interface Editable
type Editable struct {
	Ptr unsafe.Pointer
}

func (v Editable) native() *C.GtkEditable {
	return (*C.GtkEditable)(v.Ptr)
}
func wrapEditable(p *C.GtkEditable) Editable {
	return Editable{unsafe.Pointer(p)}
}
func WrapEditable(p unsafe.Pointer) Editable {
	return Editable{p}
}
func (v Editable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEditable(p unsafe.Pointer) interface{} {
	return WrapEditable(p)
}
func (v Editable) GetType() gobject.Type {
	return gobject.Type(C.gtk_editable_get_type())
}
func (v Editable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEditable(unsafe.Pointer(ptr)), nil
	}
}

// CopyClipboard is a wrapper around gtk_editable_copy_clipboard().
func (editable Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(editable.native())
}

// CutClipboard is a wrapper around gtk_editable_cut_clipboard().
func (editable Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard(editable.native())
}

// DeleteSelection is a wrapper around gtk_editable_delete_selection().
func (editable Editable) DeleteSelection() {
	C.gtk_editable_delete_selection(editable.native())
}

// DeleteText is a wrapper around gtk_editable_delete_text().
func (editable Editable) DeleteText(start_pos int, end_pos int) {
	C.gtk_editable_delete_text(editable.native(), C.gint(start_pos), C.gint(end_pos))
}

// GetChars is a wrapper around gtk_editable_get_chars().
func (editable Editable) GetChars(start_pos int, end_pos int) string {
	ret0 := C.gtk_editable_get_chars(editable.native(), C.gint(start_pos), C.gint(end_pos))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetEditable is a wrapper around gtk_editable_get_editable().
func (editable Editable) GetEditable() bool {
	ret0 := C.gtk_editable_get_editable(editable.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPosition is a wrapper around gtk_editable_get_position().
func (editable Editable) GetPosition() int {
	ret0 := C.gtk_editable_get_position(editable.native())
	return int(ret0)
}

// GetSelectionBounds is a wrapper around gtk_editable_get_selection_bounds().
func (editable Editable) GetSelectionBounds() (bool, int, int) {
	var start_pos0 C.gint
	var end_pos0 C.gint
	ret0 := C.gtk_editable_get_selection_bounds(editable.native(), &start_pos0, &end_pos0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(start_pos0), int(end_pos0)
}

// PasteClipboard is a wrapper around gtk_editable_paste_clipboard().
func (editable Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(editable.native())
}

// SelectRegion is a wrapper around gtk_editable_select_region().
func (editable Editable) SelectRegion(start_pos int, end_pos int) {
	C.gtk_editable_select_region(editable.native(), C.gint(start_pos), C.gint(end_pos))
}

// SetEditable is a wrapper around gtk_editable_set_editable().
func (editable Editable) SetEditable(is_editable bool) {
	C.gtk_editable_set_editable(editable.native(), C.gboolean(util.Bool2Int(is_editable)) /*go:.util*/)
}

// SetPosition is a wrapper around gtk_editable_set_position().
func (editable Editable) SetPosition(position int) {
	C.gtk_editable_set_position(editable.native(), C.gint(position))
}

// Interface FileChooser
type FileChooser struct {
	Ptr unsafe.Pointer
}

func (v FileChooser) native() *C.GtkFileChooser {
	return (*C.GtkFileChooser)(v.Ptr)
}
func wrapFileChooser(p *C.GtkFileChooser) FileChooser {
	return FileChooser{unsafe.Pointer(p)}
}
func WrapFileChooser(p unsafe.Pointer) FileChooser {
	return FileChooser{p}
}
func (v FileChooser) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileChooser(p unsafe.Pointer) interface{} {
	return WrapFileChooser(p)
}
func (v FileChooser) GetType() gobject.Type {
	return gobject.Type(C.gtk_file_chooser_get_type())
}
func (v FileChooser) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileChooser(unsafe.Pointer(ptr)), nil
	}
}

// AddFilter is a wrapper around gtk_file_chooser_add_filter().
func (chooser FileChooser) AddFilter(filter FileFilter) {
	C.gtk_file_chooser_add_filter(chooser.native(), filter.native())
}

// AddShortcutFolder is a wrapper around gtk_file_chooser_add_shortcut_folder().
func (chooser FileChooser) AddShortcutFolder(folder string) (bool, error) {
	folder0 := C.CString(folder)
	var err glib.Error
	ret0 := C.gtk_file_chooser_add_shortcut_folder(chooser.native(), folder0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(folder0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// AddShortcutFolderUri is a wrapper around gtk_file_chooser_add_shortcut_folder_uri().
func (chooser FileChooser) AddShortcutFolderUri(uri string) (bool, error) {
	uri0 := C.CString(uri)
	var err glib.Error
	ret0 := C.gtk_file_chooser_add_shortcut_folder_uri(chooser.native(), uri0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetAction is a wrapper around gtk_file_chooser_get_action().
func (chooser FileChooser) GetAction() FileChooserAction {
	ret0 := C.gtk_file_chooser_get_action(chooser.native())
	return FileChooserAction(ret0)
}

// GetChoice is a wrapper around gtk_file_chooser_get_choice().
func (chooser FileChooser) GetChoice(id string) string {
	id0 := C.CString(id)
	ret0 := C.gtk_file_chooser_get_choice(chooser.native(), id0)
	C.free(unsafe.Pointer(id0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	return ret
}

// GetCreateFolders is a wrapper around gtk_file_chooser_get_create_folders().
func (chooser FileChooser) GetCreateFolders() bool {
	ret0 := C.gtk_file_chooser_get_create_folders(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCurrentFolder is a wrapper around gtk_file_chooser_get_current_folder().
func (chooser FileChooser) GetCurrentFolder() string {
	ret0 := C.gtk_file_chooser_get_current_folder(chooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetCurrentFolderFile is a wrapper around gtk_file_chooser_get_current_folder_file().
func (chooser FileChooser) GetCurrentFolderFile() gio.File {
	ret0 := C.gtk_file_chooser_get_current_folder_file(chooser.native())
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetCurrentFolderUri is a wrapper around gtk_file_chooser_get_current_folder_uri().
func (chooser FileChooser) GetCurrentFolderUri() string {
	ret0 := C.gtk_file_chooser_get_current_folder_uri(chooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetCurrentName is a wrapper around gtk_file_chooser_get_current_name().
func (chooser FileChooser) GetCurrentName() string {
	ret0 := C.gtk_file_chooser_get_current_name(chooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetDoOverwriteConfirmation is a wrapper around gtk_file_chooser_get_do_overwrite_confirmation().
func (chooser FileChooser) GetDoOverwriteConfirmation() bool {
	ret0 := C.gtk_file_chooser_get_do_overwrite_confirmation(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetExtraWidget is a wrapper around gtk_file_chooser_get_extra_widget().
func (chooser FileChooser) GetExtraWidget() Widget {
	ret0 := C.gtk_file_chooser_get_extra_widget(chooser.native())
	return wrapWidget(ret0)
}

// GetFile is a wrapper around gtk_file_chooser_get_file().
func (chooser FileChooser) GetFile() gio.File {
	ret0 := C.gtk_file_chooser_get_file(chooser.native())
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetFilename is a wrapper around gtk_file_chooser_get_filename().
func (chooser FileChooser) GetFilename() string {
	ret0 := C.gtk_file_chooser_get_filename(chooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetFilter is a wrapper around gtk_file_chooser_get_filter().
func (chooser FileChooser) GetFilter() FileFilter {
	ret0 := C.gtk_file_chooser_get_filter(chooser.native())
	return wrapFileFilter(ret0)
}

// GetLocalOnly is a wrapper around gtk_file_chooser_get_local_only().
func (chooser FileChooser) GetLocalOnly() bool {
	ret0 := C.gtk_file_chooser_get_local_only(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPreviewFile is a wrapper around gtk_file_chooser_get_preview_file().
func (chooser FileChooser) GetPreviewFile() gio.File {
	ret0 := C.gtk_file_chooser_get_preview_file(chooser.native())
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetPreviewFilename is a wrapper around gtk_file_chooser_get_preview_filename().
func (chooser FileChooser) GetPreviewFilename() string {
	ret0 := C.gtk_file_chooser_get_preview_filename(chooser.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetPreviewUri is a wrapper around gtk_file_chooser_get_preview_uri().
func (chooser FileChooser) GetPreviewUri() string {
	ret0 := C.gtk_file_chooser_get_preview_uri(chooser.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetPreviewWidget is a wrapper around gtk_file_chooser_get_preview_widget().
func (chooser FileChooser) GetPreviewWidget() Widget {
	ret0 := C.gtk_file_chooser_get_preview_widget(chooser.native())
	return wrapWidget(ret0)
}

// GetPreviewWidgetActive is a wrapper around gtk_file_chooser_get_preview_widget_active().
func (chooser FileChooser) GetPreviewWidgetActive() bool {
	ret0 := C.gtk_file_chooser_get_preview_widget_active(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSelectMultiple is a wrapper around gtk_file_chooser_get_select_multiple().
func (chooser FileChooser) GetSelectMultiple() bool {
	ret0 := C.gtk_file_chooser_get_select_multiple(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowHidden is a wrapper around gtk_file_chooser_get_show_hidden().
func (chooser FileChooser) GetShowHidden() bool {
	ret0 := C.gtk_file_chooser_get_show_hidden(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUri is a wrapper around gtk_file_chooser_get_uri().
func (chooser FileChooser) GetUri() string {
	ret0 := C.gtk_file_chooser_get_uri(chooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUsePreviewLabel is a wrapper around gtk_file_chooser_get_use_preview_label().
func (chooser FileChooser) GetUsePreviewLabel() bool {
	ret0 := C.gtk_file_chooser_get_use_preview_label(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RemoveChoice is a wrapper around gtk_file_chooser_remove_choice().
func (chooser FileChooser) RemoveChoice(id string) {
	id0 := C.CString(id)
	C.gtk_file_chooser_remove_choice(chooser.native(), id0)
	C.free(unsafe.Pointer(id0)) /*ch:<stdlib.h>*/
}

// RemoveFilter is a wrapper around gtk_file_chooser_remove_filter().
func (chooser FileChooser) RemoveFilter(filter FileFilter) {
	C.gtk_file_chooser_remove_filter(chooser.native(), filter.native())
}

// RemoveShortcutFolder is a wrapper around gtk_file_chooser_remove_shortcut_folder().
func (chooser FileChooser) RemoveShortcutFolder(folder string) (bool, error) {
	folder0 := C.CString(folder)
	var err glib.Error
	ret0 := C.gtk_file_chooser_remove_shortcut_folder(chooser.native(), folder0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(folder0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// RemoveShortcutFolderUri is a wrapper around gtk_file_chooser_remove_shortcut_folder_uri().
func (chooser FileChooser) RemoveShortcutFolderUri(uri string) (bool, error) {
	uri0 := C.CString(uri)
	var err glib.Error
	ret0 := C.gtk_file_chooser_remove_shortcut_folder_uri(chooser.native(), uri0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SelectAll is a wrapper around gtk_file_chooser_select_all().
func (chooser FileChooser) SelectAll() {
	C.gtk_file_chooser_select_all(chooser.native())
}

// SelectFile is a wrapper around gtk_file_chooser_select_file().
func (chooser FileChooser) SelectFile(file gio.File) (bool, error) {
	var err glib.Error
	ret0 := C.gtk_file_chooser_select_file(chooser.native(), (*C.GFile)(file.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SelectFilename is a wrapper around gtk_file_chooser_select_filename().
func (chooser FileChooser) SelectFilename(filename string) bool {
	filename0 := C.CString(filename)
	ret0 := C.gtk_file_chooser_select_filename(chooser.native(), filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))   /*go:.util*/
}

// SelectUri is a wrapper around gtk_file_chooser_select_uri().
func (chooser FileChooser) SelectUri(uri string) bool {
	uri0 := C.CString(uri)
	ret0 := C.gtk_file_chooser_select_uri(chooser.native(), uri0)
	C.free(unsafe.Pointer(uri0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetChoice is a wrapper around gtk_file_chooser_set_choice().
func (chooser FileChooser) SetChoice(id string, option string) {
	id0 := C.CString(id)
	option0 := C.CString(option)
	C.gtk_file_chooser_set_choice(chooser.native(), id0, option0)
	C.free(unsafe.Pointer(id0))     /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(option0)) /*ch:<stdlib.h>*/
}

// SetCreateFolders is a wrapper around gtk_file_chooser_set_create_folders().
func (chooser FileChooser) SetCreateFolders(create_folders bool) {
	C.gtk_file_chooser_set_create_folders(chooser.native(), C.gboolean(util.Bool2Int(create_folders)) /*go:.util*/)
}

// SetCurrentFolderFile is a wrapper around gtk_file_chooser_set_current_folder_file().
func (chooser FileChooser) SetCurrentFolderFile(file gio.File) (bool, error) {
	var err glib.Error
	ret0 := C.gtk_file_chooser_set_current_folder_file(chooser.native(), (*C.GFile)(file.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetCurrentFolderUri is a wrapper around gtk_file_chooser_set_current_folder_uri().
func (chooser FileChooser) SetCurrentFolderUri(uri string) bool {
	uri0 := (*C.gchar)(C.CString(uri))
	ret0 := C.gtk_file_chooser_set_current_folder_uri(chooser.native(), uri0)
	C.free(unsafe.Pointer(uri0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetCurrentName is a wrapper around gtk_file_chooser_set_current_name().
func (chooser FileChooser) SetCurrentName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_file_chooser_set_current_name(chooser.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetDoOverwriteConfirmation is a wrapper around gtk_file_chooser_set_do_overwrite_confirmation().
func (chooser FileChooser) SetDoOverwriteConfirmation(do_overwrite_confirmation bool) {
	C.gtk_file_chooser_set_do_overwrite_confirmation(chooser.native(), C.gboolean(util.Bool2Int(do_overwrite_confirmation)) /*go:.util*/)
}

// SetExtraWidget is a wrapper around gtk_file_chooser_set_extra_widget().
func (chooser FileChooser) SetExtraWidget(extra_widget Widget) {
	C.gtk_file_chooser_set_extra_widget(chooser.native(), extra_widget.native())
}

// SetFile is a wrapper around gtk_file_chooser_set_file().
func (chooser FileChooser) SetFile(file gio.File) (bool, error) {
	var err glib.Error
	ret0 := C.gtk_file_chooser_set_file(chooser.native(), (*C.GFile)(file.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetFilter is a wrapper around gtk_file_chooser_set_filter().
func (chooser FileChooser) SetFilter(filter FileFilter) {
	C.gtk_file_chooser_set_filter(chooser.native(), filter.native())
}

// SetLocalOnly is a wrapper around gtk_file_chooser_set_local_only().
func (chooser FileChooser) SetLocalOnly(local_only bool) {
	C.gtk_file_chooser_set_local_only(chooser.native(), C.gboolean(util.Bool2Int(local_only)) /*go:.util*/)
}

// SetPreviewWidget is a wrapper around gtk_file_chooser_set_preview_widget().
func (chooser FileChooser) SetPreviewWidget(preview_widget Widget) {
	C.gtk_file_chooser_set_preview_widget(chooser.native(), preview_widget.native())
}

// SetPreviewWidgetActive is a wrapper around gtk_file_chooser_set_preview_widget_active().
func (chooser FileChooser) SetPreviewWidgetActive(active bool) {
	C.gtk_file_chooser_set_preview_widget_active(chooser.native(), C.gboolean(util.Bool2Int(active)) /*go:.util*/)
}

// SetSelectMultiple is a wrapper around gtk_file_chooser_set_select_multiple().
func (chooser FileChooser) SetSelectMultiple(select_multiple bool) {
	C.gtk_file_chooser_set_select_multiple(chooser.native(), C.gboolean(util.Bool2Int(select_multiple)) /*go:.util*/)
}

// SetShowHidden is a wrapper around gtk_file_chooser_set_show_hidden().
func (chooser FileChooser) SetShowHidden(show_hidden bool) {
	C.gtk_file_chooser_set_show_hidden(chooser.native(), C.gboolean(util.Bool2Int(show_hidden)) /*go:.util*/)
}

// SetUri is a wrapper around gtk_file_chooser_set_uri().
func (chooser FileChooser) SetUri(uri string) bool {
	uri0 := C.CString(uri)
	ret0 := C.gtk_file_chooser_set_uri(chooser.native(), uri0)
	C.free(unsafe.Pointer(uri0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUsePreviewLabel is a wrapper around gtk_file_chooser_set_use_preview_label().
func (chooser FileChooser) SetUsePreviewLabel(use_label bool) {
	C.gtk_file_chooser_set_use_preview_label(chooser.native(), C.gboolean(util.Bool2Int(use_label)) /*go:.util*/)
}

// UnselectAll is a wrapper around gtk_file_chooser_unselect_all().
func (chooser FileChooser) UnselectAll() {
	C.gtk_file_chooser_unselect_all(chooser.native())
}

// UnselectFile is a wrapper around gtk_file_chooser_unselect_file().
func (chooser FileChooser) UnselectFile(file gio.File) {
	C.gtk_file_chooser_unselect_file(chooser.native(), (*C.GFile)(file.Ptr))
}

// UnselectFilename is a wrapper around gtk_file_chooser_unselect_filename().
func (chooser FileChooser) UnselectFilename(filename string) {
	filename0 := C.CString(filename)
	C.gtk_file_chooser_unselect_filename(chooser.native(), filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
}

// UnselectUri is a wrapper around gtk_file_chooser_unselect_uri().
func (chooser FileChooser) UnselectUri(uri string) {
	uri0 := C.CString(uri)
	C.gtk_file_chooser_unselect_uri(chooser.native(), uri0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
}

// Interface FontChooser
type FontChooser struct {
	Ptr unsafe.Pointer
}

func (v FontChooser) native() *C.GtkFontChooser {
	return (*C.GtkFontChooser)(v.Ptr)
}
func wrapFontChooser(p *C.GtkFontChooser) FontChooser {
	return FontChooser{unsafe.Pointer(p)}
}
func WrapFontChooser(p unsafe.Pointer) FontChooser {
	return FontChooser{p}
}
func (v FontChooser) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontChooser(p unsafe.Pointer) interface{} {
	return WrapFontChooser(p)
}
func (v FontChooser) GetType() gobject.Type {
	return gobject.Type(C.gtk_font_chooser_get_type())
}
func (v FontChooser) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontChooser(unsafe.Pointer(ptr)), nil
	}
}

// GetFont is a wrapper around gtk_font_chooser_get_font().
func (fontchooser FontChooser) GetFont() string {
	ret0 := C.gtk_font_chooser_get_font(fontchooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetFontDesc is a wrapper around gtk_font_chooser_get_font_desc().
func (fontchooser FontChooser) GetFontDesc() pango.FontDescription {
	ret0 := C.gtk_font_chooser_get_font_desc(fontchooser.native())
	return pango.WrapFontDescription(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetFontFace is a wrapper around gtk_font_chooser_get_font_face().
func (fontchooser FontChooser) GetFontFace() pango.FontFace {
	ret0 := C.gtk_font_chooser_get_font_face(fontchooser.native())
	return pango.WrapFontFace(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetFontFamily is a wrapper around gtk_font_chooser_get_font_family().
func (fontchooser FontChooser) GetFontFamily() pango.FontFamily {
	ret0 := C.gtk_font_chooser_get_font_family(fontchooser.native())
	return pango.WrapFontFamily(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetFontMap is a wrapper around gtk_font_chooser_get_font_map().
func (fontchooser FontChooser) GetFontMap() pango.FontMap {
	ret0 := C.gtk_font_chooser_get_font_map(fontchooser.native())
	return pango.WrapFontMap(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetFontSize is a wrapper around gtk_font_chooser_get_font_size().
func (fontchooser FontChooser) GetFontSize() int {
	ret0 := C.gtk_font_chooser_get_font_size(fontchooser.native())
	return int(ret0)
}

// GetPreviewText is a wrapper around gtk_font_chooser_get_preview_text().
func (fontchooser FontChooser) GetPreviewText() string {
	ret0 := C.gtk_font_chooser_get_preview_text(fontchooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetShowPreviewEntry is a wrapper around gtk_font_chooser_get_show_preview_entry().
func (fontchooser FontChooser) GetShowPreviewEntry() bool {
	ret0 := C.gtk_font_chooser_get_show_preview_entry(fontchooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetFont is a wrapper around gtk_font_chooser_set_font().
func (fontchooser FontChooser) SetFont(fontname string) {
	fontname0 := (*C.gchar)(C.CString(fontname))
	C.gtk_font_chooser_set_font(fontchooser.native(), fontname0)
	C.free(unsafe.Pointer(fontname0)) /*ch:<stdlib.h>*/
}

// SetFontDesc is a wrapper around gtk_font_chooser_set_font_desc().
func (fontchooser FontChooser) SetFontDesc(font_desc pango.FontDescription) {
	C.gtk_font_chooser_set_font_desc(fontchooser.native(), (*C.PangoFontDescription)(font_desc.Ptr))
}

// SetFontMap is a wrapper around gtk_font_chooser_set_font_map().
func (fontchooser FontChooser) SetFontMap(fontmap pango.FontMap) {
	C.gtk_font_chooser_set_font_map(fontchooser.native(), (*C.PangoFontMap)(fontmap.Ptr))
}

// SetPreviewText is a wrapper around gtk_font_chooser_set_preview_text().
func (fontchooser FontChooser) SetPreviewText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_font_chooser_set_preview_text(fontchooser.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetShowPreviewEntry is a wrapper around gtk_font_chooser_set_show_preview_entry().
func (fontchooser FontChooser) SetShowPreviewEntry(show_preview_entry bool) {
	C.gtk_font_chooser_set_show_preview_entry(fontchooser.native(), C.gboolean(util.Bool2Int(show_preview_entry)) /*go:.util*/)
}

// Interface Orientable
type Orientable struct {
	Ptr unsafe.Pointer
}

func (v Orientable) native() *C.GtkOrientable {
	return (*C.GtkOrientable)(v.Ptr)
}
func wrapOrientable(p *C.GtkOrientable) Orientable {
	return Orientable{unsafe.Pointer(p)}
}
func WrapOrientable(p unsafe.Pointer) Orientable {
	return Orientable{p}
}
func (v Orientable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOrientable(p unsafe.Pointer) interface{} {
	return WrapOrientable(p)
}
func (v Orientable) GetType() gobject.Type {
	return gobject.Type(C.gtk_orientable_get_type())
}
func (v Orientable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapOrientable(unsafe.Pointer(ptr)), nil
	}
}

// GetOrientation is a wrapper around gtk_orientable_get_orientation().
func (orientable Orientable) GetOrientation() Orientation {
	ret0 := C.gtk_orientable_get_orientation(orientable.native())
	return Orientation(ret0)
}

// SetOrientation is a wrapper around gtk_orientable_set_orientation().
func (orientable Orientable) SetOrientation(orientation Orientation) {
	C.gtk_orientable_set_orientation(orientable.native(), C.GtkOrientation(orientation))
}

// Interface PrintOperationPreview
type PrintOperationPreview struct {
	Ptr unsafe.Pointer
}

func (v PrintOperationPreview) native() *C.GtkPrintOperationPreview {
	return (*C.GtkPrintOperationPreview)(v.Ptr)
}
func wrapPrintOperationPreview(p *C.GtkPrintOperationPreview) PrintOperationPreview {
	return PrintOperationPreview{unsafe.Pointer(p)}
}
func WrapPrintOperationPreview(p unsafe.Pointer) PrintOperationPreview {
	return PrintOperationPreview{p}
}
func (v PrintOperationPreview) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPrintOperationPreview(p unsafe.Pointer) interface{} {
	return WrapPrintOperationPreview(p)
}
func (v PrintOperationPreview) GetType() gobject.Type {
	return gobject.Type(C.gtk_print_operation_preview_get_type())
}
func (v PrintOperationPreview) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPrintOperationPreview(unsafe.Pointer(ptr)), nil
	}
}

// EndPreview is a wrapper around gtk_print_operation_preview_end_preview().
func (preview PrintOperationPreview) EndPreview() {
	C.gtk_print_operation_preview_end_preview(preview.native())
}

// IsSelected is a wrapper around gtk_print_operation_preview_is_selected().
func (preview PrintOperationPreview) IsSelected(page_nr int) bool {
	ret0 := C.gtk_print_operation_preview_is_selected(preview.native(), C.gint(page_nr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RenderPage is a wrapper around gtk_print_operation_preview_render_page().
func (preview PrintOperationPreview) RenderPage(page_nr int) {
	C.gtk_print_operation_preview_render_page(preview.native(), C.gint(page_nr))
}

// Interface Scrollable
type Scrollable struct {
	Ptr unsafe.Pointer
}

func (v Scrollable) native() *C.GtkScrollable {
	return (*C.GtkScrollable)(v.Ptr)
}
func wrapScrollable(p *C.GtkScrollable) Scrollable {
	return Scrollable{unsafe.Pointer(p)}
}
func WrapScrollable(p unsafe.Pointer) Scrollable {
	return Scrollable{p}
}
func (v Scrollable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScrollable(p unsafe.Pointer) interface{} {
	return WrapScrollable(p)
}
func (v Scrollable) GetType() gobject.Type {
	return gobject.Type(C.gtk_scrollable_get_type())
}
func (v Scrollable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScrollable(unsafe.Pointer(ptr)), nil
	}
}

// GetHadjustment is a wrapper around gtk_scrollable_get_hadjustment().
func (scrollable Scrollable) GetHadjustment() Adjustment {
	ret0 := C.gtk_scrollable_get_hadjustment(scrollable.native())
	return wrapAdjustment(ret0)
}

// GetHscrollPolicy is a wrapper around gtk_scrollable_get_hscroll_policy().
func (scrollable Scrollable) GetHscrollPolicy() ScrollablePolicy {
	ret0 := C.gtk_scrollable_get_hscroll_policy(scrollable.native())
	return ScrollablePolicy(ret0)
}

// GetVadjustment is a wrapper around gtk_scrollable_get_vadjustment().
func (scrollable Scrollable) GetVadjustment() Adjustment {
	ret0 := C.gtk_scrollable_get_vadjustment(scrollable.native())
	return wrapAdjustment(ret0)
}

// GetVscrollPolicy is a wrapper around gtk_scrollable_get_vscroll_policy().
func (scrollable Scrollable) GetVscrollPolicy() ScrollablePolicy {
	ret0 := C.gtk_scrollable_get_vscroll_policy(scrollable.native())
	return ScrollablePolicy(ret0)
}

// SetHadjustment is a wrapper around gtk_scrollable_set_hadjustment().
func (scrollable Scrollable) SetHadjustment(hadjustment Adjustment) {
	C.gtk_scrollable_set_hadjustment(scrollable.native(), hadjustment.native())
}

// SetHscrollPolicy is a wrapper around gtk_scrollable_set_hscroll_policy().
func (scrollable Scrollable) SetHscrollPolicy(policy ScrollablePolicy) {
	C.gtk_scrollable_set_hscroll_policy(scrollable.native(), C.GtkScrollablePolicy(policy))
}

// SetVadjustment is a wrapper around gtk_scrollable_set_vadjustment().
func (scrollable Scrollable) SetVadjustment(vadjustment Adjustment) {
	C.gtk_scrollable_set_vadjustment(scrollable.native(), vadjustment.native())
}

// SetVscrollPolicy is a wrapper around gtk_scrollable_set_vscroll_policy().
func (scrollable Scrollable) SetVscrollPolicy(policy ScrollablePolicy) {
	C.gtk_scrollable_set_vscroll_policy(scrollable.native(), C.GtkScrollablePolicy(policy))
}

// Interface StyleProvider
type StyleProvider struct {
	Ptr unsafe.Pointer
}

func (v StyleProvider) native() *C.GtkStyleProvider {
	return (*C.GtkStyleProvider)(v.Ptr)
}
func wrapStyleProvider(p *C.GtkStyleProvider) StyleProvider {
	return StyleProvider{unsafe.Pointer(p)}
}
func WrapStyleProvider(p unsafe.Pointer) StyleProvider {
	return StyleProvider{p}
}
func (v StyleProvider) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStyleProvider(p unsafe.Pointer) interface{} {
	return WrapStyleProvider(p)
}
func (v StyleProvider) GetType() gobject.Type {
	return gobject.Type(C.gtk_style_provider_get_type())
}
func (v StyleProvider) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStyleProvider(unsafe.Pointer(ptr)), nil
	}
}

// Interface ToolShell
type ToolShell struct {
	Ptr unsafe.Pointer
}

func (v ToolShell) native() *C.GtkToolShell {
	return (*C.GtkToolShell)(v.Ptr)
}
func wrapToolShell(p *C.GtkToolShell) ToolShell {
	return ToolShell{unsafe.Pointer(p)}
}
func WrapToolShell(p unsafe.Pointer) ToolShell {
	return ToolShell{p}
}
func (v ToolShell) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToolShell(p unsafe.Pointer) interface{} {
	return WrapToolShell(p)
}
func (v ToolShell) GetType() gobject.Type {
	return gobject.Type(C.gtk_tool_shell_get_type())
}
func (v ToolShell) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToolShell(unsafe.Pointer(ptr)), nil
	}
}

// GetEllipsizeMode is a wrapper around gtk_tool_shell_get_ellipsize_mode().
func (shell ToolShell) GetEllipsizeMode() /*gir:Pango*/ pango.EllipsizeMode {
	ret0 := C.gtk_tool_shell_get_ellipsize_mode(shell.native())
	return /*gir:Pango*/ pango.EllipsizeMode(ret0)
}

// GetOrientation is a wrapper around gtk_tool_shell_get_orientation().
func (shell ToolShell) GetOrientation() Orientation {
	ret0 := C.gtk_tool_shell_get_orientation(shell.native())
	return Orientation(ret0)
}

// GetReliefStyle is a wrapper around gtk_tool_shell_get_relief_style().
func (shell ToolShell) GetReliefStyle() ReliefStyle {
	ret0 := C.gtk_tool_shell_get_relief_style(shell.native())
	return ReliefStyle(ret0)
}

// GetStyle is a wrapper around gtk_tool_shell_get_style().
func (shell ToolShell) GetStyle() ToolbarStyle {
	ret0 := C.gtk_tool_shell_get_style(shell.native())
	return ToolbarStyle(ret0)
}

// GetTextAlignment is a wrapper around gtk_tool_shell_get_text_alignment().
func (shell ToolShell) GetTextAlignment() float32 {
	ret0 := C.gtk_tool_shell_get_text_alignment(shell.native())
	return float32(ret0)
}

// GetTextOrientation is a wrapper around gtk_tool_shell_get_text_orientation().
func (shell ToolShell) GetTextOrientation() Orientation {
	ret0 := C.gtk_tool_shell_get_text_orientation(shell.native())
	return Orientation(ret0)
}

// GetTextSizeGroup is a wrapper around gtk_tool_shell_get_text_size_group().
func (shell ToolShell) GetTextSizeGroup() SizeGroup {
	ret0 := C.gtk_tool_shell_get_text_size_group(shell.native())
	return wrapSizeGroup(ret0)
}

// RebuildMenu is a wrapper around gtk_tool_shell_rebuild_menu().
func (shell ToolShell) RebuildMenu() {
	C.gtk_tool_shell_rebuild_menu(shell.native())
}

// Interface TreeDragDest
type TreeDragDest struct {
	Ptr unsafe.Pointer
}

func (v TreeDragDest) native() *C.GtkTreeDragDest {
	return (*C.GtkTreeDragDest)(v.Ptr)
}
func wrapTreeDragDest(p *C.GtkTreeDragDest) TreeDragDest {
	return TreeDragDest{unsafe.Pointer(p)}
}
func WrapTreeDragDest(p unsafe.Pointer) TreeDragDest {
	return TreeDragDest{p}
}
func (v TreeDragDest) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeDragDest(p unsafe.Pointer) interface{} {
	return WrapTreeDragDest(p)
}
func (v TreeDragDest) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_drag_dest_get_type())
}
func (v TreeDragDest) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeDragDest(unsafe.Pointer(ptr)), nil
	}
}

// DragDataReceived is a wrapper around gtk_tree_drag_dest_drag_data_received().
func (drag_dest TreeDragDest) DragDataReceived(dest TreePath, selection_data SelectionData) bool {
	ret0 := C.gtk_tree_drag_dest_drag_data_received(drag_dest.native(), dest.native(), selection_data.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RowDropPossible is a wrapper around gtk_tree_drag_dest_row_drop_possible().
func (drag_dest TreeDragDest) RowDropPossible(dest_path TreePath, selection_data SelectionData) bool {
	ret0 := C.gtk_tree_drag_dest_row_drop_possible(drag_dest.native(), dest_path.native(), selection_data.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Interface TreeDragSource
type TreeDragSource struct {
	Ptr unsafe.Pointer
}

func (v TreeDragSource) native() *C.GtkTreeDragSource {
	return (*C.GtkTreeDragSource)(v.Ptr)
}
func wrapTreeDragSource(p *C.GtkTreeDragSource) TreeDragSource {
	return TreeDragSource{unsafe.Pointer(p)}
}
func WrapTreeDragSource(p unsafe.Pointer) TreeDragSource {
	return TreeDragSource{p}
}
func (v TreeDragSource) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeDragSource(p unsafe.Pointer) interface{} {
	return WrapTreeDragSource(p)
}
func (v TreeDragSource) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_drag_source_get_type())
}
func (v TreeDragSource) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeDragSource(unsafe.Pointer(ptr)), nil
	}
}

// DragDataDelete is a wrapper around gtk_tree_drag_source_drag_data_delete().
func (drag_source TreeDragSource) DragDataDelete(path TreePath) bool {
	ret0 := C.gtk_tree_drag_source_drag_data_delete(drag_source.native(), path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// DragDataGet is a wrapper around gtk_tree_drag_source_drag_data_get().
func (drag_source TreeDragSource) DragDataGet(path TreePath, selection_data SelectionData) bool {
	ret0 := C.gtk_tree_drag_source_drag_data_get(drag_source.native(), path.native(), selection_data.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RowDraggable is a wrapper around gtk_tree_drag_source_row_draggable().
func (drag_source TreeDragSource) RowDraggable(path TreePath) bool {
	ret0 := C.gtk_tree_drag_source_row_draggable(drag_source.native(), path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Interface TreeSortable
type TreeSortable struct {
	Ptr unsafe.Pointer
}

func (v TreeSortable) native() *C.GtkTreeSortable {
	return (*C.GtkTreeSortable)(v.Ptr)
}
func wrapTreeSortable(p *C.GtkTreeSortable) TreeSortable {
	return TreeSortable{unsafe.Pointer(p)}
}
func WrapTreeSortable(p unsafe.Pointer) TreeSortable {
	return TreeSortable{p}
}
func (v TreeSortable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeSortable(p unsafe.Pointer) interface{} {
	return WrapTreeSortable(p)
}
func (v TreeSortable) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_sortable_get_type())
}
func (v TreeSortable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeSortable(unsafe.Pointer(ptr)), nil
	}
}

// GetSortColumnId is a wrapper around gtk_tree_sortable_get_sort_column_id().
func (sortable TreeSortable) GetSortColumnId() (bool, int, SortType) {
	var sort_column_id0 C.gint
	var order0 C.GtkSortType
	ret0 := C.gtk_tree_sortable_get_sort_column_id(sortable.native(), &sort_column_id0, &order0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(sort_column_id0), SortType(order0)
}

// HasDefaultSortFunc is a wrapper around gtk_tree_sortable_has_default_sort_func().
func (sortable TreeSortable) HasDefaultSortFunc() bool {
	ret0 := C.gtk_tree_sortable_has_default_sort_func(sortable.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetSortColumnId is a wrapper around gtk_tree_sortable_set_sort_column_id().
func (sortable TreeSortable) SetSortColumnId(sort_column_id int, order SortType) {
	C.gtk_tree_sortable_set_sort_column_id(sortable.native(), C.gint(sort_column_id), C.GtkSortType(order))
}

// SortColumnChanged is a wrapper around gtk_tree_sortable_sort_column_changed().
func (sortable TreeSortable) SortColumnChanged() {
	C.gtk_tree_sortable_sort_column_changed(sortable.native())
}

// Object AccelGroup
type AccelGroup struct {
	gobject.Object
}

func (v AccelGroup) native() *C.GtkAccelGroup {
	return (*C.GtkAccelGroup)(v.Ptr)
}
func wrapAccelGroup(p *C.GtkAccelGroup) (v AccelGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAccelGroup(p unsafe.Pointer) (v AccelGroup) {
	v.Ptr = p
	return
}
func (v AccelGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAccelGroup(p unsafe.Pointer) interface{} {
	return WrapAccelGroup(p)
}
func (v AccelGroup) GetType() gobject.Type {
	return gobject.Type(C.gtk_accel_group_get_type())
}
func (v AccelGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAccelGroup(unsafe.Pointer(ptr)), nil
	}
}

// AccelGroupNew is a wrapper around gtk_accel_group_new().
func AccelGroupNew() AccelGroup {
	ret0 := C.gtk_accel_group_new()
	return wrapAccelGroup(ret0)
}

// Activate is a wrapper around gtk_accel_group_activate().
func (accel_group AccelGroup) Activate(accel_quark /*gir:GLib*/ glib.Quark, acceleratable gobject.Object, accel_key uint, accel_mods /*gir:Gdk*/ gdk.ModifierType) bool {
	ret0 := C.gtk_accel_group_activate(accel_group.native(), C.GQuark(accel_quark), (*C.GObject)(acceleratable.Ptr), C.guint(accel_key), C.GdkModifierType(accel_mods))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Connect is a wrapper around gtk_accel_group_connect().
func (accel_group AccelGroup) Connect(accel_key uint, accel_mods /*gir:Gdk*/ gdk.ModifierType, accel_flags AccelFlags, closure gobject.Closure) {
	C.gtk_accel_group_connect(accel_group.native(), C.guint(accel_key), C.GdkModifierType(accel_mods), C.GtkAccelFlags(accel_flags), (*C.GClosure)(closure.Ptr))
}

// Disconnect is a wrapper around gtk_accel_group_disconnect().
func (accel_group AccelGroup) Disconnect(closure gobject.Closure) bool {
	ret0 := C.gtk_accel_group_disconnect(accel_group.native(), (*C.GClosure)(closure.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// DisconnectKey is a wrapper around gtk_accel_group_disconnect_key().
func (accel_group AccelGroup) DisconnectKey(accel_key uint, accel_mods /*gir:Gdk*/ gdk.ModifierType) bool {
	ret0 := C.gtk_accel_group_disconnect_key(accel_group.native(), C.guint(accel_key), C.GdkModifierType(accel_mods))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsLocked is a wrapper around gtk_accel_group_get_is_locked().
func (accel_group AccelGroup) GetIsLocked() bool {
	ret0 := C.gtk_accel_group_get_is_locked(accel_group.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetModifierMask is a wrapper around gtk_accel_group_get_modifier_mask().
func (accel_group AccelGroup) GetModifierMask() /*gir:Gdk*/ gdk.ModifierType {
	ret0 := C.gtk_accel_group_get_modifier_mask(accel_group.native())
	return /*gir:Gdk*/ gdk.ModifierType(ret0)
}

// Lock is a wrapper around gtk_accel_group_lock().
func (accel_group AccelGroup) Lock() {
	C.gtk_accel_group_lock(accel_group.native())
}

// Unlock is a wrapper around gtk_accel_group_unlock().
func (accel_group AccelGroup) Unlock() {
	C.gtk_accel_group_unlock(accel_group.native())
}

// AccelGroupFromAccelClosure is a wrapper around gtk_accel_group_from_accel_closure().
func AccelGroupFromAccelClosure(closure gobject.Closure) AccelGroup {
	ret0 := C.gtk_accel_group_from_accel_closure((*C.GClosure)(closure.Ptr))
	return wrapAccelGroup(ret0)
}

// Object AccelMap
type AccelMap struct {
	gobject.Object
}

func (v AccelMap) native() *C.GtkAccelMap {
	return (*C.GtkAccelMap)(v.Ptr)
}
func wrapAccelMap(p *C.GtkAccelMap) (v AccelMap) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAccelMap(p unsafe.Pointer) (v AccelMap) {
	v.Ptr = p
	return
}
func (v AccelMap) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAccelMap(p unsafe.Pointer) interface{} {
	return WrapAccelMap(p)
}
func (v AccelMap) GetType() gobject.Type {
	return gobject.Type(C.gtk_accel_map_get_type())
}
func (v AccelMap) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAccelMap(unsafe.Pointer(ptr)), nil
	}
}

// AccelMapAddEntry is a wrapper around gtk_accel_map_add_entry().
func AccelMapAddEntry(accel_path string, accel_key uint, accel_mods /*gir:Gdk*/ gdk.ModifierType) {
	accel_path0 := (*C.gchar)(C.CString(accel_path))
	C.gtk_accel_map_add_entry(accel_path0, C.guint(accel_key), C.GdkModifierType(accel_mods))
	C.free(unsafe.Pointer(accel_path0)) /*ch:<stdlib.h>*/
}

// AccelMapAddFilter is a wrapper around gtk_accel_map_add_filter().
func AccelMapAddFilter(filter_pattern string) {
	filter_pattern0 := (*C.gchar)(C.CString(filter_pattern))
	C.gtk_accel_map_add_filter(filter_pattern0)
	C.free(unsafe.Pointer(filter_pattern0)) /*ch:<stdlib.h>*/
}

// AccelMapChangeEntry is a wrapper around gtk_accel_map_change_entry().
func AccelMapChangeEntry(accel_path string, accel_key uint, accel_mods /*gir:Gdk*/ gdk.ModifierType, replace bool) bool {
	accel_path0 := (*C.gchar)(C.CString(accel_path))
	ret0 := C.gtk_accel_map_change_entry(accel_path0, C.guint(accel_key), C.GdkModifierType(accel_mods), C.gboolean(util.Bool2Int(replace)) /*go:.util*/)
	C.free(unsafe.Pointer(accel_path0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// AccelMapGet is a wrapper around gtk_accel_map_get().
func AccelMapGet() AccelMap {
	ret0 := C.gtk_accel_map_get()
	return wrapAccelMap(ret0)
}

// AccelMapLoad is a wrapper around gtk_accel_map_load().
func AccelMapLoad(file_name string) {
	file_name0 := (*C.gchar)(C.CString(file_name))
	C.gtk_accel_map_load(file_name0)
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
}

// AccelMapLoadFd is a wrapper around gtk_accel_map_load_fd().
func AccelMapLoadFd(fd int) {
	C.gtk_accel_map_load_fd(C.gint(fd))
}

// AccelMapLockPath is a wrapper around gtk_accel_map_lock_path().
func AccelMapLockPath(accel_path string) {
	accel_path0 := (*C.gchar)(C.CString(accel_path))
	C.gtk_accel_map_lock_path(accel_path0)
	C.free(unsafe.Pointer(accel_path0)) /*ch:<stdlib.h>*/
}

// AccelMapSave is a wrapper around gtk_accel_map_save().
func AccelMapSave(file_name string) {
	file_name0 := (*C.gchar)(C.CString(file_name))
	C.gtk_accel_map_save(file_name0)
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
}

// AccelMapSaveFd is a wrapper around gtk_accel_map_save_fd().
func AccelMapSaveFd(fd int) {
	C.gtk_accel_map_save_fd(C.gint(fd))
}

// AccelMapUnlockPath is a wrapper around gtk_accel_map_unlock_path().
func AccelMapUnlockPath(accel_path string) {
	accel_path0 := (*C.gchar)(C.CString(accel_path))
	C.gtk_accel_map_unlock_path(accel_path0)
	C.free(unsafe.Pointer(accel_path0)) /*ch:<stdlib.h>*/
}

// Object Accessible
type Accessible struct {
	atk.Object
}

func (v Accessible) native() *C.GtkAccessible {
	return (*C.GtkAccessible)(v.Ptr)
}
func wrapAccessible(p *C.GtkAccessible) (v Accessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAccessible(p unsafe.Pointer) (v Accessible) {
	v.Ptr = p
	return
}
func (v Accessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAccessible(p unsafe.Pointer) interface{} {
	return WrapAccessible(p)
}
func (v Accessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_accessible_get_type())
}
func (v Accessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAccessible(unsafe.Pointer(ptr)), nil
	}
}

// GetWidget is a wrapper around gtk_accessible_get_widget().
func (accessible Accessible) GetWidget() Widget {
	ret0 := C.gtk_accessible_get_widget(accessible.native())
	return wrapWidget(ret0)
}

// SetWidget is a wrapper around gtk_accessible_set_widget().
func (accessible Accessible) SetWidget(widget Widget) {
	C.gtk_accessible_set_widget(accessible.native(), widget.native())
}

// Object Adjustment
type Adjustment struct {
	gobject.InitiallyUnowned
}

func (v Adjustment) native() *C.GtkAdjustment {
	return (*C.GtkAdjustment)(v.Ptr)
}
func wrapAdjustment(p *C.GtkAdjustment) (v Adjustment) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAdjustment(p unsafe.Pointer) (v Adjustment) {
	v.Ptr = p
	return
}
func (v Adjustment) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAdjustment(p unsafe.Pointer) interface{} {
	return WrapAdjustment(p)
}
func (v Adjustment) GetType() gobject.Type {
	return gobject.Type(C.gtk_adjustment_get_type())
}
func (v Adjustment) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAdjustment(unsafe.Pointer(ptr)), nil
	}
}

// AdjustmentNew is a wrapper around gtk_adjustment_new().
func AdjustmentNew(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) Adjustment {
	ret0 := C.gtk_adjustment_new(C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size))
	return wrapAdjustment(ret0)
}

// ClampPage is a wrapper around gtk_adjustment_clamp_page().
func (adjustment Adjustment) ClampPage(lower float64, upper float64) {
	C.gtk_adjustment_clamp_page(adjustment.native(), C.gdouble(lower), C.gdouble(upper))
}

// Configure is a wrapper around gtk_adjustment_configure().
func (adjustment Adjustment) Configure(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) {
	C.gtk_adjustment_configure(adjustment.native(), C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size))
}

// GetLower is a wrapper around gtk_adjustment_get_lower().
func (adjustment Adjustment) GetLower() float64 {
	ret0 := C.gtk_adjustment_get_lower(adjustment.native())
	return float64(ret0)
}

// GetMinimumIncrement is a wrapper around gtk_adjustment_get_minimum_increment().
func (adjustment Adjustment) GetMinimumIncrement() float64 {
	ret0 := C.gtk_adjustment_get_minimum_increment(adjustment.native())
	return float64(ret0)
}

// GetPageSize is a wrapper around gtk_adjustment_get_page_size().
func (adjustment Adjustment) GetPageSize() float64 {
	ret0 := C.gtk_adjustment_get_page_size(adjustment.native())
	return float64(ret0)
}

// GetStepIncrement is a wrapper around gtk_adjustment_get_step_increment().
func (adjustment Adjustment) GetStepIncrement() float64 {
	ret0 := C.gtk_adjustment_get_step_increment(adjustment.native())
	return float64(ret0)
}

// GetUpper is a wrapper around gtk_adjustment_get_upper().
func (adjustment Adjustment) GetUpper() float64 {
	ret0 := C.gtk_adjustment_get_upper(adjustment.native())
	return float64(ret0)
}

// GetValue is a wrapper around gtk_adjustment_get_value().
func (adjustment Adjustment) GetValue() float64 {
	ret0 := C.gtk_adjustment_get_value(adjustment.native())
	return float64(ret0)
}

// SetPageSize is a wrapper around gtk_adjustment_set_page_size().
func (adjustment Adjustment) SetPageSize(page_size float64) {
	C.gtk_adjustment_set_page_size(adjustment.native(), C.gdouble(page_size))
}

// SetStepIncrement is a wrapper around gtk_adjustment_set_step_increment().
func (adjustment Adjustment) SetStepIncrement(step_increment float64) {
	C.gtk_adjustment_set_step_increment(adjustment.native(), C.gdouble(step_increment))
}

// SetUpper is a wrapper around gtk_adjustment_set_upper().
func (adjustment Adjustment) SetUpper(upper float64) {
	C.gtk_adjustment_set_upper(adjustment.native(), C.gdouble(upper))
}

// SetValue is a wrapper around gtk_adjustment_set_value().
func (adjustment Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(adjustment.native(), C.gdouble(value))
}

// Object AppChooserDialog
type AppChooserDialog struct {
	Dialog
}

func (v AppChooserDialog) native() *C.GtkAppChooserDialog {
	return (*C.GtkAppChooserDialog)(v.Ptr)
}
func wrapAppChooserDialog(p *C.GtkAppChooserDialog) (v AppChooserDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAppChooserDialog(p unsafe.Pointer) (v AppChooserDialog) {
	v.Ptr = p
	return
}
func (v AppChooserDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppChooserDialog(p unsafe.Pointer) interface{} {
	return WrapAppChooserDialog(p)
}
func (v AppChooserDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_app_chooser_dialog_get_type())
}
func (v AppChooserDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppChooserDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v AppChooserDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v AppChooserDialog) AppChooser() AppChooser {
	return WrapAppChooser(v.Ptr)
}
func (v AppChooserDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// AppChooserDialogNewForContentType is a wrapper around gtk_app_chooser_dialog_new_for_content_type().
func AppChooserDialogNewForContentType(parent Window, flags DialogFlags, content_type string) Widget {
	content_type0 := (*C.gchar)(C.CString(content_type))
	ret0 := C.gtk_app_chooser_dialog_new_for_content_type(parent.native(), C.GtkDialogFlags(flags), content_type0)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetHeading is a wrapper around gtk_app_chooser_dialog_get_heading().
func (self AppChooserDialog) GetHeading() string {
	ret0 := C.gtk_app_chooser_dialog_get_heading(self.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWidget is a wrapper around gtk_app_chooser_dialog_get_widget().
func (self AppChooserDialog) GetWidget() Widget {
	ret0 := C.gtk_app_chooser_dialog_get_widget(self.native())
	return wrapWidget(ret0)
}

// SetHeading is a wrapper around gtk_app_chooser_dialog_set_heading().
func (self AppChooserDialog) SetHeading(heading string) {
	heading0 := (*C.gchar)(C.CString(heading))
	C.gtk_app_chooser_dialog_set_heading(self.native(), heading0)
	C.free(unsafe.Pointer(heading0)) /*ch:<stdlib.h>*/
}

// Object Application
type Application struct {
	gio.Application
}

func (v Application) native() *C.GtkApplication {
	return (*C.GtkApplication)(v.Ptr)
}
func wrapApplication(p *C.GtkApplication) (v Application) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapApplication(p unsafe.Pointer) (v Application) {
	v.Ptr = p
	return
}
func (v Application) IsNil() bool {
	return v.Ptr == nil
}
func IWrapApplication(p unsafe.Pointer) interface{} {
	return WrapApplication(p)
}
func (v Application) GetType() gobject.Type {
	return gobject.Type(C.gtk_application_get_type())
}
func (v Application) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapApplication(unsafe.Pointer(ptr)), nil
	}
}
func (v Application) ActionGroup() gio.ActionGroup {
	return gio.WrapActionGroup(v.Ptr) /*gir:Gio*/
}
func (v Application) ActionMap() gio.ActionMap {
	return gio.WrapActionMap(v.Ptr) /*gir:Gio*/
}

// AddWindow is a wrapper around gtk_application_add_window().
func (application Application) AddWindow(window Window) {
	C.gtk_application_add_window(application.native(), window.native())
}

// GetAccelsForAction is a wrapper around gtk_application_get_accels_for_action().
func (application Application) GetAccelsForAction(detailed_action_name string) []string {
	detailed_action_name0 := (*C.gchar)(C.CString(detailed_action_name))
	ret0 := C.gtk_application_get_accels_for_action(application.native(), detailed_action_name0)
	C.free(unsafe.Pointer(detailed_action_name0)) /*ch:<stdlib.h>*/
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

// GetActionsForAccel is a wrapper around gtk_application_get_actions_for_accel().
func (application Application) GetActionsForAccel(accel string) []string {
	accel0 := (*C.gchar)(C.CString(accel))
	ret0 := C.gtk_application_get_actions_for_accel(application.native(), accel0)
	C.free(unsafe.Pointer(accel0)) /*ch:<stdlib.h>*/
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

// GetActiveWindow is a wrapper around gtk_application_get_active_window().
func (application Application) GetActiveWindow() Window {
	ret0 := C.gtk_application_get_active_window(application.native())
	return wrapWindow(ret0)
}

// GetAppMenu is a wrapper around gtk_application_get_app_menu().
func (application Application) GetAppMenu() gio.MenuModel {
	ret0 := C.gtk_application_get_app_menu(application.native())
	return gio.WrapMenuModel(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetMenuById is a wrapper around gtk_application_get_menu_by_id().
func (application Application) GetMenuById(id string) gio.Menu {
	id0 := (*C.gchar)(C.CString(id))
	ret0 := C.gtk_application_get_menu_by_id(application.native(), id0)
	C.free(unsafe.Pointer(id0))               /*ch:<stdlib.h>*/
	return gio.WrapMenu(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetMenubar is a wrapper around gtk_application_get_menubar().
func (application Application) GetMenubar() gio.MenuModel {
	ret0 := C.gtk_application_get_menubar(application.native())
	return gio.WrapMenuModel(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetWindowById is a wrapper around gtk_application_get_window_by_id().
func (application Application) GetWindowById(id uint) Window {
	ret0 := C.gtk_application_get_window_by_id(application.native(), C.guint(id))
	return wrapWindow(ret0)
}

// GetWindows is a wrapper around gtk_application_get_windows().
func (application Application) GetWindows() glib.List {
	ret0 := C.gtk_application_get_windows(application.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// Inhibit is a wrapper around gtk_application_inhibit().
func (application Application) Inhibit(window Window, flags ApplicationInhibitFlags, reason string) uint {
	reason0 := (*C.gchar)(C.CString(reason))
	ret0 := C.gtk_application_inhibit(application.native(), window.native(), C.GtkApplicationInhibitFlags(flags), reason0)
	C.free(unsafe.Pointer(reason0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// IsInhibited is a wrapper around gtk_application_is_inhibited().
func (application Application) IsInhibited(flags ApplicationInhibitFlags) bool {
	ret0 := C.gtk_application_is_inhibited(application.native(), C.GtkApplicationInhibitFlags(flags))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ListActionDescriptions is a wrapper around gtk_application_list_action_descriptions().
func (application Application) ListActionDescriptions() []string {
	ret0 := C.gtk_application_list_action_descriptions(application.native())
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

// PrefersAppMenu is a wrapper around gtk_application_prefers_app_menu().
func (application Application) PrefersAppMenu() bool {
	ret0 := C.gtk_application_prefers_app_menu(application.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RemoveWindow is a wrapper around gtk_application_remove_window().
func (application Application) RemoveWindow(window Window) {
	C.gtk_application_remove_window(application.native(), window.native())
}

// SetAccelsForAction is a wrapper around gtk_application_set_accels_for_action().
func (application Application) SetAccelsForAction(detailed_action_name string, accels []string) {
	detailed_action_name0 := (*C.gchar)(C.CString(detailed_action_name))
	accels0 := make([]*C.gchar, len(accels))
	for idx, elemG := range accels {
		elem := (*C.gchar)(C.CString(elemG))
		accels0[idx] = elem
	}
	var accels0Ptr **C.gchar
	if len(accels0) > 0 {
		accels0Ptr = &accels0[0]
	}
	C.gtk_application_set_accels_for_action(application.native(), detailed_action_name0, accels0Ptr)
	C.free(unsafe.Pointer(detailed_action_name0)) /*ch:<stdlib.h>*/
	for _, elem := range accels0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// SetAppMenu is a wrapper around gtk_application_set_app_menu().
func (application Application) SetAppMenu(app_menu gio.MenuModel) {
	C.gtk_application_set_app_menu(application.native(), (*C.GMenuModel)(app_menu.Ptr))
}

// SetMenubar is a wrapper around gtk_application_set_menubar().
func (application Application) SetMenubar(menubar gio.MenuModel) {
	C.gtk_application_set_menubar(application.native(), (*C.GMenuModel)(menubar.Ptr))
}

// Uninhibit is a wrapper around gtk_application_uninhibit().
func (application Application) Uninhibit(cookie uint) {
	C.gtk_application_uninhibit(application.native(), C.guint(cookie))
}

// Object Bin
type Bin struct {
	Container
}

func (v Bin) native() *C.GtkBin {
	return (*C.GtkBin)(v.Ptr)
}
func wrapBin(p *C.GtkBin) (v Bin) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBin(p unsafe.Pointer) (v Bin) {
	v.Ptr = p
	return
}
func (v Bin) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBin(p unsafe.Pointer) interface{} {
	return WrapBin(p)
}
func (v Bin) GetType() gobject.Type {
	return gobject.Type(C.gtk_bin_get_type())
}
func (v Bin) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBin(unsafe.Pointer(ptr)), nil
	}
}
func (v Bin) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Bin) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// GetChild is a wrapper around gtk_bin_get_child().
func (bin Bin) GetChild() Widget {
	ret0 := C.gtk_bin_get_child(bin.native())
	return wrapWidget(ret0)
}

// Object Container
type Container struct {
	Widget
}

func (v Container) native() *C.GtkContainer {
	return (*C.GtkContainer)(v.Ptr)
}
func wrapContainer(p *C.GtkContainer) (v Container) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapContainer(p unsafe.Pointer) (v Container) {
	v.Ptr = p
	return
}
func (v Container) IsNil() bool {
	return v.Ptr == nil
}
func IWrapContainer(p unsafe.Pointer) interface{} {
	return WrapContainer(p)
}
func (v Container) GetType() gobject.Type {
	return gobject.Type(C.gtk_container_get_type())
}
func (v Container) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapContainer(unsafe.Pointer(ptr)), nil
	}
}
func (v Container) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Container) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Add is a wrapper around gtk_container_add().
func (container Container) Add(widget Widget) {
	C.gtk_container_add(container.native(), widget.native())
}

// CheckResize is a wrapper around gtk_container_check_resize().
func (container Container) CheckResize() {
	C.gtk_container_check_resize(container.native())
}

// ChildGetProperty is a wrapper around gtk_container_child_get_property().
func (container Container) ChildGetProperty(child Widget, property_name string, value gobject.Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.gtk_container_child_get_property(container.native(), child.native(), property_name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// ChildNotify is a wrapper around gtk_container_child_notify().
func (container Container) ChildNotify(child Widget, child_property string) {
	child_property0 := (*C.gchar)(C.CString(child_property))
	C.gtk_container_child_notify(container.native(), child.native(), child_property0)
	C.free(unsafe.Pointer(child_property0)) /*ch:<stdlib.h>*/
}

// ChildNotifyByPspec is a wrapper around gtk_container_child_notify_by_pspec().
func (container Container) ChildNotifyByPspec(child Widget, pspec gobject.ParamSpec) {
	C.gtk_container_child_notify_by_pspec(container.native(), child.native(), (*C.GParamSpec)(pspec.Ptr))
}

// ChildSetProperty is a wrapper around gtk_container_child_set_property().
func (container Container) ChildSetProperty(child Widget, property_name string, value gobject.Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.gtk_container_child_set_property(container.native(), child.native(), property_name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// ChildType is a wrapper around gtk_container_child_type().
func (container Container) ChildType() /*Gir:GObject*/ gobject.Type {
	ret0 := C.gtk_container_child_type(container.native())
	return /*Gir:GObject*/ gobject.Type(ret0)
}

// GetBorderWidth is a wrapper around gtk_container_get_border_width().
func (container Container) GetBorderWidth() uint {
	ret0 := C.gtk_container_get_border_width(container.native())
	return uint(ret0)
}

// GetChildren is a wrapper around gtk_container_get_children().
func (container Container) GetChildren() glib.List {
	ret0 := C.gtk_container_get_children(container.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWidget(p) }) /*gir:GLib*/
}

// GetFocusChild is a wrapper around gtk_container_get_focus_child().
func (container Container) GetFocusChild() Widget {
	ret0 := C.gtk_container_get_focus_child(container.native())
	return wrapWidget(ret0)
}

// GetFocusHadjustment is a wrapper around gtk_container_get_focus_hadjustment().
func (container Container) GetFocusHadjustment() Adjustment {
	ret0 := C.gtk_container_get_focus_hadjustment(container.native())
	return wrapAdjustment(ret0)
}

// GetFocusVadjustment is a wrapper around gtk_container_get_focus_vadjustment().
func (container Container) GetFocusVadjustment() Adjustment {
	ret0 := C.gtk_container_get_focus_vadjustment(container.native())
	return wrapAdjustment(ret0)
}

// GetPathForChild is a wrapper around gtk_container_get_path_for_child().
func (container Container) GetPathForChild(child Widget) WidgetPath {
	ret0 := C.gtk_container_get_path_for_child(container.native(), child.native())
	return wrapWidgetPath(ret0)
}

// Remove is a wrapper around gtk_container_remove().
func (container Container) Remove(widget Widget) {
	C.gtk_container_remove(container.native(), widget.native())
}

// SetBorderWidth is a wrapper around gtk_container_set_border_width().
func (container Container) SetBorderWidth(border_width uint) {
	C.gtk_container_set_border_width(container.native(), C.guint(border_width))
}

// SetFocusChain is a wrapper around gtk_container_set_focus_chain().
func (container Container) SetFocusChain(focusable_widgets glib.List) {
	C.gtk_container_set_focus_chain(container.native(), (*C.GList)(focusable_widgets.Ptr))
}

// SetFocusChild is a wrapper around gtk_container_set_focus_child().
func (container Container) SetFocusChild(child Widget) {
	C.gtk_container_set_focus_child(container.native(), child.native())
}

// SetFocusHadjustment is a wrapper around gtk_container_set_focus_hadjustment().
func (container Container) SetFocusHadjustment(adjustment Adjustment) {
	C.gtk_container_set_focus_hadjustment(container.native(), adjustment.native())
}

// SetFocusVadjustment is a wrapper around gtk_container_set_focus_vadjustment().
func (container Container) SetFocusVadjustment(adjustment Adjustment) {
	C.gtk_container_set_focus_vadjustment(container.native(), adjustment.native())
}

// UnsetFocusChain is a wrapper around gtk_container_unset_focus_chain().
func (container Container) UnsetFocusChain() {
	C.gtk_container_unset_focus_chain(container.native())
}

// Object Dialog
type Dialog struct {
	Window
}

func (v Dialog) native() *C.GtkDialog {
	return (*C.GtkDialog)(v.Ptr)
}
func wrapDialog(p *C.GtkDialog) (v Dialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDialog(p unsafe.Pointer) (v Dialog) {
	v.Ptr = p
	return
}
func (v Dialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDialog(p unsafe.Pointer) interface{} {
	return WrapDialog(p)
}
func (v Dialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_dialog_get_type())
}
func (v Dialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v Dialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Dialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// DialogNew is a wrapper around gtk_dialog_new().
func DialogNew() Widget {
	ret0 := C.gtk_dialog_new()
	return wrapWidget(ret0)
}

// AddActionWidget is a wrapper around gtk_dialog_add_action_widget().
func (dialog Dialog) AddActionWidget(child Widget, response_id int) {
	C.gtk_dialog_add_action_widget(dialog.native(), child.native(), C.gint(response_id))
}

// AddButton is a wrapper around gtk_dialog_add_button().
func (dialog Dialog) AddButton(button_text string, response_id int) Widget {
	button_text0 := (*C.gchar)(C.CString(button_text))
	ret0 := C.gtk_dialog_add_button(dialog.native(), button_text0, C.gint(response_id))
	C.free(unsafe.Pointer(button_text0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetHeaderBar is a wrapper around gtk_dialog_get_header_bar().
func (dialog Dialog) GetHeaderBar() Widget {
	ret0 := C.gtk_dialog_get_header_bar(dialog.native())
	return wrapWidget(ret0)
}

// GetResponseForWidget is a wrapper around gtk_dialog_get_response_for_widget().
func (dialog Dialog) GetResponseForWidget(widget Widget) int {
	ret0 := C.gtk_dialog_get_response_for_widget(dialog.native(), widget.native())
	return int(ret0)
}

// GetWidgetForResponse is a wrapper around gtk_dialog_get_widget_for_response().
func (dialog Dialog) GetWidgetForResponse(response_id int) Widget {
	ret0 := C.gtk_dialog_get_widget_for_response(dialog.native(), C.gint(response_id))
	return wrapWidget(ret0)
}

// Response is a wrapper around gtk_dialog_response().
func (dialog Dialog) Response(response_id int) {
	C.gtk_dialog_response(dialog.native(), C.gint(response_id))
}

// Run is a wrapper around gtk_dialog_run().
func (dialog Dialog) Run() int {
	ret0 := C.gtk_dialog_run(dialog.native())
	return int(ret0)
}

// SetDefaultResponse is a wrapper around gtk_dialog_set_default_response().
func (dialog Dialog) SetDefaultResponse(response_id int) {
	C.gtk_dialog_set_default_response(dialog.native(), C.gint(response_id))
}

// SetResponseSensitive is a wrapper around gtk_dialog_set_response_sensitive().
func (dialog Dialog) SetResponseSensitive(response_id int, setting bool) {
	C.gtk_dialog_set_response_sensitive(dialog.native(), C.gint(response_id), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// Object Window
type Window struct {
	Bin
}

func (v Window) native() *C.GtkWindow {
	return (*C.GtkWindow)(v.Ptr)
}
func wrapWindow(p *C.GtkWindow) (v Window) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapWindow(p unsafe.Pointer) (v Window) {
	v.Ptr = p
	return
}
func (v Window) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWindow(p unsafe.Pointer) interface{} {
	return WrapWindow(p)
}
func (v Window) GetType() gobject.Type {
	return gobject.Type(C.gtk_window_get_type())
}
func (v Window) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapWindow(unsafe.Pointer(ptr)), nil
	}
}
func (v Window) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Window) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// WindowNew is a wrapper around gtk_window_new().
func WindowNew(type_ WindowType) Widget {
	ret0 := C.gtk_window_new(C.GtkWindowType(type_))
	return wrapWidget(ret0)
}

// ActivateDefault is a wrapper around gtk_window_activate_default().
func (window Window) ActivateDefault() bool {
	ret0 := C.gtk_window_activate_default(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ActivateFocus is a wrapper around gtk_window_activate_focus().
func (window Window) ActivateFocus() bool {
	ret0 := C.gtk_window_activate_focus(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ActivateKey is a wrapper around gtk_window_activate_key().
func (window Window) ActivateKey(event gdk.EventKey) bool {
	ret0 := C.gtk_window_activate_key(window.native(), (*C.GdkEventKey)(event.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// AddAccelGroup is a wrapper around gtk_window_add_accel_group().
func (window Window) AddAccelGroup(accel_group AccelGroup) {
	C.gtk_window_add_accel_group(window.native(), accel_group.native())
}

// AddMnemonic is a wrapper around gtk_window_add_mnemonic().
func (window Window) AddMnemonic(keyval uint, target Widget) {
	C.gtk_window_add_mnemonic(window.native(), C.guint(keyval), target.native())
}

// BeginMoveDrag is a wrapper around gtk_window_begin_move_drag().
func (window Window) BeginMoveDrag(button int, root_x int, root_y int, timestamp uint32) {
	C.gtk_window_begin_move_drag(window.native(), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
}

// BeginResizeDrag is a wrapper around gtk_window_begin_resize_drag().
func (window Window) BeginResizeDrag(edge /*gir:Gdk*/ gdk.WindowEdge, button int, root_x int, root_y int, timestamp uint32) {
	C.gtk_window_begin_resize_drag(window.native(), C.GdkWindowEdge(edge), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
}

// Deiconify is a wrapper around gtk_window_deiconify().
func (window Window) Deiconify() {
	C.gtk_window_deiconify(window.native())
}

// Fullscreen is a wrapper around gtk_window_fullscreen().
func (window Window) Fullscreen() {
	C.gtk_window_fullscreen(window.native())
}

// FullscreenOnMonitor is a wrapper around gtk_window_fullscreen_on_monitor().
func (window Window) FullscreenOnMonitor(screen gdk.Screen, monitor int) {
	C.gtk_window_fullscreen_on_monitor(window.native(), (*C.GdkScreen)(screen.Ptr), C.gint(monitor))
}

// GetDecorated is a wrapper around gtk_window_get_decorated().
func (window Window) GetDecorated() bool {
	ret0 := C.gtk_window_get_decorated(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDefaultSize is a wrapper around gtk_window_get_default_size().
func (window Window) GetDefaultSize() (int, int) {
	var width0 C.gint
	var height0 C.gint
	C.gtk_window_get_default_size(window.native(), &width0, &height0)
	return int(width0), int(height0)
}

// GetDefaultWidget is a wrapper around gtk_window_get_default_widget().
func (window Window) GetDefaultWidget() Widget {
	ret0 := C.gtk_window_get_default_widget(window.native())
	return wrapWidget(ret0)
}

// GetDestroyWithParent is a wrapper around gtk_window_get_destroy_with_parent().
func (window Window) GetDestroyWithParent() bool {
	ret0 := C.gtk_window_get_destroy_with_parent(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFocus is a wrapper around gtk_window_get_focus().
func (window Window) GetFocus() Widget {
	ret0 := C.gtk_window_get_focus(window.native())
	return wrapWidget(ret0)
}

// GetFocusOnMap is a wrapper around gtk_window_get_focus_on_map().
func (window Window) GetFocusOnMap() bool {
	ret0 := C.gtk_window_get_focus_on_map(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFocusVisible is a wrapper around gtk_window_get_focus_visible().
func (window Window) GetFocusVisible() bool {
	ret0 := C.gtk_window_get_focus_visible(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetGravity is a wrapper around gtk_window_get_gravity().
func (window Window) GetGravity() /*gir:Gdk*/ gdk.Gravity {
	ret0 := C.gtk_window_get_gravity(window.native())
	return /*gir:Gdk*/ gdk.Gravity(ret0)
}

// GetGroup is a wrapper around gtk_window_get_group().
func (window Window) GetGroup() WindowGroup {
	ret0 := C.gtk_window_get_group(window.native())
	return wrapWindowGroup(ret0)
}

// GetHideTitlebarWhenMaximized is a wrapper around gtk_window_get_hide_titlebar_when_maximized().
func (window Window) GetHideTitlebarWhenMaximized() bool {
	ret0 := C.gtk_window_get_hide_titlebar_when_maximized(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIcon is a wrapper around gtk_window_get_icon().
func (window Window) GetIcon() gdkpixbuf.Pixbuf {
	ret0 := C.gtk_window_get_icon(window.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetIconName is a wrapper around gtk_window_get_icon_name().
func (window Window) GetIconName() string {
	ret0 := C.gtk_window_get_icon_name(window.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetMnemonicsVisible is a wrapper around gtk_window_get_mnemonics_visible().
func (window Window) GetMnemonicsVisible() bool {
	ret0 := C.gtk_window_get_mnemonics_visible(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetModal is a wrapper around gtk_window_get_modal().
func (window Window) GetModal() bool {
	ret0 := C.gtk_window_get_modal(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPosition is a wrapper around gtk_window_get_position().
func (window Window) GetPosition() (int, int) {
	var root_x0 C.gint
	var root_y0 C.gint
	C.gtk_window_get_position(window.native(), &root_x0, &root_y0)
	return int(root_x0), int(root_y0)
}

// GetResizable is a wrapper around gtk_window_get_resizable().
func (window Window) GetResizable() bool {
	ret0 := C.gtk_window_get_resizable(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRole is a wrapper around gtk_window_get_role().
func (window Window) GetRole() string {
	ret0 := C.gtk_window_get_role(window.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSize is a wrapper around gtk_window_get_size().
func (window Window) GetSize() (int, int) {
	var width0 C.gint
	var height0 C.gint
	C.gtk_window_get_size(window.native(), &width0, &height0)
	return int(width0), int(height0)
}

// GetSkipPagerHint is a wrapper around gtk_window_get_skip_pager_hint().
func (window Window) GetSkipPagerHint() bool {
	ret0 := C.gtk_window_get_skip_pager_hint(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSkipTaskbarHint is a wrapper around gtk_window_get_skip_taskbar_hint().
func (window Window) GetSkipTaskbarHint() bool {
	ret0 := C.gtk_window_get_skip_taskbar_hint(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTitle is a wrapper around gtk_window_get_title().
func (window Window) GetTitle() string {
	ret0 := C.gtk_window_get_title(window.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTitlebar is a wrapper around gtk_window_get_titlebar().
func (window Window) GetTitlebar() Widget {
	ret0 := C.gtk_window_get_titlebar(window.native())
	return wrapWidget(ret0)
}

// GetTransientFor is a wrapper around gtk_window_get_transient_for().
func (window Window) GetTransientFor() Window {
	ret0 := C.gtk_window_get_transient_for(window.native())
	return wrapWindow(ret0)
}

// GetUrgencyHint is a wrapper around gtk_window_get_urgency_hint().
func (window Window) GetUrgencyHint() bool {
	ret0 := C.gtk_window_get_urgency_hint(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetWindowType is a wrapper around gtk_window_get_window_type().
func (window Window) GetWindowType() WindowType {
	ret0 := C.gtk_window_get_window_type(window.native())
	return WindowType(ret0)
}

// HasGroup is a wrapper around gtk_window_has_group().
func (window Window) HasGroup() bool {
	ret0 := C.gtk_window_has_group(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasToplevelFocus is a wrapper around gtk_window_has_toplevel_focus().
func (window Window) HasToplevelFocus() bool {
	ret0 := C.gtk_window_has_toplevel_focus(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Iconify is a wrapper around gtk_window_iconify().
func (window Window) Iconify() {
	C.gtk_window_iconify(window.native())
}

// IsActive is a wrapper around gtk_window_is_active().
func (window Window) IsActive() bool {
	ret0 := C.gtk_window_is_active(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsMaximized is a wrapper around gtk_window_is_maximized().
func (window Window) IsMaximized() bool {
	ret0 := C.gtk_window_is_maximized(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Maximize is a wrapper around gtk_window_maximize().
func (window Window) Maximize() {
	C.gtk_window_maximize(window.native())
}

// Move is a wrapper around gtk_window_move().
func (window Window) Move(x int, y int) {
	C.gtk_window_move(window.native(), C.gint(x), C.gint(y))
}

// Present is a wrapper around gtk_window_present().
func (window Window) Present() {
	C.gtk_window_present(window.native())
}

// PresentWithTime is a wrapper around gtk_window_present_with_time().
func (window Window) PresentWithTime(timestamp uint32) {
	C.gtk_window_present_with_time(window.native(), C.guint32(timestamp))
}

// RemoveAccelGroup is a wrapper around gtk_window_remove_accel_group().
func (window Window) RemoveAccelGroup(accel_group AccelGroup) {
	C.gtk_window_remove_accel_group(window.native(), accel_group.native())
}

// RemoveMnemonic is a wrapper around gtk_window_remove_mnemonic().
func (window Window) RemoveMnemonic(keyval uint, target Widget) {
	C.gtk_window_remove_mnemonic(window.native(), C.guint(keyval), target.native())
}

// Resize is a wrapper around gtk_window_resize().
func (window Window) Resize(width int, height int) {
	C.gtk_window_resize(window.native(), C.gint(width), C.gint(height))
}

// SetAcceptFocus is a wrapper around gtk_window_set_accept_focus().
func (window Window) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetApplication is a wrapper around gtk_window_set_application().
func (window Window) SetApplication(application Application) {
	C.gtk_window_set_application(window.native(), application.native())
}

// SetAttachedTo is a wrapper around gtk_window_set_attached_to().
func (window Window) SetAttachedTo(attach_widget Widget) {
	C.gtk_window_set_attached_to(window.native(), attach_widget.native())
}

// SetDecorated is a wrapper around gtk_window_set_decorated().
func (window Window) SetDecorated(setting bool) {
	C.gtk_window_set_decorated(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetDefault is a wrapper around gtk_window_set_default().
func (window Window) SetDefault(default_widget Widget) {
	C.gtk_window_set_default(window.native(), default_widget.native())
}

// SetDefaultSize is a wrapper around gtk_window_set_default_size().
func (window Window) SetDefaultSize(width int, height int) {
	C.gtk_window_set_default_size(window.native(), C.gint(width), C.gint(height))
}

// SetDeletable is a wrapper around gtk_window_set_deletable().
func (window Window) SetDeletable(setting bool) {
	C.gtk_window_set_deletable(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetDestroyWithParent is a wrapper around gtk_window_set_destroy_with_parent().
func (window Window) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetFocus is a wrapper around gtk_window_set_focus().
func (window Window) SetFocus(focus Widget) {
	C.gtk_window_set_focus(window.native(), focus.native())
}

// SetFocusOnMap is a wrapper around gtk_window_set_focus_on_map().
func (window Window) SetFocusOnMap(setting bool) {
	C.gtk_window_set_focus_on_map(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetFocusVisible is a wrapper around gtk_window_set_focus_visible().
func (window Window) SetFocusVisible(setting bool) {
	C.gtk_window_set_focus_visible(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetHasUserRefCount is a wrapper around gtk_window_set_has_user_ref_count().
func (window Window) SetHasUserRefCount(setting bool) {
	C.gtk_window_set_has_user_ref_count(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetHideTitlebarWhenMaximized is a wrapper around gtk_window_set_hide_titlebar_when_maximized().
func (window Window) SetHideTitlebarWhenMaximized(setting bool) {
	C.gtk_window_set_hide_titlebar_when_maximized(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetIconFromFile is a wrapper around gtk_window_set_icon_from_file().
func (window Window) SetIconFromFile(filename string) (bool, error) {
	filename0 := (*C.gchar)(C.CString(filename))
	var err glib.Error
	ret0 := C.gtk_window_set_icon_from_file(window.native(), filename0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetIconName is a wrapper around gtk_window_set_icon_name().
func (window Window) SetIconName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_window_set_icon_name(window.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetKeepAbove is a wrapper around gtk_window_set_keep_above().
func (window Window) SetKeepAbove(setting bool) {
	C.gtk_window_set_keep_above(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetKeepBelow is a wrapper around gtk_window_set_keep_below().
func (window Window) SetKeepBelow(setting bool) {
	C.gtk_window_set_keep_below(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetMnemonicModifier is a wrapper around gtk_window_set_mnemonic_modifier().
func (window Window) SetMnemonicModifier(modifier /*gir:Gdk*/ gdk.ModifierType) {
	C.gtk_window_set_mnemonic_modifier(window.native(), C.GdkModifierType(modifier))
}

// SetMnemonicsVisible is a wrapper around gtk_window_set_mnemonics_visible().
func (window Window) SetMnemonicsVisible(setting bool) {
	C.gtk_window_set_mnemonics_visible(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetModal is a wrapper around gtk_window_set_modal().
func (window Window) SetModal(modal bool) {
	C.gtk_window_set_modal(window.native(), C.gboolean(util.Bool2Int(modal)) /*go:.util*/)
}

// SetPosition is a wrapper around gtk_window_set_position().
func (window Window) SetPosition(position WindowPosition) {
	C.gtk_window_set_position(window.native(), C.GtkWindowPosition(position))
}

// SetResizable is a wrapper around gtk_window_set_resizable().
func (window Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(window.native(), C.gboolean(util.Bool2Int(resizable)) /*go:.util*/)
}

// SetRole is a wrapper around gtk_window_set_role().
func (window Window) SetRole(role string) {
	role0 := (*C.gchar)(C.CString(role))
	C.gtk_window_set_role(window.native(), role0)
	C.free(unsafe.Pointer(role0)) /*ch:<stdlib.h>*/
}

// SetSkipPagerHint is a wrapper around gtk_window_set_skip_pager_hint().
func (window Window) SetSkipPagerHint(setting bool) {
	C.gtk_window_set_skip_pager_hint(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetSkipTaskbarHint is a wrapper around gtk_window_set_skip_taskbar_hint().
func (window Window) SetSkipTaskbarHint(setting bool) {
	C.gtk_window_set_skip_taskbar_hint(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetStartupId is a wrapper around gtk_window_set_startup_id().
func (window Window) SetStartupId(startup_id string) {
	startup_id0 := (*C.gchar)(C.CString(startup_id))
	C.gtk_window_set_startup_id(window.native(), startup_id0)
	C.free(unsafe.Pointer(startup_id0)) /*ch:<stdlib.h>*/
}

// SetTitle is a wrapper around gtk_window_set_title().
func (window Window) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.gtk_window_set_title(window.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetTitlebar is a wrapper around gtk_window_set_titlebar().
func (window Window) SetTitlebar(titlebar Widget) {
	C.gtk_window_set_titlebar(window.native(), titlebar.native())
}

// SetTransientFor is a wrapper around gtk_window_set_transient_for().
func (window Window) SetTransientFor(parent Window) {
	C.gtk_window_set_transient_for(window.native(), parent.native())
}

// SetUrgencyHint is a wrapper around gtk_window_set_urgency_hint().
func (window Window) SetUrgencyHint(setting bool) {
	C.gtk_window_set_urgency_hint(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// Stick is a wrapper around gtk_window_stick().
func (window Window) Stick() {
	C.gtk_window_stick(window.native())
}

// Unfullscreen is a wrapper around gtk_window_unfullscreen().
func (window Window) Unfullscreen() {
	C.gtk_window_unfullscreen(window.native())
}

// Unmaximize is a wrapper around gtk_window_unmaximize().
func (window Window) Unmaximize() {
	C.gtk_window_unmaximize(window.native())
}

// Unstick is a wrapper around gtk_window_unstick().
func (window Window) Unstick() {
	C.gtk_window_unstick(window.native())
}

// WindowGetDefaultIconName is a wrapper around gtk_window_get_default_icon_name().
func WindowGetDefaultIconName() string {
	ret0 := C.gtk_window_get_default_icon_name()
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// WindowListToplevels is a wrapper around gtk_window_list_toplevels().
func WindowListToplevels() glib.List {
	ret0 := C.gtk_window_list_toplevels()
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWidget(p) }) /*gir:GLib*/
}

// WindowSetAutoStartupNotification is a wrapper around gtk_window_set_auto_startup_notification().
func WindowSetAutoStartupNotification(setting bool) {
	C.gtk_window_set_auto_startup_notification(C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// WindowSetDefaultIconFromFile is a wrapper around gtk_window_set_default_icon_from_file().
func WindowSetDefaultIconFromFile(filename string) (bool, error) {
	filename0 := (*C.gchar)(C.CString(filename))
	var err glib.Error
	ret0 := C.gtk_window_set_default_icon_from_file(filename0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// WindowSetDefaultIconName is a wrapper around gtk_window_set_default_icon_name().
func WindowSetDefaultIconName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_window_set_default_icon_name(name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// WindowSetInteractiveDebugging is a wrapper around gtk_window_set_interactive_debugging().
func WindowSetInteractiveDebugging(enable bool) {
	C.gtk_window_set_interactive_debugging(C.gboolean(util.Bool2Int(enable)) /*go:.util*/)
}

// Object WindowGroup
type WindowGroup struct {
	gobject.Object
}

func (v WindowGroup) native() *C.GtkWindowGroup {
	return (*C.GtkWindowGroup)(v.Ptr)
}
func wrapWindowGroup(p *C.GtkWindowGroup) (v WindowGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapWindowGroup(p unsafe.Pointer) (v WindowGroup) {
	v.Ptr = p
	return
}
func (v WindowGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWindowGroup(p unsafe.Pointer) interface{} {
	return WrapWindowGroup(p)
}
func (v WindowGroup) GetType() gobject.Type {
	return gobject.Type(C.gtk_window_group_get_type())
}
func (v WindowGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapWindowGroup(unsafe.Pointer(ptr)), nil
	}
}

// WindowGroupNew is a wrapper around gtk_window_group_new().
func WindowGroupNew() WindowGroup {
	ret0 := C.gtk_window_group_new()
	return wrapWindowGroup(ret0)
}

// AddWindow is a wrapper around gtk_window_group_add_window().
func (window_group WindowGroup) AddWindow(window Window) {
	C.gtk_window_group_add_window(window_group.native(), window.native())
}

// GetCurrentDeviceGrab is a wrapper around gtk_window_group_get_current_device_grab().
func (window_group WindowGroup) GetCurrentDeviceGrab(device gdk.Device) Widget {
	ret0 := C.gtk_window_group_get_current_device_grab(window_group.native(), (*C.GdkDevice)(device.Ptr))
	return wrapWidget(ret0)
}

// GetCurrentGrab is a wrapper around gtk_window_group_get_current_grab().
func (window_group WindowGroup) GetCurrentGrab() Widget {
	ret0 := C.gtk_window_group_get_current_grab(window_group.native())
	return wrapWidget(ret0)
}

// ListWindows is a wrapper around gtk_window_group_list_windows().
func (window_group WindowGroup) ListWindows() glib.List {
	ret0 := C.gtk_window_group_list_windows(window_group.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// RemoveWindow is a wrapper around gtk_window_group_remove_window().
func (window_group WindowGroup) RemoveWindow(window Window) {
	C.gtk_window_group_remove_window(window_group.native(), window.native())
}

// Object AboutDialog
type AboutDialog struct {
	Dialog
}

func (v AboutDialog) native() *C.GtkAboutDialog {
	return (*C.GtkAboutDialog)(v.Ptr)
}
func wrapAboutDialog(p *C.GtkAboutDialog) (v AboutDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAboutDialog(p unsafe.Pointer) (v AboutDialog) {
	v.Ptr = p
	return
}
func (v AboutDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAboutDialog(p unsafe.Pointer) interface{} {
	return WrapAboutDialog(p)
}
func (v AboutDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_about_dialog_get_type())
}
func (v AboutDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAboutDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v AboutDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v AboutDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// AboutDialogNew is a wrapper around gtk_about_dialog_new().
func AboutDialogNew() Widget {
	ret0 := C.gtk_about_dialog_new()
	return wrapWidget(ret0)
}

// AddCreditSection is a wrapper around gtk_about_dialog_add_credit_section().
func (about AboutDialog) AddCreditSection(section_name string, people []string) {
	section_name0 := (*C.gchar)(C.CString(section_name))
	people0 := make([]*C.gchar, len(people))
	for idx, elemG := range people {
		elem := (*C.gchar)(C.CString(elemG))
		people0[idx] = elem
	}
	var people0Ptr **C.gchar
	if len(people0) > 0 {
		people0Ptr = &people0[0]
	}
	C.gtk_about_dialog_add_credit_section(about.native(), section_name0, people0Ptr)
	C.free(unsafe.Pointer(section_name0)) /*ch:<stdlib.h>*/
	for _, elem := range people0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// GetArtists is a wrapper around gtk_about_dialog_get_artists().
func (about AboutDialog) GetArtists() []string {
	ret0 := C.gtk_about_dialog_get_artists(about.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	return ret
}

// GetAuthors is a wrapper around gtk_about_dialog_get_authors().
func (about AboutDialog) GetAuthors() []string {
	ret0 := C.gtk_about_dialog_get_authors(about.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	return ret
}

// GetComments is a wrapper around gtk_about_dialog_get_comments().
func (about AboutDialog) GetComments() string {
	ret0 := C.gtk_about_dialog_get_comments(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetCopyright is a wrapper around gtk_about_dialog_get_copyright().
func (about AboutDialog) GetCopyright() string {
	ret0 := C.gtk_about_dialog_get_copyright(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDocumenters is a wrapper around gtk_about_dialog_get_documenters().
func (about AboutDialog) GetDocumenters() []string {
	ret0 := C.gtk_about_dialog_get_documenters(about.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	return ret
}

// GetLicense is a wrapper around gtk_about_dialog_get_license().
func (about AboutDialog) GetLicense() string {
	ret0 := C.gtk_about_dialog_get_license(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLicenseType is a wrapper around gtk_about_dialog_get_license_type().
func (about AboutDialog) GetLicenseType() License {
	ret0 := C.gtk_about_dialog_get_license_type(about.native())
	return License(ret0)
}

// GetLogo is a wrapper around gtk_about_dialog_get_logo().
func (about AboutDialog) GetLogo() gdkpixbuf.Pixbuf {
	ret0 := C.gtk_about_dialog_get_logo(about.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetProgramName is a wrapper around gtk_about_dialog_get_program_name().
func (about AboutDialog) GetProgramName() string {
	ret0 := C.gtk_about_dialog_get_program_name(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTranslatorCredits is a wrapper around gtk_about_dialog_get_translator_credits().
func (about AboutDialog) GetTranslatorCredits() string {
	ret0 := C.gtk_about_dialog_get_translator_credits(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetVersion is a wrapper around gtk_about_dialog_get_version().
func (about AboutDialog) GetVersion() string {
	ret0 := C.gtk_about_dialog_get_version(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWebsite is a wrapper around gtk_about_dialog_get_website().
func (about AboutDialog) GetWebsite() string {
	ret0 := C.gtk_about_dialog_get_website(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWebsiteLabel is a wrapper around gtk_about_dialog_get_website_label().
func (about AboutDialog) GetWebsiteLabel() string {
	ret0 := C.gtk_about_dialog_get_website_label(about.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWrapLicense is a wrapper around gtk_about_dialog_get_wrap_license().
func (about AboutDialog) GetWrapLicense() bool {
	ret0 := C.gtk_about_dialog_get_wrap_license(about.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetArtists is a wrapper around gtk_about_dialog_set_artists().
func (about AboutDialog) SetArtists(artists []string) {
	artists0 := make([]*C.gchar, len(artists))
	for idx, elemG := range artists {
		elem := (*C.gchar)(C.CString(elemG))
		artists0[idx] = elem
	}
	var artists0Ptr **C.gchar
	if len(artists0) > 0 {
		artists0Ptr = &artists0[0]
	}
	C.gtk_about_dialog_set_artists(about.native(), artists0Ptr)
	for _, elem := range artists0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// SetAuthors is a wrapper around gtk_about_dialog_set_authors().
func (about AboutDialog) SetAuthors(authors []string) {
	authors0 := make([]*C.gchar, len(authors))
	for idx, elemG := range authors {
		elem := (*C.gchar)(C.CString(elemG))
		authors0[idx] = elem
	}
	var authors0Ptr **C.gchar
	if len(authors0) > 0 {
		authors0Ptr = &authors0[0]
	}
	C.gtk_about_dialog_set_authors(about.native(), authors0Ptr)
	for _, elem := range authors0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// SetComments is a wrapper around gtk_about_dialog_set_comments().
func (about AboutDialog) SetComments(comments string) {
	comments0 := (*C.gchar)(C.CString(comments))
	C.gtk_about_dialog_set_comments(about.native(), comments0)
	C.free(unsafe.Pointer(comments0)) /*ch:<stdlib.h>*/
}

// SetCopyright is a wrapper around gtk_about_dialog_set_copyright().
func (about AboutDialog) SetCopyright(copyright string) {
	copyright0 := (*C.gchar)(C.CString(copyright))
	C.gtk_about_dialog_set_copyright(about.native(), copyright0)
	C.free(unsafe.Pointer(copyright0)) /*ch:<stdlib.h>*/
}

// SetDocumenters is a wrapper around gtk_about_dialog_set_documenters().
func (about AboutDialog) SetDocumenters(documenters []string) {
	documenters0 := make([]*C.gchar, len(documenters))
	for idx, elemG := range documenters {
		elem := (*C.gchar)(C.CString(elemG))
		documenters0[idx] = elem
	}
	var documenters0Ptr **C.gchar
	if len(documenters0) > 0 {
		documenters0Ptr = &documenters0[0]
	}
	C.gtk_about_dialog_set_documenters(about.native(), documenters0Ptr)
	for _, elem := range documenters0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// SetLicense is a wrapper around gtk_about_dialog_set_license().
func (about AboutDialog) SetLicense(license string) {
	license0 := (*C.gchar)(C.CString(license))
	C.gtk_about_dialog_set_license(about.native(), license0)
	C.free(unsafe.Pointer(license0)) /*ch:<stdlib.h>*/
}

// SetLicenseType is a wrapper around gtk_about_dialog_set_license_type().
func (about AboutDialog) SetLicenseType(license_type License) {
	C.gtk_about_dialog_set_license_type(about.native(), C.GtkLicense(license_type))
}

// SetLogo is a wrapper around gtk_about_dialog_set_logo().
func (about AboutDialog) SetLogo(logo gdkpixbuf.Pixbuf) {
	C.gtk_about_dialog_set_logo(about.native(), (*C.GdkPixbuf)(logo.Ptr))
}

// SetLogoIconName is a wrapper around gtk_about_dialog_set_logo_icon_name().
func (about AboutDialog) SetLogoIconName(icon_name string) {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	C.gtk_about_dialog_set_logo_icon_name(about.native(), icon_name0)
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
}

// SetProgramName is a wrapper around gtk_about_dialog_set_program_name().
func (about AboutDialog) SetProgramName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_about_dialog_set_program_name(about.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetTranslatorCredits is a wrapper around gtk_about_dialog_set_translator_credits().
func (about AboutDialog) SetTranslatorCredits(translator_credits string) {
	translator_credits0 := (*C.gchar)(C.CString(translator_credits))
	C.gtk_about_dialog_set_translator_credits(about.native(), translator_credits0)
	C.free(unsafe.Pointer(translator_credits0)) /*ch:<stdlib.h>*/
}

// SetVersion is a wrapper around gtk_about_dialog_set_version().
func (about AboutDialog) SetVersion(version string) {
	version0 := (*C.gchar)(C.CString(version))
	C.gtk_about_dialog_set_version(about.native(), version0)
	C.free(unsafe.Pointer(version0)) /*ch:<stdlib.h>*/
}

// SetWebsite is a wrapper around gtk_about_dialog_set_website().
func (about AboutDialog) SetWebsite(website string) {
	website0 := (*C.gchar)(C.CString(website))
	C.gtk_about_dialog_set_website(about.native(), website0)
	C.free(unsafe.Pointer(website0)) /*ch:<stdlib.h>*/
}

// SetWebsiteLabel is a wrapper around gtk_about_dialog_set_website_label().
func (about AboutDialog) SetWebsiteLabel(website_label string) {
	website_label0 := (*C.gchar)(C.CString(website_label))
	C.gtk_about_dialog_set_website_label(about.native(), website_label0)
	C.free(unsafe.Pointer(website_label0)) /*ch:<stdlib.h>*/
}

// SetWrapLicense is a wrapper around gtk_about_dialog_set_wrap_license().
func (about AboutDialog) SetWrapLicense(wrap_license bool) {
	C.gtk_about_dialog_set_wrap_license(about.native(), C.gboolean(util.Bool2Int(wrap_license)) /*go:.util*/)
}

// Object ActionBar
type ActionBar struct {
	Bin
}

func (v ActionBar) native() *C.GtkActionBar {
	return (*C.GtkActionBar)(v.Ptr)
}
func wrapActionBar(p *C.GtkActionBar) (v ActionBar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapActionBar(p unsafe.Pointer) (v ActionBar) {
	v.Ptr = p
	return
}
func (v ActionBar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapActionBar(p unsafe.Pointer) interface{} {
	return WrapActionBar(p)
}
func (v ActionBar) GetType() gobject.Type {
	return gobject.Type(C.gtk_action_bar_get_type())
}
func (v ActionBar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapActionBar(unsafe.Pointer(ptr)), nil
	}
}
func (v ActionBar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ActionBar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ActionBarNew is a wrapper around gtk_action_bar_new().
func ActionBarNew() Widget {
	ret0 := C.gtk_action_bar_new()
	return wrapWidget(ret0)
}

// GetCenterWidget is a wrapper around gtk_action_bar_get_center_widget().
func (action_bar ActionBar) GetCenterWidget() Widget {
	ret0 := C.gtk_action_bar_get_center_widget(action_bar.native())
	return wrapWidget(ret0)
}

// PackEnd is a wrapper around gtk_action_bar_pack_end().
func (action_bar ActionBar) PackEnd(child Widget) {
	C.gtk_action_bar_pack_end(action_bar.native(), child.native())
}

// PackStart is a wrapper around gtk_action_bar_pack_start().
func (action_bar ActionBar) PackStart(child Widget) {
	C.gtk_action_bar_pack_start(action_bar.native(), child.native())
}

// SetCenterWidget is a wrapper around gtk_action_bar_set_center_widget().
func (action_bar ActionBar) SetCenterWidget(center_widget Widget) {
	C.gtk_action_bar_set_center_widget(action_bar.native(), center_widget.native())
}

// Object ApplicationWindow
type ApplicationWindow struct {
	Window
}

func (v ApplicationWindow) native() *C.GtkApplicationWindow {
	return (*C.GtkApplicationWindow)(v.Ptr)
}
func wrapApplicationWindow(p *C.GtkApplicationWindow) (v ApplicationWindow) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapApplicationWindow(p unsafe.Pointer) (v ApplicationWindow) {
	v.Ptr = p
	return
}
func (v ApplicationWindow) IsNil() bool {
	return v.Ptr == nil
}
func IWrapApplicationWindow(p unsafe.Pointer) interface{} {
	return WrapApplicationWindow(p)
}
func (v ApplicationWindow) GetType() gobject.Type {
	return gobject.Type(C.gtk_application_window_get_type())
}
func (v ApplicationWindow) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapApplicationWindow(unsafe.Pointer(ptr)), nil
	}
}
func (v ApplicationWindow) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ApplicationWindow) ActionGroup() gio.ActionGroup {
	return gio.WrapActionGroup(v.Ptr) /*gir:Gio*/
}
func (v ApplicationWindow) ActionMap() gio.ActionMap {
	return gio.WrapActionMap(v.Ptr) /*gir:Gio*/
}
func (v ApplicationWindow) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ApplicationWindowNew is a wrapper around gtk_application_window_new().
func ApplicationWindowNew(application Application) Widget {
	ret0 := C.gtk_application_window_new(application.native())
	return wrapWidget(ret0)
}

// GetHelpOverlay is a wrapper around gtk_application_window_get_help_overlay().
func (window ApplicationWindow) GetHelpOverlay() ShortcutsWindow {
	ret0 := C.gtk_application_window_get_help_overlay(window.native())
	return wrapShortcutsWindow(ret0)
}

// GetId is a wrapper around gtk_application_window_get_id().
func (window ApplicationWindow) GetId() uint {
	ret0 := C.gtk_application_window_get_id(window.native())
	return uint(ret0)
}

// GetShowMenubar is a wrapper around gtk_application_window_get_show_menubar().
func (window ApplicationWindow) GetShowMenubar() bool {
	ret0 := C.gtk_application_window_get_show_menubar(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetShowMenubar is a wrapper around gtk_application_window_set_show_menubar().
func (window ApplicationWindow) SetShowMenubar(show_menubar bool) {
	C.gtk_application_window_set_show_menubar(window.native(), C.gboolean(util.Bool2Int(show_menubar)) /*go:.util*/)
}

// Object AspectFrame
type AspectFrame struct {
	Frame
}

func (v AspectFrame) native() *C.GtkAspectFrame {
	return (*C.GtkAspectFrame)(v.Ptr)
}
func wrapAspectFrame(p *C.GtkAspectFrame) (v AspectFrame) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAspectFrame(p unsafe.Pointer) (v AspectFrame) {
	v.Ptr = p
	return
}
func (v AspectFrame) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAspectFrame(p unsafe.Pointer) interface{} {
	return WrapAspectFrame(p)
}
func (v AspectFrame) GetType() gobject.Type {
	return gobject.Type(C.gtk_aspect_frame_get_type())
}
func (v AspectFrame) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAspectFrame(unsafe.Pointer(ptr)), nil
	}
}
func (v AspectFrame) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v AspectFrame) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// AspectFrameNew is a wrapper around gtk_aspect_frame_new().
func AspectFrameNew(label string, xalign float32, yalign float32, ratio float32, obey_child bool) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_aspect_frame_new(label0, C.gfloat(xalign), C.gfloat(yalign), C.gfloat(ratio), C.gboolean(util.Bool2Int(obey_child)) /*go:.util*/)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// Set is a wrapper around gtk_aspect_frame_set().
func (aspect_frame AspectFrame) Set(xalign float32, yalign float32, ratio float32, obey_child bool) {
	C.gtk_aspect_frame_set(aspect_frame.native(), C.gfloat(xalign), C.gfloat(yalign), C.gfloat(ratio), C.gboolean(util.Bool2Int(obey_child)) /*go:.util*/)
}

// Object Frame
type Frame struct {
	Bin
}

func (v Frame) native() *C.GtkFrame {
	return (*C.GtkFrame)(v.Ptr)
}
func wrapFrame(p *C.GtkFrame) (v Frame) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFrame(p unsafe.Pointer) (v Frame) {
	v.Ptr = p
	return
}
func (v Frame) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFrame(p unsafe.Pointer) interface{} {
	return WrapFrame(p)
}
func (v Frame) GetType() gobject.Type {
	return gobject.Type(C.gtk_frame_get_type())
}
func (v Frame) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFrame(unsafe.Pointer(ptr)), nil
	}
}
func (v Frame) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Frame) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// FrameNew is a wrapper around gtk_frame_new().
func FrameNew(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_frame_new(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetLabel is a wrapper around gtk_frame_get_label().
func (frame Frame) GetLabel() string {
	ret0 := C.gtk_frame_get_label(frame.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLabelAlign is a wrapper around gtk_frame_get_label_align().
func (frame Frame) GetLabelAlign() (float32, float32) {
	var xalign0 C.gfloat
	var yalign0 C.gfloat
	C.gtk_frame_get_label_align(frame.native(), &xalign0, &yalign0)
	return float32(xalign0), float32(yalign0)
}

// GetLabelWidget is a wrapper around gtk_frame_get_label_widget().
func (frame Frame) GetLabelWidget() Widget {
	ret0 := C.gtk_frame_get_label_widget(frame.native())
	return wrapWidget(ret0)
}

// GetShadowType is a wrapper around gtk_frame_get_shadow_type().
func (frame Frame) GetShadowType() ShadowType {
	ret0 := C.gtk_frame_get_shadow_type(frame.native())
	return ShadowType(ret0)
}

// SetLabel is a wrapper around gtk_frame_set_label().
func (frame Frame) SetLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_frame_set_label(frame.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetLabelAlign is a wrapper around gtk_frame_set_label_align().
func (frame Frame) SetLabelAlign(xalign float32, yalign float32) {
	C.gtk_frame_set_label_align(frame.native(), C.gfloat(xalign), C.gfloat(yalign))
}

// SetLabelWidget is a wrapper around gtk_frame_set_label_widget().
func (frame Frame) SetLabelWidget(label_widget Widget) {
	C.gtk_frame_set_label_widget(frame.native(), label_widget.native())
}

// SetShadowType is a wrapper around gtk_frame_set_shadow_type().
func (frame Frame) SetShadowType(type_ ShadowType) {
	C.gtk_frame_set_shadow_type(frame.native(), C.GtkShadowType(type_))
}

// Object FileFilter
type FileFilter struct {
	gobject.InitiallyUnowned
}

func (v FileFilter) native() *C.GtkFileFilter {
	return (*C.GtkFileFilter)(v.Ptr)
}
func wrapFileFilter(p *C.GtkFileFilter) (v FileFilter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileFilter(p unsafe.Pointer) (v FileFilter) {
	v.Ptr = p
	return
}
func (v FileFilter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileFilter(p unsafe.Pointer) interface{} {
	return WrapFileFilter(p)
}
func (v FileFilter) GetType() gobject.Type {
	return gobject.Type(C.gtk_file_filter_get_type())
}
func (v FileFilter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileFilter(unsafe.Pointer(ptr)), nil
	}
}
func (v FileFilter) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// FileFilterNew is a wrapper around gtk_file_filter_new().
func FileFilterNew() FileFilter {
	ret0 := C.gtk_file_filter_new()
	return wrapFileFilter(ret0)
}

// FileFilterNewFromGvariant is a wrapper around gtk_file_filter_new_from_gvariant().
func FileFilterNewFromGvariant(variant glib.Variant) FileFilter {
	ret0 := C.gtk_file_filter_new_from_gvariant((*C.GVariant)(variant.Ptr))
	return wrapFileFilter(ret0)
}

// AddMimeType is a wrapper around gtk_file_filter_add_mime_type().
func (filter FileFilter) AddMimeType(mime_type string) {
	mime_type0 := (*C.gchar)(C.CString(mime_type))
	C.gtk_file_filter_add_mime_type(filter.native(), mime_type0)
	C.free(unsafe.Pointer(mime_type0)) /*ch:<stdlib.h>*/
}

// AddPattern is a wrapper around gtk_file_filter_add_pattern().
func (filter FileFilter) AddPattern(pattern string) {
	pattern0 := (*C.gchar)(C.CString(pattern))
	C.gtk_file_filter_add_pattern(filter.native(), pattern0)
	C.free(unsafe.Pointer(pattern0)) /*ch:<stdlib.h>*/
}

// AddPixbufFormats is a wrapper around gtk_file_filter_add_pixbuf_formats().
func (filter FileFilter) AddPixbufFormats() {
	C.gtk_file_filter_add_pixbuf_formats(filter.native())
}

// Filter is a wrapper around gtk_file_filter_filter().
func (filter FileFilter) Filter(filter_info FileFilterInfo) bool {
	ret0 := C.gtk_file_filter_filter(filter.native(), filter_info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetName is a wrapper around gtk_file_filter_get_name().
func (filter FileFilter) GetName() string {
	ret0 := C.gtk_file_filter_get_name(filter.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNeeded is a wrapper around gtk_file_filter_get_needed().
func (filter FileFilter) GetNeeded() FileFilterFlags {
	ret0 := C.gtk_file_filter_get_needed(filter.native())
	return FileFilterFlags(ret0)
}

// ToGvariant is a wrapper around gtk_file_filter_to_gvariant().
func (filter FileFilter) ToGvariant() glib.Variant {
	ret0 := C.gtk_file_filter_to_gvariant(filter.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// Interface RecentChooser
type RecentChooser struct {
	Ptr unsafe.Pointer
}

func (v RecentChooser) native() *C.GtkRecentChooser {
	return (*C.GtkRecentChooser)(v.Ptr)
}
func wrapRecentChooser(p *C.GtkRecentChooser) RecentChooser {
	return RecentChooser{unsafe.Pointer(p)}
}
func WrapRecentChooser(p unsafe.Pointer) RecentChooser {
	return RecentChooser{p}
}
func (v RecentChooser) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentChooser(p unsafe.Pointer) interface{} {
	return WrapRecentChooser(p)
}
func (v RecentChooser) GetType() gobject.Type {
	return gobject.Type(C.gtk_recent_chooser_get_type())
}
func (v RecentChooser) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRecentChooser(unsafe.Pointer(ptr)), nil
	}
}

// AddFilter is a wrapper around gtk_recent_chooser_add_filter().
func (chooser RecentChooser) AddFilter(filter RecentFilter) {
	C.gtk_recent_chooser_add_filter(chooser.native(), filter.native())
}

// GetCurrentItem is a wrapper around gtk_recent_chooser_get_current_item().
func (chooser RecentChooser) GetCurrentItem() RecentInfo {
	ret0 := C.gtk_recent_chooser_get_current_item(chooser.native())
	return wrapRecentInfo(ret0)
}

// GetCurrentUri is a wrapper around gtk_recent_chooser_get_current_uri().
func (chooser RecentChooser) GetCurrentUri() string {
	ret0 := C.gtk_recent_chooser_get_current_uri(chooser.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetFilter is a wrapper around gtk_recent_chooser_get_filter().
func (chooser RecentChooser) GetFilter() RecentFilter {
	ret0 := C.gtk_recent_chooser_get_filter(chooser.native())
	return wrapRecentFilter(ret0)
}

// GetItems is a wrapper around gtk_recent_chooser_get_items().
func (chooser RecentChooser) GetItems() glib.List {
	ret0 := C.gtk_recent_chooser_get_items(chooser.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapRecentInfo(p) }) /*gir:GLib*/
}

// GetLimit is a wrapper around gtk_recent_chooser_get_limit().
func (chooser RecentChooser) GetLimit() int {
	ret0 := C.gtk_recent_chooser_get_limit(chooser.native())
	return int(ret0)
}

// GetLocalOnly is a wrapper around gtk_recent_chooser_get_local_only().
func (chooser RecentChooser) GetLocalOnly() bool {
	ret0 := C.gtk_recent_chooser_get_local_only(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSelectMultiple is a wrapper around gtk_recent_chooser_get_select_multiple().
func (chooser RecentChooser) GetSelectMultiple() bool {
	ret0 := C.gtk_recent_chooser_get_select_multiple(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowIcons is a wrapper around gtk_recent_chooser_get_show_icons().
func (chooser RecentChooser) GetShowIcons() bool {
	ret0 := C.gtk_recent_chooser_get_show_icons(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowNotFound is a wrapper around gtk_recent_chooser_get_show_not_found().
func (chooser RecentChooser) GetShowNotFound() bool {
	ret0 := C.gtk_recent_chooser_get_show_not_found(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowPrivate is a wrapper around gtk_recent_chooser_get_show_private().
func (chooser RecentChooser) GetShowPrivate() bool {
	ret0 := C.gtk_recent_chooser_get_show_private(chooser.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSortType is a wrapper around gtk_recent_chooser_get_sort_type().
func (chooser RecentChooser) GetSortType() RecentSortType {
	ret0 := C.gtk_recent_chooser_get_sort_type(chooser.native())
	return RecentSortType(ret0)
}

// GetUris is a wrapper around gtk_recent_chooser_get_uris().
func (chooser RecentChooser) GetUris() []string {
	var length0 C.gsize
	ret0 := C.gtk_recent_chooser_get_uris(chooser.native(), &length0)
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

// SelectAll is a wrapper around gtk_recent_chooser_select_all().
func (chooser RecentChooser) SelectAll() {
	C.gtk_recent_chooser_select_all(chooser.native())
}

// SelectUri is a wrapper around gtk_recent_chooser_select_uri().
func (chooser RecentChooser) SelectUri(uri string) (bool, error) {
	uri0 := (*C.gchar)(C.CString(uri))
	var err glib.Error
	ret0 := C.gtk_recent_chooser_select_uri(chooser.native(), uri0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetCurrentUri is a wrapper around gtk_recent_chooser_set_current_uri().
func (chooser RecentChooser) SetCurrentUri(uri string) (bool, error) {
	uri0 := (*C.gchar)(C.CString(uri))
	var err glib.Error
	ret0 := C.gtk_recent_chooser_set_current_uri(chooser.native(), uri0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetFilter is a wrapper around gtk_recent_chooser_set_filter().
func (chooser RecentChooser) SetFilter(filter RecentFilter) {
	C.gtk_recent_chooser_set_filter(chooser.native(), filter.native())
}

// SetLimit is a wrapper around gtk_recent_chooser_set_limit().
func (chooser RecentChooser) SetLimit(limit int) {
	C.gtk_recent_chooser_set_limit(chooser.native(), C.gint(limit))
}

// SetLocalOnly is a wrapper around gtk_recent_chooser_set_local_only().
func (chooser RecentChooser) SetLocalOnly(local_only bool) {
	C.gtk_recent_chooser_set_local_only(chooser.native(), C.gboolean(util.Bool2Int(local_only)) /*go:.util*/)
}

// SetSelectMultiple is a wrapper around gtk_recent_chooser_set_select_multiple().
func (chooser RecentChooser) SetSelectMultiple(select_multiple bool) {
	C.gtk_recent_chooser_set_select_multiple(chooser.native(), C.gboolean(util.Bool2Int(select_multiple)) /*go:.util*/)
}

// SetShowIcons is a wrapper around gtk_recent_chooser_set_show_icons().
func (chooser RecentChooser) SetShowIcons(show_icons bool) {
	C.gtk_recent_chooser_set_show_icons(chooser.native(), C.gboolean(util.Bool2Int(show_icons)) /*go:.util*/)
}

// SetShowNotFound is a wrapper around gtk_recent_chooser_set_show_not_found().
func (chooser RecentChooser) SetShowNotFound(show_not_found bool) {
	C.gtk_recent_chooser_set_show_not_found(chooser.native(), C.gboolean(util.Bool2Int(show_not_found)) /*go:.util*/)
}

// SetShowPrivate is a wrapper around gtk_recent_chooser_set_show_private().
func (chooser RecentChooser) SetShowPrivate(show_private bool) {
	C.gtk_recent_chooser_set_show_private(chooser.native(), C.gboolean(util.Bool2Int(show_private)) /*go:.util*/)
}

// SetShowTips is a wrapper around gtk_recent_chooser_set_show_tips().
func (chooser RecentChooser) SetShowTips(show_tips bool) {
	C.gtk_recent_chooser_set_show_tips(chooser.native(), C.gboolean(util.Bool2Int(show_tips)) /*go:.util*/)
}

// UnselectAll is a wrapper around gtk_recent_chooser_unselect_all().
func (chooser RecentChooser) UnselectAll() {
	C.gtk_recent_chooser_unselect_all(chooser.native())
}

// UnselectUri is a wrapper around gtk_recent_chooser_unselect_uri().
func (chooser RecentChooser) UnselectUri(uri string) {
	uri0 := (*C.gchar)(C.CString(uri))
	C.gtk_recent_chooser_unselect_uri(chooser.native(), uri0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
}

// Object RecentFilter
type RecentFilter struct {
	gobject.InitiallyUnowned
}

func (v RecentFilter) native() *C.GtkRecentFilter {
	return (*C.GtkRecentFilter)(v.Ptr)
}
func wrapRecentFilter(p *C.GtkRecentFilter) (v RecentFilter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRecentFilter(p unsafe.Pointer) (v RecentFilter) {
	v.Ptr = p
	return
}
func (v RecentFilter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentFilter(p unsafe.Pointer) interface{} {
	return WrapRecentFilter(p)
}
func (v RecentFilter) GetType() gobject.Type {
	return gobject.Type(C.gtk_recent_filter_get_type())
}
func (v RecentFilter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRecentFilter(unsafe.Pointer(ptr)), nil
	}
}
func (v RecentFilter) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// RecentFilterNew is a wrapper around gtk_recent_filter_new().
func RecentFilterNew() RecentFilter {
	ret0 := C.gtk_recent_filter_new()
	return wrapRecentFilter(ret0)
}

// AddAge is a wrapper around gtk_recent_filter_add_age().
func (filter RecentFilter) AddAge(days int) {
	C.gtk_recent_filter_add_age(filter.native(), C.gint(days))
}

// AddApplication is a wrapper around gtk_recent_filter_add_application().
func (filter RecentFilter) AddApplication(application string) {
	application0 := (*C.gchar)(C.CString(application))
	C.gtk_recent_filter_add_application(filter.native(), application0)
	C.free(unsafe.Pointer(application0)) /*ch:<stdlib.h>*/
}

// AddGroup is a wrapper around gtk_recent_filter_add_group().
func (filter RecentFilter) AddGroup(group string) {
	group0 := (*C.gchar)(C.CString(group))
	C.gtk_recent_filter_add_group(filter.native(), group0)
	C.free(unsafe.Pointer(group0)) /*ch:<stdlib.h>*/
}

// AddMimeType is a wrapper around gtk_recent_filter_add_mime_type().
func (filter RecentFilter) AddMimeType(mime_type string) {
	mime_type0 := (*C.gchar)(C.CString(mime_type))
	C.gtk_recent_filter_add_mime_type(filter.native(), mime_type0)
	C.free(unsafe.Pointer(mime_type0)) /*ch:<stdlib.h>*/
}

// AddPattern is a wrapper around gtk_recent_filter_add_pattern().
func (filter RecentFilter) AddPattern(pattern string) {
	pattern0 := (*C.gchar)(C.CString(pattern))
	C.gtk_recent_filter_add_pattern(filter.native(), pattern0)
	C.free(unsafe.Pointer(pattern0)) /*ch:<stdlib.h>*/
}

// AddPixbufFormats is a wrapper around gtk_recent_filter_add_pixbuf_formats().
func (filter RecentFilter) AddPixbufFormats() {
	C.gtk_recent_filter_add_pixbuf_formats(filter.native())
}

// Filter is a wrapper around gtk_recent_filter_filter().
func (filter RecentFilter) Filter(filter_info RecentFilterInfo) bool {
	ret0 := C.gtk_recent_filter_filter(filter.native(), filter_info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetName is a wrapper around gtk_recent_filter_get_name().
func (filter RecentFilter) GetName() string {
	ret0 := C.gtk_recent_filter_get_name(filter.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetName is a wrapper around gtk_recent_filter_set_name().
func (filter RecentFilter) SetName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_recent_filter_set_name(filter.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// Object SizeGroup
type SizeGroup struct {
	gobject.Object
}

func (v SizeGroup) native() *C.GtkSizeGroup {
	return (*C.GtkSizeGroup)(v.Ptr)
}
func wrapSizeGroup(p *C.GtkSizeGroup) (v SizeGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSizeGroup(p unsafe.Pointer) (v SizeGroup) {
	v.Ptr = p
	return
}
func (v SizeGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSizeGroup(p unsafe.Pointer) interface{} {
	return WrapSizeGroup(p)
}
func (v SizeGroup) GetType() gobject.Type {
	return gobject.Type(C.gtk_size_group_get_type())
}
func (v SizeGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSizeGroup(unsafe.Pointer(ptr)), nil
	}
}
func (v SizeGroup) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// SizeGroupNew is a wrapper around gtk_size_group_new().
func SizeGroupNew(mode SizeGroupMode) SizeGroup {
	ret0 := C.gtk_size_group_new(C.GtkSizeGroupMode(mode))
	return wrapSizeGroup(ret0)
}

// AddWidget is a wrapper around gtk_size_group_add_widget().
func (size_group SizeGroup) AddWidget(widget Widget) {
	C.gtk_size_group_add_widget(size_group.native(), widget.native())
}

// GetMode is a wrapper around gtk_size_group_get_mode().
func (size_group SizeGroup) GetMode() SizeGroupMode {
	ret0 := C.gtk_size_group_get_mode(size_group.native())
	return SizeGroupMode(ret0)
}

// RemoveWidget is a wrapper around gtk_size_group_remove_widget().
func (size_group SizeGroup) RemoveWidget(widget Widget) {
	C.gtk_size_group_remove_widget(size_group.native(), widget.native())
}

// SetMode is a wrapper around gtk_size_group_set_mode().
func (size_group SizeGroup) SetMode(mode SizeGroupMode) {
	C.gtk_size_group_set_mode(size_group.native(), C.GtkSizeGroupMode(mode))
}

// Object AccelLabel
type AccelLabel struct {
	Label
}

func (v AccelLabel) native() *C.GtkAccelLabel {
	return (*C.GtkAccelLabel)(v.Ptr)
}
func wrapAccelLabel(p *C.GtkAccelLabel) (v AccelLabel) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAccelLabel(p unsafe.Pointer) (v AccelLabel) {
	v.Ptr = p
	return
}
func (v AccelLabel) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAccelLabel(p unsafe.Pointer) interface{} {
	return WrapAccelLabel(p)
}
func (v AccelLabel) GetType() gobject.Type {
	return gobject.Type(C.gtk_accel_label_get_type())
}
func (v AccelLabel) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAccelLabel(unsafe.Pointer(ptr)), nil
	}
}
func (v AccelLabel) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v AccelLabel) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// AccelLabelNew is a wrapper around gtk_accel_label_new().
func AccelLabelNew(string string) Widget {
	string0 := (*C.gchar)(C.CString(string))
	ret0 := C.gtk_accel_label_new(string0)
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetAccel is a wrapper around gtk_accel_label_get_accel().
func (accel_label AccelLabel) GetAccel() (uint /*gir:Gdk*/, gdk.ModifierType) {
	var accelerator_key0 C.guint
	var accelerator_mods0 C.GdkModifierType
	C.gtk_accel_label_get_accel(accel_label.native(), &accelerator_key0, &accelerator_mods0)
	return uint(accelerator_key0) /*gir:Gdk*/, gdk.ModifierType(accelerator_mods0)
}

// GetAccelWidget is a wrapper around gtk_accel_label_get_accel_widget().
func (accel_label AccelLabel) GetAccelWidget() Widget {
	ret0 := C.gtk_accel_label_get_accel_widget(accel_label.native())
	return wrapWidget(ret0)
}

// GetAccelWidth is a wrapper around gtk_accel_label_get_accel_width().
func (accel_label AccelLabel) GetAccelWidth() uint {
	ret0 := C.gtk_accel_label_get_accel_width(accel_label.native())
	return uint(ret0)
}

// Refetch is a wrapper around gtk_accel_label_refetch().
func (accel_label AccelLabel) Refetch() bool {
	ret0 := C.gtk_accel_label_refetch(accel_label.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetAccel is a wrapper around gtk_accel_label_set_accel().
func (accel_label AccelLabel) SetAccel(accelerator_key uint, accelerator_mods /*gir:Gdk*/ gdk.ModifierType) {
	C.gtk_accel_label_set_accel(accel_label.native(), C.guint(accelerator_key), C.GdkModifierType(accelerator_mods))
}

// SetAccelClosure is a wrapper around gtk_accel_label_set_accel_closure().
func (accel_label AccelLabel) SetAccelClosure(accel_closure gobject.Closure) {
	C.gtk_accel_label_set_accel_closure(accel_label.native(), (*C.GClosure)(accel_closure.Ptr))
}

// SetAccelWidget is a wrapper around gtk_accel_label_set_accel_widget().
func (accel_label AccelLabel) SetAccelWidget(accel_widget Widget) {
	C.gtk_accel_label_set_accel_widget(accel_label.native(), accel_widget.native())
}

// Object Label
type Label struct {
	Misc
}

func (v Label) native() *C.GtkLabel {
	return (*C.GtkLabel)(v.Ptr)
}
func wrapLabel(p *C.GtkLabel) (v Label) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLabel(p unsafe.Pointer) (v Label) {
	v.Ptr = p
	return
}
func (v Label) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLabel(p unsafe.Pointer) interface{} {
	return WrapLabel(p)
}
func (v Label) GetType() gobject.Type {
	return gobject.Type(C.gtk_label_get_type())
}
func (v Label) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLabel(unsafe.Pointer(ptr)), nil
	}
}
func (v Label) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Label) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// LabelNew is a wrapper around gtk_label_new().
func LabelNew(str string) Widget {
	str0 := (*C.gchar)(C.CString(str))
	ret0 := C.gtk_label_new(str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// LabelNewWithMnemonic is a wrapper around gtk_label_new_with_mnemonic().
func LabelNewWithMnemonic(str string) Widget {
	str0 := (*C.gchar)(C.CString(str))
	ret0 := C.gtk_label_new_with_mnemonic(str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetAngle is a wrapper around gtk_label_get_angle().
func (label Label) GetAngle() float64 {
	ret0 := C.gtk_label_get_angle(label.native())
	return float64(ret0)
}

// GetAttributes is a wrapper around gtk_label_get_attributes().
func (label Label) GetAttributes() pango.AttrList {
	ret0 := C.gtk_label_get_attributes(label.native())
	return pango.WrapAttrList(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetCurrentUri is a wrapper around gtk_label_get_current_uri().
func (label Label) GetCurrentUri() string {
	ret0 := C.gtk_label_get_current_uri(label.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetEllipsize is a wrapper around gtk_label_get_ellipsize().
func (label Label) GetEllipsize() /*gir:Pango*/ pango.EllipsizeMode {
	ret0 := C.gtk_label_get_ellipsize(label.native())
	return /*gir:Pango*/ pango.EllipsizeMode(ret0)
}

// GetJustify is a wrapper around gtk_label_get_justify().
func (label Label) GetJustify() Justification {
	ret0 := C.gtk_label_get_justify(label.native())
	return Justification(ret0)
}

// GetLabel is a wrapper around gtk_label_get_label().
func (label Label) GetLabel() string {
	ret0 := C.gtk_label_get_label(label.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLayout is a wrapper around gtk_label_get_layout().
func (label Label) GetLayout() pango.Layout {
	ret0 := C.gtk_label_get_layout(label.native())
	return pango.WrapLayout(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetLayoutOffsets is a wrapper around gtk_label_get_layout_offsets().
func (label Label) GetLayoutOffsets() (int, int) {
	var x0 C.gint
	var y0 C.gint
	C.gtk_label_get_layout_offsets(label.native(), &x0, &y0)
	return int(x0), int(y0)
}

// GetLineWrap is a wrapper around gtk_label_get_line_wrap().
func (label Label) GetLineWrap() bool {
	ret0 := C.gtk_label_get_line_wrap(label.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLineWrapMode is a wrapper around gtk_label_get_line_wrap_mode().
func (label Label) GetLineWrapMode() /*gir:Pango*/ pango.WrapMode {
	ret0 := C.gtk_label_get_line_wrap_mode(label.native())
	return /*gir:Pango*/ pango.WrapMode(ret0)
}

// GetLines is a wrapper around gtk_label_get_lines().
func (label Label) GetLines() int {
	ret0 := C.gtk_label_get_lines(label.native())
	return int(ret0)
}

// GetMaxWidthChars is a wrapper around gtk_label_get_max_width_chars().
func (label Label) GetMaxWidthChars() int {
	ret0 := C.gtk_label_get_max_width_chars(label.native())
	return int(ret0)
}

// GetMnemonicKeyval is a wrapper around gtk_label_get_mnemonic_keyval().
func (label Label) GetMnemonicKeyval() uint {
	ret0 := C.gtk_label_get_mnemonic_keyval(label.native())
	return uint(ret0)
}

// GetMnemonicWidget is a wrapper around gtk_label_get_mnemonic_widget().
func (label Label) GetMnemonicWidget() Widget {
	ret0 := C.gtk_label_get_mnemonic_widget(label.native())
	return wrapWidget(ret0)
}

// GetSelectable is a wrapper around gtk_label_get_selectable().
func (label Label) GetSelectable() bool {
	ret0 := C.gtk_label_get_selectable(label.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSelectionBounds is a wrapper around gtk_label_get_selection_bounds().
func (label Label) GetSelectionBounds() (bool, int, int) {
	var start0 C.gint
	var end0 C.gint
	ret0 := C.gtk_label_get_selection_bounds(label.native(), &start0, &end0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(start0), int(end0)
}

// GetSingleLineMode is a wrapper around gtk_label_get_single_line_mode().
func (label Label) GetSingleLineMode() bool {
	ret0 := C.gtk_label_get_single_line_mode(label.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetText is a wrapper around gtk_label_get_text().
func (label Label) GetText() string {
	ret0 := C.gtk_label_get_text(label.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTrackVisitedLinks is a wrapper around gtk_label_get_track_visited_links().
func (label Label) GetTrackVisitedLinks() bool {
	ret0 := C.gtk_label_get_track_visited_links(label.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUseMarkup is a wrapper around gtk_label_get_use_markup().
func (label Label) GetUseMarkup() bool {
	ret0 := C.gtk_label_get_use_markup(label.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUseUnderline is a wrapper around gtk_label_get_use_underline().
func (label Label) GetUseUnderline() bool {
	ret0 := C.gtk_label_get_use_underline(label.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetWidthChars is a wrapper around gtk_label_get_width_chars().
func (label Label) GetWidthChars() int {
	ret0 := C.gtk_label_get_width_chars(label.native())
	return int(ret0)
}

// GetXalign is a wrapper around gtk_label_get_xalign().
func (label Label) GetXalign() float32 {
	ret0 := C.gtk_label_get_xalign(label.native())
	return float32(ret0)
}

// GetYalign is a wrapper around gtk_label_get_yalign().
func (label Label) GetYalign() float32 {
	ret0 := C.gtk_label_get_yalign(label.native())
	return float32(ret0)
}

// SelectRegion is a wrapper around gtk_label_select_region().
func (label Label) SelectRegion(start_offset int, end_offset int) {
	C.gtk_label_select_region(label.native(), C.gint(start_offset), C.gint(end_offset))
}

// SetAngle is a wrapper around gtk_label_set_angle().
func (label Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(label.native(), C.gdouble(angle))
}

// SetAttributes is a wrapper around gtk_label_set_attributes().
func (label Label) SetAttributes(attrs pango.AttrList) {
	C.gtk_label_set_attributes(label.native(), (*C.PangoAttrList)(attrs.Ptr))
}

// SetEllipsize is a wrapper around gtk_label_set_ellipsize().
func (label Label) SetEllipsize(mode /*gir:Pango*/ pango.EllipsizeMode) {
	C.gtk_label_set_ellipsize(label.native(), C.PangoEllipsizeMode(mode))
}

// SetJustify is a wrapper around gtk_label_set_justify().
func (label Label) SetJustify(jtype Justification) {
	C.gtk_label_set_justify(label.native(), C.GtkJustification(jtype))
}

// SetLabel is a wrapper around gtk_label_set_label().
func (label Label) SetLabel(str string) {
	str0 := (*C.gchar)(C.CString(str))
	C.gtk_label_set_label(label.native(), str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
}

// SetLineWrap is a wrapper around gtk_label_set_line_wrap().
func (label Label) SetLineWrap(wrap bool) {
	C.gtk_label_set_line_wrap(label.native(), C.gboolean(util.Bool2Int(wrap)) /*go:.util*/)
}

// SetLineWrapMode is a wrapper around gtk_label_set_line_wrap_mode().
func (label Label) SetLineWrapMode(wrap_mode /*gir:Pango*/ pango.WrapMode) {
	C.gtk_label_set_line_wrap_mode(label.native(), C.PangoWrapMode(wrap_mode))
}

// SetLines is a wrapper around gtk_label_set_lines().
func (label Label) SetLines(lines int) {
	C.gtk_label_set_lines(label.native(), C.gint(lines))
}

// SetMarkup is a wrapper around gtk_label_set_markup().
func (label Label) SetMarkup(str string) {
	str0 := (*C.gchar)(C.CString(str))
	C.gtk_label_set_markup(label.native(), str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
}

// SetMarkupWithMnemonic is a wrapper around gtk_label_set_markup_with_mnemonic().
func (label Label) SetMarkupWithMnemonic(str string) {
	str0 := (*C.gchar)(C.CString(str))
	C.gtk_label_set_markup_with_mnemonic(label.native(), str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
}

// SetMaxWidthChars is a wrapper around gtk_label_set_max_width_chars().
func (label Label) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(label.native(), C.gint(n_chars))
}

// SetMnemonicWidget is a wrapper around gtk_label_set_mnemonic_widget().
func (label Label) SetMnemonicWidget(widget Widget) {
	C.gtk_label_set_mnemonic_widget(label.native(), widget.native())
}

// SetPattern is a wrapper around gtk_label_set_pattern().
func (label Label) SetPattern(pattern string) {
	pattern0 := (*C.gchar)(C.CString(pattern))
	C.gtk_label_set_pattern(label.native(), pattern0)
	C.free(unsafe.Pointer(pattern0)) /*ch:<stdlib.h>*/
}

// SetSelectable is a wrapper around gtk_label_set_selectable().
func (label Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(label.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetSingleLineMode is a wrapper around gtk_label_set_single_line_mode().
func (label Label) SetSingleLineMode(single_line_mode bool) {
	C.gtk_label_set_single_line_mode(label.native(), C.gboolean(util.Bool2Int(single_line_mode)) /*go:.util*/)
}

// SetText is a wrapper around gtk_label_set_text().
func (label Label) SetText(str string) {
	str0 := (*C.gchar)(C.CString(str))
	C.gtk_label_set_text(label.native(), str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
}

// SetTextWithMnemonic is a wrapper around gtk_label_set_text_with_mnemonic().
func (label Label) SetTextWithMnemonic(str string) {
	str0 := (*C.gchar)(C.CString(str))
	C.gtk_label_set_text_with_mnemonic(label.native(), str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
}

// SetTrackVisitedLinks is a wrapper around gtk_label_set_track_visited_links().
func (label Label) SetTrackVisitedLinks(track_links bool) {
	C.gtk_label_set_track_visited_links(label.native(), C.gboolean(util.Bool2Int(track_links)) /*go:.util*/)
}

// SetUseMarkup is a wrapper around gtk_label_set_use_markup().
func (label Label) SetUseMarkup(setting bool) {
	C.gtk_label_set_use_markup(label.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetUseUnderline is a wrapper around gtk_label_set_use_underline().
func (label Label) SetUseUnderline(setting bool) {
	C.gtk_label_set_use_underline(label.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetWidthChars is a wrapper around gtk_label_set_width_chars().
func (label Label) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(label.native(), C.gint(n_chars))
}

// SetXalign is a wrapper around gtk_label_set_xalign().
func (label Label) SetXalign(xalign float32) {
	C.gtk_label_set_xalign(label.native(), C.gfloat(xalign))
}

// SetYalign is a wrapper around gtk_label_set_yalign().
func (label Label) SetYalign(yalign float32) {
	C.gtk_label_set_yalign(label.native(), C.gfloat(yalign))
}

// Object Misc
type Misc struct {
	Widget
}

func (v Misc) native() *C.GtkMisc {
	return (*C.GtkMisc)(v.Ptr)
}
func wrapMisc(p *C.GtkMisc) (v Misc) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMisc(p unsafe.Pointer) (v Misc) {
	v.Ptr = p
	return
}
func (v Misc) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMisc(p unsafe.Pointer) interface{} {
	return WrapMisc(p)
}
func (v Misc) GetType() gobject.Type {
	return gobject.Type(C.gtk_misc_get_type())
}
func (v Misc) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMisc(unsafe.Pointer(ptr)), nil
	}
}
func (v Misc) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Misc) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object AppChooserButton
type AppChooserButton struct {
	ComboBox
}

func (v AppChooserButton) native() *C.GtkAppChooserButton {
	return (*C.GtkAppChooserButton)(v.Ptr)
}
func wrapAppChooserButton(p *C.GtkAppChooserButton) (v AppChooserButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAppChooserButton(p unsafe.Pointer) (v AppChooserButton) {
	v.Ptr = p
	return
}
func (v AppChooserButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppChooserButton(p unsafe.Pointer) interface{} {
	return WrapAppChooserButton(p)
}
func (v AppChooserButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_app_chooser_button_get_type())
}
func (v AppChooserButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppChooserButton(unsafe.Pointer(ptr)), nil
	}
}
func (v AppChooserButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v AppChooserButton) AppChooser() AppChooser {
	return WrapAppChooser(v.Ptr)
}
func (v AppChooserButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v AppChooserButton) CellEditable() CellEditable {
	return WrapCellEditable(v.Ptr)
}
func (v AppChooserButton) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}

// AppendCustomItem is a wrapper around gtk_app_chooser_button_append_custom_item().
func (self AppChooserButton) AppendCustomItem(name string, label string, icon gio.Icon) {
	name0 := (*C.gchar)(C.CString(name))
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_app_chooser_button_append_custom_item(self.native(), name0, label0, (*C.GIcon)(icon.Ptr))
	C.free(unsafe.Pointer(name0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// AppendSeparator is a wrapper around gtk_app_chooser_button_append_separator().
func (self AppChooserButton) AppendSeparator() {
	C.gtk_app_chooser_button_append_separator(self.native())
}

// GetHeading is a wrapper around gtk_app_chooser_button_get_heading().
func (self AppChooserButton) GetHeading() string {
	ret0 := C.gtk_app_chooser_button_get_heading(self.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetShowDefaultItem is a wrapper around gtk_app_chooser_button_get_show_default_item().
func (self AppChooserButton) GetShowDefaultItem() bool {
	ret0 := C.gtk_app_chooser_button_get_show_default_item(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowDialogItem is a wrapper around gtk_app_chooser_button_get_show_dialog_item().
func (self AppChooserButton) GetShowDialogItem() bool {
	ret0 := C.gtk_app_chooser_button_get_show_dialog_item(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetActiveCustomItem is a wrapper around gtk_app_chooser_button_set_active_custom_item().
func (self AppChooserButton) SetActiveCustomItem(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_app_chooser_button_set_active_custom_item(self.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetHeading is a wrapper around gtk_app_chooser_button_set_heading().
func (self AppChooserButton) SetHeading(heading string) {
	heading0 := (*C.gchar)(C.CString(heading))
	C.gtk_app_chooser_button_set_heading(self.native(), heading0)
	C.free(unsafe.Pointer(heading0)) /*ch:<stdlib.h>*/
}

// SetShowDefaultItem is a wrapper around gtk_app_chooser_button_set_show_default_item().
func (self AppChooserButton) SetShowDefaultItem(setting bool) {
	C.gtk_app_chooser_button_set_show_default_item(self.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetShowDialogItem is a wrapper around gtk_app_chooser_button_set_show_dialog_item().
func (self AppChooserButton) SetShowDialogItem(setting bool) {
	C.gtk_app_chooser_button_set_show_dialog_item(self.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// Object ComboBox
type ComboBox struct {
	Bin
}

func (v ComboBox) native() *C.GtkComboBox {
	return (*C.GtkComboBox)(v.Ptr)
}
func wrapComboBox(p *C.GtkComboBox) (v ComboBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapComboBox(p unsafe.Pointer) (v ComboBox) {
	v.Ptr = p
	return
}
func (v ComboBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapComboBox(p unsafe.Pointer) interface{} {
	return WrapComboBox(p)
}
func (v ComboBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_combo_box_get_type())
}
func (v ComboBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapComboBox(unsafe.Pointer(ptr)), nil
	}
}
func (v ComboBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ComboBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ComboBox) CellEditable() CellEditable {
	return WrapCellEditable(v.Ptr)
}
func (v ComboBox) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}

// ComboBoxNew is a wrapper around gtk_combo_box_new().
func ComboBoxNew() Widget {
	ret0 := C.gtk_combo_box_new()
	return wrapWidget(ret0)
}

// ComboBoxNewWithArea is a wrapper around gtk_combo_box_new_with_area().
func ComboBoxNewWithArea(area CellArea) Widget {
	ret0 := C.gtk_combo_box_new_with_area(area.native())
	return wrapWidget(ret0)
}

// ComboBoxNewWithAreaAndEntry is a wrapper around gtk_combo_box_new_with_area_and_entry().
func ComboBoxNewWithAreaAndEntry(area CellArea) Widget {
	ret0 := C.gtk_combo_box_new_with_area_and_entry(area.native())
	return wrapWidget(ret0)
}

// ComboBoxNewWithEntry is a wrapper around gtk_combo_box_new_with_entry().
func ComboBoxNewWithEntry() Widget {
	ret0 := C.gtk_combo_box_new_with_entry()
	return wrapWidget(ret0)
}

// ComboBoxNewWithModel is a wrapper around gtk_combo_box_new_with_model().
func ComboBoxNewWithModel(model TreeModel) Widget {
	ret0 := C.gtk_combo_box_new_with_model(model.native())
	return wrapWidget(ret0)
}

// ComboBoxNewWithModelAndEntry is a wrapper around gtk_combo_box_new_with_model_and_entry().
func ComboBoxNewWithModelAndEntry(model TreeModel) Widget {
	ret0 := C.gtk_combo_box_new_with_model_and_entry(model.native())
	return wrapWidget(ret0)
}

// GetActive is a wrapper around gtk_combo_box_get_active().
func (combo_box ComboBox) GetActive() int {
	ret0 := C.gtk_combo_box_get_active(combo_box.native())
	return int(ret0)
}

// GetActiveId is a wrapper around gtk_combo_box_get_active_id().
func (combo_box ComboBox) GetActiveId() string {
	ret0 := C.gtk_combo_box_get_active_id(combo_box.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetButtonSensitivity is a wrapper around gtk_combo_box_get_button_sensitivity().
func (combo_box ComboBox) GetButtonSensitivity() SensitivityType {
	ret0 := C.gtk_combo_box_get_button_sensitivity(combo_box.native())
	return SensitivityType(ret0)
}

// GetColumnSpanColumn is a wrapper around gtk_combo_box_get_column_span_column().
func (combo_box ComboBox) GetColumnSpanColumn() int {
	ret0 := C.gtk_combo_box_get_column_span_column(combo_box.native())
	return int(ret0)
}

// GetEntryTextColumn is a wrapper around gtk_combo_box_get_entry_text_column().
func (combo_box ComboBox) GetEntryTextColumn() int {
	ret0 := C.gtk_combo_box_get_entry_text_column(combo_box.native())
	return int(ret0)
}

// GetHasEntry is a wrapper around gtk_combo_box_get_has_entry().
func (combo_box ComboBox) GetHasEntry() bool {
	ret0 := C.gtk_combo_box_get_has_entry(combo_box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIdColumn is a wrapper around gtk_combo_box_get_id_column().
func (combo_box ComboBox) GetIdColumn() int {
	ret0 := C.gtk_combo_box_get_id_column(combo_box.native())
	return int(ret0)
}

// GetModel is a wrapper around gtk_combo_box_get_model().
func (combo_box ComboBox) GetModel() TreeModel {
	ret0 := C.gtk_combo_box_get_model(combo_box.native())
	return wrapTreeModel(ret0)
}

// GetPopupAccessible is a wrapper around gtk_combo_box_get_popup_accessible().
func (combo_box ComboBox) GetPopupAccessible() atk.Object {
	ret0 := C.gtk_combo_box_get_popup_accessible(combo_box.native())
	return atk.WrapObject(unsafe.Pointer(ret0)) /*gir:Atk*/
}

// GetPopupFixedWidth is a wrapper around gtk_combo_box_get_popup_fixed_width().
func (combo_box ComboBox) GetPopupFixedWidth() bool {
	ret0 := C.gtk_combo_box_get_popup_fixed_width(combo_box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRowSpanColumn is a wrapper around gtk_combo_box_get_row_span_column().
func (combo_box ComboBox) GetRowSpanColumn() int {
	ret0 := C.gtk_combo_box_get_row_span_column(combo_box.native())
	return int(ret0)
}

// GetWrapWidth is a wrapper around gtk_combo_box_get_wrap_width().
func (combo_box ComboBox) GetWrapWidth() int {
	ret0 := C.gtk_combo_box_get_wrap_width(combo_box.native())
	return int(ret0)
}

// Popdown is a wrapper around gtk_combo_box_popdown().
func (combo_box ComboBox) Popdown() {
	C.gtk_combo_box_popdown(combo_box.native())
}

// Popup is a wrapper around gtk_combo_box_popup().
func (combo_box ComboBox) Popup() {
	C.gtk_combo_box_popup(combo_box.native())
}

// PopupForDevice is a wrapper around gtk_combo_box_popup_for_device().
func (combo_box ComboBox) PopupForDevice(device gdk.Device) {
	C.gtk_combo_box_popup_for_device(combo_box.native(), (*C.GdkDevice)(device.Ptr))
}

// SetActive is a wrapper around gtk_combo_box_set_active().
func (combo_box ComboBox) SetActive(index_ int) {
	C.gtk_combo_box_set_active(combo_box.native(), C.gint(index_))
}

// SetActiveId is a wrapper around gtk_combo_box_set_active_id().
func (combo_box ComboBox) SetActiveId(active_id string) bool {
	active_id0 := (*C.gchar)(C.CString(active_id))
	ret0 := C.gtk_combo_box_set_active_id(combo_box.native(), active_id0)
	C.free(unsafe.Pointer(active_id0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// SetActiveIter is a wrapper around gtk_combo_box_set_active_iter().
func (combo_box ComboBox) SetActiveIter(iter TreeIter) {
	C.gtk_combo_box_set_active_iter(combo_box.native(), iter.native())
}

// SetButtonSensitivity is a wrapper around gtk_combo_box_set_button_sensitivity().
func (combo_box ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	C.gtk_combo_box_set_button_sensitivity(combo_box.native(), C.GtkSensitivityType(sensitivity))
}

// SetColumnSpanColumn is a wrapper around gtk_combo_box_set_column_span_column().
func (combo_box ComboBox) SetColumnSpanColumn(column_span int) {
	C.gtk_combo_box_set_column_span_column(combo_box.native(), C.gint(column_span))
}

// SetEntryTextColumn is a wrapper around gtk_combo_box_set_entry_text_column().
func (combo_box ComboBox) SetEntryTextColumn(text_column int) {
	C.gtk_combo_box_set_entry_text_column(combo_box.native(), C.gint(text_column))
}

// SetIdColumn is a wrapper around gtk_combo_box_set_id_column().
func (combo_box ComboBox) SetIdColumn(id_column int) {
	C.gtk_combo_box_set_id_column(combo_box.native(), C.gint(id_column))
}

// SetModel is a wrapper around gtk_combo_box_set_model().
func (combo_box ComboBox) SetModel(model TreeModel) {
	C.gtk_combo_box_set_model(combo_box.native(), model.native())
}

// SetPopupFixedWidth is a wrapper around gtk_combo_box_set_popup_fixed_width().
func (combo_box ComboBox) SetPopupFixedWidth(fixed bool) {
	C.gtk_combo_box_set_popup_fixed_width(combo_box.native(), C.gboolean(util.Bool2Int(fixed)) /*go:.util*/)
}

// SetRowSpanColumn is a wrapper around gtk_combo_box_set_row_span_column().
func (combo_box ComboBox) SetRowSpanColumn(row_span int) {
	C.gtk_combo_box_set_row_span_column(combo_box.native(), C.gint(row_span))
}

// SetWrapWidth is a wrapper around gtk_combo_box_set_wrap_width().
func (combo_box ComboBox) SetWrapWidth(width int) {
	C.gtk_combo_box_set_wrap_width(combo_box.native(), C.gint(width))
}

// Object AppChooserWidget
type AppChooserWidget struct {
	Box
}

func (v AppChooserWidget) native() *C.GtkAppChooserWidget {
	return (*C.GtkAppChooserWidget)(v.Ptr)
}
func wrapAppChooserWidget(p *C.GtkAppChooserWidget) (v AppChooserWidget) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAppChooserWidget(p unsafe.Pointer) (v AppChooserWidget) {
	v.Ptr = p
	return
}
func (v AppChooserWidget) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppChooserWidget(p unsafe.Pointer) interface{} {
	return WrapAppChooserWidget(p)
}
func (v AppChooserWidget) GetType() gobject.Type {
	return gobject.Type(C.gtk_app_chooser_widget_get_type())
}
func (v AppChooserWidget) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppChooserWidget(unsafe.Pointer(ptr)), nil
	}
}
func (v AppChooserWidget) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v AppChooserWidget) AppChooser() AppChooser {
	return WrapAppChooser(v.Ptr)
}
func (v AppChooserWidget) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v AppChooserWidget) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// AppChooserWidgetNew is a wrapper around gtk_app_chooser_widget_new().
func AppChooserWidgetNew(content_type string) Widget {
	content_type0 := (*C.gchar)(C.CString(content_type))
	ret0 := C.gtk_app_chooser_widget_new(content_type0)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetDefaultText is a wrapper around gtk_app_chooser_widget_get_default_text().
func (self AppChooserWidget) GetDefaultText() string {
	ret0 := C.gtk_app_chooser_widget_get_default_text(self.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetShowAll is a wrapper around gtk_app_chooser_widget_get_show_all().
func (self AppChooserWidget) GetShowAll() bool {
	ret0 := C.gtk_app_chooser_widget_get_show_all(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowDefault is a wrapper around gtk_app_chooser_widget_get_show_default().
func (self AppChooserWidget) GetShowDefault() bool {
	ret0 := C.gtk_app_chooser_widget_get_show_default(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowFallback is a wrapper around gtk_app_chooser_widget_get_show_fallback().
func (self AppChooserWidget) GetShowFallback() bool {
	ret0 := C.gtk_app_chooser_widget_get_show_fallback(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowOther is a wrapper around gtk_app_chooser_widget_get_show_other().
func (self AppChooserWidget) GetShowOther() bool {
	ret0 := C.gtk_app_chooser_widget_get_show_other(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowRecommended is a wrapper around gtk_app_chooser_widget_get_show_recommended().
func (self AppChooserWidget) GetShowRecommended() bool {
	ret0 := C.gtk_app_chooser_widget_get_show_recommended(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetDefaultText is a wrapper around gtk_app_chooser_widget_set_default_text().
func (self AppChooserWidget) SetDefaultText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_app_chooser_widget_set_default_text(self.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetShowAll is a wrapper around gtk_app_chooser_widget_set_show_all().
func (self AppChooserWidget) SetShowAll(setting bool) {
	C.gtk_app_chooser_widget_set_show_all(self.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetShowDefault is a wrapper around gtk_app_chooser_widget_set_show_default().
func (self AppChooserWidget) SetShowDefault(setting bool) {
	C.gtk_app_chooser_widget_set_show_default(self.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetShowFallback is a wrapper around gtk_app_chooser_widget_set_show_fallback().
func (self AppChooserWidget) SetShowFallback(setting bool) {
	C.gtk_app_chooser_widget_set_show_fallback(self.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetShowOther is a wrapper around gtk_app_chooser_widget_set_show_other().
func (self AppChooserWidget) SetShowOther(setting bool) {
	C.gtk_app_chooser_widget_set_show_other(self.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetShowRecommended is a wrapper around gtk_app_chooser_widget_set_show_recommended().
func (self AppChooserWidget) SetShowRecommended(setting bool) {
	C.gtk_app_chooser_widget_set_show_recommended(self.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// Object Box
type Box struct {
	Container
}

func (v Box) native() *C.GtkBox {
	return (*C.GtkBox)(v.Ptr)
}
func wrapBox(p *C.GtkBox) (v Box) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBox(p unsafe.Pointer) (v Box) {
	v.Ptr = p
	return
}
func (v Box) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBox(p unsafe.Pointer) interface{} {
	return WrapBox(p)
}
func (v Box) GetType() gobject.Type {
	return gobject.Type(C.gtk_box_get_type())
}
func (v Box) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBox(unsafe.Pointer(ptr)), nil
	}
}
func (v Box) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Box) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Box) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// BoxNew is a wrapper around gtk_box_new().
func BoxNew(orientation Orientation, spacing int) Widget {
	ret0 := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	return wrapWidget(ret0)
}

// GetBaselinePosition is a wrapper around gtk_box_get_baseline_position().
func (box Box) GetBaselinePosition() BaselinePosition {
	ret0 := C.gtk_box_get_baseline_position(box.native())
	return BaselinePosition(ret0)
}

// GetCenterWidget is a wrapper around gtk_box_get_center_widget().
func (box Box) GetCenterWidget() Widget {
	ret0 := C.gtk_box_get_center_widget(box.native())
	return wrapWidget(ret0)
}

// GetHomogeneous is a wrapper around gtk_box_get_homogeneous().
func (box Box) GetHomogeneous() bool {
	ret0 := C.gtk_box_get_homogeneous(box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSpacing is a wrapper around gtk_box_get_spacing().
func (box Box) GetSpacing() int {
	ret0 := C.gtk_box_get_spacing(box.native())
	return int(ret0)
}

// PackEnd is a wrapper around gtk_box_pack_end().
func (box Box) PackEnd(child Widget, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_end(box.native(), child.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/, C.gboolean(util.Bool2Int(fill)) /*go:.util*/, C.guint(padding))
}

// PackStart is a wrapper around gtk_box_pack_start().
func (box Box) PackStart(child Widget, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_start(box.native(), child.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/, C.gboolean(util.Bool2Int(fill)) /*go:.util*/, C.guint(padding))
}

// QueryChildPacking is a wrapper around gtk_box_query_child_packing().
func (box Box) QueryChildPacking(child Widget) (bool, bool, uint, PackType) {
	var expand0 C.gboolean
	var fill0 C.gboolean
	var padding0 C.guint
	var pack_type0 C.GtkPackType
	C.gtk_box_query_child_packing(box.native(), child.native(), &expand0, &fill0, &padding0, &pack_type0)
	return util.Int2Bool(int(expand0)) /*go:.util*/, util.Int2Bool(int(fill0)) /*go:.util*/, uint(padding0), PackType(pack_type0)
}

// ReorderChild is a wrapper around gtk_box_reorder_child().
func (box Box) ReorderChild(child Widget, position int) {
	C.gtk_box_reorder_child(box.native(), child.native(), C.gint(position))
}

// SetBaselinePosition is a wrapper around gtk_box_set_baseline_position().
func (box Box) SetBaselinePosition(position BaselinePosition) {
	C.gtk_box_set_baseline_position(box.native(), C.GtkBaselinePosition(position))
}

// SetChildPacking is a wrapper around gtk_box_set_child_packing().
func (box Box) SetChildPacking(child Widget, expand bool, fill bool, padding uint, pack_type PackType) {
	C.gtk_box_set_child_packing(box.native(), child.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/, C.gboolean(util.Bool2Int(fill)) /*go:.util*/, C.guint(padding), C.GtkPackType(pack_type))
}

// SetHomogeneous is a wrapper around gtk_box_set_homogeneous().
func (box Box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(box.native(), C.gboolean(util.Bool2Int(homogeneous)) /*go:.util*/)
}

// SetSpacing is a wrapper around gtk_box_set_spacing().
func (box Box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(box.native(), C.gint(spacing))
}

// Object ShortcutsWindow
type ShortcutsWindow struct {
	Window
}

func (v ShortcutsWindow) native() *C.GtkShortcutsWindow {
	return (*C.GtkShortcutsWindow)(v.Ptr)
}
func wrapShortcutsWindow(p *C.GtkShortcutsWindow) (v ShortcutsWindow) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapShortcutsWindow(p unsafe.Pointer) (v ShortcutsWindow) {
	v.Ptr = p
	return
}
func (v ShortcutsWindow) IsNil() bool {
	return v.Ptr == nil
}
func IWrapShortcutsWindow(p unsafe.Pointer) interface{} {
	return WrapShortcutsWindow(p)
}
func (v ShortcutsWindow) GetType() gobject.Type {
	return gobject.Type(C.gtk_shortcuts_window_get_type())
}
func (v ShortcutsWindow) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapShortcutsWindow(unsafe.Pointer(ptr)), nil
	}
}
func (v ShortcutsWindow) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ShortcutsWindow) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object Assistant
type Assistant struct {
	Window
}

func (v Assistant) native() *C.GtkAssistant {
	return (*C.GtkAssistant)(v.Ptr)
}
func wrapAssistant(p *C.GtkAssistant) (v Assistant) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAssistant(p unsafe.Pointer) (v Assistant) {
	v.Ptr = p
	return
}
func (v Assistant) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAssistant(p unsafe.Pointer) interface{} {
	return WrapAssistant(p)
}
func (v Assistant) GetType() gobject.Type {
	return gobject.Type(C.gtk_assistant_get_type())
}
func (v Assistant) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAssistant(unsafe.Pointer(ptr)), nil
	}
}
func (v Assistant) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Assistant) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// AssistantNew is a wrapper around gtk_assistant_new().
func AssistantNew() Widget {
	ret0 := C.gtk_assistant_new()
	return wrapWidget(ret0)
}

// AddActionWidget is a wrapper around gtk_assistant_add_action_widget().
func (assistant Assistant) AddActionWidget(child Widget) {
	C.gtk_assistant_add_action_widget(assistant.native(), child.native())
}

// AppendPage is a wrapper around gtk_assistant_append_page().
func (assistant Assistant) AppendPage(page Widget) int {
	ret0 := C.gtk_assistant_append_page(assistant.native(), page.native())
	return int(ret0)
}

// Commit is a wrapper around gtk_assistant_commit().
func (assistant Assistant) Commit() {
	C.gtk_assistant_commit(assistant.native())
}

// GetCurrentPage is a wrapper around gtk_assistant_get_current_page().
func (assistant Assistant) GetCurrentPage() int {
	ret0 := C.gtk_assistant_get_current_page(assistant.native())
	return int(ret0)
}

// GetNPages is a wrapper around gtk_assistant_get_n_pages().
func (assistant Assistant) GetNPages() int {
	ret0 := C.gtk_assistant_get_n_pages(assistant.native())
	return int(ret0)
}

// GetNthPage is a wrapper around gtk_assistant_get_nth_page().
func (assistant Assistant) GetNthPage(page_num int) Widget {
	ret0 := C.gtk_assistant_get_nth_page(assistant.native(), C.gint(page_num))
	return wrapWidget(ret0)
}

// GetPageComplete is a wrapper around gtk_assistant_get_page_complete().
func (assistant Assistant) GetPageComplete(page Widget) bool {
	ret0 := C.gtk_assistant_get_page_complete(assistant.native(), page.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPageHasPadding is a wrapper around gtk_assistant_get_page_has_padding().
func (assistant Assistant) GetPageHasPadding(page Widget) bool {
	ret0 := C.gtk_assistant_get_page_has_padding(assistant.native(), page.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPageTitle is a wrapper around gtk_assistant_get_page_title().
func (assistant Assistant) GetPageTitle(page Widget) string {
	ret0 := C.gtk_assistant_get_page_title(assistant.native(), page.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPageType is a wrapper around gtk_assistant_get_page_type().
func (assistant Assistant) GetPageType(page Widget) AssistantPageType {
	ret0 := C.gtk_assistant_get_page_type(assistant.native(), page.native())
	return AssistantPageType(ret0)
}

// InsertPage is a wrapper around gtk_assistant_insert_page().
func (assistant Assistant) InsertPage(page Widget, position int) int {
	ret0 := C.gtk_assistant_insert_page(assistant.native(), page.native(), C.gint(position))
	return int(ret0)
}

// NextPage is a wrapper around gtk_assistant_next_page().
func (assistant Assistant) NextPage() {
	C.gtk_assistant_next_page(assistant.native())
}

// PrependPage is a wrapper around gtk_assistant_prepend_page().
func (assistant Assistant) PrependPage(page Widget) int {
	ret0 := C.gtk_assistant_prepend_page(assistant.native(), page.native())
	return int(ret0)
}

// PreviousPage is a wrapper around gtk_assistant_previous_page().
func (assistant Assistant) PreviousPage() {
	C.gtk_assistant_previous_page(assistant.native())
}

// RemoveActionWidget is a wrapper around gtk_assistant_remove_action_widget().
func (assistant Assistant) RemoveActionWidget(child Widget) {
	C.gtk_assistant_remove_action_widget(assistant.native(), child.native())
}

// RemovePage is a wrapper around gtk_assistant_remove_page().
func (assistant Assistant) RemovePage(page_num int) {
	C.gtk_assistant_remove_page(assistant.native(), C.gint(page_num))
}

// SetCurrentPage is a wrapper around gtk_assistant_set_current_page().
func (assistant Assistant) SetCurrentPage(page_num int) {
	C.gtk_assistant_set_current_page(assistant.native(), C.gint(page_num))
}

// SetPageComplete is a wrapper around gtk_assistant_set_page_complete().
func (assistant Assistant) SetPageComplete(page Widget, complete bool) {
	C.gtk_assistant_set_page_complete(assistant.native(), page.native(), C.gboolean(util.Bool2Int(complete)) /*go:.util*/)
}

// SetPageHasPadding is a wrapper around gtk_assistant_set_page_has_padding().
func (assistant Assistant) SetPageHasPadding(page Widget, has_padding bool) {
	C.gtk_assistant_set_page_has_padding(assistant.native(), page.native(), C.gboolean(util.Bool2Int(has_padding)) /*go:.util*/)
}

// SetPageTitle is a wrapper around gtk_assistant_set_page_title().
func (assistant Assistant) SetPageTitle(page Widget, title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.gtk_assistant_set_page_title(assistant.native(), page.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetPageType is a wrapper around gtk_assistant_set_page_type().
func (assistant Assistant) SetPageType(page Widget, type_ AssistantPageType) {
	C.gtk_assistant_set_page_type(assistant.native(), page.native(), C.GtkAssistantPageType(type_))
}

// UpdateButtonsState is a wrapper around gtk_assistant_update_buttons_state().
func (assistant Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(assistant.native())
}

// Object Button
type Button struct {
	Bin
}

func (v Button) native() *C.GtkButton {
	return (*C.GtkButton)(v.Ptr)
}
func wrapButton(p *C.GtkButton) (v Button) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapButton(p unsafe.Pointer) (v Button) {
	v.Ptr = p
	return
}
func (v Button) IsNil() bool {
	return v.Ptr == nil
}
func IWrapButton(p unsafe.Pointer) interface{} {
	return WrapButton(p)
}
func (v Button) GetType() gobject.Type {
	return gobject.Type(C.gtk_button_get_type())
}
func (v Button) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapButton(unsafe.Pointer(ptr)), nil
	}
}
func (v Button) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Button) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v Button) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v Button) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ButtonNew is a wrapper around gtk_button_new().
func ButtonNew() Widget {
	ret0 := C.gtk_button_new()
	return wrapWidget(ret0)
}

// ButtonNewWithLabel is a wrapper around gtk_button_new_with_label().
func ButtonNewWithLabel(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_button_new_with_label(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// ButtonNewWithMnemonic is a wrapper around gtk_button_new_with_mnemonic().
func ButtonNewWithMnemonic(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_button_new_with_mnemonic(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// Clicked is a wrapper around gtk_button_clicked().
func (button Button) Clicked() {
	C.gtk_button_clicked(button.native())
}

// GetAlwaysShowImage is a wrapper around gtk_button_get_always_show_image().
func (button Button) GetAlwaysShowImage() bool {
	ret0 := C.gtk_button_get_always_show_image(button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetEventWindow is a wrapper around gtk_button_get_event_window().
func (button Button) GetEventWindow() gdk.Window {
	ret0 := C.gtk_button_get_event_window(button.native())
	return gdk.WrapWindow(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetImage is a wrapper around gtk_button_get_image().
func (button Button) GetImage() Widget {
	ret0 := C.gtk_button_get_image(button.native())
	return wrapWidget(ret0)
}

// GetImagePosition is a wrapper around gtk_button_get_image_position().
func (button Button) GetImagePosition() PositionType {
	ret0 := C.gtk_button_get_image_position(button.native())
	return PositionType(ret0)
}

// GetLabel is a wrapper around gtk_button_get_label().
func (button Button) GetLabel() string {
	ret0 := C.gtk_button_get_label(button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetRelief is a wrapper around gtk_button_get_relief().
func (button Button) GetRelief() ReliefStyle {
	ret0 := C.gtk_button_get_relief(button.native())
	return ReliefStyle(ret0)
}

// GetUseUnderline is a wrapper around gtk_button_get_use_underline().
func (button Button) GetUseUnderline() bool {
	ret0 := C.gtk_button_get_use_underline(button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetAlwaysShowImage is a wrapper around gtk_button_set_always_show_image().
func (button Button) SetAlwaysShowImage(always_show bool) {
	C.gtk_button_set_always_show_image(button.native(), C.gboolean(util.Bool2Int(always_show)) /*go:.util*/)
}

// SetImage is a wrapper around gtk_button_set_image().
func (button Button) SetImage(image Widget) {
	C.gtk_button_set_image(button.native(), image.native())
}

// SetImagePosition is a wrapper around gtk_button_set_image_position().
func (button Button) SetImagePosition(position PositionType) {
	C.gtk_button_set_image_position(button.native(), C.GtkPositionType(position))
}

// SetLabel is a wrapper around gtk_button_set_label().
func (button Button) SetLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_button_set_label(button.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetRelief is a wrapper around gtk_button_set_relief().
func (button Button) SetRelief(relief ReliefStyle) {
	C.gtk_button_set_relief(button.native(), C.GtkReliefStyle(relief))
}

// SetUseUnderline is a wrapper around gtk_button_set_use_underline().
func (button Button) SetUseUnderline(use_underline bool) {
	C.gtk_button_set_use_underline(button.native(), C.gboolean(util.Bool2Int(use_underline)) /*go:.util*/)
}

// Interface Activatable
type Activatable struct {
	Ptr unsafe.Pointer
}

func (v Activatable) native() *C.GtkActivatable {
	return (*C.GtkActivatable)(v.Ptr)
}
func wrapActivatable(p *C.GtkActivatable) Activatable {
	return Activatable{unsafe.Pointer(p)}
}
func WrapActivatable(p unsafe.Pointer) Activatable {
	return Activatable{p}
}
func (v Activatable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapActivatable(p unsafe.Pointer) interface{} {
	return WrapActivatable(p)
}
func (v Activatable) GetType() gobject.Type {
	return gobject.Type(C.gtk_activatable_get_type())
}
func (v Activatable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapActivatable(unsafe.Pointer(ptr)), nil
	}
}

// Object ButtonBox
type ButtonBox struct {
	Box
}

func (v ButtonBox) native() *C.GtkButtonBox {
	return (*C.GtkButtonBox)(v.Ptr)
}
func wrapButtonBox(p *C.GtkButtonBox) (v ButtonBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapButtonBox(p unsafe.Pointer) (v ButtonBox) {
	v.Ptr = p
	return
}
func (v ButtonBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapButtonBox(p unsafe.Pointer) interface{} {
	return WrapButtonBox(p)
}
func (v ButtonBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_button_box_get_type())
}
func (v ButtonBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapButtonBox(unsafe.Pointer(ptr)), nil
	}
}
func (v ButtonBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ButtonBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ButtonBox) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// ButtonBoxNew is a wrapper around gtk_button_box_new().
func ButtonBoxNew(orientation Orientation) Widget {
	ret0 := C.gtk_button_box_new(C.GtkOrientation(orientation))
	return wrapWidget(ret0)
}

// GetChildNonHomogeneous is a wrapper around gtk_button_box_get_child_non_homogeneous().
func (widget ButtonBox) GetChildNonHomogeneous(child Widget) bool {
	ret0 := C.gtk_button_box_get_child_non_homogeneous(widget.native(), child.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChildSecondary is a wrapper around gtk_button_box_get_child_secondary().
func (widget ButtonBox) GetChildSecondary(child Widget) bool {
	ret0 := C.gtk_button_box_get_child_secondary(widget.native(), child.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLayout is a wrapper around gtk_button_box_get_layout().
func (widget ButtonBox) GetLayout() ButtonBoxStyle {
	ret0 := C.gtk_button_box_get_layout(widget.native())
	return ButtonBoxStyle(ret0)
}

// SetChildNonHomogeneous is a wrapper around gtk_button_box_set_child_non_homogeneous().
func (widget ButtonBox) SetChildNonHomogeneous(child Widget, non_homogeneous bool) {
	C.gtk_button_box_set_child_non_homogeneous(widget.native(), child.native(), C.gboolean(util.Bool2Int(non_homogeneous)) /*go:.util*/)
}

// SetChildSecondary is a wrapper around gtk_button_box_set_child_secondary().
func (widget ButtonBox) SetChildSecondary(child Widget, is_secondary bool) {
	C.gtk_button_box_set_child_secondary(widget.native(), child.native(), C.gboolean(util.Bool2Int(is_secondary)) /*go:.util*/)
}

// SetLayout is a wrapper around gtk_button_box_set_layout().
func (widget ButtonBox) SetLayout(layout_style ButtonBoxStyle) {
	C.gtk_button_box_set_layout(widget.native(), C.GtkButtonBoxStyle(layout_style))
}

// Object Calendar
type Calendar struct {
	Widget
}

func (v Calendar) native() *C.GtkCalendar {
	return (*C.GtkCalendar)(v.Ptr)
}
func wrapCalendar(p *C.GtkCalendar) (v Calendar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCalendar(p unsafe.Pointer) (v Calendar) {
	v.Ptr = p
	return
}
func (v Calendar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCalendar(p unsafe.Pointer) interface{} {
	return WrapCalendar(p)
}
func (v Calendar) GetType() gobject.Type {
	return gobject.Type(C.gtk_calendar_get_type())
}
func (v Calendar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCalendar(unsafe.Pointer(ptr)), nil
	}
}
func (v Calendar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Calendar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// CalendarNew is a wrapper around gtk_calendar_new().
func CalendarNew() Widget {
	ret0 := C.gtk_calendar_new()
	return wrapWidget(ret0)
}

// ClearMarks is a wrapper around gtk_calendar_clear_marks().
func (calendar Calendar) ClearMarks() {
	C.gtk_calendar_clear_marks(calendar.native())
}

// GetDate is a wrapper around gtk_calendar_get_date().
func (calendar Calendar) GetDate() (uint, uint, uint) {
	var year0 C.guint
	var month0 C.guint
	var day0 C.guint
	C.gtk_calendar_get_date(calendar.native(), &year0, &month0, &day0)
	return uint(year0), uint(month0), uint(day0)
}

// GetDayIsMarked is a wrapper around gtk_calendar_get_day_is_marked().
func (calendar Calendar) GetDayIsMarked(day uint) bool {
	ret0 := C.gtk_calendar_get_day_is_marked(calendar.native(), C.guint(day))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDetailHeightRows is a wrapper around gtk_calendar_get_detail_height_rows().
func (calendar Calendar) GetDetailHeightRows() int {
	ret0 := C.gtk_calendar_get_detail_height_rows(calendar.native())
	return int(ret0)
}

// GetDetailWidthChars is a wrapper around gtk_calendar_get_detail_width_chars().
func (calendar Calendar) GetDetailWidthChars() int {
	ret0 := C.gtk_calendar_get_detail_width_chars(calendar.native())
	return int(ret0)
}

// GetDisplayOptions is a wrapper around gtk_calendar_get_display_options().
func (calendar Calendar) GetDisplayOptions() CalendarDisplayOptions {
	ret0 := C.gtk_calendar_get_display_options(calendar.native())
	return CalendarDisplayOptions(ret0)
}

// MarkDay is a wrapper around gtk_calendar_mark_day().
func (calendar Calendar) MarkDay(day uint) {
	C.gtk_calendar_mark_day(calendar.native(), C.guint(day))
}

// SelectDay is a wrapper around gtk_calendar_select_day().
func (calendar Calendar) SelectDay(day uint) {
	C.gtk_calendar_select_day(calendar.native(), C.guint(day))
}

// SelectMonth is a wrapper around gtk_calendar_select_month().
func (calendar Calendar) SelectMonth(month uint, year uint) {
	C.gtk_calendar_select_month(calendar.native(), C.guint(month), C.guint(year))
}

// SetDetailHeightRows is a wrapper around gtk_calendar_set_detail_height_rows().
func (calendar Calendar) SetDetailHeightRows(rows int) {
	C.gtk_calendar_set_detail_height_rows(calendar.native(), C.gint(rows))
}

// SetDetailWidthChars is a wrapper around gtk_calendar_set_detail_width_chars().
func (calendar Calendar) SetDetailWidthChars(chars int) {
	C.gtk_calendar_set_detail_width_chars(calendar.native(), C.gint(chars))
}

// SetDisplayOptions is a wrapper around gtk_calendar_set_display_options().
func (calendar Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	C.gtk_calendar_set_display_options(calendar.native(), C.GtkCalendarDisplayOptions(flags))
}

// UnmarkDay is a wrapper around gtk_calendar_unmark_day().
func (calendar Calendar) UnmarkDay(day uint) {
	C.gtk_calendar_unmark_day(calendar.native(), C.guint(day))
}

// Object CellAreaBox
type CellAreaBox struct {
	CellArea
}

func (v CellAreaBox) native() *C.GtkCellAreaBox {
	return (*C.GtkCellAreaBox)(v.Ptr)
}
func wrapCellAreaBox(p *C.GtkCellAreaBox) (v CellAreaBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellAreaBox(p unsafe.Pointer) (v CellAreaBox) {
	v.Ptr = p
	return
}
func (v CellAreaBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellAreaBox(p unsafe.Pointer) interface{} {
	return WrapCellAreaBox(p)
}
func (v CellAreaBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_area_box_get_type())
}
func (v CellAreaBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellAreaBox(unsafe.Pointer(ptr)), nil
	}
}
func (v CellAreaBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v CellAreaBox) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}
func (v CellAreaBox) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// CellAreaBoxNew is a wrapper around gtk_cell_area_box_new().
func CellAreaBoxNew() CellArea {
	ret0 := C.gtk_cell_area_box_new()
	return wrapCellArea(ret0)
}

// GetSpacing is a wrapper around gtk_cell_area_box_get_spacing().
func (box CellAreaBox) GetSpacing() int {
	ret0 := C.gtk_cell_area_box_get_spacing(box.native())
	return int(ret0)
}

// PackEnd is a wrapper around gtk_cell_area_box_pack_end().
func (box CellAreaBox) PackEnd(renderer CellRenderer, expand bool, align bool, fixed bool) {
	C.gtk_cell_area_box_pack_end(box.native(), renderer.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/, C.gboolean(util.Bool2Int(align)) /*go:.util*/, C.gboolean(util.Bool2Int(fixed)) /*go:.util*/)
}

// PackStart is a wrapper around gtk_cell_area_box_pack_start().
func (box CellAreaBox) PackStart(renderer CellRenderer, expand bool, align bool, fixed bool) {
	C.gtk_cell_area_box_pack_start(box.native(), renderer.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/, C.gboolean(util.Bool2Int(align)) /*go:.util*/, C.gboolean(util.Bool2Int(fixed)) /*go:.util*/)
}

// SetSpacing is a wrapper around gtk_cell_area_box_set_spacing().
func (box CellAreaBox) SetSpacing(spacing int) {
	C.gtk_cell_area_box_set_spacing(box.native(), C.gint(spacing))
}

// Object CellRendererAccel
type CellRendererAccel struct {
	CellRendererText
}

func (v CellRendererAccel) native() *C.GtkCellRendererAccel {
	return (*C.GtkCellRendererAccel)(v.Ptr)
}
func wrapCellRendererAccel(p *C.GtkCellRendererAccel) (v CellRendererAccel) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererAccel(p unsafe.Pointer) (v CellRendererAccel) {
	v.Ptr = p
	return
}
func (v CellRendererAccel) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererAccel(p unsafe.Pointer) interface{} {
	return WrapCellRendererAccel(p)
}
func (v CellRendererAccel) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_accel_get_type())
}
func (v CellRendererAccel) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererAccel(unsafe.Pointer(ptr)), nil
	}
}

// CellRendererAccelNew is a wrapper around gtk_cell_renderer_accel_new().
func CellRendererAccelNew() CellRenderer {
	ret0 := C.gtk_cell_renderer_accel_new()
	return wrapCellRenderer(ret0)
}

// Object CellRendererText
type CellRendererText struct {
	CellRenderer
}

func (v CellRendererText) native() *C.GtkCellRendererText {
	return (*C.GtkCellRendererText)(v.Ptr)
}
func wrapCellRendererText(p *C.GtkCellRendererText) (v CellRendererText) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererText(p unsafe.Pointer) (v CellRendererText) {
	v.Ptr = p
	return
}
func (v CellRendererText) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererText(p unsafe.Pointer) interface{} {
	return WrapCellRendererText(p)
}
func (v CellRendererText) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_text_get_type())
}
func (v CellRendererText) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererText(unsafe.Pointer(ptr)), nil
	}
}

// CellRendererTextNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererTextNew() CellRenderer {
	ret0 := C.gtk_cell_renderer_text_new()
	return wrapCellRenderer(ret0)
}

// SetFixedHeightFromFont is a wrapper around gtk_cell_renderer_text_set_fixed_height_from_font().
func (renderer CellRendererText) SetFixedHeightFromFont(number_of_rows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(renderer.native(), C.gint(number_of_rows))
}

// Object CellRendererCombo
type CellRendererCombo struct {
	CellRendererText
}

func (v CellRendererCombo) native() *C.GtkCellRendererCombo {
	return (*C.GtkCellRendererCombo)(v.Ptr)
}
func wrapCellRendererCombo(p *C.GtkCellRendererCombo) (v CellRendererCombo) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererCombo(p unsafe.Pointer) (v CellRendererCombo) {
	v.Ptr = p
	return
}
func (v CellRendererCombo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererCombo(p unsafe.Pointer) interface{} {
	return WrapCellRendererCombo(p)
}
func (v CellRendererCombo) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_combo_get_type())
}
func (v CellRendererCombo) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererCombo(unsafe.Pointer(ptr)), nil
	}
}

// CellRendererComboNew is a wrapper around gtk_cell_renderer_combo_new().
func CellRendererComboNew() CellRenderer {
	ret0 := C.gtk_cell_renderer_combo_new()
	return wrapCellRenderer(ret0)
}

// Object CellRendererPixbuf
type CellRendererPixbuf struct {
	CellRenderer
}

func (v CellRendererPixbuf) native() *C.GtkCellRendererPixbuf {
	return (*C.GtkCellRendererPixbuf)(v.Ptr)
}
func wrapCellRendererPixbuf(p *C.GtkCellRendererPixbuf) (v CellRendererPixbuf) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererPixbuf(p unsafe.Pointer) (v CellRendererPixbuf) {
	v.Ptr = p
	return
}
func (v CellRendererPixbuf) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererPixbuf(p unsafe.Pointer) interface{} {
	return WrapCellRendererPixbuf(p)
}
func (v CellRendererPixbuf) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_pixbuf_get_type())
}
func (v CellRendererPixbuf) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererPixbuf(unsafe.Pointer(ptr)), nil
	}
}

// CellRendererPixbufNew is a wrapper around gtk_cell_renderer_pixbuf_new().
func CellRendererPixbufNew() CellRenderer {
	ret0 := C.gtk_cell_renderer_pixbuf_new()
	return wrapCellRenderer(ret0)
}

// Object CellRendererProgress
type CellRendererProgress struct {
	CellRenderer
}

func (v CellRendererProgress) native() *C.GtkCellRendererProgress {
	return (*C.GtkCellRendererProgress)(v.Ptr)
}
func wrapCellRendererProgress(p *C.GtkCellRendererProgress) (v CellRendererProgress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererProgress(p unsafe.Pointer) (v CellRendererProgress) {
	v.Ptr = p
	return
}
func (v CellRendererProgress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererProgress(p unsafe.Pointer) interface{} {
	return WrapCellRendererProgress(p)
}
func (v CellRendererProgress) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_progress_get_type())
}
func (v CellRendererProgress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererProgress(unsafe.Pointer(ptr)), nil
	}
}
func (v CellRendererProgress) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// CellRendererProgressNew is a wrapper around gtk_cell_renderer_progress_new().
func CellRendererProgressNew() CellRenderer {
	ret0 := C.gtk_cell_renderer_progress_new()
	return wrapCellRenderer(ret0)
}

// Object CellRendererSpin
type CellRendererSpin struct {
	CellRendererText
}

func (v CellRendererSpin) native() *C.GtkCellRendererSpin {
	return (*C.GtkCellRendererSpin)(v.Ptr)
}
func wrapCellRendererSpin(p *C.GtkCellRendererSpin) (v CellRendererSpin) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererSpin(p unsafe.Pointer) (v CellRendererSpin) {
	v.Ptr = p
	return
}
func (v CellRendererSpin) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererSpin(p unsafe.Pointer) interface{} {
	return WrapCellRendererSpin(p)
}
func (v CellRendererSpin) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_spin_get_type())
}
func (v CellRendererSpin) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererSpin(unsafe.Pointer(ptr)), nil
	}
}

// CellRendererSpinNew is a wrapper around gtk_cell_renderer_spin_new().
func CellRendererSpinNew() CellRenderer {
	ret0 := C.gtk_cell_renderer_spin_new()
	return wrapCellRenderer(ret0)
}

// Object CellRendererSpinner
type CellRendererSpinner struct {
	CellRenderer
}

func (v CellRendererSpinner) native() *C.GtkCellRendererSpinner {
	return (*C.GtkCellRendererSpinner)(v.Ptr)
}
func wrapCellRendererSpinner(p *C.GtkCellRendererSpinner) (v CellRendererSpinner) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererSpinner(p unsafe.Pointer) (v CellRendererSpinner) {
	v.Ptr = p
	return
}
func (v CellRendererSpinner) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererSpinner(p unsafe.Pointer) interface{} {
	return WrapCellRendererSpinner(p)
}
func (v CellRendererSpinner) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_spinner_get_type())
}
func (v CellRendererSpinner) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererSpinner(unsafe.Pointer(ptr)), nil
	}
}

// CellRendererSpinnerNew is a wrapper around gtk_cell_renderer_spinner_new().
func CellRendererSpinnerNew() CellRenderer {
	ret0 := C.gtk_cell_renderer_spinner_new()
	return wrapCellRenderer(ret0)
}

// Object CellRendererToggle
type CellRendererToggle struct {
	CellRenderer
}

func (v CellRendererToggle) native() *C.GtkCellRendererToggle {
	return (*C.GtkCellRendererToggle)(v.Ptr)
}
func wrapCellRendererToggle(p *C.GtkCellRendererToggle) (v CellRendererToggle) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellRendererToggle(p unsafe.Pointer) (v CellRendererToggle) {
	v.Ptr = p
	return
}
func (v CellRendererToggle) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellRendererToggle(p unsafe.Pointer) interface{} {
	return WrapCellRendererToggle(p)
}
func (v CellRendererToggle) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_renderer_toggle_get_type())
}
func (v CellRendererToggle) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellRendererToggle(unsafe.Pointer(ptr)), nil
	}
}

// GetActivatable is a wrapper around gtk_cell_renderer_toggle_get_activatable().
func (toggle CellRendererToggle) GetActivatable() bool {
	ret0 := C.gtk_cell_renderer_toggle_get_activatable(toggle.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetActive is a wrapper around gtk_cell_renderer_toggle_get_active().
func (toggle CellRendererToggle) GetActive() bool {
	ret0 := C.gtk_cell_renderer_toggle_get_active(toggle.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRadio is a wrapper around gtk_cell_renderer_toggle_get_radio().
func (toggle CellRendererToggle) GetRadio() bool {
	ret0 := C.gtk_cell_renderer_toggle_get_radio(toggle.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetActivatable is a wrapper around gtk_cell_renderer_toggle_set_activatable().
func (toggle CellRendererToggle) SetActivatable(setting bool) {
	C.gtk_cell_renderer_toggle_set_activatable(toggle.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetActive is a wrapper around gtk_cell_renderer_toggle_set_active().
func (toggle CellRendererToggle) SetActive(setting bool) {
	C.gtk_cell_renderer_toggle_set_active(toggle.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetRadio is a wrapper around gtk_cell_renderer_toggle_set_radio().
func (toggle CellRendererToggle) SetRadio(radio bool) {
	C.gtk_cell_renderer_toggle_set_radio(toggle.native(), C.gboolean(util.Bool2Int(radio)) /*go:.util*/)
}

// Object CellView
type CellView struct {
	Widget
}

func (v CellView) native() *C.GtkCellView {
	return (*C.GtkCellView)(v.Ptr)
}
func wrapCellView(p *C.GtkCellView) (v CellView) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellView(p unsafe.Pointer) (v CellView) {
	v.Ptr = p
	return
}
func (v CellView) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellView(p unsafe.Pointer) interface{} {
	return WrapCellView(p)
}
func (v CellView) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_view_get_type())
}
func (v CellView) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellView(unsafe.Pointer(ptr)), nil
	}
}
func (v CellView) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v CellView) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v CellView) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}
func (v CellView) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// CellViewNew is a wrapper around gtk_cell_view_new().
func CellViewNew() Widget {
	ret0 := C.gtk_cell_view_new()
	return wrapWidget(ret0)
}

// CellViewNewWithContext is a wrapper around gtk_cell_view_new_with_context().
func CellViewNewWithContext(area CellArea, context CellAreaContext) Widget {
	ret0 := C.gtk_cell_view_new_with_context(area.native(), context.native())
	return wrapWidget(ret0)
}

// CellViewNewWithMarkup is a wrapper around gtk_cell_view_new_with_markup().
func CellViewNewWithMarkup(markup string) Widget {
	markup0 := (*C.gchar)(C.CString(markup))
	ret0 := C.gtk_cell_view_new_with_markup(markup0)
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// CellViewNewWithPixbuf is a wrapper around gtk_cell_view_new_with_pixbuf().
func CellViewNewWithPixbuf(pixbuf gdkpixbuf.Pixbuf) Widget {
	ret0 := C.gtk_cell_view_new_with_pixbuf((*C.GdkPixbuf)(pixbuf.Ptr))
	return wrapWidget(ret0)
}

// CellViewNewWithText is a wrapper around gtk_cell_view_new_with_text().
func CellViewNewWithText(text string) Widget {
	text0 := (*C.gchar)(C.CString(text))
	ret0 := C.gtk_cell_view_new_with_text(text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetDisplayedRow is a wrapper around gtk_cell_view_get_displayed_row().
func (cell_view CellView) GetDisplayedRow() TreePath {
	ret0 := C.gtk_cell_view_get_displayed_row(cell_view.native())
	return wrapTreePath(ret0)
}

// GetDrawSensitive is a wrapper around gtk_cell_view_get_draw_sensitive().
func (cell_view CellView) GetDrawSensitive() bool {
	ret0 := C.gtk_cell_view_get_draw_sensitive(cell_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFitModel is a wrapper around gtk_cell_view_get_fit_model().
func (cell_view CellView) GetFitModel() bool {
	ret0 := C.gtk_cell_view_get_fit_model(cell_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetModel is a wrapper around gtk_cell_view_get_model().
func (cell_view CellView) GetModel() TreeModel {
	ret0 := C.gtk_cell_view_get_model(cell_view.native())
	return wrapTreeModel(ret0)
}

// SetBackgroundRgba is a wrapper around gtk_cell_view_set_background_rgba().
func (cell_view CellView) SetBackgroundRgba(rgba gdk.RGBA) {
	C.gtk_cell_view_set_background_rgba(cell_view.native(), (*C.GdkRGBA)(rgba.Ptr))
}

// SetDisplayedRow is a wrapper around gtk_cell_view_set_displayed_row().
func (cell_view CellView) SetDisplayedRow(path TreePath) {
	C.gtk_cell_view_set_displayed_row(cell_view.native(), path.native())
}

// SetDrawSensitive is a wrapper around gtk_cell_view_set_draw_sensitive().
func (cell_view CellView) SetDrawSensitive(draw_sensitive bool) {
	C.gtk_cell_view_set_draw_sensitive(cell_view.native(), C.gboolean(util.Bool2Int(draw_sensitive)) /*go:.util*/)
}

// SetFitModel is a wrapper around gtk_cell_view_set_fit_model().
func (cell_view CellView) SetFitModel(fit_model bool) {
	C.gtk_cell_view_set_fit_model(cell_view.native(), C.gboolean(util.Bool2Int(fit_model)) /*go:.util*/)
}

// SetModel is a wrapper around gtk_cell_view_set_model().
func (cell_view CellView) SetModel(model TreeModel) {
	C.gtk_cell_view_set_model(cell_view.native(), model.native())
}

// Object CheckButton
type CheckButton struct {
	ToggleButton
}

func (v CheckButton) native() *C.GtkCheckButton {
	return (*C.GtkCheckButton)(v.Ptr)
}
func wrapCheckButton(p *C.GtkCheckButton) (v CheckButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCheckButton(p unsafe.Pointer) (v CheckButton) {
	v.Ptr = p
	return
}
func (v CheckButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCheckButton(p unsafe.Pointer) interface{} {
	return WrapCheckButton(p)
}
func (v CheckButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_check_button_get_type())
}
func (v CheckButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCheckButton(unsafe.Pointer(ptr)), nil
	}
}
func (v CheckButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v CheckButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v CheckButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v CheckButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// CheckButtonNew is a wrapper around gtk_check_button_new().
func CheckButtonNew() Widget {
	ret0 := C.gtk_check_button_new()
	return wrapWidget(ret0)
}

// CheckButtonNewWithLabel is a wrapper around gtk_check_button_new_with_label().
func CheckButtonNewWithLabel(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_check_button_new_with_label(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// CheckButtonNewWithMnemonic is a wrapper around gtk_check_button_new_with_mnemonic().
func CheckButtonNewWithMnemonic(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_check_button_new_with_mnemonic(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// Object ToggleButton
type ToggleButton struct {
	Button
}

func (v ToggleButton) native() *C.GtkToggleButton {
	return (*C.GtkToggleButton)(v.Ptr)
}
func wrapToggleButton(p *C.GtkToggleButton) (v ToggleButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToggleButton(p unsafe.Pointer) (v ToggleButton) {
	v.Ptr = p
	return
}
func (v ToggleButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToggleButton(p unsafe.Pointer) interface{} {
	return WrapToggleButton(p)
}
func (v ToggleButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_toggle_button_get_type())
}
func (v ToggleButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToggleButton(unsafe.Pointer(ptr)), nil
	}
}
func (v ToggleButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ToggleButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v ToggleButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ToggleButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ToggleButtonNew is a wrapper around gtk_toggle_button_new().
func ToggleButtonNew() Widget {
	ret0 := C.gtk_toggle_button_new()
	return wrapWidget(ret0)
}

// ToggleButtonNewWithLabel is a wrapper around gtk_toggle_button_new_with_label().
func ToggleButtonNewWithLabel(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_toggle_button_new_with_label(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// ToggleButtonNewWithMnemonic is a wrapper around gtk_toggle_button_new_with_mnemonic().
func ToggleButtonNewWithMnemonic(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_toggle_button_new_with_mnemonic(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (toggle_button ToggleButton) GetActive() bool {
	ret0 := C.gtk_toggle_button_get_active(toggle_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInconsistent is a wrapper around gtk_toggle_button_get_inconsistent().
func (toggle_button ToggleButton) GetInconsistent() bool {
	ret0 := C.gtk_toggle_button_get_inconsistent(toggle_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMode is a wrapper around gtk_toggle_button_get_mode().
func (toggle_button ToggleButton) GetMode() bool {
	ret0 := C.gtk_toggle_button_get_mode(toggle_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (toggle_button ToggleButton) SetActive(is_active bool) {
	C.gtk_toggle_button_set_active(toggle_button.native(), C.gboolean(util.Bool2Int(is_active)) /*go:.util*/)
}

// SetInconsistent is a wrapper around gtk_toggle_button_set_inconsistent().
func (toggle_button ToggleButton) SetInconsistent(setting bool) {
	C.gtk_toggle_button_set_inconsistent(toggle_button.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetMode is a wrapper around gtk_toggle_button_set_mode().
func (toggle_button ToggleButton) SetMode(draw_indicator bool) {
	C.gtk_toggle_button_set_mode(toggle_button.native(), C.gboolean(util.Bool2Int(draw_indicator)) /*go:.util*/)
}

// Toggled is a wrapper around gtk_toggle_button_toggled().
func (toggle_button ToggleButton) Toggled() {
	C.gtk_toggle_button_toggled(toggle_button.native())
}

// Object CheckMenuItem
type CheckMenuItem struct {
	MenuItem
}

func (v CheckMenuItem) native() *C.GtkCheckMenuItem {
	return (*C.GtkCheckMenuItem)(v.Ptr)
}
func wrapCheckMenuItem(p *C.GtkCheckMenuItem) (v CheckMenuItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCheckMenuItem(p unsafe.Pointer) (v CheckMenuItem) {
	v.Ptr = p
	return
}
func (v CheckMenuItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCheckMenuItem(p unsafe.Pointer) interface{} {
	return WrapCheckMenuItem(p)
}
func (v CheckMenuItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_check_menu_item_get_type())
}
func (v CheckMenuItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCheckMenuItem(unsafe.Pointer(ptr)), nil
	}
}
func (v CheckMenuItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v CheckMenuItem) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v CheckMenuItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v CheckMenuItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// CheckMenuItemNew is a wrapper around gtk_check_menu_item_new().
func CheckMenuItemNew() Widget {
	ret0 := C.gtk_check_menu_item_new()
	return wrapWidget(ret0)
}

// CheckMenuItemNewWithLabel is a wrapper around gtk_check_menu_item_new_with_label().
func CheckMenuItemNewWithLabel(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_check_menu_item_new_with_label(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// CheckMenuItemNewWithMnemonic is a wrapper around gtk_check_menu_item_new_with_mnemonic().
func CheckMenuItemNewWithMnemonic(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_check_menu_item_new_with_mnemonic(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetActive is a wrapper around gtk_check_menu_item_get_active().
func (check_menu_item CheckMenuItem) GetActive() bool {
	ret0 := C.gtk_check_menu_item_get_active(check_menu_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDrawAsRadio is a wrapper around gtk_check_menu_item_get_draw_as_radio().
func (check_menu_item CheckMenuItem) GetDrawAsRadio() bool {
	ret0 := C.gtk_check_menu_item_get_draw_as_radio(check_menu_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInconsistent is a wrapper around gtk_check_menu_item_get_inconsistent().
func (check_menu_item CheckMenuItem) GetInconsistent() bool {
	ret0 := C.gtk_check_menu_item_get_inconsistent(check_menu_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetActive is a wrapper around gtk_check_menu_item_set_active().
func (check_menu_item CheckMenuItem) SetActive(is_active bool) {
	C.gtk_check_menu_item_set_active(check_menu_item.native(), C.gboolean(util.Bool2Int(is_active)) /*go:.util*/)
}

// SetDrawAsRadio is a wrapper around gtk_check_menu_item_set_draw_as_radio().
func (check_menu_item CheckMenuItem) SetDrawAsRadio(draw_as_radio bool) {
	C.gtk_check_menu_item_set_draw_as_radio(check_menu_item.native(), C.gboolean(util.Bool2Int(draw_as_radio)) /*go:.util*/)
}

// SetInconsistent is a wrapper around gtk_check_menu_item_set_inconsistent().
func (check_menu_item CheckMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(check_menu_item.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// Toggled is a wrapper around gtk_check_menu_item_toggled().
func (check_menu_item CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(check_menu_item.native())
}

// Object MenuItem
type MenuItem struct {
	Bin
}

func (v MenuItem) native() *C.GtkMenuItem {
	return (*C.GtkMenuItem)(v.Ptr)
}
func wrapMenuItem(p *C.GtkMenuItem) (v MenuItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuItem(p unsafe.Pointer) (v MenuItem) {
	v.Ptr = p
	return
}
func (v MenuItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuItem(p unsafe.Pointer) interface{} {
	return WrapMenuItem(p)
}
func (v MenuItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_item_get_type())
}
func (v MenuItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuItem(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v MenuItem) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v MenuItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v MenuItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// MenuItemNew is a wrapper around gtk_menu_item_new().
func MenuItemNew() Widget {
	ret0 := C.gtk_menu_item_new()
	return wrapWidget(ret0)
}

// MenuItemNewWithLabel is a wrapper around gtk_menu_item_new_with_label().
func MenuItemNewWithLabel(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_menu_item_new_with_label(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// MenuItemNewWithMnemonic is a wrapper around gtk_menu_item_new_with_mnemonic().
func MenuItemNewWithMnemonic(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_menu_item_new_with_mnemonic(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// Activate is a wrapper around gtk_menu_item_activate().
func (menu_item MenuItem) Activate() {
	C.gtk_menu_item_activate(menu_item.native())
}

// Deselect is a wrapper around gtk_menu_item_deselect().
func (menu_item MenuItem) Deselect() {
	C.gtk_menu_item_deselect(menu_item.native())
}

// GetAccelPath is a wrapper around gtk_menu_item_get_accel_path().
func (menu_item MenuItem) GetAccelPath() string {
	ret0 := C.gtk_menu_item_get_accel_path(menu_item.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLabel is a wrapper around gtk_menu_item_get_label().
func (menu_item MenuItem) GetLabel() string {
	ret0 := C.gtk_menu_item_get_label(menu_item.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetReserveIndicator is a wrapper around gtk_menu_item_get_reserve_indicator().
func (menu_item MenuItem) GetReserveIndicator() bool {
	ret0 := C.gtk_menu_item_get_reserve_indicator(menu_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSubmenu is a wrapper around gtk_menu_item_get_submenu().
func (menu_item MenuItem) GetSubmenu() Widget {
	ret0 := C.gtk_menu_item_get_submenu(menu_item.native())
	return wrapWidget(ret0)
}

// GetUseUnderline is a wrapper around gtk_menu_item_get_use_underline().
func (menu_item MenuItem) GetUseUnderline() bool {
	ret0 := C.gtk_menu_item_get_use_underline(menu_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Select is a wrapper around gtk_menu_item_select().
func (menu_item MenuItem) Select() {
	C.gtk_menu_item_select(menu_item.native())
}

// SetAccelPath is a wrapper around gtk_menu_item_set_accel_path().
func (menu_item MenuItem) SetAccelPath(accel_path string) {
	accel_path0 := (*C.gchar)(C.CString(accel_path))
	C.gtk_menu_item_set_accel_path(menu_item.native(), accel_path0)
	C.free(unsafe.Pointer(accel_path0)) /*ch:<stdlib.h>*/
}

// SetLabel is a wrapper around gtk_menu_item_set_label().
func (menu_item MenuItem) SetLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_menu_item_set_label(menu_item.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetReserveIndicator is a wrapper around gtk_menu_item_set_reserve_indicator().
func (menu_item MenuItem) SetReserveIndicator(reserve bool) {
	C.gtk_menu_item_set_reserve_indicator(menu_item.native(), C.gboolean(util.Bool2Int(reserve)) /*go:.util*/)
}

// SetUseUnderline is a wrapper around gtk_menu_item_set_use_underline().
func (menu_item MenuItem) SetUseUnderline(setting bool) {
	C.gtk_menu_item_set_use_underline(menu_item.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// ToggleSizeAllocate is a wrapper around gtk_menu_item_toggle_size_allocate().
func (menu_item MenuItem) ToggleSizeAllocate(allocation int) {
	C.gtk_menu_item_toggle_size_allocate(menu_item.native(), C.gint(allocation))
}

// Object Menu
type Menu struct {
	MenuShell
}

func (v Menu) native() *C.GtkMenu {
	return (*C.GtkMenu)(v.Ptr)
}
func wrapMenu(p *C.GtkMenu) (v Menu) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenu(p unsafe.Pointer) (v Menu) {
	v.Ptr = p
	return
}
func (v Menu) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenu(p unsafe.Pointer) interface{} {
	return WrapMenu(p)
}
func (v Menu) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_get_type())
}
func (v Menu) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenu(unsafe.Pointer(ptr)), nil
	}
}
func (v Menu) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Menu) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// MenuNew is a wrapper around gtk_menu_new().
func MenuNew() Widget {
	ret0 := C.gtk_menu_new()
	return wrapWidget(ret0)
}

// MenuNewFromModel is a wrapper around gtk_menu_new_from_model().
func MenuNewFromModel(model gio.MenuModel) Widget {
	ret0 := C.gtk_menu_new_from_model((*C.GMenuModel)(model.Ptr))
	return wrapWidget(ret0)
}

// Attach is a wrapper around gtk_menu_attach().
func (menu Menu) Attach(child Widget, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint) {
	C.gtk_menu_attach(menu.native(), child.native(), C.guint(left_attach), C.guint(right_attach), C.guint(top_attach), C.guint(bottom_attach))
}

// Detach is a wrapper around gtk_menu_detach().
func (menu Menu) Detach() {
	C.gtk_menu_detach(menu.native())
}

// GetAccelGroup is a wrapper around gtk_menu_get_accel_group().
func (menu Menu) GetAccelGroup() AccelGroup {
	ret0 := C.gtk_menu_get_accel_group(menu.native())
	return wrapAccelGroup(ret0)
}

// GetAccelPath is a wrapper around gtk_menu_get_accel_path().
func (menu Menu) GetAccelPath() string {
	ret0 := C.gtk_menu_get_accel_path(menu.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetActive is a wrapper around gtk_menu_get_active().
func (menu Menu) GetActive() Widget {
	ret0 := C.gtk_menu_get_active(menu.native())
	return wrapWidget(ret0)
}

// GetAttachWidget is a wrapper around gtk_menu_get_attach_widget().
func (menu Menu) GetAttachWidget() Widget {
	ret0 := C.gtk_menu_get_attach_widget(menu.native())
	return wrapWidget(ret0)
}

// GetMonitor is a wrapper around gtk_menu_get_monitor().
func (menu Menu) GetMonitor() int {
	ret0 := C.gtk_menu_get_monitor(menu.native())
	return int(ret0)
}

// GetReserveToggleSize is a wrapper around gtk_menu_get_reserve_toggle_size().
func (menu Menu) GetReserveToggleSize() bool {
	ret0 := C.gtk_menu_get_reserve_toggle_size(menu.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Popdown is a wrapper around gtk_menu_popdown().
func (menu Menu) Popdown() {
	C.gtk_menu_popdown(menu.native())
}

// ReorderChild is a wrapper around gtk_menu_reorder_child().
func (menu Menu) ReorderChild(child Widget, position int) {
	C.gtk_menu_reorder_child(menu.native(), child.native(), C.gint(position))
}

// Reposition is a wrapper around gtk_menu_reposition().
func (menu Menu) Reposition() {
	C.gtk_menu_reposition(menu.native())
}

// SetAccelGroup is a wrapper around gtk_menu_set_accel_group().
func (menu Menu) SetAccelGroup(accel_group AccelGroup) {
	C.gtk_menu_set_accel_group(menu.native(), accel_group.native())
}

// SetAccelPath is a wrapper around gtk_menu_set_accel_path().
func (menu Menu) SetAccelPath(accel_path string) {
	accel_path0 := (*C.gchar)(C.CString(accel_path))
	C.gtk_menu_set_accel_path(menu.native(), accel_path0)
	C.free(unsafe.Pointer(accel_path0)) /*ch:<stdlib.h>*/
}

// SetActive is a wrapper around gtk_menu_set_active().
func (menu Menu) SetActive(index uint) {
	C.gtk_menu_set_active(menu.native(), C.guint(index))
}

// SetMonitor is a wrapper around gtk_menu_set_monitor().
func (menu Menu) SetMonitor(monitor_num int) {
	C.gtk_menu_set_monitor(menu.native(), C.gint(monitor_num))
}

// SetReserveToggleSize is a wrapper around gtk_menu_set_reserve_toggle_size().
func (menu Menu) SetReserveToggleSize(reserve_toggle_size bool) {
	C.gtk_menu_set_reserve_toggle_size(menu.native(), C.gboolean(util.Bool2Int(reserve_toggle_size)) /*go:.util*/)
}

// MenuGetForAttachWidget is a wrapper around gtk_menu_get_for_attach_widget().
func MenuGetForAttachWidget(widget Widget) glib.List {
	ret0 := C.gtk_menu_get_for_attach_widget(widget.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWidget(p) }) /*gir:GLib*/
}

// Object MenuShell
type MenuShell struct {
	Container
}

func (v MenuShell) native() *C.GtkMenuShell {
	return (*C.GtkMenuShell)(v.Ptr)
}
func wrapMenuShell(p *C.GtkMenuShell) (v MenuShell) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuShell(p unsafe.Pointer) (v MenuShell) {
	v.Ptr = p
	return
}
func (v MenuShell) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuShell(p unsafe.Pointer) interface{} {
	return WrapMenuShell(p)
}
func (v MenuShell) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_shell_get_type())
}
func (v MenuShell) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuShell(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuShell) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v MenuShell) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ActivateItem is a wrapper around gtk_menu_shell_activate_item().
func (menu_shell MenuShell) ActivateItem(menu_item Widget, force_deactivate bool) {
	C.gtk_menu_shell_activate_item(menu_shell.native(), menu_item.native(), C.gboolean(util.Bool2Int(force_deactivate)) /*go:.util*/)
}

// BindModel is a wrapper around gtk_menu_shell_bind_model().
func (menu_shell MenuShell) BindModel(model gio.MenuModel, action_namespace string, with_separators bool) {
	action_namespace0 := (*C.gchar)(C.CString(action_namespace))
	C.gtk_menu_shell_bind_model(menu_shell.native(), (*C.GMenuModel)(model.Ptr), action_namespace0, C.gboolean(util.Bool2Int(with_separators)) /*go:.util*/)
	C.free(unsafe.Pointer(action_namespace0)) /*ch:<stdlib.h>*/
}

// Cancel is a wrapper around gtk_menu_shell_cancel().
func (menu_shell MenuShell) Cancel() {
	C.gtk_menu_shell_cancel(menu_shell.native())
}

// Deselect is a wrapper around gtk_menu_shell_deselect().
func (menu_shell MenuShell) Deselect() {
	C.gtk_menu_shell_deselect(menu_shell.native())
}

// GetParentShell is a wrapper around gtk_menu_shell_get_parent_shell().
func (menu_shell MenuShell) GetParentShell() Widget {
	ret0 := C.gtk_menu_shell_get_parent_shell(menu_shell.native())
	return wrapWidget(ret0)
}

// GetSelectedItem is a wrapper around gtk_menu_shell_get_selected_item().
func (menu_shell MenuShell) GetSelectedItem() Widget {
	ret0 := C.gtk_menu_shell_get_selected_item(menu_shell.native())
	return wrapWidget(ret0)
}

// GetTakeFocus is a wrapper around gtk_menu_shell_get_take_focus().
func (menu_shell MenuShell) GetTakeFocus() bool {
	ret0 := C.gtk_menu_shell_get_take_focus(menu_shell.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Insert is a wrapper around gtk_menu_shell_insert().
func (menu_shell MenuShell) Insert(child Widget, position int) {
	C.gtk_menu_shell_insert(menu_shell.native(), child.native(), C.gint(position))
}

// Prepend is a wrapper around gtk_menu_shell_prepend().
func (menu_shell MenuShell) Prepend(child Widget) {
	C.gtk_menu_shell_prepend(menu_shell.native(), child.native())
}

// SelectFirst is a wrapper around gtk_menu_shell_select_first().
func (menu_shell MenuShell) SelectFirst(search_sensitive bool) {
	C.gtk_menu_shell_select_first(menu_shell.native(), C.gboolean(util.Bool2Int(search_sensitive)) /*go:.util*/)
}

// SelectItem is a wrapper around gtk_menu_shell_select_item().
func (menu_shell MenuShell) SelectItem(menu_item Widget) {
	C.gtk_menu_shell_select_item(menu_shell.native(), menu_item.native())
}

// SetTakeFocus is a wrapper around gtk_menu_shell_set_take_focus().
func (menu_shell MenuShell) SetTakeFocus(take_focus bool) {
	C.gtk_menu_shell_set_take_focus(menu_shell.native(), C.gboolean(util.Bool2Int(take_focus)) /*go:.util*/)
}

// Object Clipboard
type Clipboard struct {
	gobject.Object
}

func (v Clipboard) native() *C.GtkClipboard {
	return (*C.GtkClipboard)(v.Ptr)
}
func wrapClipboard(p *C.GtkClipboard) (v Clipboard) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapClipboard(p unsafe.Pointer) (v Clipboard) {
	v.Ptr = p
	return
}
func (v Clipboard) IsNil() bool {
	return v.Ptr == nil
}
func IWrapClipboard(p unsafe.Pointer) interface{} {
	return WrapClipboard(p)
}
func (v Clipboard) GetType() gobject.Type {
	return gobject.Type(C.gtk_clipboard_get_type())
}
func (v Clipboard) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapClipboard(unsafe.Pointer(ptr)), nil
	}
}

// Clear is a wrapper around gtk_clipboard_clear().
func (clipboard Clipboard) Clear() {
	C.gtk_clipboard_clear(clipboard.native())
}

// GetDisplay is a wrapper around gtk_clipboard_get_display().
func (clipboard Clipboard) GetDisplay() gdk.Display {
	ret0 := C.gtk_clipboard_get_display(clipboard.native())
	return gdk.WrapDisplay(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetOwner is a wrapper around gtk_clipboard_get_owner().
func (clipboard Clipboard) GetOwner() gobject.Object {
	ret0 := C.gtk_clipboard_get_owner(clipboard.native())
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// SetImage is a wrapper around gtk_clipboard_set_image().
func (clipboard Clipboard) SetImage(pixbuf gdkpixbuf.Pixbuf) {
	C.gtk_clipboard_set_image(clipboard.native(), (*C.GdkPixbuf)(pixbuf.Ptr))
}

// SetText is a wrapper around gtk_clipboard_set_text().
func (clipboard Clipboard) SetText(text string, len int) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_clipboard_set_text(clipboard.native(), text0, C.gint(len))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// Store is a wrapper around gtk_clipboard_store().
func (clipboard Clipboard) Store() {
	C.gtk_clipboard_store(clipboard.native())
}

// WaitForImage is a wrapper around gtk_clipboard_wait_for_image().
func (clipboard Clipboard) WaitForImage() gdkpixbuf.Pixbuf {
	ret0 := C.gtk_clipboard_wait_for_image(clipboard.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// WaitForText is a wrapper around gtk_clipboard_wait_for_text().
func (clipboard Clipboard) WaitForText() string {
	ret0 := C.gtk_clipboard_wait_for_text(clipboard.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// WaitForUris is a wrapper around gtk_clipboard_wait_for_uris().
func (clipboard Clipboard) WaitForUris() []string {
	ret0 := C.gtk_clipboard_wait_for_uris(clipboard.native())
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

// WaitIsImageAvailable is a wrapper around gtk_clipboard_wait_is_image_available().
func (clipboard Clipboard) WaitIsImageAvailable() bool {
	ret0 := C.gtk_clipboard_wait_is_image_available(clipboard.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// WaitIsRichTextAvailable is a wrapper around gtk_clipboard_wait_is_rich_text_available().
func (clipboard Clipboard) WaitIsRichTextAvailable(buffer TextBuffer) bool {
	ret0 := C.gtk_clipboard_wait_is_rich_text_available(clipboard.native(), buffer.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// WaitIsTextAvailable is a wrapper around gtk_clipboard_wait_is_text_available().
func (clipboard Clipboard) WaitIsTextAvailable() bool {
	ret0 := C.gtk_clipboard_wait_is_text_available(clipboard.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// WaitIsUrisAvailable is a wrapper around gtk_clipboard_wait_is_uris_available().
func (clipboard Clipboard) WaitIsUrisAvailable() bool {
	ret0 := C.gtk_clipboard_wait_is_uris_available(clipboard.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ClipboardGetDefault is a wrapper around gtk_clipboard_get_default().
func ClipboardGetDefault(display gdk.Display) Clipboard {
	ret0 := C.gtk_clipboard_get_default((*C.GdkDisplay)(display.Ptr))
	return wrapClipboard(ret0)
}

// Object ColorButton
type ColorButton struct {
	Button
}

func (v ColorButton) native() *C.GtkColorButton {
	return (*C.GtkColorButton)(v.Ptr)
}
func wrapColorButton(p *C.GtkColorButton) (v ColorButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapColorButton(p unsafe.Pointer) (v ColorButton) {
	v.Ptr = p
	return
}
func (v ColorButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColorButton(p unsafe.Pointer) interface{} {
	return WrapColorButton(p)
}
func (v ColorButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_color_button_get_type())
}
func (v ColorButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapColorButton(unsafe.Pointer(ptr)), nil
	}
}
func (v ColorButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ColorButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v ColorButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ColorButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ColorButton) ColorChooser() ColorChooser {
	return WrapColorChooser(v.Ptr)
}

// ColorButtonNew is a wrapper around gtk_color_button_new().
func ColorButtonNew() Widget {
	ret0 := C.gtk_color_button_new()
	return wrapWidget(ret0)
}

// ColorButtonNewWithRgba is a wrapper around gtk_color_button_new_with_rgba().
func ColorButtonNewWithRgba(rgba gdk.RGBA) Widget {
	ret0 := C.gtk_color_button_new_with_rgba((*C.GdkRGBA)(rgba.Ptr))
	return wrapWidget(ret0)
}

// GetTitle is a wrapper around gtk_color_button_get_title().
func (button ColorButton) GetTitle() string {
	ret0 := C.gtk_color_button_get_title(button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Object ColorChooserDialog
type ColorChooserDialog struct {
	Dialog
}

func (v ColorChooserDialog) native() *C.GtkColorChooserDialog {
	return (*C.GtkColorChooserDialog)(v.Ptr)
}
func wrapColorChooserDialog(p *C.GtkColorChooserDialog) (v ColorChooserDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapColorChooserDialog(p unsafe.Pointer) (v ColorChooserDialog) {
	v.Ptr = p
	return
}
func (v ColorChooserDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColorChooserDialog(p unsafe.Pointer) interface{} {
	return WrapColorChooserDialog(p)
}
func (v ColorChooserDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_color_chooser_dialog_get_type())
}
func (v ColorChooserDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapColorChooserDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v ColorChooserDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ColorChooserDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ColorChooserDialog) ColorChooser() ColorChooser {
	return WrapColorChooser(v.Ptr)
}

// ColorChooserDialogNew is a wrapper around gtk_color_chooser_dialog_new().
func ColorChooserDialogNew(title string, parent Window) Widget {
	title0 := (*C.gchar)(C.CString(title))
	ret0 := C.gtk_color_chooser_dialog_new(title0, parent.native())
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// Object ColorChooserWidget
type ColorChooserWidget struct {
	Box
}

func (v ColorChooserWidget) native() *C.GtkColorChooserWidget {
	return (*C.GtkColorChooserWidget)(v.Ptr)
}
func wrapColorChooserWidget(p *C.GtkColorChooserWidget) (v ColorChooserWidget) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapColorChooserWidget(p unsafe.Pointer) (v ColorChooserWidget) {
	v.Ptr = p
	return
}
func (v ColorChooserWidget) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColorChooserWidget(p unsafe.Pointer) interface{} {
	return WrapColorChooserWidget(p)
}
func (v ColorChooserWidget) GetType() gobject.Type {
	return gobject.Type(C.gtk_color_chooser_widget_get_type())
}
func (v ColorChooserWidget) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapColorChooserWidget(unsafe.Pointer(ptr)), nil
	}
}
func (v ColorChooserWidget) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ColorChooserWidget) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ColorChooserWidget) ColorChooser() ColorChooser {
	return WrapColorChooser(v.Ptr)
}
func (v ColorChooserWidget) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// ColorChooserWidgetNew is a wrapper around gtk_color_chooser_widget_new().
func ColorChooserWidgetNew() Widget {
	ret0 := C.gtk_color_chooser_widget_new()
	return wrapWidget(ret0)
}

// Object ColorSelection
type ColorSelection struct {
	Box
}

func (v ColorSelection) native() *C.GtkColorSelection {
	return (*C.GtkColorSelection)(v.Ptr)
}
func wrapColorSelection(p *C.GtkColorSelection) (v ColorSelection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapColorSelection(p unsafe.Pointer) (v ColorSelection) {
	v.Ptr = p
	return
}
func (v ColorSelection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColorSelection(p unsafe.Pointer) interface{} {
	return WrapColorSelection(p)
}
func (v ColorSelection) GetType() gobject.Type {
	return gobject.Type(C.gtk_color_selection_get_type())
}
func (v ColorSelection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapColorSelection(unsafe.Pointer(ptr)), nil
	}
}
func (v ColorSelection) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ColorSelection) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ColorSelection) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// ColorSelectionNew is a wrapper around gtk_color_selection_new().
func ColorSelectionNew() Widget {
	ret0 := C.gtk_color_selection_new()
	return wrapWidget(ret0)
}

// GetCurrentAlpha is a wrapper around gtk_color_selection_get_current_alpha().
func (colorsel ColorSelection) GetCurrentAlpha() uint16 {
	ret0 := C.gtk_color_selection_get_current_alpha(colorsel.native())
	return uint16(ret0)
}

// GetHasOpacityControl is a wrapper around gtk_color_selection_get_has_opacity_control().
func (colorsel ColorSelection) GetHasOpacityControl() bool {
	ret0 := C.gtk_color_selection_get_has_opacity_control(colorsel.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHasPalette is a wrapper around gtk_color_selection_get_has_palette().
func (colorsel ColorSelection) GetHasPalette() bool {
	ret0 := C.gtk_color_selection_get_has_palette(colorsel.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPreviousAlpha is a wrapper around gtk_color_selection_get_previous_alpha().
func (colorsel ColorSelection) GetPreviousAlpha() uint16 {
	ret0 := C.gtk_color_selection_get_previous_alpha(colorsel.native())
	return uint16(ret0)
}

// IsAdjusting is a wrapper around gtk_color_selection_is_adjusting().
func (colorsel ColorSelection) IsAdjusting() bool {
	ret0 := C.gtk_color_selection_is_adjusting(colorsel.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetCurrentAlpha is a wrapper around gtk_color_selection_set_current_alpha().
func (colorsel ColorSelection) SetCurrentAlpha(alpha uint16) {
	C.gtk_color_selection_set_current_alpha(colorsel.native(), C.guint16(alpha))
}

// SetCurrentRgba is a wrapper around gtk_color_selection_set_current_rgba().
func (colorsel ColorSelection) SetCurrentRgba(rgba gdk.RGBA) {
	C.gtk_color_selection_set_current_rgba(colorsel.native(), (*C.GdkRGBA)(rgba.Ptr))
}

// SetHasOpacityControl is a wrapper around gtk_color_selection_set_has_opacity_control().
func (colorsel ColorSelection) SetHasOpacityControl(has_opacity bool) {
	C.gtk_color_selection_set_has_opacity_control(colorsel.native(), C.gboolean(util.Bool2Int(has_opacity)) /*go:.util*/)
}

// SetHasPalette is a wrapper around gtk_color_selection_set_has_palette().
func (colorsel ColorSelection) SetHasPalette(has_palette bool) {
	C.gtk_color_selection_set_has_palette(colorsel.native(), C.gboolean(util.Bool2Int(has_palette)) /*go:.util*/)
}

// SetPreviousAlpha is a wrapper around gtk_color_selection_set_previous_alpha().
func (colorsel ColorSelection) SetPreviousAlpha(alpha uint16) {
	C.gtk_color_selection_set_previous_alpha(colorsel.native(), C.guint16(alpha))
}

// SetPreviousRgba is a wrapper around gtk_color_selection_set_previous_rgba().
func (colorsel ColorSelection) SetPreviousRgba(rgba gdk.RGBA) {
	C.gtk_color_selection_set_previous_rgba(colorsel.native(), (*C.GdkRGBA)(rgba.Ptr))
}

// Object ColorSelectionDialog
type ColorSelectionDialog struct {
	Dialog
}

func (v ColorSelectionDialog) native() *C.GtkColorSelectionDialog {
	return (*C.GtkColorSelectionDialog)(v.Ptr)
}
func wrapColorSelectionDialog(p *C.GtkColorSelectionDialog) (v ColorSelectionDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapColorSelectionDialog(p unsafe.Pointer) (v ColorSelectionDialog) {
	v.Ptr = p
	return
}
func (v ColorSelectionDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColorSelectionDialog(p unsafe.Pointer) interface{} {
	return WrapColorSelectionDialog(p)
}
func (v ColorSelectionDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_color_selection_dialog_get_type())
}
func (v ColorSelectionDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapColorSelectionDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v ColorSelectionDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ColorSelectionDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ColorSelectionDialogNew is a wrapper around gtk_color_selection_dialog_new().
func ColorSelectionDialogNew(title string) Widget {
	title0 := (*C.gchar)(C.CString(title))
	ret0 := C.gtk_color_selection_dialog_new(title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetColorSelection is a wrapper around gtk_color_selection_dialog_get_color_selection().
func (colorsel ColorSelectionDialog) GetColorSelection() Widget {
	ret0 := C.gtk_color_selection_dialog_get_color_selection(colorsel.native())
	return wrapWidget(ret0)
}

// Object ComboBoxAccessible
type ComboBoxAccessible struct {
	ContainerAccessible
}

func (v ComboBoxAccessible) native() *C.GtkComboBoxAccessible {
	return (*C.GtkComboBoxAccessible)(v.Ptr)
}
func wrapComboBoxAccessible(p *C.GtkComboBoxAccessible) (v ComboBoxAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapComboBoxAccessible(p unsafe.Pointer) (v ComboBoxAccessible) {
	v.Ptr = p
	return
}
func (v ComboBoxAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapComboBoxAccessible(p unsafe.Pointer) interface{} {
	return WrapComboBoxAccessible(p)
}
func (v ComboBoxAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_combo_box_accessible_get_type())
}
func (v ComboBoxAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapComboBoxAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ComboBoxAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v ComboBoxAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ComboBoxAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object ComboBoxText
type ComboBoxText struct {
	ComboBox
}

func (v ComboBoxText) native() *C.GtkComboBoxText {
	return (*C.GtkComboBoxText)(v.Ptr)
}
func wrapComboBoxText(p *C.GtkComboBoxText) (v ComboBoxText) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapComboBoxText(p unsafe.Pointer) (v ComboBoxText) {
	v.Ptr = p
	return
}
func (v ComboBoxText) IsNil() bool {
	return v.Ptr == nil
}
func IWrapComboBoxText(p unsafe.Pointer) interface{} {
	return WrapComboBoxText(p)
}
func (v ComboBoxText) GetType() gobject.Type {
	return gobject.Type(C.gtk_combo_box_text_get_type())
}
func (v ComboBoxText) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapComboBoxText(unsafe.Pointer(ptr)), nil
	}
}
func (v ComboBoxText) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ComboBoxText) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ComboBoxText) CellEditable() CellEditable {
	return WrapCellEditable(v.Ptr)
}
func (v ComboBoxText) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}

// ComboBoxTextNew is a wrapper around gtk_combo_box_text_new().
func ComboBoxTextNew() Widget {
	ret0 := C.gtk_combo_box_text_new()
	return wrapWidget(ret0)
}

// ComboBoxTextNewWithEntry is a wrapper around gtk_combo_box_text_new_with_entry().
func ComboBoxTextNewWithEntry() Widget {
	ret0 := C.gtk_combo_box_text_new_with_entry()
	return wrapWidget(ret0)
}

// Append is a wrapper around gtk_combo_box_text_append().
func (combo_box ComboBoxText) Append(id string, text string) {
	id0 := (*C.gchar)(C.CString(id))
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_combo_box_text_append(combo_box.native(), id0, text0)
	C.free(unsafe.Pointer(id0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// AppendText is a wrapper around gtk_combo_box_text_append_text().
func (combo_box ComboBoxText) AppendText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_combo_box_text_append_text(combo_box.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// GetActiveText is a wrapper around gtk_combo_box_text_get_active_text().
func (combo_box ComboBoxText) GetActiveText() string {
	ret0 := C.gtk_combo_box_text_get_active_text(combo_box.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Insert is a wrapper around gtk_combo_box_text_insert().
func (combo_box ComboBoxText) Insert(position int, id string, text string) {
	id0 := (*C.gchar)(C.CString(id))
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_combo_box_text_insert(combo_box.native(), C.gint(position), id0, text0)
	C.free(unsafe.Pointer(id0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// Prepend is a wrapper around gtk_combo_box_text_prepend().
func (combo_box ComboBoxText) Prepend(id string, text string) {
	id0 := (*C.gchar)(C.CString(id))
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_combo_box_text_prepend(combo_box.native(), id0, text0)
	C.free(unsafe.Pointer(id0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// PrependText is a wrapper around gtk_combo_box_text_prepend_text().
func (combo_box ComboBoxText) PrependText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_combo_box_text_prepend_text(combo_box.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// Remove is a wrapper around gtk_combo_box_text_remove().
func (combo_box ComboBoxText) Remove(position int) {
	C.gtk_combo_box_text_remove(combo_box.native(), C.gint(position))
}

// RemoveAll is a wrapper around gtk_combo_box_text_remove_all().
func (combo_box ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all(combo_box.native())
}

// Object ContainerAccessible
type ContainerAccessible struct {
	WidgetAccessible
}

func (v ContainerAccessible) native() *C.GtkContainerAccessible {
	return (*C.GtkContainerAccessible)(v.Ptr)
}
func wrapContainerAccessible(p *C.GtkContainerAccessible) (v ContainerAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapContainerAccessible(p unsafe.Pointer) (v ContainerAccessible) {
	v.Ptr = p
	return
}
func (v ContainerAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapContainerAccessible(p unsafe.Pointer) interface{} {
	return WrapContainerAccessible(p)
}
func (v ContainerAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_container_accessible_get_type())
}
func (v ContainerAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapContainerAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ContainerAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object WidgetAccessible
type WidgetAccessible struct {
	Accessible
}

func (v WidgetAccessible) native() *C.GtkWidgetAccessible {
	return (*C.GtkWidgetAccessible)(v.Ptr)
}
func wrapWidgetAccessible(p *C.GtkWidgetAccessible) (v WidgetAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapWidgetAccessible(p unsafe.Pointer) (v WidgetAccessible) {
	v.Ptr = p
	return
}
func (v WidgetAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWidgetAccessible(p unsafe.Pointer) interface{} {
	return WrapWidgetAccessible(p)
}
func (v WidgetAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_widget_accessible_get_type())
}
func (v WidgetAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapWidgetAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v WidgetAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object CheckMenuItemAccessible
type CheckMenuItemAccessible struct {
	MenuItemAccessible
}

func (v CheckMenuItemAccessible) native() *C.GtkCheckMenuItemAccessible {
	return (*C.GtkCheckMenuItemAccessible)(v.Ptr)
}
func wrapCheckMenuItemAccessible(p *C.GtkCheckMenuItemAccessible) (v CheckMenuItemAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCheckMenuItemAccessible(p unsafe.Pointer) (v CheckMenuItemAccessible) {
	v.Ptr = p
	return
}
func (v CheckMenuItemAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCheckMenuItemAccessible(p unsafe.Pointer) interface{} {
	return WrapCheckMenuItemAccessible(p)
}
func (v CheckMenuItemAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_check_menu_item_accessible_get_type())
}
func (v CheckMenuItemAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCheckMenuItemAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v CheckMenuItemAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v CheckMenuItemAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v CheckMenuItemAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object MenuItemAccessible
type MenuItemAccessible struct {
	ContainerAccessible
}

func (v MenuItemAccessible) native() *C.GtkMenuItemAccessible {
	return (*C.GtkMenuItemAccessible)(v.Ptr)
}
func wrapMenuItemAccessible(p *C.GtkMenuItemAccessible) (v MenuItemAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuItemAccessible(p unsafe.Pointer) (v MenuItemAccessible) {
	v.Ptr = p
	return
}
func (v MenuItemAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuItemAccessible(p unsafe.Pointer) interface{} {
	return WrapMenuItemAccessible(p)
}
func (v MenuItemAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_item_accessible_get_type())
}
func (v MenuItemAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuItemAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuItemAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v MenuItemAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v MenuItemAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object ButtonAccessible
type ButtonAccessible struct {
	ContainerAccessible
}

func (v ButtonAccessible) native() *C.GtkButtonAccessible {
	return (*C.GtkButtonAccessible)(v.Ptr)
}
func wrapButtonAccessible(p *C.GtkButtonAccessible) (v ButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapButtonAccessible(p unsafe.Pointer) (v ButtonAccessible) {
	v.Ptr = p
	return
}
func (v ButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapButtonAccessible(p)
}
func (v ButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_button_accessible_get_type())
}
func (v ButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v ButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ButtonAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object CellAccessible
type CellAccessible struct {
	Accessible
}

func (v CellAccessible) native() *C.GtkCellAccessible {
	return (*C.GtkCellAccessible)(v.Ptr)
}
func wrapCellAccessible(p *C.GtkCellAccessible) (v CellAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCellAccessible(p unsafe.Pointer) (v CellAccessible) {
	v.Ptr = p
	return
}
func (v CellAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellAccessible(p unsafe.Pointer) interface{} {
	return WrapCellAccessible(p)
}
func (v CellAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_accessible_get_type())
}
func (v CellAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v CellAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v CellAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object ContainerCellAccessible
type ContainerCellAccessible struct {
	CellAccessible
}

func (v ContainerCellAccessible) native() *C.GtkContainerCellAccessible {
	return (*C.GtkContainerCellAccessible)(v.Ptr)
}
func wrapContainerCellAccessible(p *C.GtkContainerCellAccessible) (v ContainerCellAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapContainerCellAccessible(p unsafe.Pointer) (v ContainerCellAccessible) {
	v.Ptr = p
	return
}
func (v ContainerCellAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapContainerCellAccessible(p unsafe.Pointer) interface{} {
	return WrapContainerCellAccessible(p)
}
func (v ContainerCellAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_container_cell_accessible_get_type())
}
func (v ContainerCellAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapContainerCellAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ContainerCellAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v ContainerCellAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// ContainerCellAccessibleNew is a wrapper around gtk_container_cell_accessible_new().
func ContainerCellAccessibleNew() ContainerCellAccessible {
	ret0 := C.gtk_container_cell_accessible_new()
	return wrapContainerCellAccessible(ret0)
}

// AddChild is a wrapper around gtk_container_cell_accessible_add_child().
func (container ContainerCellAccessible) AddChild(child CellAccessible) {
	C.gtk_container_cell_accessible_add_child(container.native(), child.native())
}

// GetChildren is a wrapper around gtk_container_cell_accessible_get_children().
func (container ContainerCellAccessible) GetChildren() glib.List {
	ret0 := C.gtk_container_cell_accessible_get_children(container.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapCellAccessible(p) }) /*gir:GLib*/
}

// RemoveChild is a wrapper around gtk_container_cell_accessible_remove_child().
func (container ContainerCellAccessible) RemoveChild(child CellAccessible) {
	C.gtk_container_cell_accessible_remove_child(container.native(), child.native())
}

// Object CssProvider
type CssProvider struct {
	gobject.Object
}

func (v CssProvider) native() *C.GtkCssProvider {
	return (*C.GtkCssProvider)(v.Ptr)
}
func wrapCssProvider(p *C.GtkCssProvider) (v CssProvider) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCssProvider(p unsafe.Pointer) (v CssProvider) {
	v.Ptr = p
	return
}
func (v CssProvider) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCssProvider(p unsafe.Pointer) interface{} {
	return WrapCssProvider(p)
}
func (v CssProvider) GetType() gobject.Type {
	return gobject.Type(C.gtk_css_provider_get_type())
}
func (v CssProvider) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCssProvider(unsafe.Pointer(ptr)), nil
	}
}
func (v CssProvider) StyleProvider() StyleProvider {
	return WrapStyleProvider(v.Ptr)
}

// CssProviderNew is a wrapper around gtk_css_provider_new().
func CssProviderNew() CssProvider {
	ret0 := C.gtk_css_provider_new()
	return wrapCssProvider(ret0)
}

// LoadFromData is a wrapper around gtk_css_provider_load_from_data().
func (css_provider CssProvider) LoadFromData(data []byte) (bool, error) {
	data0 := make([]C.gchar, len(data))
	for idx, elemG := range data {
		data0[idx] = C.gchar(elemG)
	}
	var data0Ptr *C.gchar
	if len(data0) > 0 {
		data0Ptr = &data0[0]
	}
	var err glib.Error
	ret0 := C.gtk_css_provider_load_from_data(css_provider.native(), data0Ptr, C.gssize(len(data)), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadFromFile is a wrapper around gtk_css_provider_load_from_file().
func (css_provider CssProvider) LoadFromFile(file gio.File) (bool, error) {
	var err glib.Error
	ret0 := C.gtk_css_provider_load_from_file(css_provider.native(), (*C.GFile)(file.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadFromPath is a wrapper around gtk_css_provider_load_from_path().
func (css_provider CssProvider) LoadFromPath(path string) (bool, error) {
	path0 := (*C.gchar)(C.CString(path))
	var err glib.Error
	ret0 := C.gtk_css_provider_load_from_path(css_provider.native(), path0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadFromResource is a wrapper around gtk_css_provider_load_from_resource().
func (css_provider CssProvider) LoadFromResource(resource_path string) {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	C.gtk_css_provider_load_from_resource(css_provider.native(), resource_path0)
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
}

// ToString is a wrapper around gtk_css_provider_to_string().
func (provider CssProvider) ToString() string {
	ret0 := C.gtk_css_provider_to_string(provider.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// CssProviderGetDefault is a wrapper around gtk_css_provider_get_default().
func CssProviderGetDefault() CssProvider {
	ret0 := C.gtk_css_provider_get_default()
	return wrapCssProvider(ret0)
}

// CssProviderGetNamed is a wrapper around gtk_css_provider_get_named().
func CssProviderGetNamed(name string, variant string) CssProvider {
	name0 := (*C.gchar)(C.CString(name))
	variant0 := (*C.gchar)(C.CString(variant))
	ret0 := C.gtk_css_provider_get_named(name0, variant0)
	C.free(unsafe.Pointer(name0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(variant0)) /*ch:<stdlib.h>*/
	return wrapCssProvider(ret0)
}

// Object DrawingArea
type DrawingArea struct {
	Widget
}

func (v DrawingArea) native() *C.GtkDrawingArea {
	return (*C.GtkDrawingArea)(v.Ptr)
}
func wrapDrawingArea(p *C.GtkDrawingArea) (v DrawingArea) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDrawingArea(p unsafe.Pointer) (v DrawingArea) {
	v.Ptr = p
	return
}
func (v DrawingArea) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDrawingArea(p unsafe.Pointer) interface{} {
	return WrapDrawingArea(p)
}
func (v DrawingArea) GetType() gobject.Type {
	return gobject.Type(C.gtk_drawing_area_get_type())
}
func (v DrawingArea) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDrawingArea(unsafe.Pointer(ptr)), nil
	}
}
func (v DrawingArea) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v DrawingArea) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// DrawingAreaNew is a wrapper around gtk_drawing_area_new().
func DrawingAreaNew() Widget {
	ret0 := C.gtk_drawing_area_new()
	return wrapWidget(ret0)
}

// Object Entry
type Entry struct {
	Widget
}

func (v Entry) native() *C.GtkEntry {
	return (*C.GtkEntry)(v.Ptr)
}
func wrapEntry(p *C.GtkEntry) (v Entry) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEntry(p unsafe.Pointer) (v Entry) {
	v.Ptr = p
	return
}
func (v Entry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEntry(p unsafe.Pointer) interface{} {
	return WrapEntry(p)
}
func (v Entry) GetType() gobject.Type {
	return gobject.Type(C.gtk_entry_get_type())
}
func (v Entry) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEntry(unsafe.Pointer(ptr)), nil
	}
}
func (v Entry) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Entry) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Entry) CellEditable() CellEditable {
	return WrapCellEditable(v.Ptr)
}
func (v Entry) Editable() Editable {
	return WrapEditable(v.Ptr)
}

// EntryNew is a wrapper around gtk_entry_new().
func EntryNew() Widget {
	ret0 := C.gtk_entry_new()
	return wrapWidget(ret0)
}

// EntryNewWithBuffer is a wrapper around gtk_entry_new_with_buffer().
func EntryNewWithBuffer(buffer EntryBuffer) Widget {
	ret0 := C.gtk_entry_new_with_buffer(buffer.native())
	return wrapWidget(ret0)
}

// GetActivatesDefault is a wrapper around gtk_entry_get_activates_default().
func (entry Entry) GetActivatesDefault() bool {
	ret0 := C.gtk_entry_get_activates_default(entry.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetAlignment is a wrapper around gtk_entry_get_alignment().
func (entry Entry) GetAlignment() float32 {
	ret0 := C.gtk_entry_get_alignment(entry.native())
	return float32(ret0)
}

// GetAttributes is a wrapper around gtk_entry_get_attributes().
func (entry Entry) GetAttributes() pango.AttrList {
	ret0 := C.gtk_entry_get_attributes(entry.native())
	return pango.WrapAttrList(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetBuffer is a wrapper around gtk_entry_get_buffer().
func (entry Entry) GetBuffer() EntryBuffer {
	ret0 := C.gtk_entry_get_buffer(entry.native())
	return wrapEntryBuffer(ret0)
}

// GetCompletion is a wrapper around gtk_entry_get_completion().
func (entry Entry) GetCompletion() EntryCompletion {
	ret0 := C.gtk_entry_get_completion(entry.native())
	return wrapEntryCompletion(ret0)
}

// GetCurrentIconDragSource is a wrapper around gtk_entry_get_current_icon_drag_source().
func (entry Entry) GetCurrentIconDragSource() int {
	ret0 := C.gtk_entry_get_current_icon_drag_source(entry.native())
	return int(ret0)
}

// GetHasFrame is a wrapper around gtk_entry_get_has_frame().
func (entry Entry) GetHasFrame() bool {
	ret0 := C.gtk_entry_get_has_frame(entry.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIconActivatable is a wrapper around gtk_entry_get_icon_activatable().
func (entry Entry) GetIconActivatable(icon_pos EntryIconPosition) bool {
	ret0 := C.gtk_entry_get_icon_activatable(entry.native(), C.GtkEntryIconPosition(icon_pos))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIconAtPos is a wrapper around gtk_entry_get_icon_at_pos().
func (entry Entry) GetIconAtPos(x int, y int) int {
	ret0 := C.gtk_entry_get_icon_at_pos(entry.native(), C.gint(x), C.gint(y))
	return int(ret0)
}

// GetIconGicon is a wrapper around gtk_entry_get_icon_gicon().
func (entry Entry) GetIconGicon(icon_pos EntryIconPosition) gio.Icon {
	ret0 := C.gtk_entry_get_icon_gicon(entry.native(), C.GtkEntryIconPosition(icon_pos))
	return gio.WrapIcon(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetIconName is a wrapper around gtk_entry_get_icon_name().
func (entry Entry) GetIconName(icon_pos EntryIconPosition) string {
	ret0 := C.gtk_entry_get_icon_name(entry.native(), C.GtkEntryIconPosition(icon_pos))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetIconPixbuf is a wrapper around gtk_entry_get_icon_pixbuf().
func (entry Entry) GetIconPixbuf(icon_pos EntryIconPosition) gdkpixbuf.Pixbuf {
	ret0 := C.gtk_entry_get_icon_pixbuf(entry.native(), C.GtkEntryIconPosition(icon_pos))
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetIconSensitive is a wrapper around gtk_entry_get_icon_sensitive().
func (entry Entry) GetIconSensitive(icon_pos EntryIconPosition) bool {
	ret0 := C.gtk_entry_get_icon_sensitive(entry.native(), C.GtkEntryIconPosition(icon_pos))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIconStorageType is a wrapper around gtk_entry_get_icon_storage_type().
func (entry Entry) GetIconStorageType(icon_pos EntryIconPosition) ImageType {
	ret0 := C.gtk_entry_get_icon_storage_type(entry.native(), C.GtkEntryIconPosition(icon_pos))
	return ImageType(ret0)
}

// GetIconTooltipMarkup is a wrapper around gtk_entry_get_icon_tooltip_markup().
func (entry Entry) GetIconTooltipMarkup(icon_pos EntryIconPosition) string {
	ret0 := C.gtk_entry_get_icon_tooltip_markup(entry.native(), C.GtkEntryIconPosition(icon_pos))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetIconTooltipText is a wrapper around gtk_entry_get_icon_tooltip_text().
func (entry Entry) GetIconTooltipText(icon_pos EntryIconPosition) string {
	ret0 := C.gtk_entry_get_icon_tooltip_text(entry.native(), C.GtkEntryIconPosition(icon_pos))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetInputHints is a wrapper around gtk_entry_get_input_hints().
func (entry Entry) GetInputHints() InputHints {
	ret0 := C.gtk_entry_get_input_hints(entry.native())
	return InputHints(ret0)
}

// GetInputPurpose is a wrapper around gtk_entry_get_input_purpose().
func (entry Entry) GetInputPurpose() InputPurpose {
	ret0 := C.gtk_entry_get_input_purpose(entry.native())
	return InputPurpose(ret0)
}

// GetLayout is a wrapper around gtk_entry_get_layout().
func (entry Entry) GetLayout() pango.Layout {
	ret0 := C.gtk_entry_get_layout(entry.native())
	return pango.WrapLayout(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetLayoutOffsets is a wrapper around gtk_entry_get_layout_offsets().
func (entry Entry) GetLayoutOffsets() (int, int) {
	var x0 C.gint
	var y0 C.gint
	C.gtk_entry_get_layout_offsets(entry.native(), &x0, &y0)
	return int(x0), int(y0)
}

// GetMaxLength is a wrapper around gtk_entry_get_max_length().
func (entry Entry) GetMaxLength() int {
	ret0 := C.gtk_entry_get_max_length(entry.native())
	return int(ret0)
}

// GetMaxWidthChars is a wrapper around gtk_entry_get_max_width_chars().
func (entry Entry) GetMaxWidthChars() int {
	ret0 := C.gtk_entry_get_max_width_chars(entry.native())
	return int(ret0)
}

// GetOverwriteMode is a wrapper around gtk_entry_get_overwrite_mode().
func (entry Entry) GetOverwriteMode() bool {
	ret0 := C.gtk_entry_get_overwrite_mode(entry.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPlaceholderText is a wrapper around gtk_entry_get_placeholder_text().
func (entry Entry) GetPlaceholderText() string {
	ret0 := C.gtk_entry_get_placeholder_text(entry.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetProgressFraction is a wrapper around gtk_entry_get_progress_fraction().
func (entry Entry) GetProgressFraction() float64 {
	ret0 := C.gtk_entry_get_progress_fraction(entry.native())
	return float64(ret0)
}

// GetProgressPulseStep is a wrapper around gtk_entry_get_progress_pulse_step().
func (entry Entry) GetProgressPulseStep() float64 {
	ret0 := C.gtk_entry_get_progress_pulse_step(entry.native())
	return float64(ret0)
}

// GetTabs is a wrapper around gtk_entry_get_tabs().
func (entry Entry) GetTabs() pango.TabArray {
	ret0 := C.gtk_entry_get_tabs(entry.native())
	return pango.WrapTabArray(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetText is a wrapper around gtk_entry_get_text().
func (entry Entry) GetText() string {
	ret0 := C.gtk_entry_get_text(entry.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTextLength is a wrapper around gtk_entry_get_text_length().
func (entry Entry) GetTextLength() uint16 {
	ret0 := C.gtk_entry_get_text_length(entry.native())
	return uint16(ret0)
}

// GetVisibility is a wrapper around gtk_entry_get_visibility().
func (entry Entry) GetVisibility() bool {
	ret0 := C.gtk_entry_get_visibility(entry.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetWidthChars is a wrapper around gtk_entry_get_width_chars().
func (entry Entry) GetWidthChars() int {
	ret0 := C.gtk_entry_get_width_chars(entry.native())
	return int(ret0)
}

// GrabFocusWithoutSelecting is a wrapper around gtk_entry_grab_focus_without_selecting().
func (entry Entry) GrabFocusWithoutSelecting() {
	C.gtk_entry_grab_focus_without_selecting(entry.native())
}

// ImContextFilterKeypress is a wrapper around gtk_entry_im_context_filter_keypress().
func (entry Entry) ImContextFilterKeypress(event gdk.EventKey) bool {
	ret0 := C.gtk_entry_im_context_filter_keypress(entry.native(), (*C.GdkEventKey)(event.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LayoutIndexToTextIndex is a wrapper around gtk_entry_layout_index_to_text_index().
func (entry Entry) LayoutIndexToTextIndex(layout_index int) int {
	ret0 := C.gtk_entry_layout_index_to_text_index(entry.native(), C.gint(layout_index))
	return int(ret0)
}

// ProgressPulse is a wrapper around gtk_entry_progress_pulse().
func (entry Entry) ProgressPulse() {
	C.gtk_entry_progress_pulse(entry.native())
}

// ResetImContext is a wrapper around gtk_entry_reset_im_context().
func (entry Entry) ResetImContext() {
	C.gtk_entry_reset_im_context(entry.native())
}

// SetActivatesDefault is a wrapper around gtk_entry_set_activates_default().
func (entry Entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(entry.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetAlignment is a wrapper around gtk_entry_set_alignment().
func (entry Entry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(entry.native(), C.gfloat(xalign))
}

// SetAttributes is a wrapper around gtk_entry_set_attributes().
func (entry Entry) SetAttributes(attrs pango.AttrList) {
	C.gtk_entry_set_attributes(entry.native(), (*C.PangoAttrList)(attrs.Ptr))
}

// SetBuffer is a wrapper around gtk_entry_set_buffer().
func (entry Entry) SetBuffer(buffer EntryBuffer) {
	C.gtk_entry_set_buffer(entry.native(), buffer.native())
}

// SetCompletion is a wrapper around gtk_entry_set_completion().
func (entry Entry) SetCompletion(completion EntryCompletion) {
	C.gtk_entry_set_completion(entry.native(), completion.native())
}

// SetCursorHadjustment is a wrapper around gtk_entry_set_cursor_hadjustment().
func (entry Entry) SetCursorHadjustment(adjustment Adjustment) {
	C.gtk_entry_set_cursor_hadjustment(entry.native(), adjustment.native())
}

// SetHasFrame is a wrapper around gtk_entry_set_has_frame().
func (entry Entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(entry.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetIconActivatable is a wrapper around gtk_entry_set_icon_activatable().
func (entry Entry) SetIconActivatable(icon_pos EntryIconPosition, activatable bool) {
	C.gtk_entry_set_icon_activatable(entry.native(), C.GtkEntryIconPosition(icon_pos), C.gboolean(util.Bool2Int(activatable)) /*go:.util*/)
}

// SetIconDragSource is a wrapper around gtk_entry_set_icon_drag_source().
func (entry Entry) SetIconDragSource(icon_pos EntryIconPosition, target_list TargetList, actions /*gir:Gdk*/ gdk.DragAction) {
	C.gtk_entry_set_icon_drag_source(entry.native(), C.GtkEntryIconPosition(icon_pos), target_list.native(), C.GdkDragAction(actions))
}

// SetIconFromGicon is a wrapper around gtk_entry_set_icon_from_gicon().
func (entry Entry) SetIconFromGicon(icon_pos EntryIconPosition, icon gio.Icon) {
	C.gtk_entry_set_icon_from_gicon(entry.native(), C.GtkEntryIconPosition(icon_pos), (*C.GIcon)(icon.Ptr))
}

// SetIconFromIconName is a wrapper around gtk_entry_set_icon_from_icon_name().
func (entry Entry) SetIconFromIconName(icon_pos EntryIconPosition, icon_name string) {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	C.gtk_entry_set_icon_from_icon_name(entry.native(), C.GtkEntryIconPosition(icon_pos), icon_name0)
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
}

// SetIconFromPixbuf is a wrapper around gtk_entry_set_icon_from_pixbuf().
func (entry Entry) SetIconFromPixbuf(icon_pos EntryIconPosition, pixbuf gdkpixbuf.Pixbuf) {
	C.gtk_entry_set_icon_from_pixbuf(entry.native(), C.GtkEntryIconPosition(icon_pos), (*C.GdkPixbuf)(pixbuf.Ptr))
}

// SetIconSensitive is a wrapper around gtk_entry_set_icon_sensitive().
func (entry Entry) SetIconSensitive(icon_pos EntryIconPosition, sensitive bool) {
	C.gtk_entry_set_icon_sensitive(entry.native(), C.GtkEntryIconPosition(icon_pos), C.gboolean(util.Bool2Int(sensitive)) /*go:.util*/)
}

// SetIconTooltipMarkup is a wrapper around gtk_entry_set_icon_tooltip_markup().
func (entry Entry) SetIconTooltipMarkup(icon_pos EntryIconPosition, tooltip string) {
	tooltip0 := (*C.gchar)(C.CString(tooltip))
	C.gtk_entry_set_icon_tooltip_markup(entry.native(), C.GtkEntryIconPosition(icon_pos), tooltip0)
	C.free(unsafe.Pointer(tooltip0)) /*ch:<stdlib.h>*/
}

// SetIconTooltipText is a wrapper around gtk_entry_set_icon_tooltip_text().
func (entry Entry) SetIconTooltipText(icon_pos EntryIconPosition, tooltip string) {
	tooltip0 := (*C.gchar)(C.CString(tooltip))
	C.gtk_entry_set_icon_tooltip_text(entry.native(), C.GtkEntryIconPosition(icon_pos), tooltip0)
	C.free(unsafe.Pointer(tooltip0)) /*ch:<stdlib.h>*/
}

// SetInputHints is a wrapper around gtk_entry_set_input_hints().
func (entry Entry) SetInputHints(hints InputHints) {
	C.gtk_entry_set_input_hints(entry.native(), C.GtkInputHints(hints))
}

// SetInputPurpose is a wrapper around gtk_entry_set_input_purpose().
func (entry Entry) SetInputPurpose(purpose InputPurpose) {
	C.gtk_entry_set_input_purpose(entry.native(), C.GtkInputPurpose(purpose))
}

// SetMaxLength is a wrapper around gtk_entry_set_max_length().
func (entry Entry) SetMaxLength(max int) {
	C.gtk_entry_set_max_length(entry.native(), C.gint(max))
}

// SetMaxWidthChars is a wrapper around gtk_entry_set_max_width_chars().
func (entry Entry) SetMaxWidthChars(n_chars int) {
	C.gtk_entry_set_max_width_chars(entry.native(), C.gint(n_chars))
}

// SetOverwriteMode is a wrapper around gtk_entry_set_overwrite_mode().
func (entry Entry) SetOverwriteMode(overwrite bool) {
	C.gtk_entry_set_overwrite_mode(entry.native(), C.gboolean(util.Bool2Int(overwrite)) /*go:.util*/)
}

// SetPlaceholderText is a wrapper around gtk_entry_set_placeholder_text().
func (entry Entry) SetPlaceholderText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_entry_set_placeholder_text(entry.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetProgressFraction is a wrapper around gtk_entry_set_progress_fraction().
func (entry Entry) SetProgressFraction(fraction float64) {
	C.gtk_entry_set_progress_fraction(entry.native(), C.gdouble(fraction))
}

// SetTabs is a wrapper around gtk_entry_set_tabs().
func (entry Entry) SetTabs(tabs pango.TabArray) {
	C.gtk_entry_set_tabs(entry.native(), (*C.PangoTabArray)(tabs.Ptr))
}

// SetText is a wrapper around gtk_entry_set_text().
func (entry Entry) SetText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_entry_set_text(entry.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetVisibility is a wrapper around gtk_entry_set_visibility().
func (entry Entry) SetVisibility(visible bool) {
	C.gtk_entry_set_visibility(entry.native(), C.gboolean(util.Bool2Int(visible)) /*go:.util*/)
}

// SetWidthChars is a wrapper around gtk_entry_set_width_chars().
func (entry Entry) SetWidthChars(n_chars int) {
	C.gtk_entry_set_width_chars(entry.native(), C.gint(n_chars))
}

// TextIndexToLayoutIndex is a wrapper around gtk_entry_text_index_to_layout_index().
func (entry Entry) TextIndexToLayoutIndex(text_index int) int {
	ret0 := C.gtk_entry_text_index_to_layout_index(entry.native(), C.gint(text_index))
	return int(ret0)
}

// UnsetInvisibleChar is a wrapper around gtk_entry_unset_invisible_char().
func (entry Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(entry.native())
}

// Object EntryBuffer
type EntryBuffer struct {
	gobject.Object
}

func (v EntryBuffer) native() *C.GtkEntryBuffer {
	return (*C.GtkEntryBuffer)(v.Ptr)
}
func wrapEntryBuffer(p *C.GtkEntryBuffer) (v EntryBuffer) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEntryBuffer(p unsafe.Pointer) (v EntryBuffer) {
	v.Ptr = p
	return
}
func (v EntryBuffer) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEntryBuffer(p unsafe.Pointer) interface{} {
	return WrapEntryBuffer(p)
}
func (v EntryBuffer) GetType() gobject.Type {
	return gobject.Type(C.gtk_entry_buffer_get_type())
}
func (v EntryBuffer) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEntryBuffer(unsafe.Pointer(ptr)), nil
	}
}

// EntryBufferNew is a wrapper around gtk_entry_buffer_new().
func EntryBufferNew(initial_chars string, n_initial_chars int) EntryBuffer {
	initial_chars0 := (*C.gchar)(C.CString(initial_chars))
	ret0 := C.gtk_entry_buffer_new(initial_chars0, C.gint(n_initial_chars))
	C.free(unsafe.Pointer(initial_chars0)) /*ch:<stdlib.h>*/
	return wrapEntryBuffer(ret0)
}

// DeleteText is a wrapper around gtk_entry_buffer_delete_text().
func (buffer EntryBuffer) DeleteText(position uint, n_chars int) uint {
	ret0 := C.gtk_entry_buffer_delete_text(buffer.native(), C.guint(position), C.gint(n_chars))
	return uint(ret0)
}

// EmitDeletedText is a wrapper around gtk_entry_buffer_emit_deleted_text().
func (buffer EntryBuffer) EmitDeletedText(position uint, n_chars uint) {
	C.gtk_entry_buffer_emit_deleted_text(buffer.native(), C.guint(position), C.guint(n_chars))
}

// EmitInsertedText is a wrapper around gtk_entry_buffer_emit_inserted_text().
func (buffer EntryBuffer) EmitInsertedText(position uint, chars string, n_chars uint) {
	chars0 := (*C.gchar)(C.CString(chars))
	C.gtk_entry_buffer_emit_inserted_text(buffer.native(), C.guint(position), chars0, C.guint(n_chars))
	C.free(unsafe.Pointer(chars0)) /*ch:<stdlib.h>*/
}

// GetBytes is a wrapper around gtk_entry_buffer_get_bytes().
func (buffer EntryBuffer) GetBytes() uint {
	ret0 := C.gtk_entry_buffer_get_bytes(buffer.native())
	return uint(ret0)
}

// GetLength is a wrapper around gtk_entry_buffer_get_length().
func (buffer EntryBuffer) GetLength() uint {
	ret0 := C.gtk_entry_buffer_get_length(buffer.native())
	return uint(ret0)
}

// GetMaxLength is a wrapper around gtk_entry_buffer_get_max_length().
func (buffer EntryBuffer) GetMaxLength() int {
	ret0 := C.gtk_entry_buffer_get_max_length(buffer.native())
	return int(ret0)
}

// GetText is a wrapper around gtk_entry_buffer_get_text().
func (buffer EntryBuffer) GetText() string {
	ret0 := C.gtk_entry_buffer_get_text(buffer.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// InsertText is a wrapper around gtk_entry_buffer_insert_text().
func (buffer EntryBuffer) InsertText(position uint, chars string, n_chars int) uint {
	chars0 := (*C.gchar)(C.CString(chars))
	ret0 := C.gtk_entry_buffer_insert_text(buffer.native(), C.guint(position), chars0, C.gint(n_chars))
	C.free(unsafe.Pointer(chars0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// SetMaxLength is a wrapper around gtk_entry_buffer_set_max_length().
func (buffer EntryBuffer) SetMaxLength(max_length int) {
	C.gtk_entry_buffer_set_max_length(buffer.native(), C.gint(max_length))
}

// SetText is a wrapper around gtk_entry_buffer_set_text().
func (buffer EntryBuffer) SetText(chars string, n_chars int) {
	chars0 := (*C.gchar)(C.CString(chars))
	C.gtk_entry_buffer_set_text(buffer.native(), chars0, C.gint(n_chars))
	C.free(unsafe.Pointer(chars0)) /*ch:<stdlib.h>*/
}

// Object EntryCompletion
type EntryCompletion struct {
	gobject.Object
}

func (v EntryCompletion) native() *C.GtkEntryCompletion {
	return (*C.GtkEntryCompletion)(v.Ptr)
}
func wrapEntryCompletion(p *C.GtkEntryCompletion) (v EntryCompletion) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEntryCompletion(p unsafe.Pointer) (v EntryCompletion) {
	v.Ptr = p
	return
}
func (v EntryCompletion) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEntryCompletion(p unsafe.Pointer) interface{} {
	return WrapEntryCompletion(p)
}
func (v EntryCompletion) GetType() gobject.Type {
	return gobject.Type(C.gtk_entry_completion_get_type())
}
func (v EntryCompletion) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEntryCompletion(unsafe.Pointer(ptr)), nil
	}
}
func (v EntryCompletion) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v EntryCompletion) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}

// EntryCompletionNew is a wrapper around gtk_entry_completion_new().
func EntryCompletionNew() EntryCompletion {
	ret0 := C.gtk_entry_completion_new()
	return wrapEntryCompletion(ret0)
}

// EntryCompletionNewWithArea is a wrapper around gtk_entry_completion_new_with_area().
func EntryCompletionNewWithArea(area CellArea) EntryCompletion {
	ret0 := C.gtk_entry_completion_new_with_area(area.native())
	return wrapEntryCompletion(ret0)
}

// Complete is a wrapper around gtk_entry_completion_complete().
func (completion EntryCompletion) Complete() {
	C.gtk_entry_completion_complete(completion.native())
}

// ComputePrefix is a wrapper around gtk_entry_completion_compute_prefix().
func (completion EntryCompletion) ComputePrefix(key string) string {
	key0 := C.CString(key)
	ret0 := C.gtk_entry_completion_compute_prefix(completion.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// DeleteAction is a wrapper around gtk_entry_completion_delete_action().
func (completion EntryCompletion) DeleteAction(index_ int) {
	C.gtk_entry_completion_delete_action(completion.native(), C.gint(index_))
}

// GetCompletionPrefix is a wrapper around gtk_entry_completion_get_completion_prefix().
func (completion EntryCompletion) GetCompletionPrefix() string {
	ret0 := C.gtk_entry_completion_get_completion_prefix(completion.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetEntry is a wrapper around gtk_entry_completion_get_entry().
func (completion EntryCompletion) GetEntry() Widget {
	ret0 := C.gtk_entry_completion_get_entry(completion.native())
	return wrapWidget(ret0)
}

// GetInlineCompletion is a wrapper around gtk_entry_completion_get_inline_completion().
func (completion EntryCompletion) GetInlineCompletion() bool {
	ret0 := C.gtk_entry_completion_get_inline_completion(completion.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInlineSelection is a wrapper around gtk_entry_completion_get_inline_selection().
func (completion EntryCompletion) GetInlineSelection() bool {
	ret0 := C.gtk_entry_completion_get_inline_selection(completion.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMinimumKeyLength is a wrapper around gtk_entry_completion_get_minimum_key_length().
func (completion EntryCompletion) GetMinimumKeyLength() int {
	ret0 := C.gtk_entry_completion_get_minimum_key_length(completion.native())
	return int(ret0)
}

// GetModel is a wrapper around gtk_entry_completion_get_model().
func (completion EntryCompletion) GetModel() TreeModel {
	ret0 := C.gtk_entry_completion_get_model(completion.native())
	return wrapTreeModel(ret0)
}

// GetPopupCompletion is a wrapper around gtk_entry_completion_get_popup_completion().
func (completion EntryCompletion) GetPopupCompletion() bool {
	ret0 := C.gtk_entry_completion_get_popup_completion(completion.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPopupSetWidth is a wrapper around gtk_entry_completion_get_popup_set_width().
func (completion EntryCompletion) GetPopupSetWidth() bool {
	ret0 := C.gtk_entry_completion_get_popup_set_width(completion.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPopupSingleMatch is a wrapper around gtk_entry_completion_get_popup_single_match().
func (completion EntryCompletion) GetPopupSingleMatch() bool {
	ret0 := C.gtk_entry_completion_get_popup_single_match(completion.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTextColumn is a wrapper around gtk_entry_completion_get_text_column().
func (completion EntryCompletion) GetTextColumn() int {
	ret0 := C.gtk_entry_completion_get_text_column(completion.native())
	return int(ret0)
}

// InsertActionMarkup is a wrapper around gtk_entry_completion_insert_action_markup().
func (completion EntryCompletion) InsertActionMarkup(index_ int, markup string) {
	markup0 := (*C.gchar)(C.CString(markup))
	C.gtk_entry_completion_insert_action_markup(completion.native(), C.gint(index_), markup0)
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// InsertActionText is a wrapper around gtk_entry_completion_insert_action_text().
func (completion EntryCompletion) InsertActionText(index_ int, text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_entry_completion_insert_action_text(completion.native(), C.gint(index_), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// InsertPrefix is a wrapper around gtk_entry_completion_insert_prefix().
func (completion EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(completion.native())
}

// SetInlineCompletion is a wrapper around gtk_entry_completion_set_inline_completion().
func (completion EntryCompletion) SetInlineCompletion(inline_completion bool) {
	C.gtk_entry_completion_set_inline_completion(completion.native(), C.gboolean(util.Bool2Int(inline_completion)) /*go:.util*/)
}

// SetInlineSelection is a wrapper around gtk_entry_completion_set_inline_selection().
func (completion EntryCompletion) SetInlineSelection(inline_selection bool) {
	C.gtk_entry_completion_set_inline_selection(completion.native(), C.gboolean(util.Bool2Int(inline_selection)) /*go:.util*/)
}

// SetMinimumKeyLength is a wrapper around gtk_entry_completion_set_minimum_key_length().
func (completion EntryCompletion) SetMinimumKeyLength(length int) {
	C.gtk_entry_completion_set_minimum_key_length(completion.native(), C.gint(length))
}

// SetModel is a wrapper around gtk_entry_completion_set_model().
func (completion EntryCompletion) SetModel(model TreeModel) {
	C.gtk_entry_completion_set_model(completion.native(), model.native())
}

// SetPopupCompletion is a wrapper around gtk_entry_completion_set_popup_completion().
func (completion EntryCompletion) SetPopupCompletion(popup_completion bool) {
	C.gtk_entry_completion_set_popup_completion(completion.native(), C.gboolean(util.Bool2Int(popup_completion)) /*go:.util*/)
}

// SetPopupSetWidth is a wrapper around gtk_entry_completion_set_popup_set_width().
func (completion EntryCompletion) SetPopupSetWidth(popup_set_width bool) {
	C.gtk_entry_completion_set_popup_set_width(completion.native(), C.gboolean(util.Bool2Int(popup_set_width)) /*go:.util*/)
}

// SetPopupSingleMatch is a wrapper around gtk_entry_completion_set_popup_single_match().
func (completion EntryCompletion) SetPopupSingleMatch(popup_single_match bool) {
	C.gtk_entry_completion_set_popup_single_match(completion.native(), C.gboolean(util.Bool2Int(popup_single_match)) /*go:.util*/)
}

// SetTextColumn is a wrapper around gtk_entry_completion_set_text_column().
func (completion EntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(completion.native(), C.gint(column))
}

// Object EntryAccessible
type EntryAccessible struct {
	WidgetAccessible
}

func (v EntryAccessible) native() *C.GtkEntryAccessible {
	return (*C.GtkEntryAccessible)(v.Ptr)
}
func wrapEntryAccessible(p *C.GtkEntryAccessible) (v EntryAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEntryAccessible(p unsafe.Pointer) (v EntryAccessible) {
	v.Ptr = p
	return
}
func (v EntryAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEntryAccessible(p unsafe.Pointer) interface{} {
	return WrapEntryAccessible(p)
}
func (v EntryAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_entry_accessible_get_type())
}
func (v EntryAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEntryAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v EntryAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v EntryAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v EntryAccessible) EditableText() atk.EditableText {
	return atk.WrapEditableText(v.Ptr) /*gir:Atk*/
}
func (v EntryAccessible) Text() atk.Text {
	return atk.WrapText(v.Ptr) /*gir:Atk*/
}

// Object EventBox
type EventBox struct {
	Bin
}

func (v EventBox) native() *C.GtkEventBox {
	return (*C.GtkEventBox)(v.Ptr)
}
func wrapEventBox(p *C.GtkEventBox) (v EventBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEventBox(p unsafe.Pointer) (v EventBox) {
	v.Ptr = p
	return
}
func (v EventBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventBox(p unsafe.Pointer) interface{} {
	return WrapEventBox(p)
}
func (v EventBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_event_box_get_type())
}
func (v EventBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEventBox(unsafe.Pointer(ptr)), nil
	}
}
func (v EventBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v EventBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// EventBoxNew is a wrapper around gtk_event_box_new().
func EventBoxNew() Widget {
	ret0 := C.gtk_event_box_new()
	return wrapWidget(ret0)
}

// GetAboveChild is a wrapper around gtk_event_box_get_above_child().
func (event_box EventBox) GetAboveChild() bool {
	ret0 := C.gtk_event_box_get_above_child(event_box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVisibleWindow is a wrapper around gtk_event_box_get_visible_window().
func (event_box EventBox) GetVisibleWindow() bool {
	ret0 := C.gtk_event_box_get_visible_window(event_box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetAboveChild is a wrapper around gtk_event_box_set_above_child().
func (event_box EventBox) SetAboveChild(above_child bool) {
	C.gtk_event_box_set_above_child(event_box.native(), C.gboolean(util.Bool2Int(above_child)) /*go:.util*/)
}

// SetVisibleWindow is a wrapper around gtk_event_box_set_visible_window().
func (event_box EventBox) SetVisibleWindow(visible_window bool) {
	C.gtk_event_box_set_visible_window(event_box.native(), C.gboolean(util.Bool2Int(visible_window)) /*go:.util*/)
}

// Object EventController
type EventController struct {
	gobject.Object
}

func (v EventController) native() *C.GtkEventController {
	return (*C.GtkEventController)(v.Ptr)
}
func wrapEventController(p *C.GtkEventController) (v EventController) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEventController(p unsafe.Pointer) (v EventController) {
	v.Ptr = p
	return
}
func (v EventController) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventController(p unsafe.Pointer) interface{} {
	return WrapEventController(p)
}
func (v EventController) GetType() gobject.Type {
	return gobject.Type(C.gtk_event_controller_get_type())
}
func (v EventController) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEventController(unsafe.Pointer(ptr)), nil
	}
}

// GetPropagationPhase is a wrapper around gtk_event_controller_get_propagation_phase().
func (controller EventController) GetPropagationPhase() PropagationPhase {
	ret0 := C.gtk_event_controller_get_propagation_phase(controller.native())
	return PropagationPhase(ret0)
}

// GetWidget is a wrapper around gtk_event_controller_get_widget().
func (controller EventController) GetWidget() Widget {
	ret0 := C.gtk_event_controller_get_widget(controller.native())
	return wrapWidget(ret0)
}

// Reset is a wrapper around gtk_event_controller_reset().
func (controller EventController) Reset() {
	C.gtk_event_controller_reset(controller.native())
}

// SetPropagationPhase is a wrapper around gtk_event_controller_set_propagation_phase().
func (controller EventController) SetPropagationPhase(phase PropagationPhase) {
	C.gtk_event_controller_set_propagation_phase(controller.native(), C.GtkPropagationPhase(phase))
}

// Object Expander
type Expander struct {
	Bin
}

func (v Expander) native() *C.GtkExpander {
	return (*C.GtkExpander)(v.Ptr)
}
func wrapExpander(p *C.GtkExpander) (v Expander) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapExpander(p unsafe.Pointer) (v Expander) {
	v.Ptr = p
	return
}
func (v Expander) IsNil() bool {
	return v.Ptr == nil
}
func IWrapExpander(p unsafe.Pointer) interface{} {
	return WrapExpander(p)
}
func (v Expander) GetType() gobject.Type {
	return gobject.Type(C.gtk_expander_get_type())
}
func (v Expander) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapExpander(unsafe.Pointer(ptr)), nil
	}
}
func (v Expander) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Expander) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ExpanderNewWithMnemonic is a wrapper around gtk_expander_new_with_mnemonic().
func ExpanderNewWithMnemonic(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_expander_new_with_mnemonic(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetExpanded is a wrapper around gtk_expander_get_expanded().
func (expander Expander) GetExpanded() bool {
	ret0 := C.gtk_expander_get_expanded(expander.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLabel is a wrapper around gtk_expander_get_label().
func (expander Expander) GetLabel() string {
	ret0 := C.gtk_expander_get_label(expander.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLabelFill is a wrapper around gtk_expander_get_label_fill().
func (expander Expander) GetLabelFill() bool {
	ret0 := C.gtk_expander_get_label_fill(expander.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLabelWidget is a wrapper around gtk_expander_get_label_widget().
func (expander Expander) GetLabelWidget() Widget {
	ret0 := C.gtk_expander_get_label_widget(expander.native())
	return wrapWidget(ret0)
}

// GetResizeToplevel is a wrapper around gtk_expander_get_resize_toplevel().
func (expander Expander) GetResizeToplevel() bool {
	ret0 := C.gtk_expander_get_resize_toplevel(expander.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUseMarkup is a wrapper around gtk_expander_get_use_markup().
func (expander Expander) GetUseMarkup() bool {
	ret0 := C.gtk_expander_get_use_markup(expander.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUseUnderline is a wrapper around gtk_expander_get_use_underline().
func (expander Expander) GetUseUnderline() bool {
	ret0 := C.gtk_expander_get_use_underline(expander.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetExpanded is a wrapper around gtk_expander_set_expanded().
func (expander Expander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(expander.native(), C.gboolean(util.Bool2Int(expanded)) /*go:.util*/)
}

// SetLabel is a wrapper around gtk_expander_set_label().
func (expander Expander) SetLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_expander_set_label(expander.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetLabelFill is a wrapper around gtk_expander_set_label_fill().
func (expander Expander) SetLabelFill(label_fill bool) {
	C.gtk_expander_set_label_fill(expander.native(), C.gboolean(util.Bool2Int(label_fill)) /*go:.util*/)
}

// SetLabelWidget is a wrapper around gtk_expander_set_label_widget().
func (expander Expander) SetLabelWidget(label_widget Widget) {
	C.gtk_expander_set_label_widget(expander.native(), label_widget.native())
}

// SetResizeToplevel is a wrapper around gtk_expander_set_resize_toplevel().
func (expander Expander) SetResizeToplevel(resize_toplevel bool) {
	C.gtk_expander_set_resize_toplevel(expander.native(), C.gboolean(util.Bool2Int(resize_toplevel)) /*go:.util*/)
}

// SetUseMarkup is a wrapper around gtk_expander_set_use_markup().
func (expander Expander) SetUseMarkup(use_markup bool) {
	C.gtk_expander_set_use_markup(expander.native(), C.gboolean(util.Bool2Int(use_markup)) /*go:.util*/)
}

// SetUseUnderline is a wrapper around gtk_expander_set_use_underline().
func (expander Expander) SetUseUnderline(use_underline bool) {
	C.gtk_expander_set_use_underline(expander.native(), C.gboolean(util.Bool2Int(use_underline)) /*go:.util*/)
}

// Object ExpanderAccessible
type ExpanderAccessible struct {
	ContainerAccessible
}

func (v ExpanderAccessible) native() *C.GtkExpanderAccessible {
	return (*C.GtkExpanderAccessible)(v.Ptr)
}
func wrapExpanderAccessible(p *C.GtkExpanderAccessible) (v ExpanderAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapExpanderAccessible(p unsafe.Pointer) (v ExpanderAccessible) {
	v.Ptr = p
	return
}
func (v ExpanderAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapExpanderAccessible(p unsafe.Pointer) interface{} {
	return WrapExpanderAccessible(p)
}
func (v ExpanderAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_expander_accessible_get_type())
}
func (v ExpanderAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapExpanderAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ExpanderAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v ExpanderAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object FileChooserButton
type FileChooserButton struct {
	Box
}

func (v FileChooserButton) native() *C.GtkFileChooserButton {
	return (*C.GtkFileChooserButton)(v.Ptr)
}
func wrapFileChooserButton(p *C.GtkFileChooserButton) (v FileChooserButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileChooserButton(p unsafe.Pointer) (v FileChooserButton) {
	v.Ptr = p
	return
}
func (v FileChooserButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileChooserButton(p unsafe.Pointer) interface{} {
	return WrapFileChooserButton(p)
}
func (v FileChooserButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_file_chooser_button_get_type())
}
func (v FileChooserButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileChooserButton(unsafe.Pointer(ptr)), nil
	}
}
func (v FileChooserButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FileChooserButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FileChooserButton) FileChooser() FileChooser {
	return WrapFileChooser(v.Ptr)
}
func (v FileChooserButton) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// FileChooserButtonNew is a wrapper around gtk_file_chooser_button_new().
func FileChooserButtonNew(title string, action FileChooserAction) Widget {
	title0 := (*C.gchar)(C.CString(title))
	ret0 := C.gtk_file_chooser_button_new(title0, C.GtkFileChooserAction(action))
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetTitle is a wrapper around gtk_file_chooser_button_get_title().
func (button FileChooserButton) GetTitle() string {
	ret0 := C.gtk_file_chooser_button_get_title(button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWidthChars is a wrapper around gtk_file_chooser_button_get_width_chars().
func (button FileChooserButton) GetWidthChars() int {
	ret0 := C.gtk_file_chooser_button_get_width_chars(button.native())
	return int(ret0)
}

// SetTitle is a wrapper around gtk_file_chooser_button_set_title().
func (button FileChooserButton) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.gtk_file_chooser_button_set_title(button.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetWidthChars is a wrapper around gtk_file_chooser_button_set_width_chars().
func (button FileChooserButton) SetWidthChars(n_chars int) {
	C.gtk_file_chooser_button_set_width_chars(button.native(), C.gint(n_chars))
}

// Object FileChooserDialog
type FileChooserDialog struct {
	Dialog
}

func (v FileChooserDialog) native() *C.GtkFileChooserDialog {
	return (*C.GtkFileChooserDialog)(v.Ptr)
}
func wrapFileChooserDialog(p *C.GtkFileChooserDialog) (v FileChooserDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileChooserDialog(p unsafe.Pointer) (v FileChooserDialog) {
	v.Ptr = p
	return
}
func (v FileChooserDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileChooserDialog(p unsafe.Pointer) interface{} {
	return WrapFileChooserDialog(p)
}
func (v FileChooserDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_file_chooser_dialog_get_type())
}
func (v FileChooserDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileChooserDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v FileChooserDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FileChooserDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FileChooserDialog) FileChooser() FileChooser {
	return WrapFileChooser(v.Ptr)
}

// Object FileChooserNative
type FileChooserNative struct {
	NativeDialog
}

func (v FileChooserNative) native() *C.GtkFileChooserNative {
	return (*C.GtkFileChooserNative)(v.Ptr)
}
func wrapFileChooserNative(p *C.GtkFileChooserNative) (v FileChooserNative) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileChooserNative(p unsafe.Pointer) (v FileChooserNative) {
	v.Ptr = p
	return
}
func (v FileChooserNative) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileChooserNative(p unsafe.Pointer) interface{} {
	return WrapFileChooserNative(p)
}
func (v FileChooserNative) GetType() gobject.Type {
	return gobject.Type(C.gtk_file_chooser_native_get_type())
}
func (v FileChooserNative) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileChooserNative(unsafe.Pointer(ptr)), nil
	}
}
func (v FileChooserNative) FileChooser() FileChooser {
	return WrapFileChooser(v.Ptr)
}

// GetAcceptLabel is a wrapper around gtk_file_chooser_native_get_accept_label().
func (self FileChooserNative) GetAcceptLabel() string {
	ret0 := C.gtk_file_chooser_native_get_accept_label(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetCancelLabel is a wrapper around gtk_file_chooser_native_get_cancel_label().
func (self FileChooserNative) GetCancelLabel() string {
	ret0 := C.gtk_file_chooser_native_get_cancel_label(self.native())
	ret := C.GoString(ret0)
	return ret
}

// SetAcceptLabel is a wrapper around gtk_file_chooser_native_set_accept_label().
func (self FileChooserNative) SetAcceptLabel(accept_label string) {
	accept_label0 := C.CString(accept_label)
	C.gtk_file_chooser_native_set_accept_label(self.native(), accept_label0)
	C.free(unsafe.Pointer(accept_label0)) /*ch:<stdlib.h>*/
}

// SetCancelLabel is a wrapper around gtk_file_chooser_native_set_cancel_label().
func (self FileChooserNative) SetCancelLabel(cancel_label string) {
	cancel_label0 := C.CString(cancel_label)
	C.gtk_file_chooser_native_set_cancel_label(self.native(), cancel_label0)
	C.free(unsafe.Pointer(cancel_label0)) /*ch:<stdlib.h>*/
}

// Object NativeDialog
type NativeDialog struct {
	gobject.Object
}

func (v NativeDialog) native() *C.GtkNativeDialog {
	return (*C.GtkNativeDialog)(v.Ptr)
}
func wrapNativeDialog(p *C.GtkNativeDialog) (v NativeDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNativeDialog(p unsafe.Pointer) (v NativeDialog) {
	v.Ptr = p
	return
}
func (v NativeDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNativeDialog(p unsafe.Pointer) interface{} {
	return WrapNativeDialog(p)
}
func (v NativeDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_native_dialog_get_type())
}
func (v NativeDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNativeDialog(unsafe.Pointer(ptr)), nil
	}
}

// Destroy is a wrapper around gtk_native_dialog_destroy().
func (self NativeDialog) Destroy() {
	C.gtk_native_dialog_destroy(self.native())
}

// GetModal is a wrapper around gtk_native_dialog_get_modal().
func (self NativeDialog) GetModal() bool {
	ret0 := C.gtk_native_dialog_get_modal(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTitle is a wrapper around gtk_native_dialog_get_title().
func (self NativeDialog) GetTitle() string {
	ret0 := C.gtk_native_dialog_get_title(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetTransientFor is a wrapper around gtk_native_dialog_get_transient_for().
func (self NativeDialog) GetTransientFor() Window {
	ret0 := C.gtk_native_dialog_get_transient_for(self.native())
	return wrapWindow(ret0)
}

// GetVisible is a wrapper around gtk_native_dialog_get_visible().
func (self NativeDialog) GetVisible() bool {
	ret0 := C.gtk_native_dialog_get_visible(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Hide is a wrapper around gtk_native_dialog_hide().
func (self NativeDialog) Hide() {
	C.gtk_native_dialog_hide(self.native())
}

// Run is a wrapper around gtk_native_dialog_run().
func (self NativeDialog) Run() int {
	ret0 := C.gtk_native_dialog_run(self.native())
	return int(ret0)
}

// SetModal is a wrapper around gtk_native_dialog_set_modal().
func (self NativeDialog) SetModal(modal bool) {
	C.gtk_native_dialog_set_modal(self.native(), C.gboolean(util.Bool2Int(modal)) /*go:.util*/)
}

// SetTitle is a wrapper around gtk_native_dialog_set_title().
func (self NativeDialog) SetTitle(title string) {
	title0 := C.CString(title)
	C.gtk_native_dialog_set_title(self.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetTransientFor is a wrapper around gtk_native_dialog_set_transient_for().
func (self NativeDialog) SetTransientFor(parent Window) {
	C.gtk_native_dialog_set_transient_for(self.native(), parent.native())
}

// Show is a wrapper around gtk_native_dialog_show().
func (self NativeDialog) Show() {
	C.gtk_native_dialog_show(self.native())
}

// Object FileChooserWidget
type FileChooserWidget struct {
	Box
}

func (v FileChooserWidget) native() *C.GtkFileChooserWidget {
	return (*C.GtkFileChooserWidget)(v.Ptr)
}
func wrapFileChooserWidget(p *C.GtkFileChooserWidget) (v FileChooserWidget) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileChooserWidget(p unsafe.Pointer) (v FileChooserWidget) {
	v.Ptr = p
	return
}
func (v FileChooserWidget) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileChooserWidget(p unsafe.Pointer) interface{} {
	return WrapFileChooserWidget(p)
}
func (v FileChooserWidget) GetType() gobject.Type {
	return gobject.Type(C.gtk_file_chooser_widget_get_type())
}
func (v FileChooserWidget) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileChooserWidget(unsafe.Pointer(ptr)), nil
	}
}
func (v FileChooserWidget) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FileChooserWidget) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FileChooserWidget) FileChooser() FileChooser {
	return WrapFileChooser(v.Ptr)
}
func (v FileChooserWidget) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// FileChooserWidgetNew is a wrapper around gtk_file_chooser_widget_new().
func FileChooserWidgetNew(action FileChooserAction) Widget {
	ret0 := C.gtk_file_chooser_widget_new(C.GtkFileChooserAction(action))
	return wrapWidget(ret0)
}

// Object Fixed
type Fixed struct {
	Container
}

func (v Fixed) native() *C.GtkFixed {
	return (*C.GtkFixed)(v.Ptr)
}
func wrapFixed(p *C.GtkFixed) (v Fixed) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFixed(p unsafe.Pointer) (v Fixed) {
	v.Ptr = p
	return
}
func (v Fixed) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFixed(p unsafe.Pointer) interface{} {
	return WrapFixed(p)
}
func (v Fixed) GetType() gobject.Type {
	return gobject.Type(C.gtk_fixed_get_type())
}
func (v Fixed) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFixed(unsafe.Pointer(ptr)), nil
	}
}
func (v Fixed) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Fixed) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// FixedNew is a wrapper around gtk_fixed_new().
func FixedNew() Widget {
	ret0 := C.gtk_fixed_new()
	return wrapWidget(ret0)
}

// Move is a wrapper around gtk_fixed_move().
func (fixed Fixed) Move(widget Widget, x int, y int) {
	C.gtk_fixed_move(fixed.native(), widget.native(), C.gint(x), C.gint(y))
}

// Put is a wrapper around gtk_fixed_put().
func (fixed Fixed) Put(widget Widget, x int, y int) {
	C.gtk_fixed_put(fixed.native(), widget.native(), C.gint(x), C.gint(y))
}

// Object FlowBox
type FlowBox struct {
	Container
}

func (v FlowBox) native() *C.GtkFlowBox {
	return (*C.GtkFlowBox)(v.Ptr)
}
func wrapFlowBox(p *C.GtkFlowBox) (v FlowBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFlowBox(p unsafe.Pointer) (v FlowBox) {
	v.Ptr = p
	return
}
func (v FlowBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFlowBox(p unsafe.Pointer) interface{} {
	return WrapFlowBox(p)
}
func (v FlowBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_flow_box_get_type())
}
func (v FlowBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFlowBox(unsafe.Pointer(ptr)), nil
	}
}
func (v FlowBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FlowBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FlowBox) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// FlowBoxNew is a wrapper around gtk_flow_box_new().
func FlowBoxNew() Widget {
	ret0 := C.gtk_flow_box_new()
	return wrapWidget(ret0)
}

// GetActivateOnSingleClick is a wrapper around gtk_flow_box_get_activate_on_single_click().
func (box FlowBox) GetActivateOnSingleClick() bool {
	ret0 := C.gtk_flow_box_get_activate_on_single_click(box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChildAtIndex is a wrapper around gtk_flow_box_get_child_at_index().
func (box FlowBox) GetChildAtIndex(idx int) FlowBoxChild {
	ret0 := C.gtk_flow_box_get_child_at_index(box.native(), C.gint(idx))
	return wrapFlowBoxChild(ret0)
}

// GetChildAtPos is a wrapper around gtk_flow_box_get_child_at_pos().
func (box FlowBox) GetChildAtPos(x int, y int) FlowBoxChild {
	ret0 := C.gtk_flow_box_get_child_at_pos(box.native(), C.gint(x), C.gint(y))
	return wrapFlowBoxChild(ret0)
}

// GetColumnSpacing is a wrapper around gtk_flow_box_get_column_spacing().
func (box FlowBox) GetColumnSpacing() uint {
	ret0 := C.gtk_flow_box_get_column_spacing(box.native())
	return uint(ret0)
}

// GetHomogeneous is a wrapper around gtk_flow_box_get_homogeneous().
func (box FlowBox) GetHomogeneous() bool {
	ret0 := C.gtk_flow_box_get_homogeneous(box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMaxChildrenPerLine is a wrapper around gtk_flow_box_get_max_children_per_line().
func (box FlowBox) GetMaxChildrenPerLine() uint {
	ret0 := C.gtk_flow_box_get_max_children_per_line(box.native())
	return uint(ret0)
}

// GetMinChildrenPerLine is a wrapper around gtk_flow_box_get_min_children_per_line().
func (box FlowBox) GetMinChildrenPerLine() uint {
	ret0 := C.gtk_flow_box_get_min_children_per_line(box.native())
	return uint(ret0)
}

// GetRowSpacing is a wrapper around gtk_flow_box_get_row_spacing().
func (box FlowBox) GetRowSpacing() uint {
	ret0 := C.gtk_flow_box_get_row_spacing(box.native())
	return uint(ret0)
}

// GetSelectedChildren is a wrapper around gtk_flow_box_get_selected_children().
func (box FlowBox) GetSelectedChildren() glib.List {
	ret0 := C.gtk_flow_box_get_selected_children(box.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapFlowBoxChild(p) }) /*gir:GLib*/
}

// GetSelectionMode is a wrapper around gtk_flow_box_get_selection_mode().
func (box FlowBox) GetSelectionMode() SelectionMode {
	ret0 := C.gtk_flow_box_get_selection_mode(box.native())
	return SelectionMode(ret0)
}

// Insert is a wrapper around gtk_flow_box_insert().
func (box FlowBox) Insert(widget Widget, position int) {
	C.gtk_flow_box_insert(box.native(), widget.native(), C.gint(position))
}

// InvalidateFilter is a wrapper around gtk_flow_box_invalidate_filter().
func (box FlowBox) InvalidateFilter() {
	C.gtk_flow_box_invalidate_filter(box.native())
}

// InvalidateSort is a wrapper around gtk_flow_box_invalidate_sort().
func (box FlowBox) InvalidateSort() {
	C.gtk_flow_box_invalidate_sort(box.native())
}

// SelectAll is a wrapper around gtk_flow_box_select_all().
func (box FlowBox) SelectAll() {
	C.gtk_flow_box_select_all(box.native())
}

// SelectChild is a wrapper around gtk_flow_box_select_child().
func (box FlowBox) SelectChild(child FlowBoxChild) {
	C.gtk_flow_box_select_child(box.native(), child.native())
}

// SetActivateOnSingleClick is a wrapper around gtk_flow_box_set_activate_on_single_click().
func (box FlowBox) SetActivateOnSingleClick(single bool) {
	C.gtk_flow_box_set_activate_on_single_click(box.native(), C.gboolean(util.Bool2Int(single)) /*go:.util*/)
}

// SetColumnSpacing is a wrapper around gtk_flow_box_set_column_spacing().
func (box FlowBox) SetColumnSpacing(spacing uint) {
	C.gtk_flow_box_set_column_spacing(box.native(), C.guint(spacing))
}

// SetHadjustment is a wrapper around gtk_flow_box_set_hadjustment().
func (box FlowBox) SetHadjustment(adjustment Adjustment) {
	C.gtk_flow_box_set_hadjustment(box.native(), adjustment.native())
}

// SetHomogeneous is a wrapper around gtk_flow_box_set_homogeneous().
func (box FlowBox) SetHomogeneous(homogeneous bool) {
	C.gtk_flow_box_set_homogeneous(box.native(), C.gboolean(util.Bool2Int(homogeneous)) /*go:.util*/)
}

// SetMaxChildrenPerLine is a wrapper around gtk_flow_box_set_max_children_per_line().
func (box FlowBox) SetMaxChildrenPerLine(n_children uint) {
	C.gtk_flow_box_set_max_children_per_line(box.native(), C.guint(n_children))
}

// SetMinChildrenPerLine is a wrapper around gtk_flow_box_set_min_children_per_line().
func (box FlowBox) SetMinChildrenPerLine(n_children uint) {
	C.gtk_flow_box_set_min_children_per_line(box.native(), C.guint(n_children))
}

// SetRowSpacing is a wrapper around gtk_flow_box_set_row_spacing().
func (box FlowBox) SetRowSpacing(spacing uint) {
	C.gtk_flow_box_set_row_spacing(box.native(), C.guint(spacing))
}

// SetSelectionMode is a wrapper around gtk_flow_box_set_selection_mode().
func (box FlowBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_flow_box_set_selection_mode(box.native(), C.GtkSelectionMode(mode))
}

// SetVadjustment is a wrapper around gtk_flow_box_set_vadjustment().
func (box FlowBox) SetVadjustment(adjustment Adjustment) {
	C.gtk_flow_box_set_vadjustment(box.native(), adjustment.native())
}

// UnselectAll is a wrapper around gtk_flow_box_unselect_all().
func (box FlowBox) UnselectAll() {
	C.gtk_flow_box_unselect_all(box.native())
}

// UnselectChild is a wrapper around gtk_flow_box_unselect_child().
func (box FlowBox) UnselectChild(child FlowBoxChild) {
	C.gtk_flow_box_unselect_child(box.native(), child.native())
}

// Object FlowBoxChild
type FlowBoxChild struct {
	Bin
}

func (v FlowBoxChild) native() *C.GtkFlowBoxChild {
	return (*C.GtkFlowBoxChild)(v.Ptr)
}
func wrapFlowBoxChild(p *C.GtkFlowBoxChild) (v FlowBoxChild) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFlowBoxChild(p unsafe.Pointer) (v FlowBoxChild) {
	v.Ptr = p
	return
}
func (v FlowBoxChild) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFlowBoxChild(p unsafe.Pointer) interface{} {
	return WrapFlowBoxChild(p)
}
func (v FlowBoxChild) GetType() gobject.Type {
	return gobject.Type(C.gtk_flow_box_child_get_type())
}
func (v FlowBoxChild) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFlowBoxChild(unsafe.Pointer(ptr)), nil
	}
}
func (v FlowBoxChild) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FlowBoxChild) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// FlowBoxChildNew is a wrapper around gtk_flow_box_child_new().
func FlowBoxChildNew() Widget {
	ret0 := C.gtk_flow_box_child_new()
	return wrapWidget(ret0)
}

// Changed is a wrapper around gtk_flow_box_child_changed().
func (child FlowBoxChild) Changed() {
	C.gtk_flow_box_child_changed(child.native())
}

// GetIndex is a wrapper around gtk_flow_box_child_get_index().
func (child FlowBoxChild) GetIndex() int {
	ret0 := C.gtk_flow_box_child_get_index(child.native())
	return int(ret0)
}

// IsSelected is a wrapper around gtk_flow_box_child_is_selected().
func (child FlowBoxChild) IsSelected() bool {
	ret0 := C.gtk_flow_box_child_is_selected(child.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object FlowBoxAccessible
type FlowBoxAccessible struct {
	ContainerAccessible
}

func (v FlowBoxAccessible) native() *C.GtkFlowBoxAccessible {
	return (*C.GtkFlowBoxAccessible)(v.Ptr)
}
func wrapFlowBoxAccessible(p *C.GtkFlowBoxAccessible) (v FlowBoxAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFlowBoxAccessible(p unsafe.Pointer) (v FlowBoxAccessible) {
	v.Ptr = p
	return
}
func (v FlowBoxAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFlowBoxAccessible(p unsafe.Pointer) interface{} {
	return WrapFlowBoxAccessible(p)
}
func (v FlowBoxAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_flow_box_accessible_get_type())
}
func (v FlowBoxAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFlowBoxAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v FlowBoxAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v FlowBoxAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object FlowBoxChildAccessible
type FlowBoxChildAccessible struct {
	ContainerAccessible
}

func (v FlowBoxChildAccessible) native() *C.GtkFlowBoxChildAccessible {
	return (*C.GtkFlowBoxChildAccessible)(v.Ptr)
}
func wrapFlowBoxChildAccessible(p *C.GtkFlowBoxChildAccessible) (v FlowBoxChildAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFlowBoxChildAccessible(p unsafe.Pointer) (v FlowBoxChildAccessible) {
	v.Ptr = p
	return
}
func (v FlowBoxChildAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFlowBoxChildAccessible(p unsafe.Pointer) interface{} {
	return WrapFlowBoxChildAccessible(p)
}
func (v FlowBoxChildAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_flow_box_child_accessible_get_type())
}
func (v FlowBoxChildAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFlowBoxChildAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v FlowBoxChildAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object FontButton
type FontButton struct {
	Button
}

func (v FontButton) native() *C.GtkFontButton {
	return (*C.GtkFontButton)(v.Ptr)
}
func wrapFontButton(p *C.GtkFontButton) (v FontButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontButton(p unsafe.Pointer) (v FontButton) {
	v.Ptr = p
	return
}
func (v FontButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontButton(p unsafe.Pointer) interface{} {
	return WrapFontButton(p)
}
func (v FontButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_font_button_get_type())
}
func (v FontButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontButton(unsafe.Pointer(ptr)), nil
	}
}
func (v FontButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FontButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v FontButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v FontButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FontButton) FontChooser() FontChooser {
	return WrapFontChooser(v.Ptr)
}

// FontButtonNew is a wrapper around gtk_font_button_new().
func FontButtonNew() Widget {
	ret0 := C.gtk_font_button_new()
	return wrapWidget(ret0)
}

// FontButtonNewWithFont is a wrapper around gtk_font_button_new_with_font().
func FontButtonNewWithFont(fontname string) Widget {
	fontname0 := (*C.gchar)(C.CString(fontname))
	ret0 := C.gtk_font_button_new_with_font(fontname0)
	C.free(unsafe.Pointer(fontname0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetFontName is a wrapper around gtk_font_button_get_font_name().
func (font_button FontButton) GetFontName() string {
	ret0 := C.gtk_font_button_get_font_name(font_button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetShowSize is a wrapper around gtk_font_button_get_show_size().
func (font_button FontButton) GetShowSize() bool {
	ret0 := C.gtk_font_button_get_show_size(font_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowStyle is a wrapper around gtk_font_button_get_show_style().
func (font_button FontButton) GetShowStyle() bool {
	ret0 := C.gtk_font_button_get_show_style(font_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTitle is a wrapper around gtk_font_button_get_title().
func (font_button FontButton) GetTitle() string {
	ret0 := C.gtk_font_button_get_title(font_button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUseFont is a wrapper around gtk_font_button_get_use_font().
func (font_button FontButton) GetUseFont() bool {
	ret0 := C.gtk_font_button_get_use_font(font_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUseSize is a wrapper around gtk_font_button_get_use_size().
func (font_button FontButton) GetUseSize() bool {
	ret0 := C.gtk_font_button_get_use_size(font_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetFontName is a wrapper around gtk_font_button_set_font_name().
func (font_button FontButton) SetFontName(fontname string) bool {
	fontname0 := (*C.gchar)(C.CString(fontname))
	ret0 := C.gtk_font_button_set_font_name(font_button.native(), fontname0)
	C.free(unsafe.Pointer(fontname0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))   /*go:.util*/
}

// SetShowSize is a wrapper around gtk_font_button_set_show_size().
func (font_button FontButton) SetShowSize(show_size bool) {
	C.gtk_font_button_set_show_size(font_button.native(), C.gboolean(util.Bool2Int(show_size)) /*go:.util*/)
}

// SetShowStyle is a wrapper around gtk_font_button_set_show_style().
func (font_button FontButton) SetShowStyle(show_style bool) {
	C.gtk_font_button_set_show_style(font_button.native(), C.gboolean(util.Bool2Int(show_style)) /*go:.util*/)
}

// SetTitle is a wrapper around gtk_font_button_set_title().
func (font_button FontButton) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.gtk_font_button_set_title(font_button.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetUseFont is a wrapper around gtk_font_button_set_use_font().
func (font_button FontButton) SetUseFont(use_font bool) {
	C.gtk_font_button_set_use_font(font_button.native(), C.gboolean(util.Bool2Int(use_font)) /*go:.util*/)
}

// SetUseSize is a wrapper around gtk_font_button_set_use_size().
func (font_button FontButton) SetUseSize(use_size bool) {
	C.gtk_font_button_set_use_size(font_button.native(), C.gboolean(util.Bool2Int(use_size)) /*go:.util*/)
}

// Object FontChooserDialog
type FontChooserDialog struct {
	Dialog
}

func (v FontChooserDialog) native() *C.GtkFontChooserDialog {
	return (*C.GtkFontChooserDialog)(v.Ptr)
}
func wrapFontChooserDialog(p *C.GtkFontChooserDialog) (v FontChooserDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontChooserDialog(p unsafe.Pointer) (v FontChooserDialog) {
	v.Ptr = p
	return
}
func (v FontChooserDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontChooserDialog(p unsafe.Pointer) interface{} {
	return WrapFontChooserDialog(p)
}
func (v FontChooserDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_font_chooser_dialog_get_type())
}
func (v FontChooserDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontChooserDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v FontChooserDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FontChooserDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FontChooserDialog) FontChooser() FontChooser {
	return WrapFontChooser(v.Ptr)
}

// FontChooserDialogNew is a wrapper around gtk_font_chooser_dialog_new().
func FontChooserDialogNew(title string, parent Window) Widget {
	title0 := (*C.gchar)(C.CString(title))
	ret0 := C.gtk_font_chooser_dialog_new(title0, parent.native())
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// Object FontChooserWidget
type FontChooserWidget struct {
	Box
}

func (v FontChooserWidget) native() *C.GtkFontChooserWidget {
	return (*C.GtkFontChooserWidget)(v.Ptr)
}
func wrapFontChooserWidget(p *C.GtkFontChooserWidget) (v FontChooserWidget) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontChooserWidget(p unsafe.Pointer) (v FontChooserWidget) {
	v.Ptr = p
	return
}
func (v FontChooserWidget) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontChooserWidget(p unsafe.Pointer) interface{} {
	return WrapFontChooserWidget(p)
}
func (v FontChooserWidget) GetType() gobject.Type {
	return gobject.Type(C.gtk_font_chooser_widget_get_type())
}
func (v FontChooserWidget) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontChooserWidget(unsafe.Pointer(ptr)), nil
	}
}
func (v FontChooserWidget) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FontChooserWidget) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FontChooserWidget) FontChooser() FontChooser {
	return WrapFontChooser(v.Ptr)
}
func (v FontChooserWidget) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// FontChooserWidgetNew is a wrapper around gtk_font_chooser_widget_new().
func FontChooserWidgetNew() Widget {
	ret0 := C.gtk_font_chooser_widget_new()
	return wrapWidget(ret0)
}

// Object FontSelection
type FontSelection struct {
	Box
}

func (v FontSelection) native() *C.GtkFontSelection {
	return (*C.GtkFontSelection)(v.Ptr)
}
func wrapFontSelection(p *C.GtkFontSelection) (v FontSelection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontSelection(p unsafe.Pointer) (v FontSelection) {
	v.Ptr = p
	return
}
func (v FontSelection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontSelection(p unsafe.Pointer) interface{} {
	return WrapFontSelection(p)
}
func (v FontSelection) GetType() gobject.Type {
	return gobject.Type(C.gtk_font_selection_get_type())
}
func (v FontSelection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontSelection(unsafe.Pointer(ptr)), nil
	}
}
func (v FontSelection) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FontSelection) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v FontSelection) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object FontSelectionDialog
type FontSelectionDialog struct {
	Dialog
}

func (v FontSelectionDialog) native() *C.GtkFontSelectionDialog {
	return (*C.GtkFontSelectionDialog)(v.Ptr)
}
func wrapFontSelectionDialog(p *C.GtkFontSelectionDialog) (v FontSelectionDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontSelectionDialog(p unsafe.Pointer) (v FontSelectionDialog) {
	v.Ptr = p
	return
}
func (v FontSelectionDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontSelectionDialog(p unsafe.Pointer) interface{} {
	return WrapFontSelectionDialog(p)
}
func (v FontSelectionDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_font_selection_dialog_get_type())
}
func (v FontSelectionDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontSelectionDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v FontSelectionDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v FontSelectionDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object FrameAccessible
type FrameAccessible struct {
	ContainerAccessible
}

func (v FrameAccessible) native() *C.GtkFrameAccessible {
	return (*C.GtkFrameAccessible)(v.Ptr)
}
func wrapFrameAccessible(p *C.GtkFrameAccessible) (v FrameAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFrameAccessible(p unsafe.Pointer) (v FrameAccessible) {
	v.Ptr = p
	return
}
func (v FrameAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFrameAccessible(p unsafe.Pointer) interface{} {
	return WrapFrameAccessible(p)
}
func (v FrameAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_frame_accessible_get_type())
}
func (v FrameAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFrameAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v FrameAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object GLArea
type GLArea struct {
	Widget
}

func (v GLArea) native() *C.GtkGLArea {
	return (*C.GtkGLArea)(v.Ptr)
}
func wrapGLArea(p *C.GtkGLArea) (v GLArea) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGLArea(p unsafe.Pointer) (v GLArea) {
	v.Ptr = p
	return
}
func (v GLArea) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGLArea(p unsafe.Pointer) interface{} {
	return WrapGLArea(p)
}
func (v GLArea) GetType() gobject.Type {
	return gobject.Type(C.gtk_gl_area_get_type())
}
func (v GLArea) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGLArea(unsafe.Pointer(ptr)), nil
	}
}
func (v GLArea) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v GLArea) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// GLAreaNew is a wrapper around gtk_gl_area_new().
func GLAreaNew() Widget {
	ret0 := C.gtk_gl_area_new()
	return wrapWidget(ret0)
}

// AttachBuffers is a wrapper around gtk_gl_area_attach_buffers().
func (area GLArea) AttachBuffers() {
	C.gtk_gl_area_attach_buffers(area.native())
}

// GetAutoRender is a wrapper around gtk_gl_area_get_auto_render().
func (area GLArea) GetAutoRender() bool {
	ret0 := C.gtk_gl_area_get_auto_render(area.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetContext is a wrapper around gtk_gl_area_get_context().
func (area GLArea) GetContext() gdk.GLContext {
	ret0 := C.gtk_gl_area_get_context(area.native())
	return gdk.WrapGLContext(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetError is a wrapper around gtk_gl_area_get_error().
func (area GLArea) GetError() glib.Error {
	ret0 := C.gtk_gl_area_get_error(area.native())
	return glib.WrapError(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetHasAlpha is a wrapper around gtk_gl_area_get_has_alpha().
func (area GLArea) GetHasAlpha() bool {
	ret0 := C.gtk_gl_area_get_has_alpha(area.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHasDepthBuffer is a wrapper around gtk_gl_area_get_has_depth_buffer().
func (area GLArea) GetHasDepthBuffer() bool {
	ret0 := C.gtk_gl_area_get_has_depth_buffer(area.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHasStencilBuffer is a wrapper around gtk_gl_area_get_has_stencil_buffer().
func (area GLArea) GetHasStencilBuffer() bool {
	ret0 := C.gtk_gl_area_get_has_stencil_buffer(area.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRequiredVersion is a wrapper around gtk_gl_area_get_required_version().
func (area GLArea) GetRequiredVersion() (int, int) {
	var major0 C.gint
	var minor0 C.gint
	C.gtk_gl_area_get_required_version(area.native(), &major0, &minor0)
	return int(major0), int(minor0)
}

// GetUseEs is a wrapper around gtk_gl_area_get_use_es().
func (area GLArea) GetUseEs() bool {
	ret0 := C.gtk_gl_area_get_use_es(area.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MakeCurrent is a wrapper around gtk_gl_area_make_current().
func (area GLArea) MakeCurrent() {
	C.gtk_gl_area_make_current(area.native())
}

// QueueRender is a wrapper around gtk_gl_area_queue_render().
func (area GLArea) QueueRender() {
	C.gtk_gl_area_queue_render(area.native())
}

// SetAutoRender is a wrapper around gtk_gl_area_set_auto_render().
func (area GLArea) SetAutoRender(auto_render bool) {
	C.gtk_gl_area_set_auto_render(area.native(), C.gboolean(util.Bool2Int(auto_render)) /*go:.util*/)
}

// SetError is a wrapper around gtk_gl_area_set_error().
func (area GLArea) SetError(error glib.Error) {
	C.gtk_gl_area_set_error(area.native(), (*C.GError)(error.Ptr))
}

// SetHasAlpha is a wrapper around gtk_gl_area_set_has_alpha().
func (area GLArea) SetHasAlpha(has_alpha bool) {
	C.gtk_gl_area_set_has_alpha(area.native(), C.gboolean(util.Bool2Int(has_alpha)) /*go:.util*/)
}

// SetHasDepthBuffer is a wrapper around gtk_gl_area_set_has_depth_buffer().
func (area GLArea) SetHasDepthBuffer(has_depth_buffer bool) {
	C.gtk_gl_area_set_has_depth_buffer(area.native(), C.gboolean(util.Bool2Int(has_depth_buffer)) /*go:.util*/)
}

// SetHasStencilBuffer is a wrapper around gtk_gl_area_set_has_stencil_buffer().
func (area GLArea) SetHasStencilBuffer(has_stencil_buffer bool) {
	C.gtk_gl_area_set_has_stencil_buffer(area.native(), C.gboolean(util.Bool2Int(has_stencil_buffer)) /*go:.util*/)
}

// SetRequiredVersion is a wrapper around gtk_gl_area_set_required_version().
func (area GLArea) SetRequiredVersion(major int, minor int) {
	C.gtk_gl_area_set_required_version(area.native(), C.gint(major), C.gint(minor))
}

// SetUseEs is a wrapper around gtk_gl_area_set_use_es().
func (area GLArea) SetUseEs(use_es bool) {
	C.gtk_gl_area_set_use_es(area.native(), C.gboolean(util.Bool2Int(use_es)) /*go:.util*/)
}

// Object Gesture
type Gesture struct {
	EventController
}

func (v Gesture) native() *C.GtkGesture {
	return (*C.GtkGesture)(v.Ptr)
}
func wrapGesture(p *C.GtkGesture) (v Gesture) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGesture(p unsafe.Pointer) (v Gesture) {
	v.Ptr = p
	return
}
func (v Gesture) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGesture(p unsafe.Pointer) interface{} {
	return WrapGesture(p)
}
func (v Gesture) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_get_type())
}
func (v Gesture) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGesture(unsafe.Pointer(ptr)), nil
	}
}

// GetBoundingBoxCenter is a wrapper around gtk_gesture_get_bounding_box_center().
func (gesture Gesture) GetBoundingBoxCenter() (bool, float64, float64) {
	var x0 C.gdouble
	var y0 C.gdouble
	ret0 := C.gtk_gesture_get_bounding_box_center(gesture.native(), &x0, &y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(x0), float64(y0)
}

// GetDevice is a wrapper around gtk_gesture_get_device().
func (gesture Gesture) GetDevice() gdk.Device {
	ret0 := C.gtk_gesture_get_device(gesture.native())
	return gdk.WrapDevice(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetGroup is a wrapper around gtk_gesture_get_group().
func (gesture Gesture) GetGroup() glib.List {
	ret0 := C.gtk_gesture_get_group(gesture.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapGesture(p) }) /*gir:GLib*/
}

// GetLastUpdatedSequence is a wrapper around gtk_gesture_get_last_updated_sequence().
func (gesture Gesture) GetLastUpdatedSequence() gdk.EventSequence {
	ret0 := C.gtk_gesture_get_last_updated_sequence(gesture.native())
	return gdk.WrapEventSequence(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetPoint is a wrapper around gtk_gesture_get_point().
func (gesture Gesture) GetPoint(sequence gdk.EventSequence) (bool, float64, float64) {
	var x0 C.gdouble
	var y0 C.gdouble
	ret0 := C.gtk_gesture_get_point(gesture.native(), (*C.GdkEventSequence)(sequence.Ptr), &x0, &y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(x0), float64(y0)
}

// GetSequenceState is a wrapper around gtk_gesture_get_sequence_state().
func (gesture Gesture) GetSequenceState(sequence gdk.EventSequence) EventSequenceState {
	ret0 := C.gtk_gesture_get_sequence_state(gesture.native(), (*C.GdkEventSequence)(sequence.Ptr))
	return EventSequenceState(ret0)
}

// GetWindow is a wrapper around gtk_gesture_get_window().
func (gesture Gesture) GetWindow() gdk.Window {
	ret0 := C.gtk_gesture_get_window(gesture.native())
	return gdk.WrapWindow(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// Group is a wrapper around gtk_gesture_group().
func (group_gesture Gesture) Group(gesture Gesture) {
	C.gtk_gesture_group(group_gesture.native(), gesture.native())
}

// HandlesSequence is a wrapper around gtk_gesture_handles_sequence().
func (gesture Gesture) HandlesSequence(sequence gdk.EventSequence) bool {
	ret0 := C.gtk_gesture_handles_sequence(gesture.native(), (*C.GdkEventSequence)(sequence.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsActive is a wrapper around gtk_gesture_is_active().
func (gesture Gesture) IsActive() bool {
	ret0 := C.gtk_gesture_is_active(gesture.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsGroupedWith is a wrapper around gtk_gesture_is_grouped_with().
func (gesture Gesture) IsGroupedWith(other Gesture) bool {
	ret0 := C.gtk_gesture_is_grouped_with(gesture.native(), other.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsRecognized is a wrapper around gtk_gesture_is_recognized().
func (gesture Gesture) IsRecognized() bool {
	ret0 := C.gtk_gesture_is_recognized(gesture.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetSequenceState is a wrapper around gtk_gesture_set_sequence_state().
func (gesture Gesture) SetSequenceState(sequence gdk.EventSequence, state EventSequenceState) bool {
	ret0 := C.gtk_gesture_set_sequence_state(gesture.native(), (*C.GdkEventSequence)(sequence.Ptr), C.GtkEventSequenceState(state))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetState is a wrapper around gtk_gesture_set_state().
func (gesture Gesture) SetState(state EventSequenceState) bool {
	ret0 := C.gtk_gesture_set_state(gesture.native(), C.GtkEventSequenceState(state))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetWindow is a wrapper around gtk_gesture_set_window().
func (gesture Gesture) SetWindow(window gdk.Window) {
	C.gtk_gesture_set_window(gesture.native(), (*C.GdkWindow)(window.Ptr))
}

// Ungroup is a wrapper around gtk_gesture_ungroup().
func (gesture Gesture) Ungroup() {
	C.gtk_gesture_ungroup(gesture.native())
}

// Object GestureDrag
type GestureDrag struct {
	GestureSingle
}

func (v GestureDrag) native() *C.GtkGestureDrag {
	return (*C.GtkGestureDrag)(v.Ptr)
}
func wrapGestureDrag(p *C.GtkGestureDrag) (v GestureDrag) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGestureDrag(p unsafe.Pointer) (v GestureDrag) {
	v.Ptr = p
	return
}
func (v GestureDrag) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGestureDrag(p unsafe.Pointer) interface{} {
	return WrapGestureDrag(p)
}
func (v GestureDrag) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_drag_get_type())
}
func (v GestureDrag) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGestureDrag(unsafe.Pointer(ptr)), nil
	}
}

// GestureDragNew is a wrapper around gtk_gesture_drag_new().
func GestureDragNew(widget Widget) Gesture {
	ret0 := C.gtk_gesture_drag_new(widget.native())
	return wrapGesture(ret0)
}

// GetOffset is a wrapper around gtk_gesture_drag_get_offset().
func (gesture GestureDrag) GetOffset() (bool, float64, float64) {
	var x0 C.gdouble
	var y0 C.gdouble
	ret0 := C.gtk_gesture_drag_get_offset(gesture.native(), &x0, &y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(x0), float64(y0)
}

// GetStartPoint is a wrapper around gtk_gesture_drag_get_start_point().
func (gesture GestureDrag) GetStartPoint() (bool, float64, float64) {
	var x0 C.gdouble
	var y0 C.gdouble
	ret0 := C.gtk_gesture_drag_get_start_point(gesture.native(), &x0, &y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(x0), float64(y0)
}

// Object GestureSingle
type GestureSingle struct {
	Gesture
}

func (v GestureSingle) native() *C.GtkGestureSingle {
	return (*C.GtkGestureSingle)(v.Ptr)
}
func wrapGestureSingle(p *C.GtkGestureSingle) (v GestureSingle) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGestureSingle(p unsafe.Pointer) (v GestureSingle) {
	v.Ptr = p
	return
}
func (v GestureSingle) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGestureSingle(p unsafe.Pointer) interface{} {
	return WrapGestureSingle(p)
}
func (v GestureSingle) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_single_get_type())
}
func (v GestureSingle) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGestureSingle(unsafe.Pointer(ptr)), nil
	}
}

// GetButton is a wrapper around gtk_gesture_single_get_button().
func (gesture GestureSingle) GetButton() uint {
	ret0 := C.gtk_gesture_single_get_button(gesture.native())
	return uint(ret0)
}

// GetCurrentButton is a wrapper around gtk_gesture_single_get_current_button().
func (gesture GestureSingle) GetCurrentButton() uint {
	ret0 := C.gtk_gesture_single_get_current_button(gesture.native())
	return uint(ret0)
}

// GetCurrentSequence is a wrapper around gtk_gesture_single_get_current_sequence().
func (gesture GestureSingle) GetCurrentSequence() gdk.EventSequence {
	ret0 := C.gtk_gesture_single_get_current_sequence(gesture.native())
	return gdk.WrapEventSequence(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetExclusive is a wrapper around gtk_gesture_single_get_exclusive().
func (gesture GestureSingle) GetExclusive() bool {
	ret0 := C.gtk_gesture_single_get_exclusive(gesture.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTouchOnly is a wrapper around gtk_gesture_single_get_touch_only().
func (gesture GestureSingle) GetTouchOnly() bool {
	ret0 := C.gtk_gesture_single_get_touch_only(gesture.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetButton is a wrapper around gtk_gesture_single_set_button().
func (gesture GestureSingle) SetButton(button uint) {
	C.gtk_gesture_single_set_button(gesture.native(), C.guint(button))
}

// SetExclusive is a wrapper around gtk_gesture_single_set_exclusive().
func (gesture GestureSingle) SetExclusive(exclusive bool) {
	C.gtk_gesture_single_set_exclusive(gesture.native(), C.gboolean(util.Bool2Int(exclusive)) /*go:.util*/)
}

// SetTouchOnly is a wrapper around gtk_gesture_single_set_touch_only().
func (gesture GestureSingle) SetTouchOnly(touch_only bool) {
	C.gtk_gesture_single_set_touch_only(gesture.native(), C.gboolean(util.Bool2Int(touch_only)) /*go:.util*/)
}

// Object GestureLongPress
type GestureLongPress struct {
	GestureSingle
}

func (v GestureLongPress) native() *C.GtkGestureLongPress {
	return (*C.GtkGestureLongPress)(v.Ptr)
}
func wrapGestureLongPress(p *C.GtkGestureLongPress) (v GestureLongPress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGestureLongPress(p unsafe.Pointer) (v GestureLongPress) {
	v.Ptr = p
	return
}
func (v GestureLongPress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGestureLongPress(p unsafe.Pointer) interface{} {
	return WrapGestureLongPress(p)
}
func (v GestureLongPress) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_long_press_get_type())
}
func (v GestureLongPress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGestureLongPress(unsafe.Pointer(ptr)), nil
	}
}

// GestureLongPressNew is a wrapper around gtk_gesture_long_press_new().
func GestureLongPressNew(widget Widget) Gesture {
	ret0 := C.gtk_gesture_long_press_new(widget.native())
	return wrapGesture(ret0)
}

// Object GestureMultiPress
type GestureMultiPress struct {
	GestureSingle
}

func (v GestureMultiPress) native() *C.GtkGestureMultiPress {
	return (*C.GtkGestureMultiPress)(v.Ptr)
}
func wrapGestureMultiPress(p *C.GtkGestureMultiPress) (v GestureMultiPress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGestureMultiPress(p unsafe.Pointer) (v GestureMultiPress) {
	v.Ptr = p
	return
}
func (v GestureMultiPress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGestureMultiPress(p unsafe.Pointer) interface{} {
	return WrapGestureMultiPress(p)
}
func (v GestureMultiPress) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_multi_press_get_type())
}
func (v GestureMultiPress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGestureMultiPress(unsafe.Pointer(ptr)), nil
	}
}

// GestureMultiPressNew is a wrapper around gtk_gesture_multi_press_new().
func GestureMultiPressNew(widget Widget) Gesture {
	ret0 := C.gtk_gesture_multi_press_new(widget.native())
	return wrapGesture(ret0)
}

// SetArea is a wrapper around gtk_gesture_multi_press_set_area().
func (gesture GestureMultiPress) SetArea(rect gdk.Rectangle) {
	C.gtk_gesture_multi_press_set_area(gesture.native(), (*C.GdkRectangle)(rect.Ptr))
}

// Object GesturePan
type GesturePan struct {
	GestureDrag
}

func (v GesturePan) native() *C.GtkGesturePan {
	return (*C.GtkGesturePan)(v.Ptr)
}
func wrapGesturePan(p *C.GtkGesturePan) (v GesturePan) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGesturePan(p unsafe.Pointer) (v GesturePan) {
	v.Ptr = p
	return
}
func (v GesturePan) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGesturePan(p unsafe.Pointer) interface{} {
	return WrapGesturePan(p)
}
func (v GesturePan) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_pan_get_type())
}
func (v GesturePan) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGesturePan(unsafe.Pointer(ptr)), nil
	}
}

// GesturePanNew is a wrapper around gtk_gesture_pan_new().
func GesturePanNew(widget Widget, orientation Orientation) Gesture {
	ret0 := C.gtk_gesture_pan_new(widget.native(), C.GtkOrientation(orientation))
	return wrapGesture(ret0)
}

// GetOrientation is a wrapper around gtk_gesture_pan_get_orientation().
func (gesture GesturePan) GetOrientation() Orientation {
	ret0 := C.gtk_gesture_pan_get_orientation(gesture.native())
	return Orientation(ret0)
}

// SetOrientation is a wrapper around gtk_gesture_pan_set_orientation().
func (gesture GesturePan) SetOrientation(orientation Orientation) {
	C.gtk_gesture_pan_set_orientation(gesture.native(), C.GtkOrientation(orientation))
}

// Object GestureRotate
type GestureRotate struct {
	Gesture
}

func (v GestureRotate) native() *C.GtkGestureRotate {
	return (*C.GtkGestureRotate)(v.Ptr)
}
func wrapGestureRotate(p *C.GtkGestureRotate) (v GestureRotate) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGestureRotate(p unsafe.Pointer) (v GestureRotate) {
	v.Ptr = p
	return
}
func (v GestureRotate) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGestureRotate(p unsafe.Pointer) interface{} {
	return WrapGestureRotate(p)
}
func (v GestureRotate) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_rotate_get_type())
}
func (v GestureRotate) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGestureRotate(unsafe.Pointer(ptr)), nil
	}
}

// GestureRotateNew is a wrapper around gtk_gesture_rotate_new().
func GestureRotateNew(widget Widget) Gesture {
	ret0 := C.gtk_gesture_rotate_new(widget.native())
	return wrapGesture(ret0)
}

// GetAngleDelta is a wrapper around gtk_gesture_rotate_get_angle_delta().
func (gesture GestureRotate) GetAngleDelta() float64 {
	ret0 := C.gtk_gesture_rotate_get_angle_delta(gesture.native())
	return float64(ret0)
}

// Object GestureSwipe
type GestureSwipe struct {
	GestureSingle
}

func (v GestureSwipe) native() *C.GtkGestureSwipe {
	return (*C.GtkGestureSwipe)(v.Ptr)
}
func wrapGestureSwipe(p *C.GtkGestureSwipe) (v GestureSwipe) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGestureSwipe(p unsafe.Pointer) (v GestureSwipe) {
	v.Ptr = p
	return
}
func (v GestureSwipe) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGestureSwipe(p unsafe.Pointer) interface{} {
	return WrapGestureSwipe(p)
}
func (v GestureSwipe) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_swipe_get_type())
}
func (v GestureSwipe) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGestureSwipe(unsafe.Pointer(ptr)), nil
	}
}

// GestureSwipeNew is a wrapper around gtk_gesture_swipe_new().
func GestureSwipeNew(widget Widget) Gesture {
	ret0 := C.gtk_gesture_swipe_new(widget.native())
	return wrapGesture(ret0)
}

// GetVelocity is a wrapper around gtk_gesture_swipe_get_velocity().
func (gesture GestureSwipe) GetVelocity() (bool, float64, float64) {
	var velocity_x0 C.gdouble
	var velocity_y0 C.gdouble
	ret0 := C.gtk_gesture_swipe_get_velocity(gesture.native(), &velocity_x0, &velocity_y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(velocity_x0), float64(velocity_y0)
}

// Object GestureZoom
type GestureZoom struct {
	Gesture
}

func (v GestureZoom) native() *C.GtkGestureZoom {
	return (*C.GtkGestureZoom)(v.Ptr)
}
func wrapGestureZoom(p *C.GtkGestureZoom) (v GestureZoom) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGestureZoom(p unsafe.Pointer) (v GestureZoom) {
	v.Ptr = p
	return
}
func (v GestureZoom) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGestureZoom(p unsafe.Pointer) interface{} {
	return WrapGestureZoom(p)
}
func (v GestureZoom) GetType() gobject.Type {
	return gobject.Type(C.gtk_gesture_zoom_get_type())
}
func (v GestureZoom) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGestureZoom(unsafe.Pointer(ptr)), nil
	}
}

// GestureZoomNew is a wrapper around gtk_gesture_zoom_new().
func GestureZoomNew(widget Widget) Gesture {
	ret0 := C.gtk_gesture_zoom_new(widget.native())
	return wrapGesture(ret0)
}

// GetScaleDelta is a wrapper around gtk_gesture_zoom_get_scale_delta().
func (gesture GestureZoom) GetScaleDelta() float64 {
	ret0 := C.gtk_gesture_zoom_get_scale_delta(gesture.native())
	return float64(ret0)
}

// Object Grid
type Grid struct {
	Container
}

func (v Grid) native() *C.GtkGrid {
	return (*C.GtkGrid)(v.Ptr)
}
func wrapGrid(p *C.GtkGrid) (v Grid) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGrid(p unsafe.Pointer) (v Grid) {
	v.Ptr = p
	return
}
func (v Grid) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGrid(p unsafe.Pointer) interface{} {
	return WrapGrid(p)
}
func (v Grid) GetType() gobject.Type {
	return gobject.Type(C.gtk_grid_get_type())
}
func (v Grid) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGrid(unsafe.Pointer(ptr)), nil
	}
}
func (v Grid) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Grid) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Grid) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// GridNew is a wrapper around gtk_grid_new().
func GridNew() Widget {
	ret0 := C.gtk_grid_new()
	return wrapWidget(ret0)
}

// Attach is a wrapper around gtk_grid_attach().
func (grid Grid) Attach(child Widget, left int, top int, width int, height int) {
	C.gtk_grid_attach(grid.native(), child.native(), C.gint(left), C.gint(top), C.gint(width), C.gint(height))
}

// AttachNextTo is a wrapper around gtk_grid_attach_next_to().
func (grid Grid) AttachNextTo(child Widget, sibling Widget, side PositionType, width int, height int) {
	C.gtk_grid_attach_next_to(grid.native(), child.native(), sibling.native(), C.GtkPositionType(side), C.gint(width), C.gint(height))
}

// GetChildAt is a wrapper around gtk_grid_get_child_at().
func (grid Grid) GetChildAt(left int, top int) Widget {
	ret0 := C.gtk_grid_get_child_at(grid.native(), C.gint(left), C.gint(top))
	return wrapWidget(ret0)
}

// GetColumnHomogeneous is a wrapper around gtk_grid_get_column_homogeneous().
func (grid Grid) GetColumnHomogeneous() bool {
	ret0 := C.gtk_grid_get_column_homogeneous(grid.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetColumnSpacing is a wrapper around gtk_grid_get_column_spacing().
func (grid Grid) GetColumnSpacing() uint {
	ret0 := C.gtk_grid_get_column_spacing(grid.native())
	return uint(ret0)
}

// GetRowBaselinePosition is a wrapper around gtk_grid_get_row_baseline_position().
func (grid Grid) GetRowBaselinePosition(row int) BaselinePosition {
	ret0 := C.gtk_grid_get_row_baseline_position(grid.native(), C.gint(row))
	return BaselinePosition(ret0)
}

// GetRowHomogeneous is a wrapper around gtk_grid_get_row_homogeneous().
func (grid Grid) GetRowHomogeneous() bool {
	ret0 := C.gtk_grid_get_row_homogeneous(grid.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRowSpacing is a wrapper around gtk_grid_get_row_spacing().
func (grid Grid) GetRowSpacing() uint {
	ret0 := C.gtk_grid_get_row_spacing(grid.native())
	return uint(ret0)
}

// InsertColumn is a wrapper around gtk_grid_insert_column().
func (grid Grid) InsertColumn(position int) {
	C.gtk_grid_insert_column(grid.native(), C.gint(position))
}

// InsertNextTo is a wrapper around gtk_grid_insert_next_to().
func (grid Grid) InsertNextTo(sibling Widget, side PositionType) {
	C.gtk_grid_insert_next_to(grid.native(), sibling.native(), C.GtkPositionType(side))
}

// InsertRow is a wrapper around gtk_grid_insert_row().
func (grid Grid) InsertRow(position int) {
	C.gtk_grid_insert_row(grid.native(), C.gint(position))
}

// RemoveColumn is a wrapper around gtk_grid_remove_column().
func (grid Grid) RemoveColumn(position int) {
	C.gtk_grid_remove_column(grid.native(), C.gint(position))
}

// RemoveRow is a wrapper around gtk_grid_remove_row().
func (grid Grid) RemoveRow(position int) {
	C.gtk_grid_remove_row(grid.native(), C.gint(position))
}

// SetBaselineRow is a wrapper around gtk_grid_set_baseline_row().
func (grid Grid) SetBaselineRow(row int) {
	C.gtk_grid_set_baseline_row(grid.native(), C.gint(row))
}

// SetColumnHomogeneous is a wrapper around gtk_grid_set_column_homogeneous().
func (grid Grid) SetColumnHomogeneous(homogeneous bool) {
	C.gtk_grid_set_column_homogeneous(grid.native(), C.gboolean(util.Bool2Int(homogeneous)) /*go:.util*/)
}

// SetColumnSpacing is a wrapper around gtk_grid_set_column_spacing().
func (grid Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(grid.native(), C.guint(spacing))
}

// SetRowBaselinePosition is a wrapper around gtk_grid_set_row_baseline_position().
func (grid Grid) SetRowBaselinePosition(row int, pos BaselinePosition) {
	C.gtk_grid_set_row_baseline_position(grid.native(), C.gint(row), C.GtkBaselinePosition(pos))
}

// SetRowHomogeneous is a wrapper around gtk_grid_set_row_homogeneous().
func (grid Grid) SetRowHomogeneous(homogeneous bool) {
	C.gtk_grid_set_row_homogeneous(grid.native(), C.gboolean(util.Bool2Int(homogeneous)) /*go:.util*/)
}

// SetRowSpacing is a wrapper around gtk_grid_set_row_spacing().
func (grid Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(grid.native(), C.guint(spacing))
}

// Object HBox
type HBox struct {
	Box
}

func (v HBox) native() *C.GtkHBox {
	return (*C.GtkHBox)(v.Ptr)
}
func wrapHBox(p *C.GtkHBox) (v HBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHBox(p unsafe.Pointer) (v HBox) {
	v.Ptr = p
	return
}
func (v HBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHBox(p unsafe.Pointer) interface{} {
	return WrapHBox(p)
}
func (v HBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_hbox_get_type())
}
func (v HBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHBox(unsafe.Pointer(ptr)), nil
	}
}
func (v HBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v HBox) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object HButtonBox
type HButtonBox struct {
	ButtonBox
}

func (v HButtonBox) native() *C.GtkHButtonBox {
	return (*C.GtkHButtonBox)(v.Ptr)
}
func wrapHButtonBox(p *C.GtkHButtonBox) (v HButtonBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHButtonBox(p unsafe.Pointer) (v HButtonBox) {
	v.Ptr = p
	return
}
func (v HButtonBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHButtonBox(p unsafe.Pointer) interface{} {
	return WrapHButtonBox(p)
}
func (v HButtonBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_hbutton_box_get_type())
}
func (v HButtonBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHButtonBox(unsafe.Pointer(ptr)), nil
	}
}
func (v HButtonBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HButtonBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v HButtonBox) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object HPaned
type HPaned struct {
	Paned
}

func (v HPaned) native() *C.GtkHPaned {
	return (*C.GtkHPaned)(v.Ptr)
}
func wrapHPaned(p *C.GtkHPaned) (v HPaned) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHPaned(p unsafe.Pointer) (v HPaned) {
	v.Ptr = p
	return
}
func (v HPaned) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHPaned(p unsafe.Pointer) interface{} {
	return WrapHPaned(p)
}
func (v HPaned) GetType() gobject.Type {
	return gobject.Type(C.gtk_hpaned_get_type())
}
func (v HPaned) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHPaned(unsafe.Pointer(ptr)), nil
	}
}
func (v HPaned) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HPaned) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v HPaned) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object HSV
type HSV struct {
	Widget
}

func (v HSV) native() *C.GtkHSV {
	return (*C.GtkHSV)(v.Ptr)
}
func wrapHSV(p *C.GtkHSV) (v HSV) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHSV(p unsafe.Pointer) (v HSV) {
	v.Ptr = p
	return
}
func (v HSV) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHSV(p unsafe.Pointer) interface{} {
	return WrapHSV(p)
}
func (v HSV) GetType() gobject.Type {
	return gobject.Type(C.gtk_hsv_get_type())
}
func (v HSV) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHSV(unsafe.Pointer(ptr)), nil
	}
}
func (v HSV) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HSV) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// HSVNew is a wrapper around gtk_hsv_new().
func HSVNew() Widget {
	ret0 := C.gtk_hsv_new()
	return wrapWidget(ret0)
}

// GetColor is a wrapper around gtk_hsv_get_color().
func (hsv HSV) GetColor() (float64, float64, float64) {
	var h0 C.gdouble
	var s0 C.gdouble
	var v0 C.gdouble
	C.gtk_hsv_get_color(hsv.native(), &h0, &s0, &v0)
	return float64(h0), float64(s0), float64(v0)
}

// GetMetrics is a wrapper around gtk_hsv_get_metrics().
func (hsv HSV) GetMetrics() (int, int) {
	var size0 C.gint
	var ring_width0 C.gint
	C.gtk_hsv_get_metrics(hsv.native(), &size0, &ring_width0)
	return int(size0), int(ring_width0)
}

// IsAdjusting is a wrapper around gtk_hsv_is_adjusting().
func (hsv HSV) IsAdjusting() bool {
	ret0 := C.gtk_hsv_is_adjusting(hsv.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetColor is a wrapper around gtk_hsv_set_color().
func (hsv HSV) SetColor(h float64, s float64, v float64) {
	C.gtk_hsv_set_color(hsv.native(), C.double(h), C.double(s), C.double(v))
}

// SetMetrics is a wrapper around gtk_hsv_set_metrics().
func (hsv HSV) SetMetrics(size int, ring_width int) {
	C.gtk_hsv_set_metrics(hsv.native(), C.gint(size), C.gint(ring_width))
}

// HSVToRgb is a wrapper around gtk_hsv_to_rgb().
func HSVToRgb(h float64, s float64, v float64) (float64, float64, float64) {
	var r0 C.gdouble
	var g0 C.gdouble
	var b0 C.gdouble
	C.gtk_hsv_to_rgb(C.gdouble(h), C.gdouble(s), C.gdouble(v), &r0, &g0, &b0)
	return float64(r0), float64(g0), float64(b0)
}

// Object Paned
type Paned struct {
	Container
}

func (v Paned) native() *C.GtkPaned {
	return (*C.GtkPaned)(v.Ptr)
}
func wrapPaned(p *C.GtkPaned) (v Paned) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPaned(p unsafe.Pointer) (v Paned) {
	v.Ptr = p
	return
}
func (v Paned) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPaned(p unsafe.Pointer) interface{} {
	return WrapPaned(p)
}
func (v Paned) GetType() gobject.Type {
	return gobject.Type(C.gtk_paned_get_type())
}
func (v Paned) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPaned(unsafe.Pointer(ptr)), nil
	}
}
func (v Paned) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Paned) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Paned) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// PanedNew is a wrapper around gtk_paned_new().
func PanedNew(orientation Orientation) Widget {
	ret0 := C.gtk_paned_new(C.GtkOrientation(orientation))
	return wrapWidget(ret0)
}

// Add1 is a wrapper around gtk_paned_add1().
func (paned Paned) Add1(child Widget) {
	C.gtk_paned_add1(paned.native(), child.native())
}

// Add2 is a wrapper around gtk_paned_add2().
func (paned Paned) Add2(child Widget) {
	C.gtk_paned_add2(paned.native(), child.native())
}

// GetChild1 is a wrapper around gtk_paned_get_child1().
func (paned Paned) GetChild1() Widget {
	ret0 := C.gtk_paned_get_child1(paned.native())
	return wrapWidget(ret0)
}

// GetChild2 is a wrapper around gtk_paned_get_child2().
func (paned Paned) GetChild2() Widget {
	ret0 := C.gtk_paned_get_child2(paned.native())
	return wrapWidget(ret0)
}

// GetHandleWindow is a wrapper around gtk_paned_get_handle_window().
func (paned Paned) GetHandleWindow() gdk.Window {
	ret0 := C.gtk_paned_get_handle_window(paned.native())
	return gdk.WrapWindow(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetPosition is a wrapper around gtk_paned_get_position().
func (paned Paned) GetPosition() int {
	ret0 := C.gtk_paned_get_position(paned.native())
	return int(ret0)
}

// GetWideHandle is a wrapper around gtk_paned_get_wide_handle().
func (paned Paned) GetWideHandle() bool {
	ret0 := C.gtk_paned_get_wide_handle(paned.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Pack1 is a wrapper around gtk_paned_pack1().
func (paned Paned) Pack1(child Widget, resize bool, shrink bool) {
	C.gtk_paned_pack1(paned.native(), child.native(), C.gboolean(util.Bool2Int(resize)) /*go:.util*/, C.gboolean(util.Bool2Int(shrink)) /*go:.util*/)
}

// Pack2 is a wrapper around gtk_paned_pack2().
func (paned Paned) Pack2(child Widget, resize bool, shrink bool) {
	C.gtk_paned_pack2(paned.native(), child.native(), C.gboolean(util.Bool2Int(resize)) /*go:.util*/, C.gboolean(util.Bool2Int(shrink)) /*go:.util*/)
}

// SetPosition is a wrapper around gtk_paned_set_position().
func (paned Paned) SetPosition(position int) {
	C.gtk_paned_set_position(paned.native(), C.gint(position))
}

// SetWideHandle is a wrapper around gtk_paned_set_wide_handle().
func (paned Paned) SetWideHandle(wide bool) {
	C.gtk_paned_set_wide_handle(paned.native(), C.gboolean(util.Bool2Int(wide)) /*go:.util*/)
}

// Object HScale
type HScale struct {
	Scale
}

func (v HScale) native() *C.GtkHScale {
	return (*C.GtkHScale)(v.Ptr)
}
func wrapHScale(p *C.GtkHScale) (v HScale) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHScale(p unsafe.Pointer) (v HScale) {
	v.Ptr = p
	return
}
func (v HScale) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHScale(p unsafe.Pointer) interface{} {
	return WrapHScale(p)
}
func (v HScale) GetType() gobject.Type {
	return gobject.Type(C.gtk_hscale_get_type())
}
func (v HScale) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHScale(unsafe.Pointer(ptr)), nil
	}
}
func (v HScale) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HScale) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v HScale) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object HScrollbar
type HScrollbar struct {
	Scrollbar
}

func (v HScrollbar) native() *C.GtkHScrollbar {
	return (*C.GtkHScrollbar)(v.Ptr)
}
func wrapHScrollbar(p *C.GtkHScrollbar) (v HScrollbar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHScrollbar(p unsafe.Pointer) (v HScrollbar) {
	v.Ptr = p
	return
}
func (v HScrollbar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHScrollbar(p unsafe.Pointer) interface{} {
	return WrapHScrollbar(p)
}
func (v HScrollbar) GetType() gobject.Type {
	return gobject.Type(C.gtk_hscrollbar_get_type())
}
func (v HScrollbar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHScrollbar(unsafe.Pointer(ptr)), nil
	}
}
func (v HScrollbar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HScrollbar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v HScrollbar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object HSeparator
type HSeparator struct {
	Separator
}

func (v HSeparator) native() *C.GtkHSeparator {
	return (*C.GtkHSeparator)(v.Ptr)
}
func wrapHSeparator(p *C.GtkHSeparator) (v HSeparator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHSeparator(p unsafe.Pointer) (v HSeparator) {
	v.Ptr = p
	return
}
func (v HSeparator) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHSeparator(p unsafe.Pointer) interface{} {
	return WrapHSeparator(p)
}
func (v HSeparator) GetType() gobject.Type {
	return gobject.Type(C.gtk_hseparator_get_type())
}
func (v HSeparator) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHSeparator(unsafe.Pointer(ptr)), nil
	}
}
func (v HSeparator) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HSeparator) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v HSeparator) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object HandleBox
type HandleBox struct {
	Bin
}

func (v HandleBox) native() *C.GtkHandleBox {
	return (*C.GtkHandleBox)(v.Ptr)
}
func wrapHandleBox(p *C.GtkHandleBox) (v HandleBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHandleBox(p unsafe.Pointer) (v HandleBox) {
	v.Ptr = p
	return
}
func (v HandleBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHandleBox(p unsafe.Pointer) interface{} {
	return WrapHandleBox(p)
}
func (v HandleBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_handle_box_get_type())
}
func (v HandleBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHandleBox(unsafe.Pointer(ptr)), nil
	}
}
func (v HandleBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HandleBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object HeaderBar
type HeaderBar struct {
	Container
}

func (v HeaderBar) native() *C.GtkHeaderBar {
	return (*C.GtkHeaderBar)(v.Ptr)
}
func wrapHeaderBar(p *C.GtkHeaderBar) (v HeaderBar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHeaderBar(p unsafe.Pointer) (v HeaderBar) {
	v.Ptr = p
	return
}
func (v HeaderBar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHeaderBar(p unsafe.Pointer) interface{} {
	return WrapHeaderBar(p)
}
func (v HeaderBar) GetType() gobject.Type {
	return gobject.Type(C.gtk_header_bar_get_type())
}
func (v HeaderBar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHeaderBar(unsafe.Pointer(ptr)), nil
	}
}
func (v HeaderBar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v HeaderBar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// HeaderBarNew is a wrapper around gtk_header_bar_new().
func HeaderBarNew() Widget {
	ret0 := C.gtk_header_bar_new()
	return wrapWidget(ret0)
}

// GetCustomTitle is a wrapper around gtk_header_bar_get_custom_title().
func (bar HeaderBar) GetCustomTitle() Widget {
	ret0 := C.gtk_header_bar_get_custom_title(bar.native())
	return wrapWidget(ret0)
}

// GetDecorationLayout is a wrapper around gtk_header_bar_get_decoration_layout().
func (bar HeaderBar) GetDecorationLayout() string {
	ret0 := C.gtk_header_bar_get_decoration_layout(bar.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetHasSubtitle is a wrapper around gtk_header_bar_get_has_subtitle().
func (bar HeaderBar) GetHasSubtitle() bool {
	ret0 := C.gtk_header_bar_get_has_subtitle(bar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowCloseButton is a wrapper around gtk_header_bar_get_show_close_button().
func (bar HeaderBar) GetShowCloseButton() bool {
	ret0 := C.gtk_header_bar_get_show_close_button(bar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTitle is a wrapper around gtk_header_bar_get_title().
func (bar HeaderBar) GetTitle() string {
	ret0 := C.gtk_header_bar_get_title(bar.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// PackEnd is a wrapper around gtk_header_bar_pack_end().
func (bar HeaderBar) PackEnd(child Widget) {
	C.gtk_header_bar_pack_end(bar.native(), child.native())
}

// PackStart is a wrapper around gtk_header_bar_pack_start().
func (bar HeaderBar) PackStart(child Widget) {
	C.gtk_header_bar_pack_start(bar.native(), child.native())
}

// SetCustomTitle is a wrapper around gtk_header_bar_set_custom_title().
func (bar HeaderBar) SetCustomTitle(title_widget Widget) {
	C.gtk_header_bar_set_custom_title(bar.native(), title_widget.native())
}

// SetDecorationLayout is a wrapper around gtk_header_bar_set_decoration_layout().
func (bar HeaderBar) SetDecorationLayout(layout string) {
	layout0 := (*C.gchar)(C.CString(layout))
	C.gtk_header_bar_set_decoration_layout(bar.native(), layout0)
	C.free(unsafe.Pointer(layout0)) /*ch:<stdlib.h>*/
}

// SetHasSubtitle is a wrapper around gtk_header_bar_set_has_subtitle().
func (bar HeaderBar) SetHasSubtitle(setting bool) {
	C.gtk_header_bar_set_has_subtitle(bar.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetShowCloseButton is a wrapper around gtk_header_bar_set_show_close_button().
func (bar HeaderBar) SetShowCloseButton(setting bool) {
	C.gtk_header_bar_set_show_close_button(bar.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetSubtitle is a wrapper around gtk_header_bar_set_subtitle().
func (bar HeaderBar) SetSubtitle(subtitle string) {
	subtitle0 := (*C.gchar)(C.CString(subtitle))
	C.gtk_header_bar_set_subtitle(bar.native(), subtitle0)
	C.free(unsafe.Pointer(subtitle0)) /*ch:<stdlib.h>*/
}

// SetTitle is a wrapper around gtk_header_bar_set_title().
func (bar HeaderBar) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.gtk_header_bar_set_title(bar.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// Object Scale
type Scale struct {
	Range
}

func (v Scale) native() *C.GtkScale {
	return (*C.GtkScale)(v.Ptr)
}
func wrapScale(p *C.GtkScale) (v Scale) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScale(p unsafe.Pointer) (v Scale) {
	v.Ptr = p
	return
}
func (v Scale) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScale(p unsafe.Pointer) interface{} {
	return WrapScale(p)
}
func (v Scale) GetType() gobject.Type {
	return gobject.Type(C.gtk_scale_get_type())
}
func (v Scale) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScale(unsafe.Pointer(ptr)), nil
	}
}
func (v Scale) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Scale) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Scale) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// ScaleNew is a wrapper around gtk_scale_new().
func ScaleNew(orientation Orientation, adjustment Adjustment) Widget {
	ret0 := C.gtk_scale_new(C.GtkOrientation(orientation), adjustment.native())
	return wrapWidget(ret0)
}

// ScaleNewWithRange is a wrapper around gtk_scale_new_with_range().
func ScaleNewWithRange(orientation Orientation, min float64, max float64, step float64) Widget {
	ret0 := C.gtk_scale_new_with_range(C.GtkOrientation(orientation), C.gdouble(min), C.gdouble(max), C.gdouble(step))
	return wrapWidget(ret0)
}

// AddMark is a wrapper around gtk_scale_add_mark().
func (scale Scale) AddMark(value float64, position PositionType, markup string) {
	markup0 := (*C.gchar)(C.CString(markup))
	C.gtk_scale_add_mark(scale.native(), C.gdouble(value), C.GtkPositionType(position), markup0)
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// ClearMarks is a wrapper around gtk_scale_clear_marks().
func (scale Scale) ClearMarks() {
	C.gtk_scale_clear_marks(scale.native())
}

// GetDigits is a wrapper around gtk_scale_get_digits().
func (scale Scale) GetDigits() int {
	ret0 := C.gtk_scale_get_digits(scale.native())
	return int(ret0)
}

// GetDrawValue is a wrapper around gtk_scale_get_draw_value().
func (scale Scale) GetDrawValue() bool {
	ret0 := C.gtk_scale_get_draw_value(scale.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHasOrigin is a wrapper around gtk_scale_get_has_origin().
func (scale Scale) GetHasOrigin() bool {
	ret0 := C.gtk_scale_get_has_origin(scale.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLayout is a wrapper around gtk_scale_get_layout().
func (scale Scale) GetLayout() pango.Layout {
	ret0 := C.gtk_scale_get_layout(scale.native())
	return pango.WrapLayout(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetLayoutOffsets is a wrapper around gtk_scale_get_layout_offsets().
func (scale Scale) GetLayoutOffsets() (int, int) {
	var x0 C.gint
	var y0 C.gint
	C.gtk_scale_get_layout_offsets(scale.native(), &x0, &y0)
	return int(x0), int(y0)
}

// GetValuePos is a wrapper around gtk_scale_get_value_pos().
func (scale Scale) GetValuePos() PositionType {
	ret0 := C.gtk_scale_get_value_pos(scale.native())
	return PositionType(ret0)
}

// SetDigits is a wrapper around gtk_scale_set_digits().
func (scale Scale) SetDigits(digits int) {
	C.gtk_scale_set_digits(scale.native(), C.gint(digits))
}

// SetDrawValue is a wrapper around gtk_scale_set_draw_value().
func (scale Scale) SetDrawValue(draw_value bool) {
	C.gtk_scale_set_draw_value(scale.native(), C.gboolean(util.Bool2Int(draw_value)) /*go:.util*/)
}

// SetHasOrigin is a wrapper around gtk_scale_set_has_origin().
func (scale Scale) SetHasOrigin(has_origin bool) {
	C.gtk_scale_set_has_origin(scale.native(), C.gboolean(util.Bool2Int(has_origin)) /*go:.util*/)
}

// SetValuePos is a wrapper around gtk_scale_set_value_pos().
func (scale Scale) SetValuePos(pos PositionType) {
	C.gtk_scale_set_value_pos(scale.native(), C.GtkPositionType(pos))
}

// Object Range
type Range struct {
	Widget
}

func (v Range) native() *C.GtkRange {
	return (*C.GtkRange)(v.Ptr)
}
func wrapRange(p *C.GtkRange) (v Range) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRange(p unsafe.Pointer) (v Range) {
	v.Ptr = p
	return
}
func (v Range) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRange(p unsafe.Pointer) interface{} {
	return WrapRange(p)
}
func (v Range) GetType() gobject.Type {
	return gobject.Type(C.gtk_range_get_type())
}
func (v Range) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRange(unsafe.Pointer(ptr)), nil
	}
}
func (v Range) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Range) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Range) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// GetAdjustment is a wrapper around gtk_range_get_adjustment().
func (range_ Range) GetAdjustment() Adjustment {
	ret0 := C.gtk_range_get_adjustment(range_.native())
	return wrapAdjustment(ret0)
}

// GetFillLevel is a wrapper around gtk_range_get_fill_level().
func (range_ Range) GetFillLevel() float64 {
	ret0 := C.gtk_range_get_fill_level(range_.native())
	return float64(ret0)
}

// GetFlippable is a wrapper around gtk_range_get_flippable().
func (range_ Range) GetFlippable() bool {
	ret0 := C.gtk_range_get_flippable(range_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInverted is a wrapper around gtk_range_get_inverted().
func (range_ Range) GetInverted() bool {
	ret0 := C.gtk_range_get_inverted(range_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLowerStepperSensitivity is a wrapper around gtk_range_get_lower_stepper_sensitivity().
func (range_ Range) GetLowerStepperSensitivity() SensitivityType {
	ret0 := C.gtk_range_get_lower_stepper_sensitivity(range_.native())
	return SensitivityType(ret0)
}

// GetRestrictToFillLevel is a wrapper around gtk_range_get_restrict_to_fill_level().
func (range_ Range) GetRestrictToFillLevel() bool {
	ret0 := C.gtk_range_get_restrict_to_fill_level(range_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRoundDigits is a wrapper around gtk_range_get_round_digits().
func (range_ Range) GetRoundDigits() int {
	ret0 := C.gtk_range_get_round_digits(range_.native())
	return int(ret0)
}

// GetShowFillLevel is a wrapper around gtk_range_get_show_fill_level().
func (range_ Range) GetShowFillLevel() bool {
	ret0 := C.gtk_range_get_show_fill_level(range_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSliderRange is a wrapper around gtk_range_get_slider_range().
func (range_ Range) GetSliderRange() (int, int) {
	var slider_start0 C.gint
	var slider_end0 C.gint
	C.gtk_range_get_slider_range(range_.native(), &slider_start0, &slider_end0)
	return int(slider_start0), int(slider_end0)
}

// GetSliderSizeFixed is a wrapper around gtk_range_get_slider_size_fixed().
func (range_ Range) GetSliderSizeFixed() bool {
	ret0 := C.gtk_range_get_slider_size_fixed(range_.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUpperStepperSensitivity is a wrapper around gtk_range_get_upper_stepper_sensitivity().
func (range_ Range) GetUpperStepperSensitivity() SensitivityType {
	ret0 := C.gtk_range_get_upper_stepper_sensitivity(range_.native())
	return SensitivityType(ret0)
}

// GetValue is a wrapper around gtk_range_get_value().
func (range_ Range) GetValue() float64 {
	ret0 := C.gtk_range_get_value(range_.native())
	return float64(ret0)
}

// SetAdjustment is a wrapper around gtk_range_set_adjustment().
func (range_ Range) SetAdjustment(adjustment Adjustment) {
	C.gtk_range_set_adjustment(range_.native(), adjustment.native())
}

// SetFillLevel is a wrapper around gtk_range_set_fill_level().
func (range_ Range) SetFillLevel(fill_level float64) {
	C.gtk_range_set_fill_level(range_.native(), C.gdouble(fill_level))
}

// SetFlippable is a wrapper around gtk_range_set_flippable().
func (range_ Range) SetFlippable(flippable bool) {
	C.gtk_range_set_flippable(range_.native(), C.gboolean(util.Bool2Int(flippable)) /*go:.util*/)
}

// SetIncrements is a wrapper around gtk_range_set_increments().
func (range_ Range) SetIncrements(step float64, page float64) {
	C.gtk_range_set_increments(range_.native(), C.gdouble(step), C.gdouble(page))
}

// SetInverted is a wrapper around gtk_range_set_inverted().
func (range_ Range) SetInverted(setting bool) {
	C.gtk_range_set_inverted(range_.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetLowerStepperSensitivity is a wrapper around gtk_range_set_lower_stepper_sensitivity().
func (range_ Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	C.gtk_range_set_lower_stepper_sensitivity(range_.native(), C.GtkSensitivityType(sensitivity))
}

// SetRange is a wrapper around gtk_range_set_range().
func (range_ Range) SetRange(min float64, max float64) {
	C.gtk_range_set_range(range_.native(), C.gdouble(min), C.gdouble(max))
}

// SetRestrictToFillLevel is a wrapper around gtk_range_set_restrict_to_fill_level().
func (range_ Range) SetRestrictToFillLevel(restrict_to_fill_level bool) {
	C.gtk_range_set_restrict_to_fill_level(range_.native(), C.gboolean(util.Bool2Int(restrict_to_fill_level)) /*go:.util*/)
}

// SetRoundDigits is a wrapper around gtk_range_set_round_digits().
func (range_ Range) SetRoundDigits(round_digits int) {
	C.gtk_range_set_round_digits(range_.native(), C.gint(round_digits))
}

// SetShowFillLevel is a wrapper around gtk_range_set_show_fill_level().
func (range_ Range) SetShowFillLevel(show_fill_level bool) {
	C.gtk_range_set_show_fill_level(range_.native(), C.gboolean(util.Bool2Int(show_fill_level)) /*go:.util*/)
}

// SetSliderSizeFixed is a wrapper around gtk_range_set_slider_size_fixed().
func (range_ Range) SetSliderSizeFixed(size_fixed bool) {
	C.gtk_range_set_slider_size_fixed(range_.native(), C.gboolean(util.Bool2Int(size_fixed)) /*go:.util*/)
}

// SetUpperStepperSensitivity is a wrapper around gtk_range_set_upper_stepper_sensitivity().
func (range_ Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	C.gtk_range_set_upper_stepper_sensitivity(range_.native(), C.GtkSensitivityType(sensitivity))
}

// SetValue is a wrapper around gtk_range_set_value().
func (range_ Range) SetValue(value float64) {
	C.gtk_range_set_value(range_.native(), C.gdouble(value))
}

// Object Scrollbar
type Scrollbar struct {
	Range
}

func (v Scrollbar) native() *C.GtkScrollbar {
	return (*C.GtkScrollbar)(v.Ptr)
}
func wrapScrollbar(p *C.GtkScrollbar) (v Scrollbar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScrollbar(p unsafe.Pointer) (v Scrollbar) {
	v.Ptr = p
	return
}
func (v Scrollbar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScrollbar(p unsafe.Pointer) interface{} {
	return WrapScrollbar(p)
}
func (v Scrollbar) GetType() gobject.Type {
	return gobject.Type(C.gtk_scrollbar_get_type())
}
func (v Scrollbar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScrollbar(unsafe.Pointer(ptr)), nil
	}
}
func (v Scrollbar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Scrollbar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Scrollbar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// ScrollbarNew is a wrapper around gtk_scrollbar_new().
func ScrollbarNew(orientation Orientation, adjustment Adjustment) Widget {
	ret0 := C.gtk_scrollbar_new(C.GtkOrientation(orientation), adjustment.native())
	return wrapWidget(ret0)
}

// Object Separator
type Separator struct {
	Widget
}

func (v Separator) native() *C.GtkSeparator {
	return (*C.GtkSeparator)(v.Ptr)
}
func wrapSeparator(p *C.GtkSeparator) (v Separator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSeparator(p unsafe.Pointer) (v Separator) {
	v.Ptr = p
	return
}
func (v Separator) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSeparator(p unsafe.Pointer) interface{} {
	return WrapSeparator(p)
}
func (v Separator) GetType() gobject.Type {
	return gobject.Type(C.gtk_separator_get_type())
}
func (v Separator) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSeparator(unsafe.Pointer(ptr)), nil
	}
}
func (v Separator) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Separator) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Separator) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// SeparatorNew is a wrapper around gtk_separator_new().
func SeparatorNew(orientation Orientation) Widget {
	ret0 := C.gtk_separator_new(C.GtkOrientation(orientation))
	return wrapWidget(ret0)
}

// Object IMContext
type IMContext struct {
	gobject.Object
}

func (v IMContext) native() *C.GtkIMContext {
	return (*C.GtkIMContext)(v.Ptr)
}
func wrapIMContext(p *C.GtkIMContext) (v IMContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIMContext(p unsafe.Pointer) (v IMContext) {
	v.Ptr = p
	return
}
func (v IMContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIMContext(p unsafe.Pointer) interface{} {
	return WrapIMContext(p)
}
func (v IMContext) GetType() gobject.Type {
	return gobject.Type(C.gtk_im_context_get_type())
}
func (v IMContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIMContext(unsafe.Pointer(ptr)), nil
	}
}

// DeleteSurrounding is a wrapper around gtk_im_context_delete_surrounding().
func (context IMContext) DeleteSurrounding(offset int, n_chars int) bool {
	ret0 := C.gtk_im_context_delete_surrounding(context.native(), C.gint(offset), C.gint(n_chars))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// FilterKeypress is a wrapper around gtk_im_context_filter_keypress().
func (context IMContext) FilterKeypress(event gdk.EventKey) bool {
	ret0 := C.gtk_im_context_filter_keypress(context.native(), (*C.GdkEventKey)(event.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// FocusIn is a wrapper around gtk_im_context_focus_in().
func (context IMContext) FocusIn() {
	C.gtk_im_context_focus_in(context.native())
}

// FocusOut is a wrapper around gtk_im_context_focus_out().
func (context IMContext) FocusOut() {
	C.gtk_im_context_focus_out(context.native())
}

// GetPreeditString is a wrapper around gtk_im_context_get_preedit_string().
func (context IMContext) GetPreeditString() (string, pango.AttrList, int) {
	var str0 *C.gchar
	var attrs0 *C.PangoAttrList
	var cursor_pos0 C.gint
	C.gtk_im_context_get_preedit_string(context.native(), &str0, &attrs0, &cursor_pos0)
	str := C.GoString((*C.char)(str0))
	defer C.g_free(C.gpointer(str0))
	return str, pango.WrapAttrList(unsafe.Pointer(attrs0)) /*gir:Pango*/, int(cursor_pos0)
}

// GetSurrounding is a wrapper around gtk_im_context_get_surrounding().
func (context IMContext) GetSurrounding() (bool, string, int) {
	var text0 *C.gchar
	var cursor_index0 C.gint
	ret0 := C.gtk_im_context_get_surrounding(context.native(), &text0, &cursor_index0)
	text := C.GoString((*C.char)(text0))
	defer C.g_free(C.gpointer(text0))
	return util.Int2Bool(int(ret0)) /*go:.util*/, text, int(cursor_index0)
}

// Reset is a wrapper around gtk_im_context_reset().
func (context IMContext) Reset() {
	C.gtk_im_context_reset(context.native())
}

// SetClientWindow is a wrapper around gtk_im_context_set_client_window().
func (context IMContext) SetClientWindow(window gdk.Window) {
	C.gtk_im_context_set_client_window(context.native(), (*C.GdkWindow)(window.Ptr))
}

// SetCursorLocation is a wrapper around gtk_im_context_set_cursor_location().
func (context IMContext) SetCursorLocation(area gdk.Rectangle) {
	C.gtk_im_context_set_cursor_location(context.native(), (*C.GdkRectangle)(area.Ptr))
}

// SetSurrounding is a wrapper around gtk_im_context_set_surrounding().
func (context IMContext) SetSurrounding(text string, len int, cursor_index int) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_im_context_set_surrounding(context.native(), text0, C.gint(len), C.gint(cursor_index))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetUsePreedit is a wrapper around gtk_im_context_set_use_preedit().
func (context IMContext) SetUsePreedit(use_preedit bool) {
	C.gtk_im_context_set_use_preedit(context.native(), C.gboolean(util.Bool2Int(use_preedit)) /*go:.util*/)
}

// Object IMContextSimple
type IMContextSimple struct {
	IMContext
}

func (v IMContextSimple) native() *C.GtkIMContextSimple {
	return (*C.GtkIMContextSimple)(v.Ptr)
}
func wrapIMContextSimple(p *C.GtkIMContextSimple) (v IMContextSimple) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIMContextSimple(p unsafe.Pointer) (v IMContextSimple) {
	v.Ptr = p
	return
}
func (v IMContextSimple) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIMContextSimple(p unsafe.Pointer) interface{} {
	return WrapIMContextSimple(p)
}
func (v IMContextSimple) GetType() gobject.Type {
	return gobject.Type(C.gtk_im_context_simple_get_type())
}
func (v IMContextSimple) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIMContextSimple(unsafe.Pointer(ptr)), nil
	}
}

// IMContextSimpleNew is a wrapper around gtk_im_context_simple_new().
func IMContextSimpleNew() IMContext {
	ret0 := C.gtk_im_context_simple_new()
	return wrapIMContext(ret0)
}

// AddComposeFile is a wrapper around gtk_im_context_simple_add_compose_file().
func (context_simple IMContextSimple) AddComposeFile(compose_file string) {
	compose_file0 := (*C.gchar)(C.CString(compose_file))
	C.gtk_im_context_simple_add_compose_file(context_simple.native(), compose_file0)
	C.free(unsafe.Pointer(compose_file0)) /*ch:<stdlib.h>*/
}

// AddTable is a wrapper around gtk_im_context_simple_add_table().
func (context_simple IMContextSimple) AddTable(data []uint16, max_seq_len int, n_seqs int) {
	data0 := make([]C.guint16, len(data))
	for idx, elemG := range data {
		data0[idx] = C.guint16(elemG)
	}
	var data0Ptr *C.guint16
	if len(data0) > 0 {
		data0Ptr = &data0[0]
	}
	C.gtk_im_context_simple_add_table(context_simple.native(), data0Ptr, C.gint(max_seq_len), C.gint(n_seqs))
}

// Object IMMulticontext
type IMMulticontext struct {
	IMContext
}

func (v IMMulticontext) native() *C.GtkIMMulticontext {
	return (*C.GtkIMMulticontext)(v.Ptr)
}
func wrapIMMulticontext(p *C.GtkIMMulticontext) (v IMMulticontext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIMMulticontext(p unsafe.Pointer) (v IMMulticontext) {
	v.Ptr = p
	return
}
func (v IMMulticontext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIMMulticontext(p unsafe.Pointer) interface{} {
	return WrapIMMulticontext(p)
}
func (v IMMulticontext) GetType() gobject.Type {
	return gobject.Type(C.gtk_im_multicontext_get_type())
}
func (v IMMulticontext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIMMulticontext(unsafe.Pointer(ptr)), nil
	}
}

// IMMulticontextNew is a wrapper around gtk_im_multicontext_new().
func IMMulticontextNew() IMContext {
	ret0 := C.gtk_im_multicontext_new()
	return wrapIMContext(ret0)
}

// GetContextId is a wrapper around gtk_im_multicontext_get_context_id().
func (context IMMulticontext) GetContextId() string {
	ret0 := C.gtk_im_multicontext_get_context_id(context.native())
	ret := C.GoString(ret0)
	return ret
}

// SetContextId is a wrapper around gtk_im_multicontext_set_context_id().
func (context IMMulticontext) SetContextId(context_id string) {
	context_id0 := C.CString(context_id)
	C.gtk_im_multicontext_set_context_id(context.native(), context_id0)
	C.free(unsafe.Pointer(context_id0)) /*ch:<stdlib.h>*/
}

// Object IconFactory
type IconFactory struct {
	gobject.Object
}

func (v IconFactory) native() *C.GtkIconFactory {
	return (*C.GtkIconFactory)(v.Ptr)
}
func wrapIconFactory(p *C.GtkIconFactory) (v IconFactory) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIconFactory(p unsafe.Pointer) (v IconFactory) {
	v.Ptr = p
	return
}
func (v IconFactory) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIconFactory(p unsafe.Pointer) interface{} {
	return WrapIconFactory(p)
}
func (v IconFactory) GetType() gobject.Type {
	return gobject.Type(C.gtk_icon_factory_get_type())
}
func (v IconFactory) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIconFactory(unsafe.Pointer(ptr)), nil
	}
}
func (v IconFactory) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object IconInfo
type IconInfo struct {
	gobject.Object
}

func (v IconInfo) native() *C.GtkIconInfo {
	return (*C.GtkIconInfo)(v.Ptr)
}
func wrapIconInfo(p *C.GtkIconInfo) (v IconInfo) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIconInfo(p unsafe.Pointer) (v IconInfo) {
	v.Ptr = p
	return
}
func (v IconInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIconInfo(p unsafe.Pointer) interface{} {
	return WrapIconInfo(p)
}
func (v IconInfo) GetType() gobject.Type {
	return gobject.Type(C.gtk_icon_info_get_type())
}
func (v IconInfo) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIconInfo(unsafe.Pointer(ptr)), nil
	}
}

// IconInfoNewForPixbuf is a wrapper around gtk_icon_info_new_for_pixbuf().
func IconInfoNewForPixbuf(icon_theme IconTheme, pixbuf gdkpixbuf.Pixbuf) IconInfo {
	ret0 := C.gtk_icon_info_new_for_pixbuf(icon_theme.native(), (*C.GdkPixbuf)(pixbuf.Ptr))
	return wrapIconInfo(ret0)
}

// GetBaseScale is a wrapper around gtk_icon_info_get_base_scale().
func (icon_info IconInfo) GetBaseScale() int {
	ret0 := C.gtk_icon_info_get_base_scale(icon_info.native())
	return int(ret0)
}

// GetBaseSize is a wrapper around gtk_icon_info_get_base_size().
func (icon_info IconInfo) GetBaseSize() int {
	ret0 := C.gtk_icon_info_get_base_size(icon_info.native())
	return int(ret0)
}

// GetFilename is a wrapper around gtk_icon_info_get_filename().
func (icon_info IconInfo) GetFilename() string {
	ret0 := C.gtk_icon_info_get_filename(icon_info.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// IsSymbolic is a wrapper around gtk_icon_info_is_symbolic().
func (icon_info IconInfo) IsSymbolic() bool {
	ret0 := C.gtk_icon_info_is_symbolic(icon_info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LoadIcon is a wrapper around gtk_icon_info_load_icon().
func (icon_info IconInfo) LoadIcon() (gdkpixbuf.Pixbuf, error) {
	var err glib.Error
	ret0 := C.gtk_icon_info_load_icon(icon_info.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, nil
}

// LoadIconFinish is a wrapper around gtk_icon_info_load_icon_finish().
func (icon_info IconInfo) LoadIconFinish(res gio.AsyncResult) (gdkpixbuf.Pixbuf, error) {
	var err glib.Error
	ret0 := C.gtk_icon_info_load_icon_finish(icon_info.native(), (*C.GAsyncResult)(res.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, nil
}

// LoadSurface is a wrapper around gtk_icon_info_load_surface().
func (icon_info IconInfo) LoadSurface(for_window gdk.Window) (cairo.Surface, error) {
	var err glib.Error
	ret0 := C.gtk_icon_info_load_surface(icon_info.native(), (*C.GdkWindow)(for_window.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return cairo.Surface{}, err.GoValue()
	}
	return cairo.WrapSurface(unsafe.Pointer(ret0)) /*gir:cairo*/, nil
}

// LoadSymbolic is a wrapper around gtk_icon_info_load_symbolic().
func (icon_info IconInfo) LoadSymbolic(fg gdk.RGBA, success_color gdk.RGBA, warning_color gdk.RGBA, error_color gdk.RGBA) (gdkpixbuf.Pixbuf, bool, error) {
	var was_symbolic0 C.gboolean
	var err glib.Error
	ret0 := C.gtk_icon_info_load_symbolic(icon_info.native(), (*C.GdkRGBA)(fg.Ptr), (*C.GdkRGBA)(success_color.Ptr), (*C.GdkRGBA)(warning_color.Ptr), (*C.GdkRGBA)(error_color.Ptr), &was_symbolic0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, false, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, util.Int2Bool(int(was_symbolic0)) /*go:.util*/, nil
}

// LoadSymbolicFinish is a wrapper around gtk_icon_info_load_symbolic_finish().
func (icon_info IconInfo) LoadSymbolicFinish(res gio.AsyncResult) (gdkpixbuf.Pixbuf, bool, error) {
	var was_symbolic0 C.gboolean
	var err glib.Error
	ret0 := C.gtk_icon_info_load_symbolic_finish(icon_info.native(), (*C.GAsyncResult)(res.Ptr), &was_symbolic0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, false, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, util.Int2Bool(int(was_symbolic0)) /*go:.util*/, nil
}

// LoadSymbolicForContext is a wrapper around gtk_icon_info_load_symbolic_for_context().
func (icon_info IconInfo) LoadSymbolicForContext(context StyleContext) (gdkpixbuf.Pixbuf, bool, error) {
	var was_symbolic0 C.gboolean
	var err glib.Error
	ret0 := C.gtk_icon_info_load_symbolic_for_context(icon_info.native(), context.native(), &was_symbolic0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, false, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, util.Int2Bool(int(was_symbolic0)) /*go:.util*/, nil
}

// LoadSymbolicForContextFinish is a wrapper around gtk_icon_info_load_symbolic_for_context_finish().
func (icon_info IconInfo) LoadSymbolicForContextFinish(res gio.AsyncResult) (gdkpixbuf.Pixbuf, bool, error) {
	var was_symbolic0 C.gboolean
	var err glib.Error
	ret0 := C.gtk_icon_info_load_symbolic_for_context_finish(icon_info.native(), (*C.GAsyncResult)(res.Ptr), &was_symbolic0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, false, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, util.Int2Bool(int(was_symbolic0)) /*go:.util*/, nil
}

// Object IconTheme
type IconTheme struct {
	gobject.Object
}

func (v IconTheme) native() *C.GtkIconTheme {
	return (*C.GtkIconTheme)(v.Ptr)
}
func wrapIconTheme(p *C.GtkIconTheme) (v IconTheme) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIconTheme(p unsafe.Pointer) (v IconTheme) {
	v.Ptr = p
	return
}
func (v IconTheme) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIconTheme(p unsafe.Pointer) interface{} {
	return WrapIconTheme(p)
}
func (v IconTheme) GetType() gobject.Type {
	return gobject.Type(C.gtk_icon_theme_get_type())
}
func (v IconTheme) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIconTheme(unsafe.Pointer(ptr)), nil
	}
}

// IconThemeNew is a wrapper around gtk_icon_theme_new().
func IconThemeNew() IconTheme {
	ret0 := C.gtk_icon_theme_new()
	return wrapIconTheme(ret0)
}

// AddResourcePath is a wrapper around gtk_icon_theme_add_resource_path().
func (icon_theme IconTheme) AddResourcePath(path string) {
	path0 := (*C.gchar)(C.CString(path))
	C.gtk_icon_theme_add_resource_path(icon_theme.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
}

// AppendSearchPath is a wrapper around gtk_icon_theme_append_search_path().
func (icon_theme IconTheme) AppendSearchPath(path string) {
	path0 := (*C.gchar)(C.CString(path))
	C.gtk_icon_theme_append_search_path(icon_theme.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
}

// GetExampleIconName is a wrapper around gtk_icon_theme_get_example_icon_name().
func (icon_theme IconTheme) GetExampleIconName() string {
	ret0 := C.gtk_icon_theme_get_example_icon_name(icon_theme.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetIconSizes is a wrapper around gtk_icon_theme_get_icon_sizes().
func (icon_theme IconTheme) GetIconSizes(icon_name string) []int {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	ret0 := C.gtk_icon_theme_get_icon_sizes(icon_theme.native(), icon_name0)
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
	var ret0Slice []C.gint
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(C.gint(0))) /*go:.util*/) /*go:.util*/
	ret := make([]int, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = int(elem)
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// HasIcon is a wrapper around gtk_icon_theme_has_icon().
func (icon_theme IconTheme) HasIcon(icon_name string) bool {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	ret0 := C.gtk_icon_theme_has_icon(icon_theme.native(), icon_name0)
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// LoadIcon is a wrapper around gtk_icon_theme_load_icon().
func (icon_theme IconTheme) LoadIcon(icon_name string, size int, flags IconLookupFlags) (gdkpixbuf.Pixbuf, error) {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	var err glib.Error
	ret0 := C.gtk_icon_theme_load_icon(icon_theme.native(), icon_name0, C.gint(size), C.GtkIconLookupFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, nil
}

// LoadIconForScale is a wrapper around gtk_icon_theme_load_icon_for_scale().
func (icon_theme IconTheme) LoadIconForScale(icon_name string, size int, scale int, flags IconLookupFlags) (gdkpixbuf.Pixbuf, error) {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	var err glib.Error
	ret0 := C.gtk_icon_theme_load_icon_for_scale(icon_theme.native(), icon_name0, C.gint(size), C.gint(scale), C.GtkIconLookupFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return gdkpixbuf.Pixbuf{}, err.GoValue()
	}
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/, nil
}

// LoadSurface is a wrapper around gtk_icon_theme_load_surface().
func (icon_theme IconTheme) LoadSurface(icon_name string, size int, scale int, for_window gdk.Window, flags IconLookupFlags) (cairo.Surface, error) {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	var err glib.Error
	ret0 := C.gtk_icon_theme_load_surface(icon_theme.native(), icon_name0, C.gint(size), C.gint(scale), (*C.GdkWindow)(for_window.Ptr), C.GtkIconLookupFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return cairo.Surface{}, err.GoValue()
	}
	return cairo.WrapSurface(unsafe.Pointer(ret0)) /*gir:cairo*/, nil
}

// LookupByGicon is a wrapper around gtk_icon_theme_lookup_by_gicon().
func (icon_theme IconTheme) LookupByGicon(icon gio.Icon, size int, flags IconLookupFlags) IconInfo {
	ret0 := C.gtk_icon_theme_lookup_by_gicon(icon_theme.native(), (*C.GIcon)(icon.Ptr), C.gint(size), C.GtkIconLookupFlags(flags))
	return wrapIconInfo(ret0)
}

// LookupByGiconForScale is a wrapper around gtk_icon_theme_lookup_by_gicon_for_scale().
func (icon_theme IconTheme) LookupByGiconForScale(icon gio.Icon, size int, scale int, flags IconLookupFlags) IconInfo {
	ret0 := C.gtk_icon_theme_lookup_by_gicon_for_scale(icon_theme.native(), (*C.GIcon)(icon.Ptr), C.gint(size), C.gint(scale), C.GtkIconLookupFlags(flags))
	return wrapIconInfo(ret0)
}

// LookupIcon is a wrapper around gtk_icon_theme_lookup_icon().
func (icon_theme IconTheme) LookupIcon(icon_name string, size int, flags IconLookupFlags) IconInfo {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	ret0 := C.gtk_icon_theme_lookup_icon(icon_theme.native(), icon_name0, C.gint(size), C.GtkIconLookupFlags(flags))
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
	return wrapIconInfo(ret0)
}

// LookupIconForScale is a wrapper around gtk_icon_theme_lookup_icon_for_scale().
func (icon_theme IconTheme) LookupIconForScale(icon_name string, size int, scale int, flags IconLookupFlags) IconInfo {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	ret0 := C.gtk_icon_theme_lookup_icon_for_scale(icon_theme.native(), icon_name0, C.gint(size), C.gint(scale), C.GtkIconLookupFlags(flags))
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
	return wrapIconInfo(ret0)
}

// PrependSearchPath is a wrapper around gtk_icon_theme_prepend_search_path().
func (icon_theme IconTheme) PrependSearchPath(path string) {
	path0 := (*C.gchar)(C.CString(path))
	C.gtk_icon_theme_prepend_search_path(icon_theme.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
}

// RescanIfNeeded is a wrapper around gtk_icon_theme_rescan_if_needed().
func (icon_theme IconTheme) RescanIfNeeded() bool {
	ret0 := C.gtk_icon_theme_rescan_if_needed(icon_theme.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetCustomTheme is a wrapper around gtk_icon_theme_set_custom_theme().
func (icon_theme IconTheme) SetCustomTheme(theme_name string) {
	theme_name0 := (*C.gchar)(C.CString(theme_name))
	C.gtk_icon_theme_set_custom_theme(icon_theme.native(), theme_name0)
	C.free(unsafe.Pointer(theme_name0)) /*ch:<stdlib.h>*/
}

// SetScreen is a wrapper around gtk_icon_theme_set_screen().
func (icon_theme IconTheme) SetScreen(screen gdk.Screen) {
	C.gtk_icon_theme_set_screen(icon_theme.native(), (*C.GdkScreen)(screen.Ptr))
}

// IconThemeGetDefault is a wrapper around gtk_icon_theme_get_default().
func IconThemeGetDefault() IconTheme {
	ret0 := C.gtk_icon_theme_get_default()
	return wrapIconTheme(ret0)
}

// IconThemeGetForScreen is a wrapper around gtk_icon_theme_get_for_screen().
func IconThemeGetForScreen(screen gdk.Screen) IconTheme {
	ret0 := C.gtk_icon_theme_get_for_screen((*C.GdkScreen)(screen.Ptr))
	return wrapIconTheme(ret0)
}

// Object StyleContext
type StyleContext struct {
	gobject.Object
}

func (v StyleContext) native() *C.GtkStyleContext {
	return (*C.GtkStyleContext)(v.Ptr)
}
func wrapStyleContext(p *C.GtkStyleContext) (v StyleContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStyleContext(p unsafe.Pointer) (v StyleContext) {
	v.Ptr = p
	return
}
func (v StyleContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStyleContext(p unsafe.Pointer) interface{} {
	return WrapStyleContext(p)
}
func (v StyleContext) GetType() gobject.Type {
	return gobject.Type(C.gtk_style_context_get_type())
}
func (v StyleContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStyleContext(unsafe.Pointer(ptr)), nil
	}
}

// StyleContextNew is a wrapper around gtk_style_context_new().
func StyleContextNew() StyleContext {
	ret0 := C.gtk_style_context_new()
	return wrapStyleContext(ret0)
}

// AddClass is a wrapper around gtk_style_context_add_class().
func (context StyleContext) AddClass(class_name string) {
	class_name0 := (*C.gchar)(C.CString(class_name))
	C.gtk_style_context_add_class(context.native(), class_name0)
	C.free(unsafe.Pointer(class_name0)) /*ch:<stdlib.h>*/
}

// AddProvider is a wrapper around gtk_style_context_add_provider().
func (context StyleContext) AddProvider(provider StyleProvider, priority uint) {
	C.gtk_style_context_add_provider(context.native(), provider.native(), C.guint(priority))
}

// GetFrameClock is a wrapper around gtk_style_context_get_frame_clock().
func (context StyleContext) GetFrameClock() gdk.FrameClock {
	ret0 := C.gtk_style_context_get_frame_clock(context.native())
	return gdk.WrapFrameClock(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetJunctionSides is a wrapper around gtk_style_context_get_junction_sides().
func (context StyleContext) GetJunctionSides() JunctionSides {
	ret0 := C.gtk_style_context_get_junction_sides(context.native())
	return JunctionSides(ret0)
}

// GetParent is a wrapper around gtk_style_context_get_parent().
func (context StyleContext) GetParent() StyleContext {
	ret0 := C.gtk_style_context_get_parent(context.native())
	return wrapStyleContext(ret0)
}

// GetPath is a wrapper around gtk_style_context_get_path().
func (context StyleContext) GetPath() WidgetPath {
	ret0 := C.gtk_style_context_get_path(context.native())
	return wrapWidgetPath(ret0)
}

// GetScale is a wrapper around gtk_style_context_get_scale().
func (context StyleContext) GetScale() int {
	ret0 := C.gtk_style_context_get_scale(context.native())
	return int(ret0)
}

// GetScreen is a wrapper around gtk_style_context_get_screen().
func (context StyleContext) GetScreen() gdk.Screen {
	ret0 := C.gtk_style_context_get_screen(context.native())
	return gdk.WrapScreen(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetSection is a wrapper around gtk_style_context_get_section().
func (context StyleContext) GetSection(property string) CssSection {
	property0 := (*C.gchar)(C.CString(property))
	ret0 := C.gtk_style_context_get_section(context.native(), property0)
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
	return wrapCssSection(ret0)
}

// GetState is a wrapper around gtk_style_context_get_state().
func (context StyleContext) GetState() StateFlags {
	ret0 := C.gtk_style_context_get_state(context.native())
	return StateFlags(ret0)
}

// GetStyleProperty is a wrapper around gtk_style_context_get_style_property().
func (context StyleContext) GetStyleProperty(property_name string, value gobject.Value) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.gtk_style_context_get_style_property(context.native(), property_name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// HasClass is a wrapper around gtk_style_context_has_class().
func (context StyleContext) HasClass(class_name string) bool {
	class_name0 := (*C.gchar)(C.CString(class_name))
	ret0 := C.gtk_style_context_has_class(context.native(), class_name0)
	C.free(unsafe.Pointer(class_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// RemoveClass is a wrapper around gtk_style_context_remove_class().
func (context StyleContext) RemoveClass(class_name string) {
	class_name0 := (*C.gchar)(C.CString(class_name))
	C.gtk_style_context_remove_class(context.native(), class_name0)
	C.free(unsafe.Pointer(class_name0)) /*ch:<stdlib.h>*/
}

// RemoveProvider is a wrapper around gtk_style_context_remove_provider().
func (context StyleContext) RemoveProvider(provider StyleProvider) {
	C.gtk_style_context_remove_provider(context.native(), provider.native())
}

// Restore is a wrapper around gtk_style_context_restore().
func (context StyleContext) Restore() {
	C.gtk_style_context_restore(context.native())
}

// Save is a wrapper around gtk_style_context_save().
func (context StyleContext) Save() {
	C.gtk_style_context_save(context.native())
}

// SetFrameClock is a wrapper around gtk_style_context_set_frame_clock().
func (context StyleContext) SetFrameClock(frame_clock gdk.FrameClock) {
	C.gtk_style_context_set_frame_clock(context.native(), (*C.GdkFrameClock)(frame_clock.Ptr))
}

// SetJunctionSides is a wrapper around gtk_style_context_set_junction_sides().
func (context StyleContext) SetJunctionSides(sides JunctionSides) {
	C.gtk_style_context_set_junction_sides(context.native(), C.GtkJunctionSides(sides))
}

// SetParent is a wrapper around gtk_style_context_set_parent().
func (context StyleContext) SetParent(parent StyleContext) {
	C.gtk_style_context_set_parent(context.native(), parent.native())
}

// SetPath is a wrapper around gtk_style_context_set_path().
func (context StyleContext) SetPath(path WidgetPath) {
	C.gtk_style_context_set_path(context.native(), path.native())
}

// SetScale is a wrapper around gtk_style_context_set_scale().
func (context StyleContext) SetScale(scale int) {
	C.gtk_style_context_set_scale(context.native(), C.gint(scale))
}

// SetScreen is a wrapper around gtk_style_context_set_screen().
func (context StyleContext) SetScreen(screen gdk.Screen) {
	C.gtk_style_context_set_screen(context.native(), (*C.GdkScreen)(screen.Ptr))
}

// SetState is a wrapper around gtk_style_context_set_state().
func (context StyleContext) SetState(flags StateFlags) {
	C.gtk_style_context_set_state(context.native(), C.GtkStateFlags(flags))
}

// ToString is a wrapper around gtk_style_context_to_string().
func (context StyleContext) ToString(flags StyleContextPrintFlags) string {
	ret0 := C.gtk_style_context_to_string(context.native(), C.GtkStyleContextPrintFlags(flags))
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// StyleContextAddProviderForScreen is a wrapper around gtk_style_context_add_provider_for_screen().
func StyleContextAddProviderForScreen(screen gdk.Screen, provider StyleProvider, priority uint) {
	C.gtk_style_context_add_provider_for_screen((*C.GdkScreen)(screen.Ptr), provider.native(), C.guint(priority))
}

// StyleContextRemoveProviderForScreen is a wrapper around gtk_style_context_remove_provider_for_screen().
func StyleContextRemoveProviderForScreen(screen gdk.Screen, provider StyleProvider) {
	C.gtk_style_context_remove_provider_for_screen((*C.GdkScreen)(screen.Ptr), provider.native())
}

// StyleContextResetWidgets is a wrapper around gtk_style_context_reset_widgets().
func StyleContextResetWidgets(screen gdk.Screen) {
	C.gtk_style_context_reset_widgets((*C.GdkScreen)(screen.Ptr))
}

// Object IconView
type IconView struct {
	Container
}

func (v IconView) native() *C.GtkIconView {
	return (*C.GtkIconView)(v.Ptr)
}
func wrapIconView(p *C.GtkIconView) (v IconView) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIconView(p unsafe.Pointer) (v IconView) {
	v.Ptr = p
	return
}
func (v IconView) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIconView(p unsafe.Pointer) interface{} {
	return WrapIconView(p)
}
func (v IconView) GetType() gobject.Type {
	return gobject.Type(C.gtk_icon_view_get_type())
}
func (v IconView) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIconView(unsafe.Pointer(ptr)), nil
	}
}
func (v IconView) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v IconView) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v IconView) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}
func (v IconView) Scrollable() Scrollable {
	return WrapScrollable(v.Ptr)
}

// IconViewNew is a wrapper around gtk_icon_view_new().
func IconViewNew() Widget {
	ret0 := C.gtk_icon_view_new()
	return wrapWidget(ret0)
}

// IconViewNewWithArea is a wrapper around gtk_icon_view_new_with_area().
func IconViewNewWithArea(area CellArea) Widget {
	ret0 := C.gtk_icon_view_new_with_area(area.native())
	return wrapWidget(ret0)
}

// IconViewNewWithModel is a wrapper around gtk_icon_view_new_with_model().
func IconViewNewWithModel(model TreeModel) Widget {
	ret0 := C.gtk_icon_view_new_with_model(model.native())
	return wrapWidget(ret0)
}

// ConvertWidgetToBinWindowCoords is a wrapper around gtk_icon_view_convert_widget_to_bin_window_coords().
func (icon_view IconView) ConvertWidgetToBinWindowCoords(wx int, wy int) (int, int) {
	var bx0 C.gint
	var by0 C.gint
	C.gtk_icon_view_convert_widget_to_bin_window_coords(icon_view.native(), C.gint(wx), C.gint(wy), &bx0, &by0)
	return int(bx0), int(by0)
}

// CreateDragIcon is a wrapper around gtk_icon_view_create_drag_icon().
func (icon_view IconView) CreateDragIcon(path TreePath) cairo.Surface {
	ret0 := C.gtk_icon_view_create_drag_icon(icon_view.native(), path.native())
	return cairo.WrapSurface(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetActivateOnSingleClick is a wrapper around gtk_icon_view_get_activate_on_single_click().
func (icon_view IconView) GetActivateOnSingleClick() bool {
	ret0 := C.gtk_icon_view_get_activate_on_single_click(icon_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetColumnSpacing is a wrapper around gtk_icon_view_get_column_spacing().
func (icon_view IconView) GetColumnSpacing() int {
	ret0 := C.gtk_icon_view_get_column_spacing(icon_view.native())
	return int(ret0)
}

// GetColumns is a wrapper around gtk_icon_view_get_columns().
func (icon_view IconView) GetColumns() int {
	ret0 := C.gtk_icon_view_get_columns(icon_view.native())
	return int(ret0)
}

// GetCursor is a wrapper around gtk_icon_view_get_cursor().
func (icon_view IconView) GetCursor() (bool, TreePath, CellRenderer) {
	var path0 *C.GtkTreePath
	var cell0 *C.GtkCellRenderer
	ret0 := C.gtk_icon_view_get_cursor(icon_view.native(), &path0, &cell0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(path0), wrapCellRenderer(cell0)
}

// GetDestItemAtPos is a wrapper around gtk_icon_view_get_dest_item_at_pos().
func (icon_view IconView) GetDestItemAtPos(drag_x int, drag_y int) (bool, TreePath, IconViewDropPosition) {
	var path0 *C.GtkTreePath
	var pos0 C.GtkIconViewDropPosition
	ret0 := C.gtk_icon_view_get_dest_item_at_pos(icon_view.native(), C.gint(drag_x), C.gint(drag_y), &path0, &pos0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(path0), IconViewDropPosition(pos0)
}

// GetDragDestItem is a wrapper around gtk_icon_view_get_drag_dest_item().
func (icon_view IconView) GetDragDestItem() (TreePath, IconViewDropPosition) {
	var path0 *C.GtkTreePath
	var pos0 C.GtkIconViewDropPosition
	C.gtk_icon_view_get_drag_dest_item(icon_view.native(), &path0, &pos0)
	return wrapTreePath(path0), IconViewDropPosition(pos0)
}

// GetItemAtPos is a wrapper around gtk_icon_view_get_item_at_pos().
func (icon_view IconView) GetItemAtPos(x int, y int) (bool, TreePath, CellRenderer) {
	var path0 *C.GtkTreePath
	var cell0 *C.GtkCellRenderer
	ret0 := C.gtk_icon_view_get_item_at_pos(icon_view.native(), C.gint(x), C.gint(y), &path0, &cell0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(path0), wrapCellRenderer(cell0)
}

// GetItemColumn is a wrapper around gtk_icon_view_get_item_column().
func (icon_view IconView) GetItemColumn(path TreePath) int {
	ret0 := C.gtk_icon_view_get_item_column(icon_view.native(), path.native())
	return int(ret0)
}

// GetItemOrientation is a wrapper around gtk_icon_view_get_item_orientation().
func (icon_view IconView) GetItemOrientation() Orientation {
	ret0 := C.gtk_icon_view_get_item_orientation(icon_view.native())
	return Orientation(ret0)
}

// GetItemPadding is a wrapper around gtk_icon_view_get_item_padding().
func (icon_view IconView) GetItemPadding() int {
	ret0 := C.gtk_icon_view_get_item_padding(icon_view.native())
	return int(ret0)
}

// GetItemRow is a wrapper around gtk_icon_view_get_item_row().
func (icon_view IconView) GetItemRow(path TreePath) int {
	ret0 := C.gtk_icon_view_get_item_row(icon_view.native(), path.native())
	return int(ret0)
}

// GetItemWidth is a wrapper around gtk_icon_view_get_item_width().
func (icon_view IconView) GetItemWidth() int {
	ret0 := C.gtk_icon_view_get_item_width(icon_view.native())
	return int(ret0)
}

// GetMargin is a wrapper around gtk_icon_view_get_margin().
func (icon_view IconView) GetMargin() int {
	ret0 := C.gtk_icon_view_get_margin(icon_view.native())
	return int(ret0)
}

// GetMarkupColumn is a wrapper around gtk_icon_view_get_markup_column().
func (icon_view IconView) GetMarkupColumn() int {
	ret0 := C.gtk_icon_view_get_markup_column(icon_view.native())
	return int(ret0)
}

// GetModel is a wrapper around gtk_icon_view_get_model().
func (icon_view IconView) GetModel() TreeModel {
	ret0 := C.gtk_icon_view_get_model(icon_view.native())
	return wrapTreeModel(ret0)
}

// GetPathAtPos is a wrapper around gtk_icon_view_get_path_at_pos().
func (icon_view IconView) GetPathAtPos(x int, y int) TreePath {
	ret0 := C.gtk_icon_view_get_path_at_pos(icon_view.native(), C.gint(x), C.gint(y))
	return wrapTreePath(ret0)
}

// GetPixbufColumn is a wrapper around gtk_icon_view_get_pixbuf_column().
func (icon_view IconView) GetPixbufColumn() int {
	ret0 := C.gtk_icon_view_get_pixbuf_column(icon_view.native())
	return int(ret0)
}

// GetReorderable is a wrapper around gtk_icon_view_get_reorderable().
func (icon_view IconView) GetReorderable() bool {
	ret0 := C.gtk_icon_view_get_reorderable(icon_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRowSpacing is a wrapper around gtk_icon_view_get_row_spacing().
func (icon_view IconView) GetRowSpacing() int {
	ret0 := C.gtk_icon_view_get_row_spacing(icon_view.native())
	return int(ret0)
}

// GetSelectedItems is a wrapper around gtk_icon_view_get_selected_items().
func (icon_view IconView) GetSelectedItems() glib.List {
	ret0 := C.gtk_icon_view_get_selected_items(icon_view.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapTreePath(p) }) /*gir:GLib*/
}

// GetSelectionMode is a wrapper around gtk_icon_view_get_selection_mode().
func (icon_view IconView) GetSelectionMode() SelectionMode {
	ret0 := C.gtk_icon_view_get_selection_mode(icon_view.native())
	return SelectionMode(ret0)
}

// GetSpacing is a wrapper around gtk_icon_view_get_spacing().
func (icon_view IconView) GetSpacing() int {
	ret0 := C.gtk_icon_view_get_spacing(icon_view.native())
	return int(ret0)
}

// GetTextColumn is a wrapper around gtk_icon_view_get_text_column().
func (icon_view IconView) GetTextColumn() int {
	ret0 := C.gtk_icon_view_get_text_column(icon_view.native())
	return int(ret0)
}

// GetTooltipColumn is a wrapper around gtk_icon_view_get_tooltip_column().
func (icon_view IconView) GetTooltipColumn() int {
	ret0 := C.gtk_icon_view_get_tooltip_column(icon_view.native())
	return int(ret0)
}

// GetVisibleRange is a wrapper around gtk_icon_view_get_visible_range().
func (icon_view IconView) GetVisibleRange() (bool, TreePath, TreePath) {
	var start_path0 *C.GtkTreePath
	var end_path0 *C.GtkTreePath
	ret0 := C.gtk_icon_view_get_visible_range(icon_view.native(), &start_path0, &end_path0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(start_path0), wrapTreePath(end_path0)
}

// ItemActivated is a wrapper around gtk_icon_view_item_activated().
func (icon_view IconView) ItemActivated(path TreePath) {
	C.gtk_icon_view_item_activated(icon_view.native(), path.native())
}

// PathIsSelected is a wrapper around gtk_icon_view_path_is_selected().
func (icon_view IconView) PathIsSelected(path TreePath) bool {
	ret0 := C.gtk_icon_view_path_is_selected(icon_view.native(), path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ScrollToPath is a wrapper around gtk_icon_view_scroll_to_path().
func (icon_view IconView) ScrollToPath(path TreePath, use_align bool, row_align float32, col_align float32) {
	C.gtk_icon_view_scroll_to_path(icon_view.native(), path.native(), C.gboolean(util.Bool2Int(use_align)) /*go:.util*/, C.gfloat(row_align), C.gfloat(col_align))
}

// SelectAll is a wrapper around gtk_icon_view_select_all().
func (icon_view IconView) SelectAll() {
	C.gtk_icon_view_select_all(icon_view.native())
}

// SelectPath is a wrapper around gtk_icon_view_select_path().
func (icon_view IconView) SelectPath(path TreePath) {
	C.gtk_icon_view_select_path(icon_view.native(), path.native())
}

// SetActivateOnSingleClick is a wrapper around gtk_icon_view_set_activate_on_single_click().
func (icon_view IconView) SetActivateOnSingleClick(single bool) {
	C.gtk_icon_view_set_activate_on_single_click(icon_view.native(), C.gboolean(util.Bool2Int(single)) /*go:.util*/)
}

// SetColumnSpacing is a wrapper around gtk_icon_view_set_column_spacing().
func (icon_view IconView) SetColumnSpacing(column_spacing int) {
	C.gtk_icon_view_set_column_spacing(icon_view.native(), C.gint(column_spacing))
}

// SetColumns is a wrapper around gtk_icon_view_set_columns().
func (icon_view IconView) SetColumns(columns int) {
	C.gtk_icon_view_set_columns(icon_view.native(), C.gint(columns))
}

// SetCursor is a wrapper around gtk_icon_view_set_cursor().
func (icon_view IconView) SetCursor(path TreePath, cell CellRenderer, start_editing bool) {
	C.gtk_icon_view_set_cursor(icon_view.native(), path.native(), cell.native(), C.gboolean(util.Bool2Int(start_editing)) /*go:.util*/)
}

// SetDragDestItem is a wrapper around gtk_icon_view_set_drag_dest_item().
func (icon_view IconView) SetDragDestItem(path TreePath, pos IconViewDropPosition) {
	C.gtk_icon_view_set_drag_dest_item(icon_view.native(), path.native(), C.GtkIconViewDropPosition(pos))
}

// SetItemOrientation is a wrapper around gtk_icon_view_set_item_orientation().
func (icon_view IconView) SetItemOrientation(orientation Orientation) {
	C.gtk_icon_view_set_item_orientation(icon_view.native(), C.GtkOrientation(orientation))
}

// SetItemPadding is a wrapper around gtk_icon_view_set_item_padding().
func (icon_view IconView) SetItemPadding(item_padding int) {
	C.gtk_icon_view_set_item_padding(icon_view.native(), C.gint(item_padding))
}

// SetItemWidth is a wrapper around gtk_icon_view_set_item_width().
func (icon_view IconView) SetItemWidth(item_width int) {
	C.gtk_icon_view_set_item_width(icon_view.native(), C.gint(item_width))
}

// SetMargin is a wrapper around gtk_icon_view_set_margin().
func (icon_view IconView) SetMargin(margin int) {
	C.gtk_icon_view_set_margin(icon_view.native(), C.gint(margin))
}

// SetMarkupColumn is a wrapper around gtk_icon_view_set_markup_column().
func (icon_view IconView) SetMarkupColumn(column int) {
	C.gtk_icon_view_set_markup_column(icon_view.native(), C.gint(column))
}

// SetModel is a wrapper around gtk_icon_view_set_model().
func (icon_view IconView) SetModel(model TreeModel) {
	C.gtk_icon_view_set_model(icon_view.native(), model.native())
}

// SetPixbufColumn is a wrapper around gtk_icon_view_set_pixbuf_column().
func (icon_view IconView) SetPixbufColumn(column int) {
	C.gtk_icon_view_set_pixbuf_column(icon_view.native(), C.gint(column))
}

// SetReorderable is a wrapper around gtk_icon_view_set_reorderable().
func (icon_view IconView) SetReorderable(reorderable bool) {
	C.gtk_icon_view_set_reorderable(icon_view.native(), C.gboolean(util.Bool2Int(reorderable)) /*go:.util*/)
}

// SetRowSpacing is a wrapper around gtk_icon_view_set_row_spacing().
func (icon_view IconView) SetRowSpacing(row_spacing int) {
	C.gtk_icon_view_set_row_spacing(icon_view.native(), C.gint(row_spacing))
}

// SetSelectionMode is a wrapper around gtk_icon_view_set_selection_mode().
func (icon_view IconView) SetSelectionMode(mode SelectionMode) {
	C.gtk_icon_view_set_selection_mode(icon_view.native(), C.GtkSelectionMode(mode))
}

// SetSpacing is a wrapper around gtk_icon_view_set_spacing().
func (icon_view IconView) SetSpacing(spacing int) {
	C.gtk_icon_view_set_spacing(icon_view.native(), C.gint(spacing))
}

// SetTextColumn is a wrapper around gtk_icon_view_set_text_column().
func (icon_view IconView) SetTextColumn(column int) {
	C.gtk_icon_view_set_text_column(icon_view.native(), C.gint(column))
}

// SetTooltipCell is a wrapper around gtk_icon_view_set_tooltip_cell().
func (icon_view IconView) SetTooltipCell(tooltip Tooltip, path TreePath, cell CellRenderer) {
	C.gtk_icon_view_set_tooltip_cell(icon_view.native(), tooltip.native(), path.native(), cell.native())
}

// SetTooltipColumn is a wrapper around gtk_icon_view_set_tooltip_column().
func (icon_view IconView) SetTooltipColumn(column int) {
	C.gtk_icon_view_set_tooltip_column(icon_view.native(), C.gint(column))
}

// SetTooltipItem is a wrapper around gtk_icon_view_set_tooltip_item().
func (icon_view IconView) SetTooltipItem(tooltip Tooltip, path TreePath) {
	C.gtk_icon_view_set_tooltip_item(icon_view.native(), tooltip.native(), path.native())
}

// UnselectAll is a wrapper around gtk_icon_view_unselect_all().
func (icon_view IconView) UnselectAll() {
	C.gtk_icon_view_unselect_all(icon_view.native())
}

// UnselectPath is a wrapper around gtk_icon_view_unselect_path().
func (icon_view IconView) UnselectPath(path TreePath) {
	C.gtk_icon_view_unselect_path(icon_view.native(), path.native())
}

// UnsetModelDragDest is a wrapper around gtk_icon_view_unset_model_drag_dest().
func (icon_view IconView) UnsetModelDragDest() {
	C.gtk_icon_view_unset_model_drag_dest(icon_view.native())
}

// UnsetModelDragSource is a wrapper around gtk_icon_view_unset_model_drag_source().
func (icon_view IconView) UnsetModelDragSource() {
	C.gtk_icon_view_unset_model_drag_source(icon_view.native())
}

// Object Tooltip
type Tooltip struct {
	gobject.Object
}

func (v Tooltip) native() *C.GtkTooltip {
	return (*C.GtkTooltip)(v.Ptr)
}
func wrapTooltip(p *C.GtkTooltip) (v Tooltip) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTooltip(p unsafe.Pointer) (v Tooltip) {
	v.Ptr = p
	return
}
func (v Tooltip) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTooltip(p unsafe.Pointer) interface{} {
	return WrapTooltip(p)
}
func (v Tooltip) GetType() gobject.Type {
	return gobject.Type(C.gtk_tooltip_get_type())
}
func (v Tooltip) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTooltip(unsafe.Pointer(ptr)), nil
	}
}

// SetCustom is a wrapper around gtk_tooltip_set_custom().
func (tooltip Tooltip) SetCustom(custom_widget Widget) {
	C.gtk_tooltip_set_custom(tooltip.native(), custom_widget.native())
}

// SetIcon is a wrapper around gtk_tooltip_set_icon().
func (tooltip Tooltip) SetIcon(pixbuf gdkpixbuf.Pixbuf) {
	C.gtk_tooltip_set_icon(tooltip.native(), (*C.GdkPixbuf)(pixbuf.Ptr))
}

// SetMarkup is a wrapper around gtk_tooltip_set_markup().
func (tooltip Tooltip) SetMarkup(markup string) {
	markup0 := (*C.gchar)(C.CString(markup))
	C.gtk_tooltip_set_markup(tooltip.native(), markup0)
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// SetText is a wrapper around gtk_tooltip_set_text().
func (tooltip Tooltip) SetText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_tooltip_set_text(tooltip.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetTipArea is a wrapper around gtk_tooltip_set_tip_area().
func (tooltip Tooltip) SetTipArea(rect gdk.Rectangle) {
	C.gtk_tooltip_set_tip_area(tooltip.native(), (*C.GdkRectangle)(rect.Ptr))
}

// TooltipTriggerTooltipQuery is a wrapper around gtk_tooltip_trigger_tooltip_query().
func TooltipTriggerTooltipQuery(display gdk.Display) {
	C.gtk_tooltip_trigger_tooltip_query((*C.GdkDisplay)(display.Ptr))
}

// Object IconViewAccessible
type IconViewAccessible struct {
	ContainerAccessible
}

func (v IconViewAccessible) native() *C.GtkIconViewAccessible {
	return (*C.GtkIconViewAccessible)(v.Ptr)
}
func wrapIconViewAccessible(p *C.GtkIconViewAccessible) (v IconViewAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIconViewAccessible(p unsafe.Pointer) (v IconViewAccessible) {
	v.Ptr = p
	return
}
func (v IconViewAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIconViewAccessible(p unsafe.Pointer) interface{} {
	return WrapIconViewAccessible(p)
}
func (v IconViewAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_icon_view_accessible_get_type())
}
func (v IconViewAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIconViewAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v IconViewAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v IconViewAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object Image
type Image struct {
	Misc
}

func (v Image) native() *C.GtkImage {
	return (*C.GtkImage)(v.Ptr)
}
func wrapImage(p *C.GtkImage) (v Image) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapImage(p unsafe.Pointer) (v Image) {
	v.Ptr = p
	return
}
func (v Image) IsNil() bool {
	return v.Ptr == nil
}
func IWrapImage(p unsafe.Pointer) interface{} {
	return WrapImage(p)
}
func (v Image) GetType() gobject.Type {
	return gobject.Type(C.gtk_image_get_type())
}
func (v Image) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapImage(unsafe.Pointer(ptr)), nil
	}
}
func (v Image) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Image) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ImageNew is a wrapper around gtk_image_new().
func ImageNew() Widget {
	ret0 := C.gtk_image_new()
	return wrapWidget(ret0)
}

// ImageNewFromFile is a wrapper around gtk_image_new_from_file().
func ImageNewFromFile(filename string) Widget {
	filename0 := (*C.gchar)(C.CString(filename))
	ret0 := C.gtk_image_new_from_file(filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// ImageNewFromResource is a wrapper around gtk_image_new_from_resource().
func ImageNewFromResource(resource_path string) Widget {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	ret0 := C.gtk_image_new_from_resource(resource_path0)
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// Clear is a wrapper around gtk_image_clear().
func (image Image) Clear() {
	C.gtk_image_clear(image.native())
}

// GetAnimation is a wrapper around gtk_image_get_animation().
func (image Image) GetAnimation() gdkpixbuf.PixbufAnimation {
	ret0 := C.gtk_image_get_animation(image.native())
	return gdkpixbuf.WrapPixbufAnimation(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetPixbuf is a wrapper around gtk_image_get_pixbuf().
func (image Image) GetPixbuf() gdkpixbuf.Pixbuf {
	ret0 := C.gtk_image_get_pixbuf(image.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetPixelSize is a wrapper around gtk_image_get_pixel_size().
func (image Image) GetPixelSize() int {
	ret0 := C.gtk_image_get_pixel_size(image.native())
	return int(ret0)
}

// GetStorageType is a wrapper around gtk_image_get_storage_type().
func (image Image) GetStorageType() ImageType {
	ret0 := C.gtk_image_get_storage_type(image.native())
	return ImageType(ret0)
}

// SetFromFile is a wrapper around gtk_image_set_from_file().
func (image Image) SetFromFile(filename string) {
	filename0 := (*C.gchar)(C.CString(filename))
	C.gtk_image_set_from_file(image.native(), filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
}

// SetFromResource is a wrapper around gtk_image_set_from_resource().
func (image Image) SetFromResource(resource_path string) {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	C.gtk_image_set_from_resource(image.native(), resource_path0)
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
}

// SetPixelSize is a wrapper around gtk_image_set_pixel_size().
func (image Image) SetPixelSize(pixel_size int) {
	C.gtk_image_set_pixel_size(image.native(), C.gint(pixel_size))
}

// Object ImageAccessible
type ImageAccessible struct {
	WidgetAccessible
}

func (v ImageAccessible) native() *C.GtkImageAccessible {
	return (*C.GtkImageAccessible)(v.Ptr)
}
func wrapImageAccessible(p *C.GtkImageAccessible) (v ImageAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapImageAccessible(p unsafe.Pointer) (v ImageAccessible) {
	v.Ptr = p
	return
}
func (v ImageAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapImageAccessible(p unsafe.Pointer) interface{} {
	return WrapImageAccessible(p)
}
func (v ImageAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_image_accessible_get_type())
}
func (v ImageAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapImageAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ImageAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ImageAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object ImageCellAccessible
type ImageCellAccessible struct {
	RendererCellAccessible
}

func (v ImageCellAccessible) native() *C.GtkImageCellAccessible {
	return (*C.GtkImageCellAccessible)(v.Ptr)
}
func wrapImageCellAccessible(p *C.GtkImageCellAccessible) (v ImageCellAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapImageCellAccessible(p unsafe.Pointer) (v ImageCellAccessible) {
	v.Ptr = p
	return
}
func (v ImageCellAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapImageCellAccessible(p unsafe.Pointer) interface{} {
	return WrapImageCellAccessible(p)
}
func (v ImageCellAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_image_cell_accessible_get_type())
}
func (v ImageCellAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapImageCellAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ImageCellAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v ImageCellAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ImageCellAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object ImageMenuItem
type ImageMenuItem struct {
	MenuItem
}

func (v ImageMenuItem) native() *C.GtkImageMenuItem {
	return (*C.GtkImageMenuItem)(v.Ptr)
}
func wrapImageMenuItem(p *C.GtkImageMenuItem) (v ImageMenuItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapImageMenuItem(p unsafe.Pointer) (v ImageMenuItem) {
	v.Ptr = p
	return
}
func (v ImageMenuItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapImageMenuItem(p unsafe.Pointer) interface{} {
	return WrapImageMenuItem(p)
}
func (v ImageMenuItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_image_menu_item_get_type())
}
func (v ImageMenuItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapImageMenuItem(unsafe.Pointer(ptr)), nil
	}
}
func (v ImageMenuItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ImageMenuItem) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v ImageMenuItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ImageMenuItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object RendererCellAccessible
type RendererCellAccessible struct {
	CellAccessible
}

func (v RendererCellAccessible) native() *C.GtkRendererCellAccessible {
	return (*C.GtkRendererCellAccessible)(v.Ptr)
}
func wrapRendererCellAccessible(p *C.GtkRendererCellAccessible) (v RendererCellAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRendererCellAccessible(p unsafe.Pointer) (v RendererCellAccessible) {
	v.Ptr = p
	return
}
func (v RendererCellAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRendererCellAccessible(p unsafe.Pointer) interface{} {
	return WrapRendererCellAccessible(p)
}
func (v RendererCellAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_renderer_cell_accessible_get_type())
}
func (v RendererCellAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRendererCellAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v RendererCellAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v RendererCellAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// RendererCellAccessibleNew is a wrapper around gtk_renderer_cell_accessible_new().
func RendererCellAccessibleNew(renderer CellRenderer) atk.Object {
	ret0 := C.gtk_renderer_cell_accessible_new(renderer.native())
	return atk.WrapObject(unsafe.Pointer(ret0)) /*gir:Atk*/
}

// Object InfoBar
type InfoBar struct {
	Box
}

func (v InfoBar) native() *C.GtkInfoBar {
	return (*C.GtkInfoBar)(v.Ptr)
}
func wrapInfoBar(p *C.GtkInfoBar) (v InfoBar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInfoBar(p unsafe.Pointer) (v InfoBar) {
	v.Ptr = p
	return
}
func (v InfoBar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInfoBar(p unsafe.Pointer) interface{} {
	return WrapInfoBar(p)
}
func (v InfoBar) GetType() gobject.Type {
	return gobject.Type(C.gtk_info_bar_get_type())
}
func (v InfoBar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInfoBar(unsafe.Pointer(ptr)), nil
	}
}
func (v InfoBar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v InfoBar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v InfoBar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// InfoBarNew is a wrapper around gtk_info_bar_new().
func InfoBarNew() Widget {
	ret0 := C.gtk_info_bar_new()
	return wrapWidget(ret0)
}

// AddActionWidget is a wrapper around gtk_info_bar_add_action_widget().
func (info_bar InfoBar) AddActionWidget(child Widget, response_id int) {
	C.gtk_info_bar_add_action_widget(info_bar.native(), child.native(), C.gint(response_id))
}

// GetActionArea is a wrapper around gtk_info_bar_get_action_area().
func (info_bar InfoBar) GetActionArea() Widget {
	ret0 := C.gtk_info_bar_get_action_area(info_bar.native())
	return wrapWidget(ret0)
}

// GetContentArea is a wrapper around gtk_info_bar_get_content_area().
func (info_bar InfoBar) GetContentArea() Widget {
	ret0 := C.gtk_info_bar_get_content_area(info_bar.native())
	return wrapWidget(ret0)
}

// GetMessageType is a wrapper around gtk_info_bar_get_message_type().
func (info_bar InfoBar) GetMessageType() MessageType {
	ret0 := C.gtk_info_bar_get_message_type(info_bar.native())
	return MessageType(ret0)
}

// GetShowCloseButton is a wrapper around gtk_info_bar_get_show_close_button().
func (info_bar InfoBar) GetShowCloseButton() bool {
	ret0 := C.gtk_info_bar_get_show_close_button(info_bar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Response is a wrapper around gtk_info_bar_response().
func (info_bar InfoBar) Response(response_id int) {
	C.gtk_info_bar_response(info_bar.native(), C.gint(response_id))
}

// SetDefaultResponse is a wrapper around gtk_info_bar_set_default_response().
func (info_bar InfoBar) SetDefaultResponse(response_id int) {
	C.gtk_info_bar_set_default_response(info_bar.native(), C.gint(response_id))
}

// SetMessageType is a wrapper around gtk_info_bar_set_message_type().
func (info_bar InfoBar) SetMessageType(message_type MessageType) {
	C.gtk_info_bar_set_message_type(info_bar.native(), C.GtkMessageType(message_type))
}

// SetResponseSensitive is a wrapper around gtk_info_bar_set_response_sensitive().
func (info_bar InfoBar) SetResponseSensitive(response_id int, setting bool) {
	C.gtk_info_bar_set_response_sensitive(info_bar.native(), C.gint(response_id), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetShowCloseButton is a wrapper around gtk_info_bar_set_show_close_button().
func (info_bar InfoBar) SetShowCloseButton(setting bool) {
	C.gtk_info_bar_set_show_close_button(info_bar.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// Object Invisible
type Invisible struct {
	Widget
}

func (v Invisible) native() *C.GtkInvisible {
	return (*C.GtkInvisible)(v.Ptr)
}
func wrapInvisible(p *C.GtkInvisible) (v Invisible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInvisible(p unsafe.Pointer) (v Invisible) {
	v.Ptr = p
	return
}
func (v Invisible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInvisible(p unsafe.Pointer) interface{} {
	return WrapInvisible(p)
}
func (v Invisible) GetType() gobject.Type {
	return gobject.Type(C.gtk_invisible_get_type())
}
func (v Invisible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInvisible(unsafe.Pointer(ptr)), nil
	}
}
func (v Invisible) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Invisible) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// InvisibleNew is a wrapper around gtk_invisible_new().
func InvisibleNew() Widget {
	ret0 := C.gtk_invisible_new()
	return wrapWidget(ret0)
}

// GetScreen is a wrapper around gtk_invisible_get_screen().
func (invisible Invisible) GetScreen() gdk.Screen {
	ret0 := C.gtk_invisible_get_screen(invisible.native())
	return gdk.WrapScreen(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// Object LabelAccessible
type LabelAccessible struct {
	WidgetAccessible
}

func (v LabelAccessible) native() *C.GtkLabelAccessible {
	return (*C.GtkLabelAccessible)(v.Ptr)
}
func wrapLabelAccessible(p *C.GtkLabelAccessible) (v LabelAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLabelAccessible(p unsafe.Pointer) (v LabelAccessible) {
	v.Ptr = p
	return
}
func (v LabelAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLabelAccessible(p unsafe.Pointer) interface{} {
	return WrapLabelAccessible(p)
}
func (v LabelAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_label_accessible_get_type())
}
func (v LabelAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLabelAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v LabelAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v LabelAccessible) Hypertext() atk.Hypertext {
	return atk.WrapHypertext(v.Ptr) /*gir:Atk*/
}
func (v LabelAccessible) Text() atk.Text {
	return atk.WrapText(v.Ptr) /*gir:Atk*/
}

// Object Layout
type Layout struct {
	Container
}

func (v Layout) native() *C.GtkLayout {
	return (*C.GtkLayout)(v.Ptr)
}
func wrapLayout(p *C.GtkLayout) (v Layout) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLayout(p unsafe.Pointer) (v Layout) {
	v.Ptr = p
	return
}
func (v Layout) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLayout(p unsafe.Pointer) interface{} {
	return WrapLayout(p)
}
func (v Layout) GetType() gobject.Type {
	return gobject.Type(C.gtk_layout_get_type())
}
func (v Layout) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLayout(unsafe.Pointer(ptr)), nil
	}
}
func (v Layout) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Layout) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Layout) Scrollable() Scrollable {
	return WrapScrollable(v.Ptr)
}

// LayoutNew is a wrapper around gtk_layout_new().
func LayoutNew(hadjustment Adjustment, vadjustment Adjustment) Widget {
	ret0 := C.gtk_layout_new(hadjustment.native(), vadjustment.native())
	return wrapWidget(ret0)
}

// GetSize is a wrapper around gtk_layout_get_size().
func (layout Layout) GetSize() (uint, uint) {
	var width0 C.guint
	var height0 C.guint
	C.gtk_layout_get_size(layout.native(), &width0, &height0)
	return uint(width0), uint(height0)
}

// Move is a wrapper around gtk_layout_move().
func (layout Layout) Move(child_widget Widget, x int, y int) {
	C.gtk_layout_move(layout.native(), child_widget.native(), C.gint(x), C.gint(y))
}

// Put is a wrapper around gtk_layout_put().
func (layout Layout) Put(child_widget Widget, x int, y int) {
	C.gtk_layout_put(layout.native(), child_widget.native(), C.gint(x), C.gint(y))
}

// SetSize is a wrapper around gtk_layout_set_size().
func (layout Layout) SetSize(width uint, height uint) {
	C.gtk_layout_set_size(layout.native(), C.guint(width), C.guint(height))
}

// Object LevelBar
type LevelBar struct {
	Widget
}

func (v LevelBar) native() *C.GtkLevelBar {
	return (*C.GtkLevelBar)(v.Ptr)
}
func wrapLevelBar(p *C.GtkLevelBar) (v LevelBar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLevelBar(p unsafe.Pointer) (v LevelBar) {
	v.Ptr = p
	return
}
func (v LevelBar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLevelBar(p unsafe.Pointer) interface{} {
	return WrapLevelBar(p)
}
func (v LevelBar) GetType() gobject.Type {
	return gobject.Type(C.gtk_level_bar_get_type())
}
func (v LevelBar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLevelBar(unsafe.Pointer(ptr)), nil
	}
}
func (v LevelBar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v LevelBar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v LevelBar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// LevelBarNew is a wrapper around gtk_level_bar_new().
func LevelBarNew() Widget {
	ret0 := C.gtk_level_bar_new()
	return wrapWidget(ret0)
}

// LevelBarNewForInterval is a wrapper around gtk_level_bar_new_for_interval().
func LevelBarNewForInterval(min_value float64, max_value float64) Widget {
	ret0 := C.gtk_level_bar_new_for_interval(C.gdouble(min_value), C.gdouble(max_value))
	return wrapWidget(ret0)
}

// AddOffsetValue is a wrapper around gtk_level_bar_add_offset_value().
func (self LevelBar) AddOffsetValue(name string, value float64) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_level_bar_add_offset_value(self.native(), name0, C.gdouble(value))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// GetInverted is a wrapper around gtk_level_bar_get_inverted().
func (self LevelBar) GetInverted() bool {
	ret0 := C.gtk_level_bar_get_inverted(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMaxValue is a wrapper around gtk_level_bar_get_max_value().
func (self LevelBar) GetMaxValue() float64 {
	ret0 := C.gtk_level_bar_get_max_value(self.native())
	return float64(ret0)
}

// GetMinValue is a wrapper around gtk_level_bar_get_min_value().
func (self LevelBar) GetMinValue() float64 {
	ret0 := C.gtk_level_bar_get_min_value(self.native())
	return float64(ret0)
}

// GetMode is a wrapper around gtk_level_bar_get_mode().
func (self LevelBar) GetMode() LevelBarMode {
	ret0 := C.gtk_level_bar_get_mode(self.native())
	return LevelBarMode(ret0)
}

// GetOffsetValue is a wrapper around gtk_level_bar_get_offset_value().
func (self LevelBar) GetOffsetValue(name string) (bool, float64) {
	name0 := (*C.gchar)(C.CString(name))
	var value0 C.gdouble
	ret0 := C.gtk_level_bar_get_offset_value(self.native(), name0, &value0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(value0)
}

// GetValue is a wrapper around gtk_level_bar_get_value().
func (self LevelBar) GetValue() float64 {
	ret0 := C.gtk_level_bar_get_value(self.native())
	return float64(ret0)
}

// RemoveOffsetValue is a wrapper around gtk_level_bar_remove_offset_value().
func (self LevelBar) RemoveOffsetValue(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_level_bar_remove_offset_value(self.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetInverted is a wrapper around gtk_level_bar_set_inverted().
func (self LevelBar) SetInverted(inverted bool) {
	C.gtk_level_bar_set_inverted(self.native(), C.gboolean(util.Bool2Int(inverted)) /*go:.util*/)
}

// SetMaxValue is a wrapper around gtk_level_bar_set_max_value().
func (self LevelBar) SetMaxValue(value float64) {
	C.gtk_level_bar_set_max_value(self.native(), C.gdouble(value))
}

// SetMinValue is a wrapper around gtk_level_bar_set_min_value().
func (self LevelBar) SetMinValue(value float64) {
	C.gtk_level_bar_set_min_value(self.native(), C.gdouble(value))
}

// SetMode is a wrapper around gtk_level_bar_set_mode().
func (self LevelBar) SetMode(mode LevelBarMode) {
	C.gtk_level_bar_set_mode(self.native(), C.GtkLevelBarMode(mode))
}

// SetValue is a wrapper around gtk_level_bar_set_value().
func (self LevelBar) SetValue(value float64) {
	C.gtk_level_bar_set_value(self.native(), C.gdouble(value))
}

// Object LevelBarAccessible
type LevelBarAccessible struct {
	WidgetAccessible
}

func (v LevelBarAccessible) native() *C.GtkLevelBarAccessible {
	return (*C.GtkLevelBarAccessible)(v.Ptr)
}
func wrapLevelBarAccessible(p *C.GtkLevelBarAccessible) (v LevelBarAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLevelBarAccessible(p unsafe.Pointer) (v LevelBarAccessible) {
	v.Ptr = p
	return
}
func (v LevelBarAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLevelBarAccessible(p unsafe.Pointer) interface{} {
	return WrapLevelBarAccessible(p)
}
func (v LevelBarAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_level_bar_accessible_get_type())
}
func (v LevelBarAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLevelBarAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v LevelBarAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v LevelBarAccessible) Value() atk.Value {
	return atk.WrapValue(v.Ptr) /*gir:Atk*/
}

// Object LinkButton
type LinkButton struct {
	Button
}

func (v LinkButton) native() *C.GtkLinkButton {
	return (*C.GtkLinkButton)(v.Ptr)
}
func wrapLinkButton(p *C.GtkLinkButton) (v LinkButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLinkButton(p unsafe.Pointer) (v LinkButton) {
	v.Ptr = p
	return
}
func (v LinkButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLinkButton(p unsafe.Pointer) interface{} {
	return WrapLinkButton(p)
}
func (v LinkButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_link_button_get_type())
}
func (v LinkButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLinkButton(unsafe.Pointer(ptr)), nil
	}
}
func (v LinkButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v LinkButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v LinkButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v LinkButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// LinkButtonNew is a wrapper around gtk_link_button_new().
func LinkButtonNew(uri string) Widget {
	uri0 := (*C.gchar)(C.CString(uri))
	ret0 := C.gtk_link_button_new(uri0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// LinkButtonNewWithLabel is a wrapper around gtk_link_button_new_with_label().
func LinkButtonNewWithLabel(uri string, label string) Widget {
	uri0 := (*C.gchar)(C.CString(uri))
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_link_button_new_with_label(uri0, label0)
	C.free(unsafe.Pointer(uri0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetUri is a wrapper around gtk_link_button_get_uri().
func (link_button LinkButton) GetUri() string {
	ret0 := C.gtk_link_button_get_uri(link_button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetVisited is a wrapper around gtk_link_button_get_visited().
func (link_button LinkButton) GetVisited() bool {
	ret0 := C.gtk_link_button_get_visited(link_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUri is a wrapper around gtk_link_button_set_uri().
func (link_button LinkButton) SetUri(uri string) {
	uri0 := (*C.gchar)(C.CString(uri))
	C.gtk_link_button_set_uri(link_button.native(), uri0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
}

// SetVisited is a wrapper around gtk_link_button_set_visited().
func (link_button LinkButton) SetVisited(visited bool) {
	C.gtk_link_button_set_visited(link_button.native(), C.gboolean(util.Bool2Int(visited)) /*go:.util*/)
}

// Object LinkButtonAccessible
type LinkButtonAccessible struct {
	ButtonAccessible
}

func (v LinkButtonAccessible) native() *C.GtkLinkButtonAccessible {
	return (*C.GtkLinkButtonAccessible)(v.Ptr)
}
func wrapLinkButtonAccessible(p *C.GtkLinkButtonAccessible) (v LinkButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLinkButtonAccessible(p unsafe.Pointer) (v LinkButtonAccessible) {
	v.Ptr = p
	return
}
func (v LinkButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLinkButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapLinkButtonAccessible(p)
}
func (v LinkButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_link_button_accessible_get_type())
}
func (v LinkButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLinkButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v LinkButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v LinkButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v LinkButtonAccessible) HyperlinkImpl() atk.HyperlinkImpl {
	return atk.WrapHyperlinkImpl(v.Ptr) /*gir:Atk*/
}
func (v LinkButtonAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object ListBox
type ListBox struct {
	Container
}

func (v ListBox) native() *C.GtkListBox {
	return (*C.GtkListBox)(v.Ptr)
}
func wrapListBox(p *C.GtkListBox) (v ListBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapListBox(p unsafe.Pointer) (v ListBox) {
	v.Ptr = p
	return
}
func (v ListBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapListBox(p unsafe.Pointer) interface{} {
	return WrapListBox(p)
}
func (v ListBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_list_box_get_type())
}
func (v ListBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapListBox(unsafe.Pointer(ptr)), nil
	}
}
func (v ListBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ListBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ListBoxNew is a wrapper around gtk_list_box_new().
func ListBoxNew() Widget {
	ret0 := C.gtk_list_box_new()
	return wrapWidget(ret0)
}

// DragHighlightRow is a wrapper around gtk_list_box_drag_highlight_row().
func (box ListBox) DragHighlightRow(row ListBoxRow) {
	C.gtk_list_box_drag_highlight_row(box.native(), row.native())
}

// DragUnhighlightRow is a wrapper around gtk_list_box_drag_unhighlight_row().
func (box ListBox) DragUnhighlightRow() {
	C.gtk_list_box_drag_unhighlight_row(box.native())
}

// GetActivateOnSingleClick is a wrapper around gtk_list_box_get_activate_on_single_click().
func (box ListBox) GetActivateOnSingleClick() bool {
	ret0 := C.gtk_list_box_get_activate_on_single_click(box.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetAdjustment is a wrapper around gtk_list_box_get_adjustment().
func (box ListBox) GetAdjustment() Adjustment {
	ret0 := C.gtk_list_box_get_adjustment(box.native())
	return wrapAdjustment(ret0)
}

// GetRowAtIndex is a wrapper around gtk_list_box_get_row_at_index().
func (box ListBox) GetRowAtIndex(index_ int) ListBoxRow {
	ret0 := C.gtk_list_box_get_row_at_index(box.native(), C.gint(index_))
	return wrapListBoxRow(ret0)
}

// GetRowAtY is a wrapper around gtk_list_box_get_row_at_y().
func (box ListBox) GetRowAtY(y int) ListBoxRow {
	ret0 := C.gtk_list_box_get_row_at_y(box.native(), C.gint(y))
	return wrapListBoxRow(ret0)
}

// GetSelectedRow is a wrapper around gtk_list_box_get_selected_row().
func (box ListBox) GetSelectedRow() ListBoxRow {
	ret0 := C.gtk_list_box_get_selected_row(box.native())
	return wrapListBoxRow(ret0)
}

// GetSelectedRows is a wrapper around gtk_list_box_get_selected_rows().
func (box ListBox) GetSelectedRows() glib.List {
	ret0 := C.gtk_list_box_get_selected_rows(box.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapListBoxRow(p) }) /*gir:GLib*/
}

// GetSelectionMode is a wrapper around gtk_list_box_get_selection_mode().
func (box ListBox) GetSelectionMode() SelectionMode {
	ret0 := C.gtk_list_box_get_selection_mode(box.native())
	return SelectionMode(ret0)
}

// Insert is a wrapper around gtk_list_box_insert().
func (box ListBox) Insert(child Widget, position int) {
	C.gtk_list_box_insert(box.native(), child.native(), C.gint(position))
}

// InvalidateFilter is a wrapper around gtk_list_box_invalidate_filter().
func (box ListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter(box.native())
}

// InvalidateHeaders is a wrapper around gtk_list_box_invalidate_headers().
func (box ListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers(box.native())
}

// InvalidateSort is a wrapper around gtk_list_box_invalidate_sort().
func (box ListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort(box.native())
}

// Prepend is a wrapper around gtk_list_box_prepend().
func (box ListBox) Prepend(child Widget) {
	C.gtk_list_box_prepend(box.native(), child.native())
}

// SelectAll is a wrapper around gtk_list_box_select_all().
func (box ListBox) SelectAll() {
	C.gtk_list_box_select_all(box.native())
}

// SelectRow is a wrapper around gtk_list_box_select_row().
func (box ListBox) SelectRow(row ListBoxRow) {
	C.gtk_list_box_select_row(box.native(), row.native())
}

// SetActivateOnSingleClick is a wrapper around gtk_list_box_set_activate_on_single_click().
func (box ListBox) SetActivateOnSingleClick(single bool) {
	C.gtk_list_box_set_activate_on_single_click(box.native(), C.gboolean(util.Bool2Int(single)) /*go:.util*/)
}

// SetAdjustment is a wrapper around gtk_list_box_set_adjustment().
func (box ListBox) SetAdjustment(adjustment Adjustment) {
	C.gtk_list_box_set_adjustment(box.native(), adjustment.native())
}

// SetPlaceholder is a wrapper around gtk_list_box_set_placeholder().
func (box ListBox) SetPlaceholder(placeholder Widget) {
	C.gtk_list_box_set_placeholder(box.native(), placeholder.native())
}

// SetSelectionMode is a wrapper around gtk_list_box_set_selection_mode().
func (box ListBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_list_box_set_selection_mode(box.native(), C.GtkSelectionMode(mode))
}

// UnselectAll is a wrapper around gtk_list_box_unselect_all().
func (box ListBox) UnselectAll() {
	C.gtk_list_box_unselect_all(box.native())
}

// UnselectRow is a wrapper around gtk_list_box_unselect_row().
func (box ListBox) UnselectRow(row ListBoxRow) {
	C.gtk_list_box_unselect_row(box.native(), row.native())
}

// Object ListBoxRow
type ListBoxRow struct {
	Bin
}

func (v ListBoxRow) native() *C.GtkListBoxRow {
	return (*C.GtkListBoxRow)(v.Ptr)
}
func wrapListBoxRow(p *C.GtkListBoxRow) (v ListBoxRow) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapListBoxRow(p unsafe.Pointer) (v ListBoxRow) {
	v.Ptr = p
	return
}
func (v ListBoxRow) IsNil() bool {
	return v.Ptr == nil
}
func IWrapListBoxRow(p unsafe.Pointer) interface{} {
	return WrapListBoxRow(p)
}
func (v ListBoxRow) GetType() gobject.Type {
	return gobject.Type(C.gtk_list_box_row_get_type())
}
func (v ListBoxRow) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapListBoxRow(unsafe.Pointer(ptr)), nil
	}
}
func (v ListBoxRow) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ListBoxRow) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ListBoxRowNew is a wrapper around gtk_list_box_row_new().
func ListBoxRowNew() Widget {
	ret0 := C.gtk_list_box_row_new()
	return wrapWidget(ret0)
}

// Changed is a wrapper around gtk_list_box_row_changed().
func (row ListBoxRow) Changed() {
	C.gtk_list_box_row_changed(row.native())
}

// GetActivatable is a wrapper around gtk_list_box_row_get_activatable().
func (row ListBoxRow) GetActivatable() bool {
	ret0 := C.gtk_list_box_row_get_activatable(row.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHeader is a wrapper around gtk_list_box_row_get_header().
func (row ListBoxRow) GetHeader() Widget {
	ret0 := C.gtk_list_box_row_get_header(row.native())
	return wrapWidget(ret0)
}

// GetIndex is a wrapper around gtk_list_box_row_get_index().
func (row ListBoxRow) GetIndex() int {
	ret0 := C.gtk_list_box_row_get_index(row.native())
	return int(ret0)
}

// GetSelectable is a wrapper around gtk_list_box_row_get_selectable().
func (row ListBoxRow) GetSelectable() bool {
	ret0 := C.gtk_list_box_row_get_selectable(row.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsSelected is a wrapper around gtk_list_box_row_is_selected().
func (row ListBoxRow) IsSelected() bool {
	ret0 := C.gtk_list_box_row_is_selected(row.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetActivatable is a wrapper around gtk_list_box_row_set_activatable().
func (row ListBoxRow) SetActivatable(activatable bool) {
	C.gtk_list_box_row_set_activatable(row.native(), C.gboolean(util.Bool2Int(activatable)) /*go:.util*/)
}

// SetHeader is a wrapper around gtk_list_box_row_set_header().
func (row ListBoxRow) SetHeader(header Widget) {
	C.gtk_list_box_row_set_header(row.native(), header.native())
}

// SetSelectable is a wrapper around gtk_list_box_row_set_selectable().
func (row ListBoxRow) SetSelectable(selectable bool) {
	C.gtk_list_box_row_set_selectable(row.native(), C.gboolean(util.Bool2Int(selectable)) /*go:.util*/)
}

// Object ListBoxAccessible
type ListBoxAccessible struct {
	ContainerAccessible
}

func (v ListBoxAccessible) native() *C.GtkListBoxAccessible {
	return (*C.GtkListBoxAccessible)(v.Ptr)
}
func wrapListBoxAccessible(p *C.GtkListBoxAccessible) (v ListBoxAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapListBoxAccessible(p unsafe.Pointer) (v ListBoxAccessible) {
	v.Ptr = p
	return
}
func (v ListBoxAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapListBoxAccessible(p unsafe.Pointer) interface{} {
	return WrapListBoxAccessible(p)
}
func (v ListBoxAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_list_box_accessible_get_type())
}
func (v ListBoxAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapListBoxAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ListBoxAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ListBoxAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object ListBoxRowAccessible
type ListBoxRowAccessible struct {
	ContainerAccessible
}

func (v ListBoxRowAccessible) native() *C.GtkListBoxRowAccessible {
	return (*C.GtkListBoxRowAccessible)(v.Ptr)
}
func wrapListBoxRowAccessible(p *C.GtkListBoxRowAccessible) (v ListBoxRowAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapListBoxRowAccessible(p unsafe.Pointer) (v ListBoxRowAccessible) {
	v.Ptr = p
	return
}
func (v ListBoxRowAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapListBoxRowAccessible(p unsafe.Pointer) interface{} {
	return WrapListBoxRowAccessible(p)
}
func (v ListBoxRowAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_list_box_row_accessible_get_type())
}
func (v ListBoxRowAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapListBoxRowAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ListBoxRowAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object ListStore
type ListStore struct {
	gobject.Object
}

func (v ListStore) native() *C.GtkListStore {
	return (*C.GtkListStore)(v.Ptr)
}
func wrapListStore(p *C.GtkListStore) (v ListStore) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapListStore(p unsafe.Pointer) (v ListStore) {
	v.Ptr = p
	return
}
func (v ListStore) IsNil() bool {
	return v.Ptr == nil
}
func IWrapListStore(p unsafe.Pointer) interface{} {
	return WrapListStore(p)
}
func (v ListStore) GetType() gobject.Type {
	return gobject.Type(C.gtk_list_store_get_type())
}
func (v ListStore) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapListStore(unsafe.Pointer(ptr)), nil
	}
}
func (v ListStore) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ListStore) TreeDragDest() TreeDragDest {
	return WrapTreeDragDest(v.Ptr)
}
func (v ListStore) TreeDragSource() TreeDragSource {
	return WrapTreeDragSource(v.Ptr)
}
func (v ListStore) TreeModel() TreeModel {
	return WrapTreeModel(v.Ptr)
}
func (v ListStore) TreeSortable() TreeSortable {
	return WrapTreeSortable(v.Ptr)
}

// gtk_list_store_new shadowed by newv

// ListStoreNew is a wrapper around gtk_list_store_newv().
func ListStoreNew(types [] /*Gir:GObject*/ gobject.Type) ListStore {
	types0 := make([]C.GType, len(types))
	for idx, elemG := range types {
		types0[idx] = C.GType(elemG)
	}
	var types0Ptr *C.GType
	if len(types0) > 0 {
		types0Ptr = &types0[0]
	}
	ret0 := C.gtk_list_store_newv(C.gint(len(types)), types0Ptr)
	return wrapListStore(ret0)
}

// Clear is a wrapper around gtk_list_store_clear().
func (list_store ListStore) Clear() {
	C.gtk_list_store_clear(list_store.native())
}

// IterIsValid is a wrapper around gtk_list_store_iter_is_valid().
func (list_store ListStore) IterIsValid(iter TreeIter) bool {
	ret0 := C.gtk_list_store_iter_is_valid(list_store.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MoveAfter is a wrapper around gtk_list_store_move_after().
func (store ListStore) MoveAfter(iter TreeIter, position TreeIter) {
	C.gtk_list_store_move_after(store.native(), iter.native(), position.native())
}

// MoveBefore is a wrapper around gtk_list_store_move_before().
func (store ListStore) MoveBefore(iter TreeIter, position TreeIter) {
	C.gtk_list_store_move_before(store.native(), iter.native(), position.native())
}

// Remove is a wrapper around gtk_list_store_remove().
func (list_store ListStore) Remove(iter TreeIter) bool {
	ret0 := C.gtk_list_store_remove(list_store.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Reorder is a wrapper around gtk_list_store_reorder().
func (store ListStore) Reorder(new_order []int) {
	new_order0 := make([]C.gint, len(new_order))
	for idx, elemG := range new_order {
		new_order0[idx] = C.gint(elemG)
	}
	var new_order0Ptr *C.gint
	if len(new_order0) > 0 {
		new_order0Ptr = &new_order0[0]
	}
	C.gtk_list_store_reorder(store.native(), new_order0Ptr)
}

// gtk_list_store_set shadowed by set_valuesv

// SetColumnTypes is a wrapper around gtk_list_store_set_column_types().
func (list_store ListStore) SetColumnTypes(types [] /*Gir:GObject*/ gobject.Type) {
	types0 := make([]C.GType, len(types))
	for idx, elemG := range types {
		types0[idx] = C.GType(elemG)
	}
	var types0Ptr *C.GType
	if len(types0) > 0 {
		types0Ptr = &types0[0]
	}
	C.gtk_list_store_set_column_types(list_store.native(), C.gint(len(types)), types0Ptr)
}

// SetValue is a wrapper around gtk_list_store_set_value().
func (list_store ListStore) SetValue(iter TreeIter, column int, value gobject.Value) {
	C.gtk_list_store_set_value(list_store.native(), iter.native(), C.gint(column), (*C.GValue)(value.Ptr))
}

// Swap is a wrapper around gtk_list_store_swap().
func (store ListStore) Swap(a TreeIter, b TreeIter) {
	C.gtk_list_store_swap(store.native(), a.native(), b.native())
}

// Object LockButton
type LockButton struct {
	Button
}

func (v LockButton) native() *C.GtkLockButton {
	return (*C.GtkLockButton)(v.Ptr)
}
func wrapLockButton(p *C.GtkLockButton) (v LockButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLockButton(p unsafe.Pointer) (v LockButton) {
	v.Ptr = p
	return
}
func (v LockButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLockButton(p unsafe.Pointer) interface{} {
	return WrapLockButton(p)
}
func (v LockButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_lock_button_get_type())
}
func (v LockButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLockButton(unsafe.Pointer(ptr)), nil
	}
}
func (v LockButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v LockButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v LockButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v LockButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// LockButtonNew is a wrapper around gtk_lock_button_new().
func LockButtonNew(permission gio.Permission) Widget {
	ret0 := C.gtk_lock_button_new((*C.GPermission)(permission.Ptr))
	return wrapWidget(ret0)
}

// GetPermission is a wrapper around gtk_lock_button_get_permission().
func (button LockButton) GetPermission() gio.Permission {
	ret0 := C.gtk_lock_button_get_permission(button.native())
	return gio.WrapPermission(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// SetPermission is a wrapper around gtk_lock_button_set_permission().
func (button LockButton) SetPermission(permission gio.Permission) {
	C.gtk_lock_button_set_permission(button.native(), (*C.GPermission)(permission.Ptr))
}

// Object LockButtonAccessible
type LockButtonAccessible struct {
	ButtonAccessible
}

func (v LockButtonAccessible) native() *C.GtkLockButtonAccessible {
	return (*C.GtkLockButtonAccessible)(v.Ptr)
}
func wrapLockButtonAccessible(p *C.GtkLockButtonAccessible) (v LockButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapLockButtonAccessible(p unsafe.Pointer) (v LockButtonAccessible) {
	v.Ptr = p
	return
}
func (v LockButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLockButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapLockButtonAccessible(p)
}
func (v LockButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_lock_button_accessible_get_type())
}
func (v LockButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLockButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v LockButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v LockButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v LockButtonAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object MenuAccessible
type MenuAccessible struct {
	MenuShellAccessible
}

func (v MenuAccessible) native() *C.GtkMenuAccessible {
	return (*C.GtkMenuAccessible)(v.Ptr)
}
func wrapMenuAccessible(p *C.GtkMenuAccessible) (v MenuAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuAccessible(p unsafe.Pointer) (v MenuAccessible) {
	v.Ptr = p
	return
}
func (v MenuAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuAccessible(p unsafe.Pointer) interface{} {
	return WrapMenuAccessible(p)
}
func (v MenuAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_accessible_get_type())
}
func (v MenuAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v MenuAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object MenuShellAccessible
type MenuShellAccessible struct {
	ContainerAccessible
}

func (v MenuShellAccessible) native() *C.GtkMenuShellAccessible {
	return (*C.GtkMenuShellAccessible)(v.Ptr)
}
func wrapMenuShellAccessible(p *C.GtkMenuShellAccessible) (v MenuShellAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuShellAccessible(p unsafe.Pointer) (v MenuShellAccessible) {
	v.Ptr = p
	return
}
func (v MenuShellAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuShellAccessible(p unsafe.Pointer) interface{} {
	return WrapMenuShellAccessible(p)
}
func (v MenuShellAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_shell_accessible_get_type())
}
func (v MenuShellAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuShellAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuShellAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v MenuShellAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object MenuBar
type MenuBar struct {
	MenuShell
}

func (v MenuBar) native() *C.GtkMenuBar {
	return (*C.GtkMenuBar)(v.Ptr)
}
func wrapMenuBar(p *C.GtkMenuBar) (v MenuBar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuBar(p unsafe.Pointer) (v MenuBar) {
	v.Ptr = p
	return
}
func (v MenuBar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuBar(p unsafe.Pointer) interface{} {
	return WrapMenuBar(p)
}
func (v MenuBar) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_bar_get_type())
}
func (v MenuBar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuBar(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuBar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v MenuBar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// MenuBarNew is a wrapper around gtk_menu_bar_new().
func MenuBarNew() Widget {
	ret0 := C.gtk_menu_bar_new()
	return wrapWidget(ret0)
}

// MenuBarNewFromModel is a wrapper around gtk_menu_bar_new_from_model().
func MenuBarNewFromModel(model gio.MenuModel) Widget {
	ret0 := C.gtk_menu_bar_new_from_model((*C.GMenuModel)(model.Ptr))
	return wrapWidget(ret0)
}

// GetChildPackDirection is a wrapper around gtk_menu_bar_get_child_pack_direction().
func (menubar MenuBar) GetChildPackDirection() PackDirection {
	ret0 := C.gtk_menu_bar_get_child_pack_direction(menubar.native())
	return PackDirection(ret0)
}

// GetPackDirection is a wrapper around gtk_menu_bar_get_pack_direction().
func (menubar MenuBar) GetPackDirection() PackDirection {
	ret0 := C.gtk_menu_bar_get_pack_direction(menubar.native())
	return PackDirection(ret0)
}

// SetChildPackDirection is a wrapper around gtk_menu_bar_set_child_pack_direction().
func (menubar MenuBar) SetChildPackDirection(child_pack_dir PackDirection) {
	C.gtk_menu_bar_set_child_pack_direction(menubar.native(), C.GtkPackDirection(child_pack_dir))
}

// SetPackDirection is a wrapper around gtk_menu_bar_set_pack_direction().
func (menubar MenuBar) SetPackDirection(pack_dir PackDirection) {
	C.gtk_menu_bar_set_pack_direction(menubar.native(), C.GtkPackDirection(pack_dir))
}

// Object MenuButton
type MenuButton struct {
	ToggleButton
}

func (v MenuButton) native() *C.GtkMenuButton {
	return (*C.GtkMenuButton)(v.Ptr)
}
func wrapMenuButton(p *C.GtkMenuButton) (v MenuButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuButton(p unsafe.Pointer) (v MenuButton) {
	v.Ptr = p
	return
}
func (v MenuButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuButton(p unsafe.Pointer) interface{} {
	return WrapMenuButton(p)
}
func (v MenuButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_button_get_type())
}
func (v MenuButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuButton(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v MenuButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v MenuButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v MenuButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// MenuButtonNew is a wrapper around gtk_menu_button_new().
func MenuButtonNew() Widget {
	ret0 := C.gtk_menu_button_new()
	return wrapWidget(ret0)
}

// GetAlignWidget is a wrapper around gtk_menu_button_get_align_widget().
func (menu_button MenuButton) GetAlignWidget() Widget {
	ret0 := C.gtk_menu_button_get_align_widget(menu_button.native())
	return wrapWidget(ret0)
}

// GetDirection is a wrapper around gtk_menu_button_get_direction().
func (menu_button MenuButton) GetDirection() ArrowType {
	ret0 := C.gtk_menu_button_get_direction(menu_button.native())
	return ArrowType(ret0)
}

// GetMenuModel is a wrapper around gtk_menu_button_get_menu_model().
func (menu_button MenuButton) GetMenuModel() gio.MenuModel {
	ret0 := C.gtk_menu_button_get_menu_model(menu_button.native())
	return gio.WrapMenuModel(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetPopover is a wrapper around gtk_menu_button_get_popover().
func (menu_button MenuButton) GetPopover() Popover {
	ret0 := C.gtk_menu_button_get_popover(menu_button.native())
	return wrapPopover(ret0)
}

// GetPopup is a wrapper around gtk_menu_button_get_popup().
func (menu_button MenuButton) GetPopup() Menu {
	ret0 := C.gtk_menu_button_get_popup(menu_button.native())
	return wrapMenu(ret0)
}

// GetUsePopover is a wrapper around gtk_menu_button_get_use_popover().
func (menu_button MenuButton) GetUsePopover() bool {
	ret0 := C.gtk_menu_button_get_use_popover(menu_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetAlignWidget is a wrapper around gtk_menu_button_set_align_widget().
func (menu_button MenuButton) SetAlignWidget(align_widget Widget) {
	C.gtk_menu_button_set_align_widget(menu_button.native(), align_widget.native())
}

// SetDirection is a wrapper around gtk_menu_button_set_direction().
func (menu_button MenuButton) SetDirection(direction ArrowType) {
	C.gtk_menu_button_set_direction(menu_button.native(), C.GtkArrowType(direction))
}

// SetMenuModel is a wrapper around gtk_menu_button_set_menu_model().
func (menu_button MenuButton) SetMenuModel(menu_model gio.MenuModel) {
	C.gtk_menu_button_set_menu_model(menu_button.native(), (*C.GMenuModel)(menu_model.Ptr))
}

// SetPopover is a wrapper around gtk_menu_button_set_popover().
func (menu_button MenuButton) SetPopover(popover Widget) {
	C.gtk_menu_button_set_popover(menu_button.native(), popover.native())
}

// SetPopup is a wrapper around gtk_menu_button_set_popup().
func (menu_button MenuButton) SetPopup(menu Widget) {
	C.gtk_menu_button_set_popup(menu_button.native(), menu.native())
}

// SetUsePopover is a wrapper around gtk_menu_button_set_use_popover().
func (menu_button MenuButton) SetUsePopover(use_popover bool) {
	C.gtk_menu_button_set_use_popover(menu_button.native(), C.gboolean(util.Bool2Int(use_popover)) /*go:.util*/)
}

// Object Popover
type Popover struct {
	Bin
}

func (v Popover) native() *C.GtkPopover {
	return (*C.GtkPopover)(v.Ptr)
}
func wrapPopover(p *C.GtkPopover) (v Popover) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPopover(p unsafe.Pointer) (v Popover) {
	v.Ptr = p
	return
}
func (v Popover) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPopover(p unsafe.Pointer) interface{} {
	return WrapPopover(p)
}
func (v Popover) GetType() gobject.Type {
	return gobject.Type(C.gtk_popover_get_type())
}
func (v Popover) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPopover(unsafe.Pointer(ptr)), nil
	}
}
func (v Popover) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Popover) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// PopoverNew is a wrapper around gtk_popover_new().
func PopoverNew(relative_to Widget) Widget {
	ret0 := C.gtk_popover_new(relative_to.native())
	return wrapWidget(ret0)
}

// PopoverNewFromModel is a wrapper around gtk_popover_new_from_model().
func PopoverNewFromModel(relative_to Widget, model gio.MenuModel) Widget {
	ret0 := C.gtk_popover_new_from_model(relative_to.native(), (*C.GMenuModel)(model.Ptr))
	return wrapWidget(ret0)
}

// BindModel is a wrapper around gtk_popover_bind_model().
func (popover Popover) BindModel(model gio.MenuModel, action_namespace string) {
	action_namespace0 := (*C.gchar)(C.CString(action_namespace))
	C.gtk_popover_bind_model(popover.native(), (*C.GMenuModel)(model.Ptr), action_namespace0)
	C.free(unsafe.Pointer(action_namespace0)) /*ch:<stdlib.h>*/
}

// GetConstrainTo is a wrapper around gtk_popover_get_constrain_to().
func (popover Popover) GetConstrainTo() PopoverConstraint {
	ret0 := C.gtk_popover_get_constrain_to(popover.native())
	return PopoverConstraint(ret0)
}

// GetDefaultWidget is a wrapper around gtk_popover_get_default_widget().
func (popover Popover) GetDefaultWidget() Widget {
	ret0 := C.gtk_popover_get_default_widget(popover.native())
	return wrapWidget(ret0)
}

// GetModal is a wrapper around gtk_popover_get_modal().
func (popover Popover) GetModal() bool {
	ret0 := C.gtk_popover_get_modal(popover.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPosition is a wrapper around gtk_popover_get_position().
func (popover Popover) GetPosition() PositionType {
	ret0 := C.gtk_popover_get_position(popover.native())
	return PositionType(ret0)
}

// GetRelativeTo is a wrapper around gtk_popover_get_relative_to().
func (popover Popover) GetRelativeTo() Widget {
	ret0 := C.gtk_popover_get_relative_to(popover.native())
	return wrapWidget(ret0)
}

// Popdown is a wrapper around gtk_popover_popdown().
func (popover Popover) Popdown() {
	C.gtk_popover_popdown(popover.native())
}

// Popup is a wrapper around gtk_popover_popup().
func (popover Popover) Popup() {
	C.gtk_popover_popup(popover.native())
}

// SetConstrainTo is a wrapper around gtk_popover_set_constrain_to().
func (popover Popover) SetConstrainTo(constraint PopoverConstraint) {
	C.gtk_popover_set_constrain_to(popover.native(), C.GtkPopoverConstraint(constraint))
}

// SetDefaultWidget is a wrapper around gtk_popover_set_default_widget().
func (popover Popover) SetDefaultWidget(widget Widget) {
	C.gtk_popover_set_default_widget(popover.native(), widget.native())
}

// SetModal is a wrapper around gtk_popover_set_modal().
func (popover Popover) SetModal(modal bool) {
	C.gtk_popover_set_modal(popover.native(), C.gboolean(util.Bool2Int(modal)) /*go:.util*/)
}

// SetPosition is a wrapper around gtk_popover_set_position().
func (popover Popover) SetPosition(position PositionType) {
	C.gtk_popover_set_position(popover.native(), C.GtkPositionType(position))
}

// SetRelativeTo is a wrapper around gtk_popover_set_relative_to().
func (popover Popover) SetRelativeTo(relative_to Widget) {
	C.gtk_popover_set_relative_to(popover.native(), relative_to.native())
}

// Object MenuButtonAccessible
type MenuButtonAccessible struct {
	ToggleButtonAccessible
}

func (v MenuButtonAccessible) native() *C.GtkMenuButtonAccessible {
	return (*C.GtkMenuButtonAccessible)(v.Ptr)
}
func wrapMenuButtonAccessible(p *C.GtkMenuButtonAccessible) (v MenuButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuButtonAccessible(p unsafe.Pointer) (v MenuButtonAccessible) {
	v.Ptr = p
	return
}
func (v MenuButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapMenuButtonAccessible(p)
}
func (v MenuButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_button_accessible_get_type())
}
func (v MenuButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v MenuButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v MenuButtonAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object ToggleButtonAccessible
type ToggleButtonAccessible struct {
	ButtonAccessible
}

func (v ToggleButtonAccessible) native() *C.GtkToggleButtonAccessible {
	return (*C.GtkToggleButtonAccessible)(v.Ptr)
}
func wrapToggleButtonAccessible(p *C.GtkToggleButtonAccessible) (v ToggleButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToggleButtonAccessible(p unsafe.Pointer) (v ToggleButtonAccessible) {
	v.Ptr = p
	return
}
func (v ToggleButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToggleButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapToggleButtonAccessible(p)
}
func (v ToggleButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_toggle_button_accessible_get_type())
}
func (v ToggleButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToggleButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ToggleButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v ToggleButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ToggleButtonAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object ToolItem
type ToolItem struct {
	Bin
}

func (v ToolItem) native() *C.GtkToolItem {
	return (*C.GtkToolItem)(v.Ptr)
}
func wrapToolItem(p *C.GtkToolItem) (v ToolItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToolItem(p unsafe.Pointer) (v ToolItem) {
	v.Ptr = p
	return
}
func (v ToolItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToolItem(p unsafe.Pointer) interface{} {
	return WrapToolItem(p)
}
func (v ToolItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_tool_item_get_type())
}
func (v ToolItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToolItem(unsafe.Pointer(ptr)), nil
	}
}
func (v ToolItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ToolItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ToolItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ToolItemNew is a wrapper around gtk_tool_item_new().
func ToolItemNew() ToolItem {
	ret0 := C.gtk_tool_item_new()
	return wrapToolItem(ret0)
}

// GetEllipsizeMode is a wrapper around gtk_tool_item_get_ellipsize_mode().
func (tool_item ToolItem) GetEllipsizeMode() /*gir:Pango*/ pango.EllipsizeMode {
	ret0 := C.gtk_tool_item_get_ellipsize_mode(tool_item.native())
	return /*gir:Pango*/ pango.EllipsizeMode(ret0)
}

// GetExpand is a wrapper around gtk_tool_item_get_expand().
func (tool_item ToolItem) GetExpand() bool {
	ret0 := C.gtk_tool_item_get_expand(tool_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHomogeneous is a wrapper around gtk_tool_item_get_homogeneous().
func (tool_item ToolItem) GetHomogeneous() bool {
	ret0 := C.gtk_tool_item_get_homogeneous(tool_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsImportant is a wrapper around gtk_tool_item_get_is_important().
func (tool_item ToolItem) GetIsImportant() bool {
	ret0 := C.gtk_tool_item_get_is_important(tool_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetOrientation is a wrapper around gtk_tool_item_get_orientation().
func (tool_item ToolItem) GetOrientation() Orientation {
	ret0 := C.gtk_tool_item_get_orientation(tool_item.native())
	return Orientation(ret0)
}

// GetProxyMenuItem is a wrapper around gtk_tool_item_get_proxy_menu_item().
func (tool_item ToolItem) GetProxyMenuItem(menu_item_id string) Widget {
	menu_item_id0 := (*C.gchar)(C.CString(menu_item_id))
	ret0 := C.gtk_tool_item_get_proxy_menu_item(tool_item.native(), menu_item_id0)
	C.free(unsafe.Pointer(menu_item_id0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetReliefStyle is a wrapper around gtk_tool_item_get_relief_style().
func (tool_item ToolItem) GetReliefStyle() ReliefStyle {
	ret0 := C.gtk_tool_item_get_relief_style(tool_item.native())
	return ReliefStyle(ret0)
}

// GetTextAlignment is a wrapper around gtk_tool_item_get_text_alignment().
func (tool_item ToolItem) GetTextAlignment() float32 {
	ret0 := C.gtk_tool_item_get_text_alignment(tool_item.native())
	return float32(ret0)
}

// GetTextOrientation is a wrapper around gtk_tool_item_get_text_orientation().
func (tool_item ToolItem) GetTextOrientation() Orientation {
	ret0 := C.gtk_tool_item_get_text_orientation(tool_item.native())
	return Orientation(ret0)
}

// GetTextSizeGroup is a wrapper around gtk_tool_item_get_text_size_group().
func (tool_item ToolItem) GetTextSizeGroup() SizeGroup {
	ret0 := C.gtk_tool_item_get_text_size_group(tool_item.native())
	return wrapSizeGroup(ret0)
}

// GetToolbarStyle is a wrapper around gtk_tool_item_get_toolbar_style().
func (tool_item ToolItem) GetToolbarStyle() ToolbarStyle {
	ret0 := C.gtk_tool_item_get_toolbar_style(tool_item.native())
	return ToolbarStyle(ret0)
}

// GetUseDragWindow is a wrapper around gtk_tool_item_get_use_drag_window().
func (tool_item ToolItem) GetUseDragWindow() bool {
	ret0 := C.gtk_tool_item_get_use_drag_window(tool_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVisibleHorizontal is a wrapper around gtk_tool_item_get_visible_horizontal().
func (tool_item ToolItem) GetVisibleHorizontal() bool {
	ret0 := C.gtk_tool_item_get_visible_horizontal(tool_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVisibleVertical is a wrapper around gtk_tool_item_get_visible_vertical().
func (tool_item ToolItem) GetVisibleVertical() bool {
	ret0 := C.gtk_tool_item_get_visible_vertical(tool_item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RebuildMenu is a wrapper around gtk_tool_item_rebuild_menu().
func (tool_item ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu(tool_item.native())
}

// RetrieveProxyMenuItem is a wrapper around gtk_tool_item_retrieve_proxy_menu_item().
func (tool_item ToolItem) RetrieveProxyMenuItem() Widget {
	ret0 := C.gtk_tool_item_retrieve_proxy_menu_item(tool_item.native())
	return wrapWidget(ret0)
}

// SetExpand is a wrapper around gtk_tool_item_set_expand().
func (tool_item ToolItem) SetExpand(expand bool) {
	C.gtk_tool_item_set_expand(tool_item.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// SetHomogeneous is a wrapper around gtk_tool_item_set_homogeneous().
func (tool_item ToolItem) SetHomogeneous(homogeneous bool) {
	C.gtk_tool_item_set_homogeneous(tool_item.native(), C.gboolean(util.Bool2Int(homogeneous)) /*go:.util*/)
}

// SetIsImportant is a wrapper around gtk_tool_item_set_is_important().
func (tool_item ToolItem) SetIsImportant(is_important bool) {
	C.gtk_tool_item_set_is_important(tool_item.native(), C.gboolean(util.Bool2Int(is_important)) /*go:.util*/)
}

// SetProxyMenuItem is a wrapper around gtk_tool_item_set_proxy_menu_item().
func (tool_item ToolItem) SetProxyMenuItem(menu_item_id string, menu_item Widget) {
	menu_item_id0 := (*C.gchar)(C.CString(menu_item_id))
	C.gtk_tool_item_set_proxy_menu_item(tool_item.native(), menu_item_id0, menu_item.native())
	C.free(unsafe.Pointer(menu_item_id0)) /*ch:<stdlib.h>*/
}

// SetTooltipMarkup is a wrapper around gtk_tool_item_set_tooltip_markup().
func (tool_item ToolItem) SetTooltipMarkup(markup string) {
	markup0 := (*C.gchar)(C.CString(markup))
	C.gtk_tool_item_set_tooltip_markup(tool_item.native(), markup0)
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// SetTooltipText is a wrapper around gtk_tool_item_set_tooltip_text().
func (tool_item ToolItem) SetTooltipText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_tool_item_set_tooltip_text(tool_item.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetUseDragWindow is a wrapper around gtk_tool_item_set_use_drag_window().
func (tool_item ToolItem) SetUseDragWindow(use_drag_window bool) {
	C.gtk_tool_item_set_use_drag_window(tool_item.native(), C.gboolean(util.Bool2Int(use_drag_window)) /*go:.util*/)
}

// SetVisibleHorizontal is a wrapper around gtk_tool_item_set_visible_horizontal().
func (tool_item ToolItem) SetVisibleHorizontal(visible_horizontal bool) {
	C.gtk_tool_item_set_visible_horizontal(tool_item.native(), C.gboolean(util.Bool2Int(visible_horizontal)) /*go:.util*/)
}

// SetVisibleVertical is a wrapper around gtk_tool_item_set_visible_vertical().
func (tool_item ToolItem) SetVisibleVertical(visible_vertical bool) {
	C.gtk_tool_item_set_visible_vertical(tool_item.native(), C.gboolean(util.Bool2Int(visible_vertical)) /*go:.util*/)
}

// ToolbarReconfigured is a wrapper around gtk_tool_item_toolbar_reconfigured().
func (tool_item ToolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured(tool_item.native())
}

// Object ToolButton
type ToolButton struct {
	ToolItem
}

func (v ToolButton) native() *C.GtkToolButton {
	return (*C.GtkToolButton)(v.Ptr)
}
func wrapToolButton(p *C.GtkToolButton) (v ToolButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToolButton(p unsafe.Pointer) (v ToolButton) {
	v.Ptr = p
	return
}
func (v ToolButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToolButton(p unsafe.Pointer) interface{} {
	return WrapToolButton(p)
}
func (v ToolButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_tool_button_get_type())
}
func (v ToolButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToolButton(unsafe.Pointer(ptr)), nil
	}
}
func (v ToolButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ToolButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v ToolButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ToolButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ToolButtonNew is a wrapper around gtk_tool_button_new().
func ToolButtonNew(icon_widget Widget, label string) ToolItem {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_tool_button_new(icon_widget.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapToolItem(ret0)
}

// GetIconName is a wrapper around gtk_tool_button_get_icon_name().
func (button ToolButton) GetIconName() string {
	ret0 := C.gtk_tool_button_get_icon_name(button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetIconWidget is a wrapper around gtk_tool_button_get_icon_widget().
func (button ToolButton) GetIconWidget() Widget {
	ret0 := C.gtk_tool_button_get_icon_widget(button.native())
	return wrapWidget(ret0)
}

// GetLabel is a wrapper around gtk_tool_button_get_label().
func (button ToolButton) GetLabel() string {
	ret0 := C.gtk_tool_button_get_label(button.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLabelWidget is a wrapper around gtk_tool_button_get_label_widget().
func (button ToolButton) GetLabelWidget() Widget {
	ret0 := C.gtk_tool_button_get_label_widget(button.native())
	return wrapWidget(ret0)
}

// GetUseUnderline is a wrapper around gtk_tool_button_get_use_underline().
func (button ToolButton) GetUseUnderline() bool {
	ret0 := C.gtk_tool_button_get_use_underline(button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetIconName is a wrapper around gtk_tool_button_set_icon_name().
func (button ToolButton) SetIconName(icon_name string) {
	icon_name0 := (*C.gchar)(C.CString(icon_name))
	C.gtk_tool_button_set_icon_name(button.native(), icon_name0)
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
}

// SetIconWidget is a wrapper around gtk_tool_button_set_icon_widget().
func (button ToolButton) SetIconWidget(icon_widget Widget) {
	C.gtk_tool_button_set_icon_widget(button.native(), icon_widget.native())
}

// SetLabel is a wrapper around gtk_tool_button_set_label().
func (button ToolButton) SetLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_tool_button_set_label(button.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetLabelWidget is a wrapper around gtk_tool_button_set_label_widget().
func (button ToolButton) SetLabelWidget(label_widget Widget) {
	C.gtk_tool_button_set_label_widget(button.native(), label_widget.native())
}

// SetUseUnderline is a wrapper around gtk_tool_button_set_use_underline().
func (button ToolButton) SetUseUnderline(use_underline bool) {
	C.gtk_tool_button_set_use_underline(button.native(), C.gboolean(util.Bool2Int(use_underline)) /*go:.util*/)
}

// Object MenuToolButton
type MenuToolButton struct {
	ToolButton
}

func (v MenuToolButton) native() *C.GtkMenuToolButton {
	return (*C.GtkMenuToolButton)(v.Ptr)
}
func wrapMenuToolButton(p *C.GtkMenuToolButton) (v MenuToolButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuToolButton(p unsafe.Pointer) (v MenuToolButton) {
	v.Ptr = p
	return
}
func (v MenuToolButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuToolButton(p unsafe.Pointer) interface{} {
	return WrapMenuToolButton(p)
}
func (v MenuToolButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_menu_tool_button_get_type())
}
func (v MenuToolButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuToolButton(unsafe.Pointer(ptr)), nil
	}
}
func (v MenuToolButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v MenuToolButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v MenuToolButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v MenuToolButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// MenuToolButtonNew is a wrapper around gtk_menu_tool_button_new().
func MenuToolButtonNew(icon_widget Widget, label string) ToolItem {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_menu_tool_button_new(icon_widget.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapToolItem(ret0)
}

// GetMenu is a wrapper around gtk_menu_tool_button_get_menu().
func (button MenuToolButton) GetMenu() Widget {
	ret0 := C.gtk_menu_tool_button_get_menu(button.native())
	return wrapWidget(ret0)
}

// SetArrowTooltipMarkup is a wrapper around gtk_menu_tool_button_set_arrow_tooltip_markup().
func (button MenuToolButton) SetArrowTooltipMarkup(markup string) {
	markup0 := (*C.gchar)(C.CString(markup))
	C.gtk_menu_tool_button_set_arrow_tooltip_markup(button.native(), markup0)
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// SetArrowTooltipText is a wrapper around gtk_menu_tool_button_set_arrow_tooltip_text().
func (button MenuToolButton) SetArrowTooltipText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_menu_tool_button_set_arrow_tooltip_text(button.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetMenu is a wrapper around gtk_menu_tool_button_set_menu().
func (button MenuToolButton) SetMenu(menu Widget) {
	C.gtk_menu_tool_button_set_menu(button.native(), menu.native())
}

// Object MessageDialog
type MessageDialog struct {
	Dialog
}

func (v MessageDialog) native() *C.GtkMessageDialog {
	return (*C.GtkMessageDialog)(v.Ptr)
}
func wrapMessageDialog(p *C.GtkMessageDialog) (v MessageDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMessageDialog(p unsafe.Pointer) (v MessageDialog) {
	v.Ptr = p
	return
}
func (v MessageDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMessageDialog(p unsafe.Pointer) interface{} {
	return WrapMessageDialog(p)
}
func (v MessageDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_message_dialog_get_type())
}
func (v MessageDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMessageDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v MessageDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v MessageDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// GetMessageArea is a wrapper around gtk_message_dialog_get_message_area().
func (message_dialog MessageDialog) GetMessageArea() Widget {
	ret0 := C.gtk_message_dialog_get_message_area(message_dialog.native())
	return wrapWidget(ret0)
}

// SetMarkup is a wrapper around gtk_message_dialog_set_markup().
func (message_dialog MessageDialog) SetMarkup(str string) {
	str0 := (*C.gchar)(C.CString(str))
	C.gtk_message_dialog_set_markup(message_dialog.native(), str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
}

// Object ModelButton
type ModelButton struct {
	Button
}

func (v ModelButton) native() *C.GtkModelButton {
	return (*C.GtkModelButton)(v.Ptr)
}
func wrapModelButton(p *C.GtkModelButton) (v ModelButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapModelButton(p unsafe.Pointer) (v ModelButton) {
	v.Ptr = p
	return
}
func (v ModelButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapModelButton(p unsafe.Pointer) interface{} {
	return WrapModelButton(p)
}
func (v ModelButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_model_button_get_type())
}
func (v ModelButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapModelButton(unsafe.Pointer(ptr)), nil
	}
}
func (v ModelButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ModelButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v ModelButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ModelButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ModelButtonNew is a wrapper around gtk_model_button_new().
func ModelButtonNew() Widget {
	ret0 := C.gtk_model_button_new()
	return wrapWidget(ret0)
}

// Object MountOperation
type MountOperation struct {
	gio.MountOperation
}

func (v MountOperation) native() *C.GtkMountOperation {
	return (*C.GtkMountOperation)(v.Ptr)
}
func wrapMountOperation(p *C.GtkMountOperation) (v MountOperation) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMountOperation(p unsafe.Pointer) (v MountOperation) {
	v.Ptr = p
	return
}
func (v MountOperation) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMountOperation(p unsafe.Pointer) interface{} {
	return WrapMountOperation(p)
}
func (v MountOperation) GetType() gobject.Type {
	return gobject.Type(C.gtk_mount_operation_get_type())
}
func (v MountOperation) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMountOperation(unsafe.Pointer(ptr)), nil
	}
}

// MountOperationNew is a wrapper around gtk_mount_operation_new().
func MountOperationNew(parent Window) gio.MountOperation {
	ret0 := C.gtk_mount_operation_new(parent.native())
	return gio.WrapMountOperation(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetParent is a wrapper around gtk_mount_operation_get_parent().
func (op MountOperation) GetParent() Window {
	ret0 := C.gtk_mount_operation_get_parent(op.native())
	return wrapWindow(ret0)
}

// IsShowing is a wrapper around gtk_mount_operation_is_showing().
func (op MountOperation) IsShowing() bool {
	ret0 := C.gtk_mount_operation_is_showing(op.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetParent is a wrapper around gtk_mount_operation_set_parent().
func (op MountOperation) SetParent(parent Window) {
	C.gtk_mount_operation_set_parent(op.native(), parent.native())
}

// Object Notebook
type Notebook struct {
	Container
}

func (v Notebook) native() *C.GtkNotebook {
	return (*C.GtkNotebook)(v.Ptr)
}
func wrapNotebook(p *C.GtkNotebook) (v Notebook) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNotebook(p unsafe.Pointer) (v Notebook) {
	v.Ptr = p
	return
}
func (v Notebook) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNotebook(p unsafe.Pointer) interface{} {
	return WrapNotebook(p)
}
func (v Notebook) GetType() gobject.Type {
	return gobject.Type(C.gtk_notebook_get_type())
}
func (v Notebook) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNotebook(unsafe.Pointer(ptr)), nil
	}
}
func (v Notebook) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Notebook) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// NotebookNew is a wrapper around gtk_notebook_new().
func NotebookNew() Widget {
	ret0 := C.gtk_notebook_new()
	return wrapWidget(ret0)
}

// AppendPage is a wrapper around gtk_notebook_append_page().
func (notebook Notebook) AppendPage(child Widget, tab_label Widget) int {
	ret0 := C.gtk_notebook_append_page(notebook.native(), child.native(), tab_label.native())
	return int(ret0)
}

// AppendPageMenu is a wrapper around gtk_notebook_append_page_menu().
func (notebook Notebook) AppendPageMenu(child Widget, tab_label Widget, menu_label Widget) int {
	ret0 := C.gtk_notebook_append_page_menu(notebook.native(), child.native(), tab_label.native(), menu_label.native())
	return int(ret0)
}

// DetachTab is a wrapper around gtk_notebook_detach_tab().
func (notebook Notebook) DetachTab(child Widget) {
	C.gtk_notebook_detach_tab(notebook.native(), child.native())
}

// GetActionWidget is a wrapper around gtk_notebook_get_action_widget().
func (notebook Notebook) GetActionWidget(pack_type PackType) Widget {
	ret0 := C.gtk_notebook_get_action_widget(notebook.native(), C.GtkPackType(pack_type))
	return wrapWidget(ret0)
}

// GetCurrentPage is a wrapper around gtk_notebook_get_current_page().
func (notebook Notebook) GetCurrentPage() int {
	ret0 := C.gtk_notebook_get_current_page(notebook.native())
	return int(ret0)
}

// GetGroupName is a wrapper around gtk_notebook_get_group_name().
func (notebook Notebook) GetGroupName() string {
	ret0 := C.gtk_notebook_get_group_name(notebook.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetMenuLabel is a wrapper around gtk_notebook_get_menu_label().
func (notebook Notebook) GetMenuLabel(child Widget) Widget {
	ret0 := C.gtk_notebook_get_menu_label(notebook.native(), child.native())
	return wrapWidget(ret0)
}

// GetMenuLabelText is a wrapper around gtk_notebook_get_menu_label_text().
func (notebook Notebook) GetMenuLabelText(child Widget) string {
	ret0 := C.gtk_notebook_get_menu_label_text(notebook.native(), child.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNPages is a wrapper around gtk_notebook_get_n_pages().
func (notebook Notebook) GetNPages() int {
	ret0 := C.gtk_notebook_get_n_pages(notebook.native())
	return int(ret0)
}

// GetNthPage is a wrapper around gtk_notebook_get_nth_page().
func (notebook Notebook) GetNthPage(page_num int) Widget {
	ret0 := C.gtk_notebook_get_nth_page(notebook.native(), C.gint(page_num))
	return wrapWidget(ret0)
}

// GetScrollable is a wrapper around gtk_notebook_get_scrollable().
func (notebook Notebook) GetScrollable() bool {
	ret0 := C.gtk_notebook_get_scrollable(notebook.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowBorder is a wrapper around gtk_notebook_get_show_border().
func (notebook Notebook) GetShowBorder() bool {
	ret0 := C.gtk_notebook_get_show_border(notebook.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowTabs is a wrapper around gtk_notebook_get_show_tabs().
func (notebook Notebook) GetShowTabs() bool {
	ret0 := C.gtk_notebook_get_show_tabs(notebook.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTabDetachable is a wrapper around gtk_notebook_get_tab_detachable().
func (notebook Notebook) GetTabDetachable(child Widget) bool {
	ret0 := C.gtk_notebook_get_tab_detachable(notebook.native(), child.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTabLabel is a wrapper around gtk_notebook_get_tab_label().
func (notebook Notebook) GetTabLabel(child Widget) Widget {
	ret0 := C.gtk_notebook_get_tab_label(notebook.native(), child.native())
	return wrapWidget(ret0)
}

// GetTabLabelText is a wrapper around gtk_notebook_get_tab_label_text().
func (notebook Notebook) GetTabLabelText(child Widget) string {
	ret0 := C.gtk_notebook_get_tab_label_text(notebook.native(), child.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTabPos is a wrapper around gtk_notebook_get_tab_pos().
func (notebook Notebook) GetTabPos() PositionType {
	ret0 := C.gtk_notebook_get_tab_pos(notebook.native())
	return PositionType(ret0)
}

// GetTabReorderable is a wrapper around gtk_notebook_get_tab_reorderable().
func (notebook Notebook) GetTabReorderable(child Widget) bool {
	ret0 := C.gtk_notebook_get_tab_reorderable(notebook.native(), child.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InsertPage is a wrapper around gtk_notebook_insert_page().
func (notebook Notebook) InsertPage(child Widget, tab_label Widget, position int) int {
	ret0 := C.gtk_notebook_insert_page(notebook.native(), child.native(), tab_label.native(), C.gint(position))
	return int(ret0)
}

// InsertPageMenu is a wrapper around gtk_notebook_insert_page_menu().
func (notebook Notebook) InsertPageMenu(child Widget, tab_label Widget, menu_label Widget, position int) int {
	ret0 := C.gtk_notebook_insert_page_menu(notebook.native(), child.native(), tab_label.native(), menu_label.native(), C.gint(position))
	return int(ret0)
}

// NextPage is a wrapper around gtk_notebook_next_page().
func (notebook Notebook) NextPage() {
	C.gtk_notebook_next_page(notebook.native())
}

// PageNum is a wrapper around gtk_notebook_page_num().
func (notebook Notebook) PageNum(child Widget) int {
	ret0 := C.gtk_notebook_page_num(notebook.native(), child.native())
	return int(ret0)
}

// PopupDisable is a wrapper around gtk_notebook_popup_disable().
func (notebook Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(notebook.native())
}

// PopupEnable is a wrapper around gtk_notebook_popup_enable().
func (notebook Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(notebook.native())
}

// PrependPage is a wrapper around gtk_notebook_prepend_page().
func (notebook Notebook) PrependPage(child Widget, tab_label Widget) int {
	ret0 := C.gtk_notebook_prepend_page(notebook.native(), child.native(), tab_label.native())
	return int(ret0)
}

// PrependPageMenu is a wrapper around gtk_notebook_prepend_page_menu().
func (notebook Notebook) PrependPageMenu(child Widget, tab_label Widget, menu_label Widget) int {
	ret0 := C.gtk_notebook_prepend_page_menu(notebook.native(), child.native(), tab_label.native(), menu_label.native())
	return int(ret0)
}

// PrevPage is a wrapper around gtk_notebook_prev_page().
func (notebook Notebook) PrevPage() {
	C.gtk_notebook_prev_page(notebook.native())
}

// RemovePage is a wrapper around gtk_notebook_remove_page().
func (notebook Notebook) RemovePage(page_num int) {
	C.gtk_notebook_remove_page(notebook.native(), C.gint(page_num))
}

// ReorderChild is a wrapper around gtk_notebook_reorder_child().
func (notebook Notebook) ReorderChild(child Widget, position int) {
	C.gtk_notebook_reorder_child(notebook.native(), child.native(), C.gint(position))
}

// SetActionWidget is a wrapper around gtk_notebook_set_action_widget().
func (notebook Notebook) SetActionWidget(widget Widget, pack_type PackType) {
	C.gtk_notebook_set_action_widget(notebook.native(), widget.native(), C.GtkPackType(pack_type))
}

// SetCurrentPage is a wrapper around gtk_notebook_set_current_page().
func (notebook Notebook) SetCurrentPage(page_num int) {
	C.gtk_notebook_set_current_page(notebook.native(), C.gint(page_num))
}

// SetGroupName is a wrapper around gtk_notebook_set_group_name().
func (notebook Notebook) SetGroupName(group_name string) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	C.gtk_notebook_set_group_name(notebook.native(), group_name0)
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
}

// SetMenuLabel is a wrapper around gtk_notebook_set_menu_label().
func (notebook Notebook) SetMenuLabel(child Widget, menu_label Widget) {
	C.gtk_notebook_set_menu_label(notebook.native(), child.native(), menu_label.native())
}

// SetMenuLabelText is a wrapper around gtk_notebook_set_menu_label_text().
func (notebook Notebook) SetMenuLabelText(child Widget, menu_text string) {
	menu_text0 := (*C.gchar)(C.CString(menu_text))
	C.gtk_notebook_set_menu_label_text(notebook.native(), child.native(), menu_text0)
	C.free(unsafe.Pointer(menu_text0)) /*ch:<stdlib.h>*/
}

// SetScrollable is a wrapper around gtk_notebook_set_scrollable().
func (notebook Notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(notebook.native(), C.gboolean(util.Bool2Int(scrollable)) /*go:.util*/)
}

// SetShowBorder is a wrapper around gtk_notebook_set_show_border().
func (notebook Notebook) SetShowBorder(show_border bool) {
	C.gtk_notebook_set_show_border(notebook.native(), C.gboolean(util.Bool2Int(show_border)) /*go:.util*/)
}

// SetShowTabs is a wrapper around gtk_notebook_set_show_tabs().
func (notebook Notebook) SetShowTabs(show_tabs bool) {
	C.gtk_notebook_set_show_tabs(notebook.native(), C.gboolean(util.Bool2Int(show_tabs)) /*go:.util*/)
}

// SetTabDetachable is a wrapper around gtk_notebook_set_tab_detachable().
func (notebook Notebook) SetTabDetachable(child Widget, detachable bool) {
	C.gtk_notebook_set_tab_detachable(notebook.native(), child.native(), C.gboolean(util.Bool2Int(detachable)) /*go:.util*/)
}

// SetTabLabel is a wrapper around gtk_notebook_set_tab_label().
func (notebook Notebook) SetTabLabel(child Widget, tab_label Widget) {
	C.gtk_notebook_set_tab_label(notebook.native(), child.native(), tab_label.native())
}

// SetTabLabelText is a wrapper around gtk_notebook_set_tab_label_text().
func (notebook Notebook) SetTabLabelText(child Widget, tab_text string) {
	tab_text0 := (*C.gchar)(C.CString(tab_text))
	C.gtk_notebook_set_tab_label_text(notebook.native(), child.native(), tab_text0)
	C.free(unsafe.Pointer(tab_text0)) /*ch:<stdlib.h>*/
}

// SetTabPos is a wrapper around gtk_notebook_set_tab_pos().
func (notebook Notebook) SetTabPos(pos PositionType) {
	C.gtk_notebook_set_tab_pos(notebook.native(), C.GtkPositionType(pos))
}

// SetTabReorderable is a wrapper around gtk_notebook_set_tab_reorderable().
func (notebook Notebook) SetTabReorderable(child Widget, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(notebook.native(), child.native(), C.gboolean(util.Bool2Int(reorderable)) /*go:.util*/)
}

// Object NotebookAccessible
type NotebookAccessible struct {
	ContainerAccessible
}

func (v NotebookAccessible) native() *C.GtkNotebookAccessible {
	return (*C.GtkNotebookAccessible)(v.Ptr)
}
func wrapNotebookAccessible(p *C.GtkNotebookAccessible) (v NotebookAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNotebookAccessible(p unsafe.Pointer) (v NotebookAccessible) {
	v.Ptr = p
	return
}
func (v NotebookAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNotebookAccessible(p unsafe.Pointer) interface{} {
	return WrapNotebookAccessible(p)
}
func (v NotebookAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_notebook_accessible_get_type())
}
func (v NotebookAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNotebookAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v NotebookAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v NotebookAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object NotebookPageAccessible
type NotebookPageAccessible struct {
	atk.Object
}

func (v NotebookPageAccessible) native() *C.GtkNotebookPageAccessible {
	return (*C.GtkNotebookPageAccessible)(v.Ptr)
}
func wrapNotebookPageAccessible(p *C.GtkNotebookPageAccessible) (v NotebookPageAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNotebookPageAccessible(p unsafe.Pointer) (v NotebookPageAccessible) {
	v.Ptr = p
	return
}
func (v NotebookPageAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNotebookPageAccessible(p unsafe.Pointer) interface{} {
	return WrapNotebookPageAccessible(p)
}
func (v NotebookPageAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_notebook_page_accessible_get_type())
}
func (v NotebookPageAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNotebookPageAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v NotebookPageAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// NotebookPageAccessibleNew is a wrapper around gtk_notebook_page_accessible_new().
func NotebookPageAccessibleNew(notebook NotebookAccessible, child Widget) atk.Object {
	ret0 := C.gtk_notebook_page_accessible_new(notebook.native(), child.native())
	return atk.WrapObject(unsafe.Pointer(ret0)) /*gir:Atk*/
}

// Invalidate is a wrapper around gtk_notebook_page_accessible_invalidate().
func (page NotebookPageAccessible) Invalidate() {
	C.gtk_notebook_page_accessible_invalidate(page.native())
}

// Object NumerableIcon
type NumerableIcon struct {
	gio.EmblemedIcon
}

func (v NumerableIcon) native() *C.GtkNumerableIcon {
	return (*C.GtkNumerableIcon)(v.Ptr)
}
func wrapNumerableIcon(p *C.GtkNumerableIcon) (v NumerableIcon) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNumerableIcon(p unsafe.Pointer) (v NumerableIcon) {
	v.Ptr = p
	return
}
func (v NumerableIcon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNumerableIcon(p unsafe.Pointer) interface{} {
	return WrapNumerableIcon(p)
}
func (v NumerableIcon) GetType() gobject.Type {
	return gobject.Type(C.gtk_numerable_icon_get_type())
}
func (v NumerableIcon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNumerableIcon(unsafe.Pointer(ptr)), nil
	}
}
func (v NumerableIcon) Icon() gio.Icon {
	return gio.WrapIcon(v.Ptr) /*gir:Gio*/
}

// Object OffscreenWindow
type OffscreenWindow struct {
	Window
}

func (v OffscreenWindow) native() *C.GtkOffscreenWindow {
	return (*C.GtkOffscreenWindow)(v.Ptr)
}
func wrapOffscreenWindow(p *C.GtkOffscreenWindow) (v OffscreenWindow) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapOffscreenWindow(p unsafe.Pointer) (v OffscreenWindow) {
	v.Ptr = p
	return
}
func (v OffscreenWindow) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOffscreenWindow(p unsafe.Pointer) interface{} {
	return WrapOffscreenWindow(p)
}
func (v OffscreenWindow) GetType() gobject.Type {
	return gobject.Type(C.gtk_offscreen_window_get_type())
}
func (v OffscreenWindow) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapOffscreenWindow(unsafe.Pointer(ptr)), nil
	}
}
func (v OffscreenWindow) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v OffscreenWindow) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// OffscreenWindowNew is a wrapper around gtk_offscreen_window_new().
func OffscreenWindowNew() Widget {
	ret0 := C.gtk_offscreen_window_new()
	return wrapWidget(ret0)
}

// GetPixbuf is a wrapper around gtk_offscreen_window_get_pixbuf().
func (offscreen OffscreenWindow) GetPixbuf() gdkpixbuf.Pixbuf {
	ret0 := C.gtk_offscreen_window_get_pixbuf(offscreen.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetSurface is a wrapper around gtk_offscreen_window_get_surface().
func (offscreen OffscreenWindow) GetSurface() cairo.Surface {
	ret0 := C.gtk_offscreen_window_get_surface(offscreen.native())
	return cairo.WrapSurface(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// Object Overlay
type Overlay struct {
	Bin
}

func (v Overlay) native() *C.GtkOverlay {
	return (*C.GtkOverlay)(v.Ptr)
}
func wrapOverlay(p *C.GtkOverlay) (v Overlay) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapOverlay(p unsafe.Pointer) (v Overlay) {
	v.Ptr = p
	return
}
func (v Overlay) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOverlay(p unsafe.Pointer) interface{} {
	return WrapOverlay(p)
}
func (v Overlay) GetType() gobject.Type {
	return gobject.Type(C.gtk_overlay_get_type())
}
func (v Overlay) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapOverlay(unsafe.Pointer(ptr)), nil
	}
}
func (v Overlay) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Overlay) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// OverlayNew is a wrapper around gtk_overlay_new().
func OverlayNew() Widget {
	ret0 := C.gtk_overlay_new()
	return wrapWidget(ret0)
}

// AddOverlay is a wrapper around gtk_overlay_add_overlay().
func (overlay Overlay) AddOverlay(widget Widget) {
	C.gtk_overlay_add_overlay(overlay.native(), widget.native())
}

// GetOverlayPassThrough is a wrapper around gtk_overlay_get_overlay_pass_through().
func (overlay Overlay) GetOverlayPassThrough(widget Widget) bool {
	ret0 := C.gtk_overlay_get_overlay_pass_through(overlay.native(), widget.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ReorderOverlay is a wrapper around gtk_overlay_reorder_overlay().
func (overlay Overlay) ReorderOverlay(child Widget, position int) {
	C.gtk_overlay_reorder_overlay(overlay.native(), child.native(), C.gint(position))
}

// SetOverlayPassThrough is a wrapper around gtk_overlay_set_overlay_pass_through().
func (overlay Overlay) SetOverlayPassThrough(widget Widget, pass_through bool) {
	C.gtk_overlay_set_overlay_pass_through(overlay.native(), widget.native(), C.gboolean(util.Bool2Int(pass_through)) /*go:.util*/)
}

// Object PadController
type PadController struct {
	EventController
}

func (v PadController) native() *C.GtkPadController {
	return (*C.GtkPadController)(v.Ptr)
}
func wrapPadController(p *C.GtkPadController) (v PadController) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPadController(p unsafe.Pointer) (v PadController) {
	v.Ptr = p
	return
}
func (v PadController) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPadController(p unsafe.Pointer) interface{} {
	return WrapPadController(p)
}
func (v PadController) GetType() gobject.Type {
	return gobject.Type(C.gtk_pad_controller_get_type())
}
func (v PadController) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPadController(unsafe.Pointer(ptr)), nil
	}
}

// SetAction is a wrapper around gtk_pad_controller_set_action().
func (controller PadController) SetAction(type_ PadActionType, index int, mode int, label string, action_name string) {
	label0 := (*C.gchar)(C.CString(label))
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.gtk_pad_controller_set_action(controller.native(), C.GtkPadActionType(type_), C.gint(index), C.gint(mode), label0, action_name0)
	C.free(unsafe.Pointer(label0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// Object PageSetup
type PageSetup struct {
	gobject.Object
}

func (v PageSetup) native() *C.GtkPageSetup {
	return (*C.GtkPageSetup)(v.Ptr)
}
func wrapPageSetup(p *C.GtkPageSetup) (v PageSetup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPageSetup(p unsafe.Pointer) (v PageSetup) {
	v.Ptr = p
	return
}
func (v PageSetup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPageSetup(p unsafe.Pointer) interface{} {
	return WrapPageSetup(p)
}
func (v PageSetup) GetType() gobject.Type {
	return gobject.Type(C.gtk_page_setup_get_type())
}
func (v PageSetup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPageSetup(unsafe.Pointer(ptr)), nil
	}
}

// PageSetupNew is a wrapper around gtk_page_setup_new().
func PageSetupNew() PageSetup {
	ret0 := C.gtk_page_setup_new()
	return wrapPageSetup(ret0)
}

// PageSetupNewFromFile is a wrapper around gtk_page_setup_new_from_file().
func PageSetupNewFromFile(file_name string) (PageSetup, error) {
	file_name0 := (*C.gchar)(C.CString(file_name))
	var err glib.Error
	ret0 := C.gtk_page_setup_new_from_file(file_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PageSetup{}, err.GoValue()
	}
	return wrapPageSetup(ret0), nil
}

// PageSetupNewFromKeyFile is a wrapper around gtk_page_setup_new_from_key_file().
func PageSetupNewFromKeyFile(key_file glib.KeyFile, group_name string) (PageSetup, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	var err glib.Error
	ret0 := C.gtk_page_setup_new_from_key_file((*C.GKeyFile)(key_file.Ptr), group_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PageSetup{}, err.GoValue()
	}
	return wrapPageSetup(ret0), nil
}

// Copy is a wrapper around gtk_page_setup_copy().
func (other PageSetup) Copy() PageSetup {
	ret0 := C.gtk_page_setup_copy(other.native())
	return wrapPageSetup(ret0)
}

// GetBottomMargin is a wrapper around gtk_page_setup_get_bottom_margin().
func (setup PageSetup) GetBottomMargin(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_bottom_margin(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetLeftMargin is a wrapper around gtk_page_setup_get_left_margin().
func (setup PageSetup) GetLeftMargin(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_left_margin(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetOrientation is a wrapper around gtk_page_setup_get_orientation().
func (setup PageSetup) GetOrientation() PageOrientation {
	ret0 := C.gtk_page_setup_get_orientation(setup.native())
	return PageOrientation(ret0)
}

// GetPageHeight is a wrapper around gtk_page_setup_get_page_height().
func (setup PageSetup) GetPageHeight(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_page_height(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetPageWidth is a wrapper around gtk_page_setup_get_page_width().
func (setup PageSetup) GetPageWidth(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_page_width(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetPaperHeight is a wrapper around gtk_page_setup_get_paper_height().
func (setup PageSetup) GetPaperHeight(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_paper_height(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetPaperSize is a wrapper around gtk_page_setup_get_paper_size().
func (setup PageSetup) GetPaperSize() PaperSize {
	ret0 := C.gtk_page_setup_get_paper_size(setup.native())
	return wrapPaperSize(ret0)
}

// GetPaperWidth is a wrapper around gtk_page_setup_get_paper_width().
func (setup PageSetup) GetPaperWidth(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_paper_width(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetRightMargin is a wrapper around gtk_page_setup_get_right_margin().
func (setup PageSetup) GetRightMargin(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_right_margin(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetTopMargin is a wrapper around gtk_page_setup_get_top_margin().
func (setup PageSetup) GetTopMargin(unit Unit) float64 {
	ret0 := C.gtk_page_setup_get_top_margin(setup.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// LoadFile is a wrapper around gtk_page_setup_load_file().
func (setup PageSetup) LoadFile(file_name string) (bool, error) {
	file_name0 := C.CString(file_name)
	var err glib.Error
	ret0 := C.gtk_page_setup_load_file(setup.native(), file_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadKeyFile is a wrapper around gtk_page_setup_load_key_file().
func (setup PageSetup) LoadKeyFile(key_file glib.KeyFile, group_name string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	var err glib.Error
	ret0 := C.gtk_page_setup_load_key_file(setup.native(), (*C.GKeyFile)(key_file.Ptr), group_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetBottomMargin is a wrapper around gtk_page_setup_set_bottom_margin().
func (setup PageSetup) SetBottomMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_bottom_margin(setup.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// SetLeftMargin is a wrapper around gtk_page_setup_set_left_margin().
func (setup PageSetup) SetLeftMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_left_margin(setup.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// SetOrientation is a wrapper around gtk_page_setup_set_orientation().
func (setup PageSetup) SetOrientation(orientation PageOrientation) {
	C.gtk_page_setup_set_orientation(setup.native(), C.GtkPageOrientation(orientation))
}

// SetPaperSize is a wrapper around gtk_page_setup_set_paper_size().
func (setup PageSetup) SetPaperSize(size PaperSize) {
	C.gtk_page_setup_set_paper_size(setup.native(), size.native())
}

// SetPaperSizeAndDefaultMargins is a wrapper around gtk_page_setup_set_paper_size_and_default_margins().
func (setup PageSetup) SetPaperSizeAndDefaultMargins(size PaperSize) {
	C.gtk_page_setup_set_paper_size_and_default_margins(setup.native(), size.native())
}

// SetRightMargin is a wrapper around gtk_page_setup_set_right_margin().
func (setup PageSetup) SetRightMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_right_margin(setup.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// SetTopMargin is a wrapper around gtk_page_setup_set_top_margin().
func (setup PageSetup) SetTopMargin(margin float64, unit Unit) {
	C.gtk_page_setup_set_top_margin(setup.native(), C.gdouble(margin), C.GtkUnit(unit))
}

// ToFile is a wrapper around gtk_page_setup_to_file().
func (setup PageSetup) ToFile(file_name string) (bool, error) {
	file_name0 := C.CString(file_name)
	var err glib.Error
	ret0 := C.gtk_page_setup_to_file(setup.native(), file_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ToGvariant is a wrapper around gtk_page_setup_to_gvariant().
func (setup PageSetup) ToGvariant() glib.Variant {
	ret0 := C.gtk_page_setup_to_gvariant(setup.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// Object PanedAccessible
type PanedAccessible struct {
	ContainerAccessible
}

func (v PanedAccessible) native() *C.GtkPanedAccessible {
	return (*C.GtkPanedAccessible)(v.Ptr)
}
func wrapPanedAccessible(p *C.GtkPanedAccessible) (v PanedAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPanedAccessible(p unsafe.Pointer) (v PanedAccessible) {
	v.Ptr = p
	return
}
func (v PanedAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPanedAccessible(p unsafe.Pointer) interface{} {
	return WrapPanedAccessible(p)
}
func (v PanedAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_paned_accessible_get_type())
}
func (v PanedAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPanedAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v PanedAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v PanedAccessible) Value() atk.Value {
	return atk.WrapValue(v.Ptr) /*gir:Atk*/
}

// Object ScrolledWindow
type ScrolledWindow struct {
	Bin
}

func (v ScrolledWindow) native() *C.GtkScrolledWindow {
	return (*C.GtkScrolledWindow)(v.Ptr)
}
func wrapScrolledWindow(p *C.GtkScrolledWindow) (v ScrolledWindow) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScrolledWindow(p unsafe.Pointer) (v ScrolledWindow) {
	v.Ptr = p
	return
}
func (v ScrolledWindow) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScrolledWindow(p unsafe.Pointer) interface{} {
	return WrapScrolledWindow(p)
}
func (v ScrolledWindow) GetType() gobject.Type {
	return gobject.Type(C.gtk_scrolled_window_get_type())
}
func (v ScrolledWindow) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScrolledWindow(unsafe.Pointer(ptr)), nil
	}
}
func (v ScrolledWindow) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ScrolledWindow) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ScrolledWindowNew is a wrapper around gtk_scrolled_window_new().
func ScrolledWindowNew(hadjustment Adjustment, vadjustment Adjustment) Widget {
	ret0 := C.gtk_scrolled_window_new(hadjustment.native(), vadjustment.native())
	return wrapWidget(ret0)
}

// GetCaptureButtonPress is a wrapper around gtk_scrolled_window_get_capture_button_press().
func (scrolled_window ScrolledWindow) GetCaptureButtonPress() bool {
	ret0 := C.gtk_scrolled_window_get_capture_button_press(scrolled_window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHadjustment is a wrapper around gtk_scrolled_window_get_hadjustment().
func (scrolled_window ScrolledWindow) GetHadjustment() Adjustment {
	ret0 := C.gtk_scrolled_window_get_hadjustment(scrolled_window.native())
	return wrapAdjustment(ret0)
}

// GetHscrollbar is a wrapper around gtk_scrolled_window_get_hscrollbar().
func (scrolled_window ScrolledWindow) GetHscrollbar() Widget {
	ret0 := C.gtk_scrolled_window_get_hscrollbar(scrolled_window.native())
	return wrapWidget(ret0)
}

// GetKineticScrolling is a wrapper around gtk_scrolled_window_get_kinetic_scrolling().
func (scrolled_window ScrolledWindow) GetKineticScrolling() bool {
	ret0 := C.gtk_scrolled_window_get_kinetic_scrolling(scrolled_window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMaxContentHeight is a wrapper around gtk_scrolled_window_get_max_content_height().
func (scrolled_window ScrolledWindow) GetMaxContentHeight() int {
	ret0 := C.gtk_scrolled_window_get_max_content_height(scrolled_window.native())
	return int(ret0)
}

// GetMaxContentWidth is a wrapper around gtk_scrolled_window_get_max_content_width().
func (scrolled_window ScrolledWindow) GetMaxContentWidth() int {
	ret0 := C.gtk_scrolled_window_get_max_content_width(scrolled_window.native())
	return int(ret0)
}

// GetMinContentHeight is a wrapper around gtk_scrolled_window_get_min_content_height().
func (scrolled_window ScrolledWindow) GetMinContentHeight() int {
	ret0 := C.gtk_scrolled_window_get_min_content_height(scrolled_window.native())
	return int(ret0)
}

// GetMinContentWidth is a wrapper around gtk_scrolled_window_get_min_content_width().
func (scrolled_window ScrolledWindow) GetMinContentWidth() int {
	ret0 := C.gtk_scrolled_window_get_min_content_width(scrolled_window.native())
	return int(ret0)
}

// GetOverlayScrolling is a wrapper around gtk_scrolled_window_get_overlay_scrolling().
func (scrolled_window ScrolledWindow) GetOverlayScrolling() bool {
	ret0 := C.gtk_scrolled_window_get_overlay_scrolling(scrolled_window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPlacement is a wrapper around gtk_scrolled_window_get_placement().
func (scrolled_window ScrolledWindow) GetPlacement() CornerType {
	ret0 := C.gtk_scrolled_window_get_placement(scrolled_window.native())
	return CornerType(ret0)
}

// GetPolicy is a wrapper around gtk_scrolled_window_get_policy().
func (scrolled_window ScrolledWindow) GetPolicy() (PolicyType, PolicyType) {
	var hscrollbar_policy0 C.GtkPolicyType
	var vscrollbar_policy0 C.GtkPolicyType
	C.gtk_scrolled_window_get_policy(scrolled_window.native(), &hscrollbar_policy0, &vscrollbar_policy0)
	return PolicyType(hscrollbar_policy0), PolicyType(vscrollbar_policy0)
}

// GetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_get_propagate_natural_height().
func (scrolled_window ScrolledWindow) GetPropagateNaturalHeight() bool {
	ret0 := C.gtk_scrolled_window_get_propagate_natural_height(scrolled_window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_get_propagate_natural_width().
func (scrolled_window ScrolledWindow) GetPropagateNaturalWidth() bool {
	ret0 := C.gtk_scrolled_window_get_propagate_natural_width(scrolled_window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShadowType is a wrapper around gtk_scrolled_window_get_shadow_type().
func (scrolled_window ScrolledWindow) GetShadowType() ShadowType {
	ret0 := C.gtk_scrolled_window_get_shadow_type(scrolled_window.native())
	return ShadowType(ret0)
}

// GetVadjustment is a wrapper around gtk_scrolled_window_get_vadjustment().
func (scrolled_window ScrolledWindow) GetVadjustment() Adjustment {
	ret0 := C.gtk_scrolled_window_get_vadjustment(scrolled_window.native())
	return wrapAdjustment(ret0)
}

// GetVscrollbar is a wrapper around gtk_scrolled_window_get_vscrollbar().
func (scrolled_window ScrolledWindow) GetVscrollbar() Widget {
	ret0 := C.gtk_scrolled_window_get_vscrollbar(scrolled_window.native())
	return wrapWidget(ret0)
}

// SetCaptureButtonPress is a wrapper around gtk_scrolled_window_set_capture_button_press().
func (scrolled_window ScrolledWindow) SetCaptureButtonPress(capture_button_press bool) {
	C.gtk_scrolled_window_set_capture_button_press(scrolled_window.native(), C.gboolean(util.Bool2Int(capture_button_press)) /*go:.util*/)
}

// SetHadjustment is a wrapper around gtk_scrolled_window_set_hadjustment().
func (scrolled_window ScrolledWindow) SetHadjustment(hadjustment Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(scrolled_window.native(), hadjustment.native())
}

// SetKineticScrolling is a wrapper around gtk_scrolled_window_set_kinetic_scrolling().
func (scrolled_window ScrolledWindow) SetKineticScrolling(kinetic_scrolling bool) {
	C.gtk_scrolled_window_set_kinetic_scrolling(scrolled_window.native(), C.gboolean(util.Bool2Int(kinetic_scrolling)) /*go:.util*/)
}

// SetMaxContentHeight is a wrapper around gtk_scrolled_window_set_max_content_height().
func (scrolled_window ScrolledWindow) SetMaxContentHeight(height int) {
	C.gtk_scrolled_window_set_max_content_height(scrolled_window.native(), C.gint(height))
}

// SetMaxContentWidth is a wrapper around gtk_scrolled_window_set_max_content_width().
func (scrolled_window ScrolledWindow) SetMaxContentWidth(width int) {
	C.gtk_scrolled_window_set_max_content_width(scrolled_window.native(), C.gint(width))
}

// SetMinContentHeight is a wrapper around gtk_scrolled_window_set_min_content_height().
func (scrolled_window ScrolledWindow) SetMinContentHeight(height int) {
	C.gtk_scrolled_window_set_min_content_height(scrolled_window.native(), C.gint(height))
}

// SetMinContentWidth is a wrapper around gtk_scrolled_window_set_min_content_width().
func (scrolled_window ScrolledWindow) SetMinContentWidth(width int) {
	C.gtk_scrolled_window_set_min_content_width(scrolled_window.native(), C.gint(width))
}

// SetOverlayScrolling is a wrapper around gtk_scrolled_window_set_overlay_scrolling().
func (scrolled_window ScrolledWindow) SetOverlayScrolling(overlay_scrolling bool) {
	C.gtk_scrolled_window_set_overlay_scrolling(scrolled_window.native(), C.gboolean(util.Bool2Int(overlay_scrolling)) /*go:.util*/)
}

// SetPlacement is a wrapper around gtk_scrolled_window_set_placement().
func (scrolled_window ScrolledWindow) SetPlacement(window_placement CornerType) {
	C.gtk_scrolled_window_set_placement(scrolled_window.native(), C.GtkCornerType(window_placement))
}

// SetPolicy is a wrapper around gtk_scrolled_window_set_policy().
func (scrolled_window ScrolledWindow) SetPolicy(hscrollbar_policy PolicyType, vscrollbar_policy PolicyType) {
	C.gtk_scrolled_window_set_policy(scrolled_window.native(), C.GtkPolicyType(hscrollbar_policy), C.GtkPolicyType(vscrollbar_policy))
}

// SetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_set_propagate_natural_height().
func (scrolled_window ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_height(scrolled_window.native(), C.gboolean(util.Bool2Int(propagate)) /*go:.util*/)
}

// SetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_set_propagate_natural_width().
func (scrolled_window ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_width(scrolled_window.native(), C.gboolean(util.Bool2Int(propagate)) /*go:.util*/)
}

// SetShadowType is a wrapper around gtk_scrolled_window_set_shadow_type().
func (scrolled_window ScrolledWindow) SetShadowType(type_ ShadowType) {
	C.gtk_scrolled_window_set_shadow_type(scrolled_window.native(), C.GtkShadowType(type_))
}

// SetVadjustment is a wrapper around gtk_scrolled_window_set_vadjustment().
func (scrolled_window ScrolledWindow) SetVadjustment(vadjustment Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(scrolled_window.native(), vadjustment.native())
}

// UnsetPlacement is a wrapper around gtk_scrolled_window_unset_placement().
func (scrolled_window ScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(scrolled_window.native())
}

// Object PlacesSidebar
type PlacesSidebar struct {
	ScrolledWindow
}

func (v PlacesSidebar) native() *C.GtkPlacesSidebar {
	return (*C.GtkPlacesSidebar)(v.Ptr)
}
func wrapPlacesSidebar(p *C.GtkPlacesSidebar) (v PlacesSidebar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPlacesSidebar(p unsafe.Pointer) (v PlacesSidebar) {
	v.Ptr = p
	return
}
func (v PlacesSidebar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPlacesSidebar(p unsafe.Pointer) interface{} {
	return WrapPlacesSidebar(p)
}
func (v PlacesSidebar) GetType() gobject.Type {
	return gobject.Type(C.gtk_places_sidebar_get_type())
}
func (v PlacesSidebar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPlacesSidebar(unsafe.Pointer(ptr)), nil
	}
}
func (v PlacesSidebar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v PlacesSidebar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// PlacesSidebarNew is a wrapper around gtk_places_sidebar_new().
func PlacesSidebarNew() Widget {
	ret0 := C.gtk_places_sidebar_new()
	return wrapWidget(ret0)
}

// AddShortcut is a wrapper around gtk_places_sidebar_add_shortcut().
func (sidebar PlacesSidebar) AddShortcut(location gio.File) {
	C.gtk_places_sidebar_add_shortcut(sidebar.native(), (*C.GFile)(location.Ptr))
}

// GetLocalOnly is a wrapper around gtk_places_sidebar_get_local_only().
func (sidebar PlacesSidebar) GetLocalOnly() bool {
	ret0 := C.gtk_places_sidebar_get_local_only(sidebar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLocation is a wrapper around gtk_places_sidebar_get_location().
func (sidebar PlacesSidebar) GetLocation() gio.File {
	ret0 := C.gtk_places_sidebar_get_location(sidebar.native())
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetNthBookmark is a wrapper around gtk_places_sidebar_get_nth_bookmark().
func (sidebar PlacesSidebar) GetNthBookmark(n int) gio.File {
	ret0 := C.gtk_places_sidebar_get_nth_bookmark(sidebar.native(), C.gint(n))
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetOpenFlags is a wrapper around gtk_places_sidebar_get_open_flags().
func (sidebar PlacesSidebar) GetOpenFlags() PlacesOpenFlags {
	ret0 := C.gtk_places_sidebar_get_open_flags(sidebar.native())
	return PlacesOpenFlags(ret0)
}

// GetShowDesktop is a wrapper around gtk_places_sidebar_get_show_desktop().
func (sidebar PlacesSidebar) GetShowDesktop() bool {
	ret0 := C.gtk_places_sidebar_get_show_desktop(sidebar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowEnterLocation is a wrapper around gtk_places_sidebar_get_show_enter_location().
func (sidebar PlacesSidebar) GetShowEnterLocation() bool {
	ret0 := C.gtk_places_sidebar_get_show_enter_location(sidebar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowOtherLocations is a wrapper around gtk_places_sidebar_get_show_other_locations().
func (sidebar PlacesSidebar) GetShowOtherLocations() bool {
	ret0 := C.gtk_places_sidebar_get_show_other_locations(sidebar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowRecent is a wrapper around gtk_places_sidebar_get_show_recent().
func (sidebar PlacesSidebar) GetShowRecent() bool {
	ret0 := C.gtk_places_sidebar_get_show_recent(sidebar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowTrash is a wrapper around gtk_places_sidebar_get_show_trash().
func (sidebar PlacesSidebar) GetShowTrash() bool {
	ret0 := C.gtk_places_sidebar_get_show_trash(sidebar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RemoveShortcut is a wrapper around gtk_places_sidebar_remove_shortcut().
func (sidebar PlacesSidebar) RemoveShortcut(location gio.File) {
	C.gtk_places_sidebar_remove_shortcut(sidebar.native(), (*C.GFile)(location.Ptr))
}

// SetLocalOnly is a wrapper around gtk_places_sidebar_set_local_only().
func (sidebar PlacesSidebar) SetLocalOnly(local_only bool) {
	C.gtk_places_sidebar_set_local_only(sidebar.native(), C.gboolean(util.Bool2Int(local_only)) /*go:.util*/)
}

// SetLocation is a wrapper around gtk_places_sidebar_set_location().
func (sidebar PlacesSidebar) SetLocation(location gio.File) {
	C.gtk_places_sidebar_set_location(sidebar.native(), (*C.GFile)(location.Ptr))
}

// SetOpenFlags is a wrapper around gtk_places_sidebar_set_open_flags().
func (sidebar PlacesSidebar) SetOpenFlags(flags PlacesOpenFlags) {
	C.gtk_places_sidebar_set_open_flags(sidebar.native(), C.GtkPlacesOpenFlags(flags))
}

// SetShowDesktop is a wrapper around gtk_places_sidebar_set_show_desktop().
func (sidebar PlacesSidebar) SetShowDesktop(show_desktop bool) {
	C.gtk_places_sidebar_set_show_desktop(sidebar.native(), C.gboolean(util.Bool2Int(show_desktop)) /*go:.util*/)
}

// SetShowEnterLocation is a wrapper around gtk_places_sidebar_set_show_enter_location().
func (sidebar PlacesSidebar) SetShowEnterLocation(show_enter_location bool) {
	C.gtk_places_sidebar_set_show_enter_location(sidebar.native(), C.gboolean(util.Bool2Int(show_enter_location)) /*go:.util*/)
}

// SetShowOtherLocations is a wrapper around gtk_places_sidebar_set_show_other_locations().
func (sidebar PlacesSidebar) SetShowOtherLocations(show_other_locations bool) {
	C.gtk_places_sidebar_set_show_other_locations(sidebar.native(), C.gboolean(util.Bool2Int(show_other_locations)) /*go:.util*/)
}

// SetShowRecent is a wrapper around gtk_places_sidebar_set_show_recent().
func (sidebar PlacesSidebar) SetShowRecent(show_recent bool) {
	C.gtk_places_sidebar_set_show_recent(sidebar.native(), C.gboolean(util.Bool2Int(show_recent)) /*go:.util*/)
}

// SetShowTrash is a wrapper around gtk_places_sidebar_set_show_trash().
func (sidebar PlacesSidebar) SetShowTrash(show_trash bool) {
	C.gtk_places_sidebar_set_show_trash(sidebar.native(), C.gboolean(util.Bool2Int(show_trash)) /*go:.util*/)
}

// Object Plug
type Plug struct {
	Window
}

func (v Plug) native() *C.GtkPlug {
	return (*C.GtkPlug)(v.Ptr)
}
func wrapPlug(p *C.GtkPlug) (v Plug) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPlug(p unsafe.Pointer) (v Plug) {
	v.Ptr = p
	return
}
func (v Plug) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPlug(p unsafe.Pointer) interface{} {
	return WrapPlug(p)
}
func (v Plug) GetType() gobject.Type {
	return gobject.Type(C.gtk_plug_get_type())
}
func (v Plug) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPlug(unsafe.Pointer(ptr)), nil
	}
}
func (v Plug) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Plug) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object PopoverAccessible
type PopoverAccessible struct {
	ContainerAccessible
}

func (v PopoverAccessible) native() *C.GtkPopoverAccessible {
	return (*C.GtkPopoverAccessible)(v.Ptr)
}
func wrapPopoverAccessible(p *C.GtkPopoverAccessible) (v PopoverAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPopoverAccessible(p unsafe.Pointer) (v PopoverAccessible) {
	v.Ptr = p
	return
}
func (v PopoverAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPopoverAccessible(p unsafe.Pointer) interface{} {
	return WrapPopoverAccessible(p)
}
func (v PopoverAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_popover_accessible_get_type())
}
func (v PopoverAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPopoverAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v PopoverAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object PopoverMenu
type PopoverMenu struct {
	Popover
}

func (v PopoverMenu) native() *C.GtkPopoverMenu {
	return (*C.GtkPopoverMenu)(v.Ptr)
}
func wrapPopoverMenu(p *C.GtkPopoverMenu) (v PopoverMenu) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPopoverMenu(p unsafe.Pointer) (v PopoverMenu) {
	v.Ptr = p
	return
}
func (v PopoverMenu) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPopoverMenu(p unsafe.Pointer) interface{} {
	return WrapPopoverMenu(p)
}
func (v PopoverMenu) GetType() gobject.Type {
	return gobject.Type(C.gtk_popover_menu_get_type())
}
func (v PopoverMenu) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPopoverMenu(unsafe.Pointer(ptr)), nil
	}
}
func (v PopoverMenu) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v PopoverMenu) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object PrintContext
type PrintContext struct {
	gobject.Object
}

func (v PrintContext) native() *C.GtkPrintContext {
	return (*C.GtkPrintContext)(v.Ptr)
}
func wrapPrintContext(p *C.GtkPrintContext) (v PrintContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPrintContext(p unsafe.Pointer) (v PrintContext) {
	v.Ptr = p
	return
}
func (v PrintContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPrintContext(p unsafe.Pointer) interface{} {
	return WrapPrintContext(p)
}
func (v PrintContext) GetType() gobject.Type {
	return gobject.Type(C.gtk_print_context_get_type())
}
func (v PrintContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPrintContext(unsafe.Pointer(ptr)), nil
	}
}

// CreatePangoContext is a wrapper around gtk_print_context_create_pango_context().
func (context PrintContext) CreatePangoContext() pango.Context {
	ret0 := C.gtk_print_context_create_pango_context(context.native())
	return pango.WrapContext(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// CreatePangoLayout is a wrapper around gtk_print_context_create_pango_layout().
func (context PrintContext) CreatePangoLayout() pango.Layout {
	ret0 := C.gtk_print_context_create_pango_layout(context.native())
	return pango.WrapLayout(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetCairoContext is a wrapper around gtk_print_context_get_cairo_context().
func (context PrintContext) GetCairoContext() cairo.Context {
	ret0 := C.gtk_print_context_get_cairo_context(context.native())
	return cairo.WrapContext(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetDpiX is a wrapper around gtk_print_context_get_dpi_x().
func (context PrintContext) GetDpiX() float64 {
	ret0 := C.gtk_print_context_get_dpi_x(context.native())
	return float64(ret0)
}

// GetDpiY is a wrapper around gtk_print_context_get_dpi_y().
func (context PrintContext) GetDpiY() float64 {
	ret0 := C.gtk_print_context_get_dpi_y(context.native())
	return float64(ret0)
}

// GetHardMargins is a wrapper around gtk_print_context_get_hard_margins().
func (context PrintContext) GetHardMargins() (bool, float64, float64, float64, float64) {
	var top0 C.gdouble
	var bottom0 C.gdouble
	var left0 C.gdouble
	var right0 C.gdouble
	ret0 := C.gtk_print_context_get_hard_margins(context.native(), &top0, &bottom0, &left0, &right0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(top0), float64(bottom0), float64(left0), float64(right0)
}

// GetHeight is a wrapper around gtk_print_context_get_height().
func (context PrintContext) GetHeight() float64 {
	ret0 := C.gtk_print_context_get_height(context.native())
	return float64(ret0)
}

// GetPageSetup is a wrapper around gtk_print_context_get_page_setup().
func (context PrintContext) GetPageSetup() PageSetup {
	ret0 := C.gtk_print_context_get_page_setup(context.native())
	return wrapPageSetup(ret0)
}

// GetPangoFontmap is a wrapper around gtk_print_context_get_pango_fontmap().
func (context PrintContext) GetPangoFontmap() pango.FontMap {
	ret0 := C.gtk_print_context_get_pango_fontmap(context.native())
	return pango.WrapFontMap(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetWidth is a wrapper around gtk_print_context_get_width().
func (context PrintContext) GetWidth() float64 {
	ret0 := C.gtk_print_context_get_width(context.native())
	return float64(ret0)
}

// Object PrintOperation
type PrintOperation struct {
	gobject.Object
}

func (v PrintOperation) native() *C.GtkPrintOperation {
	return (*C.GtkPrintOperation)(v.Ptr)
}
func wrapPrintOperation(p *C.GtkPrintOperation) (v PrintOperation) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPrintOperation(p unsafe.Pointer) (v PrintOperation) {
	v.Ptr = p
	return
}
func (v PrintOperation) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPrintOperation(p unsafe.Pointer) interface{} {
	return WrapPrintOperation(p)
}
func (v PrintOperation) GetType() gobject.Type {
	return gobject.Type(C.gtk_print_operation_get_type())
}
func (v PrintOperation) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPrintOperation(unsafe.Pointer(ptr)), nil
	}
}
func (v PrintOperation) PrintOperationPreview() PrintOperationPreview {
	return WrapPrintOperationPreview(v.Ptr)
}

// PrintOperationNew is a wrapper around gtk_print_operation_new().
func PrintOperationNew() PrintOperation {
	ret0 := C.gtk_print_operation_new()
	return wrapPrintOperation(ret0)
}

// Cancel is a wrapper around gtk_print_operation_cancel().
func (op PrintOperation) Cancel() {
	C.gtk_print_operation_cancel(op.native())
}

// DrawPageFinish is a wrapper around gtk_print_operation_draw_page_finish().
func (op PrintOperation) DrawPageFinish() {
	C.gtk_print_operation_draw_page_finish(op.native())
}

// GetDefaultPageSetup is a wrapper around gtk_print_operation_get_default_page_setup().
func (op PrintOperation) GetDefaultPageSetup() PageSetup {
	ret0 := C.gtk_print_operation_get_default_page_setup(op.native())
	return wrapPageSetup(ret0)
}

// GetEmbedPageSetup is a wrapper around gtk_print_operation_get_embed_page_setup().
func (op PrintOperation) GetEmbedPageSetup() bool {
	ret0 := C.gtk_print_operation_get_embed_page_setup(op.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHasSelection is a wrapper around gtk_print_operation_get_has_selection().
func (op PrintOperation) GetHasSelection() bool {
	ret0 := C.gtk_print_operation_get_has_selection(op.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetNPagesToPrint is a wrapper around gtk_print_operation_get_n_pages_to_print().
func (op PrintOperation) GetNPagesToPrint() int {
	ret0 := C.gtk_print_operation_get_n_pages_to_print(op.native())
	return int(ret0)
}

// GetPrintSettings is a wrapper around gtk_print_operation_get_print_settings().
func (op PrintOperation) GetPrintSettings() PrintSettings {
	ret0 := C.gtk_print_operation_get_print_settings(op.native())
	return wrapPrintSettings(ret0)
}

// GetStatus is a wrapper around gtk_print_operation_get_status().
func (op PrintOperation) GetStatus() PrintStatus {
	ret0 := C.gtk_print_operation_get_status(op.native())
	return PrintStatus(ret0)
}

// GetStatusString is a wrapper around gtk_print_operation_get_status_string().
func (op PrintOperation) GetStatusString() string {
	ret0 := C.gtk_print_operation_get_status_string(op.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSupportSelection is a wrapper around gtk_print_operation_get_support_selection().
func (op PrintOperation) GetSupportSelection() bool {
	ret0 := C.gtk_print_operation_get_support_selection(op.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsFinished is a wrapper around gtk_print_operation_is_finished().
func (op PrintOperation) IsFinished() bool {
	ret0 := C.gtk_print_operation_is_finished(op.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetAllowAsync is a wrapper around gtk_print_operation_set_allow_async().
func (op PrintOperation) SetAllowAsync(allow_async bool) {
	C.gtk_print_operation_set_allow_async(op.native(), C.gboolean(util.Bool2Int(allow_async)) /*go:.util*/)
}

// SetCurrentPage is a wrapper around gtk_print_operation_set_current_page().
func (op PrintOperation) SetCurrentPage(current_page int) {
	C.gtk_print_operation_set_current_page(op.native(), C.gint(current_page))
}

// SetCustomTabLabel is a wrapper around gtk_print_operation_set_custom_tab_label().
func (op PrintOperation) SetCustomTabLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_print_operation_set_custom_tab_label(op.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetDefaultPageSetup is a wrapper around gtk_print_operation_set_default_page_setup().
func (op PrintOperation) SetDefaultPageSetup(default_page_setup PageSetup) {
	C.gtk_print_operation_set_default_page_setup(op.native(), default_page_setup.native())
}

// SetDeferDrawing is a wrapper around gtk_print_operation_set_defer_drawing().
func (op PrintOperation) SetDeferDrawing() {
	C.gtk_print_operation_set_defer_drawing(op.native())
}

// SetEmbedPageSetup is a wrapper around gtk_print_operation_set_embed_page_setup().
func (op PrintOperation) SetEmbedPageSetup(embed bool) {
	C.gtk_print_operation_set_embed_page_setup(op.native(), C.gboolean(util.Bool2Int(embed)) /*go:.util*/)
}

// SetExportFilename is a wrapper around gtk_print_operation_set_export_filename().
func (op PrintOperation) SetExportFilename(filename string) {
	filename0 := (*C.gchar)(C.CString(filename))
	C.gtk_print_operation_set_export_filename(op.native(), filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
}

// SetHasSelection is a wrapper around gtk_print_operation_set_has_selection().
func (op PrintOperation) SetHasSelection(has_selection bool) {
	C.gtk_print_operation_set_has_selection(op.native(), C.gboolean(util.Bool2Int(has_selection)) /*go:.util*/)
}

// SetJobName is a wrapper around gtk_print_operation_set_job_name().
func (op PrintOperation) SetJobName(job_name string) {
	job_name0 := (*C.gchar)(C.CString(job_name))
	C.gtk_print_operation_set_job_name(op.native(), job_name0)
	C.free(unsafe.Pointer(job_name0)) /*ch:<stdlib.h>*/
}

// SetNPages is a wrapper around gtk_print_operation_set_n_pages().
func (op PrintOperation) SetNPages(n_pages int) {
	C.gtk_print_operation_set_n_pages(op.native(), C.gint(n_pages))
}

// SetPrintSettings is a wrapper around gtk_print_operation_set_print_settings().
func (op PrintOperation) SetPrintSettings(print_settings PrintSettings) {
	C.gtk_print_operation_set_print_settings(op.native(), print_settings.native())
}

// SetShowProgress is a wrapper around gtk_print_operation_set_show_progress().
func (op PrintOperation) SetShowProgress(show_progress bool) {
	C.gtk_print_operation_set_show_progress(op.native(), C.gboolean(util.Bool2Int(show_progress)) /*go:.util*/)
}

// SetSupportSelection is a wrapper around gtk_print_operation_set_support_selection().
func (op PrintOperation) SetSupportSelection(support_selection bool) {
	C.gtk_print_operation_set_support_selection(op.native(), C.gboolean(util.Bool2Int(support_selection)) /*go:.util*/)
}

// SetTrackPrintStatus is a wrapper around gtk_print_operation_set_track_print_status().
func (op PrintOperation) SetTrackPrintStatus(track_status bool) {
	C.gtk_print_operation_set_track_print_status(op.native(), C.gboolean(util.Bool2Int(track_status)) /*go:.util*/)
}

// SetUnit is a wrapper around gtk_print_operation_set_unit().
func (op PrintOperation) SetUnit(unit Unit) {
	C.gtk_print_operation_set_unit(op.native(), C.GtkUnit(unit))
}

// SetUseFullPage is a wrapper around gtk_print_operation_set_use_full_page().
func (op PrintOperation) SetUseFullPage(full_page bool) {
	C.gtk_print_operation_set_use_full_page(op.native(), C.gboolean(util.Bool2Int(full_page)) /*go:.util*/)
}

// Object PrintSettings
type PrintSettings struct {
	gobject.Object
}

func (v PrintSettings) native() *C.GtkPrintSettings {
	return (*C.GtkPrintSettings)(v.Ptr)
}
func wrapPrintSettings(p *C.GtkPrintSettings) (v PrintSettings) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPrintSettings(p unsafe.Pointer) (v PrintSettings) {
	v.Ptr = p
	return
}
func (v PrintSettings) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPrintSettings(p unsafe.Pointer) interface{} {
	return WrapPrintSettings(p)
}
func (v PrintSettings) GetType() gobject.Type {
	return gobject.Type(C.gtk_print_settings_get_type())
}
func (v PrintSettings) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPrintSettings(unsafe.Pointer(ptr)), nil
	}
}

// PrintSettingsNew is a wrapper around gtk_print_settings_new().
func PrintSettingsNew() PrintSettings {
	ret0 := C.gtk_print_settings_new()
	return wrapPrintSettings(ret0)
}

// PrintSettingsNewFromFile is a wrapper around gtk_print_settings_new_from_file().
func PrintSettingsNewFromFile(file_name string) (PrintSettings, error) {
	file_name0 := (*C.gchar)(C.CString(file_name))
	var err glib.Error
	ret0 := C.gtk_print_settings_new_from_file(file_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PrintSettings{}, err.GoValue()
	}
	return wrapPrintSettings(ret0), nil
}

// PrintSettingsNewFromKeyFile is a wrapper around gtk_print_settings_new_from_key_file().
func PrintSettingsNewFromKeyFile(key_file glib.KeyFile, group_name string) (PrintSettings, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	var err glib.Error
	ret0 := C.gtk_print_settings_new_from_key_file((*C.GKeyFile)(key_file.Ptr), group_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return PrintSettings{}, err.GoValue()
	}
	return wrapPrintSettings(ret0), nil
}

// Copy is a wrapper around gtk_print_settings_copy().
func (other PrintSettings) Copy() PrintSettings {
	ret0 := C.gtk_print_settings_copy(other.native())
	return wrapPrintSettings(ret0)
}

// Get is a wrapper around gtk_print_settings_get().
func (settings PrintSettings) Get(key string) string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_get(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetBool is a wrapper around gtk_print_settings_get_bool().
func (settings PrintSettings) GetBool(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_get_bool(settings.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCollate is a wrapper around gtk_print_settings_get_collate().
func (settings PrintSettings) GetCollate() bool {
	ret0 := C.gtk_print_settings_get_collate(settings.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDefaultSource is a wrapper around gtk_print_settings_get_default_source().
func (settings PrintSettings) GetDefaultSource() string {
	ret0 := C.gtk_print_settings_get_default_source(settings.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDither is a wrapper around gtk_print_settings_get_dither().
func (settings PrintSettings) GetDither() string {
	ret0 := C.gtk_print_settings_get_dither(settings.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDouble is a wrapper around gtk_print_settings_get_double().
func (settings PrintSettings) GetDouble(key string) float64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_get_double(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return float64(ret0)
}

// GetDoubleWithDefault is a wrapper around gtk_print_settings_get_double_with_default().
func (settings PrintSettings) GetDoubleWithDefault(key string, def float64) float64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_get_double_with_default(settings.native(), key0, C.gdouble(def))
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return float64(ret0)
}

// GetDuplex is a wrapper around gtk_print_settings_get_duplex().
func (settings PrintSettings) GetDuplex() PrintDuplex {
	ret0 := C.gtk_print_settings_get_duplex(settings.native())
	return PrintDuplex(ret0)
}

// GetFinishings is a wrapper around gtk_print_settings_get_finishings().
func (settings PrintSettings) GetFinishings() string {
	ret0 := C.gtk_print_settings_get_finishings(settings.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetInt is a wrapper around gtk_print_settings_get_int().
func (settings PrintSettings) GetInt(key string) int {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_get_int(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetIntWithDefault is a wrapper around gtk_print_settings_get_int_with_default().
func (settings PrintSettings) GetIntWithDefault(key string, def int) int {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_get_int_with_default(settings.native(), key0, C.gint(def))
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetLength is a wrapper around gtk_print_settings_get_length().
func (settings PrintSettings) GetLength(key string, unit Unit) float64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_get_length(settings.native(), key0, C.GtkUnit(unit))
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return float64(ret0)
}

// GetMediaType is a wrapper around gtk_print_settings_get_media_type().
func (settings PrintSettings) GetMediaType() string {
	ret0 := C.gtk_print_settings_get_media_type(settings.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNCopies is a wrapper around gtk_print_settings_get_n_copies().
func (settings PrintSettings) GetNCopies() int {
	ret0 := C.gtk_print_settings_get_n_copies(settings.native())
	return int(ret0)
}

// GetNumberUp is a wrapper around gtk_print_settings_get_number_up().
func (settings PrintSettings) GetNumberUp() int {
	ret0 := C.gtk_print_settings_get_number_up(settings.native())
	return int(ret0)
}

// GetNumberUpLayout is a wrapper around gtk_print_settings_get_number_up_layout().
func (settings PrintSettings) GetNumberUpLayout() NumberUpLayout {
	ret0 := C.gtk_print_settings_get_number_up_layout(settings.native())
	return NumberUpLayout(ret0)
}

// GetOrientation is a wrapper around gtk_print_settings_get_orientation().
func (settings PrintSettings) GetOrientation() PageOrientation {
	ret0 := C.gtk_print_settings_get_orientation(settings.native())
	return PageOrientation(ret0)
}

// GetOutputBin is a wrapper around gtk_print_settings_get_output_bin().
func (settings PrintSettings) GetOutputBin() string {
	ret0 := C.gtk_print_settings_get_output_bin(settings.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPageSet is a wrapper around gtk_print_settings_get_page_set().
func (settings PrintSettings) GetPageSet() PageSet {
	ret0 := C.gtk_print_settings_get_page_set(settings.native())
	return PageSet(ret0)
}

// GetPaperHeight is a wrapper around gtk_print_settings_get_paper_height().
func (settings PrintSettings) GetPaperHeight(unit Unit) float64 {
	ret0 := C.gtk_print_settings_get_paper_height(settings.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetPaperSize is a wrapper around gtk_print_settings_get_paper_size().
func (settings PrintSettings) GetPaperSize() PaperSize {
	ret0 := C.gtk_print_settings_get_paper_size(settings.native())
	return wrapPaperSize(ret0)
}

// GetPaperWidth is a wrapper around gtk_print_settings_get_paper_width().
func (settings PrintSettings) GetPaperWidth(unit Unit) float64 {
	ret0 := C.gtk_print_settings_get_paper_width(settings.native(), C.GtkUnit(unit))
	return float64(ret0)
}

// GetPrintPages is a wrapper around gtk_print_settings_get_print_pages().
func (settings PrintSettings) GetPrintPages() PrintPages {
	ret0 := C.gtk_print_settings_get_print_pages(settings.native())
	return PrintPages(ret0)
}

// GetPrinter is a wrapper around gtk_print_settings_get_printer().
func (settings PrintSettings) GetPrinter() string {
	ret0 := C.gtk_print_settings_get_printer(settings.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPrinterLpi is a wrapper around gtk_print_settings_get_printer_lpi().
func (settings PrintSettings) GetPrinterLpi() float64 {
	ret0 := C.gtk_print_settings_get_printer_lpi(settings.native())
	return float64(ret0)
}

// GetQuality is a wrapper around gtk_print_settings_get_quality().
func (settings PrintSettings) GetQuality() PrintQuality {
	ret0 := C.gtk_print_settings_get_quality(settings.native())
	return PrintQuality(ret0)
}

// GetResolution is a wrapper around gtk_print_settings_get_resolution().
func (settings PrintSettings) GetResolution() int {
	ret0 := C.gtk_print_settings_get_resolution(settings.native())
	return int(ret0)
}

// GetResolutionX is a wrapper around gtk_print_settings_get_resolution_x().
func (settings PrintSettings) GetResolutionX() int {
	ret0 := C.gtk_print_settings_get_resolution_x(settings.native())
	return int(ret0)
}

// GetResolutionY is a wrapper around gtk_print_settings_get_resolution_y().
func (settings PrintSettings) GetResolutionY() int {
	ret0 := C.gtk_print_settings_get_resolution_y(settings.native())
	return int(ret0)
}

// GetReverse is a wrapper around gtk_print_settings_get_reverse().
func (settings PrintSettings) GetReverse() bool {
	ret0 := C.gtk_print_settings_get_reverse(settings.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetScale is a wrapper around gtk_print_settings_get_scale().
func (settings PrintSettings) GetScale() float64 {
	ret0 := C.gtk_print_settings_get_scale(settings.native())
	return float64(ret0)
}

// GetUseColor is a wrapper around gtk_print_settings_get_use_color().
func (settings PrintSettings) GetUseColor() bool {
	ret0 := C.gtk_print_settings_get_use_color(settings.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasKey is a wrapper around gtk_print_settings_has_key().
func (settings PrintSettings) HasKey(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.gtk_print_settings_has_key(settings.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LoadFile is a wrapper around gtk_print_settings_load_file().
func (settings PrintSettings) LoadFile(file_name string) (bool, error) {
	file_name0 := (*C.gchar)(C.CString(file_name))
	var err glib.Error
	ret0 := C.gtk_print_settings_load_file(settings.native(), file_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadKeyFile is a wrapper around gtk_print_settings_load_key_file().
func (settings PrintSettings) LoadKeyFile(key_file glib.KeyFile, group_name string) (bool, error) {
	group_name0 := (*C.gchar)(C.CString(group_name))
	var err glib.Error
	ret0 := C.gtk_print_settings_load_key_file(settings.native(), (*C.GKeyFile)(key_file.Ptr), group_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(group_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Set is a wrapper around gtk_print_settings_set().
func (settings PrintSettings) Set(key string, value string) {
	key0 := (*C.gchar)(C.CString(key))
	value0 := (*C.gchar)(C.CString(value))
	C.gtk_print_settings_set(settings.native(), key0, value0)
	C.free(unsafe.Pointer(key0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetBool is a wrapper around gtk_print_settings_set_bool().
func (settings PrintSettings) SetBool(key string, value bool) {
	key0 := (*C.gchar)(C.CString(key))
	C.gtk_print_settings_set_bool(settings.native(), key0, C.gboolean(util.Bool2Int(value)) /*go:.util*/)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// SetCollate is a wrapper around gtk_print_settings_set_collate().
func (settings PrintSettings) SetCollate(collate bool) {
	C.gtk_print_settings_set_collate(settings.native(), C.gboolean(util.Bool2Int(collate)) /*go:.util*/)
}

// SetDefaultSource is a wrapper around gtk_print_settings_set_default_source().
func (settings PrintSettings) SetDefaultSource(default_source string) {
	default_source0 := (*C.gchar)(C.CString(default_source))
	C.gtk_print_settings_set_default_source(settings.native(), default_source0)
	C.free(unsafe.Pointer(default_source0)) /*ch:<stdlib.h>*/
}

// SetDither is a wrapper around gtk_print_settings_set_dither().
func (settings PrintSettings) SetDither(dither string) {
	dither0 := (*C.gchar)(C.CString(dither))
	C.gtk_print_settings_set_dither(settings.native(), dither0)
	C.free(unsafe.Pointer(dither0)) /*ch:<stdlib.h>*/
}

// SetDouble is a wrapper around gtk_print_settings_set_double().
func (settings PrintSettings) SetDouble(key string, value float64) {
	key0 := (*C.gchar)(C.CString(key))
	C.gtk_print_settings_set_double(settings.native(), key0, C.gdouble(value))
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// SetDuplex is a wrapper around gtk_print_settings_set_duplex().
func (settings PrintSettings) SetDuplex(duplex PrintDuplex) {
	C.gtk_print_settings_set_duplex(settings.native(), C.GtkPrintDuplex(duplex))
}

// SetFinishings is a wrapper around gtk_print_settings_set_finishings().
func (settings PrintSettings) SetFinishings(finishings string) {
	finishings0 := (*C.gchar)(C.CString(finishings))
	C.gtk_print_settings_set_finishings(settings.native(), finishings0)
	C.free(unsafe.Pointer(finishings0)) /*ch:<stdlib.h>*/
}

// SetInt is a wrapper around gtk_print_settings_set_int().
func (settings PrintSettings) SetInt(key string, value int) {
	key0 := (*C.gchar)(C.CString(key))
	C.gtk_print_settings_set_int(settings.native(), key0, C.gint(value))
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// SetLength is a wrapper around gtk_print_settings_set_length().
func (settings PrintSettings) SetLength(key string, value float64, unit Unit) {
	key0 := (*C.gchar)(C.CString(key))
	C.gtk_print_settings_set_length(settings.native(), key0, C.gdouble(value), C.GtkUnit(unit))
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// SetMediaType is a wrapper around gtk_print_settings_set_media_type().
func (settings PrintSettings) SetMediaType(media_type string) {
	media_type0 := (*C.gchar)(C.CString(media_type))
	C.gtk_print_settings_set_media_type(settings.native(), media_type0)
	C.free(unsafe.Pointer(media_type0)) /*ch:<stdlib.h>*/
}

// SetNCopies is a wrapper around gtk_print_settings_set_n_copies().
func (settings PrintSettings) SetNCopies(num_copies int) {
	C.gtk_print_settings_set_n_copies(settings.native(), C.gint(num_copies))
}

// SetNumberUp is a wrapper around gtk_print_settings_set_number_up().
func (settings PrintSettings) SetNumberUp(number_up int) {
	C.gtk_print_settings_set_number_up(settings.native(), C.gint(number_up))
}

// SetNumberUpLayout is a wrapper around gtk_print_settings_set_number_up_layout().
func (settings PrintSettings) SetNumberUpLayout(number_up_layout NumberUpLayout) {
	C.gtk_print_settings_set_number_up_layout(settings.native(), C.GtkNumberUpLayout(number_up_layout))
}

// SetOrientation is a wrapper around gtk_print_settings_set_orientation().
func (settings PrintSettings) SetOrientation(orientation PageOrientation) {
	C.gtk_print_settings_set_orientation(settings.native(), C.GtkPageOrientation(orientation))
}

// SetOutputBin is a wrapper around gtk_print_settings_set_output_bin().
func (settings PrintSettings) SetOutputBin(output_bin string) {
	output_bin0 := (*C.gchar)(C.CString(output_bin))
	C.gtk_print_settings_set_output_bin(settings.native(), output_bin0)
	C.free(unsafe.Pointer(output_bin0)) /*ch:<stdlib.h>*/
}

// SetPageSet is a wrapper around gtk_print_settings_set_page_set().
func (settings PrintSettings) SetPageSet(page_set PageSet) {
	C.gtk_print_settings_set_page_set(settings.native(), C.GtkPageSet(page_set))
}

// SetPaperHeight is a wrapper around gtk_print_settings_set_paper_height().
func (settings PrintSettings) SetPaperHeight(height float64, unit Unit) {
	C.gtk_print_settings_set_paper_height(settings.native(), C.gdouble(height), C.GtkUnit(unit))
}

// SetPaperSize is a wrapper around gtk_print_settings_set_paper_size().
func (settings PrintSettings) SetPaperSize(paper_size PaperSize) {
	C.gtk_print_settings_set_paper_size(settings.native(), paper_size.native())
}

// SetPaperWidth is a wrapper around gtk_print_settings_set_paper_width().
func (settings PrintSettings) SetPaperWidth(width float64, unit Unit) {
	C.gtk_print_settings_set_paper_width(settings.native(), C.gdouble(width), C.GtkUnit(unit))
}

// SetPrintPages is a wrapper around gtk_print_settings_set_print_pages().
func (settings PrintSettings) SetPrintPages(pages PrintPages) {
	C.gtk_print_settings_set_print_pages(settings.native(), C.GtkPrintPages(pages))
}

// SetPrinter is a wrapper around gtk_print_settings_set_printer().
func (settings PrintSettings) SetPrinter(printer string) {
	printer0 := (*C.gchar)(C.CString(printer))
	C.gtk_print_settings_set_printer(settings.native(), printer0)
	C.free(unsafe.Pointer(printer0)) /*ch:<stdlib.h>*/
}

// SetPrinterLpi is a wrapper around gtk_print_settings_set_printer_lpi().
func (settings PrintSettings) SetPrinterLpi(lpi float64) {
	C.gtk_print_settings_set_printer_lpi(settings.native(), C.gdouble(lpi))
}

// SetQuality is a wrapper around gtk_print_settings_set_quality().
func (settings PrintSettings) SetQuality(quality PrintQuality) {
	C.gtk_print_settings_set_quality(settings.native(), C.GtkPrintQuality(quality))
}

// SetResolution is a wrapper around gtk_print_settings_set_resolution().
func (settings PrintSettings) SetResolution(resolution int) {
	C.gtk_print_settings_set_resolution(settings.native(), C.gint(resolution))
}

// SetResolutionXy is a wrapper around gtk_print_settings_set_resolution_xy().
func (settings PrintSettings) SetResolutionXy(resolution_x int, resolution_y int) {
	C.gtk_print_settings_set_resolution_xy(settings.native(), C.gint(resolution_x), C.gint(resolution_y))
}

// SetReverse is a wrapper around gtk_print_settings_set_reverse().
func (settings PrintSettings) SetReverse(reverse bool) {
	C.gtk_print_settings_set_reverse(settings.native(), C.gboolean(util.Bool2Int(reverse)) /*go:.util*/)
}

// SetScale is a wrapper around gtk_print_settings_set_scale().
func (settings PrintSettings) SetScale(scale float64) {
	C.gtk_print_settings_set_scale(settings.native(), C.gdouble(scale))
}

// SetUseColor is a wrapper around gtk_print_settings_set_use_color().
func (settings PrintSettings) SetUseColor(use_color bool) {
	C.gtk_print_settings_set_use_color(settings.native(), C.gboolean(util.Bool2Int(use_color)) /*go:.util*/)
}

// ToFile is a wrapper around gtk_print_settings_to_file().
func (settings PrintSettings) ToFile(file_name string) (bool, error) {
	file_name0 := (*C.gchar)(C.CString(file_name))
	var err glib.Error
	ret0 := C.gtk_print_settings_to_file(settings.native(), file_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ToGvariant is a wrapper around gtk_print_settings_to_gvariant().
func (settings PrintSettings) ToGvariant() glib.Variant {
	ret0 := C.gtk_print_settings_to_gvariant(settings.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// Unset is a wrapper around gtk_print_settings_unset().
func (settings PrintSettings) Unset(key string) {
	key0 := (*C.gchar)(C.CString(key))
	C.gtk_print_settings_unset(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// Object ProgressBar
type ProgressBar struct {
	Widget
}

func (v ProgressBar) native() *C.GtkProgressBar {
	return (*C.GtkProgressBar)(v.Ptr)
}
func wrapProgressBar(p *C.GtkProgressBar) (v ProgressBar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapProgressBar(p unsafe.Pointer) (v ProgressBar) {
	v.Ptr = p
	return
}
func (v ProgressBar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapProgressBar(p unsafe.Pointer) interface{} {
	return WrapProgressBar(p)
}
func (v ProgressBar) GetType() gobject.Type {
	return gobject.Type(C.gtk_progress_bar_get_type())
}
func (v ProgressBar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapProgressBar(unsafe.Pointer(ptr)), nil
	}
}
func (v ProgressBar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ProgressBar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ProgressBar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// ProgressBarNew is a wrapper around gtk_progress_bar_new().
func ProgressBarNew() Widget {
	ret0 := C.gtk_progress_bar_new()
	return wrapWidget(ret0)
}

// GetEllipsize is a wrapper around gtk_progress_bar_get_ellipsize().
func (pbar ProgressBar) GetEllipsize() /*gir:Pango*/ pango.EllipsizeMode {
	ret0 := C.gtk_progress_bar_get_ellipsize(pbar.native())
	return /*gir:Pango*/ pango.EllipsizeMode(ret0)
}

// GetFraction is a wrapper around gtk_progress_bar_get_fraction().
func (pbar ProgressBar) GetFraction() float64 {
	ret0 := C.gtk_progress_bar_get_fraction(pbar.native())
	return float64(ret0)
}

// GetInverted is a wrapper around gtk_progress_bar_get_inverted().
func (pbar ProgressBar) GetInverted() bool {
	ret0 := C.gtk_progress_bar_get_inverted(pbar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPulseStep is a wrapper around gtk_progress_bar_get_pulse_step().
func (pbar ProgressBar) GetPulseStep() float64 {
	ret0 := C.gtk_progress_bar_get_pulse_step(pbar.native())
	return float64(ret0)
}

// GetShowText is a wrapper around gtk_progress_bar_get_show_text().
func (pbar ProgressBar) GetShowText() bool {
	ret0 := C.gtk_progress_bar_get_show_text(pbar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetText is a wrapper around gtk_progress_bar_get_text().
func (pbar ProgressBar) GetText() string {
	ret0 := C.gtk_progress_bar_get_text(pbar.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Pulse is a wrapper around gtk_progress_bar_pulse().
func (pbar ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(pbar.native())
}

// SetEllipsize is a wrapper around gtk_progress_bar_set_ellipsize().
func (pbar ProgressBar) SetEllipsize(mode /*gir:Pango*/ pango.EllipsizeMode) {
	C.gtk_progress_bar_set_ellipsize(pbar.native(), C.PangoEllipsizeMode(mode))
}

// SetFraction is a wrapper around gtk_progress_bar_set_fraction().
func (pbar ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(pbar.native(), C.gdouble(fraction))
}

// SetInverted is a wrapper around gtk_progress_bar_set_inverted().
func (pbar ProgressBar) SetInverted(inverted bool) {
	C.gtk_progress_bar_set_inverted(pbar.native(), C.gboolean(util.Bool2Int(inverted)) /*go:.util*/)
}

// SetPulseStep is a wrapper around gtk_progress_bar_set_pulse_step().
func (pbar ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(pbar.native(), C.gdouble(fraction))
}

// SetShowText is a wrapper around gtk_progress_bar_set_show_text().
func (pbar ProgressBar) SetShowText(show_text bool) {
	C.gtk_progress_bar_set_show_text(pbar.native(), C.gboolean(util.Bool2Int(show_text)) /*go:.util*/)
}

// SetText is a wrapper around gtk_progress_bar_set_text().
func (pbar ProgressBar) SetText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_progress_bar_set_text(pbar.native(), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// Object ProgressBarAccessible
type ProgressBarAccessible struct {
	WidgetAccessible
}

func (v ProgressBarAccessible) native() *C.GtkProgressBarAccessible {
	return (*C.GtkProgressBarAccessible)(v.Ptr)
}
func wrapProgressBarAccessible(p *C.GtkProgressBarAccessible) (v ProgressBarAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapProgressBarAccessible(p unsafe.Pointer) (v ProgressBarAccessible) {
	v.Ptr = p
	return
}
func (v ProgressBarAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapProgressBarAccessible(p unsafe.Pointer) interface{} {
	return WrapProgressBarAccessible(p)
}
func (v ProgressBarAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_progress_bar_accessible_get_type())
}
func (v ProgressBarAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapProgressBarAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ProgressBarAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ProgressBarAccessible) Value() atk.Value {
	return atk.WrapValue(v.Ptr) /*gir:Atk*/
}

// Object RadioAction
type RadioAction struct {
	ToggleAction
}

func (v RadioAction) native() *C.GtkRadioAction {
	return (*C.GtkRadioAction)(v.Ptr)
}
func wrapRadioAction(p *C.GtkRadioAction) (v RadioAction) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRadioAction(p unsafe.Pointer) (v RadioAction) {
	v.Ptr = p
	return
}
func (v RadioAction) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRadioAction(p unsafe.Pointer) interface{} {
	return WrapRadioAction(p)
}
func (v RadioAction) GetType() gobject.Type {
	return gobject.Type(C.gtk_radio_action_get_type())
}
func (v RadioAction) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRadioAction(unsafe.Pointer(ptr)), nil
	}
}
func (v RadioAction) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object ToggleAction
type ToggleAction struct {
	Action
}

func (v ToggleAction) native() *C.GtkToggleAction {
	return (*C.GtkToggleAction)(v.Ptr)
}
func wrapToggleAction(p *C.GtkToggleAction) (v ToggleAction) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToggleAction(p unsafe.Pointer) (v ToggleAction) {
	v.Ptr = p
	return
}
func (v ToggleAction) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToggleAction(p unsafe.Pointer) interface{} {
	return WrapToggleAction(p)
}
func (v ToggleAction) GetType() gobject.Type {
	return gobject.Type(C.gtk_toggle_action_get_type())
}
func (v ToggleAction) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToggleAction(unsafe.Pointer(ptr)), nil
	}
}
func (v ToggleAction) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object Action
type Action struct {
	gobject.Object
}

func (v Action) native() *C.GtkAction {
	return (*C.GtkAction)(v.Ptr)
}
func wrapAction(p *C.GtkAction) (v Action) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAction(p unsafe.Pointer) (v Action) {
	v.Ptr = p
	return
}
func (v Action) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAction(p unsafe.Pointer) interface{} {
	return WrapAction(p)
}
func (v Action) GetType() gobject.Type {
	return gobject.Type(C.gtk_action_get_type())
}
func (v Action) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAction(unsafe.Pointer(ptr)), nil
	}
}
func (v Action) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object RadioButton
type RadioButton struct {
	CheckButton
}

func (v RadioButton) native() *C.GtkRadioButton {
	return (*C.GtkRadioButton)(v.Ptr)
}
func wrapRadioButton(p *C.GtkRadioButton) (v RadioButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRadioButton(p unsafe.Pointer) (v RadioButton) {
	v.Ptr = p
	return
}
func (v RadioButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRadioButton(p unsafe.Pointer) interface{} {
	return WrapRadioButton(p)
}
func (v RadioButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_radio_button_get_type())
}
func (v RadioButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRadioButton(unsafe.Pointer(ptr)), nil
	}
}
func (v RadioButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v RadioButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v RadioButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v RadioButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// RadioButtonNewFromWidget is a wrapper around gtk_radio_button_new_from_widget().
func RadioButtonNewFromWidget(radio_group_member RadioButton) Widget {
	ret0 := C.gtk_radio_button_new_from_widget(radio_group_member.native())
	return wrapWidget(ret0)
}

// RadioButtonNewWithLabelFromWidget is a wrapper around gtk_radio_button_new_with_label_from_widget().
func RadioButtonNewWithLabelFromWidget(radio_group_member RadioButton, label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_radio_button_new_with_label_from_widget(radio_group_member.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// RadioButtonNewWithMnemonicFromWidget is a wrapper around gtk_radio_button_new_with_mnemonic_from_widget().
func RadioButtonNewWithMnemonicFromWidget(radio_group_member RadioButton, label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_radio_button_new_with_mnemonic_from_widget(radio_group_member.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// JoinGroup is a wrapper around gtk_radio_button_join_group().
func (radio_button RadioButton) JoinGroup(group_source RadioButton) {
	C.gtk_radio_button_join_group(radio_button.native(), group_source.native())
}

// Object RadioButtonAccessible
type RadioButtonAccessible struct {
	ToggleButtonAccessible
}

func (v RadioButtonAccessible) native() *C.GtkRadioButtonAccessible {
	return (*C.GtkRadioButtonAccessible)(v.Ptr)
}
func wrapRadioButtonAccessible(p *C.GtkRadioButtonAccessible) (v RadioButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRadioButtonAccessible(p unsafe.Pointer) (v RadioButtonAccessible) {
	v.Ptr = p
	return
}
func (v RadioButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRadioButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapRadioButtonAccessible(p)
}
func (v RadioButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_radio_button_accessible_get_type())
}
func (v RadioButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRadioButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v RadioButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v RadioButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v RadioButtonAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object RadioMenuItem
type RadioMenuItem struct {
	CheckMenuItem
}

func (v RadioMenuItem) native() *C.GtkRadioMenuItem {
	return (*C.GtkRadioMenuItem)(v.Ptr)
}
func wrapRadioMenuItem(p *C.GtkRadioMenuItem) (v RadioMenuItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRadioMenuItem(p unsafe.Pointer) (v RadioMenuItem) {
	v.Ptr = p
	return
}
func (v RadioMenuItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRadioMenuItem(p unsafe.Pointer) interface{} {
	return WrapRadioMenuItem(p)
}
func (v RadioMenuItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_radio_menu_item_get_type())
}
func (v RadioMenuItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRadioMenuItem(unsafe.Pointer(ptr)), nil
	}
}
func (v RadioMenuItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v RadioMenuItem) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v RadioMenuItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v RadioMenuItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// RadioMenuItemNewFromWidget is a wrapper around gtk_radio_menu_item_new_from_widget().
func RadioMenuItemNewFromWidget(group RadioMenuItem) Widget {
	ret0 := C.gtk_radio_menu_item_new_from_widget(group.native())
	return wrapWidget(ret0)
}

// RadioMenuItemNewWithLabelFromWidget is a wrapper around gtk_radio_menu_item_new_with_label_from_widget().
func RadioMenuItemNewWithLabelFromWidget(group RadioMenuItem, label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_radio_menu_item_new_with_label_from_widget(group.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// RadioMenuItemNewWithMnemonicFromWidget is a wrapper around gtk_radio_menu_item_new_with_mnemonic_from_widget().
func RadioMenuItemNewWithMnemonicFromWidget(group RadioMenuItem, label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(group.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// JoinGroup is a wrapper around gtk_radio_menu_item_join_group().
func (radio_menu_item RadioMenuItem) JoinGroup(group_source RadioMenuItem) {
	C.gtk_radio_menu_item_join_group(radio_menu_item.native(), group_source.native())
}

// Object RadioMenuItemAccessible
type RadioMenuItemAccessible struct {
	CheckMenuItemAccessible
}

func (v RadioMenuItemAccessible) native() *C.GtkRadioMenuItemAccessible {
	return (*C.GtkRadioMenuItemAccessible)(v.Ptr)
}
func wrapRadioMenuItemAccessible(p *C.GtkRadioMenuItemAccessible) (v RadioMenuItemAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRadioMenuItemAccessible(p unsafe.Pointer) (v RadioMenuItemAccessible) {
	v.Ptr = p
	return
}
func (v RadioMenuItemAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRadioMenuItemAccessible(p unsafe.Pointer) interface{} {
	return WrapRadioMenuItemAccessible(p)
}
func (v RadioMenuItemAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_radio_menu_item_accessible_get_type())
}
func (v RadioMenuItemAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRadioMenuItemAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v RadioMenuItemAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v RadioMenuItemAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v RadioMenuItemAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}

// Object ToggleToolButton
type ToggleToolButton struct {
	ToolButton
}

func (v ToggleToolButton) native() *C.GtkToggleToolButton {
	return (*C.GtkToggleToolButton)(v.Ptr)
}
func wrapToggleToolButton(p *C.GtkToggleToolButton) (v ToggleToolButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToggleToolButton(p unsafe.Pointer) (v ToggleToolButton) {
	v.Ptr = p
	return
}
func (v ToggleToolButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToggleToolButton(p unsafe.Pointer) interface{} {
	return WrapToggleToolButton(p)
}
func (v ToggleToolButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_toggle_tool_button_get_type())
}
func (v ToggleToolButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToggleToolButton(unsafe.Pointer(ptr)), nil
	}
}
func (v ToggleToolButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ToggleToolButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v ToggleToolButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ToggleToolButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// ToggleToolButtonNew is a wrapper around gtk_toggle_tool_button_new().
func ToggleToolButtonNew() ToolItem {
	ret0 := C.gtk_toggle_tool_button_new()
	return wrapToolItem(ret0)
}

// GetActive is a wrapper around gtk_toggle_tool_button_get_active().
func (button ToggleToolButton) GetActive() bool {
	ret0 := C.gtk_toggle_tool_button_get_active(button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetActive is a wrapper around gtk_toggle_tool_button_set_active().
func (button ToggleToolButton) SetActive(is_active bool) {
	C.gtk_toggle_tool_button_set_active(button.native(), C.gboolean(util.Bool2Int(is_active)) /*go:.util*/)
}

// Object RadioToolButton
type RadioToolButton struct {
	ToggleToolButton
}

func (v RadioToolButton) native() *C.GtkRadioToolButton {
	return (*C.GtkRadioToolButton)(v.Ptr)
}
func wrapRadioToolButton(p *C.GtkRadioToolButton) (v RadioToolButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRadioToolButton(p unsafe.Pointer) (v RadioToolButton) {
	v.Ptr = p
	return
}
func (v RadioToolButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRadioToolButton(p unsafe.Pointer) interface{} {
	return WrapRadioToolButton(p)
}
func (v RadioToolButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_radio_tool_button_get_type())
}
func (v RadioToolButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRadioToolButton(unsafe.Pointer(ptr)), nil
	}
}
func (v RadioToolButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v RadioToolButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v RadioToolButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v RadioToolButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// RadioToolButtonNewFromWidget is a wrapper around gtk_radio_tool_button_new_from_widget().
func RadioToolButtonNewFromWidget(group RadioToolButton) ToolItem {
	ret0 := C.gtk_radio_tool_button_new_from_widget(group.native())
	return wrapToolItem(ret0)
}

// Object RangeAccessible
type RangeAccessible struct {
	WidgetAccessible
}

func (v RangeAccessible) native() *C.GtkRangeAccessible {
	return (*C.GtkRangeAccessible)(v.Ptr)
}
func wrapRangeAccessible(p *C.GtkRangeAccessible) (v RangeAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRangeAccessible(p unsafe.Pointer) (v RangeAccessible) {
	v.Ptr = p
	return
}
func (v RangeAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRangeAccessible(p unsafe.Pointer) interface{} {
	return WrapRangeAccessible(p)
}
func (v RangeAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_range_accessible_get_type())
}
func (v RangeAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRangeAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v RangeAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v RangeAccessible) Value() atk.Value {
	return atk.WrapValue(v.Ptr) /*gir:Atk*/
}

// Object RcStyle
type RcStyle struct {
	gobject.Object
}

func (v RcStyle) native() *C.GtkRcStyle {
	return (*C.GtkRcStyle)(v.Ptr)
}
func wrapRcStyle(p *C.GtkRcStyle) (v RcStyle) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRcStyle(p unsafe.Pointer) (v RcStyle) {
	v.Ptr = p
	return
}
func (v RcStyle) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRcStyle(p unsafe.Pointer) interface{} {
	return WrapRcStyle(p)
}
func (v RcStyle) GetType() gobject.Type {
	return gobject.Type(C.gtk_rc_style_get_type())
}
func (v RcStyle) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRcStyle(unsafe.Pointer(ptr)), nil
	}
}

// Object RecentAction
type RecentAction struct {
	Action
}

func (v RecentAction) native() *C.GtkRecentAction {
	return (*C.GtkRecentAction)(v.Ptr)
}
func wrapRecentAction(p *C.GtkRecentAction) (v RecentAction) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRecentAction(p unsafe.Pointer) (v RecentAction) {
	v.Ptr = p
	return
}
func (v RecentAction) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentAction(p unsafe.Pointer) interface{} {
	return WrapRecentAction(p)
}
func (v RecentAction) GetType() gobject.Type {
	return gobject.Type(C.gtk_recent_action_get_type())
}
func (v RecentAction) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRecentAction(unsafe.Pointer(ptr)), nil
	}
}
func (v RecentAction) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v RecentAction) RecentChooser() RecentChooser {
	return WrapRecentChooser(v.Ptr)
}

// Object RecentChooserDialog
type RecentChooserDialog struct {
	Dialog
}

func (v RecentChooserDialog) native() *C.GtkRecentChooserDialog {
	return (*C.GtkRecentChooserDialog)(v.Ptr)
}
func wrapRecentChooserDialog(p *C.GtkRecentChooserDialog) (v RecentChooserDialog) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRecentChooserDialog(p unsafe.Pointer) (v RecentChooserDialog) {
	v.Ptr = p
	return
}
func (v RecentChooserDialog) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentChooserDialog(p unsafe.Pointer) interface{} {
	return WrapRecentChooserDialog(p)
}
func (v RecentChooserDialog) GetType() gobject.Type {
	return gobject.Type(C.gtk_recent_chooser_dialog_get_type())
}
func (v RecentChooserDialog) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRecentChooserDialog(unsafe.Pointer(ptr)), nil
	}
}
func (v RecentChooserDialog) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v RecentChooserDialog) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v RecentChooserDialog) RecentChooser() RecentChooser {
	return WrapRecentChooser(v.Ptr)
}

// Object RecentChooserMenu
type RecentChooserMenu struct {
	Menu
}

func (v RecentChooserMenu) native() *C.GtkRecentChooserMenu {
	return (*C.GtkRecentChooserMenu)(v.Ptr)
}
func wrapRecentChooserMenu(p *C.GtkRecentChooserMenu) (v RecentChooserMenu) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRecentChooserMenu(p unsafe.Pointer) (v RecentChooserMenu) {
	v.Ptr = p
	return
}
func (v RecentChooserMenu) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentChooserMenu(p unsafe.Pointer) interface{} {
	return WrapRecentChooserMenu(p)
}
func (v RecentChooserMenu) GetType() gobject.Type {
	return gobject.Type(C.gtk_recent_chooser_menu_get_type())
}
func (v RecentChooserMenu) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRecentChooserMenu(unsafe.Pointer(ptr)), nil
	}
}
func (v RecentChooserMenu) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v RecentChooserMenu) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v RecentChooserMenu) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v RecentChooserMenu) RecentChooser() RecentChooser {
	return WrapRecentChooser(v.Ptr)
}

// RecentChooserMenuNew is a wrapper around gtk_recent_chooser_menu_new().
func RecentChooserMenuNew() Widget {
	ret0 := C.gtk_recent_chooser_menu_new()
	return wrapWidget(ret0)
}

// RecentChooserMenuNewForManager is a wrapper around gtk_recent_chooser_menu_new_for_manager().
func RecentChooserMenuNewForManager(manager RecentManager) Widget {
	ret0 := C.gtk_recent_chooser_menu_new_for_manager(manager.native())
	return wrapWidget(ret0)
}

// GetShowNumbers is a wrapper around gtk_recent_chooser_menu_get_show_numbers().
func (menu RecentChooserMenu) GetShowNumbers() bool {
	ret0 := C.gtk_recent_chooser_menu_get_show_numbers(menu.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetShowNumbers is a wrapper around gtk_recent_chooser_menu_set_show_numbers().
func (menu RecentChooserMenu) SetShowNumbers(show_numbers bool) {
	C.gtk_recent_chooser_menu_set_show_numbers(menu.native(), C.gboolean(util.Bool2Int(show_numbers)) /*go:.util*/)
}

// Object RecentManager
type RecentManager struct {
	gobject.Object
}

func (v RecentManager) native() *C.GtkRecentManager {
	return (*C.GtkRecentManager)(v.Ptr)
}
func wrapRecentManager(p *C.GtkRecentManager) (v RecentManager) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRecentManager(p unsafe.Pointer) (v RecentManager) {
	v.Ptr = p
	return
}
func (v RecentManager) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentManager(p unsafe.Pointer) interface{} {
	return WrapRecentManager(p)
}
func (v RecentManager) GetType() gobject.Type {
	return gobject.Type(C.gtk_recent_manager_get_type())
}
func (v RecentManager) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRecentManager(unsafe.Pointer(ptr)), nil
	}
}

// RecentManagerNew is a wrapper around gtk_recent_manager_new().
func RecentManagerNew() RecentManager {
	ret0 := C.gtk_recent_manager_new()
	return wrapRecentManager(ret0)
}

// AddFull is a wrapper around gtk_recent_manager_add_full().
func (manager RecentManager) AddFull(uri string, recent_data RecentData) bool {
	uri0 := (*C.gchar)(C.CString(uri))
	ret0 := C.gtk_recent_manager_add_full(manager.native(), uri0, recent_data.native())
	C.free(unsafe.Pointer(uri0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// AddItem is a wrapper around gtk_recent_manager_add_item().
func (manager RecentManager) AddItem(uri string) bool {
	uri0 := (*C.gchar)(C.CString(uri))
	ret0 := C.gtk_recent_manager_add_item(manager.native(), uri0)
	C.free(unsafe.Pointer(uri0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetItems is a wrapper around gtk_recent_manager_get_items().
func (manager RecentManager) GetItems() glib.List {
	ret0 := C.gtk_recent_manager_get_items(manager.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapRecentInfo(p) }) /*gir:GLib*/
}

// HasItem is a wrapper around gtk_recent_manager_has_item().
func (manager RecentManager) HasItem(uri string) bool {
	uri0 := (*C.gchar)(C.CString(uri))
	ret0 := C.gtk_recent_manager_has_item(manager.native(), uri0)
	C.free(unsafe.Pointer(uri0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LookupItem is a wrapper around gtk_recent_manager_lookup_item().
func (manager RecentManager) LookupItem(uri string) (RecentInfo, error) {
	uri0 := (*C.gchar)(C.CString(uri))
	var err glib.Error
	ret0 := C.gtk_recent_manager_lookup_item(manager.native(), uri0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return RecentInfo{}, err.GoValue()
	}
	return wrapRecentInfo(ret0), nil
}

// MoveItem is a wrapper around gtk_recent_manager_move_item().
func (manager RecentManager) MoveItem(uri string, new_uri string) (bool, error) {
	uri0 := (*C.gchar)(C.CString(uri))
	new_uri0 := (*C.gchar)(C.CString(new_uri))
	var err glib.Error
	ret0 := C.gtk_recent_manager_move_item(manager.native(), uri0, new_uri0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0))     /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(new_uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PurgeItems is a wrapper around gtk_recent_manager_purge_items().
func (manager RecentManager) PurgeItems() (int, error) {
	var err glib.Error
	ret0 := C.gtk_recent_manager_purge_items(manager.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// RemoveItem is a wrapper around gtk_recent_manager_remove_item().
func (manager RecentManager) RemoveItem(uri string) (bool, error) {
	uri0 := (*C.gchar)(C.CString(uri))
	var err glib.Error
	ret0 := C.gtk_recent_manager_remove_item(manager.native(), uri0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// RecentManagerGetDefault is a wrapper around gtk_recent_manager_get_default().
func RecentManagerGetDefault() RecentManager {
	ret0 := C.gtk_recent_manager_get_default()
	return wrapRecentManager(ret0)
}

// Object RecentChooserWidget
type RecentChooserWidget struct {
	Box
}

func (v RecentChooserWidget) native() *C.GtkRecentChooserWidget {
	return (*C.GtkRecentChooserWidget)(v.Ptr)
}
func wrapRecentChooserWidget(p *C.GtkRecentChooserWidget) (v RecentChooserWidget) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRecentChooserWidget(p unsafe.Pointer) (v RecentChooserWidget) {
	v.Ptr = p
	return
}
func (v RecentChooserWidget) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRecentChooserWidget(p unsafe.Pointer) interface{} {
	return WrapRecentChooserWidget(p)
}
func (v RecentChooserWidget) GetType() gobject.Type {
	return gobject.Type(C.gtk_recent_chooser_widget_get_type())
}
func (v RecentChooserWidget) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRecentChooserWidget(unsafe.Pointer(ptr)), nil
	}
}
func (v RecentChooserWidget) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v RecentChooserWidget) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v RecentChooserWidget) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}
func (v RecentChooserWidget) RecentChooser() RecentChooser {
	return WrapRecentChooser(v.Ptr)
}

// RecentChooserWidgetNew is a wrapper around gtk_recent_chooser_widget_new().
func RecentChooserWidgetNew() Widget {
	ret0 := C.gtk_recent_chooser_widget_new()
	return wrapWidget(ret0)
}

// RecentChooserWidgetNewForManager is a wrapper around gtk_recent_chooser_widget_new_for_manager().
func RecentChooserWidgetNewForManager(manager RecentManager) Widget {
	ret0 := C.gtk_recent_chooser_widget_new_for_manager(manager.native())
	return wrapWidget(ret0)
}

// Object Revealer
type Revealer struct {
	Bin
}

func (v Revealer) native() *C.GtkRevealer {
	return (*C.GtkRevealer)(v.Ptr)
}
func wrapRevealer(p *C.GtkRevealer) (v Revealer) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRevealer(p unsafe.Pointer) (v Revealer) {
	v.Ptr = p
	return
}
func (v Revealer) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRevealer(p unsafe.Pointer) interface{} {
	return WrapRevealer(p)
}
func (v Revealer) GetType() gobject.Type {
	return gobject.Type(C.gtk_revealer_get_type())
}
func (v Revealer) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRevealer(unsafe.Pointer(ptr)), nil
	}
}
func (v Revealer) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Revealer) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// RevealerNew is a wrapper around gtk_revealer_new().
func RevealerNew() Widget {
	ret0 := C.gtk_revealer_new()
	return wrapWidget(ret0)
}

// GetChildRevealed is a wrapper around gtk_revealer_get_child_revealed().
func (revealer Revealer) GetChildRevealed() bool {
	ret0 := C.gtk_revealer_get_child_revealed(revealer.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRevealChild is a wrapper around gtk_revealer_get_reveal_child().
func (revealer Revealer) GetRevealChild() bool {
	ret0 := C.gtk_revealer_get_reveal_child(revealer.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTransitionDuration is a wrapper around gtk_revealer_get_transition_duration().
func (revealer Revealer) GetTransitionDuration() uint {
	ret0 := C.gtk_revealer_get_transition_duration(revealer.native())
	return uint(ret0)
}

// GetTransitionType is a wrapper around gtk_revealer_get_transition_type().
func (revealer Revealer) GetTransitionType() RevealerTransitionType {
	ret0 := C.gtk_revealer_get_transition_type(revealer.native())
	return RevealerTransitionType(ret0)
}

// SetRevealChild is a wrapper around gtk_revealer_set_reveal_child().
func (revealer Revealer) SetRevealChild(reveal_child bool) {
	C.gtk_revealer_set_reveal_child(revealer.native(), C.gboolean(util.Bool2Int(reveal_child)) /*go:.util*/)
}

// SetTransitionDuration is a wrapper around gtk_revealer_set_transition_duration().
func (revealer Revealer) SetTransitionDuration(duration uint) {
	C.gtk_revealer_set_transition_duration(revealer.native(), C.guint(duration))
}

// SetTransitionType is a wrapper around gtk_revealer_set_transition_type().
func (revealer Revealer) SetTransitionType(transition RevealerTransitionType) {
	C.gtk_revealer_set_transition_type(revealer.native(), C.GtkRevealerTransitionType(transition))
}

// Object ScaleAccessible
type ScaleAccessible struct {
	RangeAccessible
}

func (v ScaleAccessible) native() *C.GtkScaleAccessible {
	return (*C.GtkScaleAccessible)(v.Ptr)
}
func wrapScaleAccessible(p *C.GtkScaleAccessible) (v ScaleAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScaleAccessible(p unsafe.Pointer) (v ScaleAccessible) {
	v.Ptr = p
	return
}
func (v ScaleAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScaleAccessible(p unsafe.Pointer) interface{} {
	return WrapScaleAccessible(p)
}
func (v ScaleAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_scale_accessible_get_type())
}
func (v ScaleAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScaleAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ScaleAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ScaleAccessible) Value() atk.Value {
	return atk.WrapValue(v.Ptr) /*gir:Atk*/
}

// Object ScaleButton
type ScaleButton struct {
	Button
}

func (v ScaleButton) native() *C.GtkScaleButton {
	return (*C.GtkScaleButton)(v.Ptr)
}
func wrapScaleButton(p *C.GtkScaleButton) (v ScaleButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScaleButton(p unsafe.Pointer) (v ScaleButton) {
	v.Ptr = p
	return
}
func (v ScaleButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScaleButton(p unsafe.Pointer) interface{} {
	return WrapScaleButton(p)
}
func (v ScaleButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_scale_button_get_type())
}
func (v ScaleButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScaleButton(unsafe.Pointer(ptr)), nil
	}
}
func (v ScaleButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ScaleButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v ScaleButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v ScaleButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ScaleButton) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// GetAdjustment is a wrapper around gtk_scale_button_get_adjustment().
func (button ScaleButton) GetAdjustment() Adjustment {
	ret0 := C.gtk_scale_button_get_adjustment(button.native())
	return wrapAdjustment(ret0)
}

// GetPopup is a wrapper around gtk_scale_button_get_popup().
func (button ScaleButton) GetPopup() Widget {
	ret0 := C.gtk_scale_button_get_popup(button.native())
	return wrapWidget(ret0)
}

// GetValue is a wrapper around gtk_scale_button_get_value().
func (button ScaleButton) GetValue() float64 {
	ret0 := C.gtk_scale_button_get_value(button.native())
	return float64(ret0)
}

// SetAdjustment is a wrapper around gtk_scale_button_set_adjustment().
func (button ScaleButton) SetAdjustment(adjustment Adjustment) {
	C.gtk_scale_button_set_adjustment(button.native(), adjustment.native())
}

// SetIcons is a wrapper around gtk_scale_button_set_icons().
func (button ScaleButton) SetIcons(icons []string) {
	icons0 := make([]*C.gchar, len(icons))
	for idx, elemG := range icons {
		elem := (*C.gchar)(C.CString(elemG))
		icons0[idx] = elem
	}
	var icons0Ptr **C.gchar
	if len(icons0) > 0 {
		icons0Ptr = &icons0[0]
	}
	C.gtk_scale_button_set_icons(button.native(), icons0Ptr)
	for _, elem := range icons0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// SetValue is a wrapper around gtk_scale_button_set_value().
func (button ScaleButton) SetValue(value float64) {
	C.gtk_scale_button_set_value(button.native(), C.gdouble(value))
}

// Object ScaleButtonAccessible
type ScaleButtonAccessible struct {
	ButtonAccessible
}

func (v ScaleButtonAccessible) native() *C.GtkScaleButtonAccessible {
	return (*C.GtkScaleButtonAccessible)(v.Ptr)
}
func wrapScaleButtonAccessible(p *C.GtkScaleButtonAccessible) (v ScaleButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScaleButtonAccessible(p unsafe.Pointer) (v ScaleButtonAccessible) {
	v.Ptr = p
	return
}
func (v ScaleButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScaleButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapScaleButtonAccessible(p)
}
func (v ScaleButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_scale_button_accessible_get_type())
}
func (v ScaleButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScaleButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ScaleButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v ScaleButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v ScaleButtonAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}
func (v ScaleButtonAccessible) Value() atk.Value {
	return atk.WrapValue(v.Ptr) /*gir:Atk*/
}

// Object ScrolledWindowAccessible
type ScrolledWindowAccessible struct {
	ContainerAccessible
}

func (v ScrolledWindowAccessible) native() *C.GtkScrolledWindowAccessible {
	return (*C.GtkScrolledWindowAccessible)(v.Ptr)
}
func wrapScrolledWindowAccessible(p *C.GtkScrolledWindowAccessible) (v ScrolledWindowAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScrolledWindowAccessible(p unsafe.Pointer) (v ScrolledWindowAccessible) {
	v.Ptr = p
	return
}
func (v ScrolledWindowAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScrolledWindowAccessible(p unsafe.Pointer) interface{} {
	return WrapScrolledWindowAccessible(p)
}
func (v ScrolledWindowAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_scrolled_window_accessible_get_type())
}
func (v ScrolledWindowAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScrolledWindowAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v ScrolledWindowAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object SearchBar
type SearchBar struct {
	Bin
}

func (v SearchBar) native() *C.GtkSearchBar {
	return (*C.GtkSearchBar)(v.Ptr)
}
func wrapSearchBar(p *C.GtkSearchBar) (v SearchBar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSearchBar(p unsafe.Pointer) (v SearchBar) {
	v.Ptr = p
	return
}
func (v SearchBar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSearchBar(p unsafe.Pointer) interface{} {
	return WrapSearchBar(p)
}
func (v SearchBar) GetType() gobject.Type {
	return gobject.Type(C.gtk_search_bar_get_type())
}
func (v SearchBar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSearchBar(unsafe.Pointer(ptr)), nil
	}
}
func (v SearchBar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v SearchBar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// SearchBarNew is a wrapper around gtk_search_bar_new().
func SearchBarNew() Widget {
	ret0 := C.gtk_search_bar_new()
	return wrapWidget(ret0)
}

// ConnectEntry is a wrapper around gtk_search_bar_connect_entry().
func (bar SearchBar) ConnectEntry(entry Entry) {
	C.gtk_search_bar_connect_entry(bar.native(), entry.native())
}

// GetSearchMode is a wrapper around gtk_search_bar_get_search_mode().
func (bar SearchBar) GetSearchMode() bool {
	ret0 := C.gtk_search_bar_get_search_mode(bar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowCloseButton is a wrapper around gtk_search_bar_get_show_close_button().
func (bar SearchBar) GetShowCloseButton() bool {
	ret0 := C.gtk_search_bar_get_show_close_button(bar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetSearchMode is a wrapper around gtk_search_bar_set_search_mode().
func (bar SearchBar) SetSearchMode(search_mode bool) {
	C.gtk_search_bar_set_search_mode(bar.native(), C.gboolean(util.Bool2Int(search_mode)) /*go:.util*/)
}

// SetShowCloseButton is a wrapper around gtk_search_bar_set_show_close_button().
func (bar SearchBar) SetShowCloseButton(visible bool) {
	C.gtk_search_bar_set_show_close_button(bar.native(), C.gboolean(util.Bool2Int(visible)) /*go:.util*/)
}

// Object SearchEntry
type SearchEntry struct {
	Entry
}

func (v SearchEntry) native() *C.GtkSearchEntry {
	return (*C.GtkSearchEntry)(v.Ptr)
}
func wrapSearchEntry(p *C.GtkSearchEntry) (v SearchEntry) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSearchEntry(p unsafe.Pointer) (v SearchEntry) {
	v.Ptr = p
	return
}
func (v SearchEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSearchEntry(p unsafe.Pointer) interface{} {
	return WrapSearchEntry(p)
}
func (v SearchEntry) GetType() gobject.Type {
	return gobject.Type(C.gtk_search_entry_get_type())
}
func (v SearchEntry) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSearchEntry(unsafe.Pointer(ptr)), nil
	}
}
func (v SearchEntry) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v SearchEntry) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v SearchEntry) CellEditable() CellEditable {
	return WrapCellEditable(v.Ptr)
}
func (v SearchEntry) Editable() Editable {
	return WrapEditable(v.Ptr)
}

// SearchEntryNew is a wrapper around gtk_search_entry_new().
func SearchEntryNew() Widget {
	ret0 := C.gtk_search_entry_new()
	return wrapWidget(ret0)
}

// Object SeparatorMenuItem
type SeparatorMenuItem struct {
	MenuItem
}

func (v SeparatorMenuItem) native() *C.GtkSeparatorMenuItem {
	return (*C.GtkSeparatorMenuItem)(v.Ptr)
}
func wrapSeparatorMenuItem(p *C.GtkSeparatorMenuItem) (v SeparatorMenuItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSeparatorMenuItem(p unsafe.Pointer) (v SeparatorMenuItem) {
	v.Ptr = p
	return
}
func (v SeparatorMenuItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSeparatorMenuItem(p unsafe.Pointer) interface{} {
	return WrapSeparatorMenuItem(p)
}
func (v SeparatorMenuItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_separator_menu_item_get_type())
}
func (v SeparatorMenuItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSeparatorMenuItem(unsafe.Pointer(ptr)), nil
	}
}
func (v SeparatorMenuItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v SeparatorMenuItem) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v SeparatorMenuItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v SeparatorMenuItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// SeparatorMenuItemNew is a wrapper around gtk_separator_menu_item_new().
func SeparatorMenuItemNew() Widget {
	ret0 := C.gtk_separator_menu_item_new()
	return wrapWidget(ret0)
}

// Object SeparatorToolItem
type SeparatorToolItem struct {
	ToolItem
}

func (v SeparatorToolItem) native() *C.GtkSeparatorToolItem {
	return (*C.GtkSeparatorToolItem)(v.Ptr)
}
func wrapSeparatorToolItem(p *C.GtkSeparatorToolItem) (v SeparatorToolItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSeparatorToolItem(p unsafe.Pointer) (v SeparatorToolItem) {
	v.Ptr = p
	return
}
func (v SeparatorToolItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSeparatorToolItem(p unsafe.Pointer) interface{} {
	return WrapSeparatorToolItem(p)
}
func (v SeparatorToolItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_separator_tool_item_get_type())
}
func (v SeparatorToolItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSeparatorToolItem(unsafe.Pointer(ptr)), nil
	}
}
func (v SeparatorToolItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v SeparatorToolItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v SeparatorToolItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// SeparatorToolItemNew is a wrapper around gtk_separator_tool_item_new().
func SeparatorToolItemNew() ToolItem {
	ret0 := C.gtk_separator_tool_item_new()
	return wrapToolItem(ret0)
}

// GetDraw is a wrapper around gtk_separator_tool_item_get_draw().
func (item SeparatorToolItem) GetDraw() bool {
	ret0 := C.gtk_separator_tool_item_get_draw(item.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetDraw is a wrapper around gtk_separator_tool_item_set_draw().
func (item SeparatorToolItem) SetDraw(draw bool) {
	C.gtk_separator_tool_item_set_draw(item.native(), C.gboolean(util.Bool2Int(draw)) /*go:.util*/)
}

// Object Settings
type Settings struct {
	gobject.Object
}

func (v Settings) native() *C.GtkSettings {
	return (*C.GtkSettings)(v.Ptr)
}
func wrapSettings(p *C.GtkSettings) (v Settings) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSettings(p unsafe.Pointer) (v Settings) {
	v.Ptr = p
	return
}
func (v Settings) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSettings(p unsafe.Pointer) interface{} {
	return WrapSettings(p)
}
func (v Settings) GetType() gobject.Type {
	return gobject.Type(C.gtk_settings_get_type())
}
func (v Settings) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSettings(unsafe.Pointer(ptr)), nil
	}
}
func (v Settings) StyleProvider() StyleProvider {
	return WrapStyleProvider(v.Ptr)
}

// ResetProperty is a wrapper around gtk_settings_reset_property().
func (settings Settings) ResetProperty(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_settings_reset_property(settings.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SettingsGetDefault is a wrapper around gtk_settings_get_default().
func SettingsGetDefault() Settings {
	ret0 := C.gtk_settings_get_default()
	return wrapSettings(ret0)
}

// Object ShortcutLabel
type ShortcutLabel struct {
	Box
}

func (v ShortcutLabel) native() *C.GtkShortcutLabel {
	return (*C.GtkShortcutLabel)(v.Ptr)
}
func wrapShortcutLabel(p *C.GtkShortcutLabel) (v ShortcutLabel) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapShortcutLabel(p unsafe.Pointer) (v ShortcutLabel) {
	v.Ptr = p
	return
}
func (v ShortcutLabel) IsNil() bool {
	return v.Ptr == nil
}
func IWrapShortcutLabel(p unsafe.Pointer) interface{} {
	return WrapShortcutLabel(p)
}
func (v ShortcutLabel) GetType() gobject.Type {
	return gobject.Type(C.gtk_shortcut_label_get_type())
}
func (v ShortcutLabel) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapShortcutLabel(unsafe.Pointer(ptr)), nil
	}
}
func (v ShortcutLabel) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ShortcutLabel) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ShortcutLabel) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// ShortcutLabelNew is a wrapper around gtk_shortcut_label_new().
func ShortcutLabelNew(accelerator string) Widget {
	accelerator0 := (*C.gchar)(C.CString(accelerator))
	ret0 := C.gtk_shortcut_label_new(accelerator0)
	C.free(unsafe.Pointer(accelerator0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetAccelerator is a wrapper around gtk_shortcut_label_get_accelerator().
func (self ShortcutLabel) GetAccelerator() string {
	ret0 := C.gtk_shortcut_label_get_accelerator(self.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDisabledText is a wrapper around gtk_shortcut_label_get_disabled_text().
func (self ShortcutLabel) GetDisabledText() string {
	ret0 := C.gtk_shortcut_label_get_disabled_text(self.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetAccelerator is a wrapper around gtk_shortcut_label_set_accelerator().
func (self ShortcutLabel) SetAccelerator(accelerator string) {
	accelerator0 := (*C.gchar)(C.CString(accelerator))
	C.gtk_shortcut_label_set_accelerator(self.native(), accelerator0)
	C.free(unsafe.Pointer(accelerator0)) /*ch:<stdlib.h>*/
}

// SetDisabledText is a wrapper around gtk_shortcut_label_set_disabled_text().
func (self ShortcutLabel) SetDisabledText(disabled_text string) {
	disabled_text0 := (*C.gchar)(C.CString(disabled_text))
	C.gtk_shortcut_label_set_disabled_text(self.native(), disabled_text0)
	C.free(unsafe.Pointer(disabled_text0)) /*ch:<stdlib.h>*/
}

// Object ShortcutsGroup
type ShortcutsGroup struct {
	Box
}

func (v ShortcutsGroup) native() *C.GtkShortcutsGroup {
	return (*C.GtkShortcutsGroup)(v.Ptr)
}
func wrapShortcutsGroup(p *C.GtkShortcutsGroup) (v ShortcutsGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapShortcutsGroup(p unsafe.Pointer) (v ShortcutsGroup) {
	v.Ptr = p
	return
}
func (v ShortcutsGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapShortcutsGroup(p unsafe.Pointer) interface{} {
	return WrapShortcutsGroup(p)
}
func (v ShortcutsGroup) GetType() gobject.Type {
	return gobject.Type(C.gtk_shortcuts_group_get_type())
}
func (v ShortcutsGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapShortcutsGroup(unsafe.Pointer(ptr)), nil
	}
}
func (v ShortcutsGroup) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ShortcutsGroup) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ShortcutsGroup) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object ShortcutsSection
type ShortcutsSection struct {
	Box
}

func (v ShortcutsSection) native() *C.GtkShortcutsSection {
	return (*C.GtkShortcutsSection)(v.Ptr)
}
func wrapShortcutsSection(p *C.GtkShortcutsSection) (v ShortcutsSection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapShortcutsSection(p unsafe.Pointer) (v ShortcutsSection) {
	v.Ptr = p
	return
}
func (v ShortcutsSection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapShortcutsSection(p unsafe.Pointer) interface{} {
	return WrapShortcutsSection(p)
}
func (v ShortcutsSection) GetType() gobject.Type {
	return gobject.Type(C.gtk_shortcuts_section_get_type())
}
func (v ShortcutsSection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapShortcutsSection(unsafe.Pointer(ptr)), nil
	}
}
func (v ShortcutsSection) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ShortcutsSection) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ShortcutsSection) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object ShortcutsShortcut
type ShortcutsShortcut struct {
	Box
}

func (v ShortcutsShortcut) native() *C.GtkShortcutsShortcut {
	return (*C.GtkShortcutsShortcut)(v.Ptr)
}
func wrapShortcutsShortcut(p *C.GtkShortcutsShortcut) (v ShortcutsShortcut) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapShortcutsShortcut(p unsafe.Pointer) (v ShortcutsShortcut) {
	v.Ptr = p
	return
}
func (v ShortcutsShortcut) IsNil() bool {
	return v.Ptr == nil
}
func IWrapShortcutsShortcut(p unsafe.Pointer) interface{} {
	return WrapShortcutsShortcut(p)
}
func (v ShortcutsShortcut) GetType() gobject.Type {
	return gobject.Type(C.gtk_shortcuts_shortcut_get_type())
}
func (v ShortcutsShortcut) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapShortcutsShortcut(unsafe.Pointer(ptr)), nil
	}
}
func (v ShortcutsShortcut) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ShortcutsShortcut) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ShortcutsShortcut) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object Socket
type Socket struct {
	Container
}

func (v Socket) native() *C.GtkSocket {
	return (*C.GtkSocket)(v.Ptr)
}
func wrapSocket(p *C.GtkSocket) (v Socket) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocket(p unsafe.Pointer) (v Socket) {
	v.Ptr = p
	return
}
func (v Socket) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocket(p unsafe.Pointer) interface{} {
	return WrapSocket(p)
}
func (v Socket) GetType() gobject.Type {
	return gobject.Type(C.gtk_socket_get_type())
}
func (v Socket) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocket(unsafe.Pointer(ptr)), nil
	}
}
func (v Socket) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Socket) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// SocketNew is a wrapper around gtk_socket_new().
func SocketNew() Widget {
	ret0 := C.gtk_socket_new()
	return wrapWidget(ret0)
}

// GetPlugWindow is a wrapper around gtk_socket_get_plug_window().
func (socket_ Socket) GetPlugWindow() gdk.Window {
	ret0 := C.gtk_socket_get_plug_window(socket_.native())
	return gdk.WrapWindow(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// Object SpinButton
type SpinButton struct {
	Entry
}

func (v SpinButton) native() *C.GtkSpinButton {
	return (*C.GtkSpinButton)(v.Ptr)
}
func wrapSpinButton(p *C.GtkSpinButton) (v SpinButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSpinButton(p unsafe.Pointer) (v SpinButton) {
	v.Ptr = p
	return
}
func (v SpinButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSpinButton(p unsafe.Pointer) interface{} {
	return WrapSpinButton(p)
}
func (v SpinButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_spin_button_get_type())
}
func (v SpinButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSpinButton(unsafe.Pointer(ptr)), nil
	}
}
func (v SpinButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v SpinButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v SpinButton) CellEditable() CellEditable {
	return WrapCellEditable(v.Ptr)
}
func (v SpinButton) Editable() Editable {
	return WrapEditable(v.Ptr)
}
func (v SpinButton) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// SpinButtonNew is a wrapper around gtk_spin_button_new().
func SpinButtonNew(adjustment Adjustment, climb_rate float64, digits uint) Widget {
	ret0 := C.gtk_spin_button_new(adjustment.native(), C.gdouble(climb_rate), C.guint(digits))
	return wrapWidget(ret0)
}

// SpinButtonNewWithRange is a wrapper around gtk_spin_button_new_with_range().
func SpinButtonNewWithRange(min float64, max float64, step float64) Widget {
	ret0 := C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))
	return wrapWidget(ret0)
}

// Configure is a wrapper around gtk_spin_button_configure().
func (spin_button SpinButton) Configure(adjustment Adjustment, climb_rate float64, digits uint) {
	C.gtk_spin_button_configure(spin_button.native(), adjustment.native(), C.gdouble(climb_rate), C.guint(digits))
}

// GetAdjustment is a wrapper around gtk_spin_button_get_adjustment().
func (spin_button SpinButton) GetAdjustment() Adjustment {
	ret0 := C.gtk_spin_button_get_adjustment(spin_button.native())
	return wrapAdjustment(ret0)
}

// GetDigits is a wrapper around gtk_spin_button_get_digits().
func (spin_button SpinButton) GetDigits() uint {
	ret0 := C.gtk_spin_button_get_digits(spin_button.native())
	return uint(ret0)
}

// GetIncrements is a wrapper around gtk_spin_button_get_increments().
func (spin_button SpinButton) GetIncrements() (float64, float64) {
	var step0 C.gdouble
	var page0 C.gdouble
	C.gtk_spin_button_get_increments(spin_button.native(), &step0, &page0)
	return float64(step0), float64(page0)
}

// GetNumeric is a wrapper around gtk_spin_button_get_numeric().
func (spin_button SpinButton) GetNumeric() bool {
	ret0 := C.gtk_spin_button_get_numeric(spin_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRange is a wrapper around gtk_spin_button_get_range().
func (spin_button SpinButton) GetRange() (float64, float64) {
	var min0 C.gdouble
	var max0 C.gdouble
	C.gtk_spin_button_get_range(spin_button.native(), &min0, &max0)
	return float64(min0), float64(max0)
}

// GetSnapToTicks is a wrapper around gtk_spin_button_get_snap_to_ticks().
func (spin_button SpinButton) GetSnapToTicks() bool {
	ret0 := C.gtk_spin_button_get_snap_to_ticks(spin_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetUpdatePolicy is a wrapper around gtk_spin_button_get_update_policy().
func (spin_button SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {
	ret0 := C.gtk_spin_button_get_update_policy(spin_button.native())
	return SpinButtonUpdatePolicy(ret0)
}

// GetValue is a wrapper around gtk_spin_button_get_value().
func (spin_button SpinButton) GetValue() float64 {
	ret0 := C.gtk_spin_button_get_value(spin_button.native())
	return float64(ret0)
}

// GetValueAsInt is a wrapper around gtk_spin_button_get_value_as_int().
func (spin_button SpinButton) GetValueAsInt() int {
	ret0 := C.gtk_spin_button_get_value_as_int(spin_button.native())
	return int(ret0)
}

// GetWrap is a wrapper around gtk_spin_button_get_wrap().
func (spin_button SpinButton) GetWrap() bool {
	ret0 := C.gtk_spin_button_get_wrap(spin_button.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetAdjustment is a wrapper around gtk_spin_button_set_adjustment().
func (spin_button SpinButton) SetAdjustment(adjustment Adjustment) {
	C.gtk_spin_button_set_adjustment(spin_button.native(), adjustment.native())
}

// SetDigits is a wrapper around gtk_spin_button_set_digits().
func (spin_button SpinButton) SetDigits(digits uint) {
	C.gtk_spin_button_set_digits(spin_button.native(), C.guint(digits))
}

// SetIncrements is a wrapper around gtk_spin_button_set_increments().
func (spin_button SpinButton) SetIncrements(step float64, page float64) {
	C.gtk_spin_button_set_increments(spin_button.native(), C.gdouble(step), C.gdouble(page))
}

// SetNumeric is a wrapper around gtk_spin_button_set_numeric().
func (spin_button SpinButton) SetNumeric(numeric bool) {
	C.gtk_spin_button_set_numeric(spin_button.native(), C.gboolean(util.Bool2Int(numeric)) /*go:.util*/)
}

// SetRange is a wrapper around gtk_spin_button_set_range().
func (spin_button SpinButton) SetRange(min float64, max float64) {
	C.gtk_spin_button_set_range(spin_button.native(), C.gdouble(min), C.gdouble(max))
}

// SetSnapToTicks is a wrapper around gtk_spin_button_set_snap_to_ticks().
func (spin_button SpinButton) SetSnapToTicks(snap_to_ticks bool) {
	C.gtk_spin_button_set_snap_to_ticks(spin_button.native(), C.gboolean(util.Bool2Int(snap_to_ticks)) /*go:.util*/)
}

// SetUpdatePolicy is a wrapper around gtk_spin_button_set_update_policy().
func (spin_button SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	C.gtk_spin_button_set_update_policy(spin_button.native(), C.GtkSpinButtonUpdatePolicy(policy))
}

// SetValue is a wrapper around gtk_spin_button_set_value().
func (spin_button SpinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(spin_button.native(), C.gdouble(value))
}

// SetWrap is a wrapper around gtk_spin_button_set_wrap().
func (spin_button SpinButton) SetWrap(wrap bool) {
	C.gtk_spin_button_set_wrap(spin_button.native(), C.gboolean(util.Bool2Int(wrap)) /*go:.util*/)
}

// Spin is a wrapper around gtk_spin_button_spin().
func (spin_button SpinButton) Spin(direction SpinType, increment float64) {
	C.gtk_spin_button_spin(spin_button.native(), C.GtkSpinType(direction), C.gdouble(increment))
}

// Update is a wrapper around gtk_spin_button_update().
func (spin_button SpinButton) Update() {
	C.gtk_spin_button_update(spin_button.native())
}

// Object SpinButtonAccessible
type SpinButtonAccessible struct {
	EntryAccessible
}

func (v SpinButtonAccessible) native() *C.GtkSpinButtonAccessible {
	return (*C.GtkSpinButtonAccessible)(v.Ptr)
}
func wrapSpinButtonAccessible(p *C.GtkSpinButtonAccessible) (v SpinButtonAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSpinButtonAccessible(p unsafe.Pointer) (v SpinButtonAccessible) {
	v.Ptr = p
	return
}
func (v SpinButtonAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSpinButtonAccessible(p unsafe.Pointer) interface{} {
	return WrapSpinButtonAccessible(p)
}
func (v SpinButtonAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_spin_button_accessible_get_type())
}
func (v SpinButtonAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSpinButtonAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v SpinButtonAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v SpinButtonAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v SpinButtonAccessible) EditableText() atk.EditableText {
	return atk.WrapEditableText(v.Ptr) /*gir:Atk*/
}
func (v SpinButtonAccessible) Text() atk.Text {
	return atk.WrapText(v.Ptr) /*gir:Atk*/
}
func (v SpinButtonAccessible) Value() atk.Value {
	return atk.WrapValue(v.Ptr) /*gir:Atk*/
}

// Object Spinner
type Spinner struct {
	Widget
}

func (v Spinner) native() *C.GtkSpinner {
	return (*C.GtkSpinner)(v.Ptr)
}
func wrapSpinner(p *C.GtkSpinner) (v Spinner) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSpinner(p unsafe.Pointer) (v Spinner) {
	v.Ptr = p
	return
}
func (v Spinner) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSpinner(p unsafe.Pointer) interface{} {
	return WrapSpinner(p)
}
func (v Spinner) GetType() gobject.Type {
	return gobject.Type(C.gtk_spinner_get_type())
}
func (v Spinner) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSpinner(unsafe.Pointer(ptr)), nil
	}
}
func (v Spinner) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Spinner) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// SpinnerNew is a wrapper around gtk_spinner_new().
func SpinnerNew() Widget {
	ret0 := C.gtk_spinner_new()
	return wrapWidget(ret0)
}

// Start is a wrapper around gtk_spinner_start().
func (spinner Spinner) Start() {
	C.gtk_spinner_start(spinner.native())
}

// Stop is a wrapper around gtk_spinner_stop().
func (spinner Spinner) Stop() {
	C.gtk_spinner_stop(spinner.native())
}

// Object SpinnerAccessible
type SpinnerAccessible struct {
	WidgetAccessible
}

func (v SpinnerAccessible) native() *C.GtkSpinnerAccessible {
	return (*C.GtkSpinnerAccessible)(v.Ptr)
}
func wrapSpinnerAccessible(p *C.GtkSpinnerAccessible) (v SpinnerAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSpinnerAccessible(p unsafe.Pointer) (v SpinnerAccessible) {
	v.Ptr = p
	return
}
func (v SpinnerAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSpinnerAccessible(p unsafe.Pointer) interface{} {
	return WrapSpinnerAccessible(p)
}
func (v SpinnerAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_spinner_accessible_get_type())
}
func (v SpinnerAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSpinnerAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v SpinnerAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v SpinnerAccessible) Image() atk.Image {
	return atk.WrapImage(v.Ptr) /*gir:Atk*/
}

// Object Stack
type Stack struct {
	Container
}

func (v Stack) native() *C.GtkStack {
	return (*C.GtkStack)(v.Ptr)
}
func wrapStack(p *C.GtkStack) (v Stack) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStack(p unsafe.Pointer) (v Stack) {
	v.Ptr = p
	return
}
func (v Stack) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStack(p unsafe.Pointer) interface{} {
	return WrapStack(p)
}
func (v Stack) GetType() gobject.Type {
	return gobject.Type(C.gtk_stack_get_type())
}
func (v Stack) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStack(unsafe.Pointer(ptr)), nil
	}
}
func (v Stack) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Stack) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// StackNew is a wrapper around gtk_stack_new().
func StackNew() Widget {
	ret0 := C.gtk_stack_new()
	return wrapWidget(ret0)
}

// AddNamed is a wrapper around gtk_stack_add_named().
func (stack Stack) AddNamed(child Widget, name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_stack_add_named(stack.native(), child.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// AddTitled is a wrapper around gtk_stack_add_titled().
func (stack Stack) AddTitled(child Widget, name string, title string) {
	name0 := (*C.gchar)(C.CString(name))
	title0 := (*C.gchar)(C.CString(title))
	C.gtk_stack_add_titled(stack.native(), child.native(), name0, title0)
	C.free(unsafe.Pointer(name0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// GetChildByName is a wrapper around gtk_stack_get_child_by_name().
func (stack Stack) GetChildByName(name string) Widget {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_stack_get_child_by_name(stack.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetHhomogeneous is a wrapper around gtk_stack_get_hhomogeneous().
func (stack Stack) GetHhomogeneous() bool {
	ret0 := C.gtk_stack_get_hhomogeneous(stack.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHomogeneous is a wrapper around gtk_stack_get_homogeneous().
func (stack Stack) GetHomogeneous() bool {
	ret0 := C.gtk_stack_get_homogeneous(stack.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInterpolateSize is a wrapper around gtk_stack_get_interpolate_size().
func (stack Stack) GetInterpolateSize() bool {
	ret0 := C.gtk_stack_get_interpolate_size(stack.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTransitionDuration is a wrapper around gtk_stack_get_transition_duration().
func (stack Stack) GetTransitionDuration() uint {
	ret0 := C.gtk_stack_get_transition_duration(stack.native())
	return uint(ret0)
}

// GetTransitionRunning is a wrapper around gtk_stack_get_transition_running().
func (stack Stack) GetTransitionRunning() bool {
	ret0 := C.gtk_stack_get_transition_running(stack.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTransitionType is a wrapper around gtk_stack_get_transition_type().
func (stack Stack) GetTransitionType() StackTransitionType {
	ret0 := C.gtk_stack_get_transition_type(stack.native())
	return StackTransitionType(ret0)
}

// GetVhomogeneous is a wrapper around gtk_stack_get_vhomogeneous().
func (stack Stack) GetVhomogeneous() bool {
	ret0 := C.gtk_stack_get_vhomogeneous(stack.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVisibleChild is a wrapper around gtk_stack_get_visible_child().
func (stack Stack) GetVisibleChild() Widget {
	ret0 := C.gtk_stack_get_visible_child(stack.native())
	return wrapWidget(ret0)
}

// GetVisibleChildName is a wrapper around gtk_stack_get_visible_child_name().
func (stack Stack) GetVisibleChildName() string {
	ret0 := C.gtk_stack_get_visible_child_name(stack.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetHhomogeneous is a wrapper around gtk_stack_set_hhomogeneous().
func (stack Stack) SetHhomogeneous(hhomogeneous bool) {
	C.gtk_stack_set_hhomogeneous(stack.native(), C.gboolean(util.Bool2Int(hhomogeneous)) /*go:.util*/)
}

// SetHomogeneous is a wrapper around gtk_stack_set_homogeneous().
func (stack Stack) SetHomogeneous(homogeneous bool) {
	C.gtk_stack_set_homogeneous(stack.native(), C.gboolean(util.Bool2Int(homogeneous)) /*go:.util*/)
}

// SetInterpolateSize is a wrapper around gtk_stack_set_interpolate_size().
func (stack Stack) SetInterpolateSize(interpolate_size bool) {
	C.gtk_stack_set_interpolate_size(stack.native(), C.gboolean(util.Bool2Int(interpolate_size)) /*go:.util*/)
}

// SetTransitionDuration is a wrapper around gtk_stack_set_transition_duration().
func (stack Stack) SetTransitionDuration(duration uint) {
	C.gtk_stack_set_transition_duration(stack.native(), C.guint(duration))
}

// SetTransitionType is a wrapper around gtk_stack_set_transition_type().
func (stack Stack) SetTransitionType(transition StackTransitionType) {
	C.gtk_stack_set_transition_type(stack.native(), C.GtkStackTransitionType(transition))
}

// SetVhomogeneous is a wrapper around gtk_stack_set_vhomogeneous().
func (stack Stack) SetVhomogeneous(vhomogeneous bool) {
	C.gtk_stack_set_vhomogeneous(stack.native(), C.gboolean(util.Bool2Int(vhomogeneous)) /*go:.util*/)
}

// SetVisibleChild is a wrapper around gtk_stack_set_visible_child().
func (stack Stack) SetVisibleChild(child Widget) {
	C.gtk_stack_set_visible_child(stack.native(), child.native())
}

// SetVisibleChildFull is a wrapper around gtk_stack_set_visible_child_full().
func (stack Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_stack_set_visible_child_full(stack.native(), name0, C.GtkStackTransitionType(transition))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetVisibleChildName is a wrapper around gtk_stack_set_visible_child_name().
func (stack Stack) SetVisibleChildName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gtk_stack_set_visible_child_name(stack.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// Object StackSidebar
type StackSidebar struct {
	Bin
}

func (v StackSidebar) native() *C.GtkStackSidebar {
	return (*C.GtkStackSidebar)(v.Ptr)
}
func wrapStackSidebar(p *C.GtkStackSidebar) (v StackSidebar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStackSidebar(p unsafe.Pointer) (v StackSidebar) {
	v.Ptr = p
	return
}
func (v StackSidebar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStackSidebar(p unsafe.Pointer) interface{} {
	return WrapStackSidebar(p)
}
func (v StackSidebar) GetType() gobject.Type {
	return gobject.Type(C.gtk_stack_sidebar_get_type())
}
func (v StackSidebar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStackSidebar(unsafe.Pointer(ptr)), nil
	}
}
func (v StackSidebar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v StackSidebar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// StackSidebarNew is a wrapper around gtk_stack_sidebar_new().
func StackSidebarNew() Widget {
	ret0 := C.gtk_stack_sidebar_new()
	return wrapWidget(ret0)
}

// GetStack is a wrapper around gtk_stack_sidebar_get_stack().
func (sidebar StackSidebar) GetStack() Stack {
	ret0 := C.gtk_stack_sidebar_get_stack(sidebar.native())
	return wrapStack(ret0)
}

// SetStack is a wrapper around gtk_stack_sidebar_set_stack().
func (sidebar StackSidebar) SetStack(stack Stack) {
	C.gtk_stack_sidebar_set_stack(sidebar.native(), stack.native())
}

// Object StackSwitcher
type StackSwitcher struct {
	Box
}

func (v StackSwitcher) native() *C.GtkStackSwitcher {
	return (*C.GtkStackSwitcher)(v.Ptr)
}
func wrapStackSwitcher(p *C.GtkStackSwitcher) (v StackSwitcher) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStackSwitcher(p unsafe.Pointer) (v StackSwitcher) {
	v.Ptr = p
	return
}
func (v StackSwitcher) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStackSwitcher(p unsafe.Pointer) interface{} {
	return WrapStackSwitcher(p)
}
func (v StackSwitcher) GetType() gobject.Type {
	return gobject.Type(C.gtk_stack_switcher_get_type())
}
func (v StackSwitcher) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStackSwitcher(unsafe.Pointer(ptr)), nil
	}
}
func (v StackSwitcher) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v StackSwitcher) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v StackSwitcher) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// StackSwitcherNew is a wrapper around gtk_stack_switcher_new().
func StackSwitcherNew() Widget {
	ret0 := C.gtk_stack_switcher_new()
	return wrapWidget(ret0)
}

// GetStack is a wrapper around gtk_stack_switcher_get_stack().
func (switcher StackSwitcher) GetStack() Stack {
	ret0 := C.gtk_stack_switcher_get_stack(switcher.native())
	return wrapStack(ret0)
}

// SetStack is a wrapper around gtk_stack_switcher_set_stack().
func (switcher StackSwitcher) SetStack(stack Stack) {
	C.gtk_stack_switcher_set_stack(switcher.native(), stack.native())
}

// Object StatusIcon
type StatusIcon struct {
	gobject.Object
}

func (v StatusIcon) native() *C.GtkStatusIcon {
	return (*C.GtkStatusIcon)(v.Ptr)
}
func wrapStatusIcon(p *C.GtkStatusIcon) (v StatusIcon) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStatusIcon(p unsafe.Pointer) (v StatusIcon) {
	v.Ptr = p
	return
}
func (v StatusIcon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStatusIcon(p unsafe.Pointer) interface{} {
	return WrapStatusIcon(p)
}
func (v StatusIcon) GetType() gobject.Type {
	return gobject.Type(C.gtk_status_icon_get_type())
}
func (v StatusIcon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStatusIcon(unsafe.Pointer(ptr)), nil
	}
}

// Object Statusbar
type Statusbar struct {
	Box
}

func (v Statusbar) native() *C.GtkStatusbar {
	return (*C.GtkStatusbar)(v.Ptr)
}
func wrapStatusbar(p *C.GtkStatusbar) (v Statusbar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStatusbar(p unsafe.Pointer) (v Statusbar) {
	v.Ptr = p
	return
}
func (v Statusbar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStatusbar(p unsafe.Pointer) interface{} {
	return WrapStatusbar(p)
}
func (v Statusbar) GetType() gobject.Type {
	return gobject.Type(C.gtk_statusbar_get_type())
}
func (v Statusbar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStatusbar(unsafe.Pointer(ptr)), nil
	}
}
func (v Statusbar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Statusbar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Statusbar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// StatusbarNew is a wrapper around gtk_statusbar_new().
func StatusbarNew() Widget {
	ret0 := C.gtk_statusbar_new()
	return wrapWidget(ret0)
}

// GetContextId is a wrapper around gtk_statusbar_get_context_id().
func (statusbar Statusbar) GetContextId(context_description string) uint {
	context_description0 := (*C.gchar)(C.CString(context_description))
	ret0 := C.gtk_statusbar_get_context_id(statusbar.native(), context_description0)
	C.free(unsafe.Pointer(context_description0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// Pop is a wrapper around gtk_statusbar_pop().
func (statusbar Statusbar) Pop(context_id uint) {
	C.gtk_statusbar_pop(statusbar.native(), C.guint(context_id))
}

// Push is a wrapper around gtk_statusbar_push().
func (statusbar Statusbar) Push(context_id uint, text string) uint {
	text0 := (*C.gchar)(C.CString(text))
	ret0 := C.gtk_statusbar_push(statusbar.native(), C.guint(context_id), text0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// Remove is a wrapper around gtk_statusbar_remove().
func (statusbar Statusbar) Remove(context_id uint, message_id uint) {
	C.gtk_statusbar_remove(statusbar.native(), C.guint(context_id), C.guint(message_id))
}

// RemoveAll is a wrapper around gtk_statusbar_remove_all().
func (statusbar Statusbar) RemoveAll(context_id uint) {
	C.gtk_statusbar_remove_all(statusbar.native(), C.guint(context_id))
}

// Object StatusbarAccessible
type StatusbarAccessible struct {
	ContainerAccessible
}

func (v StatusbarAccessible) native() *C.GtkStatusbarAccessible {
	return (*C.GtkStatusbarAccessible)(v.Ptr)
}
func wrapStatusbarAccessible(p *C.GtkStatusbarAccessible) (v StatusbarAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStatusbarAccessible(p unsafe.Pointer) (v StatusbarAccessible) {
	v.Ptr = p
	return
}
func (v StatusbarAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStatusbarAccessible(p unsafe.Pointer) interface{} {
	return WrapStatusbarAccessible(p)
}
func (v StatusbarAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_statusbar_accessible_get_type())
}
func (v StatusbarAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStatusbarAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v StatusbarAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object Style
type Style struct {
	gobject.Object
}

func (v Style) native() *C.GtkStyle {
	return (*C.GtkStyle)(v.Ptr)
}
func wrapStyle(p *C.GtkStyle) (v Style) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStyle(p unsafe.Pointer) (v Style) {
	v.Ptr = p
	return
}
func (v Style) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStyle(p unsafe.Pointer) interface{} {
	return WrapStyle(p)
}
func (v Style) GetType() gobject.Type {
	return gobject.Type(C.gtk_style_get_type())
}
func (v Style) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStyle(unsafe.Pointer(ptr)), nil
	}
}

// HasContext is a wrapper around gtk_style_has_context().
func (style Style) HasContext() bool {
	ret0 := C.gtk_style_has_context(style.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object StyleProperties
type StyleProperties struct {
	gobject.Object
}

func (v StyleProperties) native() *C.GtkStyleProperties {
	return (*C.GtkStyleProperties)(v.Ptr)
}
func wrapStyleProperties(p *C.GtkStyleProperties) (v StyleProperties) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStyleProperties(p unsafe.Pointer) (v StyleProperties) {
	v.Ptr = p
	return
}
func (v StyleProperties) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStyleProperties(p unsafe.Pointer) interface{} {
	return WrapStyleProperties(p)
}
func (v StyleProperties) GetType() gobject.Type {
	return gobject.Type(C.gtk_style_properties_get_type())
}
func (v StyleProperties) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStyleProperties(unsafe.Pointer(ptr)), nil
	}
}
func (v StyleProperties) StyleProvider() StyleProvider {
	return WrapStyleProvider(v.Ptr)
}

// Object Switch
type Switch struct {
	Widget
}

func (v Switch) native() *C.GtkSwitch {
	return (*C.GtkSwitch)(v.Ptr)
}
func wrapSwitch(p *C.GtkSwitch) (v Switch) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSwitch(p unsafe.Pointer) (v Switch) {
	v.Ptr = p
	return
}
func (v Switch) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSwitch(p unsafe.Pointer) interface{} {
	return WrapSwitch(p)
}
func (v Switch) GetType() gobject.Type {
	return gobject.Type(C.gtk_switch_get_type())
}
func (v Switch) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSwitch(unsafe.Pointer(ptr)), nil
	}
}
func (v Switch) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Switch) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v Switch) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v Switch) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// SwitchNew is a wrapper around gtk_switch_new().
func SwitchNew() Widget {
	ret0 := C.gtk_switch_new()
	return wrapWidget(ret0)
}

// GetActive is a wrapper around gtk_switch_get_active().
func (sw Switch) GetActive() bool {
	ret0 := C.gtk_switch_get_active(sw.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetState is a wrapper around gtk_switch_get_state().
func (sw Switch) GetState() bool {
	ret0 := C.gtk_switch_get_state(sw.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetActive is a wrapper around gtk_switch_set_active().
func (sw Switch) SetActive(is_active bool) {
	C.gtk_switch_set_active(sw.native(), C.gboolean(util.Bool2Int(is_active)) /*go:.util*/)
}

// SetState is a wrapper around gtk_switch_set_state().
func (sw Switch) SetState(state bool) {
	C.gtk_switch_set_state(sw.native(), C.gboolean(util.Bool2Int(state)) /*go:.util*/)
}

// Object SwitchAccessible
type SwitchAccessible struct {
	WidgetAccessible
}

func (v SwitchAccessible) native() *C.GtkSwitchAccessible {
	return (*C.GtkSwitchAccessible)(v.Ptr)
}
func wrapSwitchAccessible(p *C.GtkSwitchAccessible) (v SwitchAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSwitchAccessible(p unsafe.Pointer) (v SwitchAccessible) {
	v.Ptr = p
	return
}
func (v SwitchAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSwitchAccessible(p unsafe.Pointer) interface{} {
	return WrapSwitchAccessible(p)
}
func (v SwitchAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_switch_accessible_get_type())
}
func (v SwitchAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSwitchAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v SwitchAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v SwitchAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}

// Object Table
type Table struct {
	Container
}

func (v Table) native() *C.GtkTable {
	return (*C.GtkTable)(v.Ptr)
}
func wrapTable(p *C.GtkTable) (v Table) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTable(p unsafe.Pointer) (v Table) {
	v.Ptr = p
	return
}
func (v Table) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTable(p unsafe.Pointer) interface{} {
	return WrapTable(p)
}
func (v Table) GetType() gobject.Type {
	return gobject.Type(C.gtk_table_get_type())
}
func (v Table) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTable(unsafe.Pointer(ptr)), nil
	}
}
func (v Table) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Table) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object TearoffMenuItem
type TearoffMenuItem struct {
	MenuItem
}

func (v TearoffMenuItem) native() *C.GtkTearoffMenuItem {
	return (*C.GtkTearoffMenuItem)(v.Ptr)
}
func wrapTearoffMenuItem(p *C.GtkTearoffMenuItem) (v TearoffMenuItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTearoffMenuItem(p unsafe.Pointer) (v TearoffMenuItem) {
	v.Ptr = p
	return
}
func (v TearoffMenuItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTearoffMenuItem(p unsafe.Pointer) interface{} {
	return WrapTearoffMenuItem(p)
}
func (v TearoffMenuItem) GetType() gobject.Type {
	return gobject.Type(C.gtk_tearoff_menu_item_get_type())
}
func (v TearoffMenuItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTearoffMenuItem(unsafe.Pointer(ptr)), nil
	}
}
func (v TearoffMenuItem) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v TearoffMenuItem) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v TearoffMenuItem) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v TearoffMenuItem) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object TextMark
type TextMark struct {
	gobject.Object
}

func (v TextMark) native() *C.GtkTextMark {
	return (*C.GtkTextMark)(v.Ptr)
}
func wrapTextMark(p *C.GtkTextMark) (v TextMark) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextMark(p unsafe.Pointer) (v TextMark) {
	v.Ptr = p
	return
}
func (v TextMark) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextMark(p unsafe.Pointer) interface{} {
	return WrapTextMark(p)
}
func (v TextMark) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_mark_get_type())
}
func (v TextMark) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextMark(unsafe.Pointer(ptr)), nil
	}
}

// TextMarkNew is a wrapper around gtk_text_mark_new().
func TextMarkNew(name string, left_gravity bool) TextMark {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gtk_text_mark_new(name0, C.gboolean(util.Bool2Int(left_gravity)) /*go:.util*/)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapTextMark(ret0)
}

// GetBuffer is a wrapper around gtk_text_mark_get_buffer().
func (mark TextMark) GetBuffer() TextBuffer {
	ret0 := C.gtk_text_mark_get_buffer(mark.native())
	return wrapTextBuffer(ret0)
}

// GetDeleted is a wrapper around gtk_text_mark_get_deleted().
func (mark TextMark) GetDeleted() bool {
	ret0 := C.gtk_text_mark_get_deleted(mark.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLeftGravity is a wrapper around gtk_text_mark_get_left_gravity().
func (mark TextMark) GetLeftGravity() bool {
	ret0 := C.gtk_text_mark_get_left_gravity(mark.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetName is a wrapper around gtk_text_mark_get_name().
func (mark TextMark) GetName() string {
	ret0 := C.gtk_text_mark_get_name(mark.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetVisible is a wrapper around gtk_text_mark_get_visible().
func (mark TextMark) GetVisible() bool {
	ret0 := C.gtk_text_mark_get_visible(mark.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetVisible is a wrapper around gtk_text_mark_set_visible().
func (mark TextMark) SetVisible(setting bool) {
	C.gtk_text_mark_set_visible(mark.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// Object TextCellAccessible
type TextCellAccessible struct {
	RendererCellAccessible
}

func (v TextCellAccessible) native() *C.GtkTextCellAccessible {
	return (*C.GtkTextCellAccessible)(v.Ptr)
}
func wrapTextCellAccessible(p *C.GtkTextCellAccessible) (v TextCellAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextCellAccessible(p unsafe.Pointer) (v TextCellAccessible) {
	v.Ptr = p
	return
}
func (v TextCellAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextCellAccessible(p unsafe.Pointer) interface{} {
	return WrapTextCellAccessible(p)
}
func (v TextCellAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_cell_accessible_get_type())
}
func (v TextCellAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextCellAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v TextCellAccessible) Action() atk.Action {
	return atk.WrapAction(v.Ptr) /*gir:Atk*/
}
func (v TextCellAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v TextCellAccessible) Text() atk.Text {
	return atk.WrapText(v.Ptr) /*gir:Atk*/
}

// Object TextView
type TextView struct {
	Container
}

func (v TextView) native() *C.GtkTextView {
	return (*C.GtkTextView)(v.Ptr)
}
func wrapTextView(p *C.GtkTextView) (v TextView) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextView(p unsafe.Pointer) (v TextView) {
	v.Ptr = p
	return
}
func (v TextView) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextView(p unsafe.Pointer) interface{} {
	return WrapTextView(p)
}
func (v TextView) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_view_get_type())
}
func (v TextView) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextView(unsafe.Pointer(ptr)), nil
	}
}
func (v TextView) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v TextView) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v TextView) Scrollable() Scrollable {
	return WrapScrollable(v.Ptr)
}

// TextViewNew is a wrapper around gtk_text_view_new().
func TextViewNew() Widget {
	ret0 := C.gtk_text_view_new()
	return wrapWidget(ret0)
}

// TextViewNewWithBuffer is a wrapper around gtk_text_view_new_with_buffer().
func TextViewNewWithBuffer(buffer TextBuffer) Widget {
	ret0 := C.gtk_text_view_new_with_buffer(buffer.native())
	return wrapWidget(ret0)
}

// AddChildAtAnchor is a wrapper around gtk_text_view_add_child_at_anchor().
func (text_view TextView) AddChildAtAnchor(child Widget, anchor TextChildAnchor) {
	C.gtk_text_view_add_child_at_anchor(text_view.native(), child.native(), anchor.native())
}

// AddChildInWindow is a wrapper around gtk_text_view_add_child_in_window().
func (text_view TextView) AddChildInWindow(child Widget, which_window TextWindowType, xpos int, ypos int) {
	C.gtk_text_view_add_child_in_window(text_view.native(), child.native(), C.GtkTextWindowType(which_window), C.gint(xpos), C.gint(ypos))
}

// BackwardDisplayLine is a wrapper around gtk_text_view_backward_display_line().
func (text_view TextView) BackwardDisplayLine(iter TextIter) bool {
	ret0 := C.gtk_text_view_backward_display_line(text_view.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BackwardDisplayLineStart is a wrapper around gtk_text_view_backward_display_line_start().
func (text_view TextView) BackwardDisplayLineStart(iter TextIter) bool {
	ret0 := C.gtk_text_view_backward_display_line_start(text_view.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// BufferToWindowCoords is a wrapper around gtk_text_view_buffer_to_window_coords().
func (text_view TextView) BufferToWindowCoords(win TextWindowType, buffer_x int, buffer_y int) (int, int) {
	var window_x0 C.gint
	var window_y0 C.gint
	C.gtk_text_view_buffer_to_window_coords(text_view.native(), C.GtkTextWindowType(win), C.gint(buffer_x), C.gint(buffer_y), &window_x0, &window_y0)
	return int(window_x0), int(window_y0)
}

// ForwardDisplayLine is a wrapper around gtk_text_view_forward_display_line().
func (text_view TextView) ForwardDisplayLine(iter TextIter) bool {
	ret0 := C.gtk_text_view_forward_display_line(text_view.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ForwardDisplayLineEnd is a wrapper around gtk_text_view_forward_display_line_end().
func (text_view TextView) ForwardDisplayLineEnd(iter TextIter) bool {
	ret0 := C.gtk_text_view_forward_display_line_end(text_view.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetAcceptsTab is a wrapper around gtk_text_view_get_accepts_tab().
func (text_view TextView) GetAcceptsTab() bool {
	ret0 := C.gtk_text_view_get_accepts_tab(text_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetBorderWindowSize is a wrapper around gtk_text_view_get_border_window_size().
func (text_view TextView) GetBorderWindowSize(type_ TextWindowType) int {
	ret0 := C.gtk_text_view_get_border_window_size(text_view.native(), C.GtkTextWindowType(type_))
	return int(ret0)
}

// GetBottomMargin is a wrapper around gtk_text_view_get_bottom_margin().
func (text_view TextView) GetBottomMargin() int {
	ret0 := C.gtk_text_view_get_bottom_margin(text_view.native())
	return int(ret0)
}

// GetBuffer is a wrapper around gtk_text_view_get_buffer().
func (text_view TextView) GetBuffer() TextBuffer {
	ret0 := C.gtk_text_view_get_buffer(text_view.native())
	return wrapTextBuffer(ret0)
}

// GetCursorVisible is a wrapper around gtk_text_view_get_cursor_visible().
func (text_view TextView) GetCursorVisible() bool {
	ret0 := C.gtk_text_view_get_cursor_visible(text_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDefaultAttributes is a wrapper around gtk_text_view_get_default_attributes().
func (text_view TextView) GetDefaultAttributes() TextAttributes {
	ret0 := C.gtk_text_view_get_default_attributes(text_view.native())
	return wrapTextAttributes(ret0)
}

// GetEditable is a wrapper around gtk_text_view_get_editable().
func (text_view TextView) GetEditable() bool {
	ret0 := C.gtk_text_view_get_editable(text_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIndent is a wrapper around gtk_text_view_get_indent().
func (text_view TextView) GetIndent() int {
	ret0 := C.gtk_text_view_get_indent(text_view.native())
	return int(ret0)
}

// GetInputHints is a wrapper around gtk_text_view_get_input_hints().
func (text_view TextView) GetInputHints() InputHints {
	ret0 := C.gtk_text_view_get_input_hints(text_view.native())
	return InputHints(ret0)
}

// GetInputPurpose is a wrapper around gtk_text_view_get_input_purpose().
func (text_view TextView) GetInputPurpose() InputPurpose {
	ret0 := C.gtk_text_view_get_input_purpose(text_view.native())
	return InputPurpose(ret0)
}

// GetJustification is a wrapper around gtk_text_view_get_justification().
func (text_view TextView) GetJustification() Justification {
	ret0 := C.gtk_text_view_get_justification(text_view.native())
	return Justification(ret0)
}

// GetLeftMargin is a wrapper around gtk_text_view_get_left_margin().
func (text_view TextView) GetLeftMargin() int {
	ret0 := C.gtk_text_view_get_left_margin(text_view.native())
	return int(ret0)
}

// GetLineYrange is a wrapper around gtk_text_view_get_line_yrange().
func (text_view TextView) GetLineYrange(iter TextIter) (int, int) {
	var y0 C.gint
	var height0 C.gint
	C.gtk_text_view_get_line_yrange(text_view.native(), iter.native(), &y0, &height0)
	return int(y0), int(height0)
}

// GetMonospace is a wrapper around gtk_text_view_get_monospace().
func (text_view TextView) GetMonospace() bool {
	ret0 := C.gtk_text_view_get_monospace(text_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetOverwrite is a wrapper around gtk_text_view_get_overwrite().
func (text_view TextView) GetOverwrite() bool {
	ret0 := C.gtk_text_view_get_overwrite(text_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPixelsAboveLines is a wrapper around gtk_text_view_get_pixels_above_lines().
func (text_view TextView) GetPixelsAboveLines() int {
	ret0 := C.gtk_text_view_get_pixels_above_lines(text_view.native())
	return int(ret0)
}

// GetPixelsBelowLines is a wrapper around gtk_text_view_get_pixels_below_lines().
func (text_view TextView) GetPixelsBelowLines() int {
	ret0 := C.gtk_text_view_get_pixels_below_lines(text_view.native())
	return int(ret0)
}

// GetPixelsInsideWrap is a wrapper around gtk_text_view_get_pixels_inside_wrap().
func (text_view TextView) GetPixelsInsideWrap() int {
	ret0 := C.gtk_text_view_get_pixels_inside_wrap(text_view.native())
	return int(ret0)
}

// GetRightMargin is a wrapper around gtk_text_view_get_right_margin().
func (text_view TextView) GetRightMargin() int {
	ret0 := C.gtk_text_view_get_right_margin(text_view.native())
	return int(ret0)
}

// GetTabs is a wrapper around gtk_text_view_get_tabs().
func (text_view TextView) GetTabs() pango.TabArray {
	ret0 := C.gtk_text_view_get_tabs(text_view.native())
	return pango.WrapTabArray(unsafe.Pointer(ret0)) /*gir:Pango*/
}

// GetTopMargin is a wrapper around gtk_text_view_get_top_margin().
func (text_view TextView) GetTopMargin() int {
	ret0 := C.gtk_text_view_get_top_margin(text_view.native())
	return int(ret0)
}

// GetWrapMode is a wrapper around gtk_text_view_get_wrap_mode().
func (text_view TextView) GetWrapMode() WrapMode {
	ret0 := C.gtk_text_view_get_wrap_mode(text_view.native())
	return WrapMode(ret0)
}

// MoveChild is a wrapper around gtk_text_view_move_child().
func (text_view TextView) MoveChild(child Widget, xpos int, ypos int) {
	C.gtk_text_view_move_child(text_view.native(), child.native(), C.gint(xpos), C.gint(ypos))
}

// MoveMarkOnscreen is a wrapper around gtk_text_view_move_mark_onscreen().
func (text_view TextView) MoveMarkOnscreen(mark TextMark) bool {
	ret0 := C.gtk_text_view_move_mark_onscreen(text_view.native(), mark.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MoveVisually is a wrapper around gtk_text_view_move_visually().
func (text_view TextView) MoveVisually(iter TextIter, count int) bool {
	ret0 := C.gtk_text_view_move_visually(text_view.native(), iter.native(), C.gint(count))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PlaceCursorOnscreen is a wrapper around gtk_text_view_place_cursor_onscreen().
func (text_view TextView) PlaceCursorOnscreen() bool {
	ret0 := C.gtk_text_view_place_cursor_onscreen(text_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ResetCursorBlink is a wrapper around gtk_text_view_reset_cursor_blink().
func (text_view TextView) ResetCursorBlink() {
	C.gtk_text_view_reset_cursor_blink(text_view.native())
}

// ResetImContext is a wrapper around gtk_text_view_reset_im_context().
func (text_view TextView) ResetImContext() {
	C.gtk_text_view_reset_im_context(text_view.native())
}

// ScrollMarkOnscreen is a wrapper around gtk_text_view_scroll_mark_onscreen().
func (text_view TextView) ScrollMarkOnscreen(mark TextMark) {
	C.gtk_text_view_scroll_mark_onscreen(text_view.native(), mark.native())
}

// ScrollToIter is a wrapper around gtk_text_view_scroll_to_iter().
func (text_view TextView) ScrollToIter(iter TextIter, within_margin float64, use_align bool, xalign float64, yalign float64) bool {
	ret0 := C.gtk_text_view_scroll_to_iter(text_view.native(), iter.native(), C.gdouble(within_margin), C.gboolean(util.Bool2Int(use_align)) /*go:.util*/, C.gdouble(xalign), C.gdouble(yalign))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ScrollToMark is a wrapper around gtk_text_view_scroll_to_mark().
func (text_view TextView) ScrollToMark(mark TextMark, within_margin float64, use_align bool, xalign float64, yalign float64) {
	C.gtk_text_view_scroll_to_mark(text_view.native(), mark.native(), C.gdouble(within_margin), C.gboolean(util.Bool2Int(use_align)) /*go:.util*/, C.gdouble(xalign), C.gdouble(yalign))
}

// SetAcceptsTab is a wrapper around gtk_text_view_set_accepts_tab().
func (text_view TextView) SetAcceptsTab(accepts_tab bool) {
	C.gtk_text_view_set_accepts_tab(text_view.native(), C.gboolean(util.Bool2Int(accepts_tab)) /*go:.util*/)
}

// SetBorderWindowSize is a wrapper around gtk_text_view_set_border_window_size().
func (text_view TextView) SetBorderWindowSize(type_ TextWindowType, size int) {
	C.gtk_text_view_set_border_window_size(text_view.native(), C.GtkTextWindowType(type_), C.gint(size))
}

// SetBottomMargin is a wrapper around gtk_text_view_set_bottom_margin().
func (text_view TextView) SetBottomMargin(bottom_margin int) {
	C.gtk_text_view_set_bottom_margin(text_view.native(), C.gint(bottom_margin))
}

// SetBuffer is a wrapper around gtk_text_view_set_buffer().
func (text_view TextView) SetBuffer(buffer TextBuffer) {
	C.gtk_text_view_set_buffer(text_view.native(), buffer.native())
}

// SetCursorVisible is a wrapper around gtk_text_view_set_cursor_visible().
func (text_view TextView) SetCursorVisible(setting bool) {
	C.gtk_text_view_set_cursor_visible(text_view.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetEditable is a wrapper around gtk_text_view_set_editable().
func (text_view TextView) SetEditable(setting bool) {
	C.gtk_text_view_set_editable(text_view.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetIndent is a wrapper around gtk_text_view_set_indent().
func (text_view TextView) SetIndent(indent int) {
	C.gtk_text_view_set_indent(text_view.native(), C.gint(indent))
}

// SetInputHints is a wrapper around gtk_text_view_set_input_hints().
func (text_view TextView) SetInputHints(hints InputHints) {
	C.gtk_text_view_set_input_hints(text_view.native(), C.GtkInputHints(hints))
}

// SetInputPurpose is a wrapper around gtk_text_view_set_input_purpose().
func (text_view TextView) SetInputPurpose(purpose InputPurpose) {
	C.gtk_text_view_set_input_purpose(text_view.native(), C.GtkInputPurpose(purpose))
}

// SetJustification is a wrapper around gtk_text_view_set_justification().
func (text_view TextView) SetJustification(justification Justification) {
	C.gtk_text_view_set_justification(text_view.native(), C.GtkJustification(justification))
}

// SetLeftMargin is a wrapper around gtk_text_view_set_left_margin().
func (text_view TextView) SetLeftMargin(left_margin int) {
	C.gtk_text_view_set_left_margin(text_view.native(), C.gint(left_margin))
}

// SetMonospace is a wrapper around gtk_text_view_set_monospace().
func (text_view TextView) SetMonospace(monospace bool) {
	C.gtk_text_view_set_monospace(text_view.native(), C.gboolean(util.Bool2Int(monospace)) /*go:.util*/)
}

// SetOverwrite is a wrapper around gtk_text_view_set_overwrite().
func (text_view TextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(text_view.native(), C.gboolean(util.Bool2Int(overwrite)) /*go:.util*/)
}

// SetPixelsAboveLines is a wrapper around gtk_text_view_set_pixels_above_lines().
func (text_view TextView) SetPixelsAboveLines(pixels_above_lines int) {
	C.gtk_text_view_set_pixels_above_lines(text_view.native(), C.gint(pixels_above_lines))
}

// SetPixelsBelowLines is a wrapper around gtk_text_view_set_pixels_below_lines().
func (text_view TextView) SetPixelsBelowLines(pixels_below_lines int) {
	C.gtk_text_view_set_pixels_below_lines(text_view.native(), C.gint(pixels_below_lines))
}

// SetPixelsInsideWrap is a wrapper around gtk_text_view_set_pixels_inside_wrap().
func (text_view TextView) SetPixelsInsideWrap(pixels_inside_wrap int) {
	C.gtk_text_view_set_pixels_inside_wrap(text_view.native(), C.gint(pixels_inside_wrap))
}

// SetRightMargin is a wrapper around gtk_text_view_set_right_margin().
func (text_view TextView) SetRightMargin(right_margin int) {
	C.gtk_text_view_set_right_margin(text_view.native(), C.gint(right_margin))
}

// SetTopMargin is a wrapper around gtk_text_view_set_top_margin().
func (text_view TextView) SetTopMargin(top_margin int) {
	C.gtk_text_view_set_top_margin(text_view.native(), C.gint(top_margin))
}

// SetWrapMode is a wrapper around gtk_text_view_set_wrap_mode().
func (text_view TextView) SetWrapMode(wrap_mode WrapMode) {
	C.gtk_text_view_set_wrap_mode(text_view.native(), C.GtkWrapMode(wrap_mode))
}

// StartsDisplayLine is a wrapper around gtk_text_view_starts_display_line().
func (text_view TextView) StartsDisplayLine(iter TextIter) bool {
	ret0 := C.gtk_text_view_starts_display_line(text_view.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// WindowToBufferCoords is a wrapper around gtk_text_view_window_to_buffer_coords().
func (text_view TextView) WindowToBufferCoords(win TextWindowType, window_x int, window_y int) (int, int) {
	var buffer_x0 C.gint
	var buffer_y0 C.gint
	C.gtk_text_view_window_to_buffer_coords(text_view.native(), C.GtkTextWindowType(win), C.gint(window_x), C.gint(window_y), &buffer_x0, &buffer_y0)
	return int(buffer_x0), int(buffer_y0)
}

// Object TextViewAccessible
type TextViewAccessible struct {
	ContainerAccessible
}

func (v TextViewAccessible) native() *C.GtkTextViewAccessible {
	return (*C.GtkTextViewAccessible)(v.Ptr)
}
func wrapTextViewAccessible(p *C.GtkTextViewAccessible) (v TextViewAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTextViewAccessible(p unsafe.Pointer) (v TextViewAccessible) {
	v.Ptr = p
	return
}
func (v TextViewAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextViewAccessible(p unsafe.Pointer) interface{} {
	return WrapTextViewAccessible(p)
}
func (v TextViewAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_text_view_accessible_get_type())
}
func (v TextViewAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTextViewAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v TextViewAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v TextViewAccessible) EditableText() atk.EditableText {
	return atk.WrapEditableText(v.Ptr) /*gir:Atk*/
}
func (v TextViewAccessible) StreamableContent() atk.StreamableContent {
	return atk.WrapStreamableContent(v.Ptr) /*gir:Atk*/
}
func (v TextViewAccessible) Text() atk.Text {
	return atk.WrapText(v.Ptr) /*gir:Atk*/
}

// Object ThemingEngine
type ThemingEngine struct {
	gobject.Object
}

func (v ThemingEngine) native() *C.GtkThemingEngine {
	return (*C.GtkThemingEngine)(v.Ptr)
}
func wrapThemingEngine(p *C.GtkThemingEngine) (v ThemingEngine) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapThemingEngine(p unsafe.Pointer) (v ThemingEngine) {
	v.Ptr = p
	return
}
func (v ThemingEngine) IsNil() bool {
	return v.Ptr == nil
}
func IWrapThemingEngine(p unsafe.Pointer) interface{} {
	return WrapThemingEngine(p)
}
func (v ThemingEngine) GetType() gobject.Type {
	return gobject.Type(C.gtk_theming_engine_get_type())
}
func (v ThemingEngine) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapThemingEngine(unsafe.Pointer(ptr)), nil
	}
}

// Object ToolItemGroup
type ToolItemGroup struct {
	Container
}

func (v ToolItemGroup) native() *C.GtkToolItemGroup {
	return (*C.GtkToolItemGroup)(v.Ptr)
}
func wrapToolItemGroup(p *C.GtkToolItemGroup) (v ToolItemGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToolItemGroup(p unsafe.Pointer) (v ToolItemGroup) {
	v.Ptr = p
	return
}
func (v ToolItemGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToolItemGroup(p unsafe.Pointer) interface{} {
	return WrapToolItemGroup(p)
}
func (v ToolItemGroup) GetType() gobject.Type {
	return gobject.Type(C.gtk_tool_item_group_get_type())
}
func (v ToolItemGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToolItemGroup(unsafe.Pointer(ptr)), nil
	}
}
func (v ToolItemGroup) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ToolItemGroup) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ToolItemGroup) ToolShell() ToolShell {
	return WrapToolShell(v.Ptr)
}

// ToolItemGroupNew is a wrapper around gtk_tool_item_group_new().
func ToolItemGroupNew(label string) Widget {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.gtk_tool_item_group_new(label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapWidget(ret0)
}

// GetCollapsed is a wrapper around gtk_tool_item_group_get_collapsed().
func (group ToolItemGroup) GetCollapsed() bool {
	ret0 := C.gtk_tool_item_group_get_collapsed(group.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDropItem is a wrapper around gtk_tool_item_group_get_drop_item().
func (group ToolItemGroup) GetDropItem(x int, y int) ToolItem {
	ret0 := C.gtk_tool_item_group_get_drop_item(group.native(), C.gint(x), C.gint(y))
	return wrapToolItem(ret0)
}

// GetEllipsize is a wrapper around gtk_tool_item_group_get_ellipsize().
func (group ToolItemGroup) GetEllipsize() /*gir:Pango*/ pango.EllipsizeMode {
	ret0 := C.gtk_tool_item_group_get_ellipsize(group.native())
	return /*gir:Pango*/ pango.EllipsizeMode(ret0)
}

// GetHeaderRelief is a wrapper around gtk_tool_item_group_get_header_relief().
func (group ToolItemGroup) GetHeaderRelief() ReliefStyle {
	ret0 := C.gtk_tool_item_group_get_header_relief(group.native())
	return ReliefStyle(ret0)
}

// GetItemPosition is a wrapper around gtk_tool_item_group_get_item_position().
func (group ToolItemGroup) GetItemPosition(item ToolItem) int {
	ret0 := C.gtk_tool_item_group_get_item_position(group.native(), item.native())
	return int(ret0)
}

// GetLabel is a wrapper around gtk_tool_item_group_get_label().
func (group ToolItemGroup) GetLabel() string {
	ret0 := C.gtk_tool_item_group_get_label(group.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLabelWidget is a wrapper around gtk_tool_item_group_get_label_widget().
func (group ToolItemGroup) GetLabelWidget() Widget {
	ret0 := C.gtk_tool_item_group_get_label_widget(group.native())
	return wrapWidget(ret0)
}

// GetNItems is a wrapper around gtk_tool_item_group_get_n_items().
func (group ToolItemGroup) GetNItems() uint {
	ret0 := C.gtk_tool_item_group_get_n_items(group.native())
	return uint(ret0)
}

// GetNthItem is a wrapper around gtk_tool_item_group_get_nth_item().
func (group ToolItemGroup) GetNthItem(index uint) ToolItem {
	ret0 := C.gtk_tool_item_group_get_nth_item(group.native(), C.guint(index))
	return wrapToolItem(ret0)
}

// Insert is a wrapper around gtk_tool_item_group_insert().
func (group ToolItemGroup) Insert(item ToolItem, position int) {
	C.gtk_tool_item_group_insert(group.native(), item.native(), C.gint(position))
}

// SetCollapsed is a wrapper around gtk_tool_item_group_set_collapsed().
func (group ToolItemGroup) SetCollapsed(collapsed bool) {
	C.gtk_tool_item_group_set_collapsed(group.native(), C.gboolean(util.Bool2Int(collapsed)) /*go:.util*/)
}

// SetEllipsize is a wrapper around gtk_tool_item_group_set_ellipsize().
func (group ToolItemGroup) SetEllipsize(ellipsize /*gir:Pango*/ pango.EllipsizeMode) {
	C.gtk_tool_item_group_set_ellipsize(group.native(), C.PangoEllipsizeMode(ellipsize))
}

// SetHeaderRelief is a wrapper around gtk_tool_item_group_set_header_relief().
func (group ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	C.gtk_tool_item_group_set_header_relief(group.native(), C.GtkReliefStyle(style))
}

// SetItemPosition is a wrapper around gtk_tool_item_group_set_item_position().
func (group ToolItemGroup) SetItemPosition(item ToolItem, position int) {
	C.gtk_tool_item_group_set_item_position(group.native(), item.native(), C.gint(position))
}

// SetLabel is a wrapper around gtk_tool_item_group_set_label().
func (group ToolItemGroup) SetLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.gtk_tool_item_group_set_label(group.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetLabelWidget is a wrapper around gtk_tool_item_group_set_label_widget().
func (group ToolItemGroup) SetLabelWidget(label_widget Widget) {
	C.gtk_tool_item_group_set_label_widget(group.native(), label_widget.native())
}

// Object ToolPalette
type ToolPalette struct {
	Container
}

func (v ToolPalette) native() *C.GtkToolPalette {
	return (*C.GtkToolPalette)(v.Ptr)
}
func wrapToolPalette(p *C.GtkToolPalette) (v ToolPalette) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToolPalette(p unsafe.Pointer) (v ToolPalette) {
	v.Ptr = p
	return
}
func (v ToolPalette) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToolPalette(p unsafe.Pointer) interface{} {
	return WrapToolPalette(p)
}
func (v ToolPalette) GetType() gobject.Type {
	return gobject.Type(C.gtk_tool_palette_get_type())
}
func (v ToolPalette) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToolPalette(unsafe.Pointer(ptr)), nil
	}
}
func (v ToolPalette) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v ToolPalette) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v ToolPalette) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}
func (v ToolPalette) Scrollable() Scrollable {
	return WrapScrollable(v.Ptr)
}

// ToolPaletteNew is a wrapper around gtk_tool_palette_new().
func ToolPaletteNew() Widget {
	ret0 := C.gtk_tool_palette_new()
	return wrapWidget(ret0)
}

// AddDragDest is a wrapper around gtk_tool_palette_add_drag_dest().
func (palette ToolPalette) AddDragDest(widget Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions /*gir:Gdk*/ gdk.DragAction) {
	C.gtk_tool_palette_add_drag_dest(palette.native(), widget.native(), C.GtkDestDefaults(flags), C.GtkToolPaletteDragTargets(targets), C.GdkDragAction(actions))
}

// GetDragItem is a wrapper around gtk_tool_palette_get_drag_item().
func (palette ToolPalette) GetDragItem(selection SelectionData) Widget {
	ret0 := C.gtk_tool_palette_get_drag_item(palette.native(), selection.native())
	return wrapWidget(ret0)
}

// GetDropGroup is a wrapper around gtk_tool_palette_get_drop_group().
func (palette ToolPalette) GetDropGroup(x int, y int) ToolItemGroup {
	ret0 := C.gtk_tool_palette_get_drop_group(palette.native(), C.gint(x), C.gint(y))
	return wrapToolItemGroup(ret0)
}

// GetDropItem is a wrapper around gtk_tool_palette_get_drop_item().
func (palette ToolPalette) GetDropItem(x int, y int) ToolItem {
	ret0 := C.gtk_tool_palette_get_drop_item(palette.native(), C.gint(x), C.gint(y))
	return wrapToolItem(ret0)
}

// GetExclusive is a wrapper around gtk_tool_palette_get_exclusive().
func (palette ToolPalette) GetExclusive(group ToolItemGroup) bool {
	ret0 := C.gtk_tool_palette_get_exclusive(palette.native(), group.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetExpand is a wrapper around gtk_tool_palette_get_expand().
func (palette ToolPalette) GetExpand(group ToolItemGroup) bool {
	ret0 := C.gtk_tool_palette_get_expand(palette.native(), group.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetGroupPosition is a wrapper around gtk_tool_palette_get_group_position().
func (palette ToolPalette) GetGroupPosition(group ToolItemGroup) int {
	ret0 := C.gtk_tool_palette_get_group_position(palette.native(), group.native())
	return int(ret0)
}

// GetStyle is a wrapper around gtk_tool_palette_get_style().
func (palette ToolPalette) GetStyle() ToolbarStyle {
	ret0 := C.gtk_tool_palette_get_style(palette.native())
	return ToolbarStyle(ret0)
}

// SetDragSource is a wrapper around gtk_tool_palette_set_drag_source().
func (palette ToolPalette) SetDragSource(targets ToolPaletteDragTargets) {
	C.gtk_tool_palette_set_drag_source(palette.native(), C.GtkToolPaletteDragTargets(targets))
}

// SetExclusive is a wrapper around gtk_tool_palette_set_exclusive().
func (palette ToolPalette) SetExclusive(group ToolItemGroup, exclusive bool) {
	C.gtk_tool_palette_set_exclusive(palette.native(), group.native(), C.gboolean(util.Bool2Int(exclusive)) /*go:.util*/)
}

// SetExpand is a wrapper around gtk_tool_palette_set_expand().
func (palette ToolPalette) SetExpand(group ToolItemGroup, expand bool) {
	C.gtk_tool_palette_set_expand(palette.native(), group.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// SetGroupPosition is a wrapper around gtk_tool_palette_set_group_position().
func (palette ToolPalette) SetGroupPosition(group ToolItemGroup, position int) {
	C.gtk_tool_palette_set_group_position(palette.native(), group.native(), C.gint(position))
}

// SetStyle is a wrapper around gtk_tool_palette_set_style().
func (palette ToolPalette) SetStyle(style ToolbarStyle) {
	C.gtk_tool_palette_set_style(palette.native(), C.GtkToolbarStyle(style))
}

// UnsetIconSize is a wrapper around gtk_tool_palette_unset_icon_size().
func (palette ToolPalette) UnsetIconSize() {
	C.gtk_tool_palette_unset_icon_size(palette.native())
}

// UnsetStyle is a wrapper around gtk_tool_palette_unset_style().
func (palette ToolPalette) UnsetStyle() {
	C.gtk_tool_palette_unset_style(palette.native())
}

// ToolPaletteGetDragTargetGroup is a wrapper around gtk_tool_palette_get_drag_target_group().
func ToolPaletteGetDragTargetGroup() TargetEntry {
	ret0 := C.gtk_tool_palette_get_drag_target_group()
	return wrapTargetEntry(ret0)
}

// ToolPaletteGetDragTargetItem is a wrapper around gtk_tool_palette_get_drag_target_item().
func ToolPaletteGetDragTargetItem() TargetEntry {
	ret0 := C.gtk_tool_palette_get_drag_target_item()
	return wrapTargetEntry(ret0)
}

// Object Toolbar
type Toolbar struct {
	Container
}

func (v Toolbar) native() *C.GtkToolbar {
	return (*C.GtkToolbar)(v.Ptr)
}
func wrapToolbar(p *C.GtkToolbar) (v Toolbar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToolbar(p unsafe.Pointer) (v Toolbar) {
	v.Ptr = p
	return
}
func (v Toolbar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToolbar(p unsafe.Pointer) interface{} {
	return WrapToolbar(p)
}
func (v Toolbar) GetType() gobject.Type {
	return gobject.Type(C.gtk_toolbar_get_type())
}
func (v Toolbar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToolbar(unsafe.Pointer(ptr)), nil
	}
}
func (v Toolbar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Toolbar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Toolbar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}
func (v Toolbar) ToolShell() ToolShell {
	return WrapToolShell(v.Ptr)
}

// ToolbarNew is a wrapper around gtk_toolbar_new().
func ToolbarNew() Widget {
	ret0 := C.gtk_toolbar_new()
	return wrapWidget(ret0)
}

// GetDropIndex is a wrapper around gtk_toolbar_get_drop_index().
func (toolbar Toolbar) GetDropIndex(x int, y int) int {
	ret0 := C.gtk_toolbar_get_drop_index(toolbar.native(), C.gint(x), C.gint(y))
	return int(ret0)
}

// GetIconSize is a wrapper around gtk_toolbar_get_icon_size().
func (toolbar Toolbar) GetIconSize() IconSize {
	ret0 := C.gtk_toolbar_get_icon_size(toolbar.native())
	return IconSize(ret0)
}

// GetItemIndex is a wrapper around gtk_toolbar_get_item_index().
func (toolbar Toolbar) GetItemIndex(item ToolItem) int {
	ret0 := C.gtk_toolbar_get_item_index(toolbar.native(), item.native())
	return int(ret0)
}

// GetNItems is a wrapper around gtk_toolbar_get_n_items().
func (toolbar Toolbar) GetNItems() int {
	ret0 := C.gtk_toolbar_get_n_items(toolbar.native())
	return int(ret0)
}

// GetNthItem is a wrapper around gtk_toolbar_get_nth_item().
func (toolbar Toolbar) GetNthItem(n int) ToolItem {
	ret0 := C.gtk_toolbar_get_nth_item(toolbar.native(), C.gint(n))
	return wrapToolItem(ret0)
}

// GetReliefStyle is a wrapper around gtk_toolbar_get_relief_style().
func (toolbar Toolbar) GetReliefStyle() ReliefStyle {
	ret0 := C.gtk_toolbar_get_relief_style(toolbar.native())
	return ReliefStyle(ret0)
}

// GetShowArrow is a wrapper around gtk_toolbar_get_show_arrow().
func (toolbar Toolbar) GetShowArrow() bool {
	ret0 := C.gtk_toolbar_get_show_arrow(toolbar.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetStyle is a wrapper around gtk_toolbar_get_style().
func (toolbar Toolbar) GetStyle() ToolbarStyle {
	ret0 := C.gtk_toolbar_get_style(toolbar.native())
	return ToolbarStyle(ret0)
}

// Insert is a wrapper around gtk_toolbar_insert().
func (toolbar Toolbar) Insert(item ToolItem, pos int) {
	C.gtk_toolbar_insert(toolbar.native(), item.native(), C.gint(pos))
}

// SetDropHighlightItem is a wrapper around gtk_toolbar_set_drop_highlight_item().
func (toolbar Toolbar) SetDropHighlightItem(tool_item ToolItem, index_ int) {
	C.gtk_toolbar_set_drop_highlight_item(toolbar.native(), tool_item.native(), C.gint(index_))
}

// SetIconSize is a wrapper around gtk_toolbar_set_icon_size().
func (toolbar Toolbar) SetIconSize(icon_size IconSize) {
	C.gtk_toolbar_set_icon_size(toolbar.native(), C.GtkIconSize(icon_size))
}

// SetShowArrow is a wrapper around gtk_toolbar_set_show_arrow().
func (toolbar Toolbar) SetShowArrow(show_arrow bool) {
	C.gtk_toolbar_set_show_arrow(toolbar.native(), C.gboolean(util.Bool2Int(show_arrow)) /*go:.util*/)
}

// SetStyle is a wrapper around gtk_toolbar_set_style().
func (toolbar Toolbar) SetStyle(style ToolbarStyle) {
	C.gtk_toolbar_set_style(toolbar.native(), C.GtkToolbarStyle(style))
}

// UnsetIconSize is a wrapper around gtk_toolbar_unset_icon_size().
func (toolbar Toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(toolbar.native())
}

// UnsetStyle is a wrapper around gtk_toolbar_unset_style().
func (toolbar Toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(toolbar.native())
}

// Object ToplevelAccessible
type ToplevelAccessible struct {
	atk.Object
}

func (v ToplevelAccessible) native() *C.GtkToplevelAccessible {
	return (*C.GtkToplevelAccessible)(v.Ptr)
}
func wrapToplevelAccessible(p *C.GtkToplevelAccessible) (v ToplevelAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapToplevelAccessible(p unsafe.Pointer) (v ToplevelAccessible) {
	v.Ptr = p
	return
}
func (v ToplevelAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapToplevelAccessible(p unsafe.Pointer) interface{} {
	return WrapToplevelAccessible(p)
}
func (v ToplevelAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_toplevel_accessible_get_type())
}
func (v ToplevelAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapToplevelAccessible(unsafe.Pointer(ptr)), nil
	}
}

// GetChildren is a wrapper around gtk_toplevel_accessible_get_children().
func (accessible ToplevelAccessible) GetChildren() glib.List {
	ret0 := C.gtk_toplevel_accessible_get_children(accessible.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// Object TreeModelFilter
type TreeModelFilter struct {
	gobject.Object
}

func (v TreeModelFilter) native() *C.GtkTreeModelFilter {
	return (*C.GtkTreeModelFilter)(v.Ptr)
}
func wrapTreeModelFilter(p *C.GtkTreeModelFilter) (v TreeModelFilter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTreeModelFilter(p unsafe.Pointer) (v TreeModelFilter) {
	v.Ptr = p
	return
}
func (v TreeModelFilter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeModelFilter(p unsafe.Pointer) interface{} {
	return WrapTreeModelFilter(p)
}
func (v TreeModelFilter) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_model_filter_get_type())
}
func (v TreeModelFilter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeModelFilter(unsafe.Pointer(ptr)), nil
	}
}
func (v TreeModelFilter) TreeDragSource() TreeDragSource {
	return WrapTreeDragSource(v.Ptr)
}
func (v TreeModelFilter) TreeModel() TreeModel {
	return WrapTreeModel(v.Ptr)
}

// ClearCache is a wrapper around gtk_tree_model_filter_clear_cache().
func (filter TreeModelFilter) ClearCache() {
	C.gtk_tree_model_filter_clear_cache(filter.native())
}

// ConvertChildPathToPath is a wrapper around gtk_tree_model_filter_convert_child_path_to_path().
func (filter TreeModelFilter) ConvertChildPathToPath(child_path TreePath) TreePath {
	ret0 := C.gtk_tree_model_filter_convert_child_path_to_path(filter.native(), child_path.native())
	return wrapTreePath(ret0)
}

// ConvertPathToChildPath is a wrapper around gtk_tree_model_filter_convert_path_to_child_path().
func (filter TreeModelFilter) ConvertPathToChildPath(filter_path TreePath) TreePath {
	ret0 := C.gtk_tree_model_filter_convert_path_to_child_path(filter.native(), filter_path.native())
	return wrapTreePath(ret0)
}

// GetModel is a wrapper around gtk_tree_model_filter_get_model().
func (filter TreeModelFilter) GetModel() TreeModel {
	ret0 := C.gtk_tree_model_filter_get_model(filter.native())
	return wrapTreeModel(ret0)
}

// Refilter is a wrapper around gtk_tree_model_filter_refilter().
func (filter TreeModelFilter) Refilter() {
	C.gtk_tree_model_filter_refilter(filter.native())
}

// SetVisibleColumn is a wrapper around gtk_tree_model_filter_set_visible_column().
func (filter TreeModelFilter) SetVisibleColumn(column int) {
	C.gtk_tree_model_filter_set_visible_column(filter.native(), C.gint(column))
}

// Object TreeModelSort
type TreeModelSort struct {
	gobject.Object
}

func (v TreeModelSort) native() *C.GtkTreeModelSort {
	return (*C.GtkTreeModelSort)(v.Ptr)
}
func wrapTreeModelSort(p *C.GtkTreeModelSort) (v TreeModelSort) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTreeModelSort(p unsafe.Pointer) (v TreeModelSort) {
	v.Ptr = p
	return
}
func (v TreeModelSort) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeModelSort(p unsafe.Pointer) interface{} {
	return WrapTreeModelSort(p)
}
func (v TreeModelSort) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_model_sort_get_type())
}
func (v TreeModelSort) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeModelSort(unsafe.Pointer(ptr)), nil
	}
}
func (v TreeModelSort) TreeDragSource() TreeDragSource {
	return WrapTreeDragSource(v.Ptr)
}
func (v TreeModelSort) TreeModel() TreeModel {
	return WrapTreeModel(v.Ptr)
}
func (v TreeModelSort) TreeSortable() TreeSortable {
	return WrapTreeSortable(v.Ptr)
}

// ClearCache is a wrapper around gtk_tree_model_sort_clear_cache().
func (tree_model_sort TreeModelSort) ClearCache() {
	C.gtk_tree_model_sort_clear_cache(tree_model_sort.native())
}

// ConvertChildPathToPath is a wrapper around gtk_tree_model_sort_convert_child_path_to_path().
func (tree_model_sort TreeModelSort) ConvertChildPathToPath(child_path TreePath) TreePath {
	ret0 := C.gtk_tree_model_sort_convert_child_path_to_path(tree_model_sort.native(), child_path.native())
	return wrapTreePath(ret0)
}

// ConvertPathToChildPath is a wrapper around gtk_tree_model_sort_convert_path_to_child_path().
func (tree_model_sort TreeModelSort) ConvertPathToChildPath(sorted_path TreePath) TreePath {
	ret0 := C.gtk_tree_model_sort_convert_path_to_child_path(tree_model_sort.native(), sorted_path.native())
	return wrapTreePath(ret0)
}

// GetModel is a wrapper around gtk_tree_model_sort_get_model().
func (tree_model TreeModelSort) GetModel() TreeModel {
	ret0 := C.gtk_tree_model_sort_get_model(tree_model.native())
	return wrapTreeModel(ret0)
}

// IterIsValid is a wrapper around gtk_tree_model_sort_iter_is_valid().
func (tree_model_sort TreeModelSort) IterIsValid(iter TreeIter) bool {
	ret0 := C.gtk_tree_model_sort_iter_is_valid(tree_model_sort.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ResetDefaultSortFunc is a wrapper around gtk_tree_model_sort_reset_default_sort_func().
func (tree_model_sort TreeModelSort) ResetDefaultSortFunc() {
	C.gtk_tree_model_sort_reset_default_sort_func(tree_model_sort.native())
}

// Object TreeSelection
type TreeSelection struct {
	gobject.Object
}

func (v TreeSelection) native() *C.GtkTreeSelection {
	return (*C.GtkTreeSelection)(v.Ptr)
}
func wrapTreeSelection(p *C.GtkTreeSelection) (v TreeSelection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTreeSelection(p unsafe.Pointer) (v TreeSelection) {
	v.Ptr = p
	return
}
func (v TreeSelection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeSelection(p unsafe.Pointer) interface{} {
	return WrapTreeSelection(p)
}
func (v TreeSelection) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_selection_get_type())
}
func (v TreeSelection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeSelection(unsafe.Pointer(ptr)), nil
	}
}

// CountSelectedRows is a wrapper around gtk_tree_selection_count_selected_rows().
func (selection TreeSelection) CountSelectedRows() int {
	ret0 := C.gtk_tree_selection_count_selected_rows(selection.native())
	return int(ret0)
}

// GetMode is a wrapper around gtk_tree_selection_get_mode().
func (selection TreeSelection) GetMode() SelectionMode {
	ret0 := C.gtk_tree_selection_get_mode(selection.native())
	return SelectionMode(ret0)
}

// GetSelectedRows is a wrapper around gtk_tree_selection_get_selected_rows().
func (selection TreeSelection) GetSelectedRows() (glib.List, TreeModel) {
	var model0 *C.GtkTreeModel
	ret0 := C.gtk_tree_selection_get_selected_rows(selection.native(), &model0)
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapTreePath(p) }), /*gir:GLib*/ wrapTreeModel(model0)
}

// GetTreeView is a wrapper around gtk_tree_selection_get_tree_view().
func (selection TreeSelection) GetTreeView() TreeView {
	ret0 := C.gtk_tree_selection_get_tree_view(selection.native())
	return wrapTreeView(ret0)
}

// GetUserData is a wrapper around gtk_tree_selection_get_user_data().
func (selection TreeSelection) GetUserData() unsafe.Pointer {
	ret0 := C.gtk_tree_selection_get_user_data(selection.native())
	return unsafe.Pointer(ret0)
}

// IterIsSelected is a wrapper around gtk_tree_selection_iter_is_selected().
func (selection TreeSelection) IterIsSelected(iter TreeIter) bool {
	ret0 := C.gtk_tree_selection_iter_is_selected(selection.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PathIsSelected is a wrapper around gtk_tree_selection_path_is_selected().
func (selection TreeSelection) PathIsSelected(path TreePath) bool {
	ret0 := C.gtk_tree_selection_path_is_selected(selection.native(), path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SelectAll is a wrapper around gtk_tree_selection_select_all().
func (selection TreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all(selection.native())
}

// SelectIter is a wrapper around gtk_tree_selection_select_iter().
func (selection TreeSelection) SelectIter(iter TreeIter) {
	C.gtk_tree_selection_select_iter(selection.native(), iter.native())
}

// SelectPath is a wrapper around gtk_tree_selection_select_path().
func (selection TreeSelection) SelectPath(path TreePath) {
	C.gtk_tree_selection_select_path(selection.native(), path.native())
}

// SelectRange is a wrapper around gtk_tree_selection_select_range().
func (selection TreeSelection) SelectRange(start_path TreePath, end_path TreePath) {
	C.gtk_tree_selection_select_range(selection.native(), start_path.native(), end_path.native())
}

// SetMode is a wrapper around gtk_tree_selection_set_mode().
func (selection TreeSelection) SetMode(type_ SelectionMode) {
	C.gtk_tree_selection_set_mode(selection.native(), C.GtkSelectionMode(type_))
}

// UnselectAll is a wrapper around gtk_tree_selection_unselect_all().
func (selection TreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all(selection.native())
}

// UnselectIter is a wrapper around gtk_tree_selection_unselect_iter().
func (selection TreeSelection) UnselectIter(iter TreeIter) {
	C.gtk_tree_selection_unselect_iter(selection.native(), iter.native())
}

// UnselectPath is a wrapper around gtk_tree_selection_unselect_path().
func (selection TreeSelection) UnselectPath(path TreePath) {
	C.gtk_tree_selection_unselect_path(selection.native(), path.native())
}

// UnselectRange is a wrapper around gtk_tree_selection_unselect_range().
func (selection TreeSelection) UnselectRange(start_path TreePath, end_path TreePath) {
	C.gtk_tree_selection_unselect_range(selection.native(), start_path.native(), end_path.native())
}

// Object TreeView
type TreeView struct {
	Container
}

func (v TreeView) native() *C.GtkTreeView {
	return (*C.GtkTreeView)(v.Ptr)
}
func wrapTreeView(p *C.GtkTreeView) (v TreeView) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTreeView(p unsafe.Pointer) (v TreeView) {
	v.Ptr = p
	return
}
func (v TreeView) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeView(p unsafe.Pointer) interface{} {
	return WrapTreeView(p)
}
func (v TreeView) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_view_get_type())
}
func (v TreeView) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeView(unsafe.Pointer(ptr)), nil
	}
}
func (v TreeView) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v TreeView) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v TreeView) Scrollable() Scrollable {
	return WrapScrollable(v.Ptr)
}

// TreeViewNew is a wrapper around gtk_tree_view_new().
func TreeViewNew() Widget {
	ret0 := C.gtk_tree_view_new()
	return wrapWidget(ret0)
}

// TreeViewNewWithModel is a wrapper around gtk_tree_view_new_with_model().
func TreeViewNewWithModel(model TreeModel) Widget {
	ret0 := C.gtk_tree_view_new_with_model(model.native())
	return wrapWidget(ret0)
}

// AppendColumn is a wrapper around gtk_tree_view_append_column().
func (tree_view TreeView) AppendColumn(column TreeViewColumn) int {
	ret0 := C.gtk_tree_view_append_column(tree_view.native(), column.native())
	return int(ret0)
}

// CollapseAll is a wrapper around gtk_tree_view_collapse_all().
func (tree_view TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(tree_view.native())
}

// CollapseRow is a wrapper around gtk_tree_view_collapse_row().
func (tree_view TreeView) CollapseRow(path TreePath) bool {
	ret0 := C.gtk_tree_view_collapse_row(tree_view.native(), path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ColumnsAutosize is a wrapper around gtk_tree_view_columns_autosize().
func (tree_view TreeView) ColumnsAutosize() {
	C.gtk_tree_view_columns_autosize(tree_view.native())
}

// ConvertBinWindowToTreeCoords is a wrapper around gtk_tree_view_convert_bin_window_to_tree_coords().
func (tree_view TreeView) ConvertBinWindowToTreeCoords(bx int, by int) (int, int) {
	var tx0 C.gint
	var ty0 C.gint
	C.gtk_tree_view_convert_bin_window_to_tree_coords(tree_view.native(), C.gint(bx), C.gint(by), &tx0, &ty0)
	return int(tx0), int(ty0)
}

// ConvertBinWindowToWidgetCoords is a wrapper around gtk_tree_view_convert_bin_window_to_widget_coords().
func (tree_view TreeView) ConvertBinWindowToWidgetCoords(bx int, by int) (int, int) {
	var wx0 C.gint
	var wy0 C.gint
	C.gtk_tree_view_convert_bin_window_to_widget_coords(tree_view.native(), C.gint(bx), C.gint(by), &wx0, &wy0)
	return int(wx0), int(wy0)
}

// ConvertTreeToBinWindowCoords is a wrapper around gtk_tree_view_convert_tree_to_bin_window_coords().
func (tree_view TreeView) ConvertTreeToBinWindowCoords(tx int, ty int) (int, int) {
	var bx0 C.gint
	var by0 C.gint
	C.gtk_tree_view_convert_tree_to_bin_window_coords(tree_view.native(), C.gint(tx), C.gint(ty), &bx0, &by0)
	return int(bx0), int(by0)
}

// ConvertTreeToWidgetCoords is a wrapper around gtk_tree_view_convert_tree_to_widget_coords().
func (tree_view TreeView) ConvertTreeToWidgetCoords(tx int, ty int) (int, int) {
	var wx0 C.gint
	var wy0 C.gint
	C.gtk_tree_view_convert_tree_to_widget_coords(tree_view.native(), C.gint(tx), C.gint(ty), &wx0, &wy0)
	return int(wx0), int(wy0)
}

// ConvertWidgetToBinWindowCoords is a wrapper around gtk_tree_view_convert_widget_to_bin_window_coords().
func (tree_view TreeView) ConvertWidgetToBinWindowCoords(wx int, wy int) (int, int) {
	var bx0 C.gint
	var by0 C.gint
	C.gtk_tree_view_convert_widget_to_bin_window_coords(tree_view.native(), C.gint(wx), C.gint(wy), &bx0, &by0)
	return int(bx0), int(by0)
}

// ConvertWidgetToTreeCoords is a wrapper around gtk_tree_view_convert_widget_to_tree_coords().
func (tree_view TreeView) ConvertWidgetToTreeCoords(wx int, wy int) (int, int) {
	var tx0 C.gint
	var ty0 C.gint
	C.gtk_tree_view_convert_widget_to_tree_coords(tree_view.native(), C.gint(wx), C.gint(wy), &tx0, &ty0)
	return int(tx0), int(ty0)
}

// CreateRowDragIcon is a wrapper around gtk_tree_view_create_row_drag_icon().
func (tree_view TreeView) CreateRowDragIcon(path TreePath) cairo.Surface {
	ret0 := C.gtk_tree_view_create_row_drag_icon(tree_view.native(), path.native())
	return cairo.WrapSurface(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// ExpandAll is a wrapper around gtk_tree_view_expand_all().
func (tree_view TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(tree_view.native())
}

// ExpandRow is a wrapper around gtk_tree_view_expand_row().
func (tree_view TreeView) ExpandRow(path TreePath, open_all bool) bool {
	ret0 := C.gtk_tree_view_expand_row(tree_view.native(), path.native(), C.gboolean(util.Bool2Int(open_all)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ExpandToPath is a wrapper around gtk_tree_view_expand_to_path().
func (tree_view TreeView) ExpandToPath(path TreePath) {
	C.gtk_tree_view_expand_to_path(tree_view.native(), path.native())
}

// GetActivateOnSingleClick is a wrapper around gtk_tree_view_get_activate_on_single_click().
func (tree_view TreeView) GetActivateOnSingleClick() bool {
	ret0 := C.gtk_tree_view_get_activate_on_single_click(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetColumn is a wrapper around gtk_tree_view_get_column().
func (tree_view TreeView) GetColumn(n int) TreeViewColumn {
	ret0 := C.gtk_tree_view_get_column(tree_view.native(), C.gint(n))
	return wrapTreeViewColumn(ret0)
}

// GetColumns is a wrapper around gtk_tree_view_get_columns().
func (tree_view TreeView) GetColumns() glib.List {
	ret0 := C.gtk_tree_view_get_columns(tree_view.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapTreeViewColumn(p) }) /*gir:GLib*/
}

// GetCursor is a wrapper around gtk_tree_view_get_cursor().
func (tree_view TreeView) GetCursor() (TreePath, TreeViewColumn) {
	var path0 *C.GtkTreePath
	var focus_column0 *C.GtkTreeViewColumn
	C.gtk_tree_view_get_cursor(tree_view.native(), &path0, &focus_column0)
	return wrapTreePath(path0), wrapTreeViewColumn(focus_column0)
}

// GetDestRowAtPos is a wrapper around gtk_tree_view_get_dest_row_at_pos().
func (tree_view TreeView) GetDestRowAtPos(drag_x int, drag_y int) (bool, TreePath, TreeViewDropPosition) {
	var path0 *C.GtkTreePath
	var pos0 C.GtkTreeViewDropPosition
	ret0 := C.gtk_tree_view_get_dest_row_at_pos(tree_view.native(), C.gint(drag_x), C.gint(drag_y), &path0, &pos0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(path0), TreeViewDropPosition(pos0)
}

// GetDragDestRow is a wrapper around gtk_tree_view_get_drag_dest_row().
func (tree_view TreeView) GetDragDestRow() (TreePath, TreeViewDropPosition) {
	var path0 *C.GtkTreePath
	var pos0 C.GtkTreeViewDropPosition
	C.gtk_tree_view_get_drag_dest_row(tree_view.native(), &path0, &pos0)
	return wrapTreePath(path0), TreeViewDropPosition(pos0)
}

// GetEnableSearch is a wrapper around gtk_tree_view_get_enable_search().
func (tree_view TreeView) GetEnableSearch() bool {
	ret0 := C.gtk_tree_view_get_enable_search(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetEnableTreeLines is a wrapper around gtk_tree_view_get_enable_tree_lines().
func (tree_view TreeView) GetEnableTreeLines() bool {
	ret0 := C.gtk_tree_view_get_enable_tree_lines(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetExpanderColumn is a wrapper around gtk_tree_view_get_expander_column().
func (tree_view TreeView) GetExpanderColumn() TreeViewColumn {
	ret0 := C.gtk_tree_view_get_expander_column(tree_view.native())
	return wrapTreeViewColumn(ret0)
}

// GetFixedHeightMode is a wrapper around gtk_tree_view_get_fixed_height_mode().
func (tree_view TreeView) GetFixedHeightMode() bool {
	ret0 := C.gtk_tree_view_get_fixed_height_mode(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetGridLines is a wrapper around gtk_tree_view_get_grid_lines().
func (tree_view TreeView) GetGridLines() TreeViewGridLines {
	ret0 := C.gtk_tree_view_get_grid_lines(tree_view.native())
	return TreeViewGridLines(ret0)
}

// GetHeadersClickable is a wrapper around gtk_tree_view_get_headers_clickable().
func (tree_view TreeView) GetHeadersClickable() bool {
	ret0 := C.gtk_tree_view_get_headers_clickable(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHeadersVisible is a wrapper around gtk_tree_view_get_headers_visible().
func (tree_view TreeView) GetHeadersVisible() bool {
	ret0 := C.gtk_tree_view_get_headers_visible(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHoverExpand is a wrapper around gtk_tree_view_get_hover_expand().
func (tree_view TreeView) GetHoverExpand() bool {
	ret0 := C.gtk_tree_view_get_hover_expand(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetHoverSelection is a wrapper around gtk_tree_view_get_hover_selection().
func (tree_view TreeView) GetHoverSelection() bool {
	ret0 := C.gtk_tree_view_get_hover_selection(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLevelIndentation is a wrapper around gtk_tree_view_get_level_indentation().
func (tree_view TreeView) GetLevelIndentation() int {
	ret0 := C.gtk_tree_view_get_level_indentation(tree_view.native())
	return int(ret0)
}

// GetModel is a wrapper around gtk_tree_view_get_model().
func (tree_view TreeView) GetModel() TreeModel {
	ret0 := C.gtk_tree_view_get_model(tree_view.native())
	return wrapTreeModel(ret0)
}

// GetNColumns is a wrapper around gtk_tree_view_get_n_columns().
func (tree_view TreeView) GetNColumns() uint {
	ret0 := C.gtk_tree_view_get_n_columns(tree_view.native())
	return uint(ret0)
}

// GetPathAtPos is a wrapper around gtk_tree_view_get_path_at_pos().
func (tree_view TreeView) GetPathAtPos(x int, y int) (bool, TreePath, TreeViewColumn, int, int) {
	var path0 *C.GtkTreePath
	var column0 *C.GtkTreeViewColumn
	var cell_x0 C.gint
	var cell_y0 C.gint
	ret0 := C.gtk_tree_view_get_path_at_pos(tree_view.native(), C.gint(x), C.gint(y), &path0, &column0, &cell_x0, &cell_y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(path0), wrapTreeViewColumn(column0), int(cell_x0), int(cell_y0)
}

// GetReorderable is a wrapper around gtk_tree_view_get_reorderable().
func (tree_view TreeView) GetReorderable() bool {
	ret0 := C.gtk_tree_view_get_reorderable(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRubberBanding is a wrapper around gtk_tree_view_get_rubber_banding().
func (tree_view TreeView) GetRubberBanding() bool {
	ret0 := C.gtk_tree_view_get_rubber_banding(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSearchColumn is a wrapper around gtk_tree_view_get_search_column().
func (tree_view TreeView) GetSearchColumn() int {
	ret0 := C.gtk_tree_view_get_search_column(tree_view.native())
	return int(ret0)
}

// GetSearchEntry is a wrapper around gtk_tree_view_get_search_entry().
func (tree_view TreeView) GetSearchEntry() Entry {
	ret0 := C.gtk_tree_view_get_search_entry(tree_view.native())
	return wrapEntry(ret0)
}

// GetSelection is a wrapper around gtk_tree_view_get_selection().
func (tree_view TreeView) GetSelection() TreeSelection {
	ret0 := C.gtk_tree_view_get_selection(tree_view.native())
	return wrapTreeSelection(ret0)
}

// GetShowExpanders is a wrapper around gtk_tree_view_get_show_expanders().
func (tree_view TreeView) GetShowExpanders() bool {
	ret0 := C.gtk_tree_view_get_show_expanders(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTooltipColumn is a wrapper around gtk_tree_view_get_tooltip_column().
func (tree_view TreeView) GetTooltipColumn() int {
	ret0 := C.gtk_tree_view_get_tooltip_column(tree_view.native())
	return int(ret0)
}

// GetVisibleRange is a wrapper around gtk_tree_view_get_visible_range().
func (tree_view TreeView) GetVisibleRange() (bool, TreePath, TreePath) {
	var start_path0 *C.GtkTreePath
	var end_path0 *C.GtkTreePath
	ret0 := C.gtk_tree_view_get_visible_range(tree_view.native(), &start_path0, &end_path0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(start_path0), wrapTreePath(end_path0)
}

// InsertColumn is a wrapper around gtk_tree_view_insert_column().
func (tree_view TreeView) InsertColumn(column TreeViewColumn, position int) int {
	ret0 := C.gtk_tree_view_insert_column(tree_view.native(), column.native(), C.gint(position))
	return int(ret0)
}

// IsBlankAtPos is a wrapper around gtk_tree_view_is_blank_at_pos().
func (tree_view TreeView) IsBlankAtPos(x int, y int) (bool, TreePath, TreeViewColumn, int, int) {
	var path0 *C.GtkTreePath
	var column0 *C.GtkTreeViewColumn
	var cell_x0 C.gint
	var cell_y0 C.gint
	ret0 := C.gtk_tree_view_is_blank_at_pos(tree_view.native(), C.gint(x), C.gint(y), &path0, &column0, &cell_x0, &cell_y0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapTreePath(path0), wrapTreeViewColumn(column0), int(cell_x0), int(cell_y0)
}

// IsRubberBandingActive is a wrapper around gtk_tree_view_is_rubber_banding_active().
func (tree_view TreeView) IsRubberBandingActive() bool {
	ret0 := C.gtk_tree_view_is_rubber_banding_active(tree_view.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MoveColumnAfter is a wrapper around gtk_tree_view_move_column_after().
func (tree_view TreeView) MoveColumnAfter(column TreeViewColumn, base_column TreeViewColumn) {
	C.gtk_tree_view_move_column_after(tree_view.native(), column.native(), base_column.native())
}

// RemoveColumn is a wrapper around gtk_tree_view_remove_column().
func (tree_view TreeView) RemoveColumn(column TreeViewColumn) int {
	ret0 := C.gtk_tree_view_remove_column(tree_view.native(), column.native())
	return int(ret0)
}

// RowActivated is a wrapper around gtk_tree_view_row_activated().
func (tree_view TreeView) RowActivated(path TreePath, column TreeViewColumn) {
	C.gtk_tree_view_row_activated(tree_view.native(), path.native(), column.native())
}

// RowExpanded is a wrapper around gtk_tree_view_row_expanded().
func (tree_view TreeView) RowExpanded(path TreePath) bool {
	ret0 := C.gtk_tree_view_row_expanded(tree_view.native(), path.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ScrollToCell is a wrapper around gtk_tree_view_scroll_to_cell().
func (tree_view TreeView) ScrollToCell(path TreePath, column TreeViewColumn, use_align bool, row_align float32, col_align float32) {
	C.gtk_tree_view_scroll_to_cell(tree_view.native(), path.native(), column.native(), C.gboolean(util.Bool2Int(use_align)) /*go:.util*/, C.gfloat(row_align), C.gfloat(col_align))
}

// ScrollToPoint is a wrapper around gtk_tree_view_scroll_to_point().
func (tree_view TreeView) ScrollToPoint(tree_x int, tree_y int) {
	C.gtk_tree_view_scroll_to_point(tree_view.native(), C.gint(tree_x), C.gint(tree_y))
}

// SetActivateOnSingleClick is a wrapper around gtk_tree_view_set_activate_on_single_click().
func (tree_view TreeView) SetActivateOnSingleClick(single bool) {
	C.gtk_tree_view_set_activate_on_single_click(tree_view.native(), C.gboolean(util.Bool2Int(single)) /*go:.util*/)
}

// SetCursor is a wrapper around gtk_tree_view_set_cursor().
func (tree_view TreeView) SetCursor(path TreePath, focus_column TreeViewColumn, start_editing bool) {
	C.gtk_tree_view_set_cursor(tree_view.native(), path.native(), focus_column.native(), C.gboolean(util.Bool2Int(start_editing)) /*go:.util*/)
}

// SetCursorOnCell is a wrapper around gtk_tree_view_set_cursor_on_cell().
func (tree_view TreeView) SetCursorOnCell(path TreePath, focus_column TreeViewColumn, focus_cell CellRenderer, start_editing bool) {
	C.gtk_tree_view_set_cursor_on_cell(tree_view.native(), path.native(), focus_column.native(), focus_cell.native(), C.gboolean(util.Bool2Int(start_editing)) /*go:.util*/)
}

// SetDragDestRow is a wrapper around gtk_tree_view_set_drag_dest_row().
func (tree_view TreeView) SetDragDestRow(path TreePath, pos TreeViewDropPosition) {
	C.gtk_tree_view_set_drag_dest_row(tree_view.native(), path.native(), C.GtkTreeViewDropPosition(pos))
}

// SetEnableSearch is a wrapper around gtk_tree_view_set_enable_search().
func (tree_view TreeView) SetEnableSearch(enable_search bool) {
	C.gtk_tree_view_set_enable_search(tree_view.native(), C.gboolean(util.Bool2Int(enable_search)) /*go:.util*/)
}

// SetEnableTreeLines is a wrapper around gtk_tree_view_set_enable_tree_lines().
func (tree_view TreeView) SetEnableTreeLines(enabled bool) {
	C.gtk_tree_view_set_enable_tree_lines(tree_view.native(), C.gboolean(util.Bool2Int(enabled)) /*go:.util*/)
}

// SetExpanderColumn is a wrapper around gtk_tree_view_set_expander_column().
func (tree_view TreeView) SetExpanderColumn(column TreeViewColumn) {
	C.gtk_tree_view_set_expander_column(tree_view.native(), column.native())
}

// SetFixedHeightMode is a wrapper around gtk_tree_view_set_fixed_height_mode().
func (tree_view TreeView) SetFixedHeightMode(enable bool) {
	C.gtk_tree_view_set_fixed_height_mode(tree_view.native(), C.gboolean(util.Bool2Int(enable)) /*go:.util*/)
}

// SetGridLines is a wrapper around gtk_tree_view_set_grid_lines().
func (tree_view TreeView) SetGridLines(grid_lines TreeViewGridLines) {
	C.gtk_tree_view_set_grid_lines(tree_view.native(), C.GtkTreeViewGridLines(grid_lines))
}

// SetHeadersClickable is a wrapper around gtk_tree_view_set_headers_clickable().
func (tree_view TreeView) SetHeadersClickable(setting bool) {
	C.gtk_tree_view_set_headers_clickable(tree_view.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetHeadersVisible is a wrapper around gtk_tree_view_set_headers_visible().
func (tree_view TreeView) SetHeadersVisible(headers_visible bool) {
	C.gtk_tree_view_set_headers_visible(tree_view.native(), C.gboolean(util.Bool2Int(headers_visible)) /*go:.util*/)
}

// SetHoverExpand is a wrapper around gtk_tree_view_set_hover_expand().
func (tree_view TreeView) SetHoverExpand(expand bool) {
	C.gtk_tree_view_set_hover_expand(tree_view.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// SetHoverSelection is a wrapper around gtk_tree_view_set_hover_selection().
func (tree_view TreeView) SetHoverSelection(hover bool) {
	C.gtk_tree_view_set_hover_selection(tree_view.native(), C.gboolean(util.Bool2Int(hover)) /*go:.util*/)
}

// SetLevelIndentation is a wrapper around gtk_tree_view_set_level_indentation().
func (tree_view TreeView) SetLevelIndentation(indentation int) {
	C.gtk_tree_view_set_level_indentation(tree_view.native(), C.gint(indentation))
}

// SetModel is a wrapper around gtk_tree_view_set_model().
func (tree_view TreeView) SetModel(model TreeModel) {
	C.gtk_tree_view_set_model(tree_view.native(), model.native())
}

// SetReorderable is a wrapper around gtk_tree_view_set_reorderable().
func (tree_view TreeView) SetReorderable(reorderable bool) {
	C.gtk_tree_view_set_reorderable(tree_view.native(), C.gboolean(util.Bool2Int(reorderable)) /*go:.util*/)
}

// SetRubberBanding is a wrapper around gtk_tree_view_set_rubber_banding().
func (tree_view TreeView) SetRubberBanding(enable bool) {
	C.gtk_tree_view_set_rubber_banding(tree_view.native(), C.gboolean(util.Bool2Int(enable)) /*go:.util*/)
}

// SetSearchColumn is a wrapper around gtk_tree_view_set_search_column().
func (tree_view TreeView) SetSearchColumn(column int) {
	C.gtk_tree_view_set_search_column(tree_view.native(), C.gint(column))
}

// SetSearchEntry is a wrapper around gtk_tree_view_set_search_entry().
func (tree_view TreeView) SetSearchEntry(entry Entry) {
	C.gtk_tree_view_set_search_entry(tree_view.native(), entry.native())
}

// SetShowExpanders is a wrapper around gtk_tree_view_set_show_expanders().
func (tree_view TreeView) SetShowExpanders(enabled bool) {
	C.gtk_tree_view_set_show_expanders(tree_view.native(), C.gboolean(util.Bool2Int(enabled)) /*go:.util*/)
}

// SetTooltipCell is a wrapper around gtk_tree_view_set_tooltip_cell().
func (tree_view TreeView) SetTooltipCell(tooltip Tooltip, path TreePath, column TreeViewColumn, cell CellRenderer) {
	C.gtk_tree_view_set_tooltip_cell(tree_view.native(), tooltip.native(), path.native(), column.native(), cell.native())
}

// SetTooltipColumn is a wrapper around gtk_tree_view_set_tooltip_column().
func (tree_view TreeView) SetTooltipColumn(column int) {
	C.gtk_tree_view_set_tooltip_column(tree_view.native(), C.gint(column))
}

// SetTooltipRow is a wrapper around gtk_tree_view_set_tooltip_row().
func (tree_view TreeView) SetTooltipRow(tooltip Tooltip, path TreePath) {
	C.gtk_tree_view_set_tooltip_row(tree_view.native(), tooltip.native(), path.native())
}

// UnsetRowsDragDest is a wrapper around gtk_tree_view_unset_rows_drag_dest().
func (tree_view TreeView) UnsetRowsDragDest() {
	C.gtk_tree_view_unset_rows_drag_dest(tree_view.native())
}

// UnsetRowsDragSource is a wrapper around gtk_tree_view_unset_rows_drag_source().
func (tree_view TreeView) UnsetRowsDragSource() {
	C.gtk_tree_view_unset_rows_drag_source(tree_view.native())
}

// Object TreeViewColumn
type TreeViewColumn struct {
	gobject.InitiallyUnowned
}

func (v TreeViewColumn) native() *C.GtkTreeViewColumn {
	return (*C.GtkTreeViewColumn)(v.Ptr)
}
func wrapTreeViewColumn(p *C.GtkTreeViewColumn) (v TreeViewColumn) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTreeViewColumn(p unsafe.Pointer) (v TreeViewColumn) {
	v.Ptr = p
	return
}
func (v TreeViewColumn) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeViewColumn(p unsafe.Pointer) interface{} {
	return WrapTreeViewColumn(p)
}
func (v TreeViewColumn) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_view_column_get_type())
}
func (v TreeViewColumn) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeViewColumn(unsafe.Pointer(ptr)), nil
	}
}
func (v TreeViewColumn) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v TreeViewColumn) CellLayout() CellLayout {
	return WrapCellLayout(v.Ptr)
}

// TreeViewColumnNew is a wrapper around gtk_tree_view_column_new().
func TreeViewColumnNew() TreeViewColumn {
	ret0 := C.gtk_tree_view_column_new()
	return wrapTreeViewColumn(ret0)
}

// TreeViewColumnNewWithArea is a wrapper around gtk_tree_view_column_new_with_area().
func TreeViewColumnNewWithArea(area CellArea) TreeViewColumn {
	ret0 := C.gtk_tree_view_column_new_with_area(area.native())
	return wrapTreeViewColumn(ret0)
}

// AddAttribute is a wrapper around gtk_tree_view_column_add_attribute().
func (tree_column TreeViewColumn) AddAttribute(cell_renderer CellRenderer, attribute string, column int) {
	attribute0 := (*C.gchar)(C.CString(attribute))
	C.gtk_tree_view_column_add_attribute(tree_column.native(), cell_renderer.native(), attribute0, C.gint(column))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// CellGetPosition is a wrapper around gtk_tree_view_column_cell_get_position().
func (tree_column TreeViewColumn) CellGetPosition(cell_renderer CellRenderer) (bool, int, int) {
	var x_offset0 C.gint
	var width0 C.gint
	ret0 := C.gtk_tree_view_column_cell_get_position(tree_column.native(), cell_renderer.native(), &x_offset0, &width0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(x_offset0), int(width0)
}

// CellIsVisible is a wrapper around gtk_tree_view_column_cell_is_visible().
func (tree_column TreeViewColumn) CellIsVisible() bool {
	ret0 := C.gtk_tree_view_column_cell_is_visible(tree_column.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CellSetCellData is a wrapper around gtk_tree_view_column_cell_set_cell_data().
func (tree_column TreeViewColumn) CellSetCellData(tree_model TreeModel, iter TreeIter, is_expander bool, is_expanded bool) {
	C.gtk_tree_view_column_cell_set_cell_data(tree_column.native(), tree_model.native(), iter.native(), C.gboolean(util.Bool2Int(is_expander)) /*go:.util*/, C.gboolean(util.Bool2Int(is_expanded)) /*go:.util*/)
}

// Clear is a wrapper around gtk_tree_view_column_clear().
func (tree_column TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(tree_column.native())
}

// ClearAttributes is a wrapper around gtk_tree_view_column_clear_attributes().
func (tree_column TreeViewColumn) ClearAttributes(cell_renderer CellRenderer) {
	C.gtk_tree_view_column_clear_attributes(tree_column.native(), cell_renderer.native())
}

// Clicked is a wrapper around gtk_tree_view_column_clicked().
func (tree_column TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(tree_column.native())
}

// FocusCell is a wrapper around gtk_tree_view_column_focus_cell().
func (tree_column TreeViewColumn) FocusCell(cell CellRenderer) {
	C.gtk_tree_view_column_focus_cell(tree_column.native(), cell.native())
}

// GetAlignment is a wrapper around gtk_tree_view_column_get_alignment().
func (tree_column TreeViewColumn) GetAlignment() float32 {
	ret0 := C.gtk_tree_view_column_get_alignment(tree_column.native())
	return float32(ret0)
}

// GetButton is a wrapper around gtk_tree_view_column_get_button().
func (tree_column TreeViewColumn) GetButton() Widget {
	ret0 := C.gtk_tree_view_column_get_button(tree_column.native())
	return wrapWidget(ret0)
}

// GetClickable is a wrapper around gtk_tree_view_column_get_clickable().
func (tree_column TreeViewColumn) GetClickable() bool {
	ret0 := C.gtk_tree_view_column_get_clickable(tree_column.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetExpand is a wrapper around gtk_tree_view_column_get_expand().
func (tree_column TreeViewColumn) GetExpand() bool {
	ret0 := C.gtk_tree_view_column_get_expand(tree_column.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFixedWidth is a wrapper around gtk_tree_view_column_get_fixed_width().
func (tree_column TreeViewColumn) GetFixedWidth() int {
	ret0 := C.gtk_tree_view_column_get_fixed_width(tree_column.native())
	return int(ret0)
}

// GetMaxWidth is a wrapper around gtk_tree_view_column_get_max_width().
func (tree_column TreeViewColumn) GetMaxWidth() int {
	ret0 := C.gtk_tree_view_column_get_max_width(tree_column.native())
	return int(ret0)
}

// GetMinWidth is a wrapper around gtk_tree_view_column_get_min_width().
func (tree_column TreeViewColumn) GetMinWidth() int {
	ret0 := C.gtk_tree_view_column_get_min_width(tree_column.native())
	return int(ret0)
}

// GetReorderable is a wrapper around gtk_tree_view_column_get_reorderable().
func (tree_column TreeViewColumn) GetReorderable() bool {
	ret0 := C.gtk_tree_view_column_get_reorderable(tree_column.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetResizable is a wrapper around gtk_tree_view_column_get_resizable().
func (tree_column TreeViewColumn) GetResizable() bool {
	ret0 := C.gtk_tree_view_column_get_resizable(tree_column.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSizing is a wrapper around gtk_tree_view_column_get_sizing().
func (tree_column TreeViewColumn) GetSizing() TreeViewColumnSizing {
	ret0 := C.gtk_tree_view_column_get_sizing(tree_column.native())
	return TreeViewColumnSizing(ret0)
}

// GetSortColumnId is a wrapper around gtk_tree_view_column_get_sort_column_id().
func (tree_column TreeViewColumn) GetSortColumnId() int {
	ret0 := C.gtk_tree_view_column_get_sort_column_id(tree_column.native())
	return int(ret0)
}

// GetSortIndicator is a wrapper around gtk_tree_view_column_get_sort_indicator().
func (tree_column TreeViewColumn) GetSortIndicator() bool {
	ret0 := C.gtk_tree_view_column_get_sort_indicator(tree_column.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSortOrder is a wrapper around gtk_tree_view_column_get_sort_order().
func (tree_column TreeViewColumn) GetSortOrder() SortType {
	ret0 := C.gtk_tree_view_column_get_sort_order(tree_column.native())
	return SortType(ret0)
}

// GetSpacing is a wrapper around gtk_tree_view_column_get_spacing().
func (tree_column TreeViewColumn) GetSpacing() int {
	ret0 := C.gtk_tree_view_column_get_spacing(tree_column.native())
	return int(ret0)
}

// GetTitle is a wrapper around gtk_tree_view_column_get_title().
func (tree_column TreeViewColumn) GetTitle() string {
	ret0 := C.gtk_tree_view_column_get_title(tree_column.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTreeView is a wrapper around gtk_tree_view_column_get_tree_view().
func (tree_column TreeViewColumn) GetTreeView() Widget {
	ret0 := C.gtk_tree_view_column_get_tree_view(tree_column.native())
	return wrapWidget(ret0)
}

// GetVisible is a wrapper around gtk_tree_view_column_get_visible().
func (tree_column TreeViewColumn) GetVisible() bool {
	ret0 := C.gtk_tree_view_column_get_visible(tree_column.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetWidget is a wrapper around gtk_tree_view_column_get_widget().
func (tree_column TreeViewColumn) GetWidget() Widget {
	ret0 := C.gtk_tree_view_column_get_widget(tree_column.native())
	return wrapWidget(ret0)
}

// GetWidth is a wrapper around gtk_tree_view_column_get_width().
func (tree_column TreeViewColumn) GetWidth() int {
	ret0 := C.gtk_tree_view_column_get_width(tree_column.native())
	return int(ret0)
}

// GetXOffset is a wrapper around gtk_tree_view_column_get_x_offset().
func (tree_column TreeViewColumn) GetXOffset() int {
	ret0 := C.gtk_tree_view_column_get_x_offset(tree_column.native())
	return int(ret0)
}

// PackEnd is a wrapper around gtk_tree_view_column_pack_end().
func (tree_column TreeViewColumn) PackEnd(cell CellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_end(tree_column.native(), cell.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// PackStart is a wrapper around gtk_tree_view_column_pack_start().
func (tree_column TreeViewColumn) PackStart(cell CellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_start(tree_column.native(), cell.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// QueueResize is a wrapper around gtk_tree_view_column_queue_resize().
func (tree_column TreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize(tree_column.native())
}

// SetAlignment is a wrapper around gtk_tree_view_column_set_alignment().
func (tree_column TreeViewColumn) SetAlignment(xalign float32) {
	C.gtk_tree_view_column_set_alignment(tree_column.native(), C.gfloat(xalign))
}

// SetClickable is a wrapper around gtk_tree_view_column_set_clickable().
func (tree_column TreeViewColumn) SetClickable(clickable bool) {
	C.gtk_tree_view_column_set_clickable(tree_column.native(), C.gboolean(util.Bool2Int(clickable)) /*go:.util*/)
}

// SetExpand is a wrapper around gtk_tree_view_column_set_expand().
func (tree_column TreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(tree_column.native(), C.gboolean(util.Bool2Int(expand)) /*go:.util*/)
}

// SetFixedWidth is a wrapper around gtk_tree_view_column_set_fixed_width().
func (tree_column TreeViewColumn) SetFixedWidth(fixed_width int) {
	C.gtk_tree_view_column_set_fixed_width(tree_column.native(), C.gint(fixed_width))
}

// SetMaxWidth is a wrapper around gtk_tree_view_column_set_max_width().
func (tree_column TreeViewColumn) SetMaxWidth(max_width int) {
	C.gtk_tree_view_column_set_max_width(tree_column.native(), C.gint(max_width))
}

// SetMinWidth is a wrapper around gtk_tree_view_column_set_min_width().
func (tree_column TreeViewColumn) SetMinWidth(min_width int) {
	C.gtk_tree_view_column_set_min_width(tree_column.native(), C.gint(min_width))
}

// SetReorderable is a wrapper around gtk_tree_view_column_set_reorderable().
func (tree_column TreeViewColumn) SetReorderable(reorderable bool) {
	C.gtk_tree_view_column_set_reorderable(tree_column.native(), C.gboolean(util.Bool2Int(reorderable)) /*go:.util*/)
}

// SetResizable is a wrapper around gtk_tree_view_column_set_resizable().
func (tree_column TreeViewColumn) SetResizable(resizable bool) {
	C.gtk_tree_view_column_set_resizable(tree_column.native(), C.gboolean(util.Bool2Int(resizable)) /*go:.util*/)
}

// SetSizing is a wrapper around gtk_tree_view_column_set_sizing().
func (tree_column TreeViewColumn) SetSizing(type_ TreeViewColumnSizing) {
	C.gtk_tree_view_column_set_sizing(tree_column.native(), C.GtkTreeViewColumnSizing(type_))
}

// SetSortColumnId is a wrapper around gtk_tree_view_column_set_sort_column_id().
func (tree_column TreeViewColumn) SetSortColumnId(sort_column_id int) {
	C.gtk_tree_view_column_set_sort_column_id(tree_column.native(), C.gint(sort_column_id))
}

// SetSortIndicator is a wrapper around gtk_tree_view_column_set_sort_indicator().
func (tree_column TreeViewColumn) SetSortIndicator(setting bool) {
	C.gtk_tree_view_column_set_sort_indicator(tree_column.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetSortOrder is a wrapper around gtk_tree_view_column_set_sort_order().
func (tree_column TreeViewColumn) SetSortOrder(order SortType) {
	C.gtk_tree_view_column_set_sort_order(tree_column.native(), C.GtkSortType(order))
}

// SetSpacing is a wrapper around gtk_tree_view_column_set_spacing().
func (tree_column TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(tree_column.native(), C.gint(spacing))
}

// SetTitle is a wrapper around gtk_tree_view_column_set_title().
func (tree_column TreeViewColumn) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.gtk_tree_view_column_set_title(tree_column.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetVisible is a wrapper around gtk_tree_view_column_set_visible().
func (tree_column TreeViewColumn) SetVisible(visible bool) {
	C.gtk_tree_view_column_set_visible(tree_column.native(), C.gboolean(util.Bool2Int(visible)) /*go:.util*/)
}

// SetWidget is a wrapper around gtk_tree_view_column_set_widget().
func (tree_column TreeViewColumn) SetWidget(widget Widget) {
	C.gtk_tree_view_column_set_widget(tree_column.native(), widget.native())
}

// Object TreeStore
type TreeStore struct {
	gobject.Object
}

func (v TreeStore) native() *C.GtkTreeStore {
	return (*C.GtkTreeStore)(v.Ptr)
}
func wrapTreeStore(p *C.GtkTreeStore) (v TreeStore) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTreeStore(p unsafe.Pointer) (v TreeStore) {
	v.Ptr = p
	return
}
func (v TreeStore) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeStore(p unsafe.Pointer) interface{} {
	return WrapTreeStore(p)
}
func (v TreeStore) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_store_get_type())
}
func (v TreeStore) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeStore(unsafe.Pointer(ptr)), nil
	}
}
func (v TreeStore) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v TreeStore) TreeDragDest() TreeDragDest {
	return WrapTreeDragDest(v.Ptr)
}
func (v TreeStore) TreeDragSource() TreeDragSource {
	return WrapTreeDragSource(v.Ptr)
}
func (v TreeStore) TreeModel() TreeModel {
	return WrapTreeModel(v.Ptr)
}
func (v TreeStore) TreeSortable() TreeSortable {
	return WrapTreeSortable(v.Ptr)
}

// TreeStoreNew is a wrapper around gtk_tree_store_newv().
func TreeStoreNew(types [] /*Gir:GObject*/ gobject.Type) TreeStore {
	types0 := make([]C.GType, len(types))
	for idx, elemG := range types {
		types0[idx] = C.GType(elemG)
	}
	var types0Ptr *C.GType
	if len(types0) > 0 {
		types0Ptr = &types0[0]
	}
	ret0 := C.gtk_tree_store_newv(C.gint(len(types)), types0Ptr)
	return wrapTreeStore(ret0)
}

// Clear is a wrapper around gtk_tree_store_clear().
func (tree_store TreeStore) Clear() {
	C.gtk_tree_store_clear(tree_store.native())
}

// IsAncestor is a wrapper around gtk_tree_store_is_ancestor().
func (tree_store TreeStore) IsAncestor(iter TreeIter, descendant TreeIter) bool {
	ret0 := C.gtk_tree_store_is_ancestor(tree_store.native(), iter.native(), descendant.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IterDepth is a wrapper around gtk_tree_store_iter_depth().
func (tree_store TreeStore) IterDepth(iter TreeIter) int {
	ret0 := C.gtk_tree_store_iter_depth(tree_store.native(), iter.native())
	return int(ret0)
}

// IterIsValid is a wrapper around gtk_tree_store_iter_is_valid().
func (tree_store TreeStore) IterIsValid(iter TreeIter) bool {
	ret0 := C.gtk_tree_store_iter_is_valid(tree_store.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MoveAfter is a wrapper around gtk_tree_store_move_after().
func (tree_store TreeStore) MoveAfter(iter TreeIter, position TreeIter) {
	C.gtk_tree_store_move_after(tree_store.native(), iter.native(), position.native())
}

// MoveBefore is a wrapper around gtk_tree_store_move_before().
func (tree_store TreeStore) MoveBefore(iter TreeIter, position TreeIter) {
	C.gtk_tree_store_move_before(tree_store.native(), iter.native(), position.native())
}

// Remove is a wrapper around gtk_tree_store_remove().
func (tree_store TreeStore) Remove(iter TreeIter) bool {
	ret0 := C.gtk_tree_store_remove(tree_store.native(), iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Reorder is a wrapper around gtk_tree_store_reorder().
func (tree_store TreeStore) Reorder(parent TreeIter, new_order []int) {
	new_order0 := make([]C.gint, len(new_order))
	for idx, elemG := range new_order {
		new_order0[idx] = C.gint(elemG)
	}
	var new_order0Ptr *C.gint
	if len(new_order0) > 0 {
		new_order0Ptr = &new_order0[0]
	}
	C.gtk_tree_store_reorder(tree_store.native(), parent.native(), new_order0Ptr)
}

// SetColumnTypes is a wrapper around gtk_tree_store_set_column_types().
func (tree_store TreeStore) SetColumnTypes(types [] /*Gir:GObject*/ gobject.Type) {
	types0 := make([]C.GType, len(types))
	for idx, elemG := range types {
		types0[idx] = C.GType(elemG)
	}
	var types0Ptr *C.GType
	if len(types0) > 0 {
		types0Ptr = &types0[0]
	}
	C.gtk_tree_store_set_column_types(tree_store.native(), C.gint(len(types)), types0Ptr)
}

// SetValue is a wrapper around gtk_tree_store_set_value().
func (tree_store TreeStore) SetValue(iter TreeIter, column int, value gobject.Value) {
	C.gtk_tree_store_set_value(tree_store.native(), iter.native(), C.gint(column), (*C.GValue)(value.Ptr))
}

// Swap is a wrapper around gtk_tree_store_swap().
func (tree_store TreeStore) Swap(a TreeIter, b TreeIter) {
	C.gtk_tree_store_swap(tree_store.native(), a.native(), b.native())
}

// Object TreeViewAccessible
type TreeViewAccessible struct {
	ContainerAccessible
}

func (v TreeViewAccessible) native() *C.GtkTreeViewAccessible {
	return (*C.GtkTreeViewAccessible)(v.Ptr)
}
func wrapTreeViewAccessible(p *C.GtkTreeViewAccessible) (v TreeViewAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTreeViewAccessible(p unsafe.Pointer) (v TreeViewAccessible) {
	v.Ptr = p
	return
}
func (v TreeViewAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTreeViewAccessible(p unsafe.Pointer) interface{} {
	return WrapTreeViewAccessible(p)
}
func (v TreeViewAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_tree_view_accessible_get_type())
}
func (v TreeViewAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTreeViewAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v TreeViewAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v TreeViewAccessible) Selection() atk.Selection {
	return atk.WrapSelection(v.Ptr) /*gir:Atk*/
}
func (v TreeViewAccessible) Table() atk.Table {
	return atk.WrapTable(v.Ptr) /*gir:Atk*/
}
func (v TreeViewAccessible) CellAccessibleParent() CellAccessibleParent {
	return WrapCellAccessibleParent(v.Ptr)
}

// Object UIManager
type UIManager struct {
	gobject.Object
}

func (v UIManager) native() *C.GtkUIManager {
	return (*C.GtkUIManager)(v.Ptr)
}
func wrapUIManager(p *C.GtkUIManager) (v UIManager) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUIManager(p unsafe.Pointer) (v UIManager) {
	v.Ptr = p
	return
}
func (v UIManager) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUIManager(p unsafe.Pointer) interface{} {
	return WrapUIManager(p)
}
func (v UIManager) GetType() gobject.Type {
	return gobject.Type(C.gtk_ui_manager_get_type())
}
func (v UIManager) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUIManager(unsafe.Pointer(ptr)), nil
	}
}
func (v UIManager) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}

// Object VBox
type VBox struct {
	Box
}

func (v VBox) native() *C.GtkVBox {
	return (*C.GtkVBox)(v.Ptr)
}
func wrapVBox(p *C.GtkVBox) (v VBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVBox(p unsafe.Pointer) (v VBox) {
	v.Ptr = p
	return
}
func (v VBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVBox(p unsafe.Pointer) interface{} {
	return WrapVBox(p)
}
func (v VBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_vbox_get_type())
}
func (v VBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVBox(unsafe.Pointer(ptr)), nil
	}
}
func (v VBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v VBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v VBox) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object VButtonBox
type VButtonBox struct {
	ButtonBox
}

func (v VButtonBox) native() *C.GtkVButtonBox {
	return (*C.GtkVButtonBox)(v.Ptr)
}
func wrapVButtonBox(p *C.GtkVButtonBox) (v VButtonBox) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVButtonBox(p unsafe.Pointer) (v VButtonBox) {
	v.Ptr = p
	return
}
func (v VButtonBox) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVButtonBox(p unsafe.Pointer) interface{} {
	return WrapVButtonBox(p)
}
func (v VButtonBox) GetType() gobject.Type {
	return gobject.Type(C.gtk_vbutton_box_get_type())
}
func (v VButtonBox) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVButtonBox(unsafe.Pointer(ptr)), nil
	}
}
func (v VButtonBox) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v VButtonBox) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v VButtonBox) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object VPaned
type VPaned struct {
	Paned
}

func (v VPaned) native() *C.GtkVPaned {
	return (*C.GtkVPaned)(v.Ptr)
}
func wrapVPaned(p *C.GtkVPaned) (v VPaned) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVPaned(p unsafe.Pointer) (v VPaned) {
	v.Ptr = p
	return
}
func (v VPaned) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVPaned(p unsafe.Pointer) interface{} {
	return WrapVPaned(p)
}
func (v VPaned) GetType() gobject.Type {
	return gobject.Type(C.gtk_vpaned_get_type())
}
func (v VPaned) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVPaned(unsafe.Pointer(ptr)), nil
	}
}
func (v VPaned) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v VPaned) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v VPaned) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object VScale
type VScale struct {
	Scale
}

func (v VScale) native() *C.GtkVScale {
	return (*C.GtkVScale)(v.Ptr)
}
func wrapVScale(p *C.GtkVScale) (v VScale) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVScale(p unsafe.Pointer) (v VScale) {
	v.Ptr = p
	return
}
func (v VScale) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVScale(p unsafe.Pointer) interface{} {
	return WrapVScale(p)
}
func (v VScale) GetType() gobject.Type {
	return gobject.Type(C.gtk_vscale_get_type())
}
func (v VScale) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVScale(unsafe.Pointer(ptr)), nil
	}
}
func (v VScale) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v VScale) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v VScale) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object VScrollbar
type VScrollbar struct {
	Scrollbar
}

func (v VScrollbar) native() *C.GtkVScrollbar {
	return (*C.GtkVScrollbar)(v.Ptr)
}
func wrapVScrollbar(p *C.GtkVScrollbar) (v VScrollbar) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVScrollbar(p unsafe.Pointer) (v VScrollbar) {
	v.Ptr = p
	return
}
func (v VScrollbar) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVScrollbar(p unsafe.Pointer) interface{} {
	return WrapVScrollbar(p)
}
func (v VScrollbar) GetType() gobject.Type {
	return gobject.Type(C.gtk_vscrollbar_get_type())
}
func (v VScrollbar) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVScrollbar(unsafe.Pointer(ptr)), nil
	}
}
func (v VScrollbar) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v VScrollbar) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v VScrollbar) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Object VSeparator
type VSeparator struct {
	Separator
}

func (v VSeparator) native() *C.GtkVSeparator {
	return (*C.GtkVSeparator)(v.Ptr)
}
func wrapVSeparator(p *C.GtkVSeparator) (v VSeparator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVSeparator(p unsafe.Pointer) (v VSeparator) {
	v.Ptr = p
	return
}
func (v VSeparator) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVSeparator(p unsafe.Pointer) interface{} {
	return WrapVSeparator(p)
}
func (v VSeparator) GetType() gobject.Type {
	return gobject.Type(C.gtk_vseparator_get_type())
}
func (v VSeparator) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVSeparator(unsafe.Pointer(ptr)), nil
	}
}
func (v VSeparator) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v VSeparator) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v VSeparator) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// Interface CellAccessibleParent
type CellAccessibleParent struct {
	Ptr unsafe.Pointer
}

func (v CellAccessibleParent) native() *C.GtkCellAccessibleParent {
	return (*C.GtkCellAccessibleParent)(v.Ptr)
}
func wrapCellAccessibleParent(p *C.GtkCellAccessibleParent) CellAccessibleParent {
	return CellAccessibleParent{unsafe.Pointer(p)}
}
func WrapCellAccessibleParent(p unsafe.Pointer) CellAccessibleParent {
	return CellAccessibleParent{p}
}
func (v CellAccessibleParent) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCellAccessibleParent(p unsafe.Pointer) interface{} {
	return WrapCellAccessibleParent(p)
}
func (v CellAccessibleParent) GetType() gobject.Type {
	return gobject.Type(C.gtk_cell_accessible_parent_get_type())
}
func (v CellAccessibleParent) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCellAccessibleParent(unsafe.Pointer(ptr)), nil
	}
}

// Activate is a wrapper around gtk_cell_accessible_parent_activate().
func (parent CellAccessibleParent) Activate(cell CellAccessible) {
	C.gtk_cell_accessible_parent_activate(parent.native(), cell.native())
}

// Edit is a wrapper around gtk_cell_accessible_parent_edit().
func (parent CellAccessibleParent) Edit(cell CellAccessible) {
	C.gtk_cell_accessible_parent_edit(parent.native(), cell.native())
}

// ExpandCollapse is a wrapper around gtk_cell_accessible_parent_expand_collapse().
func (parent CellAccessibleParent) ExpandCollapse(cell CellAccessible) {
	C.gtk_cell_accessible_parent_expand_collapse(parent.native(), cell.native())
}

// GetChildIndex is a wrapper around gtk_cell_accessible_parent_get_child_index().
func (parent CellAccessibleParent) GetChildIndex(cell CellAccessible) int {
	ret0 := C.gtk_cell_accessible_parent_get_child_index(parent.native(), cell.native())
	return int(ret0)
}

// GetRendererState is a wrapper around gtk_cell_accessible_parent_get_renderer_state().
func (parent CellAccessibleParent) GetRendererState(cell CellAccessible) CellRendererState {
	ret0 := C.gtk_cell_accessible_parent_get_renderer_state(parent.native(), cell.native())
	return CellRendererState(ret0)
}

// GrabFocus is a wrapper around gtk_cell_accessible_parent_grab_focus().
func (parent CellAccessibleParent) GrabFocus(cell CellAccessible) bool {
	ret0 := C.gtk_cell_accessible_parent_grab_focus(parent.native(), cell.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// UpdateRelationset is a wrapper around gtk_cell_accessible_parent_update_relationset().
func (parent CellAccessibleParent) UpdateRelationset(cell CellAccessible, relationset atk.RelationSet) {
	C.gtk_cell_accessible_parent_update_relationset(parent.native(), cell.native(), (*C.AtkRelationSet)(relationset.Ptr))
}

// Object Viewport
type Viewport struct {
	Bin
}

func (v Viewport) native() *C.GtkViewport {
	return (*C.GtkViewport)(v.Ptr)
}
func wrapViewport(p *C.GtkViewport) (v Viewport) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapViewport(p unsafe.Pointer) (v Viewport) {
	v.Ptr = p
	return
}
func (v Viewport) IsNil() bool {
	return v.Ptr == nil
}
func IWrapViewport(p unsafe.Pointer) interface{} {
	return WrapViewport(p)
}
func (v Viewport) GetType() gobject.Type {
	return gobject.Type(C.gtk_viewport_get_type())
}
func (v Viewport) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapViewport(unsafe.Pointer(ptr)), nil
	}
}
func (v Viewport) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v Viewport) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v Viewport) Scrollable() Scrollable {
	return WrapScrollable(v.Ptr)
}

// ViewportNew is a wrapper around gtk_viewport_new().
func ViewportNew(hadjustment Adjustment, vadjustment Adjustment) Widget {
	ret0 := C.gtk_viewport_new(hadjustment.native(), vadjustment.native())
	return wrapWidget(ret0)
}

// GetBinWindow is a wrapper around gtk_viewport_get_bin_window().
func (viewport Viewport) GetBinWindow() gdk.Window {
	ret0 := C.gtk_viewport_get_bin_window(viewport.native())
	return gdk.WrapWindow(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// GetShadowType is a wrapper around gtk_viewport_get_shadow_type().
func (viewport Viewport) GetShadowType() ShadowType {
	ret0 := C.gtk_viewport_get_shadow_type(viewport.native())
	return ShadowType(ret0)
}

// GetViewWindow is a wrapper around gtk_viewport_get_view_window().
func (viewport Viewport) GetViewWindow() gdk.Window {
	ret0 := C.gtk_viewport_get_view_window(viewport.native())
	return gdk.WrapWindow(unsafe.Pointer(ret0)) /*gir:Gdk*/
}

// SetShadowType is a wrapper around gtk_viewport_set_shadow_type().
func (viewport Viewport) SetShadowType(type_ ShadowType) {
	C.gtk_viewport_set_shadow_type(viewport.native(), C.GtkShadowType(type_))
}

// Object VolumeButton
type VolumeButton struct {
	ScaleButton
}

func (v VolumeButton) native() *C.GtkVolumeButton {
	return (*C.GtkVolumeButton)(v.Ptr)
}
func wrapVolumeButton(p *C.GtkVolumeButton) (v VolumeButton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVolumeButton(p unsafe.Pointer) (v VolumeButton) {
	v.Ptr = p
	return
}
func (v VolumeButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVolumeButton(p unsafe.Pointer) interface{} {
	return WrapVolumeButton(p)
}
func (v VolumeButton) GetType() gobject.Type {
	return gobject.Type(C.gtk_volume_button_get_type())
}
func (v VolumeButton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVolumeButton(unsafe.Pointer(ptr)), nil
	}
}
func (v VolumeButton) ImplementorIface() atk.ImplementorIface {
	return atk.WrapImplementorIface(v.Ptr) /*gir:Atk*/
}
func (v VolumeButton) Actionable() Actionable {
	return WrapActionable(v.Ptr)
}
func (v VolumeButton) Activatable() Activatable {
	return WrapActivatable(v.Ptr)
}
func (v VolumeButton) Buildable() Buildable {
	return WrapBuildable(v.Ptr)
}
func (v VolumeButton) Orientable() Orientable {
	return WrapOrientable(v.Ptr)
}

// VolumeButtonNew is a wrapper around gtk_volume_button_new().
func VolumeButtonNew() Widget {
	ret0 := C.gtk_volume_button_new()
	return wrapWidget(ret0)
}

// Object WindowAccessible
type WindowAccessible struct {
	ContainerAccessible
}

func (v WindowAccessible) native() *C.GtkWindowAccessible {
	return (*C.GtkWindowAccessible)(v.Ptr)
}
func wrapWindowAccessible(p *C.GtkWindowAccessible) (v WindowAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapWindowAccessible(p unsafe.Pointer) (v WindowAccessible) {
	v.Ptr = p
	return
}
func (v WindowAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWindowAccessible(p unsafe.Pointer) interface{} {
	return WrapWindowAccessible(p)
}
func (v WindowAccessible) GetType() gobject.Type {
	return gobject.Type(C.gtk_window_accessible_get_type())
}
func (v WindowAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapWindowAccessible(unsafe.Pointer(ptr)), nil
	}
}
func (v WindowAccessible) Component() atk.Component {
	return atk.WrapComponent(v.Ptr) /*gir:Atk*/
}
func (v WindowAccessible) Window() atk.Window {
	return atk.WrapWindow(v.Ptr) /*gir:Atk*/
}

type Stock unsafe.Pointer
type Align int

const (
	AlignFill     Align = 0
	AlignStart          = 1
	AlignEnd            = 2
	AlignCenter         = 3
	AlignBaseline       = 4
)

type ArrowPlacement int

const (
	ArrowPlacementBoth  ArrowPlacement = 0
	ArrowPlacementStart                = 1
	ArrowPlacementEnd                  = 2
)

type ArrowType int

const (
	ArrowTypeUp    ArrowType = 0
	ArrowTypeDown            = 1
	ArrowTypeLeft            = 2
	ArrowTypeRight           = 3
	ArrowTypeNone            = 4
)

type AssistantPageType int

const (
	AssistantPageTypeContent  AssistantPageType = 0
	AssistantPageTypeIntro                      = 1
	AssistantPageTypeConfirm                    = 2
	AssistantPageTypeSummary                    = 3
	AssistantPageTypeProgress                   = 4
	AssistantPageTypeCustom                     = 5
)

type BaselinePosition int

const (
	BaselinePositionTop    BaselinePosition = 0
	BaselinePositionCenter                  = 1
	BaselinePositionBottom                  = 2
)

type BorderStyle int

const (
	BorderStyleNone   BorderStyle = 0
	BorderStyleSolid              = 1
	BorderStyleInset              = 2
	BorderStyleOutset             = 3
	BorderStyleHidden             = 4
	BorderStyleDotted             = 5
	BorderStyleDashed             = 6
	BorderStyleDouble             = 7
	BorderStyleGroove             = 8
	BorderStyleRidge              = 9
)

type BuilderError int

const (
	BuilderErrorInvalidTypeFunction  BuilderError = 0
	BuilderErrorUnhandledTag                      = 1
	BuilderErrorMissingAttribute                  = 2
	BuilderErrorInvalidAttribute                  = 3
	BuilderErrorInvalidTag                        = 4
	BuilderErrorMissingPropertyValue              = 5
	BuilderErrorInvalidValue                      = 6
	BuilderErrorVersionMismatch                   = 7
	BuilderErrorDuplicateId                       = 8
	BuilderErrorObjectTypeRefused                 = 9
	BuilderErrorTemplateMismatch                  = 10
	BuilderErrorInvalidProperty                   = 11
	BuilderErrorInvalidSignal                     = 12
	BuilderErrorInvalidId                         = 13
)

type ButtonBoxStyle int

const (
	ButtonBoxStyleSpread ButtonBoxStyle = 1
	ButtonBoxStyleEdge                  = 2
	ButtonBoxStyleStart                 = 3
	ButtonBoxStyleEnd                   = 4
	ButtonBoxStyleCenter                = 5
	ButtonBoxStyleExpand                = 6
)

type ButtonRole int

const (
	ButtonRoleNormal ButtonRole = 0
	ButtonRoleCheck             = 1
	ButtonRoleRadio             = 2
)

type ButtonsType int

const (
	ButtonsTypeNone     ButtonsType = 0
	ButtonsTypeOk                   = 1
	ButtonsTypeClose                = 2
	ButtonsTypeCancel               = 3
	ButtonsTypeYesNo                = 4
	ButtonsTypeOkCancel             = 5
)

type CellRendererAccelMode int

const (
	CellRendererAccelModeGtk   CellRendererAccelMode = 0
	CellRendererAccelModeOther                       = 1
)

type CellRendererMode int

const (
	CellRendererModeInert       CellRendererMode = 0
	CellRendererModeActivatable                  = 1
	CellRendererModeEditable                     = 2
)

type CornerType int

const (
	CornerTypeTopLeft     CornerType = 0
	CornerTypeBottomLeft             = 1
	CornerTypeTopRight               = 2
	CornerTypeBottomRight            = 3
)

type CssProviderError int

const (
	CssProviderErrorFailed       CssProviderError = 0
	CssProviderErrorSyntax                        = 1
	CssProviderErrorImport                        = 2
	CssProviderErrorName                          = 3
	CssProviderErrorDeprecated                    = 4
	CssProviderErrorUnknownValue                  = 5
)

type CssSectionType int

const (
	CssSectionTypeDocument        CssSectionType = 0
	CssSectionTypeImport                         = 1
	CssSectionTypeColorDefinition                = 2
	CssSectionTypeBindingSet                     = 3
	CssSectionTypeRuleset                        = 4
	CssSectionTypeSelector                       = 5
	CssSectionTypeDeclaration                    = 6
	CssSectionTypeValue                          = 7
	CssSectionTypeKeyframes                      = 8
)

type DeleteType int

const (
	DeleteTypeChars           DeleteType = 0
	DeleteTypeWordEnds                   = 1
	DeleteTypeWords                      = 2
	DeleteTypeDisplayLines               = 3
	DeleteTypeDisplayLineEnds            = 4
	DeleteTypeParagraphEnds              = 5
	DeleteTypeParagraphs                 = 6
	DeleteTypeWhitespace                 = 7
)

type DirectionType int

const (
	DirectionTypeTabForward  DirectionType = 0
	DirectionTypeTabBackward               = 1
	DirectionTypeUp                        = 2
	DirectionTypeDown                      = 3
	DirectionTypeLeft                      = 4
	DirectionTypeRight                     = 5
)

type DragResult int

const (
	DragResultSuccess        DragResult = 0
	DragResultNoTarget                  = 1
	DragResultUserCancelled             = 2
	DragResultTimeoutExpired            = 3
	DragResultGrabBroken                = 4
	DragResultError                     = 5
)

type EntryIconPosition int

const (
	EntryIconPositionPrimary   EntryIconPosition = 0
	EntryIconPositionSecondary                   = 1
)

type EventSequenceState int

const (
	EventSequenceStateNone    EventSequenceState = 0
	EventSequenceStateClaimed                    = 1
	EventSequenceStateDenied                     = 2
)

type ExpanderStyle int

const (
	ExpanderStyleCollapsed     ExpanderStyle = 0
	ExpanderStyleSemiCollapsed               = 1
	ExpanderStyleSemiExpanded                = 2
	ExpanderStyleExpanded                    = 3
)

type FileChooserAction int

const (
	FileChooserActionOpen         FileChooserAction = 0
	FileChooserActionSave                           = 1
	FileChooserActionSelectFolder                   = 2
	FileChooserActionCreateFolder                   = 3
)

type FileChooserConfirmation int

const (
	FileChooserConfirmationConfirm        FileChooserConfirmation = 0
	FileChooserConfirmationAcceptFilename                         = 1
	FileChooserConfirmationSelectAgain                            = 2
)

type FileChooserError int

const (
	FileChooserErrorNonexistent        FileChooserError = 0
	FileChooserErrorBadFilename                         = 1
	FileChooserErrorAlreadyExists                       = 2
	FileChooserErrorIncompleteHostname                  = 3
)

type IMPreeditStyle int

const (
	IMPreeditStyleNothing  IMPreeditStyle = 0
	IMPreeditStyleCallback                = 1
	IMPreeditStyleNone                    = 2
)

type IMStatusStyle int

const (
	IMStatusStyleNothing  IMStatusStyle = 0
	IMStatusStyleCallback               = 1
	IMStatusStyleNone                   = 2
)

type IconSize int

const (
	IconSizeInvalid      IconSize = 0
	IconSizeMenu                  = 1
	IconSizeSmallToolbar          = 2
	IconSizeLargeToolbar          = 3
	IconSizeButton                = 4
	IconSizeDnd                   = 5
	IconSizeDialog                = 6
)

type IconThemeError int

const (
	IconThemeErrorNotFound IconThemeError = 0
	IconThemeErrorFailed                  = 1
)

type IconViewDropPosition int

const (
	IconViewDropPositionNoDrop    IconViewDropPosition = 0
	IconViewDropPositionDropInto                       = 1
	IconViewDropPositionDropLeft                       = 2
	IconViewDropPositionDropRight                      = 3
	IconViewDropPositionDropAbove                      = 4
	IconViewDropPositionDropBelow                      = 5
)

type ImageType int

const (
	ImageTypeEmpty     ImageType = 0
	ImageTypePixbuf              = 1
	ImageTypeStock               = 2
	ImageTypeIconSet             = 3
	ImageTypeAnimation           = 4
	ImageTypeIconName            = 5
	ImageTypeGicon               = 6
	ImageTypeSurface             = 7
)

type InputPurpose int

const (
	InputPurposeFreeForm InputPurpose = 0
	InputPurposeAlpha                 = 1
	InputPurposeDigits                = 2
	InputPurposeNumber                = 3
	InputPurposePhone                 = 4
	InputPurposeUrl                   = 5
	InputPurposeEmail                 = 6
	InputPurposeName                  = 7
	InputPurposePassword              = 8
	InputPurposePin                   = 9
)

type Justification int

const (
	JustificationLeft   Justification = 0
	JustificationRight                = 1
	JustificationCenter               = 2
	JustificationFill                 = 3
)

type LevelBarMode int

const (
	LevelBarModeContinuous LevelBarMode = 0
	LevelBarModeDiscrete                = 1
)

type License int

const (
	LicenseUnknown    License = 0
	LicenseCustom             = 1
	LicenseGpl20              = 2
	LicenseGpl30              = 3
	LicenseLgpl21             = 4
	LicenseLgpl30             = 5
	LicenseBsd                = 6
	LicenseMitX11             = 7
	LicenseArtistic           = 8
	LicenseGpl20Only          = 9
	LicenseGpl30Only          = 10
	LicenseLgpl21Only         = 11
	LicenseLgpl30Only         = 12
	LicenseAgpl30             = 13
)

type MenuDirectionType int

const (
	MenuDirectionTypeParent MenuDirectionType = 0
	MenuDirectionTypeChild                    = 1
	MenuDirectionTypeNext                     = 2
	MenuDirectionTypePrev                     = 3
)

type MessageType int

const (
	MessageTypeInfo     MessageType = 0
	MessageTypeWarning              = 1
	MessageTypeQuestion             = 2
	MessageTypeError                = 3
	MessageTypeOther                = 4
)

type MovementStep int

const (
	MovementStepLogicalPositions MovementStep = 0
	MovementStepVisualPositions               = 1
	MovementStepWords                         = 2
	MovementStepDisplayLines                  = 3
	MovementStepDisplayLineEnds               = 4
	MovementStepParagraphs                    = 5
	MovementStepParagraphEnds                 = 6
	MovementStepPages                         = 7
	MovementStepBufferEnds                    = 8
	MovementStepHorizontalPages               = 9
)

type NotebookTab int

const (
	NotebookTabFirst NotebookTab = 0
	NotebookTabLast              = 1
)

type NumberUpLayout int

const (
	NumberUpLayoutLrtb NumberUpLayout = 0
	NumberUpLayoutLrbt                = 1
	NumberUpLayoutRltb                = 2
	NumberUpLayoutRlbt                = 3
	NumberUpLayoutTblr                = 4
	NumberUpLayoutTbrl                = 5
	NumberUpLayoutBtlr                = 6
	NumberUpLayoutBtrl                = 7
)

type Orientation int

const (
	OrientationHorizontal Orientation = 0
	OrientationVertical               = 1
)

type PackDirection int

const (
	PackDirectionLtr PackDirection = 0
	PackDirectionRtl               = 1
	PackDirectionTtb               = 2
	PackDirectionBtt               = 3
)

type PackType int

const (
	PackTypeStart PackType = 0
	PackTypeEnd            = 1
)

type PadActionType int

const (
	PadActionTypeButton PadActionType = 0
	PadActionTypeRing                 = 1
	PadActionTypeStrip                = 2
)

type PageOrientation int

const (
	PageOrientationPortrait         PageOrientation = 0
	PageOrientationLandscape                        = 1
	PageOrientationReversePortrait                  = 2
	PageOrientationReverseLandscape                 = 3
)

type PageSet int

const (
	PageSetAll  PageSet = 0
	PageSetEven         = 1
	PageSetOdd          = 2
)

type PanDirection int

const (
	PanDirectionLeft  PanDirection = 0
	PanDirectionRight              = 1
	PanDirectionUp                 = 2
	PanDirectionDown               = 3
)

type PathPriorityType int

const (
	PathPriorityTypeLowest      PathPriorityType = 0
	PathPriorityTypeGtk                          = 4
	PathPriorityTypeApplication                  = 8
	PathPriorityTypeTheme                        = 10
	PathPriorityTypeRc                           = 12
	PathPriorityTypeHighest                      = 15
)

type PathType int

const (
	PathTypeWidget      PathType = 0
	PathTypeWidgetClass          = 1
	PathTypeClass                = 2
)

type PolicyType int

const (
	PolicyTypeAlways    PolicyType = 0
	PolicyTypeAutomatic            = 1
	PolicyTypeNever                = 2
	PolicyTypeExternal             = 3
)

type PopoverConstraint int

const (
	PopoverConstraintNone   PopoverConstraint = 0
	PopoverConstraintWindow                   = 1
)

type PositionType int

const (
	PositionTypeLeft   PositionType = 0
	PositionTypeRight               = 1
	PositionTypeTop                 = 2
	PositionTypeBottom              = 3
)

type PrintDuplex int

const (
	PrintDuplexSimplex    PrintDuplex = 0
	PrintDuplexHorizontal             = 1
	PrintDuplexVertical               = 2
)

type PrintError int

const (
	PrintErrorGeneral       PrintError = 0
	PrintErrorInternalError            = 1
	PrintErrorNomem                    = 2
	PrintErrorInvalidFile              = 3
)

type PrintOperationAction int

const (
	PrintOperationActionPrintDialog PrintOperationAction = 0
	PrintOperationActionPrint                            = 1
	PrintOperationActionPreview                          = 2
	PrintOperationActionExport                           = 3
)

type PrintOperationResult int

const (
	PrintOperationResultError      PrintOperationResult = 0
	PrintOperationResultApply                           = 1
	PrintOperationResultCancel                          = 2
	PrintOperationResultInProgress                      = 3
)

type PrintPages int

const (
	PrintPagesAll       PrintPages = 0
	PrintPagesCurrent              = 1
	PrintPagesRanges               = 2
	PrintPagesSelection            = 3
)

type PrintQuality int

const (
	PrintQualityLow    PrintQuality = 0
	PrintQualityNormal              = 1
	PrintQualityHigh                = 2
	PrintQualityDraft               = 3
)

type PrintStatus int

const (
	PrintStatusInitial         PrintStatus = 0
	PrintStatusPreparing                   = 1
	PrintStatusGeneratingData              = 2
	PrintStatusSendingData                 = 3
	PrintStatusPending                     = 4
	PrintStatusPendingIssue                = 5
	PrintStatusPrinting                    = 6
	PrintStatusFinished                    = 7
	PrintStatusFinishedAborted             = 8
)

type PropagationPhase int

const (
	PropagationPhaseNone    PropagationPhase = 0
	PropagationPhaseCapture                  = 1
	PropagationPhaseBubble                   = 2
	PropagationPhaseTarget                   = 3
)

type RcTokenType int

const (
	RcTokenTypeInvalid      RcTokenType = 270
	RcTokenTypeInclude                  = 271
	RcTokenTypeNormal                   = 272
	RcTokenTypeActive                   = 273
	RcTokenTypePrelight                 = 274
	RcTokenTypeSelected                 = 275
	RcTokenTypeInsensitive              = 276
	RcTokenTypeFg                       = 277
	RcTokenTypeBg                       = 278
	RcTokenTypeText                     = 279
	RcTokenTypeBase                     = 280
	RcTokenTypeXthickness               = 281
	RcTokenTypeYthickness               = 282
	RcTokenTypeFont                     = 283
	RcTokenTypeFontset                  = 284
	RcTokenTypeFontName                 = 285
	RcTokenTypeBgPixmap                 = 286
	RcTokenTypePixmapPath               = 287
	RcTokenTypeStyle                    = 288
	RcTokenTypeBinding                  = 289
	RcTokenTypeBind                     = 290
	RcTokenTypeWidget                   = 291
	RcTokenTypeWidgetClass              = 292
	RcTokenTypeClass                    = 293
	RcTokenTypeLowest                   = 294
	RcTokenTypeGtk                      = 295
	RcTokenTypeApplication              = 296
	RcTokenTypeTheme                    = 297
	RcTokenTypeRc                       = 298
	RcTokenTypeHighest                  = 299
	RcTokenTypeEngine                   = 300
	RcTokenTypeModulePath               = 301
	RcTokenTypeImModulePath             = 302
	RcTokenTypeImModuleFile             = 303
	RcTokenTypeStock                    = 304
	RcTokenTypeLtr                      = 305
	RcTokenTypeRtl                      = 306
	RcTokenTypeColor                    = 307
	RcTokenTypeUnbind                   = 308
	RcTokenTypeLast                     = 309
)

type RecentChooserError int

const (
	RecentChooserErrorNotFound   RecentChooserError = 0
	RecentChooserErrorInvalidUri                    = 1
)

type RecentManagerError int

const (
	RecentManagerErrorNotFound        RecentManagerError = 0
	RecentManagerErrorInvalidUri                         = 1
	RecentManagerErrorInvalidEncoding                    = 2
	RecentManagerErrorNotRegistered                      = 3
	RecentManagerErrorRead                               = 4
	RecentManagerErrorWrite                              = 5
	RecentManagerErrorUnknown                            = 6
)

type RecentSortType int

const (
	RecentSortTypeNone   RecentSortType = 0
	RecentSortTypeMru                   = 1
	RecentSortTypeLru                   = 2
	RecentSortTypeCustom                = 3
)

type ReliefStyle int

const (
	ReliefStyleNormal ReliefStyle = 0
	ReliefStyleHalf               = 1
	ReliefStyleNone               = 2
)

type ResizeMode int

const (
	ResizeModeParent    ResizeMode = 0
	ResizeModeQueue                = 1
	ResizeModeImmediate            = 2
)

type ResponseType int

const (
	ResponseTypeNone        ResponseType = -1
	ResponseTypeReject                   = -2
	ResponseTypeAccept                   = -3
	ResponseTypeDeleteEvent              = -4
	ResponseTypeOk                       = -5
	ResponseTypeCancel                   = -6
	ResponseTypeClose                    = -7
	ResponseTypeYes                      = -8
	ResponseTypeNo                       = -9
	ResponseTypeApply                    = -10
	ResponseTypeHelp                     = -11
)

type RevealerTransitionType int

const (
	RevealerTransitionTypeNone       RevealerTransitionType = 0
	RevealerTransitionTypeCrossfade                         = 1
	RevealerTransitionTypeSlideRight                        = 2
	RevealerTransitionTypeSlideLeft                         = 3
	RevealerTransitionTypeSlideUp                           = 4
	RevealerTransitionTypeSlideDown                         = 5
)

type ScrollStep int

const (
	ScrollStepSteps           ScrollStep = 0
	ScrollStepPages                      = 1
	ScrollStepEnds                       = 2
	ScrollStepHorizontalSteps            = 3
	ScrollStepHorizontalPages            = 4
	ScrollStepHorizontalEnds             = 5
)

type ScrollType int

const (
	ScrollTypeNone         ScrollType = 0
	ScrollTypeJump                    = 1
	ScrollTypeStepBackward            = 2
	ScrollTypeStepForward             = 3
	ScrollTypePageBackward            = 4
	ScrollTypePageForward             = 5
	ScrollTypeStepUp                  = 6
	ScrollTypeStepDown                = 7
	ScrollTypePageUp                  = 8
	ScrollTypePageDown                = 9
	ScrollTypeStepLeft                = 10
	ScrollTypeStepRight               = 11
	ScrollTypePageLeft                = 12
	ScrollTypePageRight               = 13
	ScrollTypeStart                   = 14
	ScrollTypeEnd                     = 15
)

type ScrollablePolicy int

const (
	ScrollablePolicyMinimum ScrollablePolicy = 0
	ScrollablePolicyNatural                  = 1
)

type SelectionMode int

const (
	SelectionModeNone     SelectionMode = 0
	SelectionModeSingle                 = 1
	SelectionModeBrowse                 = 2
	SelectionModeMultiple               = 3
)

type SensitivityType int

const (
	SensitivityTypeAuto SensitivityType = 0
	SensitivityTypeOn                   = 1
	SensitivityTypeOff                  = 2
)

type ShadowType int

const (
	ShadowTypeNone      ShadowType = 0
	ShadowTypeIn                   = 1
	ShadowTypeOut                  = 2
	ShadowTypeEtchedIn             = 3
	ShadowTypeEtchedOut            = 4
)

type ShortcutType int

const (
	ShortcutTypeAccelerator                   ShortcutType = 0
	ShortcutTypeGesturePinch                               = 1
	ShortcutTypeGestureStretch                             = 2
	ShortcutTypeGestureRotateClockwise                     = 3
	ShortcutTypeGestureRotateCounterclockwise              = 4
	ShortcutTypeGestureTwoFingerSwipeLeft                  = 5
	ShortcutTypeGestureTwoFingerSwipeRight                 = 6
	ShortcutTypeGesture                                    = 7
)

type SizeGroupMode int

const (
	SizeGroupModeNone       SizeGroupMode = 0
	SizeGroupModeHorizontal               = 1
	SizeGroupModeVertical                 = 2
	SizeGroupModeBoth                     = 3
)

type SizeRequestMode int

const (
	SizeRequestModeHeightForWidth SizeRequestMode = 0
	SizeRequestModeWidthForHeight                 = 1
	SizeRequestModeConstantSize                   = 2
)

type SortType int

const (
	SortTypeAscending  SortType = 0
	SortTypeDescending          = 1
)

type SpinButtonUpdatePolicy int

const (
	SpinButtonUpdatePolicyAlways  SpinButtonUpdatePolicy = 0
	SpinButtonUpdatePolicyIfValid                        = 1
)

type SpinType int

const (
	SpinTypeStepForward  SpinType = 0
	SpinTypeStepBackward          = 1
	SpinTypePageForward           = 2
	SpinTypePageBackward          = 3
	SpinTypeHome                  = 4
	SpinTypeEnd                   = 5
	SpinTypeUserDefined           = 6
)

type StackTransitionType int

const (
	StackTransitionTypeNone           StackTransitionType = 0
	StackTransitionTypeCrossfade                          = 1
	StackTransitionTypeSlideRight                         = 2
	StackTransitionTypeSlideLeft                          = 3
	StackTransitionTypeSlideUp                            = 4
	StackTransitionTypeSlideDown                          = 5
	StackTransitionTypeSlideLeftRight                     = 6
	StackTransitionTypeSlideUpDown                        = 7
	StackTransitionTypeOverUp                             = 8
	StackTransitionTypeOverDown                           = 9
	StackTransitionTypeOverLeft                           = 10
	StackTransitionTypeOverRight                          = 11
	StackTransitionTypeUnderUp                            = 12
	StackTransitionTypeUnderDown                          = 13
	StackTransitionTypeUnderLeft                          = 14
	StackTransitionTypeUnderRight                         = 15
	StackTransitionTypeOverUpDown                         = 16
	StackTransitionTypeOverDownUp                         = 17
	StackTransitionTypeOverLeftRight                      = 18
	StackTransitionTypeOverRightLeft                      = 19
)

type StateType int

const (
	StateTypeNormal       StateType = 0
	StateTypeActive                 = 1
	StateTypePrelight               = 2
	StateTypeSelected               = 3
	StateTypeInsensitive            = 4
	StateTypeInconsistent           = 5
	StateTypeFocused                = 6
)

type TextBufferTargetInfo int

const (
	TextBufferTargetInfoBufferContents TextBufferTargetInfo = -1
	TextBufferTargetInfoRichText                            = -2
	TextBufferTargetInfoText                                = -3
)

type TextDirection int

const (
	TextDirectionNone TextDirection = 0
	TextDirectionLtr                = 1
	TextDirectionRtl                = 2
)

type TextExtendSelection int

const (
	TextExtendSelectionWord TextExtendSelection = 0
	TextExtendSelectionLine                     = 1
)

type TextViewLayer int

const (
	TextViewLayerBelow     TextViewLayer = 0
	TextViewLayerAbove                   = 1
	TextViewLayerBelowText               = 2
	TextViewLayerAboveText               = 3
)

type TextWindowType int

const (
	TextWindowTypePrivate TextWindowType = 0
	TextWindowTypeWidget                 = 1
	TextWindowTypeText                   = 2
	TextWindowTypeLeft                   = 3
	TextWindowTypeRight                  = 4
	TextWindowTypeTop                    = 5
	TextWindowTypeBottom                 = 6
)

type ToolbarSpaceStyle int

const (
	ToolbarSpaceStyleEmpty ToolbarSpaceStyle = 0
	ToolbarSpaceStyleLine                    = 1
)

type ToolbarStyle int

const (
	ToolbarStyleIcons     ToolbarStyle = 0
	ToolbarStyleText                   = 1
	ToolbarStyleBoth                   = 2
	ToolbarStyleBothHoriz              = 3
)

type TreeViewColumnSizing int

const (
	TreeViewColumnSizingGrowOnly TreeViewColumnSizing = 0
	TreeViewColumnSizingAutosize                      = 1
	TreeViewColumnSizingFixed                         = 2
)

type TreeViewDropPosition int

const (
	TreeViewDropPositionBefore       TreeViewDropPosition = 0
	TreeViewDropPositionAfter                             = 1
	TreeViewDropPositionIntoOrBefore                      = 2
	TreeViewDropPositionIntoOrAfter                       = 3
)

type TreeViewGridLines int

const (
	TreeViewGridLinesNone       TreeViewGridLines = 0
	TreeViewGridLinesHorizontal                   = 1
	TreeViewGridLinesVertical                     = 2
	TreeViewGridLinesBoth                         = 3
)

type Unit int

const (
	UnitNone   Unit = 0
	UnitPoints      = 1
	UnitInch        = 2
	UnitMm          = 3
)

type WidgetHelpType int

const (
	WidgetHelpTypeTooltip   WidgetHelpType = 0
	WidgetHelpTypeWhatsThis                = 1
)

type WindowPosition int

const (
	WindowPositionNone           WindowPosition = 0
	WindowPositionCenter                        = 1
	WindowPositionMouse                         = 2
	WindowPositionCenterAlways                  = 3
	WindowPositionCenterOnParent                = 4
)

type WindowType int

const (
	WindowTypeToplevel WindowType = 0
	WindowTypePopup               = 1
)

type WrapMode int

const (
	WrapModeNone     WrapMode = 0
	WrapModeChar              = 1
	WrapModeWord              = 2
	WrapModeWordChar          = 3
)

type AccelFlags int

const (
	AccelFlagsVisible AccelFlags = 1
	AccelFlagsLocked             = 2
	AccelFlagsMask               = 7
)

type ApplicationInhibitFlags int

const (
	ApplicationInhibitFlagsLogout  ApplicationInhibitFlags = 1
	ApplicationInhibitFlagsSwitch                          = 2
	ApplicationInhibitFlagsSuspend                         = 4
	ApplicationInhibitFlagsIdle                            = 8
)

type AttachOptions int

const (
	AttachOptionsExpand AttachOptions = 1
	AttachOptionsShrink               = 2
	AttachOptionsFill                 = 4
)

type CalendarDisplayOptions int

const (
	CalendarDisplayOptionsShowHeading     CalendarDisplayOptions = 1
	CalendarDisplayOptionsShowDayNames                           = 2
	CalendarDisplayOptionsNoMonthChange                          = 4
	CalendarDisplayOptionsShowWeekNumbers                        = 8
	CalendarDisplayOptionsShowDetails                            = 32
)

type CellRendererState int

const (
	CellRendererStateSelected    CellRendererState = 1
	CellRendererStatePrelit                        = 2
	CellRendererStateInsensitive                   = 4
	CellRendererStateSorted                        = 8
	CellRendererStateFocused                       = 16
	CellRendererStateExpandable                    = 32
	CellRendererStateExpanded                      = 64
)

type DebugFlag int

const (
	DebugFlagMisc         DebugFlag = 1
	DebugFlagPlugsocket             = 2
	DebugFlagText                   = 4
	DebugFlagTree                   = 8
	DebugFlagUpdates                = 16
	DebugFlagKeybindings            = 32
	DebugFlagMultihead              = 64
	DebugFlagModules                = 128
	DebugFlagGeometry               = 256
	DebugFlagIcontheme              = 512
	DebugFlagPrinting               = 1024
	DebugFlagBuilder                = 2048
	DebugFlagSizeRequest            = 4096
	DebugFlagNoCssCache             = 8192
	DebugFlagBaselines              = 16384
	DebugFlagPixelCache             = 32768
	DebugFlagNoPixelCache           = 65536
	DebugFlagInteractive            = 131072
	DebugFlagTouchscreen            = 262144
	DebugFlagActions                = 524288
	DebugFlagResize                 = 1048576
	DebugFlagLayout                 = 2097152
)

type DestDefaults int

const (
	DestDefaultsMotion    DestDefaults = 1
	DestDefaultsHighlight              = 2
	DestDefaultsDrop                   = 4
	DestDefaultsAll                    = 7
)

type DialogFlags int

const (
	DialogFlagsModal             DialogFlags = 1
	DialogFlagsDestroyWithParent             = 2
	DialogFlagsUseHeaderBar                  = 4
)

type FileFilterFlags int

const (
	FileFilterFlagsFilename    FileFilterFlags = 1
	FileFilterFlagsUri                         = 2
	FileFilterFlagsDisplayName                 = 4
	FileFilterFlagsMimeType                    = 8
)

type IconLookupFlags int

const (
	IconLookupFlagsNoSvg           IconLookupFlags = 1
	IconLookupFlagsForceSvg                        = 2
	IconLookupFlagsUseBuiltin                      = 4
	IconLookupFlagsGenericFallback                 = 8
	IconLookupFlagsForceSize                       = 16
	IconLookupFlagsForceRegular                    = 32
	IconLookupFlagsForceSymbolic                   = 64
	IconLookupFlagsDirLtr                          = 128
	IconLookupFlagsDirRtl                          = 256
)

type InputHints int

const (
	InputHintsNone               InputHints = 0
	InputHintsSpellcheck                    = 1
	InputHintsNoSpellcheck                  = 2
	InputHintsWordCompletion                = 4
	InputHintsLowercase                     = 8
	InputHintsUppercaseChars                = 16
	InputHintsUppercaseWords                = 32
	InputHintsUppercaseSentences            = 64
	InputHintsInhibitOsk                    = 128
	InputHintsVerticalWriting               = 256
)

type JunctionSides int

const (
	JunctionSidesNone              JunctionSides = 0
	JunctionSidesCornerTopleft                   = 1
	JunctionSidesCornerTopright                  = 2
	JunctionSidesCornerBottomleft                = 4
	JunctionSidesCornerBottomright               = 8
	JunctionSidesTop                             = 3
	JunctionSidesBottom                          = 12
	JunctionSidesLeft                            = 5
	JunctionSidesRight                           = 10
)

type PlacesOpenFlags int

const (
	PlacesOpenFlagsNormal    PlacesOpenFlags = 1
	PlacesOpenFlagsNewTab                    = 2
	PlacesOpenFlagsNewWindow                 = 4
)

type RcFlags int

const (
	RcFlagsFg   RcFlags = 1
	RcFlagsBg           = 2
	RcFlagsText         = 4
	RcFlagsBase         = 8
)

type RecentFilterFlags int

const (
	RecentFilterFlagsUri         RecentFilterFlags = 1
	RecentFilterFlagsDisplayName                   = 2
	RecentFilterFlagsMimeType                      = 4
	RecentFilterFlagsApplication                   = 8
	RecentFilterFlagsGroup                         = 16
	RecentFilterFlagsAge                           = 32
)

type RegionFlags int

const (
	RegionFlagsEven   RegionFlags = 1
	RegionFlagsOdd                = 2
	RegionFlagsFirst              = 4
	RegionFlagsLast               = 8
	RegionFlagsOnly               = 16
	RegionFlagsSorted             = 32
)

type StateFlags int

const (
	StateFlagsNormal       StateFlags = 0
	StateFlagsActive                  = 1
	StateFlagsPrelight                = 2
	StateFlagsSelected                = 4
	StateFlagsInsensitive             = 8
	StateFlagsInconsistent            = 16
	StateFlagsFocused                 = 32
	StateFlagsBackdrop                = 64
	StateFlagsDirLtr                  = 128
	StateFlagsDirRtl                  = 256
	StateFlagsLink                    = 512
	StateFlagsVisited                 = 1024
	StateFlagsChecked                 = 2048
	StateFlagsDropActive              = 4096
)

type StyleContextPrintFlags int

const (
	StyleContextPrintFlagsNone      StyleContextPrintFlags = 0
	StyleContextPrintFlagsRecurse                          = 1
	StyleContextPrintFlagsShowStyle                        = 2
)

type TargetFlags int

const (
	TargetFlagsSameApp     TargetFlags = 1
	TargetFlagsSameWidget              = 2
	TargetFlagsOtherApp                = 4
	TargetFlagsOtherWidget             = 8
)

type TextSearchFlags int

const (
	TextSearchFlagsVisibleOnly     TextSearchFlags = 1
	TextSearchFlagsTextOnly                        = 2
	TextSearchFlagsCaseInsensitive                 = 4
)

type ToolPaletteDragTargets int

const (
	ToolPaletteDragTargetsItems  ToolPaletteDragTargets = 1
	ToolPaletteDragTargetsGroups                        = 2
)

type TreeModelFlags int

const (
	TreeModelFlagsItersPersist TreeModelFlags = 1
	TreeModelFlagsListOnly                    = 2
)

type UIManagerItemType int

const (
	UIManagerItemTypeAuto            UIManagerItemType = 0
	UIManagerItemTypeMenubar                           = 1
	UIManagerItemTypeMenu                              = 2
	UIManagerItemTypeToolbar                           = 4
	UIManagerItemTypePlaceholder                       = 8
	UIManagerItemTypePopup                             = 16
	UIManagerItemTypeMenuitem                          = 32
	UIManagerItemTypeToolitem                          = 64
	UIManagerItemTypeSeparator                         = 128
	UIManagerItemTypeAccelerator                       = 256
	UIManagerItemTypePopupWithAccels                   = 512
)
