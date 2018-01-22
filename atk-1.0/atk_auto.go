package atk

/*
#cgo pkg-config: atk
#include <atk/atk.h>
#include <stdlib.h>
*/
import "C"
import "github.com/linuxdeepin/go-gir/glib-2.0"
import "github.com/linuxdeepin/go-gir/gobject-2.0"
import "github.com/linuxdeepin/go-gir/util"
import "unsafe"

// Struct Attribute
type Attribute struct {
	Ptr unsafe.Pointer
}

func (v Attribute) native() *C.AtkAttribute {
	return (*C.AtkAttribute)(v.Ptr)
}
func wrapAttribute(p *C.AtkAttribute) Attribute {
	return Attribute{Ptr: unsafe.Pointer(p)}
}
func WrapAttribute(p unsafe.Pointer) Attribute {
	return Attribute{Ptr: p}
}
func (v Attribute) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAttribute(p unsafe.Pointer) interface{} {
	return WrapAttribute(p)
}

// Struct KeyEventStruct
type KeyEventStruct struct {
	Ptr unsafe.Pointer
}

func (v KeyEventStruct) native() *C.AtkKeyEventStruct {
	return (*C.AtkKeyEventStruct)(v.Ptr)
}
func wrapKeyEventStruct(p *C.AtkKeyEventStruct) KeyEventStruct {
	return KeyEventStruct{Ptr: unsafe.Pointer(p)}
}
func WrapKeyEventStruct(p unsafe.Pointer) KeyEventStruct {
	return KeyEventStruct{Ptr: p}
}
func (v KeyEventStruct) IsNil() bool {
	return v.Ptr == nil
}
func IWrapKeyEventStruct(p unsafe.Pointer) interface{} {
	return WrapKeyEventStruct(p)
}

// Struct PropertyValues
type PropertyValues struct {
	Ptr unsafe.Pointer
}

func (v PropertyValues) native() *C.AtkPropertyValues {
	return (*C.AtkPropertyValues)(v.Ptr)
}
func wrapPropertyValues(p *C.AtkPropertyValues) PropertyValues {
	return PropertyValues{Ptr: unsafe.Pointer(p)}
}
func WrapPropertyValues(p unsafe.Pointer) PropertyValues {
	return PropertyValues{Ptr: p}
}
func (v PropertyValues) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPropertyValues(p unsafe.Pointer) interface{} {
	return WrapPropertyValues(p)
}

// Struct Rectangle
type Rectangle struct {
	Ptr unsafe.Pointer
}

func (v Rectangle) native() *C.AtkRectangle {
	return (*C.AtkRectangle)(v.Ptr)
}
func wrapRectangle(p *C.AtkRectangle) Rectangle {
	return Rectangle{Ptr: unsafe.Pointer(p)}
}
func WrapRectangle(p unsafe.Pointer) Rectangle {
	return Rectangle{Ptr: p}
}
func (v Rectangle) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRectangle(p unsafe.Pointer) interface{} {
	return WrapRectangle(p)
}

// Struct TextRange
type TextRange struct {
	Ptr unsafe.Pointer
}

func (v TextRange) native() *C.AtkTextRange {
	return (*C.AtkTextRange)(v.Ptr)
}
func wrapTextRange(p *C.AtkTextRange) TextRange {
	return TextRange{Ptr: unsafe.Pointer(p)}
}
func WrapTextRange(p unsafe.Pointer) TextRange {
	return TextRange{Ptr: p}
}
func (v TextRange) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextRange(p unsafe.Pointer) interface{} {
	return WrapTextRange(p)
}

// Struct TextRectangle
type TextRectangle struct {
	Ptr unsafe.Pointer
}

func (v TextRectangle) native() *C.AtkTextRectangle {
	return (*C.AtkTextRectangle)(v.Ptr)
}
func wrapTextRectangle(p *C.AtkTextRectangle) TextRectangle {
	return TextRectangle{Ptr: unsafe.Pointer(p)}
}
func WrapTextRectangle(p unsafe.Pointer) TextRectangle {
	return TextRectangle{Ptr: p}
}
func (v TextRectangle) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTextRectangle(p unsafe.Pointer) interface{} {
	return WrapTextRectangle(p)
}

// Interface Component
type Component struct {
	ComponentIface
	Ptr unsafe.Pointer
}
type ComponentIface struct{}

func (v *ComponentIface) native() *C.AtkComponent {
	return (*C.AtkComponent)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapComponent(p *C.AtkComponent) Component {
	return Component{Ptr: unsafe.Pointer(p)}
}
func WrapComponent(p unsafe.Pointer) Component {
	return Component{Ptr: p}
}
func (v Component) IsNil() bool {
	return v.Ptr == nil
}
func IWrapComponent(p unsafe.Pointer) interface{} {
	return WrapComponent(p)
}
func (v Component) GetType() gobject.Type {
	return gobject.Type(C.atk_component_get_type())
}
func (v Component) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapComponent(unsafe.Pointer(ptr)), nil
	}
}

// Contains is a wrapper around atk_component_contains().
func (component *ComponentIface) Contains(x int, y int, coord_type CoordType) bool {
	ret0 := C.atk_component_contains(component.native(), C.gint(x), C.gint(y), C.AtkCoordType(coord_type))
	return util.Int2Bool(int(ret0))
}

// GetAlpha is a wrapper around atk_component_get_alpha().
func (component *ComponentIface) GetAlpha() float64 {
	ret0 := C.atk_component_get_alpha(component.native())
	return float64(ret0)
}

// GetExtents is a wrapper around atk_component_get_extents().
func (component *ComponentIface) GetExtents(coord_type CoordType) (int, int, int, int) {
	var x0 C.gint
	var y0 C.gint
	var width0 C.gint
	var height0 C.gint
	C.atk_component_get_extents(component.native(), &x0, &y0, &width0, &height0, C.AtkCoordType(coord_type))
	return int(x0), int(y0), int(width0), int(height0)
}

// GetLayer is a wrapper around atk_component_get_layer().
func (component *ComponentIface) GetLayer() Layer {
	ret0 := C.atk_component_get_layer(component.native())
	return Layer(ret0)
}

// GetMdiZorder is a wrapper around atk_component_get_mdi_zorder().
func (component *ComponentIface) GetMdiZorder() int {
	ret0 := C.atk_component_get_mdi_zorder(component.native())
	return int(ret0)
}

// GrabFocus is a wrapper around atk_component_grab_focus().
func (component *ComponentIface) GrabFocus() bool {
	ret0 := C.atk_component_grab_focus(component.native())
	return util.Int2Bool(int(ret0))
}

// RefAccessibleAtPoint is a wrapper around atk_component_ref_accessible_at_point().
func (component *ComponentIface) RefAccessibleAtPoint(x int, y int, coord_type CoordType) Object {
	ret0 := C.atk_component_ref_accessible_at_point(component.native(), C.gint(x), C.gint(y), C.AtkCoordType(coord_type))
	return wrapObject(ret0)
}

// SetExtents is a wrapper around atk_component_set_extents().
func (component *ComponentIface) SetExtents(x int, y int, width int, height int, coord_type CoordType) bool {
	ret0 := C.atk_component_set_extents(component.native(), C.gint(x), C.gint(y), C.gint(width), C.gint(height), C.AtkCoordType(coord_type))
	return util.Int2Bool(int(ret0))
}

// SetPosition is a wrapper around atk_component_set_position().
func (component *ComponentIface) SetPosition(x int, y int, coord_type CoordType) bool {
	ret0 := C.atk_component_set_position(component.native(), C.gint(x), C.gint(y), C.AtkCoordType(coord_type))
	return util.Int2Bool(int(ret0))
}

// SetSize is a wrapper around atk_component_set_size().
func (component *ComponentIface) SetSize(width int, height int) bool {
	ret0 := C.atk_component_set_size(component.native(), C.gint(width), C.gint(height))
	return util.Int2Bool(int(ret0))
}

// Interface EditableText
type EditableText struct {
	EditableTextIface
	Ptr unsafe.Pointer
}
type EditableTextIface struct{}

func (v *EditableTextIface) native() *C.AtkEditableText {
	return (*C.AtkEditableText)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapEditableText(p *C.AtkEditableText) EditableText {
	return EditableText{Ptr: unsafe.Pointer(p)}
}
func WrapEditableText(p unsafe.Pointer) EditableText {
	return EditableText{Ptr: p}
}
func (v EditableText) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEditableText(p unsafe.Pointer) interface{} {
	return WrapEditableText(p)
}
func (v EditableText) GetType() gobject.Type {
	return gobject.Type(C.atk_editable_text_get_type())
}
func (v EditableText) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEditableText(unsafe.Pointer(ptr)), nil
	}
}

// CopyText is a wrapper around atk_editable_text_copy_text().
func (text *EditableTextIface) CopyText(start_pos int, end_pos int) {
	C.atk_editable_text_copy_text(text.native(), C.gint(start_pos), C.gint(end_pos))
}

// CutText is a wrapper around atk_editable_text_cut_text().
func (text *EditableTextIface) CutText(start_pos int, end_pos int) {
	C.atk_editable_text_cut_text(text.native(), C.gint(start_pos), C.gint(end_pos))
}

// DeleteText is a wrapper around atk_editable_text_delete_text().
func (text *EditableTextIface) DeleteText(start_pos int, end_pos int) {
	C.atk_editable_text_delete_text(text.native(), C.gint(start_pos), C.gint(end_pos))
}

// PasteText is a wrapper around atk_editable_text_paste_text().
func (text *EditableTextIface) PasteText(position int) {
	C.atk_editable_text_paste_text(text.native(), C.gint(position))
}

// SetTextContents is a wrapper around atk_editable_text_set_text_contents().
func (text *EditableTextIface) SetTextContents(string string) {
	string0 := (*C.gchar)(C.CString(string))
	C.atk_editable_text_set_text_contents(text.native(), string0)
	C.free(unsafe.Pointer(string0))
}

// Struct Range
type Range struct {
	Ptr unsafe.Pointer
}

