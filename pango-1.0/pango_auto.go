package pango

/*
#cgo pkg-config: pango
#include <pango/pango.h>
#include <stdlib.h>
*/
import "C"
import "github.com/linuxdeepin/go-gir/gobject-2.0"
import "github.com/linuxdeepin/go-gir/util"
import "unsafe"

// Struct GlyphItem
type GlyphItem struct {
	Ptr unsafe.Pointer
}

func (v GlyphItem) native() *C.PangoGlyphItem {
	return (*C.PangoGlyphItem)(v.Ptr)
}
func wrapGlyphItem(p *C.PangoGlyphItem) GlyphItem {
	return GlyphItem{unsafe.Pointer(p)}
}
func WrapGlyphItem(p unsafe.Pointer) GlyphItem {
	return GlyphItem{p}
}
func (v GlyphItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGlyphItem(p unsafe.Pointer) interface{} {
	return WrapGlyphItem(p)
}

// Copy is a wrapper around pango_glyph_item_copy().
func (orig GlyphItem) Copy() GlyphItem {
	ret0 := C.pango_glyph_item_copy(orig.native())
	return wrapGlyphItem(ret0)
}

// Free is a wrapper around pango_glyph_item_free().
func (glyph_item GlyphItem) Free() {
	C.pango_glyph_item_free(glyph_item.native())
}

// GetLogicalWidths is a wrapper around pango_glyph_item_get_logical_widths().
func (glyph_item GlyphItem) GetLogicalWidths(text string, logical_widths []int) {
	text0 := C.CString(text)
	logical_widths0 := make([]C.int, len(logical_widths))
	for idx, elemG := range logical_widths {
		logical_widths0[idx] = C.int(elemG)
	}
	var logical_widths0Ptr *C.int
	if len(logical_widths0) > 0 {
		logical_widths0Ptr = &logical_widths0[0]
	}
	C.pango_glyph_item_get_logical_widths(glyph_item.native(), text0, logical_widths0Ptr)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// Split is a wrapper around pango_glyph_item_split().
func (orig GlyphItem) Split(text string, split_index int) GlyphItem {
	text0 := C.CString(text)
	ret0 := C.pango_glyph_item_split(orig.native(), text0, C.int(split_index))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
	return wrapGlyphItem(ret0)
}

// Struct Analysis
type Analysis struct {
	Ptr unsafe.Pointer
}

func (v Analysis) native() *C.PangoAnalysis {
	return (*C.PangoAnalysis)(v.Ptr)
}
func wrapAnalysis(p *C.PangoAnalysis) Analysis {
	return Analysis{unsafe.Pointer(p)}
}
func WrapAnalysis(p unsafe.Pointer) Analysis {
	return Analysis{p}
}
func (v Analysis) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAnalysis(p unsafe.Pointer) interface{} {
	return WrapAnalysis(p)
}

// Struct AttrClass
type AttrClass struct {
	Ptr unsafe.Pointer
}

func (v AttrClass) native() *C.PangoAttrClass {
	return (*C.PangoAttrClass)(v.Ptr)
}
func wrapAttrClass(p *C.PangoAttrClass) AttrClass {
	return AttrClass{unsafe.Pointer(p)}
}
func WrapAttrClass(p unsafe.Pointer) AttrClass {
	return AttrClass{p}
}
func (v AttrClass) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrClass(p unsafe.Pointer) interface{} {
	return WrapAttrClass(p)
}

// Struct AttrColor
type AttrColor struct {
	Ptr unsafe.Pointer
}

func (v AttrColor) native() *C.PangoAttrColor {
	return (*C.PangoAttrColor)(v.Ptr)
}
func wrapAttrColor(p *C.PangoAttrColor) AttrColor {
	return AttrColor{unsafe.Pointer(p)}
}
func WrapAttrColor(p unsafe.Pointer) AttrColor {
	return AttrColor{p}
}
func (v AttrColor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrColor(p unsafe.Pointer) interface{} {
	return WrapAttrColor(p)
}

// Struct AttrFloat
type AttrFloat struct {
	Ptr unsafe.Pointer
}

func (v AttrFloat) native() *C.PangoAttrFloat {
	return (*C.PangoAttrFloat)(v.Ptr)
}
func wrapAttrFloat(p *C.PangoAttrFloat) AttrFloat {
	return AttrFloat{unsafe.Pointer(p)}
}
func WrapAttrFloat(p unsafe.Pointer) AttrFloat {
	return AttrFloat{p}
}
func (v AttrFloat) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrFloat(p unsafe.Pointer) interface{} {
	return WrapAttrFloat(p)
}

// Struct AttrInt
type AttrInt struct {
	Ptr unsafe.Pointer
}

func (v AttrInt) native() *C.PangoAttrInt {
	return (*C.PangoAttrInt)(v.Ptr)
}
func wrapAttrInt(p *C.PangoAttrInt) AttrInt {
	return AttrInt{unsafe.Pointer(p)}
}
func WrapAttrInt(p unsafe.Pointer) AttrInt {
	return AttrInt{p}
}
func (v AttrInt) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrInt(p unsafe.Pointer) interface{} {
	return WrapAttrInt(p)
}

// Struct AttrList
type AttrList struct {
	Ptr unsafe.Pointer
}

func (v AttrList) native() *C.PangoAttrList {
	return (*C.PangoAttrList)(v.Ptr)
}
func wrapAttrList(p *C.PangoAttrList) AttrList {
	return AttrList{unsafe.Pointer(p)}
}
func WrapAttrList(p unsafe.Pointer) AttrList {
	return AttrList{p}
}
func (v AttrList) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrList(p unsafe.Pointer) interface{} {
	return WrapAttrList(p)
}

// AttrListNew is a wrapper around pango_attr_list_new().
func AttrListNew() AttrList {
	ret0 := C.pango_attr_list_new()
	return wrapAttrList(ret0)
}

// Change is a wrapper around pango_attr_list_change().
func (list AttrList) Change(attr Attribute) {
	C.pango_attr_list_change(list.native(), attr.native())
}

// Copy is a wrapper around pango_attr_list_copy().
func (list AttrList) Copy() AttrList {
	ret0 := C.pango_attr_list_copy(list.native())
	return wrapAttrList(ret0)
}

// GetIterator is a wrapper around pango_attr_list_get_iterator().
func (list AttrList) GetIterator() AttrIterator {
	ret0 := C.pango_attr_list_get_iterator(list.native())
	return wrapAttrIterator(ret0)
}

// Insert is a wrapper around pango_attr_list_insert().
func (list AttrList) Insert(attr Attribute) {
	C.pango_attr_list_insert(list.native(), attr.native())
}

// InsertBefore is a wrapper around pango_attr_list_insert_before().
func (list AttrList) InsertBefore(attr Attribute) {
	C.pango_attr_list_insert_before(list.native(), attr.native())
}

// Ref is a wrapper around pango_attr_list_ref().
func (list AttrList) Ref() AttrList {
	ret0 := C.pango_attr_list_ref(list.native())
	return wrapAttrList(ret0)
}

// Splice is a wrapper around pango_attr_list_splice().
func (list AttrList) Splice(other AttrList, pos int, len_ int) {
	C.pango_attr_list_splice(list.native(), other.native(), C.gint(pos), C.gint(len_))
}

// Unref is a wrapper around pango_attr_list_unref().
func (list AttrList) Unref() {
	C.pango_attr_list_unref(list.native())
}

// Struct AttrString
type AttrString struct {
	Ptr unsafe.Pointer
}

func (v AttrString) native() *C.PangoAttrString {
	return (*C.PangoAttrString)(v.Ptr)
}
func wrapAttrString(p *C.PangoAttrString) AttrString {
	return AttrString{unsafe.Pointer(p)}
}
func WrapAttrString(p unsafe.Pointer) AttrString {
	return AttrString{p}
}
func (v AttrString) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrString(p unsafe.Pointer) interface{} {
	return WrapAttrString(p)
}

// Struct Attribute
type Attribute struct {
	Ptr unsafe.Pointer
}

func (v Attribute) native() *C.PangoAttribute {
	return (*C.PangoAttribute)(v.Ptr)
}
func wrapAttribute(p *C.PangoAttribute) Attribute {
	return Attribute{unsafe.Pointer(p)}
}
func WrapAttribute(p unsafe.Pointer) Attribute {
	return Attribute{p}
}
func (v Attribute) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttribute(p unsafe.Pointer) interface{} {
	return WrapAttribute(p)
}

// Copy is a wrapper around pango_attribute_copy().
func (attr Attribute) Copy() Attribute {
	ret0 := C.pango_attribute_copy(attr.native())
	return wrapAttribute(ret0)
}

// Destroy is a wrapper around pango_attribute_destroy().
func (attr Attribute) Destroy() {
	C.pango_attribute_destroy(attr.native())
}

// Equal is a wrapper around pango_attribute_equal().
func (attr1 Attribute) Equal(attr2 Attribute) bool {
	ret0 := C.pango_attribute_equal(attr1.native(), attr2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Init is a wrapper around pango_attribute_init().
func (attr Attribute) Init(klass AttrClass) {
	C.pango_attribute_init(attr.native(), klass.native())
}

// Struct Color
type Color struct {
	Ptr unsafe.Pointer
}

func (v Color) native() *C.PangoColor {
	return (*C.PangoColor)(v.Ptr)
}
func wrapColor(p *C.PangoColor) Color {
	return Color{unsafe.Pointer(p)}
}
func WrapColor(p unsafe.Pointer) Color {
	return Color{p}
}
func (v Color) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColor(p unsafe.Pointer) interface{} {
	return WrapColor(p)
}

// Copy is a wrapper around pango_color_copy().
func (src Color) Copy() Color {
	ret0 := C.pango_color_copy(src.native())
	return wrapColor(ret0)
}

// Free is a wrapper around pango_color_free().
func (color Color) Free() {
	C.pango_color_free(color.native())
}

// Parse is a wrapper around pango_color_parse().
func (color Color) Parse(spec string) bool {
	spec0 := C.CString(spec)
	ret0 := C.pango_color_parse(color.native(), spec0)
	C.free(unsafe.Pointer(spec0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ToString is a wrapper around pango_color_to_string().
func (color Color) ToString() string {
	ret0 := C.pango_color_to_string(color.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Struct GlyphGeometry
type GlyphGeometry struct {
	Ptr unsafe.Pointer
}

func (v GlyphGeometry) native() *C.PangoGlyphGeometry {
	return (*C.PangoGlyphGeometry)(v.Ptr)
}
func wrapGlyphGeometry(p *C.PangoGlyphGeometry) GlyphGeometry {
	return GlyphGeometry{unsafe.Pointer(p)}
}
func WrapGlyphGeometry(p unsafe.Pointer) GlyphGeometry {
	return GlyphGeometry{p}
}
func (v GlyphGeometry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGlyphGeometry(p unsafe.Pointer) interface{} {
	return WrapGlyphGeometry(p)
}

// Struct GlyphInfo
type GlyphInfo struct {
	Ptr unsafe.Pointer
}

func (v GlyphInfo) native() *C.PangoGlyphInfo {
	return (*C.PangoGlyphInfo)(v.Ptr)
}
func wrapGlyphInfo(p *C.PangoGlyphInfo) GlyphInfo {
	return GlyphInfo{unsafe.Pointer(p)}
}
func WrapGlyphInfo(p unsafe.Pointer) GlyphInfo {
	return GlyphInfo{p}
}
func (v GlyphInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGlyphInfo(p unsafe.Pointer) interface{} {
	return WrapGlyphInfo(p)
}

// Struct GlyphVisAttr
type GlyphVisAttr struct {
	Ptr unsafe.Pointer
}

func (v GlyphVisAttr) native() *C.PangoGlyphVisAttr {
	return (*C.PangoGlyphVisAttr)(v.Ptr)
}
func wrapGlyphVisAttr(p *C.PangoGlyphVisAttr) GlyphVisAttr {
	return GlyphVisAttr{unsafe.Pointer(p)}
}
func WrapGlyphVisAttr(p unsafe.Pointer) GlyphVisAttr {
	return GlyphVisAttr{p}
}
func (v GlyphVisAttr) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGlyphVisAttr(p unsafe.Pointer) interface{} {
	return WrapGlyphVisAttr(p)
}

// Struct LogAttr
type LogAttr struct {
	Ptr unsafe.Pointer
}

func (v LogAttr) native() *C.PangoLogAttr {
	return (*C.PangoLogAttr)(v.Ptr)
}
func wrapLogAttr(p *C.PangoLogAttr) LogAttr {
	return LogAttr{unsafe.Pointer(p)}
}
func WrapLogAttr(p unsafe.Pointer) LogAttr {
	return LogAttr{p}
}
func (v LogAttr) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLogAttr(p unsafe.Pointer) interface{} {
	return WrapLogAttr(p)
}

// Struct Rectangle
type Rectangle struct {
	Ptr unsafe.Pointer
}

func (v Rectangle) native() *C.PangoRectangle {
	return (*C.PangoRectangle)(v.Ptr)
}
func wrapRectangle(p *C.PangoRectangle) Rectangle {
	return Rectangle{unsafe.Pointer(p)}
}
func WrapRectangle(p unsafe.Pointer) Rectangle {
	return Rectangle{p}
}
func (v Rectangle) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRectangle(p unsafe.Pointer) interface{} {
	return WrapRectangle(p)
}

// Struct AttrFontDesc
type AttrFontDesc struct {
	Ptr unsafe.Pointer
}

func (v AttrFontDesc) native() *C.PangoAttrFontDesc {
	return (*C.PangoAttrFontDesc)(v.Ptr)
}
func wrapAttrFontDesc(p *C.PangoAttrFontDesc) AttrFontDesc {
	return AttrFontDesc{unsafe.Pointer(p)}
}
func WrapAttrFontDesc(p unsafe.Pointer) AttrFontDesc {
	return AttrFontDesc{p}
}
func (v AttrFontDesc) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrFontDesc(p unsafe.Pointer) interface{} {
	return WrapAttrFontDesc(p)
}

// AttrFontDescNew is a wrapper around pango_attr_font_desc_new().
func AttrFontDescNew(desc FontDescription) Attribute {
	ret0 := C.pango_attr_font_desc_new(desc.native())
	return wrapAttribute(ret0)
}

// Struct FontDescription
type FontDescription struct {
	Ptr unsafe.Pointer
}

func (v FontDescription) native() *C.PangoFontDescription {
	return (*C.PangoFontDescription)(v.Ptr)
}
func wrapFontDescription(p *C.PangoFontDescription) FontDescription {
	return FontDescription{unsafe.Pointer(p)}
}
func WrapFontDescription(p unsafe.Pointer) FontDescription {
	return FontDescription{p}
}
func (v FontDescription) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontDescription(p unsafe.Pointer) interface{} {
	return WrapFontDescription(p)
}

// FontDescriptionNew is a wrapper around pango_font_description_new().
func FontDescriptionNew() FontDescription {
	ret0 := C.pango_font_description_new()
	return wrapFontDescription(ret0)
}

// BetterMatch is a wrapper around pango_font_description_better_match().
func (desc FontDescription) BetterMatch(old_match FontDescription, new_match FontDescription) bool {
	ret0 := C.pango_font_description_better_match(desc.native(), old_match.native(), new_match.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Copy is a wrapper around pango_font_description_copy().
func (desc FontDescription) Copy() FontDescription {
	ret0 := C.pango_font_description_copy(desc.native())
	return wrapFontDescription(ret0)
}

// CopyStatic is a wrapper around pango_font_description_copy_static().
func (desc FontDescription) CopyStatic() FontDescription {
	ret0 := C.pango_font_description_copy_static(desc.native())
	return wrapFontDescription(ret0)
}

// Equal is a wrapper around pango_font_description_equal().
func (desc1 FontDescription) Equal(desc2 FontDescription) bool {
	ret0 := C.pango_font_description_equal(desc1.native(), desc2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Free is a wrapper around pango_font_description_free().
func (desc FontDescription) Free() {
	C.pango_font_description_free(desc.native())
}

// GetFamily is a wrapper around pango_font_description_get_family().
func (desc FontDescription) GetFamily() string {
	ret0 := C.pango_font_description_get_family(desc.native())
	ret := C.GoString(ret0)
	return ret
}

// GetGravity is a wrapper around pango_font_description_get_gravity().
func (desc FontDescription) GetGravity() Gravity {
	ret0 := C.pango_font_description_get_gravity(desc.native())
	return Gravity(ret0)
}

// GetSetFields is a wrapper around pango_font_description_get_set_fields().
func (desc FontDescription) GetSetFields() FontMask {
	ret0 := C.pango_font_description_get_set_fields(desc.native())
	return FontMask(ret0)
}

// GetSize is a wrapper around pango_font_description_get_size().
func (desc FontDescription) GetSize() int {
	ret0 := C.pango_font_description_get_size(desc.native())
	return int(ret0)
}

// GetSizeIsAbsolute is a wrapper around pango_font_description_get_size_is_absolute().
func (desc FontDescription) GetSizeIsAbsolute() bool {
	ret0 := C.pango_font_description_get_size_is_absolute(desc.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetStretch is a wrapper around pango_font_description_get_stretch().
func (desc FontDescription) GetStretch() Stretch {
	ret0 := C.pango_font_description_get_stretch(desc.native())
	return Stretch(ret0)
}

// GetStyle is a wrapper around pango_font_description_get_style().
func (desc FontDescription) GetStyle() Style {
	ret0 := C.pango_font_description_get_style(desc.native())
	return Style(ret0)
}

// GetVariant is a wrapper around pango_font_description_get_variant().
func (desc FontDescription) GetVariant() Variant {
	ret0 := C.pango_font_description_get_variant(desc.native())
	return Variant(ret0)
}

// GetWeight is a wrapper around pango_font_description_get_weight().
func (desc FontDescription) GetWeight() Weight {
	ret0 := C.pango_font_description_get_weight(desc.native())
	return Weight(ret0)
}

// Hash is a wrapper around pango_font_description_hash().
func (desc FontDescription) Hash() uint {
	ret0 := C.pango_font_description_hash(desc.native())
	return uint(ret0)
}

// Merge is a wrapper around pango_font_description_merge().
func (desc FontDescription) Merge(desc_to_merge FontDescription, replace_existing bool) {
	C.pango_font_description_merge(desc.native(), desc_to_merge.native(), C.gboolean(util.Bool2Int(replace_existing)) /*go:.util*/)
}

// MergeStatic is a wrapper around pango_font_description_merge_static().
func (desc FontDescription) MergeStatic(desc_to_merge FontDescription, replace_existing bool) {
	C.pango_font_description_merge_static(desc.native(), desc_to_merge.native(), C.gboolean(util.Bool2Int(replace_existing)) /*go:.util*/)
}

// SetAbsoluteSize is a wrapper around pango_font_description_set_absolute_size().
func (desc FontDescription) SetAbsoluteSize(size float64) {
	C.pango_font_description_set_absolute_size(desc.native(), C.double(size))
}

// SetFamily is a wrapper around pango_font_description_set_family().
func (desc FontDescription) SetFamily(family string) {
	family0 := C.CString(family)
	C.pango_font_description_set_family(desc.native(), family0)
	C.free(unsafe.Pointer(family0)) /*ch:<stdlib.h>*/
}

// SetFamilyStatic is a wrapper around pango_font_description_set_family_static().
func (desc FontDescription) SetFamilyStatic(family string) {
	family0 := C.CString(family)
	C.pango_font_description_set_family_static(desc.native(), family0)
	C.free(unsafe.Pointer(family0)) /*ch:<stdlib.h>*/
}

// SetGravity is a wrapper around pango_font_description_set_gravity().
func (desc FontDescription) SetGravity(gravity Gravity) {
	C.pango_font_description_set_gravity(desc.native(), C.PangoGravity(gravity))
}

// SetSize is a wrapper around pango_font_description_set_size().
func (desc FontDescription) SetSize(size int) {
	C.pango_font_description_set_size(desc.native(), C.gint(size))
}

// SetStretch is a wrapper around pango_font_description_set_stretch().
func (desc FontDescription) SetStretch(stretch Stretch) {
	C.pango_font_description_set_stretch(desc.native(), C.PangoStretch(stretch))
}

// SetStyle is a wrapper around pango_font_description_set_style().
func (desc FontDescription) SetStyle(style Style) {
	C.pango_font_description_set_style(desc.native(), C.PangoStyle(style))
}

// SetVariant is a wrapper around pango_font_description_set_variant().
func (desc FontDescription) SetVariant(variant Variant) {
	C.pango_font_description_set_variant(desc.native(), C.PangoVariant(variant))
}

// SetWeight is a wrapper around pango_font_description_set_weight().
func (desc FontDescription) SetWeight(weight Weight) {
	C.pango_font_description_set_weight(desc.native(), C.PangoWeight(weight))
}

// ToFilename is a wrapper around pango_font_description_to_filename().
func (desc FontDescription) ToFilename() string {
	ret0 := C.pango_font_description_to_filename(desc.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// ToString is a wrapper around pango_font_description_to_string().
func (desc FontDescription) ToString() string {
	ret0 := C.pango_font_description_to_string(desc.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// UnsetFields is a wrapper around pango_font_description_unset_fields().
func (desc FontDescription) UnsetFields(to_unset FontMask) {
	C.pango_font_description_unset_fields(desc.native(), C.PangoFontMask(to_unset))
}

// FontDescriptionFromString is a wrapper around pango_font_description_from_string().
func FontDescriptionFromString(str string) FontDescription {
	str0 := C.CString(str)
	ret0 := C.pango_font_description_from_string(str0)
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
	return wrapFontDescription(ret0)
}

// Struct AttrFontFeatures
type AttrFontFeatures struct {
	Ptr unsafe.Pointer
}

func (v AttrFontFeatures) native() *C.PangoAttrFontFeatures {
	return (*C.PangoAttrFontFeatures)(v.Ptr)
}
func wrapAttrFontFeatures(p *C.PangoAttrFontFeatures) AttrFontFeatures {
	return AttrFontFeatures{unsafe.Pointer(p)}
}
func WrapAttrFontFeatures(p unsafe.Pointer) AttrFontFeatures {
	return AttrFontFeatures{p}
}
func (v AttrFontFeatures) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrFontFeatures(p unsafe.Pointer) interface{} {
	return WrapAttrFontFeatures(p)
}

// AttrFontFeaturesNew is a wrapper around pango_attr_font_features_new().
func AttrFontFeaturesNew(features string) Attribute {
	features0 := (*C.gchar)(C.CString(features))
	ret0 := C.pango_attr_font_features_new(features0)
	C.free(unsafe.Pointer(features0)) /*ch:<stdlib.h>*/
	return wrapAttribute(ret0)
}

// Struct AttrLanguage
type AttrLanguage struct {
	Ptr unsafe.Pointer
}

func (v AttrLanguage) native() *C.PangoAttrLanguage {
	return (*C.PangoAttrLanguage)(v.Ptr)
}
func wrapAttrLanguage(p *C.PangoAttrLanguage) AttrLanguage {
	return AttrLanguage{unsafe.Pointer(p)}
}
func WrapAttrLanguage(p unsafe.Pointer) AttrLanguage {
	return AttrLanguage{p}
}
func (v AttrLanguage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrLanguage(p unsafe.Pointer) interface{} {
	return WrapAttrLanguage(p)
}

// AttrLanguageNew is a wrapper around pango_attr_language_new().
func AttrLanguageNew(language Language) Attribute {
	ret0 := C.pango_attr_language_new(language.native())
	return wrapAttribute(ret0)
}

// Struct Language
type Language struct {
	Ptr unsafe.Pointer
}

func (v Language) native() *C.PangoLanguage {
	return (*C.PangoLanguage)(v.Ptr)
}
func wrapLanguage(p *C.PangoLanguage) Language {
	return Language{unsafe.Pointer(p)}
}
func WrapLanguage(p unsafe.Pointer) Language {
	return Language{p}
}
func (v Language) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLanguage(p unsafe.Pointer) interface{} {
	return WrapLanguage(p)
}

// GetSampleString is a wrapper around pango_language_get_sample_string().
func (language Language) GetSampleString() string {
	ret0 := C.pango_language_get_sample_string(language.native())
	ret := C.GoString(ret0)
	return ret
}

// GetScripts is a wrapper around pango_language_get_scripts().
func (language Language) GetScripts() []Script {
	var num_scripts0 C.int
	ret0 := C.pango_language_get_scripts(language.native(), &num_scripts0)
	var ret0Slice []C.PangoScript
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(num_scripts0)) /*go:.util*/
	ret := make([]Script, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = Script(elem)
	}
	return ret
}

// IncludesScript is a wrapper around pango_language_includes_script().
func (language Language) IncludesScript(script Script) bool {
	ret0 := C.pango_language_includes_script(language.native(), C.PangoScript(script))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Matches is a wrapper around pango_language_matches().
func (language Language) Matches(range_list string) bool {
	range_list0 := C.CString(range_list)
	ret0 := C.pango_language_matches(language.native(), range_list0)
	C.free(unsafe.Pointer(range_list0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// ToString is a wrapper around pango_language_to_string().
func (language Language) ToString() string {
	ret0 := C.pango_language_to_string(language.native())
	ret := C.GoString(ret0)
	return ret
}

// LanguageFromString is a wrapper around pango_language_from_string().
func LanguageFromString(language string) Language {
	language0 := C.CString(language)
	ret0 := C.pango_language_from_string(language0)
	C.free(unsafe.Pointer(language0)) /*ch:<stdlib.h>*/
	return wrapLanguage(ret0)
}

// LanguageGetDefault is a wrapper around pango_language_get_default().
func LanguageGetDefault() Language {
	ret0 := C.pango_language_get_default()
	return wrapLanguage(ret0)
}

// Struct AttrIterator
type AttrIterator struct {
	Ptr unsafe.Pointer
}

func (v AttrIterator) native() *C.PangoAttrIterator {
	return (*C.PangoAttrIterator)(v.Ptr)
}
func wrapAttrIterator(p *C.PangoAttrIterator) AttrIterator {
	return AttrIterator{unsafe.Pointer(p)}
}
func WrapAttrIterator(p unsafe.Pointer) AttrIterator {
	return AttrIterator{p}
}
func (v AttrIterator) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrIterator(p unsafe.Pointer) interface{} {
	return WrapAttrIterator(p)
}

// Copy is a wrapper around pango_attr_iterator_copy().
func (iterator AttrIterator) Copy() AttrIterator {
	ret0 := C.pango_attr_iterator_copy(iterator.native())
	return wrapAttrIterator(ret0)
}

// Destroy is a wrapper around pango_attr_iterator_destroy().
func (iterator AttrIterator) Destroy() {
	C.pango_attr_iterator_destroy(iterator.native())
}

// Get is a wrapper around pango_attr_iterator_get().
func (iterator AttrIterator) Get(type_ AttrType) Attribute {
	ret0 := C.pango_attr_iterator_get(iterator.native(), C.PangoAttrType(type_))
	return wrapAttribute(ret0)
}

// Next is a wrapper around pango_attr_iterator_next().
func (iterator AttrIterator) Next() bool {
	ret0 := C.pango_attr_iterator_next(iterator.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Range is a wrapper around pango_attr_iterator_range().
func (iterator AttrIterator) Range() (int, int) {
	var start0 C.gint
	var end0 C.gint
	C.pango_attr_iterator_range(iterator.native(), &start0, &end0)
	return int(start0), int(end0)
}

// Struct AttrShape
type AttrShape struct {
	Ptr unsafe.Pointer
}

func (v AttrShape) native() *C.PangoAttrShape {
	return (*C.PangoAttrShape)(v.Ptr)
}
func wrapAttrShape(p *C.PangoAttrShape) AttrShape {
	return AttrShape{unsafe.Pointer(p)}
}
func WrapAttrShape(p unsafe.Pointer) AttrShape {
	return AttrShape{p}
}
func (v AttrShape) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrShape(p unsafe.Pointer) interface{} {
	return WrapAttrShape(p)
}

// AttrShapeNew is a wrapper around pango_attr_shape_new().
func AttrShapeNew(ink_rect Rectangle, logical_rect Rectangle) Attribute {
	ret0 := C.pango_attr_shape_new(ink_rect.native(), logical_rect.native())
	return wrapAttribute(ret0)
}

// Struct AttrSize
type AttrSize struct {
	Ptr unsafe.Pointer
}

func (v AttrSize) native() *C.PangoAttrSize {
	return (*C.PangoAttrSize)(v.Ptr)
}
func wrapAttrSize(p *C.PangoAttrSize) AttrSize {
	return AttrSize{unsafe.Pointer(p)}
}
func WrapAttrSize(p unsafe.Pointer) AttrSize {
	return AttrSize{p}
}
func (v AttrSize) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttrSize(p unsafe.Pointer) interface{} {
	return WrapAttrSize(p)
}

// AttrSizeNew is a wrapper around pango_attr_size_new().
func AttrSizeNew(size int) Attribute {
	ret0 := C.pango_attr_size_new(C.int(size))
	return wrapAttribute(ret0)
}

// AttrSizeNewAbsolute is a wrapper around pango_attr_size_new_absolute().
func AttrSizeNewAbsolute(size int) Attribute {
	ret0 := C.pango_attr_size_new_absolute(C.int(size))
	return wrapAttribute(ret0)
}

// Struct FontMetrics
type FontMetrics struct {
	Ptr unsafe.Pointer
}

func (v FontMetrics) native() *C.PangoFontMetrics {
	return (*C.PangoFontMetrics)(v.Ptr)
}
func wrapFontMetrics(p *C.PangoFontMetrics) FontMetrics {
	return FontMetrics{unsafe.Pointer(p)}
}
func WrapFontMetrics(p unsafe.Pointer) FontMetrics {
	return FontMetrics{p}
}
func (v FontMetrics) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontMetrics(p unsafe.Pointer) interface{} {
	return WrapFontMetrics(p)
}

// GetApproximateCharWidth is a wrapper around pango_font_metrics_get_approximate_char_width().
func (metrics FontMetrics) GetApproximateCharWidth() int {
	ret0 := C.pango_font_metrics_get_approximate_char_width(metrics.native())
	return int(ret0)
}

// GetApproximateDigitWidth is a wrapper around pango_font_metrics_get_approximate_digit_width().
func (metrics FontMetrics) GetApproximateDigitWidth() int {
	ret0 := C.pango_font_metrics_get_approximate_digit_width(metrics.native())
	return int(ret0)
}

// GetAscent is a wrapper around pango_font_metrics_get_ascent().
func (metrics FontMetrics) GetAscent() int {
	ret0 := C.pango_font_metrics_get_ascent(metrics.native())
	return int(ret0)
}

// GetDescent is a wrapper around pango_font_metrics_get_descent().
func (metrics FontMetrics) GetDescent() int {
	ret0 := C.pango_font_metrics_get_descent(metrics.native())
	return int(ret0)
}

// GetStrikethroughPosition is a wrapper around pango_font_metrics_get_strikethrough_position().
func (metrics FontMetrics) GetStrikethroughPosition() int {
	ret0 := C.pango_font_metrics_get_strikethrough_position(metrics.native())
	return int(ret0)
}

// GetStrikethroughThickness is a wrapper around pango_font_metrics_get_strikethrough_thickness().
func (metrics FontMetrics) GetStrikethroughThickness() int {
	ret0 := C.pango_font_metrics_get_strikethrough_thickness(metrics.native())
	return int(ret0)
}

// GetUnderlinePosition is a wrapper around pango_font_metrics_get_underline_position().
func (metrics FontMetrics) GetUnderlinePosition() int {
	ret0 := C.pango_font_metrics_get_underline_position(metrics.native())
	return int(ret0)
}

// GetUnderlineThickness is a wrapper around pango_font_metrics_get_underline_thickness().
func (metrics FontMetrics) GetUnderlineThickness() int {
	ret0 := C.pango_font_metrics_get_underline_thickness(metrics.native())
	return int(ret0)
}

// Ref is a wrapper around pango_font_metrics_ref().
func (metrics FontMetrics) Ref() FontMetrics {
	ret0 := C.pango_font_metrics_ref(metrics.native())
	return wrapFontMetrics(ret0)
}

// Unref is a wrapper around pango_font_metrics_unref().
func (metrics FontMetrics) Unref() {
	C.pango_font_metrics_unref(metrics.native())
}

// Struct GlyphItemIter
type GlyphItemIter struct {
	Ptr unsafe.Pointer
}

func (v GlyphItemIter) native() *C.PangoGlyphItemIter {
	return (*C.PangoGlyphItemIter)(v.Ptr)
}
func wrapGlyphItemIter(p *C.PangoGlyphItemIter) GlyphItemIter {
	return GlyphItemIter{unsafe.Pointer(p)}
}
func WrapGlyphItemIter(p unsafe.Pointer) GlyphItemIter {
	return GlyphItemIter{p}
}
func (v GlyphItemIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGlyphItemIter(p unsafe.Pointer) interface{} {
	return WrapGlyphItemIter(p)
}

// Copy is a wrapper around pango_glyph_item_iter_copy().
func (orig GlyphItemIter) Copy() GlyphItemIter {
	ret0 := C.pango_glyph_item_iter_copy(orig.native())
	return wrapGlyphItemIter(ret0)
}

// Free is a wrapper around pango_glyph_item_iter_free().
func (iter GlyphItemIter) Free() {
	C.pango_glyph_item_iter_free(iter.native())
}

// InitEnd is a wrapper around pango_glyph_item_iter_init_end().
func (iter GlyphItemIter) InitEnd(glyph_item GlyphItem, text string) bool {
	text0 := C.CString(text)
	ret0 := C.pango_glyph_item_iter_init_end(iter.native(), glyph_item.native(), text0)
	C.free(unsafe.Pointer(text0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// InitStart is a wrapper around pango_glyph_item_iter_init_start().
func (iter GlyphItemIter) InitStart(glyph_item GlyphItem, text string) bool {
	text0 := C.CString(text)
	ret0 := C.pango_glyph_item_iter_init_start(iter.native(), glyph_item.native(), text0)
	C.free(unsafe.Pointer(text0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// NextCluster is a wrapper around pango_glyph_item_iter_next_cluster().
func (iter GlyphItemIter) NextCluster() bool {
	ret0 := C.pango_glyph_item_iter_next_cluster(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PrevCluster is a wrapper around pango_glyph_item_iter_prev_cluster().
func (iter GlyphItemIter) PrevCluster() bool {
	ret0 := C.pango_glyph_item_iter_prev_cluster(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Struct GlyphString
type GlyphString struct {
	Ptr unsafe.Pointer
}

func (v GlyphString) native() *C.PangoGlyphString {
	return (*C.PangoGlyphString)(v.Ptr)
}
func wrapGlyphString(p *C.PangoGlyphString) GlyphString {
	return GlyphString{unsafe.Pointer(p)}
}
func WrapGlyphString(p unsafe.Pointer) GlyphString {
	return GlyphString{p}
}
func (v GlyphString) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGlyphString(p unsafe.Pointer) interface{} {
	return WrapGlyphString(p)
}

// GlyphStringNew is a wrapper around pango_glyph_string_new().
func GlyphStringNew() GlyphString {
	ret0 := C.pango_glyph_string_new()
	return wrapGlyphString(ret0)
}

// Copy is a wrapper around pango_glyph_string_copy().
func (string GlyphString) Copy() GlyphString {
	ret0 := C.pango_glyph_string_copy(string.native())
	return wrapGlyphString(ret0)
}

// Extents is a wrapper around pango_glyph_string_extents().
func (glyphs GlyphString) Extents(font Font) (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_glyph_string_extents(glyphs.native(), font.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// ExtentsRange is a wrapper around pango_glyph_string_extents_range().
func (glyphs GlyphString) ExtentsRange(start int, end int, font Font) (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_glyph_string_extents_range(glyphs.native(), C.int(start), C.int(end), font.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// Free is a wrapper around pango_glyph_string_free().
func (string GlyphString) Free() {
	C.pango_glyph_string_free(string.native())
}

// GetLogicalWidths is a wrapper around pango_glyph_string_get_logical_widths().
func (glyphs GlyphString) GetLogicalWidths(text string, length int, embedding_level int, logical_widths []int) {
	text0 := C.CString(text)
	logical_widths0 := make([]C.int, len(logical_widths))
	for idx, elemG := range logical_widths {
		logical_widths0[idx] = C.int(elemG)
	}
	var logical_widths0Ptr *C.int
	if len(logical_widths0) > 0 {
		logical_widths0Ptr = &logical_widths0[0]
	}
	C.pango_glyph_string_get_logical_widths(glyphs.native(), text0, C.int(length), C.int(embedding_level), logical_widths0Ptr)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// GetWidth is a wrapper around pango_glyph_string_get_width().
func (glyphs GlyphString) GetWidth() int {
	ret0 := C.pango_glyph_string_get_width(glyphs.native())
	return int(ret0)
}

// IndexToX is a wrapper around pango_glyph_string_index_to_x().
func (glyphs GlyphString) IndexToX(text string, length int, analysis Analysis, index_ int, trailing bool) int {
	text0 := C.CString(text)
	var x_pos0 C.int
	C.pango_glyph_string_index_to_x(glyphs.native(), text0, C.int(length), analysis.native(), C.int(index_), C.gboolean(util.Bool2Int(trailing)) /*go:.util*/, &x_pos0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
	return int(x_pos0)
}

// SetSize is a wrapper around pango_glyph_string_set_size().
func (string GlyphString) SetSize(new_len int) {
	C.pango_glyph_string_set_size(string.native(), C.gint(new_len))
}

// XToIndex is a wrapper around pango_glyph_string_x_to_index().
func (glyphs GlyphString) XToIndex(text string, length int, analysis Analysis, x_pos int) (int, int) {
	text0 := C.CString(text)
	var index_0 C.int
	var trailing0 C.int
	C.pango_glyph_string_x_to_index(glyphs.native(), text0, C.int(length), analysis.native(), C.int(x_pos), &index_0, &trailing0)
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
	return int(index_0), int(trailing0)
}

// Struct Item
type Item struct {
	Ptr unsafe.Pointer
}

func (v Item) native() *C.PangoItem {
	return (*C.PangoItem)(v.Ptr)
}
func wrapItem(p *C.PangoItem) Item {
	return Item{unsafe.Pointer(p)}
}
func WrapItem(p unsafe.Pointer) Item {
	return Item{p}
}
func (v Item) IsNil() bool {
	return v.Ptr == nil
}
func IWrapItem(p unsafe.Pointer) interface{} {
	return WrapItem(p)
}

// ItemNew is a wrapper around pango_item_new().
func ItemNew() Item {
	ret0 := C.pango_item_new()
	return wrapItem(ret0)
}

// Copy is a wrapper around pango_item_copy().
func (item Item) Copy() Item {
	ret0 := C.pango_item_copy(item.native())
	return wrapItem(ret0)
}

// Free is a wrapper around pango_item_free().
func (item Item) Free() {
	C.pango_item_free(item.native())
}

// Split is a wrapper around pango_item_split().
func (orig Item) Split(split_index int, split_offset int) Item {
	ret0 := C.pango_item_split(orig.native(), C.int(split_index), C.int(split_offset))
	return wrapItem(ret0)
}

// Struct LayoutIter
type LayoutIter struct {
	Ptr unsafe.Pointer
}

func (v LayoutIter) native() *C.PangoLayoutIter {
	return (*C.PangoLayoutIter)(v.Ptr)
}
func wrapLayoutIter(p *C.PangoLayoutIter) LayoutIter {
	return LayoutIter{unsafe.Pointer(p)}
}
func WrapLayoutIter(p unsafe.Pointer) LayoutIter {
	return LayoutIter{p}
}
func (v LayoutIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLayoutIter(p unsafe.Pointer) interface{} {
	return WrapLayoutIter(p)
}

// AtLastLine is a wrapper around pango_layout_iter_at_last_line().
func (iter LayoutIter) AtLastLine() bool {
	ret0 := C.pango_layout_iter_at_last_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Copy is a wrapper around pango_layout_iter_copy().
func (iter LayoutIter) Copy() LayoutIter {
	ret0 := C.pango_layout_iter_copy(iter.native())
	return wrapLayoutIter(ret0)
}

// Free is a wrapper around pango_layout_iter_free().
func (iter LayoutIter) Free() {
	C.pango_layout_iter_free(iter.native())
}

// GetBaseline is a wrapper around pango_layout_iter_get_baseline().
func (iter LayoutIter) GetBaseline() int {
	ret0 := C.pango_layout_iter_get_baseline(iter.native())
	return int(ret0)
}

// GetCharExtents is a wrapper around pango_layout_iter_get_char_extents().
func (iter LayoutIter) GetCharExtents() Rectangle {
	var logical_rect0 C.PangoRectangle
	C.pango_layout_iter_get_char_extents(iter.native(), &logical_rect0)
	return wrapRectangle(&logical_rect0)
}

// GetClusterExtents is a wrapper around pango_layout_iter_get_cluster_extents().
func (iter LayoutIter) GetClusterExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_iter_get_cluster_extents(iter.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// GetIndex is a wrapper around pango_layout_iter_get_index().
func (iter LayoutIter) GetIndex() int {
	ret0 := C.pango_layout_iter_get_index(iter.native())
	return int(ret0)
}

// GetLayout is a wrapper around pango_layout_iter_get_layout().
func (iter LayoutIter) GetLayout() Layout {
	ret0 := C.pango_layout_iter_get_layout(iter.native())
	return wrapLayout(ret0)
}

// GetLayoutExtents is a wrapper around pango_layout_iter_get_layout_extents().
func (iter LayoutIter) GetLayoutExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_iter_get_layout_extents(iter.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// GetLine is a wrapper around pango_layout_iter_get_line().
func (iter LayoutIter) GetLine() LayoutLine {
	ret0 := C.pango_layout_iter_get_line(iter.native())
	return wrapLayoutLine(ret0)
}

// GetLineExtents is a wrapper around pango_layout_iter_get_line_extents().
func (iter LayoutIter) GetLineExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_iter_get_line_extents(iter.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// GetLineReadonly is a wrapper around pango_layout_iter_get_line_readonly().
func (iter LayoutIter) GetLineReadonly() LayoutLine {
	ret0 := C.pango_layout_iter_get_line_readonly(iter.native())
	return wrapLayoutLine(ret0)
}

// GetLineYrange is a wrapper around pango_layout_iter_get_line_yrange().
func (iter LayoutIter) GetLineYrange() (int, int) {
	var y0_0 C.int
	var y1_0 C.int
	C.pango_layout_iter_get_line_yrange(iter.native(), &y0_0, &y1_0)
	return int(y0_0), int(y1_0)
}

// GetRunExtents is a wrapper around pango_layout_iter_get_run_extents().
func (iter LayoutIter) GetRunExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_iter_get_run_extents(iter.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// NextChar is a wrapper around pango_layout_iter_next_char().
func (iter LayoutIter) NextChar() bool {
	ret0 := C.pango_layout_iter_next_char(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// NextCluster is a wrapper around pango_layout_iter_next_cluster().
func (iter LayoutIter) NextCluster() bool {
	ret0 := C.pango_layout_iter_next_cluster(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// NextLine is a wrapper around pango_layout_iter_next_line().
func (iter LayoutIter) NextLine() bool {
	ret0 := C.pango_layout_iter_next_line(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// NextRun is a wrapper around pango_layout_iter_next_run().
func (iter LayoutIter) NextRun() bool {
	ret0 := C.pango_layout_iter_next_run(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object Layout
type Layout struct {
	gobject.Object
}

func (v Layout) native() *C.PangoLayout {
	return (*C.PangoLayout)(v.Ptr)
}
func wrapLayout(p *C.PangoLayout) (v Layout) {
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
	return gobject.Type(C.pango_layout_get_type())
}
func (v Layout) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLayout(unsafe.Pointer(ptr)), nil
	}
}

// LayoutNew is a wrapper around pango_layout_new().
func LayoutNew(context Context) Layout {
	ret0 := C.pango_layout_new(context.native())
	return wrapLayout(ret0)
}

// ContextChanged is a wrapper around pango_layout_context_changed().
func (layout Layout) ContextChanged() {
	C.pango_layout_context_changed(layout.native())
}

// Copy is a wrapper around pango_layout_copy().
func (src Layout) Copy() Layout {
	ret0 := C.pango_layout_copy(src.native())
	return wrapLayout(ret0)
}

// GetAlignment is a wrapper around pango_layout_get_alignment().
func (layout Layout) GetAlignment() Alignment {
	ret0 := C.pango_layout_get_alignment(layout.native())
	return Alignment(ret0)
}

// GetAttributes is a wrapper around pango_layout_get_attributes().
func (layout Layout) GetAttributes() AttrList {
	ret0 := C.pango_layout_get_attributes(layout.native())
	return wrapAttrList(ret0)
}

// GetAutoDir is a wrapper around pango_layout_get_auto_dir().
func (layout Layout) GetAutoDir() bool {
	ret0 := C.pango_layout_get_auto_dir(layout.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetBaseline is a wrapper around pango_layout_get_baseline().
func (layout Layout) GetBaseline() int {
	ret0 := C.pango_layout_get_baseline(layout.native())
	return int(ret0)
}

// GetCharacterCount is a wrapper around pango_layout_get_character_count().
func (layout Layout) GetCharacterCount() int {
	ret0 := C.pango_layout_get_character_count(layout.native())
	return int(ret0)
}

// GetContext is a wrapper around pango_layout_get_context().
func (layout Layout) GetContext() Context {
	ret0 := C.pango_layout_get_context(layout.native())
	return wrapContext(ret0)
}

// GetCursorPos is a wrapper around pango_layout_get_cursor_pos().
func (layout Layout) GetCursorPos(index_ int) (Rectangle, Rectangle) {
	var strong_pos0 C.PangoRectangle
	var weak_pos0 C.PangoRectangle
	C.pango_layout_get_cursor_pos(layout.native(), C.int(index_), &strong_pos0, &weak_pos0)
	return wrapRectangle(&strong_pos0), wrapRectangle(&weak_pos0)
}

// GetEllipsize is a wrapper around pango_layout_get_ellipsize().
func (layout Layout) GetEllipsize() EllipsizeMode {
	ret0 := C.pango_layout_get_ellipsize(layout.native())
	return EllipsizeMode(ret0)
}

// GetExtents is a wrapper around pango_layout_get_extents().
func (layout Layout) GetExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_get_extents(layout.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// GetFontDescription is a wrapper around pango_layout_get_font_description().
func (layout Layout) GetFontDescription() FontDescription {
	ret0 := C.pango_layout_get_font_description(layout.native())
	return wrapFontDescription(ret0)
}

// GetHeight is a wrapper around pango_layout_get_height().
func (layout Layout) GetHeight() int {
	ret0 := C.pango_layout_get_height(layout.native())
	return int(ret0)
}

// GetIndent is a wrapper around pango_layout_get_indent().
func (layout Layout) GetIndent() int {
	ret0 := C.pango_layout_get_indent(layout.native())
	return int(ret0)
}

// GetIter is a wrapper around pango_layout_get_iter().
func (layout Layout) GetIter() LayoutIter {
	ret0 := C.pango_layout_get_iter(layout.native())
	return wrapLayoutIter(ret0)
}

// GetJustify is a wrapper around pango_layout_get_justify().
func (layout Layout) GetJustify() bool {
	ret0 := C.pango_layout_get_justify(layout.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLine is a wrapper around pango_layout_get_line().
func (layout Layout) GetLine(line int) LayoutLine {
	ret0 := C.pango_layout_get_line(layout.native(), C.int(line))
	return wrapLayoutLine(ret0)
}

// GetLineCount is a wrapper around pango_layout_get_line_count().
func (layout Layout) GetLineCount() int {
	ret0 := C.pango_layout_get_line_count(layout.native())
	return int(ret0)
}

// GetLineReadonly is a wrapper around pango_layout_get_line_readonly().
func (layout Layout) GetLineReadonly(line int) LayoutLine {
	ret0 := C.pango_layout_get_line_readonly(layout.native(), C.int(line))
	return wrapLayoutLine(ret0)
}

// GetLogAttrsReadonly is a wrapper around pango_layout_get_log_attrs_readonly().
func (layout Layout) GetLogAttrsReadonly() []LogAttr {
	var n_attrs0 C.gint
	ret0 := C.pango_layout_get_log_attrs_readonly(layout.native(), &n_attrs0)
	var ret0Slice []C.PangoLogAttr
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(n_attrs0)) /*go:.util*/
	ret := make([]LogAttr, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = wrapLogAttr(&elem)
	}
	return ret
}

// GetPixelExtents is a wrapper around pango_layout_get_pixel_extents().
func (layout Layout) GetPixelExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_get_pixel_extents(layout.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// GetPixelSize is a wrapper around pango_layout_get_pixel_size().
func (layout Layout) GetPixelSize() (int, int) {
	var width0 C.int
	var height0 C.int
	C.pango_layout_get_pixel_size(layout.native(), &width0, &height0)
	return int(width0), int(height0)
}

// GetSerial is a wrapper around pango_layout_get_serial().
func (layout Layout) GetSerial() uint {
	ret0 := C.pango_layout_get_serial(layout.native())
	return uint(ret0)
}

// GetSingleParagraphMode is a wrapper around pango_layout_get_single_paragraph_mode().
func (layout Layout) GetSingleParagraphMode() bool {
	ret0 := C.pango_layout_get_single_paragraph_mode(layout.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSize is a wrapper around pango_layout_get_size().
func (layout Layout) GetSize() (int, int) {
	var width0 C.int
	var height0 C.int
	C.pango_layout_get_size(layout.native(), &width0, &height0)
	return int(width0), int(height0)
}

// GetSpacing is a wrapper around pango_layout_get_spacing().
func (layout Layout) GetSpacing() int {
	ret0 := C.pango_layout_get_spacing(layout.native())
	return int(ret0)
}

// GetTabs is a wrapper around pango_layout_get_tabs().
func (layout Layout) GetTabs() TabArray {
	ret0 := C.pango_layout_get_tabs(layout.native())
	return wrapTabArray(ret0)
}

// GetText is a wrapper around pango_layout_get_text().
func (layout Layout) GetText() string {
	ret0 := C.pango_layout_get_text(layout.native())
	ret := C.GoString(ret0)
	return ret
}

// GetUnknownGlyphsCount is a wrapper around pango_layout_get_unknown_glyphs_count().
func (layout Layout) GetUnknownGlyphsCount() int {
	ret0 := C.pango_layout_get_unknown_glyphs_count(layout.native())
	return int(ret0)
}

// GetWidth is a wrapper around pango_layout_get_width().
func (layout Layout) GetWidth() int {
	ret0 := C.pango_layout_get_width(layout.native())
	return int(ret0)
}

// GetWrap is a wrapper around pango_layout_get_wrap().
func (layout Layout) GetWrap() WrapMode {
	ret0 := C.pango_layout_get_wrap(layout.native())
	return WrapMode(ret0)
}

// IndexToLineX is a wrapper around pango_layout_index_to_line_x().
func (layout Layout) IndexToLineX(index_ int, trailing bool) (int, int) {
	var line0 C.int
	var x_pos0 C.int
	C.pango_layout_index_to_line_x(layout.native(), C.int(index_), C.gboolean(util.Bool2Int(trailing)) /*go:.util*/, &line0, &x_pos0)
	return int(line0), int(x_pos0)
}

// IndexToPos is a wrapper around pango_layout_index_to_pos().
func (layout Layout) IndexToPos(index_ int) Rectangle {
	var pos0 C.PangoRectangle
	C.pango_layout_index_to_pos(layout.native(), C.int(index_), &pos0)
	return wrapRectangle(&pos0)
}

// IsEllipsized is a wrapper around pango_layout_is_ellipsized().
func (layout Layout) IsEllipsized() bool {
	ret0 := C.pango_layout_is_ellipsized(layout.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsWrapped is a wrapper around pango_layout_is_wrapped().
func (layout Layout) IsWrapped() bool {
	ret0 := C.pango_layout_is_wrapped(layout.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MoveCursorVisually is a wrapper around pango_layout_move_cursor_visually().
func (layout Layout) MoveCursorVisually(strong bool, old_index int, old_trailing int, direction int) (int, int) {
	var new_index0 C.int
	var new_trailing0 C.int
	C.pango_layout_move_cursor_visually(layout.native(), C.gboolean(util.Bool2Int(strong)) /*go:.util*/, C.int(old_index), C.int(old_trailing), C.int(direction), &new_index0, &new_trailing0)
	return int(new_index0), int(new_trailing0)
}

// SetAlignment is a wrapper around pango_layout_set_alignment().
func (layout Layout) SetAlignment(alignment Alignment) {
	C.pango_layout_set_alignment(layout.native(), C.PangoAlignment(alignment))
}

// SetAttributes is a wrapper around pango_layout_set_attributes().
func (layout Layout) SetAttributes(attrs AttrList) {
	C.pango_layout_set_attributes(layout.native(), attrs.native())
}

// SetAutoDir is a wrapper around pango_layout_set_auto_dir().
func (layout Layout) SetAutoDir(auto_dir bool) {
	C.pango_layout_set_auto_dir(layout.native(), C.gboolean(util.Bool2Int(auto_dir)) /*go:.util*/)
}

// SetEllipsize is a wrapper around pango_layout_set_ellipsize().
func (layout Layout) SetEllipsize(ellipsize EllipsizeMode) {
	C.pango_layout_set_ellipsize(layout.native(), C.PangoEllipsizeMode(ellipsize))
}

// SetFontDescription is a wrapper around pango_layout_set_font_description().
func (layout Layout) SetFontDescription(desc FontDescription) {
	C.pango_layout_set_font_description(layout.native(), desc.native())
}

// SetHeight is a wrapper around pango_layout_set_height().
func (layout Layout) SetHeight(height int) {
	C.pango_layout_set_height(layout.native(), C.int(height))
}

// SetIndent is a wrapper around pango_layout_set_indent().
func (layout Layout) SetIndent(indent int) {
	C.pango_layout_set_indent(layout.native(), C.int(indent))
}

// SetJustify is a wrapper around pango_layout_set_justify().
func (layout Layout) SetJustify(justify bool) {
	C.pango_layout_set_justify(layout.native(), C.gboolean(util.Bool2Int(justify)) /*go:.util*/)
}

// SetMarkup is a wrapper around pango_layout_set_markup().
func (layout Layout) SetMarkup(markup string, length int) {
	markup0 := C.CString(markup)
	C.pango_layout_set_markup(layout.native(), markup0, C.int(length))
	C.free(unsafe.Pointer(markup0)) /*ch:<stdlib.h>*/
}

// SetSingleParagraphMode is a wrapper around pango_layout_set_single_paragraph_mode().
func (layout Layout) SetSingleParagraphMode(setting bool) {
	C.pango_layout_set_single_paragraph_mode(layout.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetSpacing is a wrapper around pango_layout_set_spacing().
func (layout Layout) SetSpacing(spacing int) {
	C.pango_layout_set_spacing(layout.native(), C.int(spacing))
}

// SetTabs is a wrapper around pango_layout_set_tabs().
func (layout Layout) SetTabs(tabs TabArray) {
	C.pango_layout_set_tabs(layout.native(), tabs.native())
}

// SetText is a wrapper around pango_layout_set_text().
func (layout Layout) SetText(text string, length int) {
	text0 := C.CString(text)
	C.pango_layout_set_text(layout.native(), text0, C.int(length))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// SetWidth is a wrapper around pango_layout_set_width().
func (layout Layout) SetWidth(width int) {
	C.pango_layout_set_width(layout.native(), C.int(width))
}

// SetWrap is a wrapper around pango_layout_set_wrap().
func (layout Layout) SetWrap(wrap WrapMode) {
	C.pango_layout_set_wrap(layout.native(), C.PangoWrapMode(wrap))
}

// XyToIndex is a wrapper around pango_layout_xy_to_index().
func (layout Layout) XyToIndex(x int, y int) (bool, int, int) {
	var index_0 C.int
	var trailing0 C.int
	ret0 := C.pango_layout_xy_to_index(layout.native(), C.int(x), C.int(y), &index_0, &trailing0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(index_0), int(trailing0)
}

// Object Context
type Context struct {
	gobject.Object
}

func (v Context) native() *C.PangoContext {
	return (*C.PangoContext)(v.Ptr)
}
func wrapContext(p *C.PangoContext) (v Context) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapContext(p unsafe.Pointer) (v Context) {
	v.Ptr = p
	return
}
func (v Context) IsNil() bool {
	return v.Ptr == nil
}
func IWrapContext(p unsafe.Pointer) interface{} {
	return WrapContext(p)
}
func (v Context) GetType() gobject.Type {
	return gobject.Type(C.pango_context_get_type())
}
func (v Context) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapContext(unsafe.Pointer(ptr)), nil
	}
}

// ContextNew is a wrapper around pango_context_new().
func ContextNew() Context {
	ret0 := C.pango_context_new()
	return wrapContext(ret0)
}

// Changed is a wrapper around pango_context_changed().
func (context Context) Changed() {
	C.pango_context_changed(context.native())
}

// GetBaseDir is a wrapper around pango_context_get_base_dir().
func (context Context) GetBaseDir() Direction {
	ret0 := C.pango_context_get_base_dir(context.native())
	return Direction(ret0)
}

// GetBaseGravity is a wrapper around pango_context_get_base_gravity().
func (context Context) GetBaseGravity() Gravity {
	ret0 := C.pango_context_get_base_gravity(context.native())
	return Gravity(ret0)
}

// GetFontDescription is a wrapper around pango_context_get_font_description().
func (context Context) GetFontDescription() FontDescription {
	ret0 := C.pango_context_get_font_description(context.native())
	return wrapFontDescription(ret0)
}

// GetFontMap is a wrapper around pango_context_get_font_map().
func (context Context) GetFontMap() FontMap {
	ret0 := C.pango_context_get_font_map(context.native())
	return wrapFontMap(ret0)
}

// GetGravity is a wrapper around pango_context_get_gravity().
func (context Context) GetGravity() Gravity {
	ret0 := C.pango_context_get_gravity(context.native())
	return Gravity(ret0)
}

// GetGravityHint is a wrapper around pango_context_get_gravity_hint().
func (context Context) GetGravityHint() GravityHint {
	ret0 := C.pango_context_get_gravity_hint(context.native())
	return GravityHint(ret0)
}

// GetLanguage is a wrapper around pango_context_get_language().
func (context Context) GetLanguage() Language {
	ret0 := C.pango_context_get_language(context.native())
	return wrapLanguage(ret0)
}

// GetMatrix is a wrapper around pango_context_get_matrix().
func (context Context) GetMatrix() Matrix {
	ret0 := C.pango_context_get_matrix(context.native())
	return wrapMatrix(ret0)
}

// GetMetrics is a wrapper around pango_context_get_metrics().
func (context Context) GetMetrics(desc FontDescription, language Language) FontMetrics {
	ret0 := C.pango_context_get_metrics(context.native(), desc.native(), language.native())
	return wrapFontMetrics(ret0)
}

// GetSerial is a wrapper around pango_context_get_serial().
func (context Context) GetSerial() uint {
	ret0 := C.pango_context_get_serial(context.native())
	return uint(ret0)
}

// LoadFont is a wrapper around pango_context_load_font().
func (context Context) LoadFont(desc FontDescription) Font {
	ret0 := C.pango_context_load_font(context.native(), desc.native())
	return wrapFont(ret0)
}

// LoadFontset is a wrapper around pango_context_load_fontset().
func (context Context) LoadFontset(desc FontDescription, language Language) Fontset {
	ret0 := C.pango_context_load_fontset(context.native(), desc.native(), language.native())
	return wrapFontset(ret0)
}

// SetBaseDir is a wrapper around pango_context_set_base_dir().
func (context Context) SetBaseDir(direction Direction) {
	C.pango_context_set_base_dir(context.native(), C.PangoDirection(direction))
}

// SetBaseGravity is a wrapper around pango_context_set_base_gravity().
func (context Context) SetBaseGravity(gravity Gravity) {
	C.pango_context_set_base_gravity(context.native(), C.PangoGravity(gravity))
}

// SetFontDescription is a wrapper around pango_context_set_font_description().
func (context Context) SetFontDescription(desc FontDescription) {
	C.pango_context_set_font_description(context.native(), desc.native())
}

// SetFontMap is a wrapper around pango_context_set_font_map().
func (context Context) SetFontMap(font_map FontMap) {
	C.pango_context_set_font_map(context.native(), font_map.native())
}

// SetGravityHint is a wrapper around pango_context_set_gravity_hint().
func (context Context) SetGravityHint(hint GravityHint) {
	C.pango_context_set_gravity_hint(context.native(), C.PangoGravityHint(hint))
}

// SetLanguage is a wrapper around pango_context_set_language().
func (context Context) SetLanguage(language Language) {
	C.pango_context_set_language(context.native(), language.native())
}

// SetMatrix is a wrapper around pango_context_set_matrix().
func (context Context) SetMatrix(matrix Matrix) {
	C.pango_context_set_matrix(context.native(), matrix.native())
}

// Object FontMap
type FontMap struct {
	gobject.Object
}

func (v FontMap) native() *C.PangoFontMap {
	return (*C.PangoFontMap)(v.Ptr)
}
func wrapFontMap(p *C.PangoFontMap) (v FontMap) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontMap(p unsafe.Pointer) (v FontMap) {
	v.Ptr = p
	return
}
func (v FontMap) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontMap(p unsafe.Pointer) interface{} {
	return WrapFontMap(p)
}
func (v FontMap) GetType() gobject.Type {
	return gobject.Type(C.pango_font_map_get_type())
}
func (v FontMap) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontMap(unsafe.Pointer(ptr)), nil
	}
}

// Changed is a wrapper around pango_font_map_changed().
func (fontmap FontMap) Changed() {
	C.pango_font_map_changed(fontmap.native())
}

// CreateContext is a wrapper around pango_font_map_create_context().
func (fontmap FontMap) CreateContext() Context {
	ret0 := C.pango_font_map_create_context(fontmap.native())
	return wrapContext(ret0)
}

// GetSerial is a wrapper around pango_font_map_get_serial().
func (fontmap FontMap) GetSerial() uint {
	ret0 := C.pango_font_map_get_serial(fontmap.native())
	return uint(ret0)
}

// LoadFont is a wrapper around pango_font_map_load_font().
func (fontmap FontMap) LoadFont(context Context, desc FontDescription) Font {
	ret0 := C.pango_font_map_load_font(fontmap.native(), context.native(), desc.native())
	return wrapFont(ret0)
}

// LoadFontset is a wrapper around pango_font_map_load_fontset().
func (fontmap FontMap) LoadFontset(context Context, desc FontDescription, language Language) Fontset {
	ret0 := C.pango_font_map_load_fontset(fontmap.native(), context.native(), desc.native(), language.native())
	return wrapFontset(ret0)
}

// Object Font
type Font struct {
	gobject.Object
}

func (v Font) native() *C.PangoFont {
	return (*C.PangoFont)(v.Ptr)
}
func wrapFont(p *C.PangoFont) (v Font) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFont(p unsafe.Pointer) (v Font) {
	v.Ptr = p
	return
}
func (v Font) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFont(p unsafe.Pointer) interface{} {
	return WrapFont(p)
}
func (v Font) GetType() gobject.Type {
	return gobject.Type(C.pango_font_get_type())
}
func (v Font) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFont(unsafe.Pointer(ptr)), nil
	}
}

// Describe is a wrapper around pango_font_describe().
func (font Font) Describe() FontDescription {
	ret0 := C.pango_font_describe(font.native())
	return wrapFontDescription(ret0)
}

// DescribeWithAbsoluteSize is a wrapper around pango_font_describe_with_absolute_size().
func (font Font) DescribeWithAbsoluteSize() FontDescription {
	ret0 := C.pango_font_describe_with_absolute_size(font.native())
	return wrapFontDescription(ret0)
}

// GetCoverage is a wrapper around pango_font_get_coverage().
func (font Font) GetCoverage(language Language) Coverage {
	ret0 := C.pango_font_get_coverage(font.native(), language.native())
	return wrapCoverage(ret0)
}

// GetFontMap is a wrapper around pango_font_get_font_map().
func (font Font) GetFontMap() FontMap {
	ret0 := C.pango_font_get_font_map(font.native())
	return wrapFontMap(ret0)
}

// GetGlyphExtents is a wrapper around pango_font_get_glyph_extents().
func (font Font) GetGlyphExtents(glyph Glyph) (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_font_get_glyph_extents(font.native(), C.PangoGlyph(glyph), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// GetMetrics is a wrapper around pango_font_get_metrics().
func (font Font) GetMetrics(language Language) FontMetrics {
	ret0 := C.pango_font_get_metrics(font.native(), language.native())
	return wrapFontMetrics(ret0)
}

// FontDescriptionsFree is a wrapper around pango_font_descriptions_free().
func FontDescriptionsFree(descs []FontDescription) {
	descs0 := make([]*C.PangoFontDescription, len(descs))
	for idx, elemG := range descs {
		descs0[idx] = elemG.native()
	}
	var descs0Ptr **C.PangoFontDescription
	if len(descs0) > 0 {
		descs0Ptr = &descs0[0]
	}
	C.pango_font_descriptions_free(descs0Ptr, C.int(len(descs)))
}

// Struct Coverage
type Coverage struct {
	Ptr unsafe.Pointer
}

func (v Coverage) native() *C.PangoCoverage {
	return (*C.PangoCoverage)(v.Ptr)
}
func wrapCoverage(p *C.PangoCoverage) Coverage {
	return Coverage{unsafe.Pointer(p)}
}
func WrapCoverage(p unsafe.Pointer) Coverage {
	return Coverage{p}
}
func (v Coverage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCoverage(p unsafe.Pointer) interface{} {
	return WrapCoverage(p)
}

// Copy is a wrapper around pango_coverage_copy().
func (coverage Coverage) Copy() Coverage {
	ret0 := C.pango_coverage_copy(coverage.native())
	return wrapCoverage(ret0)
}

// Get is a wrapper around pango_coverage_get().
func (coverage Coverage) Get(index_ int) CoverageLevel {
	ret0 := C.pango_coverage_get(coverage.native(), C.int(index_))
	return CoverageLevel(ret0)
}

// Max is a wrapper around pango_coverage_max().
func (coverage Coverage) Max(other Coverage) {
	C.pango_coverage_max(coverage.native(), other.native())
}

// Ref is a wrapper around pango_coverage_ref().
func (coverage Coverage) Ref() Coverage {
	ret0 := C.pango_coverage_ref(coverage.native())
	return wrapCoverage(ret0)
}

// Set is a wrapper around pango_coverage_set().
func (coverage Coverage) Set(index_ int, level CoverageLevel) {
	C.pango_coverage_set(coverage.native(), C.int(index_), C.PangoCoverageLevel(level))
}

// Unref is a wrapper around pango_coverage_unref().
func (coverage Coverage) Unref() {
	C.pango_coverage_unref(coverage.native())
}

// CoverageFromBytes is a wrapper around pango_coverage_from_bytes().
func CoverageFromBytes(bytes []byte) Coverage {
	bytes0 := make([]C.guchar, len(bytes))
	for idx, elemG := range bytes {
		bytes0[idx] = C.guchar(elemG)
	}
	var bytes0Ptr *C.guchar
	if len(bytes0) > 0 {
		bytes0Ptr = &bytes0[0]
	}
	ret0 := C.pango_coverage_from_bytes(bytes0Ptr, C.int(len(bytes)))
	return wrapCoverage(ret0)
}

// CoverageNew is a wrapper around pango_coverage_new().
func CoverageNew() Coverage {
	ret0 := C.pango_coverage_new()
	return wrapCoverage(ret0)
}

// Object Fontset
type Fontset struct {
	gobject.Object
}

func (v Fontset) native() *C.PangoFontset {
	return (*C.PangoFontset)(v.Ptr)
}
func wrapFontset(p *C.PangoFontset) (v Fontset) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontset(p unsafe.Pointer) (v Fontset) {
	v.Ptr = p
	return
}
func (v Fontset) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontset(p unsafe.Pointer) interface{} {
	return WrapFontset(p)
}
func (v Fontset) GetType() gobject.Type {
	return gobject.Type(C.pango_fontset_get_type())
}
func (v Fontset) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontset(unsafe.Pointer(ptr)), nil
	}
}

// GetFont is a wrapper around pango_fontset_get_font().
func (fontset Fontset) GetFont(wc uint) Font {
	ret0 := C.pango_fontset_get_font(fontset.native(), C.guint(wc))
	return wrapFont(ret0)
}

// GetMetrics is a wrapper around pango_fontset_get_metrics().
func (fontset Fontset) GetMetrics() FontMetrics {
	ret0 := C.pango_fontset_get_metrics(fontset.native())
	return wrapFontMetrics(ret0)
}

// Struct Matrix
type Matrix struct {
	Ptr unsafe.Pointer
}

func (v Matrix) native() *C.PangoMatrix {
	return (*C.PangoMatrix)(v.Ptr)
}
func wrapMatrix(p *C.PangoMatrix) Matrix {
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

// Concat is a wrapper around pango_matrix_concat().
func (matrix Matrix) Concat(new_matrix Matrix) {
	C.pango_matrix_concat(matrix.native(), new_matrix.native())
}

// Copy is a wrapper around pango_matrix_copy().
func (matrix Matrix) Copy() Matrix {
	ret0 := C.pango_matrix_copy(matrix.native())
	return wrapMatrix(ret0)
}

// Free is a wrapper around pango_matrix_free().
func (matrix Matrix) Free() {
	C.pango_matrix_free(matrix.native())
}

// GetFontScaleFactor is a wrapper around pango_matrix_get_font_scale_factor().
func (matrix Matrix) GetFontScaleFactor() float64 {
	ret0 := C.pango_matrix_get_font_scale_factor(matrix.native())
	return float64(ret0)
}

// GetFontScaleFactors is a wrapper around pango_matrix_get_font_scale_factors().
func (matrix Matrix) GetFontScaleFactors() (float64, float64) {
	var xscale0 C.double
	var yscale0 C.double
	C.pango_matrix_get_font_scale_factors(matrix.native(), &xscale0, &yscale0)
	return float64(xscale0), float64(yscale0)
}

// Rotate is a wrapper around pango_matrix_rotate().
func (matrix Matrix) Rotate(degrees float64) {
	C.pango_matrix_rotate(matrix.native(), C.double(degrees))
}

// Scale is a wrapper around pango_matrix_scale().
func (matrix Matrix) Scale(scale_x float64, scale_y float64) {
	C.pango_matrix_scale(matrix.native(), C.double(scale_x), C.double(scale_y))
}

// Translate is a wrapper around pango_matrix_translate().
func (matrix Matrix) Translate(tx float64, ty float64) {
	C.pango_matrix_translate(matrix.native(), C.double(tx), C.double(ty))
}

// Struct LayoutLine
type LayoutLine struct {
	Ptr unsafe.Pointer
}

func (v LayoutLine) native() *C.PangoLayoutLine {
	return (*C.PangoLayoutLine)(v.Ptr)
}
func wrapLayoutLine(p *C.PangoLayoutLine) LayoutLine {
	return LayoutLine{unsafe.Pointer(p)}
}
func WrapLayoutLine(p unsafe.Pointer) LayoutLine {
	return LayoutLine{p}
}
func (v LayoutLine) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLayoutLine(p unsafe.Pointer) interface{} {
	return WrapLayoutLine(p)
}

// GetExtents is a wrapper around pango_layout_line_get_extents().
func (line LayoutLine) GetExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_line_get_extents(line.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// GetPixelExtents is a wrapper around pango_layout_line_get_pixel_extents().
func (layout_line LayoutLine) GetPixelExtents() (Rectangle, Rectangle) {
	var ink_rect0 C.PangoRectangle
	var logical_rect0 C.PangoRectangle
	C.pango_layout_line_get_pixel_extents(layout_line.native(), &ink_rect0, &logical_rect0)
	return wrapRectangle(&ink_rect0), wrapRectangle(&logical_rect0)
}

// IndexToX is a wrapper around pango_layout_line_index_to_x().
func (line LayoutLine) IndexToX(index_ int, trailing bool) int {
	var x_pos0 C.int
	C.pango_layout_line_index_to_x(line.native(), C.int(index_), C.gboolean(util.Bool2Int(trailing)) /*go:.util*/, &x_pos0)
	return int(x_pos0)
}

// Ref is a wrapper around pango_layout_line_ref().
func (line LayoutLine) Ref() LayoutLine {
	ret0 := C.pango_layout_line_ref(line.native())
	return wrapLayoutLine(ret0)
}

// Unref is a wrapper around pango_layout_line_unref().
func (line LayoutLine) Unref() {
	C.pango_layout_line_unref(line.native())
}

// XToIndex is a wrapper around pango_layout_line_x_to_index().
func (line LayoutLine) XToIndex(x_pos int) (bool, int, int) {
	var index_0 C.int
	var trailing0 C.int
	ret0 := C.pango_layout_line_x_to_index(line.native(), C.int(x_pos), &index_0, &trailing0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(index_0), int(trailing0)
}

// Struct TabArray
type TabArray struct {
	Ptr unsafe.Pointer
}

func (v TabArray) native() *C.PangoTabArray {
	return (*C.PangoTabArray)(v.Ptr)
}
func wrapTabArray(p *C.PangoTabArray) TabArray {
	return TabArray{unsafe.Pointer(p)}
}
func WrapTabArray(p unsafe.Pointer) TabArray {
	return TabArray{p}
}
func (v TabArray) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTabArray(p unsafe.Pointer) interface{} {
	return WrapTabArray(p)
}

// TabArrayNew is a wrapper around pango_tab_array_new().
func TabArrayNew(initial_size int, positions_in_pixels bool) TabArray {
	ret0 := C.pango_tab_array_new(C.gint(initial_size), C.gboolean(util.Bool2Int(positions_in_pixels)) /*go:.util*/)
	return wrapTabArray(ret0)
}

// Copy is a wrapper around pango_tab_array_copy().
func (src TabArray) Copy() TabArray {
	ret0 := C.pango_tab_array_copy(src.native())
	return wrapTabArray(ret0)
}

// Free is a wrapper around pango_tab_array_free().
func (tab_array TabArray) Free() {
	C.pango_tab_array_free(tab_array.native())
}

// GetPositionsInPixels is a wrapper around pango_tab_array_get_positions_in_pixels().
func (tab_array TabArray) GetPositionsInPixels() bool {
	ret0 := C.pango_tab_array_get_positions_in_pixels(tab_array.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSize is a wrapper around pango_tab_array_get_size().
func (tab_array TabArray) GetSize() int {
	ret0 := C.pango_tab_array_get_size(tab_array.native())
	return int(ret0)
}

// GetTab is a wrapper around pango_tab_array_get_tab().
func (tab_array TabArray) GetTab(tab_index int) (TabAlign, int) {
	var alignment0 C.PangoTabAlign
	var location0 C.gint
	C.pango_tab_array_get_tab(tab_array.native(), C.gint(tab_index), &alignment0, &location0)
	return TabAlign(alignment0), int(location0)
}

// Resize is a wrapper around pango_tab_array_resize().
func (tab_array TabArray) Resize(new_size int) {
	C.pango_tab_array_resize(tab_array.native(), C.gint(new_size))
}

// SetTab is a wrapper around pango_tab_array_set_tab().
func (tab_array TabArray) SetTab(tab_index int, alignment TabAlign, location int) {
	C.pango_tab_array_set_tab(tab_array.native(), C.gint(tab_index), C.PangoTabAlign(alignment), C.gint(location))
}

// Object FontFace
type FontFace struct {
	gobject.Object
}

func (v FontFace) native() *C.PangoFontFace {
	return (*C.PangoFontFace)(v.Ptr)
}
func wrapFontFace(p *C.PangoFontFace) (v FontFace) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontFace(p unsafe.Pointer) (v FontFace) {
	v.Ptr = p
	return
}
func (v FontFace) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontFace(p unsafe.Pointer) interface{} {
	return WrapFontFace(p)
}
func (v FontFace) GetType() gobject.Type {
	return gobject.Type(C.pango_font_face_get_type())
}
func (v FontFace) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontFace(unsafe.Pointer(ptr)), nil
	}
}

// Describe is a wrapper around pango_font_face_describe().
func (face FontFace) Describe() FontDescription {
	ret0 := C.pango_font_face_describe(face.native())
	return wrapFontDescription(ret0)
}

// GetFaceName is a wrapper around pango_font_face_get_face_name().
func (face FontFace) GetFaceName() string {
	ret0 := C.pango_font_face_get_face_name(face.native())
	ret := C.GoString(ret0)
	return ret
}

// IsSynthesized is a wrapper around pango_font_face_is_synthesized().
func (face FontFace) IsSynthesized() bool {
	ret0 := C.pango_font_face_is_synthesized(face.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object FontFamily
type FontFamily struct {
	gobject.Object
}

func (v FontFamily) native() *C.PangoFontFamily {
	return (*C.PangoFontFamily)(v.Ptr)
}
func wrapFontFamily(p *C.PangoFontFamily) (v FontFamily) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFontFamily(p unsafe.Pointer) (v FontFamily) {
	v.Ptr = p
	return
}
func (v FontFamily) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFontFamily(p unsafe.Pointer) interface{} {
	return WrapFontFamily(p)
}
func (v FontFamily) GetType() gobject.Type {
	return gobject.Type(C.pango_font_family_get_type())
}
func (v FontFamily) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFontFamily(unsafe.Pointer(ptr)), nil
	}
}

// GetName is a wrapper around pango_font_family_get_name().
func (family FontFamily) GetName() string {
	ret0 := C.pango_font_family_get_name(family.native())
	ret := C.GoString(ret0)
	return ret
}

// IsMonospace is a wrapper around pango_font_family_is_monospace().
func (family FontFamily) IsMonospace() bool {
	ret0 := C.pango_font_family_is_monospace(family.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object Renderer
type Renderer struct {
	gobject.Object
}

func (v Renderer) native() *C.PangoRenderer {
	return (*C.PangoRenderer)(v.Ptr)
}
func wrapRenderer(p *C.PangoRenderer) (v Renderer) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRenderer(p unsafe.Pointer) (v Renderer) {
	v.Ptr = p
	return
}
func (v Renderer) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRenderer(p unsafe.Pointer) interface{} {
	return WrapRenderer(p)
}
func (v Renderer) GetType() gobject.Type {
	return gobject.Type(C.pango_renderer_get_type())
}
func (v Renderer) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRenderer(unsafe.Pointer(ptr)), nil
	}
}

// Activate is a wrapper around pango_renderer_activate().
func (renderer Renderer) Activate() {
	C.pango_renderer_activate(renderer.native())
}

// Deactivate is a wrapper around pango_renderer_deactivate().
func (renderer Renderer) Deactivate() {
	C.pango_renderer_deactivate(renderer.native())
}

// DrawErrorUnderline is a wrapper around pango_renderer_draw_error_underline().
func (renderer Renderer) DrawErrorUnderline(x int, y int, width int, height int) {
	C.pango_renderer_draw_error_underline(renderer.native(), C.int(x), C.int(y), C.int(width), C.int(height))
}

// DrawGlyph is a wrapper around pango_renderer_draw_glyph().
func (renderer Renderer) DrawGlyph(font Font, glyph Glyph, x float64, y float64) {
	C.pango_renderer_draw_glyph(renderer.native(), font.native(), C.PangoGlyph(glyph), C.double(x), C.double(y))
}

// DrawGlyphItem is a wrapper around pango_renderer_draw_glyph_item().
func (renderer Renderer) DrawGlyphItem(text string, glyph_item GlyphItem, x int, y int) {
	text0 := C.CString(text)
	C.pango_renderer_draw_glyph_item(renderer.native(), text0, glyph_item.native(), C.int(x), C.int(y))
	C.free(unsafe.Pointer(text0)) /*ch:<stdlib.h>*/
}

// DrawGlyphs is a wrapper around pango_renderer_draw_glyphs().
func (renderer Renderer) DrawGlyphs(font Font, glyphs GlyphString, x int, y int) {
	C.pango_renderer_draw_glyphs(renderer.native(), font.native(), glyphs.native(), C.int(x), C.int(y))
}

// DrawLayout is a wrapper around pango_renderer_draw_layout().
func (renderer Renderer) DrawLayout(layout Layout, x int, y int) {
	C.pango_renderer_draw_layout(renderer.native(), layout.native(), C.int(x), C.int(y))
}

// DrawLayoutLine is a wrapper around pango_renderer_draw_layout_line().
func (renderer Renderer) DrawLayoutLine(line LayoutLine, x int, y int) {
	C.pango_renderer_draw_layout_line(renderer.native(), line.native(), C.int(x), C.int(y))
}

// DrawRectangle is a wrapper around pango_renderer_draw_rectangle().
func (renderer Renderer) DrawRectangle(part RenderPart, x int, y int, width int, height int) {
	C.pango_renderer_draw_rectangle(renderer.native(), C.PangoRenderPart(part), C.int(x), C.int(y), C.int(width), C.int(height))
}

// DrawTrapezoid is a wrapper around pango_renderer_draw_trapezoid().
func (renderer Renderer) DrawTrapezoid(part RenderPart, y1_ float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	C.pango_renderer_draw_trapezoid(renderer.native(), C.PangoRenderPart(part), C.double(y1_), C.double(x11), C.double(x21), C.double(y2), C.double(x12), C.double(x22))
}

// GetAlpha is a wrapper around pango_renderer_get_alpha().
func (renderer Renderer) GetAlpha(part RenderPart) uint16 {
	ret0 := C.pango_renderer_get_alpha(renderer.native(), C.PangoRenderPart(part))
	return uint16(ret0)
}

// GetColor is a wrapper around pango_renderer_get_color().
func (renderer Renderer) GetColor(part RenderPart) Color {
	ret0 := C.pango_renderer_get_color(renderer.native(), C.PangoRenderPart(part))
	return wrapColor(ret0)
}

// GetLayout is a wrapper around pango_renderer_get_layout().
func (renderer Renderer) GetLayout() Layout {
	ret0 := C.pango_renderer_get_layout(renderer.native())
	return wrapLayout(ret0)
}

// GetLayoutLine is a wrapper around pango_renderer_get_layout_line().
func (renderer Renderer) GetLayoutLine() LayoutLine {
	ret0 := C.pango_renderer_get_layout_line(renderer.native())
	return wrapLayoutLine(ret0)
}

// GetMatrix is a wrapper around pango_renderer_get_matrix().
func (renderer Renderer) GetMatrix() Matrix {
	ret0 := C.pango_renderer_get_matrix(renderer.native())
	return wrapMatrix(ret0)
}

// PartChanged is a wrapper around pango_renderer_part_changed().
func (renderer Renderer) PartChanged(part RenderPart) {
	C.pango_renderer_part_changed(renderer.native(), C.PangoRenderPart(part))
}

// SetAlpha is a wrapper around pango_renderer_set_alpha().
func (renderer Renderer) SetAlpha(part RenderPart, alpha uint16) {
	C.pango_renderer_set_alpha(renderer.native(), C.PangoRenderPart(part), C.guint16(alpha))
}

// SetColor is a wrapper around pango_renderer_set_color().
func (renderer Renderer) SetColor(part RenderPart, color Color) {
	C.pango_renderer_set_color(renderer.native(), C.PangoRenderPart(part), color.native())
}

// SetMatrix is a wrapper around pango_renderer_set_matrix().
func (renderer Renderer) SetMatrix(matrix Matrix) {
	C.pango_renderer_set_matrix(renderer.native(), matrix.native())
}

type Glyph uint32
type GlyphUnit int32
type Alignment int

const (
	AlignmentLeft   Alignment = 0
	AlignmentCenter           = 1
	AlignmentRight            = 2
)

type AttrType int

const (
	AttrTypeInvalid            AttrType = 0
	AttrTypeLanguage                    = 1
	AttrTypeFamily                      = 2
	AttrTypeStyle                       = 3
	AttrTypeWeight                      = 4
	AttrTypeVariant                     = 5
	AttrTypeStretch                     = 6
	AttrTypeSize                        = 7
	AttrTypeFontDesc                    = 8
	AttrTypeForeground                  = 9
	AttrTypeBackground                  = 10
	AttrTypeUnderline                   = 11
	AttrTypeStrikethrough               = 12
	AttrTypeRise                        = 13
	AttrTypeShape                       = 14
	AttrTypeScale                       = 15
	AttrTypeFallback                    = 16
	AttrTypeLetterSpacing               = 17
	AttrTypeUnderlineColor              = 18
	AttrTypeStrikethroughColor          = 19
	AttrTypeAbsoluteSize                = 20
	AttrTypeGravity                     = 21
	AttrTypeGravityHint                 = 22
	AttrTypeFontFeatures                = 23
	AttrTypeForegroundAlpha             = 24
	AttrTypeBackgroundAlpha             = 25
)

type BidiType int

const (
	BidiTypeL   BidiType = 0
	BidiTypeLre          = 1
	BidiTypeLro          = 2
	BidiTypeR            = 3
	BidiTypeAl           = 4
	BidiTypeRle          = 5
	BidiTypeRlo          = 6
	BidiTypePdf          = 7
	BidiTypeEn           = 8
	BidiTypeEs           = 9
	BidiTypeEt           = 10
	BidiTypeAn           = 11
	BidiTypeCs           = 12
	BidiTypeNsm          = 13
	BidiTypeBn           = 14
	BidiTypeB            = 15
	BidiTypeS            = 16
	BidiTypeWs           = 17
	BidiTypeOn           = 18
)

type CoverageLevel int

const (
	CoverageLevelNone        CoverageLevel = 0
	CoverageLevelFallback                  = 1
	CoverageLevelApproximate               = 2
	CoverageLevelExact                     = 3
)

type Direction int

const (
	DirectionLtr     Direction = 0
	DirectionRtl               = 1
	DirectionTtbLtr            = 2
	DirectionTtbRtl            = 3
	DirectionWeakLtr           = 4
	DirectionWeakRtl           = 5
	DirectionNeutral           = 6
)

type EllipsizeMode int

const (
	EllipsizeModeNone   EllipsizeMode = 0
	EllipsizeModeStart                = 1
	EllipsizeModeMiddle               = 2
	EllipsizeModeEnd                  = 3
)

type Gravity int

const (
	GravitySouth Gravity = 0
	GravityEast          = 1
	GravityNorth         = 2
	GravityWest          = 3
	GravityAuto          = 4
)

type GravityHint int

const (
	GravityHintNatural GravityHint = 0
	GravityHintStrong              = 1
	GravityHintLine                = 2
)

type RenderPart int

const (
	RenderPartForeground    RenderPart = 0
	RenderPartBackground               = 1
	RenderPartUnderline                = 2
	RenderPartStrikethrough            = 3
)

type Script int

const (
	ScriptInvalidCode          Script = -1
	ScriptCommon                      = 0
	ScriptInherited                   = 1
	ScriptArabic                      = 2
	ScriptArmenian                    = 3
	ScriptBengali                     = 4
	ScriptBopomofo                    = 5
	ScriptCherokee                    = 6
	ScriptCoptic                      = 7
	ScriptCyrillic                    = 8
	ScriptDeseret                     = 9
	ScriptDevanagari                  = 10
	ScriptEthiopic                    = 11
	ScriptGeorgian                    = 12
	ScriptGothic                      = 13
	ScriptGreek                       = 14
	ScriptGujarati                    = 15
	ScriptGurmukhi                    = 16
	ScriptHan                         = 17
	ScriptHangul                      = 18
	ScriptHebrew                      = 19
	ScriptHiragana                    = 20
	ScriptKannada                     = 21
	ScriptKatakana                    = 22
	ScriptKhmer                       = 23
	ScriptLao                         = 24
	ScriptLatin                       = 25
	ScriptMalayalam                   = 26
	ScriptMongolian                   = 27
	ScriptMyanmar                     = 28
	ScriptOgham                       = 29
	ScriptOldItalic                   = 30
	ScriptOriya                       = 31
	ScriptRunic                       = 32
	ScriptSinhala                     = 33
	ScriptSyriac                      = 34
	ScriptTamil                       = 35
	ScriptTelugu                      = 36
	ScriptThaana                      = 37
	ScriptThai                        = 38
	ScriptTibetan                     = 39
	ScriptCanadianAboriginal          = 40
	ScriptYi                          = 41
	ScriptTagalog                     = 42
	ScriptHanunoo                     = 43
	ScriptBuhid                       = 44
	ScriptTagbanwa                    = 45
	ScriptBraille                     = 46
	ScriptCypriot                     = 47
	ScriptLimbu                       = 48
	ScriptOsmanya                     = 49
	ScriptShavian                     = 50
	ScriptLinearB                     = 51
	ScriptTaiLe                       = 52
	ScriptUgaritic                    = 53
	ScriptNewTaiLue                   = 54
	ScriptBuginese                    = 55
	ScriptGlagolitic                  = 56
	ScriptTifinagh                    = 57
	ScriptSylotiNagri                 = 58
	ScriptOldPersian                  = 59
	ScriptKharoshthi                  = 60
	ScriptUnknown                     = 61
	ScriptBalinese                    = 62
	ScriptCuneiform                   = 63
	ScriptPhoenician                  = 64
	ScriptPhagsPa                     = 65
	ScriptNko                         = 66
	ScriptKayahLi                     = 67
	ScriptLepcha                      = 68
	ScriptRejang                      = 69
	ScriptSundanese                   = 70
	ScriptSaurashtra                  = 71
	ScriptCham                        = 72
	ScriptOlChiki                     = 73
	ScriptVai                         = 74
	ScriptCarian                      = 75
	ScriptLycian                      = 76
	ScriptLydian                      = 77
	ScriptBatak                       = 78
	ScriptBrahmi                      = 79
	ScriptMandaic                     = 80
	ScriptChakma                      = 81
	ScriptMeroiticCursive             = 82
	ScriptMeroiticHieroglyphs         = 83
	ScriptMiao                        = 84
	ScriptSharada                     = 85
	ScriptSoraSompeng                 = 86
	ScriptTakri                       = 87
	ScriptBassaVah                    = 88
	ScriptCaucasianAlbanian           = 89
	ScriptDuployan                    = 90
	ScriptElbasan                     = 91
	ScriptGrantha                     = 92
	ScriptKhojki                      = 93
	ScriptKhudawadi                   = 94
	ScriptLinearA                     = 95
	ScriptMahajani                    = 96
	ScriptManichaean                  = 97
	ScriptMendeKikakui                = 98
	ScriptModi                        = 99
	ScriptMro                         = 100
	ScriptNabataean                   = 101
	ScriptOldNorthArabian             = 102
	ScriptOldPermic                   = 103
	ScriptPahawhHmong                 = 104
	ScriptPalmyrene                   = 105
	ScriptPauCinHau                   = 106
	ScriptPsalterPahlavi              = 107
	ScriptSiddham                     = 108
	ScriptTirhuta                     = 109
	ScriptWarangCiti                  = 110
	ScriptAhom                        = 111
	ScriptAnatolianHieroglyphs        = 112
	ScriptHatran                      = 113
	ScriptMultani                     = 114
	ScriptOldHungarian                = 115
	ScriptSignwriting                 = 116
)

type Stretch int

const (
	StretchUltraCondensed Stretch = 0
	StretchExtraCondensed         = 1
	StretchCondensed              = 2
	StretchSemiCondensed          = 3
	StretchNormal                 = 4
	StretchSemiExpanded           = 5
	StretchExpanded               = 6
	StretchExtraExpanded          = 7
	StretchUltraExpanded          = 8
)

type Style int

const (
	StyleNormal  Style = 0
	StyleOblique       = 1
	StyleItalic        = 2
)

type TabAlign int

const (
	TabAlignLeft TabAlign = 0
)

type Underline int

const (
	UnderlineNone   Underline = 0
	UnderlineSingle           = 1
	UnderlineDouble           = 2
	UnderlineLow              = 3
	UnderlineError            = 4
)

type Variant int

const (
	VariantNormal    Variant = 0
	VariantSmallCaps         = 1
)

type Weight int

const (
	WeightThin       Weight = 100
	WeightUltralight        = 200
	WeightLight             = 300
	WeightSemilight         = 350
	WeightBook              = 380
	WeightNormal            = 400
	WeightMedium            = 500
	WeightSemibold          = 600
	WeightBold              = 700
	WeightUltrabold         = 800
	WeightHeavy             = 900
	WeightUltraheavy        = 1000
)

type WrapMode int

const (
	WrapModeWord     WrapMode = 0
	WrapModeChar              = 1
	WrapModeWordChar          = 2
)

type FontMask int

const (
	FontMaskFamily  FontMask = 1
	FontMaskStyle            = 2
	FontMaskVariant          = 4
	FontMaskWeight           = 8
	FontMaskStretch          = 16
	FontMaskSize             = 32
	FontMaskGravity          = 64
)