func (v Range) native() *C.AtkRange {
	return (*C.AtkRange)(v.Ptr)
}
func wrapRange(p *C.AtkRange) Range {
	return Range{Ptr: unsafe.Pointer(p)}
}
func WrapRange(p unsafe.Pointer) Range {
	return Range{Ptr: p}
}
func (v Range) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRange(p unsafe.Pointer) interface{} {
	return WrapRange(p)
}

// RangeNew is a wrapper around atk_range_new().
func RangeNew(lower_limit float64, upper_limit float64, description string) Range {
	description0 := (*C.gchar)(C.CString(description))
	ret0 := C.atk_range_new(C.gdouble(lower_limit), C.gdouble(upper_limit), description0)
	C.free(unsafe.Pointer(description0))
	return wrapRange(ret0)
}

// Copy is a wrapper around atk_range_copy().
func (src Range) Copy() Range {
	ret0 := C.atk_range_copy(src.native())
	return wrapRange(ret0)
}

// Free is a wrapper around atk_range_free().
func (range_ Range) Free() {
	C.atk_range_free(range_.native())
}

// GetDescription is a wrapper around atk_range_get_description().
func (range_ Range) GetDescription() string {
	ret0 := C.atk_range_get_description(range_.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLowerLimit is a wrapper around atk_range_get_lower_limit().
func (range_ Range) GetLowerLimit() float64 {
	ret0 := C.atk_range_get_lower_limit(range_.native())
	return float64(ret0)
}

// GetUpperLimit is a wrapper around atk_range_get_upper_limit().
func (range_ Range) GetUpperLimit() float64 {
	ret0 := C.atk_range_get_upper_limit(range_.native())
	return float64(ret0)
}

// Interface Action
type Action struct {
	ActionIface
	Ptr unsafe.Pointer
}
type ActionIface struct{}

func (v *ActionIface) native() *C.AtkAction {
	return (*C.AtkAction)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapAction(p *C.AtkAction) Action {
	return Action{Ptr: unsafe.Pointer(p)}
}
func WrapAction(p unsafe.Pointer) Action {
	return Action{Ptr: p}
}
func (v Action) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAction(p unsafe.Pointer) interface{} {
	return WrapAction(p)
}
func (v Action) GetType() gobject.Type {
	return gobject.Type(C.atk_action_get_type())
}
func (v Action) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAction(unsafe.Pointer(ptr)), nil
	}
}

// DoAction is a wrapper around atk_action_do_action().
func (action *ActionIface) DoAction(i int) bool {
	ret0 := C.atk_action_do_action(action.native(), C.gint(i))
	return util.Int2Bool(int(ret0))
}

// GetDescription is a wrapper around atk_action_get_description().
func (action *ActionIface) GetDescription(i int) string {
	ret0 := C.atk_action_get_description(action.native(), C.gint(i))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetKeybinding is a wrapper around atk_action_get_keybinding().
func (action *ActionIface) GetKeybinding(i int) string {
	ret0 := C.atk_action_get_keybinding(action.native(), C.gint(i))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLocalizedName is a wrapper around atk_action_get_localized_name().
func (action *ActionIface) GetLocalizedName(i int) string {
	ret0 := C.atk_action_get_localized_name(action.native(), C.gint(i))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNActions is a wrapper around atk_action_get_n_actions().
func (action *ActionIface) GetNActions() int {
	ret0 := C.atk_action_get_n_actions(action.native())
	return int(ret0)
}

// GetName is a wrapper around atk_action_get_name().
func (action *ActionIface) GetName(i int) string {
	ret0 := C.atk_action_get_name(action.native(), C.gint(i))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetDescription is a wrapper around atk_action_set_description().
func (action *ActionIface) SetDescription(i int, desc string) bool {
	desc0 := (*C.gchar)(C.CString(desc))
	ret0 := C.atk_action_set_description(action.native(), C.gint(i), desc0)
	C.free(unsafe.Pointer(desc0))
	return util.Int2Bool(int(ret0))
}

// Object Object
type Object struct {
	gobject.Object
}

func (v Object) native() *C.AtkObject {
	return (*C.AtkObject)(v.Ptr)
}
func wrapObject(p *C.AtkObject) (v Object) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapObject(p unsafe.Pointer) (v Object) {
	v.Ptr = p
	return
}
func (v Object) IsNil() bool {
	return v.Ptr == nil
}
func IWrapObject(p unsafe.Pointer) interface{} {
	return WrapObject(p)
}
func (v Object) GetType() gobject.Type {
	return gobject.Type(C.atk_object_get_type())
}
func (v Object) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapObject(unsafe.Pointer(ptr)), nil
	}
}

// AddRelationship is a wrapper around atk_object_add_relationship().
func (object Object) AddRelationship(relationship RelationType, target Object) bool {
	ret0 := C.atk_object_add_relationship(object.native(), C.AtkRelationType(relationship), target.native())
	return util.Int2Bool(int(ret0))
}

// GetDescription is a wrapper around atk_object_get_description().
func (accessible Object) GetDescription() string {
	ret0 := C.atk_object_get_description(accessible.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetIndexInParent is a wrapper around atk_object_get_index_in_parent().
func (accessible Object) GetIndexInParent() int {
	ret0 := C.atk_object_get_index_in_parent(accessible.native())
	return int(ret0)
}

// GetNAccessibleChildren is a wrapper around atk_object_get_n_accessible_children().
func (accessible Object) GetNAccessibleChildren() int {
	ret0 := C.atk_object_get_n_accessible_children(accessible.native())
	return int(ret0)
}

// GetName is a wrapper around atk_object_get_name().
func (accessible Object) GetName() string {
	ret0 := C.atk_object_get_name(accessible.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetObjectLocale is a wrapper around atk_object_get_object_locale().
func (accessible Object) GetObjectLocale() string {
	ret0 := C.atk_object_get_object_locale(accessible.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetParent is a wrapper around atk_object_get_parent().
func (accessible Object) GetParent() Object {
	ret0 := C.atk_object_get_parent(accessible.native())
	return wrapObject(ret0)
}

// GetRole is a wrapper around atk_object_get_role().
func (accessible Object) GetRole() Role {
	ret0 := C.atk_object_get_role(accessible.native())
	return Role(ret0)
}

// Initialize is a wrapper around atk_object_initialize().
func (accessible Object) Initialize(data unsafe.Pointer) {
	C.atk_object_initialize(accessible.native(), C.gpointer(data))
}

// NotifyStateChange is a wrapper around atk_object_notify_state_change().
func (accessible Object) NotifyStateChange(state State, value bool) {
	C.atk_object_notify_state_change(accessible.native(), C.AtkState(state), C.gboolean(util.Bool2Int(value)))
}

// PeekParent is a wrapper around atk_object_peek_parent().
func (accessible Object) PeekParent() Object {
	ret0 := C.atk_object_peek_parent(accessible.native())
	return wrapObject(ret0)
}

// RefAccessibleChild is a wrapper around atk_object_ref_accessible_child().
func (accessible Object) RefAccessibleChild(i int) Object {
	ret0 := C.atk_object_ref_accessible_child(accessible.native(), C.gint(i))
	return wrapObject(ret0)
}

// RefRelationSet is a wrapper around atk_object_ref_relation_set().
func (accessible Object) RefRelationSet() RelationSet {
	ret0 := C.atk_object_ref_relation_set(accessible.native())
	return wrapRelationSet(ret0)
}

// RefStateSet is a wrapper around atk_object_ref_state_set().
func (accessible Object) RefStateSet() StateSet {
	ret0 := C.atk_object_ref_state_set(accessible.native())
	return wrapStateSet(ret0)
}

// RemoveRelationship is a wrapper around atk_object_remove_relationship().
func (object Object) RemoveRelationship(relationship RelationType, target Object) bool {
	ret0 := C.atk_object_remove_relationship(object.native(), C.AtkRelationType(relationship), target.native())
	return util.Int2Bool(int(ret0))
}

// SetDescription is a wrapper around atk_object_set_description().
func (accessible Object) SetDescription(description string) {
	description0 := (*C.gchar)(C.CString(description))
	C.atk_object_set_description(accessible.native(), description0)
	C.free(unsafe.Pointer(description0))
}

// SetName is a wrapper around atk_object_set_name().
func (accessible Object) SetName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.atk_object_set_name(accessible.native(), name0)
	C.free(unsafe.Pointer(name0))
}

// SetParent is a wrapper around atk_object_set_parent().
func (accessible Object) SetParent(parent Object) {
	C.atk_object_set_parent(accessible.native(), parent.native())
}

// SetRole is a wrapper around atk_object_set_role().
func (accessible Object) SetRole(role Role) {
	C.atk_object_set_role(accessible.native(), C.AtkRole(role))
}

// Object RelationSet
type RelationSet struct {
	gobject.Object
}

func (v RelationSet) native() *C.AtkRelationSet {
	return (*C.AtkRelationSet)(v.Ptr)
}
func wrapRelationSet(p *C.AtkRelationSet) (v RelationSet) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRelationSet(p unsafe.Pointer) (v RelationSet) {
	v.Ptr = p
	return
}
func (v RelationSet) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRelationSet(p unsafe.Pointer) interface{} {
	return WrapRelationSet(p)
}
func (v RelationSet) GetType() gobject.Type {
	return gobject.Type(C.atk_relation_set_get_type())
}
func (v RelationSet) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRelationSet(unsafe.Pointer(ptr)), nil
	}
}

// RelationSetNew is a wrapper around atk_relation_set_new().
func RelationSetNew() RelationSet {
	ret0 := C.atk_relation_set_new()
	return wrapRelationSet(ret0)
}

// Add is a wrapper around atk_relation_set_add().
func (set RelationSet) Add(relation Relation) {
	C.atk_relation_set_add(set.native(), relation.native())
}

// AddRelationByType is a wrapper around atk_relation_set_add_relation_by_type().
func (set RelationSet) AddRelationByType(relationship RelationType, target Object) {
	C.atk_relation_set_add_relation_by_type(set.native(), C.AtkRelationType(relationship), target.native())
}

// Contains is a wrapper around atk_relation_set_contains().
func (set RelationSet) Contains(relationship RelationType) bool {
	ret0 := C.atk_relation_set_contains(set.native(), C.AtkRelationType(relationship))
	return util.Int2Bool(int(ret0))
}

// ContainsTarget is a wrapper around atk_relation_set_contains_target().
func (set RelationSet) ContainsTarget(relationship RelationType, target Object) bool {
	ret0 := C.atk_relation_set_contains_target(set.native(), C.AtkRelationType(relationship), target.native())
	return util.Int2Bool(int(ret0))
}

// GetNRelations is a wrapper around atk_relation_set_get_n_relations().
func (set RelationSet) GetNRelations() int {
	ret0 := C.atk_relation_set_get_n_relations(set.native())
	return int(ret0)
}

// GetRelation is a wrapper around atk_relation_set_get_relation().
func (set RelationSet) GetRelation(i int) Relation {
	ret0 := C.atk_relation_set_get_relation(set.native(), C.gint(i))
	return wrapRelation(ret0)
}

// GetRelationByType is a wrapper around atk_relation_set_get_relation_by_type().
func (set RelationSet) GetRelationByType(relationship RelationType) Relation {
	ret0 := C.atk_relation_set_get_relation_by_type(set.native(), C.AtkRelationType(relationship))
	return wrapRelation(ret0)
}

// Remove is a wrapper around atk_relation_set_remove().
func (set RelationSet) Remove(relation Relation) {
	C.atk_relation_set_remove(set.native(), relation.native())
}

// Object Relation
type Relation struct {
	gobject.Object
}

func (v Relation) native() *C.AtkRelation {
	return (*C.AtkRelation)(v.Ptr)
}
func wrapRelation(p *C.AtkRelation) (v Relation) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRelation(p unsafe.Pointer) (v Relation) {
	v.Ptr = p
	return
}
func (v Relation) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRelation(p unsafe.Pointer) interface{} {
	return WrapRelation(p)
}
func (v Relation) GetType() gobject.Type {
	return gobject.Type(C.atk_relation_get_type())
}
func (v Relation) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRelation(unsafe.Pointer(ptr)), nil
	}
}

// RelationNew is a wrapper around atk_relation_new().
func RelationNew(targets []Object, relationship RelationType) Relation {
	targets0 := make([]*C.AtkObject, len(targets))
	for idx, elemG := range targets {
		targets0[idx] = elemG.native()
	}
	var targets0Ptr **C.AtkObject
	if len(targets0) > 0 {
		targets0Ptr = &targets0[0]
	}
	ret0 := C.atk_relation_new(targets0Ptr, C.gint(len(targets)), C.AtkRelationType(relationship))
	return wrapRelation(ret0)
}

// AddTarget is a wrapper around atk_relation_add_target().
func (relation Relation) AddTarget(target Object) {
	C.atk_relation_add_target(relation.native(), target.native())
}

// GetRelationType is a wrapper around atk_relation_get_relation_type().
func (relation Relation) GetRelationType() RelationType {
	ret0 := C.atk_relation_get_relation_type(relation.native())
	return RelationType(ret0)
}

// RemoveTarget is a wrapper around atk_relation_remove_target().
func (relation Relation) RemoveTarget(target Object) bool {
	ret0 := C.atk_relation_remove_target(relation.native(), target.native())
	return util.Int2Bool(int(ret0))
}

// Interface Document
type Document struct {
	DocumentIface
	Ptr unsafe.Pointer
}
type DocumentIface struct{}

func (v *DocumentIface) native() *C.AtkDocument {
	return (*C.AtkDocument)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapDocument(p *C.AtkDocument) Document {
	return Document{Ptr: unsafe.Pointer(p)}
}
func WrapDocument(p unsafe.Pointer) Document {
	return Document{Ptr: p}
}
func (v Document) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDocument(p unsafe.Pointer) interface{} {
	return WrapDocument(p)
}
func (v Document) GetType() gobject.Type {
	return gobject.Type(C.atk_document_get_type())
}
func (v Document) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDocument(unsafe.Pointer(ptr)), nil
	}
}

// GetAttributeValue is a wrapper around atk_document_get_attribute_value().
func (document *DocumentIface) GetAttributeValue(attribute_name string) string {
	attribute_name0 := (*C.gchar)(C.CString(attribute_name))
	ret0 := C.atk_document_get_attribute_value(document.native(), attribute_name0)
	C.free(unsafe.Pointer(attribute_name0))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetCurrentPageNumber is a wrapper around atk_document_get_current_page_number().
func (document *DocumentIface) GetCurrentPageNumber() int {
	ret0 := C.atk_document_get_current_page_number(document.native())
	return int(ret0)
}

// GetPageCount is a wrapper around atk_document_get_page_count().
func (document *DocumentIface) GetPageCount() int {
	ret0 := C.atk_document_get_page_count(document.native())
	return int(ret0)
}

// SetAttributeValue is a wrapper around atk_document_set_attribute_value().
func (document *DocumentIface) SetAttributeValue(attribute_name string, attribute_value string) bool {
	attribute_name0 := (*C.gchar)(C.CString(attribute_name))
	attribute_value0 := (*C.gchar)(C.CString(attribute_value))
	ret0 := C.atk_document_set_attribute_value(document.native(), attribute_name0, attribute_value0)
	C.free(unsafe.Pointer(attribute_name0))
	C.free(unsafe.Pointer(attribute_value0))
	return util.Int2Bool(int(ret0))
}

// Interface HyperlinkImpl
type HyperlinkImpl struct {
	HyperlinkImplIface
	Ptr unsafe.Pointer
}
type HyperlinkImplIface struct{}

func (v *HyperlinkImplIface) native() *C.AtkHyperlinkImpl {
	return (*C.AtkHyperlinkImpl)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapHyperlinkImpl(p *C.AtkHyperlinkImpl) HyperlinkImpl {
	return HyperlinkImpl{Ptr: unsafe.Pointer(p)}
}
func WrapHyperlinkImpl(p unsafe.Pointer) HyperlinkImpl {
	return HyperlinkImpl{Ptr: p}
}
func (v HyperlinkImpl) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHyperlinkImpl(p unsafe.Pointer) interface{} {
	return WrapHyperlinkImpl(p)
}
func (v HyperlinkImpl) GetType() gobject.Type {
	return gobject.Type(C.atk_hyperlink_impl_get_type())
}
func (v HyperlinkImpl) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHyperlinkImpl(unsafe.Pointer(ptr)), nil
	}
}

// GetHyperlink is a wrapper around atk_hyperlink_impl_get_hyperlink().
func (impl *HyperlinkImplIface) GetHyperlink() Hyperlink {
	ret0 := C.atk_hyperlink_impl_get_hyperlink(impl.native())
	return wrapHyperlink(ret0)
}

// Object Hyperlink
type Hyperlink struct {
	ActionIface
	gobject.Object
}

func (v Hyperlink) native() *C.AtkHyperlink {
	return (*C.AtkHyperlink)(v.Ptr)
}
func wrapHyperlink(p *C.AtkHyperlink) (v Hyperlink) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapHyperlink(p unsafe.Pointer) (v Hyperlink) {
	v.Ptr = p
	return
}
func (v Hyperlink) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHyperlink(p unsafe.Pointer) interface{} {
	return WrapHyperlink(p)
}
func (v Hyperlink) GetType() gobject.Type {
	return gobject.Type(C.atk_hyperlink_get_type())
}
func (v Hyperlink) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHyperlink(unsafe.Pointer(ptr)), nil
	}
}
func (v Hyperlink) Action() Action {
	return WrapAction(v.Ptr)
}

// GetEndIndex is a wrapper around atk_hyperlink_get_end_index().
func (link_ Hyperlink) GetEndIndex() int {
	ret0 := C.atk_hyperlink_get_end_index(link_.native())
	return int(ret0)
}

// GetNAnchors is a wrapper around atk_hyperlink_get_n_anchors().
func (link_ Hyperlink) GetNAnchors() int {
	ret0 := C.atk_hyperlink_get_n_anchors(link_.native())
	return int(ret0)
}

// GetObject is a wrapper around atk_hyperlink_get_object().
func (link_ Hyperlink) GetObject(i int) Object {
	ret0 := C.atk_hyperlink_get_object(link_.native(), C.gint(i))
	return wrapObject(ret0)
}

// GetStartIndex is a wrapper around atk_hyperlink_get_start_index().
func (link_ Hyperlink) GetStartIndex() int {
	ret0 := C.atk_hyperlink_get_start_index(link_.native())
	return int(ret0)
}

// GetUri is a wrapper around atk_hyperlink_get_uri().
func (link_ Hyperlink) GetUri(i int) string {
	ret0 := C.atk_hyperlink_get_uri(link_.native(), C.gint(i))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// IsInline is a wrapper around atk_hyperlink_is_inline().
func (link_ Hyperlink) IsInline() bool {
	ret0 := C.atk_hyperlink_is_inline(link_.native())
	return util.Int2Bool(int(ret0))
}

// IsValid is a wrapper around atk_hyperlink_is_valid().
func (link_ Hyperlink) IsValid() bool {
	ret0 := C.atk_hyperlink_is_valid(link_.native())
	return util.Int2Bool(int(ret0))
}

// Interface Hypertext
type Hypertext struct {
	HypertextIface
	Ptr unsafe.Pointer
}
type HypertextIface struct{}

func (v *HypertextIface) native() *C.AtkHypertext {
	return (*C.AtkHypertext)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapHypertext(p *C.AtkHypertext) Hypertext {
	return Hypertext{Ptr: unsafe.Pointer(p)}
}
func WrapHypertext(p unsafe.Pointer) Hypertext {
	return Hypertext{Ptr: p}
}
func (v Hypertext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapHypertext(p unsafe.Pointer) interface{} {
	return WrapHypertext(p)
}
func (v Hypertext) GetType() gobject.Type {
	return gobject.Type(C.atk_hypertext_get_type())
}
func (v Hypertext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapHypertext(unsafe.Pointer(ptr)), nil
	}
}

// GetLink is a wrapper around atk_hypertext_get_link().
func (hypertext *HypertextIface) GetLink(link_index int) Hyperlink {
	ret0 := C.atk_hypertext_get_link(hypertext.native(), C.gint(link_index))
	return wrapHyperlink(ret0)
}

// GetLinkIndex is a wrapper around atk_hypertext_get_link_index().
func (hypertext *HypertextIface) GetLinkIndex(char_index int) int {
	ret0 := C.atk_hypertext_get_link_index(hypertext.native(), C.gint(char_index))
	return int(ret0)
}

// GetNLinks is a wrapper around atk_hypertext_get_n_links().
func (hypertext *HypertextIface) GetNLinks() int {
	ret0 := C.atk_hypertext_get_n_links(hypertext.native())
	return int(ret0)
}

// Interface Image
type Image struct {
	ImageIface
	Ptr unsafe.Pointer
}
type ImageIface struct{}

func (v *ImageIface) native() *C.AtkImage {
	return (*C.AtkImage)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapImage(p *C.AtkImage) Image {
	return Image{Ptr: unsafe.Pointer(p)}
}
func WrapImage(p unsafe.Pointer) Image {
	return Image{Ptr: p}
}
func (v Image) IsNil() bool {
	return v.Ptr == nil
}
func IWrapImage(p unsafe.Pointer) interface{} {
	return WrapImage(p)
}
func (v Image) GetType() gobject.Type {
	return gobject.Type(C.atk_image_get_type())
}
func (v Image) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapImage(unsafe.Pointer(ptr)), nil
	}
}

// GetImageDescription is a wrapper around atk_image_get_image_description().
func (image *ImageIface) GetImageDescription() string {
	ret0 := C.atk_image_get_image_description(image.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetImageLocale is a wrapper around atk_image_get_image_locale().
func (image *ImageIface) GetImageLocale() string {
	ret0 := C.atk_image_get_image_locale(image.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetImagePosition is a wrapper around atk_image_get_image_position().
func (image *ImageIface) GetImagePosition(coord_type CoordType) (int, int) {
	var x0 C.gint
	var y0 C.gint
	C.atk_image_get_image_position(image.native(), &x0, &y0, C.AtkCoordType(coord_type))
	return int(x0), int(y0)
}

// GetImageSize is a wrapper around atk_image_get_image_size().
func (image *ImageIface) GetImageSize() (int, int) {
	var width0 C.gint
	var height0 C.gint
	C.atk_image_get_image_size(image.native(), &width0, &height0)
	return int(width0), int(height0)
}

// SetImageDescription is a wrapper around atk_image_set_image_description().
func (image *ImageIface) SetImageDescription(description string) bool {
	description0 := (*C.gchar)(C.CString(description))
	ret0 := C.atk_image_set_image_description(image.native(), description0)
	C.free(unsafe.Pointer(description0))
	return util.Int2Bool(int(ret0))
}

// Interface ImplementorIface
type ImplementorIface struct {
	ImplementorIfaceIface
	Ptr unsafe.Pointer
}
type ImplementorIfaceIface struct{}

func (v *ImplementorIfaceIface) native() *C.AtkImplementorIface {
	return (*C.AtkImplementorIface)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapImplementorIface(p *C.AtkImplementorIface) ImplementorIface {
	return ImplementorIface{Ptr: unsafe.Pointer(p)}
}
func WrapImplementorIface(p unsafe.Pointer) ImplementorIface {
	return ImplementorIface{Ptr: p}
}
func (v ImplementorIface) IsNil() bool {
	return v.Ptr == nil
}
func IWrapImplementorIface(p unsafe.Pointer) interface{} {
	return WrapImplementorIface(p)
}
func (v ImplementorIface) GetType() gobject.Type {
	return gobject.Type(C.atk_implementor_get_type())
}
func (v ImplementorIface) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapImplementorIface(unsafe.Pointer(ptr)), nil
	}
}

// Interface Selection
type Selection struct {
	SelectionIface
	Ptr unsafe.Pointer
}
type SelectionIface struct{}

func (v *SelectionIface) native() *C.AtkSelection {
	return (*C.AtkSelection)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapSelection(p *C.AtkSelection) Selection {
	return Selection{Ptr: unsafe.Pointer(p)}
}
func WrapSelection(p unsafe.Pointer) Selection {
	return Selection{Ptr: p}
}
func (v Selection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSelection(p unsafe.Pointer) interface{} {
	return WrapSelection(p)
}
func (v Selection) GetType() gobject.Type {
	return gobject.Type(C.atk_selection_get_type())
}
func (v Selection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSelection(unsafe.Pointer(ptr)), nil
	}
}

// AddSelection is a wrapper around atk_selection_add_selection().
func (selection *SelectionIface) AddSelection(i int) bool {
	ret0 := C.atk_selection_add_selection(selection.native(), C.gint(i))
	return util.Int2Bool(int(ret0))
}

// ClearSelection is a wrapper around atk_selection_clear_selection().
func (selection *SelectionIface) ClearSelection() bool {
	ret0 := C.atk_selection_clear_selection(selection.native())
	return util.Int2Bool(int(ret0))
}

// GetSelectionCount is a wrapper around atk_selection_get_selection_count().
func (selection *SelectionIface) GetSelectionCount() int {
	ret0 := C.atk_selection_get_selection_count(selection.native())
	return int(ret0)
}

// IsChildSelected is a wrapper around atk_selection_is_child_selected().
func (selection *SelectionIface) IsChildSelected(i int) bool {
	ret0 := C.atk_selection_is_child_selected(selection.native(), C.gint(i))
	return util.Int2Bool(int(ret0))
}

// RefSelection is a wrapper around atk_selection_ref_selection().
func (selection *SelectionIface) RefSelection(i int) Object {
	ret0 := C.atk_selection_ref_selection(selection.native(), C.gint(i))
	return wrapObject(ret0)
}

// RemoveSelection is a wrapper around atk_selection_remove_selection().
func (selection *SelectionIface) RemoveSelection(i int) bool {
	ret0 := C.atk_selection_remove_selection(selection.native(), C.gint(i))
	return util.Int2Bool(int(ret0))
}

// SelectAllSelection is a wrapper around atk_selection_select_all_selection().
func (selection *SelectionIface) SelectAllSelection() bool {
	ret0 := C.atk_selection_select_all_selection(selection.native())
	return util.Int2Bool(int(ret0))
}

// Interface StreamableContent
type StreamableContent struct {
	StreamableContentIface
	Ptr unsafe.Pointer
}
type StreamableContentIface struct{}

func (v *StreamableContentIface) native() *C.AtkStreamableContent {
	return (*C.AtkStreamableContent)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapStreamableContent(p *C.AtkStreamableContent) StreamableContent {
	return StreamableContent{Ptr: unsafe.Pointer(p)}
}
func WrapStreamableContent(p unsafe.Pointer) StreamableContent {
	return StreamableContent{Ptr: p}
}
func (v StreamableContent) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStreamableContent(p unsafe.Pointer) interface{} {
	return WrapStreamableContent(p)
}
func (v StreamableContent) GetType() gobject.Type {
	return gobject.Type(C.atk_streamable_content_get_type())
}
func (v StreamableContent) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStreamableContent(unsafe.Pointer(ptr)), nil
	}
}

// GetMimeType is a wrapper around atk_streamable_content_get_mime_type().
func (streamable *StreamableContentIface) GetMimeType(i int) string {
	ret0 := C.atk_streamable_content_get_mime_type(streamable.native(), C.gint(i))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNMimeTypes is a wrapper around atk_streamable_content_get_n_mime_types().
func (streamable *StreamableContentIface) GetNMimeTypes() int {
	ret0 := C.atk_streamable_content_get_n_mime_types(streamable.native())
	return int(ret0)
}

// GetUri is a wrapper around atk_streamable_content_get_uri().
func (streamable *StreamableContentIface) GetUri(mime_type string) string {
	mime_type0 := (*C.gchar)(C.CString(mime_type))
	ret0 := C.atk_streamable_content_get_uri(streamable.native(), mime_type0)
	C.free(unsafe.Pointer(mime_type0))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Interface Table
type Table struct {
	TableIface
	Ptr unsafe.Pointer
}
type TableIface struct{}

func (v *TableIface) native() *C.AtkTable {
	return (*C.AtkTable)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapTable(p *C.AtkTable) Table {
	return Table{Ptr: unsafe.Pointer(p)}
}
func WrapTable(p unsafe.Pointer) Table {
	return Table{Ptr: p}
}
func (v Table) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTable(p unsafe.Pointer) interface{} {
	return WrapTable(p)
}
func (v Table) GetType() gobject.Type {
	return gobject.Type(C.atk_table_get_type())
}
func (v Table) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTable(unsafe.Pointer(ptr)), nil
	}
}

// AddColumnSelection is a wrapper around atk_table_add_column_selection().
func (table *TableIface) AddColumnSelection(column int) bool {
	ret0 := C.atk_table_add_column_selection(table.native(), C.gint(column))
	return util.Int2Bool(int(ret0))
}

// AddRowSelection is a wrapper around atk_table_add_row_selection().
func (table *TableIface) AddRowSelection(row int) bool {
	ret0 := C.atk_table_add_row_selection(table.native(), C.gint(row))
	return util.Int2Bool(int(ret0))
}

// GetCaption is a wrapper around atk_table_get_caption().
func (table *TableIface) GetCaption() Object {
	ret0 := C.atk_table_get_caption(table.native())
	return wrapObject(ret0)
}

// GetColumnDescription is a wrapper around atk_table_get_column_description().
func (table *TableIface) GetColumnDescription(column int) string {
	ret0 := C.atk_table_get_column_description(table.native(), C.gint(column))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetColumnExtentAt is a wrapper around atk_table_get_column_extent_at().
func (table *TableIface) GetColumnExtentAt(row int, column int) int {
	ret0 := C.atk_table_get_column_extent_at(table.native(), C.gint(row), C.gint(column))
	return int(ret0)
}

// GetColumnHeader is a wrapper around atk_table_get_column_header().
func (table *TableIface) GetColumnHeader(column int) Object {
	ret0 := C.atk_table_get_column_header(table.native(), C.gint(column))
	return wrapObject(ret0)
}

// GetNColumns is a wrapper around atk_table_get_n_columns().
func (table *TableIface) GetNColumns() int {
	ret0 := C.atk_table_get_n_columns(table.native())
	return int(ret0)
}

// GetNRows is a wrapper around atk_table_get_n_rows().
func (table *TableIface) GetNRows() int {
	ret0 := C.atk_table_get_n_rows(table.native())
	return int(ret0)
}

// GetRowDescription is a wrapper around atk_table_get_row_description().
func (table *TableIface) GetRowDescription(row int) string {
	ret0 := C.atk_table_get_row_description(table.native(), C.gint(row))
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetRowExtentAt is a wrapper around atk_table_get_row_extent_at().
func (table *TableIface) GetRowExtentAt(row int, column int) int {
	ret0 := C.atk_table_get_row_extent_at(table.native(), C.gint(row), C.gint(column))
	return int(ret0)
}

// GetRowHeader is a wrapper around atk_table_get_row_header().
func (table *TableIface) GetRowHeader(row int) Object {
	ret0 := C.atk_table_get_row_header(table.native(), C.gint(row))
	return wrapObject(ret0)
}

// GetSummary is a wrapper around atk_table_get_summary().
func (table *TableIface) GetSummary() Object {
	ret0 := C.atk_table_get_summary(table.native())
	return wrapObject(ret0)
}

// IsColumnSelected is a wrapper around atk_table_is_column_selected().
func (table *TableIface) IsColumnSelected(column int) bool {
	ret0 := C.atk_table_is_column_selected(table.native(), C.gint(column))
	return util.Int2Bool(int(ret0))
}

// IsRowSelected is a wrapper around atk_table_is_row_selected().
func (table *TableIface) IsRowSelected(row int) bool {
	ret0 := C.atk_table_is_row_selected(table.native(), C.gint(row))
	return util.Int2Bool(int(ret0))
}

// IsSelected is a wrapper around atk_table_is_selected().
func (table *TableIface) IsSelected(row int, column int) bool {
	ret0 := C.atk_table_is_selected(table.native(), C.gint(row), C.gint(column))
	return util.Int2Bool(int(ret0))
}

// RefAt is a wrapper around atk_table_ref_at().
func (table *TableIface) RefAt(row int, column int) Object {
	ret0 := C.atk_table_ref_at(table.native(), C.gint(row), C.gint(column))
	return wrapObject(ret0)
}

// RemoveColumnSelection is a wrapper around atk_table_remove_column_selection().
func (table *TableIface) RemoveColumnSelection(column int) bool {
	ret0 := C.atk_table_remove_column_selection(table.native(), C.gint(column))
	return util.Int2Bool(int(ret0))
}

// RemoveRowSelection is a wrapper around atk_table_remove_row_selection().
func (table *TableIface) RemoveRowSelection(row int) bool {
	ret0 := C.atk_table_remove_row_selection(table.native(), C.gint(row))
	return util.Int2Bool(int(ret0))
}

// SetCaption is a wrapper around atk_table_set_caption().
func (table *TableIface) SetCaption(caption Object) {
	C.atk_table_set_caption(table.native(), caption.native())
}

// SetColumnDescription is a wrapper around atk_table_set_column_description().
func (table *TableIface) SetColumnDescription(column int, description string) {
	description0 := (*C.gchar)(C.CString(description))
	C.atk_table_set_column_description(table.native(), C.gint(column), description0)
	C.free(unsafe.Pointer(description0))
}

// SetColumnHeader is a wrapper around atk_table_set_column_header().
func (table *TableIface) SetColumnHeader(column int, header Object) {
	C.atk_table_set_column_header(table.native(), C.gint(column), header.native())
}

// SetRowDescription is a wrapper around atk_table_set_row_description().
func (table *TableIface) SetRowDescription(row int, description string) {
	description0 := (*C.gchar)(C.CString(description))
	C.atk_table_set_row_description(table.native(), C.gint(row), description0)
	C.free(unsafe.Pointer(description0))
}

// SetRowHeader is a wrapper around atk_table_set_row_header().
func (table *TableIface) SetRowHeader(row int, header Object) {
	C.atk_table_set_row_header(table.native(), C.gint(row), header.native())
}

// SetSummary is a wrapper around atk_table_set_summary().
func (table *TableIface) SetSummary(accessible Object) {
	C.atk_table_set_summary(table.native(), accessible.native())
}

// Interface TableCell
type TableCell struct {
	TableCellIface
	Ptr unsafe.Pointer
}
type TableCellIface struct{}

func (v *TableCellIface) native() *C.AtkTableCell {
	return (*C.AtkTableCell)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapTableCell(p *C.AtkTableCell) TableCell {
	return TableCell{Ptr: unsafe.Pointer(p)}
}
func WrapTableCell(p unsafe.Pointer) TableCell {
	return TableCell{Ptr: p}
}
func (v TableCell) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTableCell(p unsafe.Pointer) interface{} {
	return WrapTableCell(p)
}
func (v TableCell) GetType() gobject.Type {
	return gobject.Type(C.atk_table_cell_get_type())
}
func (v TableCell) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTableCell(unsafe.Pointer(ptr)), nil
	}
}

// GetColumnSpan is a wrapper around atk_table_cell_get_column_span().
func (cell *TableCellIface) GetColumnSpan() int {
	ret0 := C.atk_table_cell_get_column_span(cell.native())
	return int(ret0)
}

// GetPosition is a wrapper around atk_table_cell_get_position().
func (cell *TableCellIface) GetPosition() (bool, int, int) {
	var row0 C.gint
	var column0 C.gint
	ret0 := C.atk_table_cell_get_position(cell.native(), &row0, &column0)
	return util.Int2Bool(int(ret0)), int(row0), int(column0)
}

// GetRowColumnSpan is a wrapper around atk_table_cell_get_row_column_span().
func (cell *TableCellIface) GetRowColumnSpan() (bool, int, int, int, int) {
	var row0 C.gint
	var column0 C.gint
	var row_span0 C.gint
	var column_span0 C.gint
	ret0 := C.atk_table_cell_get_row_column_span(cell.native(), &row0, &column0, &row_span0, &column_span0)
	return util.Int2Bool(int(ret0)), int(row0), int(column0), int(row_span0), int(column_span0)
}

// GetRowSpan is a wrapper around atk_table_cell_get_row_span().
func (cell *TableCellIface) GetRowSpan() int {
	ret0 := C.atk_table_cell_get_row_span(cell.native())
	return int(ret0)
}

// GetTable is a wrapper around atk_table_cell_get_table().
func (cell *TableCellIface) GetTable() Object {
	ret0 := C.atk_table_cell_get_table(cell.native())
	return wrapObject(ret0)
}

// Interface Text
type Text struct {
	TextIface
	Ptr unsafe.Pointer
}
type TextIface struct{}

func (v *TextIface) native() *C.AtkText {
	return (*C.AtkText)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapText(p *C.AtkText) Text {
	return Text{Ptr: unsafe.Pointer(p)}
}
func WrapText(p unsafe.Pointer) Text {
	return Text{Ptr: p}
}
func (v Text) IsNil() bool {
	return v.Ptr == nil
}
func IWrapText(p unsafe.Pointer) interface{} {
	return WrapText(p)
}
func (v Text) GetType() gobject.Type {
	return gobject.Type(C.atk_text_get_type())
}
func (v Text) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapText(unsafe.Pointer(ptr)), nil
	}
}

// AddSelection is a wrapper around atk_text_add_selection().
func (text *TextIface) AddSelection(start_offset int, end_offset int) bool {
	ret0 := C.atk_text_add_selection(text.native(), C.gint(start_offset), C.gint(end_offset))
	return util.Int2Bool(int(ret0))
}

// GetBoundedRanges is a wrapper around atk_text_get_bounded_ranges().
func (text *TextIface) GetBoundedRanges(rect TextRectangle, coord_type CoordType, x_clip_type TextClipType, y_clip_type TextClipType) []TextRange {
	ret0 := C.atk_text_get_bounded_ranges(text.native(), rect.native(), C.AtkCoordType(coord_type), C.AtkTextClipType(x_clip_type), C.AtkTextClipType(y_clip_type))
	var ret0Slice []*C.AtkTextRange
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))))
	ret := make([]TextRange, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = wrapTextRange(elem)
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetCaretOffset is a wrapper around atk_text_get_caret_offset().
func (text *TextIface) GetCaretOffset() int {
	ret0 := C.atk_text_get_caret_offset(text.native())
	return int(ret0)
}

// GetCharacterCount is a wrapper around atk_text_get_character_count().
func (text *TextIface) GetCharacterCount() int {
	ret0 := C.atk_text_get_character_count(text.native())
	return int(ret0)
}

// GetCharacterExtents is a wrapper around atk_text_get_character_extents().
func (text *TextIface) GetCharacterExtents(offset int, coords CoordType) (int, int, int, int) {
	var x0 C.gint
	var y0 C.gint
	var width0 C.gint
	var height0 C.gint
	C.atk_text_get_character_extents(text.native(), C.gint(offset), &x0, &y0, &width0, &height0, C.AtkCoordType(coords))
	return int(x0), int(y0), int(width0), int(height0)
}

// GetNSelections is a wrapper around atk_text_get_n_selections().
func (text *TextIface) GetNSelections() int {
	ret0 := C.atk_text_get_n_selections(text.native())
	return int(ret0)
}

// GetOffsetAtPoint is a wrapper around atk_text_get_offset_at_point().
func (text *TextIface) GetOffsetAtPoint(x int, y int, coords CoordType) int {
	ret0 := C.atk_text_get_offset_at_point(text.native(), C.gint(x), C.gint(y), C.AtkCoordType(coords))
	return int(ret0)
}

// GetRangeExtents is a wrapper around atk_text_get_range_extents().
func (text *TextIface) GetRangeExtents(start_offset int, end_offset int, coord_type CoordType) TextRectangle {
	var rect0 C.AtkTextRectangle
	C.atk_text_get_range_extents(text.native(), C.gint(start_offset), C.gint(end_offset), C.AtkCoordType(coord_type), &rect0)
	return wrapTextRectangle(&rect0)
}

// GetSelection is a wrapper around atk_text_get_selection().
func (text *TextIface) GetSelection(selection_num int) (string, int, int) {
	var start_offset0 C.gint
	var end_offset0 C.gint
	ret0 := C.atk_text_get_selection(text.native(), C.gint(selection_num), &start_offset0, &end_offset0)
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret, int(start_offset0), int(end_offset0)
}

// GetStringAtOffset is a wrapper around atk_text_get_string_at_offset().
func (text *TextIface) GetStringAtOffset(offset int, granularity TextGranularity) (string, int, int) {
	var start_offset0 C.gint
	var end_offset0 C.gint
	ret0 := C.atk_text_get_string_at_offset(text.native(), C.gint(offset), C.AtkTextGranularity(granularity), &start_offset0, &end_offset0)
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret, int(start_offset0), int(end_offset0)
}

// GetText is a wrapper around atk_text_get_text().
func (text *TextIface) GetText(start_offset int, end_offset int) string {
	ret0 := C.atk_text_get_text(text.native(), C.gint(start_offset), C.gint(end_offset))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// RemoveSelection is a wrapper around atk_text_remove_selection().
func (text *TextIface) RemoveSelection(selection_num int) bool {
	ret0 := C.atk_text_remove_selection(text.native(), C.gint(selection_num))
	return util.Int2Bool(int(ret0))
}

// SetCaretOffset is a wrapper around atk_text_set_caret_offset().
func (text *TextIface) SetCaretOffset(offset int) bool {
	ret0 := C.atk_text_set_caret_offset(text.native(), C.gint(offset))
	return util.Int2Bool(int(ret0))
}

// SetSelection is a wrapper around atk_text_set_selection().
func (text *TextIface) SetSelection(selection_num int, start_offset int, end_offset int) bool {
	ret0 := C.atk_text_set_selection(text.native(), C.gint(selection_num), C.gint(start_offset), C.gint(end_offset))
	return util.Int2Bool(int(ret0))
}

// TextFreeRanges is a wrapper around atk_text_free_ranges().
func TextFreeRanges(ranges []TextRange) {
	ranges0 := make([]*C.AtkTextRange, len(ranges))
	for idx, elemG := range ranges {
		ranges0[idx] = elemG.native()
	}
	var ranges0Ptr **C.AtkTextRange
	if len(ranges0) > 0 {
		ranges0Ptr = &ranges0[0]
	}
	C.atk_text_free_ranges(ranges0Ptr)
}

// Interface Value
type Value struct {
	ValueIface
	Ptr unsafe.Pointer
}
type ValueIface struct{}

func (v *ValueIface) native() *C.AtkValue {
	return (*C.AtkValue)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapValue(p *C.AtkValue) Value {
	return Value{Ptr: unsafe.Pointer(p)}
}
func WrapValue(p unsafe.Pointer) Value {
	return Value{Ptr: p}
}
func (v Value) IsNil() bool {
	return v.Ptr == nil
}
func IWrapValue(p unsafe.Pointer) interface{} {
	return WrapValue(p)
}
func (v Value) GetType() gobject.Type {
	return gobject.Type(C.atk_value_get_type())
}
func (v Value) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapValue(unsafe.Pointer(ptr)), nil
	}
}

// GetIncrement is a wrapper around atk_value_get_increment().
func (obj *ValueIface) GetIncrement() float64 {
	ret0 := C.atk_value_get_increment(obj.native())
	return float64(ret0)
}

// GetRange is a wrapper around atk_value_get_range().
func (obj *ValueIface) GetRange() Range {
	ret0 := C.atk_value_get_range(obj.native())
	return wrapRange(ret0)
}

// GetSubRanges is a wrapper around atk_value_get_sub_ranges().
func (obj *ValueIface) GetSubRanges() glib.SList {
	ret0 := C.atk_value_get_sub_ranges(obj.native())
	return glib.WrapSList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapRange(p) })
}

// GetValueAndText is a wrapper around atk_value_get_value_and_text().
func (obj *ValueIface) GetValueAndText() (float64, string) {
	var value0 C.gdouble
	var text0 *C.gchar
	C.atk_value_get_value_and_text(obj.native(), &value0, &text0)
	text := C.GoString((*C.char)(text0))
	defer C.g_free(C.gpointer(text0))
	return float64(value0), text
}

// SetValue is a wrapper around atk_value_set_value().
func (obj *ValueIface) SetValue(new_value float64) {
	C.atk_value_set_value(obj.native(), C.gdouble(new_value))
}

// Interface Window
type Window struct {
	WindowIface
	Ptr unsafe.Pointer
}
type WindowIface struct{}

func (v *WindowIface) native() *C.AtkWindow {
	return (*C.AtkWindow)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapWindow(p *C.AtkWindow) Window {
	return Window{Ptr: unsafe.Pointer(p)}
}
func WrapWindow(p unsafe.Pointer) Window {
	return Window{Ptr: p}
}
func (v Window) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWindow(p unsafe.Pointer) interface{} {
	return WrapWindow(p)
}
func (v Window) GetType() gobject.Type {
	return gobject.Type(C.atk_window_get_type())
}
func (v Window) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapWindow(unsafe.Pointer(ptr)), nil
	}
}

// Object GObjectAccessible
type GObjectAccessible struct {
	Object
}

func (v GObjectAccessible) native() *C.AtkGObjectAccessible {
	return (*C.AtkGObjectAccessible)(v.Ptr)
}
func wrapGObjectAccessible(p *C.AtkGObjectAccessible) (v GObjectAccessible) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGObjectAccessible(p unsafe.Pointer) (v GObjectAccessible) {
	v.Ptr = p
	return
}
func (v GObjectAccessible) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGObjectAccessible(p unsafe.Pointer) interface{} {
	return WrapGObjectAccessible(p)
}
func (v GObjectAccessible) GetType() gobject.Type {
	return gobject.Type(C.atk_gobject_accessible_get_type())
}
func (v GObjectAccessible) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGObjectAccessible(unsafe.Pointer(ptr)), nil
	}
}

// GetObject is a wrapper around atk_gobject_accessible_get_object().
func (obj GObjectAccessible) GetObject() gobject.Object {
	ret0 := C.atk_gobject_accessible_get_object(obj.native())
	return gobject.WrapObject(unsafe.Pointer(ret0))
}

// GObjectAccessibleForObject is a wrapper around atk_gobject_accessible_for_object().
func GObjectAccessibleForObject(obj gobject.Object) Object {
	ret0 := C.atk_gobject_accessible_for_object((*C.GObject)(obj.Ptr))
	return wrapObject(ret0)
}

// Object NoOpObject
type NoOpObject struct {
	ActionIface
	ComponentIface
	DocumentIface
	EditableTextIface
	HypertextIface
	ImageIface
	SelectionIface
	TableIface
	TableCellIface
	TextIface
	ValueIface
	WindowIface
	Object
}

func (v NoOpObject) native() *C.AtkNoOpObject {
	return (*C.AtkNoOpObject)(v.Ptr)
}
func wrapNoOpObject(p *C.AtkNoOpObject) (v NoOpObject) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNoOpObject(p unsafe.Pointer) (v NoOpObject) {
	v.Ptr = p
	return
}
func (v NoOpObject) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNoOpObject(p unsafe.Pointer) interface{} {
	return WrapNoOpObject(p)
}
func (v NoOpObject) GetType() gobject.Type {
	return gobject.Type(C.atk_no_op_object_get_type())
}
func (v NoOpObject) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNoOpObject(unsafe.Pointer(ptr)), nil
	}
}
func (v NoOpObject) Action() Action {
	return WrapAction(v.Ptr)
}
func (v NoOpObject) Component() Component {
	return WrapComponent(v.Ptr)
}
func (v NoOpObject) Document() Document {
	return WrapDocument(v.Ptr)
}
func (v NoOpObject) EditableText() EditableText {
	return WrapEditableText(v.Ptr)
}
func (v NoOpObject) Hypertext() Hypertext {
	return WrapHypertext(v.Ptr)
}
func (v NoOpObject) Image() Image {
	return WrapImage(v.Ptr)
}
func (v NoOpObject) Selection() Selection {
	return WrapSelection(v.Ptr)
}
func (v NoOpObject) Table() Table {
	return WrapTable(v.Ptr)
}
func (v NoOpObject) TableCell() TableCell {
	return WrapTableCell(v.Ptr)
}
func (v NoOpObject) Text() Text {
	return WrapText(v.Ptr)
}
func (v NoOpObject) Value() Value {
	return WrapValue(v.Ptr)
}
func (v NoOpObject) Window() Window {
	return WrapWindow(v.Ptr)
}

// NoOpObjectNew is a wrapper around atk_no_op_object_new().
func NoOpObjectNew(obj gobject.Object) Object {
	ret0 := C.atk_no_op_object_new((*C.GObject)(obj.Ptr))
	return wrapObject(ret0)
}

// Object NoOpObjectFactory
type NoOpObjectFactory struct {
	ObjectFactory
}

func (v NoOpObjectFactory) native() *C.AtkNoOpObjectFactory {
	return (*C.AtkNoOpObjectFactory)(v.Ptr)
}
func wrapNoOpObjectFactory(p *C.AtkNoOpObjectFactory) (v NoOpObjectFactory) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNoOpObjectFactory(p unsafe.Pointer) (v NoOpObjectFactory) {
	v.Ptr = p
	return
}
func (v NoOpObjectFactory) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNoOpObjectFactory(p unsafe.Pointer) interface{} {
	return WrapNoOpObjectFactory(p)
}
func (v NoOpObjectFactory) GetType() gobject.Type {
	return gobject.Type(C.atk_no_op_object_factory_get_type())
}
func (v NoOpObjectFactory) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNoOpObjectFactory(unsafe.Pointer(ptr)), nil
	}
}

// NoOpObjectFactoryNew is a wrapper around atk_no_op_object_factory_new().
func NoOpObjectFactoryNew() ObjectFactory {
	ret0 := C.atk_no_op_object_factory_new()
	return wrapObjectFactory(ret0)
}

// Object ObjectFactory
type ObjectFactory struct {
	gobject.Object
}

func (v ObjectFactory) native() *C.AtkObjectFactory {
	return (*C.AtkObjectFactory)(v.Ptr)
}
func wrapObjectFactory(p *C.AtkObjectFactory) (v ObjectFactory) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapObjectFactory(p unsafe.Pointer) (v ObjectFactory) {
	v.Ptr = p
	return
}
func (v ObjectFactory) IsNil() bool {
	return v.Ptr == nil
}
func IWrapObjectFactory(p unsafe.Pointer) interface{} {
	return WrapObjectFactory(p)
}
func (v ObjectFactory) GetType() gobject.Type {
	return gobject.Type(C.atk_object_factory_get_type())
}
func (v ObjectFactory) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapObjectFactory(unsafe.Pointer(ptr)), nil
	}
}

// CreateAccessible is a wrapper around atk_object_factory_create_accessible().
func (factory ObjectFactory) CreateAccessible(obj gobject.Object) Object {
	ret0 := C.atk_object_factory_create_accessible(factory.native(), (*C.GObject)(obj.Ptr))
	return wrapObject(ret0)
}

// GetAccessibleType is a wrapper around atk_object_factory_get_accessible_type().
func (factory ObjectFactory) GetAccessibleType() gobject.Type {
	ret0 := C.atk_object_factory_get_accessible_type(factory.native())
	return gobject.Type(ret0)
}

// Invalidate is a wrapper around atk_object_factory_invalidate().
func (factory ObjectFactory) Invalidate() {
	C.atk_object_factory_invalidate(factory.native())
}

// Object StateSet
type StateSet struct {
	gobject.Object
}

func (v StateSet) native() *C.AtkStateSet {
	return (*C.AtkStateSet)(v.Ptr)
}
func wrapStateSet(p *C.AtkStateSet) (v StateSet) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapStateSet(p unsafe.Pointer) (v StateSet) {
	v.Ptr = p
	return
}
func (v StateSet) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStateSet(p unsafe.Pointer) interface{} {
	return WrapStateSet(p)
}
func (v StateSet) GetType() gobject.Type {
	return gobject.Type(C.atk_state_set_get_type())
}
func (v StateSet) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapStateSet(unsafe.Pointer(ptr)), nil
	}
}

// StateSetNew is a wrapper around atk_state_set_new().
func StateSetNew() StateSet {
	ret0 := C.atk_state_set_new()
	return wrapStateSet(ret0)
}

// AddState is a wrapper around atk_state_set_add_state().
func (set StateSet) AddState(type_ StateType) bool {
	ret0 := C.atk_state_set_add_state(set.native(), C.AtkStateType(type_))
	return util.Int2Bool(int(ret0))
}

// AddStates is a wrapper around atk_state_set_add_states().
func (set StateSet) AddStates(types []StateType) {
	types0 := make([]C.AtkStateType, len(types))
	for idx, elemG := range types {
		types0[idx] = C.AtkStateType(elemG)
	}
	var types0Ptr *C.AtkStateType
	if len(types0) > 0 {
		types0Ptr = &types0[0]
	}
	C.atk_state_set_add_states(set.native(), types0Ptr, C.gint(len(types)))
}

// AndSets is a wrapper around atk_state_set_and_sets().
func (set StateSet) AndSets(compare_set StateSet) StateSet {
	ret0 := C.atk_state_set_and_sets(set.native(), compare_set.native())
	return wrapStateSet(ret0)
}

// ClearStates is a wrapper around atk_state_set_clear_states().
func (set StateSet) ClearStates() {
	C.atk_state_set_clear_states(set.native())
}

// ContainsState is a wrapper around atk_state_set_contains_state().
func (set StateSet) ContainsState(type_ StateType) bool {
	ret0 := C.atk_state_set_contains_state(set.native(), C.AtkStateType(type_))
	return util.Int2Bool(int(ret0))
}

// ContainsStates is a wrapper around atk_state_set_contains_states().
func (set StateSet) ContainsStates(types []StateType) bool {
	types0 := make([]C.AtkStateType, len(types))
	for idx, elemG := range types {
		types0[idx] = C.AtkStateType(elemG)
	}
	var types0Ptr *C.AtkStateType
	if len(types0) > 0 {
		types0Ptr = &types0[0]
	}
	ret0 := C.atk_state_set_contains_states(set.native(), types0Ptr, C.gint(len(types)))
	return util.Int2Bool(int(ret0))
}

// IsEmpty is a wrapper around atk_state_set_is_empty().
func (set StateSet) IsEmpty() bool {
	ret0 := C.atk_state_set_is_empty(set.native())
	return util.Int2Bool(int(ret0))
}

// OrSets is a wrapper around atk_state_set_or_sets().
func (set StateSet) OrSets(compare_set StateSet) StateSet {
	ret0 := C.atk_state_set_or_sets(set.native(), compare_set.native())
	return wrapStateSet(ret0)
}

// RemoveState is a wrapper around atk_state_set_remove_state().
func (set StateSet) RemoveState(type_ StateType) bool {
	ret0 := C.atk_state_set_remove_state(set.native(), C.AtkStateType(type_))
	return util.Int2Bool(int(ret0))
}

// XorSets is a wrapper around atk_state_set_xor_sets().
func (set StateSet) XorSets(compare_set StateSet) StateSet {
	ret0 := C.atk_state_set_xor_sets(set.native(), compare_set.native())
	return wrapStateSet(ret0)
}

// Object Plug
type Plug struct {
	ComponentIface
	Object
}

func (v Plug) native() *C.AtkPlug {
	return (*C.AtkPlug)(v.Ptr)
}
func wrapPlug(p *C.AtkPlug) (v Plug) {
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
	return gobject.Type(C.atk_plug_get_type())
}
func (v Plug) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPlug(unsafe.Pointer(ptr)), nil
	}
}
func (v Plug) Component() Component {
	return WrapComponent(v.Ptr)
}

// PlugNew is a wrapper around atk_plug_new().
func PlugNew() Object {
	ret0 := C.atk_plug_new()
	return wrapObject(ret0)
}

// GetId is a wrapper around atk_plug_get_id().
func (plug Plug) GetId() string {
	ret0 := C.atk_plug_get_id(plug.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object Registry
type Registry struct {
	gobject.Object
}

func (v Registry) native() *C.AtkRegistry {
	return (*C.AtkRegistry)(v.Ptr)
}
func wrapRegistry(p *C.AtkRegistry) (v Registry) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRegistry(p unsafe.Pointer) (v Registry) {
	v.Ptr = p
	return
}
func (v Registry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRegistry(p unsafe.Pointer) interface{} {
	return WrapRegistry(p)
}
func (v Registry) GetType() gobject.Type {
	return gobject.Type(C.atk_registry_get_type())
}
func (v Registry) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRegistry(unsafe.Pointer(ptr)), nil
	}
}

// GetFactory is a wrapper around atk_registry_get_factory().
func (registry Registry) GetFactory(type_ gobject.Type) ObjectFactory {
	ret0 := C.atk_registry_get_factory(registry.native(), C.GType(type_))
	return wrapObjectFactory(ret0)
}

// GetFactoryType is a wrapper around atk_registry_get_factory_type().
func (registry Registry) GetFactoryType(type_ gobject.Type) gobject.Type {
	ret0 := C.atk_registry_get_factory_type(registry.native(), C.GType(type_))
	return gobject.Type(ret0)
}

// SetFactoryType is a wrapper around atk_registry_set_factory_type().
func (registry Registry) SetFactoryType(type_ gobject.Type, factory_type gobject.Type) {
	C.atk_registry_set_factory_type(registry.native(), C.GType(type_), C.GType(factory_type))
}

// Object Socket
type Socket struct {
	ComponentIface
	Object
}

func (v Socket) native() *C.AtkSocket {
	return (*C.AtkSocket)(v.Ptr)
}
func wrapSocket(p *C.AtkSocket) (v Socket) {
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
	return gobject.Type(C.atk_socket_get_type())
}
func (v Socket) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocket(unsafe.Pointer(ptr)), nil
	}
}
func (v Socket) Component() Component {
	return WrapComponent(v.Ptr)
}

// SocketNew is a wrapper around atk_socket_new().
func SocketNew() Object {
	ret0 := C.atk_socket_new()
	return wrapObject(ret0)
}

// Embed is a wrapper around atk_socket_embed().
func (obj Socket) Embed(plug_id string) {
	plug_id0 := (*C.gchar)(C.CString(plug_id))
	C.atk_socket_embed(obj.native(), plug_id0)
	C.free(unsafe.Pointer(plug_id0))
}

// IsOccupied is a wrapper around atk_socket_is_occupied().
func (obj Socket) IsOccupied() bool {
	ret0 := C.atk_socket_is_occupied(obj.native())
	return util.Int2Bool(int(ret0))
}

// Object Util
type Util struct {
	gobject.Object
}

func (v Util) native() *C.AtkUtil {
	return (*C.AtkUtil)(v.Ptr)
}
func wrapUtil(p *C.AtkUtil) (v Util) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUtil(p unsafe.Pointer) (v Util) {
	v.Ptr = p
	return
}
func (v Util) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUtil(p unsafe.Pointer) interface{} {
	return WrapUtil(p)
}
func (v Util) GetType() gobject.Type {
	return gobject.Type(C.atk_util_get_type())
}
func (v Util) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUtil(unsafe.Pointer(ptr)), nil
	}
}

type State uint64
type CoordType int

const (
	CoordTypeScreen CoordType = 0
	CoordTypeWindow           = 1
)

type KeyEventType int

const (
	KeyEventTypePress       KeyEventType = 0
	KeyEventTypeRelease                  = 1
	KeyEventTypeLastDefined              = 2
)

type Layer int

const (
	LayerInvalid    Layer = 0
	LayerBackground       = 1
	LayerCanvas           = 2
	LayerWidget           = 3
	LayerMdi              = 4
	LayerPopup            = 5
	LayerOverlay          = 6
	LayerWindow           = 7
)

type RelationType int

const (
	RelationTypeNull           RelationType = 0
	RelationTypeControlledBy                = 1
	RelationTypeControllerFor               = 2
	RelationTypeLabelFor                    = 3
	RelationTypeLabelledBy                  = 4
	RelationTypeMemberOf                    = 5
	RelationTypeNodeChildOf                 = 6
	RelationTypeFlowsTo                     = 7
	RelationTypeFlowsFrom                   = 8
	RelationTypeSubwindowOf                 = 9
	RelationTypeEmbeds                      = 10
	RelationTypeEmbeddedBy                  = 11
	RelationTypePopupFor                    = 12
	RelationTypeParentWindowOf              = 13
	RelationTypeDescribedBy                 = 14
	RelationTypeDescriptionFor              = 15
	RelationTypeNodeParentOf                = 16
	RelationTypeLastDefined                 = 17
)

type Role int

const (
	RoleInvalid              Role = 0
	RoleAcceleratorLabel          = 1
	RoleAlert                     = 2
	RoleAnimation                 = 3
	RoleArrow                     = 4
	RoleCalendar                  = 5
	RoleCanvas                    = 6
	RoleCheckBox                  = 7
	RoleCheckMenuItem             = 8
	RoleColorChooser              = 9
	RoleColumnHeader              = 10
	RoleComboBox                  = 11
	RoleDateEditor                = 12
	RoleDesktopIcon               = 13
	RoleDesktopFrame              = 14
	RoleDial                      = 15
	RoleDialog                    = 16
	RoleDirectoryPane             = 17
	RoleDrawingArea               = 18
	RoleFileChooser               = 19
	RoleFiller                    = 20
	RoleFontChooser               = 21
	RoleFrame                     = 22
	RoleGlassPane                 = 23
	RoleHtmlContainer             = 24
	RoleIcon                      = 25
	RoleImage                     = 26
	RoleInternalFrame             = 27
	RoleLabel                     = 28
	RoleLayeredPane               = 29
	RoleList                      = 30
	RoleListItem                  = 31
	RoleMenu                      = 32
	RoleMenuBar                   = 33
	RoleMenuItem                  = 34
	RoleOptionPane                = 35
	RolePageTab                   = 36
	RolePageTabList               = 37
	RolePanel                     = 38
	RolePasswordText              = 39
	RolePopupMenu                 = 40
	RoleProgressBar               = 41
	RolePushButton                = 42
	RoleRadioButton               = 43
	RoleRadioMenuItem             = 44
	RoleRootPane                  = 45
	RoleRowHeader                 = 46
	RoleScrollBar                 = 47
	RoleScrollPane                = 48
	RoleSeparator                 = 49
	RoleSlider                    = 50
	RoleSplitPane                 = 51
	RoleSpinButton                = 52
	RoleStatusbar                 = 53
	RoleTable                     = 54
	RoleTableCell                 = 55
	RoleTableColumnHeader         = 56
	RoleTableRowHeader            = 57
	RoleTearOffMenuItem           = 58
	RoleTerminal                  = 59
	RoleText                      = 60
	RoleToggleButton              = 61
	RoleToolBar                   = 62
	RoleToolTip                   = 63
	RoleTree                      = 64
	RoleTreeTable                 = 65
	RoleUnknown                   = 66
	RoleViewport                  = 67
	RoleWindow                    = 68
	RoleHeader                    = 69
	RoleFooter                    = 70
	RoleParagraph                 = 71
	RoleRuler                     = 72
	RoleApplication               = 73
	RoleAutocomplete              = 74
	RoleEditBar                   = 75
	RoleEmbedded                  = 76
	RoleEntry                     = 77
	RoleChart                     = 78
	RoleCaption                   = 79
	RoleDocumentFrame             = 80
	RoleHeading                   = 81
	RolePage                      = 82
	RoleSection                   = 83
	RoleRedundantObject           = 84
	RoleForm                      = 85
	RoleLink                      = 86
	RoleInputMethodWindow         = 87
	RoleTableRow                  = 88
	RoleTreeItem                  = 89
	RoleDocumentSpreadsheet       = 90
	RoleDocumentPresentation      = 91
	RoleDocumentText              = 92
	RoleDocumentWeb               = 93
	RoleDocumentEmail             = 94
	RoleComment                   = 95
	RoleListBox                   = 96
	RoleGrouping                  = 97
	RoleImageMap                  = 98
	RoleNotification              = 99
	RoleInfoBar                   = 100
	RoleLevelBar                  = 101
	RoleTitleBar                  = 102
	RoleBlockQuote                = 103
	RoleAudio                     = 104
	RoleVideo                     = 105
	RoleDefinition                = 106
	RoleArticle                   = 107
	RoleLandmark                  = 108
	RoleLog                       = 109
	RoleMarquee                   = 110
	RoleMath                      = 111
	RoleRating                    = 112
	RoleTimer                     = 113
	RoleDescriptionList           = 114
	RoleDescriptionTerm           = 115
	RoleDescriptionValue          = 116
	RoleStatic                    = 117
	RoleMathFraction              = 118
	RoleMathRoot                  = 119
	RoleSubscript                 = 120
	RoleSuperscript               = 121
	RoleLastDefined               = 122
)

type StateType int

const (
	StateTypeInvalid                StateType = 0
	StateTypeActive                           = 1
	StateTypeArmed                            = 2
	StateTypeBusy                             = 3
	StateTypeChecked                          = 4
	StateTypeDefunct                          = 5
	StateTypeEditable                         = 6
	StateTypeEnabled                          = 7
	StateTypeExpandable                       = 8
	StateTypeExpanded                         = 9
	StateTypeFocusable                        = 10
	StateTypeFocused                          = 11
	StateTypeHorizontal                       = 12
	StateTypeIconified                        = 13
	StateTypeModal                            = 14
	StateTypeMultiLine                        = 15
	StateTypeMultiselectable                  = 16
	StateTypeOpaque                           = 17
	StateTypePressed                          = 18
	StateTypeResizable                        = 19
	StateTypeSelectable                       = 20
	StateTypeSelected                         = 21
	StateTypeSensitive                        = 22
	StateTypeShowing                          = 23
	StateTypeSingleLine                       = 24
	StateTypeStale                            = 25
	StateTypeTransient                        = 26
	StateTypeVertical                         = 27
	StateTypeVisible                          = 28
	StateTypeManagesDescendants               = 29
	StateTypeIndeterminate                    = 30
	StateTypeTruncated                        = 31
	StateTypeRequired                         = 32
	StateTypeInvalidEntry                     = 33
	StateTypeSupportsAutocompletion           = 34
	StateTypeSelectableText                   = 35
	StateTypeDefault                          = 36
	StateTypeAnimated                         = 37
	StateTypeVisited                          = 38
	StateTypeCheckable                        = 39
	StateTypeHasPopup                         = 40
	StateTypeHasTooltip                       = 41
	StateTypeReadOnly                         = 42
	StateTypeLastDefined                      = 43
)

type TextAttribute int

const (
	TextAttributeInvalid          TextAttribute = 0
	TextAttributeLeftMargin                     = 1
	TextAttributeRightMargin                    = 2
	TextAttributeIndent                         = 3
	TextAttributeInvisible                      = 4
	TextAttributeEditable                       = 5
	TextAttributePixelsAboveLines               = 6
	TextAttributePixelsBelowLines               = 7
	TextAttributePixelsInsideWrap               = 8
	TextAttributeBgFullHeight                   = 9
	TextAttributeRise                           = 10
	TextAttributeUnderline                      = 11
	TextAttributeStrikethrough                  = 12
	TextAttributeSize                           = 13
	TextAttributeScale                          = 14
	TextAttributeWeight                         = 15
	TextAttributeLanguage                       = 16
	TextAttributeFamilyName                     = 17
	TextAttributeBgColor                        = 18
	TextAttributeFgColor                        = 19
	TextAttributeBgStipple                      = 20
	TextAttributeFgStipple                      = 21
	TextAttributeWrapMode                       = 22
	TextAttributeDirection                      = 23
	TextAttributeJustification                  = 24
	TextAttributeStretch                        = 25
	TextAttributeVariant                        = 26
	TextAttributeStyle                          = 27
	TextAttributeLastDefined                    = 28
)

type TextBoundary int

const (
	TextBoundaryChar          TextBoundary = 0
	TextBoundaryWordStart                  = 1
	TextBoundaryWordEnd                    = 2
	TextBoundarySentenceStart              = 3
	TextBoundarySentenceEnd                = 4
	TextBoundaryLineStart                  = 5
	TextBoundaryLineEnd                    = 6
)

type TextClipType int

const (
	TextClipTypeNone TextClipType = 0
	TextClipTypeMin               = 1
	TextClipTypeMax               = 2
	TextClipTypeBoth              = 3
)

type TextGranularity int

const (
	TextGranularityChar      TextGranularity = 0
	TextGranularityWord                      = 1
	TextGranularitySentence                  = 2
	TextGranularityLine                      = 3
	TextGranularityParagraph                 = 4
)

type ValueType int

const (
	ValueTypeVeryWeak    ValueType = 0
	ValueTypeWeak                  = 1
	ValueTypeAcceptable            = 2
	ValueTypeStrong                = 3
	ValueTypeVeryStrong            = 4
	ValueTypeVeryLow               = 5
	ValueTypeLow                   = 6
	ValueTypeMedium                = 7
	ValueTypeHigh                  = 8
	ValueTypeVeryHigh              = 9
	ValueTypeVeryBad               = 10
	ValueTypeBad                   = 11
	ValueTypeGood                  = 12
	ValueTypeVeryGood              = 13
	ValueTypeBest                  = 14
	ValueTypeLastDefined           = 15
)

type HyperlinkStateFlags int

const (
	HyperlinkStateFlagsInline HyperlinkStateFlags = 1
)
