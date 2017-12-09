package gdk

/*
#cgo pkg-config: gdk-3.0
#include <gdk/gdk.h>
#include <stdlib.h>
*/
import "C"
import "github.com/linuxdeepin/go-gir/cairo-1.0"
import "github.com/linuxdeepin/go-gir/gdkpixbuf-2.0"
import "github.com/linuxdeepin/go-gir/gio-2.0"
import "github.com/linuxdeepin/go-gir/glib-2.0"
import "github.com/linuxdeepin/go-gir/gobject-2.0"
import "github.com/linuxdeepin/go-gir/pango-1.0"
import "github.com/linuxdeepin/go-gir/util"
import "unsafe"

// Struct Color
type Color struct {
	Ptr unsafe.Pointer
}

func (v Color) native() *C.GdkColor {
	return (*C.GdkColor)(v.Ptr)
}
func wrapColor(p *C.GdkColor) Color {
	return Color{Ptr: unsafe.Pointer(p)}
}
func WrapColor(p unsafe.Pointer) Color {
	return Color{Ptr: p}
}
func (v Color) IsNil() bool {
	return v.Ptr == nil
}
func IWrapColor(p unsafe.Pointer) interface{} {
	return WrapColor(p)
}

// Struct EventAny
type EventAny struct {
	Ptr unsafe.Pointer
}

func (v EventAny) native() *C.GdkEventAny {
	return (*C.GdkEventAny)(v.Ptr)
}
func wrapEventAny(p *C.GdkEventAny) EventAny {
	return EventAny{Ptr: unsafe.Pointer(p)}
}
func WrapEventAny(p unsafe.Pointer) EventAny {
	return EventAny{Ptr: p}
}
func (v EventAny) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventAny(p unsafe.Pointer) interface{} {
	return WrapEventAny(p)
}

// Struct EventButton
type EventButton struct {
	Ptr unsafe.Pointer
}

func (v EventButton) native() *C.GdkEventButton {
	return (*C.GdkEventButton)(v.Ptr)
}
func wrapEventButton(p *C.GdkEventButton) EventButton {
	return EventButton{Ptr: unsafe.Pointer(p)}
}
func WrapEventButton(p unsafe.Pointer) EventButton {
	return EventButton{Ptr: p}
}
func (v EventButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventButton(p unsafe.Pointer) interface{} {
	return WrapEventButton(p)
}

// Struct EventConfigure
type EventConfigure struct {
	Ptr unsafe.Pointer
}

func (v EventConfigure) native() *C.GdkEventConfigure {
	return (*C.GdkEventConfigure)(v.Ptr)
}
func wrapEventConfigure(p *C.GdkEventConfigure) EventConfigure {
	return EventConfigure{Ptr: unsafe.Pointer(p)}
}
func WrapEventConfigure(p unsafe.Pointer) EventConfigure {
	return EventConfigure{Ptr: p}
}
func (v EventConfigure) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventConfigure(p unsafe.Pointer) interface{} {
	return WrapEventConfigure(p)
}

// Struct EventCrossing
type EventCrossing struct {
	Ptr unsafe.Pointer
}

func (v EventCrossing) native() *C.GdkEventCrossing {
	return (*C.GdkEventCrossing)(v.Ptr)
}
func wrapEventCrossing(p *C.GdkEventCrossing) EventCrossing {
	return EventCrossing{Ptr: unsafe.Pointer(p)}
}
func WrapEventCrossing(p unsafe.Pointer) EventCrossing {
	return EventCrossing{Ptr: p}
}
func (v EventCrossing) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventCrossing(p unsafe.Pointer) interface{} {
	return WrapEventCrossing(p)
}

// Struct EventDND
type EventDND struct {
	Ptr unsafe.Pointer
}

func (v EventDND) native() *C.GdkEventDND {
	return (*C.GdkEventDND)(v.Ptr)
}
func wrapEventDND(p *C.GdkEventDND) EventDND {
	return EventDND{Ptr: unsafe.Pointer(p)}
}
func WrapEventDND(p unsafe.Pointer) EventDND {
	return EventDND{Ptr: p}
}
func (v EventDND) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventDND(p unsafe.Pointer) interface{} {
	return WrapEventDND(p)
}

// Struct EventExpose
type EventExpose struct {
	Ptr unsafe.Pointer
}

func (v EventExpose) native() *C.GdkEventExpose {
	return (*C.GdkEventExpose)(v.Ptr)
}
func wrapEventExpose(p *C.GdkEventExpose) EventExpose {
	return EventExpose{Ptr: unsafe.Pointer(p)}
}
func WrapEventExpose(p unsafe.Pointer) EventExpose {
	return EventExpose{Ptr: p}
}
func (v EventExpose) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventExpose(p unsafe.Pointer) interface{} {
	return WrapEventExpose(p)
}

// Struct EventFocus
type EventFocus struct {
	Ptr unsafe.Pointer
}

func (v EventFocus) native() *C.GdkEventFocus {
	return (*C.GdkEventFocus)(v.Ptr)
}
func wrapEventFocus(p *C.GdkEventFocus) EventFocus {
	return EventFocus{Ptr: unsafe.Pointer(p)}
}
func WrapEventFocus(p unsafe.Pointer) EventFocus {
	return EventFocus{Ptr: p}
}
func (v EventFocus) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventFocus(p unsafe.Pointer) interface{} {
	return WrapEventFocus(p)
}

// Struct EventGrabBroken
type EventGrabBroken struct {
	Ptr unsafe.Pointer
}

func (v EventGrabBroken) native() *C.GdkEventGrabBroken {
	return (*C.GdkEventGrabBroken)(v.Ptr)
}
func wrapEventGrabBroken(p *C.GdkEventGrabBroken) EventGrabBroken {
	return EventGrabBroken{Ptr: unsafe.Pointer(p)}
}
func WrapEventGrabBroken(p unsafe.Pointer) EventGrabBroken {
	return EventGrabBroken{Ptr: p}
}
func (v EventGrabBroken) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventGrabBroken(p unsafe.Pointer) interface{} {
	return WrapEventGrabBroken(p)
}

// Struct EventKey
type EventKey struct {
	Ptr unsafe.Pointer
}

func (v EventKey) native() *C.GdkEventKey {
	return (*C.GdkEventKey)(v.Ptr)
}
func wrapEventKey(p *C.GdkEventKey) EventKey {
	return EventKey{Ptr: unsafe.Pointer(p)}
}
func WrapEventKey(p unsafe.Pointer) EventKey {
	return EventKey{Ptr: p}
}
func (v EventKey) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventKey(p unsafe.Pointer) interface{} {
	return WrapEventKey(p)
}

// Struct EventMotion
type EventMotion struct {
	Ptr unsafe.Pointer
}

func (v EventMotion) native() *C.GdkEventMotion {
	return (*C.GdkEventMotion)(v.Ptr)
}
func wrapEventMotion(p *C.GdkEventMotion) EventMotion {
	return EventMotion{Ptr: unsafe.Pointer(p)}
}
func WrapEventMotion(p unsafe.Pointer) EventMotion {
	return EventMotion{Ptr: p}
}
func (v EventMotion) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventMotion(p unsafe.Pointer) interface{} {
	return WrapEventMotion(p)
}

// Struct EventOwnerChange
type EventOwnerChange struct {
	Ptr unsafe.Pointer
}

func (v EventOwnerChange) native() *C.GdkEventOwnerChange {
	return (*C.GdkEventOwnerChange)(v.Ptr)
}
func wrapEventOwnerChange(p *C.GdkEventOwnerChange) EventOwnerChange {
	return EventOwnerChange{Ptr: unsafe.Pointer(p)}
}
func WrapEventOwnerChange(p unsafe.Pointer) EventOwnerChange {
	return EventOwnerChange{Ptr: p}
}
func (v EventOwnerChange) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventOwnerChange(p unsafe.Pointer) interface{} {
	return WrapEventOwnerChange(p)
}

// Struct EventPadAxis
type EventPadAxis struct {
	Ptr unsafe.Pointer
}

func (v EventPadAxis) native() *C.GdkEventPadAxis {
	return (*C.GdkEventPadAxis)(v.Ptr)
}
func wrapEventPadAxis(p *C.GdkEventPadAxis) EventPadAxis {
	return EventPadAxis{Ptr: unsafe.Pointer(p)}
}
func WrapEventPadAxis(p unsafe.Pointer) EventPadAxis {
	return EventPadAxis{Ptr: p}
}
func (v EventPadAxis) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventPadAxis(p unsafe.Pointer) interface{} {
	return WrapEventPadAxis(p)
}

// Struct EventPadButton
type EventPadButton struct {
	Ptr unsafe.Pointer
}

func (v EventPadButton) native() *C.GdkEventPadButton {
	return (*C.GdkEventPadButton)(v.Ptr)
}
func wrapEventPadButton(p *C.GdkEventPadButton) EventPadButton {
	return EventPadButton{Ptr: unsafe.Pointer(p)}
}
func WrapEventPadButton(p unsafe.Pointer) EventPadButton {
	return EventPadButton{Ptr: p}
}
func (v EventPadButton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventPadButton(p unsafe.Pointer) interface{} {
	return WrapEventPadButton(p)
}

// Struct EventPadGroupMode
type EventPadGroupMode struct {
	Ptr unsafe.Pointer
}

func (v EventPadGroupMode) native() *C.GdkEventPadGroupMode {
	return (*C.GdkEventPadGroupMode)(v.Ptr)
}
func wrapEventPadGroupMode(p *C.GdkEventPadGroupMode) EventPadGroupMode {
	return EventPadGroupMode{Ptr: unsafe.Pointer(p)}
}
func WrapEventPadGroupMode(p unsafe.Pointer) EventPadGroupMode {
	return EventPadGroupMode{Ptr: p}
}
func (v EventPadGroupMode) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventPadGroupMode(p unsafe.Pointer) interface{} {
	return WrapEventPadGroupMode(p)
}

// Struct EventProperty
type EventProperty struct {
	Ptr unsafe.Pointer
}

func (v EventProperty) native() *C.GdkEventProperty {
	return (*C.GdkEventProperty)(v.Ptr)
}
func wrapEventProperty(p *C.GdkEventProperty) EventProperty {
	return EventProperty{Ptr: unsafe.Pointer(p)}
}
func WrapEventProperty(p unsafe.Pointer) EventProperty {
	return EventProperty{Ptr: p}
}
func (v EventProperty) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventProperty(p unsafe.Pointer) interface{} {
	return WrapEventProperty(p)
}

// Struct EventProximity
type EventProximity struct {
	Ptr unsafe.Pointer
}

func (v EventProximity) native() *C.GdkEventProximity {
	return (*C.GdkEventProximity)(v.Ptr)
}
func wrapEventProximity(p *C.GdkEventProximity) EventProximity {
	return EventProximity{Ptr: unsafe.Pointer(p)}
}
func WrapEventProximity(p unsafe.Pointer) EventProximity {
	return EventProximity{Ptr: p}
}
func (v EventProximity) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventProximity(p unsafe.Pointer) interface{} {
	return WrapEventProximity(p)
}

// Struct EventScroll
type EventScroll struct {
	Ptr unsafe.Pointer
}

func (v EventScroll) native() *C.GdkEventScroll {
	return (*C.GdkEventScroll)(v.Ptr)
}
func wrapEventScroll(p *C.GdkEventScroll) EventScroll {
	return EventScroll{Ptr: unsafe.Pointer(p)}
}
func WrapEventScroll(p unsafe.Pointer) EventScroll {
	return EventScroll{Ptr: p}
}
func (v EventScroll) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventScroll(p unsafe.Pointer) interface{} {
	return WrapEventScroll(p)
}

// Struct EventSelection
type EventSelection struct {
	Ptr unsafe.Pointer
}

func (v EventSelection) native() *C.GdkEventSelection {
	return (*C.GdkEventSelection)(v.Ptr)
}
func wrapEventSelection(p *C.GdkEventSelection) EventSelection {
	return EventSelection{Ptr: unsafe.Pointer(p)}
}
func WrapEventSelection(p unsafe.Pointer) EventSelection {
	return EventSelection{Ptr: p}
}
func (v EventSelection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventSelection(p unsafe.Pointer) interface{} {
	return WrapEventSelection(p)
}

// Struct EventSequence
type EventSequence struct {
	Ptr unsafe.Pointer
}

func (v EventSequence) native() *C.GdkEventSequence {
	return (*C.GdkEventSequence)(v.Ptr)
}
func wrapEventSequence(p *C.GdkEventSequence) EventSequence {
	return EventSequence{Ptr: unsafe.Pointer(p)}
}
func WrapEventSequence(p unsafe.Pointer) EventSequence {
	return EventSequence{Ptr: p}
}
func (v EventSequence) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventSequence(p unsafe.Pointer) interface{} {
	return WrapEventSequence(p)
}

// Struct EventSetting
type EventSetting struct {
	Ptr unsafe.Pointer
}

func (v EventSetting) native() *C.GdkEventSetting {
	return (*C.GdkEventSetting)(v.Ptr)
}
func wrapEventSetting(p *C.GdkEventSetting) EventSetting {
	return EventSetting{Ptr: unsafe.Pointer(p)}
}
func WrapEventSetting(p unsafe.Pointer) EventSetting {
	return EventSetting{Ptr: p}
}
func (v EventSetting) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventSetting(p unsafe.Pointer) interface{} {
	return WrapEventSetting(p)
}

// Struct EventTouch
type EventTouch struct {
	Ptr unsafe.Pointer
}

func (v EventTouch) native() *C.GdkEventTouch {
	return (*C.GdkEventTouch)(v.Ptr)
}
func wrapEventTouch(p *C.GdkEventTouch) EventTouch {
	return EventTouch{Ptr: unsafe.Pointer(p)}
}
func WrapEventTouch(p unsafe.Pointer) EventTouch {
	return EventTouch{Ptr: p}
}
func (v EventTouch) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventTouch(p unsafe.Pointer) interface{} {
	return WrapEventTouch(p)
}

// Struct EventTouchpadPinch
type EventTouchpadPinch struct {
	Ptr unsafe.Pointer
}

func (v EventTouchpadPinch) native() *C.GdkEventTouchpadPinch {
	return (*C.GdkEventTouchpadPinch)(v.Ptr)
}
func wrapEventTouchpadPinch(p *C.GdkEventTouchpadPinch) EventTouchpadPinch {
	return EventTouchpadPinch{Ptr: unsafe.Pointer(p)}
}
func WrapEventTouchpadPinch(p unsafe.Pointer) EventTouchpadPinch {
	return EventTouchpadPinch{Ptr: p}
}
func (v EventTouchpadPinch) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventTouchpadPinch(p unsafe.Pointer) interface{} {
	return WrapEventTouchpadPinch(p)
}

// Struct EventTouchpadSwipe
type EventTouchpadSwipe struct {
	Ptr unsafe.Pointer
}

func (v EventTouchpadSwipe) native() *C.GdkEventTouchpadSwipe {
	return (*C.GdkEventTouchpadSwipe)(v.Ptr)
}
func wrapEventTouchpadSwipe(p *C.GdkEventTouchpadSwipe) EventTouchpadSwipe {
	return EventTouchpadSwipe{Ptr: unsafe.Pointer(p)}
}
func WrapEventTouchpadSwipe(p unsafe.Pointer) EventTouchpadSwipe {
	return EventTouchpadSwipe{Ptr: p}
}
func (v EventTouchpadSwipe) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventTouchpadSwipe(p unsafe.Pointer) interface{} {
	return WrapEventTouchpadSwipe(p)
}

// Struct EventVisibility
type EventVisibility struct {
	Ptr unsafe.Pointer
}

func (v EventVisibility) native() *C.GdkEventVisibility {
	return (*C.GdkEventVisibility)(v.Ptr)
}
func wrapEventVisibility(p *C.GdkEventVisibility) EventVisibility {
	return EventVisibility{Ptr: unsafe.Pointer(p)}
}
func WrapEventVisibility(p unsafe.Pointer) EventVisibility {
	return EventVisibility{Ptr: p}
}
func (v EventVisibility) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventVisibility(p unsafe.Pointer) interface{} {
	return WrapEventVisibility(p)
}

// Struct EventWindowState
type EventWindowState struct {
	Ptr unsafe.Pointer
}

func (v EventWindowState) native() *C.GdkEventWindowState {
	return (*C.GdkEventWindowState)(v.Ptr)
}
func wrapEventWindowState(p *C.GdkEventWindowState) EventWindowState {
	return EventWindowState{Ptr: unsafe.Pointer(p)}
}
func WrapEventWindowState(p unsafe.Pointer) EventWindowState {
	return EventWindowState{Ptr: p}
}
func (v EventWindowState) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEventWindowState(p unsafe.Pointer) interface{} {
	return WrapEventWindowState(p)
}

// Struct FrameTimings
type FrameTimings struct {
	Ptr unsafe.Pointer
}

func (v FrameTimings) native() *C.GdkFrameTimings {
	return (*C.GdkFrameTimings)(v.Ptr)
}
func wrapFrameTimings(p *C.GdkFrameTimings) FrameTimings {
	return FrameTimings{Ptr: unsafe.Pointer(p)}
}
func WrapFrameTimings(p unsafe.Pointer) FrameTimings {
	return FrameTimings{Ptr: p}
}
func (v FrameTimings) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFrameTimings(p unsafe.Pointer) interface{} {
	return WrapFrameTimings(p)
}

// GetComplete is a wrapper around gdk_frame_timings_get_complete().
func (timings FrameTimings) GetComplete() bool {
	ret0 := C.gdk_frame_timings_get_complete(timings.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFrameCounter is a wrapper around gdk_frame_timings_get_frame_counter().
func (timings FrameTimings) GetFrameCounter() int64 {
	ret0 := C.gdk_frame_timings_get_frame_counter(timings.native())
	return int64(ret0)
}

// GetFrameTime is a wrapper around gdk_frame_timings_get_frame_time().
func (timings FrameTimings) GetFrameTime() int64 {
	ret0 := C.gdk_frame_timings_get_frame_time(timings.native())
	return int64(ret0)
}

// GetPredictedPresentationTime is a wrapper around gdk_frame_timings_get_predicted_presentation_time().
func (timings FrameTimings) GetPredictedPresentationTime() int64 {
	ret0 := C.gdk_frame_timings_get_predicted_presentation_time(timings.native())
	return int64(ret0)
}

// GetPresentationTime is a wrapper around gdk_frame_timings_get_presentation_time().
func (timings FrameTimings) GetPresentationTime() int64 {
	ret0 := C.gdk_frame_timings_get_presentation_time(timings.native())
	return int64(ret0)
}

// GetRefreshInterval is a wrapper around gdk_frame_timings_get_refresh_interval().
func (timings FrameTimings) GetRefreshInterval() int64 {
	ret0 := C.gdk_frame_timings_get_refresh_interval(timings.native())
	return int64(ret0)
}

// Ref is a wrapper around gdk_frame_timings_ref().
func (timings FrameTimings) Ref() FrameTimings {
	ret0 := C.gdk_frame_timings_ref(timings.native())
	return wrapFrameTimings(ret0)
}

// Unref is a wrapper around gdk_frame_timings_unref().
func (timings FrameTimings) Unref() {
	C.gdk_frame_timings_unref(timings.native())
}

// Struct Geometry
type Geometry struct {
	Ptr unsafe.Pointer
}

func (v Geometry) native() *C.GdkGeometry {
	return (*C.GdkGeometry)(v.Ptr)
}
func wrapGeometry(p *C.GdkGeometry) Geometry {
	return Geometry{Ptr: unsafe.Pointer(p)}
}
func WrapGeometry(p unsafe.Pointer) Geometry {
	return Geometry{Ptr: p}
}
func (v Geometry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGeometry(p unsafe.Pointer) interface{} {
	return WrapGeometry(p)
}

// Struct KeymapKey
type KeymapKey struct {
	Ptr unsafe.Pointer
}

func (v KeymapKey) native() *C.GdkKeymapKey {
	return (*C.GdkKeymapKey)(v.Ptr)
}
func wrapKeymapKey(p *C.GdkKeymapKey) KeymapKey {
	return KeymapKey{Ptr: unsafe.Pointer(p)}
}
func WrapKeymapKey(p unsafe.Pointer) KeymapKey {
	return KeymapKey{Ptr: p}
}
func (v KeymapKey) IsNil() bool {
	return v.Ptr == nil
}
func IWrapKeymapKey(p unsafe.Pointer) interface{} {
	return WrapKeymapKey(p)
}

// Struct Point
type Point struct {
	Ptr unsafe.Pointer
}

func (v Point) native() *C.GdkPoint {
	return (*C.GdkPoint)(v.Ptr)
}
func wrapPoint(p *C.GdkPoint) Point {
	return Point{Ptr: unsafe.Pointer(p)}
}
func WrapPoint(p unsafe.Pointer) Point {
	return Point{Ptr: p}
}
func (v Point) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPoint(p unsafe.Pointer) interface{} {
	return WrapPoint(p)
}

// Struct RGBA
type RGBA struct {
	Ptr unsafe.Pointer
}

func (v RGBA) native() *C.GdkRGBA {
	return (*C.GdkRGBA)(v.Ptr)
}
func wrapRGBA(p *C.GdkRGBA) RGBA {
	return RGBA{Ptr: unsafe.Pointer(p)}
}
func WrapRGBA(p unsafe.Pointer) RGBA {
	return RGBA{Ptr: p}
}
func (v RGBA) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRGBA(p unsafe.Pointer) interface{} {
	return WrapRGBA(p)
}

// Copy is a wrapper around gdk_rgba_copy().
func (rgba RGBA) Copy() RGBA {
	ret0 := C.gdk_rgba_copy(rgba.native())
	return wrapRGBA(ret0)
}

// Equal is a wrapper around gdk_rgba_equal().
func (p1 RGBA) Equal(p2 RGBA) bool {
	ret0 := C.gdk_rgba_equal(C.gconstpointer(p1.native()), C.gconstpointer(p2.native()))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Free is a wrapper around gdk_rgba_free().
func (rgba RGBA) Free() {
	C.gdk_rgba_free(rgba.native())
}

// Hash is a wrapper around gdk_rgba_hash().
func (p RGBA) Hash() uint {
	ret0 := C.gdk_rgba_hash(C.gconstpointer(p.native()))
	return uint(ret0)
}

// Parse is a wrapper around gdk_rgba_parse().
func (rgba RGBA) Parse(spec string) bool {
	spec0 := (*C.gchar)(C.CString(spec))
	ret0 := C.gdk_rgba_parse(rgba.native(), spec0)
	C.free(unsafe.Pointer(spec0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ToString is a wrapper around gdk_rgba_to_string().
func (rgba RGBA) ToString() string {
	ret0 := C.gdk_rgba_to_string(rgba.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Struct Rectangle
type Rectangle struct {
	Ptr unsafe.Pointer
}

func (v Rectangle) native() *C.GdkRectangle {
	return (*C.GdkRectangle)(v.Ptr)
}
func wrapRectangle(p *C.GdkRectangle) Rectangle {
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

// Equal is a wrapper around gdk_rectangle_equal().
func (rect1 Rectangle) Equal(rect2 Rectangle) bool {
	ret0 := C.gdk_rectangle_equal(rect1.native(), rect2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Intersect is a wrapper around gdk_rectangle_intersect().
func (src1 Rectangle) Intersect(src2 Rectangle) (bool, Rectangle) {
	var dest0 C.GdkRectangle
	ret0 := C.gdk_rectangle_intersect(src1.native(), src2.native(), &dest0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapRectangle(&dest0)
}

// Union is a wrapper around gdk_rectangle_union().
func (src1 Rectangle) Union(src2 Rectangle) Rectangle {
	var dest0 C.GdkRectangle
	C.gdk_rectangle_union(src1.native(), src2.native(), &dest0)
	return wrapRectangle(&dest0)
}

// Struct TimeCoord
type TimeCoord struct {
	Ptr unsafe.Pointer
}

func (v TimeCoord) native() *C.GdkTimeCoord {
	return (*C.GdkTimeCoord)(v.Ptr)
}
func wrapTimeCoord(p *C.GdkTimeCoord) TimeCoord {
	return TimeCoord{Ptr: unsafe.Pointer(p)}
}
func WrapTimeCoord(p unsafe.Pointer) TimeCoord {
	return TimeCoord{Ptr: p}
}
func (v TimeCoord) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTimeCoord(p unsafe.Pointer) interface{} {
	return WrapTimeCoord(p)
}

// Struct WindowAttr
type WindowAttr struct {
	Ptr unsafe.Pointer
}

func (v WindowAttr) native() *C.GdkWindowAttr {
	return (*C.GdkWindowAttr)(v.Ptr)
}
func wrapWindowAttr(p *C.GdkWindowAttr) WindowAttr {
	return WindowAttr{Ptr: unsafe.Pointer(p)}
}
func WrapWindowAttr(p unsafe.Pointer) WindowAttr {
	return WindowAttr{Ptr: p}
}
func (v WindowAttr) IsNil() bool {
	return v.Ptr == nil
}
func IWrapWindowAttr(p unsafe.Pointer) interface{} {
	return WrapWindowAttr(p)
}

// Interface DevicePad
type DevicePad struct {
	DevicePadIface
	Ptr unsafe.Pointer
}
type DevicePadIface struct{}

func (v *DevicePadIface) native() *C.GdkDevicePad {
	return (*C.GdkDevicePad)(*(*unsafe.Pointer)(unsafe.Pointer(v)))
}
func wrapDevicePad(p *C.GdkDevicePad) DevicePad {
	return DevicePad{Ptr: unsafe.Pointer(p)}
}
func WrapDevicePad(p unsafe.Pointer) DevicePad {
	return DevicePad{Ptr: p}
}
func (v DevicePad) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDevicePad(p unsafe.Pointer) interface{} {
	return WrapDevicePad(p)
}
func (v DevicePad) GetType() gobject.Type {
	return gobject.Type(C.gdk_device_pad_get_type())
}
func (v DevicePad) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDevicePad(unsafe.Pointer(ptr)), nil
	}
}

// GetFeatureGroup is a wrapper around gdk_device_pad_get_feature_group().
func (pad *DevicePadIface) GetFeatureGroup(feature DevicePadFeature, feature_idx int) int {
	ret0 := C.gdk_device_pad_get_feature_group(pad.native(), C.GdkDevicePadFeature(feature), C.gint(feature_idx))
	return int(ret0)
}

// GetGroupNModes is a wrapper around gdk_device_pad_get_group_n_modes().
func (pad *DevicePadIface) GetGroupNModes(group_idx int) int {
	ret0 := C.gdk_device_pad_get_group_n_modes(pad.native(), C.gint(group_idx))
	return int(ret0)
}

// GetNFeatures is a wrapper around gdk_device_pad_get_n_features().
func (pad *DevicePadIface) GetNFeatures(feature DevicePadFeature) int {
	ret0 := C.gdk_device_pad_get_n_features(pad.native(), C.GdkDevicePadFeature(feature))
	return int(ret0)
}

// GetNGroups is a wrapper around gdk_device_pad_get_n_groups().
func (pad *DevicePadIface) GetNGroups() int {
	ret0 := C.gdk_device_pad_get_n_groups(pad.native())
	return int(ret0)
}

// Object AppLaunchContext
type AppLaunchContext struct {
	gio.AppLaunchContext
}

func (v AppLaunchContext) native() *C.GdkAppLaunchContext {
	return (*C.GdkAppLaunchContext)(v.Ptr)
}
func wrapAppLaunchContext(p *C.GdkAppLaunchContext) (v AppLaunchContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAppLaunchContext(p unsafe.Pointer) (v AppLaunchContext) {
	v.Ptr = p
	return
}
func (v AppLaunchContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppLaunchContext(p unsafe.Pointer) interface{} {
	return WrapAppLaunchContext(p)
}
func (v AppLaunchContext) GetType() gobject.Type {
	return gobject.Type(C.gdk_app_launch_context_get_type())
}
func (v AppLaunchContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppLaunchContext(unsafe.Pointer(ptr)), nil
	}
}

// SetDesktop is a wrapper around gdk_app_launch_context_set_desktop().
func (context AppLaunchContext) SetDesktop(desktop int) {
	C.gdk_app_launch_context_set_desktop(context.native(), C.gint(desktop))
}

// SetIcon is a wrapper around gdk_app_launch_context_set_icon().
func (context AppLaunchContext) SetIcon(icon gio.Icon) {
	C.gdk_app_launch_context_set_icon(context.native(), (*C.GIcon)(icon.Ptr))
}

// SetIconName is a wrapper around gdk_app_launch_context_set_icon_name().
func (context AppLaunchContext) SetIconName(icon_name string) {
	icon_name0 := C.CString(icon_name)
	C.gdk_app_launch_context_set_icon_name(context.native(), icon_name0)
	C.free(unsafe.Pointer(icon_name0)) /*ch:<stdlib.h>*/
}

// SetScreen is a wrapper around gdk_app_launch_context_set_screen().
func (context AppLaunchContext) SetScreen(screen Screen) {
	C.gdk_app_launch_context_set_screen(context.native(), screen.native())
}

// SetTimestamp is a wrapper around gdk_app_launch_context_set_timestamp().
func (context AppLaunchContext) SetTimestamp(timestamp uint32) {
	C.gdk_app_launch_context_set_timestamp(context.native(), C.guint32(timestamp))
}

// Object Screen
type Screen struct {
	gobject.Object
}

func (v Screen) native() *C.GdkScreen {
	return (*C.GdkScreen)(v.Ptr)
}
func wrapScreen(p *C.GdkScreen) (v Screen) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapScreen(p unsafe.Pointer) (v Screen) {
	v.Ptr = p
	return
}
func (v Screen) IsNil() bool {
	return v.Ptr == nil
}
func IWrapScreen(p unsafe.Pointer) interface{} {
	return WrapScreen(p)
}
func (v Screen) GetType() gobject.Type {
	return gobject.Type(C.gdk_screen_get_type())
}
func (v Screen) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapScreen(unsafe.Pointer(ptr)), nil
	}
}

// GetDisplay is a wrapper around gdk_screen_get_display().
func (screen Screen) GetDisplay() Display {
	ret0 := C.gdk_screen_get_display(screen.native())
	return wrapDisplay(ret0)
}

// GetFontOptions is a wrapper around gdk_screen_get_font_options().
func (screen Screen) GetFontOptions() cairo.FontOptions {
	ret0 := C.gdk_screen_get_font_options(screen.native())
	return cairo.WrapFontOptions(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetResolution is a wrapper around gdk_screen_get_resolution().
func (screen Screen) GetResolution() float64 {
	ret0 := C.gdk_screen_get_resolution(screen.native())
	return float64(ret0)
}

// GetRgbaVisual is a wrapper around gdk_screen_get_rgba_visual().
func (screen Screen) GetRgbaVisual() Visual {
	ret0 := C.gdk_screen_get_rgba_visual(screen.native())
	return wrapVisual(ret0)
}

// GetRootWindow is a wrapper around gdk_screen_get_root_window().
func (screen Screen) GetRootWindow() Window {
	ret0 := C.gdk_screen_get_root_window(screen.native())
	return wrapWindow(ret0)
}

// GetSetting is a wrapper around gdk_screen_get_setting().
func (screen Screen) GetSetting(name string, value gobject.Value) bool {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gdk_screen_get_setting(screen.native(), name0, (*C.GValue)(value.Ptr))
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSystemVisual is a wrapper around gdk_screen_get_system_visual().
func (screen Screen) GetSystemVisual() Visual {
	ret0 := C.gdk_screen_get_system_visual(screen.native())
	return wrapVisual(ret0)
}

// GetToplevelWindows is a wrapper around gdk_screen_get_toplevel_windows().
func (screen Screen) GetToplevelWindows() glib.List {
	ret0 := C.gdk_screen_get_toplevel_windows(screen.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// GetWindowStack is a wrapper around gdk_screen_get_window_stack().
func (screen Screen) GetWindowStack() glib.List {
	ret0 := C.gdk_screen_get_window_stack(screen.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// IsComposited is a wrapper around gdk_screen_is_composited().
func (screen Screen) IsComposited() bool {
	ret0 := C.gdk_screen_is_composited(screen.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ListVisuals is a wrapper around gdk_screen_list_visuals().
func (screen Screen) ListVisuals() glib.List {
	ret0 := C.gdk_screen_list_visuals(screen.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapVisual(p) }) /*gir:GLib*/
}

// SetFontOptions is a wrapper around gdk_screen_set_font_options().
func (screen Screen) SetFontOptions(options cairo.FontOptions) {
	C.gdk_screen_set_font_options(screen.native(), (*C.cairo_font_options_t)(options.Ptr))
}

// SetResolution is a wrapper around gdk_screen_set_resolution().
func (screen Screen) SetResolution(dpi float64) {
	C.gdk_screen_set_resolution(screen.native(), C.gdouble(dpi))
}

// ScreenGetDefault is a wrapper around gdk_screen_get_default().
func ScreenGetDefault() Screen {
	ret0 := C.gdk_screen_get_default()
	return wrapScreen(ret0)
}

// Object Display
type Display struct {
	gobject.Object
}

func (v Display) native() *C.GdkDisplay {
	return (*C.GdkDisplay)(v.Ptr)
}
func wrapDisplay(p *C.GdkDisplay) (v Display) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDisplay(p unsafe.Pointer) (v Display) {
	v.Ptr = p
	return
}
func (v Display) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDisplay(p unsafe.Pointer) interface{} {
	return WrapDisplay(p)
}
func (v Display) GetType() gobject.Type {
	return gobject.Type(C.gdk_display_get_type())
}
func (v Display) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDisplay(unsafe.Pointer(ptr)), nil
	}
}

// Beep is a wrapper around gdk_display_beep().
func (display Display) Beep() {
	C.gdk_display_beep(display.native())
}

// Close is a wrapper around gdk_display_close().
func (display Display) Close() {
	C.gdk_display_close(display.native())
}

// DeviceIsGrabbed is a wrapper around gdk_display_device_is_grabbed().
func (display Display) DeviceIsGrabbed(device Device) bool {
	ret0 := C.gdk_display_device_is_grabbed(display.native(), device.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Flush is a wrapper around gdk_display_flush().
func (display Display) Flush() {
	C.gdk_display_flush(display.native())
}

// GetAppLaunchContext is a wrapper around gdk_display_get_app_launch_context().
func (display Display) GetAppLaunchContext() AppLaunchContext {
	ret0 := C.gdk_display_get_app_launch_context(display.native())
	return wrapAppLaunchContext(ret0)
}

// GetDefaultCursorSize is a wrapper around gdk_display_get_default_cursor_size().
func (display Display) GetDefaultCursorSize() uint {
	ret0 := C.gdk_display_get_default_cursor_size(display.native())
	return uint(ret0)
}

// GetDefaultGroup is a wrapper around gdk_display_get_default_group().
func (display Display) GetDefaultGroup() Window {
	ret0 := C.gdk_display_get_default_group(display.native())
	return wrapWindow(ret0)
}

// GetDefaultScreen is a wrapper around gdk_display_get_default_screen().
func (display Display) GetDefaultScreen() Screen {
	ret0 := C.gdk_display_get_default_screen(display.native())
	return wrapScreen(ret0)
}

// GetDefaultSeat is a wrapper around gdk_display_get_default_seat().
func (display Display) GetDefaultSeat() Seat {
	ret0 := C.gdk_display_get_default_seat(display.native())
	return wrapSeat(ret0)
}

// GetMaximalCursorSize is a wrapper around gdk_display_get_maximal_cursor_size().
func (display Display) GetMaximalCursorSize() (uint, uint) {
	var width0 C.guint
	var height0 C.guint
	C.gdk_display_get_maximal_cursor_size(display.native(), &width0, &height0)
	return uint(width0), uint(height0)
}

// GetMonitor is a wrapper around gdk_display_get_monitor().
func (display Display) GetMonitor(monitor_num int) Monitor {
	ret0 := C.gdk_display_get_monitor(display.native(), C.int(monitor_num))
	return wrapMonitor(ret0)
}

// GetMonitorAtPoint is a wrapper around gdk_display_get_monitor_at_point().
func (display Display) GetMonitorAtPoint(x int, y int) Monitor {
	ret0 := C.gdk_display_get_monitor_at_point(display.native(), C.int(x), C.int(y))
	return wrapMonitor(ret0)
}

// GetMonitorAtWindow is a wrapper around gdk_display_get_monitor_at_window().
func (display Display) GetMonitorAtWindow(window Window) Monitor {
	ret0 := C.gdk_display_get_monitor_at_window(display.native(), window.native())
	return wrapMonitor(ret0)
}

// GetNMonitors is a wrapper around gdk_display_get_n_monitors().
func (display Display) GetNMonitors() int {
	ret0 := C.gdk_display_get_n_monitors(display.native())
	return int(ret0)
}

// GetName is a wrapper around gdk_display_get_name().
func (display Display) GetName() string {
	ret0 := C.gdk_display_get_name(display.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPrimaryMonitor is a wrapper around gdk_display_get_primary_monitor().
func (display Display) GetPrimaryMonitor() Monitor {
	ret0 := C.gdk_display_get_primary_monitor(display.native())
	return wrapMonitor(ret0)
}

// HasPending is a wrapper around gdk_display_has_pending().
func (display Display) HasPending() bool {
	ret0 := C.gdk_display_has_pending(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around gdk_display_is_closed().
func (display Display) IsClosed() bool {
	ret0 := C.gdk_display_is_closed(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ListSeats is a wrapper around gdk_display_list_seats().
func (display Display) ListSeats() glib.List {
	ret0 := C.gdk_display_list_seats(display.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapSeat(p) }) /*gir:GLib*/
}

// NotifyStartupComplete is a wrapper around gdk_display_notify_startup_complete().
func (display Display) NotifyStartupComplete(startup_id string) {
	startup_id0 := (*C.gchar)(C.CString(startup_id))
	C.gdk_display_notify_startup_complete(display.native(), startup_id0)
	C.free(unsafe.Pointer(startup_id0)) /*ch:<stdlib.h>*/
}

// SetDoubleClickDistance is a wrapper around gdk_display_set_double_click_distance().
func (display Display) SetDoubleClickDistance(distance uint) {
	C.gdk_display_set_double_click_distance(display.native(), C.guint(distance))
}

// SetDoubleClickTime is a wrapper around gdk_display_set_double_click_time().
func (display Display) SetDoubleClickTime(msec uint) {
	C.gdk_display_set_double_click_time(display.native(), C.guint(msec))
}

// SupportsClipboardPersistence is a wrapper around gdk_display_supports_clipboard_persistence().
func (display Display) SupportsClipboardPersistence() bool {
	ret0 := C.gdk_display_supports_clipboard_persistence(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsCursorAlpha is a wrapper around gdk_display_supports_cursor_alpha().
func (display Display) SupportsCursorAlpha() bool {
	ret0 := C.gdk_display_supports_cursor_alpha(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsCursorColor is a wrapper around gdk_display_supports_cursor_color().
func (display Display) SupportsCursorColor() bool {
	ret0 := C.gdk_display_supports_cursor_color(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsInputShapes is a wrapper around gdk_display_supports_input_shapes().
func (display Display) SupportsInputShapes() bool {
	ret0 := C.gdk_display_supports_input_shapes(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsSelectionNotification is a wrapper around gdk_display_supports_selection_notification().
func (display Display) SupportsSelectionNotification() bool {
	ret0 := C.gdk_display_supports_selection_notification(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsShapes is a wrapper around gdk_display_supports_shapes().
func (display Display) SupportsShapes() bool {
	ret0 := C.gdk_display_supports_shapes(display.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Sync is a wrapper around gdk_display_sync().
func (display Display) Sync() {
	C.gdk_display_sync(display.native())
}

// DisplayGetDefault is a wrapper around gdk_display_get_default().
func DisplayGetDefault() Display {
	ret0 := C.gdk_display_get_default()
	return wrapDisplay(ret0)
}

// DisplayOpen is a wrapper around gdk_display_open().
func DisplayOpen(display_name string) Display {
	display_name0 := (*C.gchar)(C.CString(display_name))
	ret0 := C.gdk_display_open(display_name0)
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
	return wrapDisplay(ret0)
}

// Object Device
type Device struct {
	gobject.Object
}

func (v Device) native() *C.GdkDevice {
	return (*C.GdkDevice)(v.Ptr)
}
func wrapDevice(p *C.GdkDevice) (v Device) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDevice(p unsafe.Pointer) (v Device) {
	v.Ptr = p
	return
}
func (v Device) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDevice(p unsafe.Pointer) interface{} {
	return WrapDevice(p)
}
func (v Device) GetType() gobject.Type {
	return gobject.Type(C.gdk_device_get_type())
}
func (v Device) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDevice(unsafe.Pointer(ptr)), nil
	}
}

// GetAssociatedDevice is a wrapper around gdk_device_get_associated_device().
func (device Device) GetAssociatedDevice() Device {
	ret0 := C.gdk_device_get_associated_device(device.native())
	return wrapDevice(ret0)
}

// GetAxes is a wrapper around gdk_device_get_axes().
func (device Device) GetAxes() AxisFlags {
	ret0 := C.gdk_device_get_axes(device.native())
	return AxisFlags(ret0)
}

// GetAxis is a wrapper around gdk_device_get_axis().
func (device Device) GetAxis(axes []float64, use AxisUse) (bool, float64) {
	axes0 := make([]C.gdouble, len(axes))
	for idx, elemG := range axes {
		axes0[idx] = C.gdouble(elemG)
	}
	var axes0Ptr *C.gdouble
	if len(axes0) > 0 {
		axes0Ptr = &axes0[0]
	}
	var value0 C.gdouble
	ret0 := C.gdk_device_get_axis(device.native(), axes0Ptr, C.GdkAxisUse(use), &value0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, float64(value0)
}

// GetAxisUse is a wrapper around gdk_device_get_axis_use().
func (device Device) GetAxisUse(index_ uint) AxisUse {
	ret0 := C.gdk_device_get_axis_use(device.native(), C.guint(index_))
	return AxisUse(ret0)
}

// GetDeviceType is a wrapper around gdk_device_get_device_type().
func (device Device) GetDeviceType() DeviceType {
	ret0 := C.gdk_device_get_device_type(device.native())
	return DeviceType(ret0)
}

// GetDisplay is a wrapper around gdk_device_get_display().
func (device Device) GetDisplay() Display {
	ret0 := C.gdk_device_get_display(device.native())
	return wrapDisplay(ret0)
}

// GetHasCursor is a wrapper around gdk_device_get_has_cursor().
func (device Device) GetHasCursor() bool {
	ret0 := C.gdk_device_get_has_cursor(device.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetKey is a wrapper around gdk_device_get_key().
func (device Device) GetKey(index_ uint) (bool, uint, ModifierType) {
	var keyval0 C.guint
	var modifiers0 C.GdkModifierType
	ret0 := C.gdk_device_get_key(device.native(), C.guint(index_), &keyval0, &modifiers0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint(keyval0), ModifierType(modifiers0)
}

// GetLastEventWindow is a wrapper around gdk_device_get_last_event_window().
func (device Device) GetLastEventWindow() Window {
	ret0 := C.gdk_device_get_last_event_window(device.native())
	return wrapWindow(ret0)
}

// GetMode is a wrapper around gdk_device_get_mode().
func (device Device) GetMode() InputMode {
	ret0 := C.gdk_device_get_mode(device.native())
	return InputMode(ret0)
}

// GetNAxes is a wrapper around gdk_device_get_n_axes().
func (device Device) GetNAxes() int {
	ret0 := C.gdk_device_get_n_axes(device.native())
	return int(ret0)
}

// GetNKeys is a wrapper around gdk_device_get_n_keys().
func (device Device) GetNKeys() int {
	ret0 := C.gdk_device_get_n_keys(device.native())
	return int(ret0)
}

// GetName is a wrapper around gdk_device_get_name().
func (device Device) GetName() string {
	ret0 := C.gdk_device_get_name(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPosition is a wrapper around gdk_device_get_position().
func (device Device) GetPosition() (Screen, int, int) {
	var screen0 *C.GdkScreen
	var x0 C.gint
	var y0 C.gint
	C.gdk_device_get_position(device.native(), &screen0, &x0, &y0)
	return wrapScreen(screen0), int(x0), int(y0)
}

// GetPositionDouble is a wrapper around gdk_device_get_position_double().
func (device Device) GetPositionDouble() (Screen, float64, float64) {
	var screen0 *C.GdkScreen
	var x0 C.gdouble
	var y0 C.gdouble
	C.gdk_device_get_position_double(device.native(), &screen0, &x0, &y0)
	return wrapScreen(screen0), float64(x0), float64(y0)
}

// GetProductId is a wrapper around gdk_device_get_product_id().
func (device Device) GetProductId() string {
	ret0 := C.gdk_device_get_product_id(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSeat is a wrapper around gdk_device_get_seat().
func (device Device) GetSeat() Seat {
	ret0 := C.gdk_device_get_seat(device.native())
	return wrapSeat(ret0)
}

// GetSource is a wrapper around gdk_device_get_source().
func (device Device) GetSource() InputSource {
	ret0 := C.gdk_device_get_source(device.native())
	return InputSource(ret0)
}

// GetState is a wrapper around gdk_device_get_state().
func (device Device) GetState(window Window, axes []float64) ModifierType {
	axes0 := make([]C.gdouble, len(axes))
	for idx, elemG := range axes {
		axes0[idx] = C.gdouble(elemG)
	}
	var axes0Ptr *C.gdouble
	if len(axes0) > 0 {
		axes0Ptr = &axes0[0]
	}
	var mask0 C.GdkModifierType
	C.gdk_device_get_state(device.native(), window.native(), axes0Ptr, &mask0)
	return ModifierType(mask0)
}

// GetVendorId is a wrapper around gdk_device_get_vendor_id().
func (device Device) GetVendorId() string {
	ret0 := C.gdk_device_get_vendor_id(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetWindowAtPosition is a wrapper around gdk_device_get_window_at_position().
func (device Device) GetWindowAtPosition() (Window, int, int) {
	var win_x0 C.gint
	var win_y0 C.gint
	ret0 := C.gdk_device_get_window_at_position(device.native(), &win_x0, &win_y0)
	return wrapWindow(ret0), int(win_x0), int(win_y0)
}

// GetWindowAtPositionDouble is a wrapper around gdk_device_get_window_at_position_double().
func (device Device) GetWindowAtPositionDouble() (Window, float64, float64) {
	var win_x0 C.gdouble
	var win_y0 C.gdouble
	ret0 := C.gdk_device_get_window_at_position_double(device.native(), &win_x0, &win_y0)
	return wrapWindow(ret0), float64(win_x0), float64(win_y0)
}

// ListAxes is a wrapper around gdk_device_list_axes().
func (device Device) ListAxes() glib.List {
	ret0 := C.gdk_device_list_axes(device.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapAtom(p) }) /*gir:GLib*/
}

// ListSlaveDevices is a wrapper around gdk_device_list_slave_devices().
func (device Device) ListSlaveDevices() glib.List {
	ret0 := C.gdk_device_list_slave_devices(device.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDevice(p) }) /*gir:GLib*/
}

// SetAxisUse is a wrapper around gdk_device_set_axis_use().
func (device Device) SetAxisUse(index_ uint, use AxisUse) {
	C.gdk_device_set_axis_use(device.native(), C.guint(index_), C.GdkAxisUse(use))
}

// SetKey is a wrapper around gdk_device_set_key().
func (device Device) SetKey(index_ uint, keyval uint, modifiers ModifierType) {
	C.gdk_device_set_key(device.native(), C.guint(index_), C.guint(keyval), C.GdkModifierType(modifiers))
}

// SetMode is a wrapper around gdk_device_set_mode().
func (device Device) SetMode(mode InputMode) bool {
	ret0 := C.gdk_device_set_mode(device.native(), C.GdkInputMode(mode))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Warp is a wrapper around gdk_device_warp().
func (device Device) Warp(screen Screen, x int, y int) {
	C.gdk_device_warp(device.native(), screen.native(), C.gint(x), C.gint(y))
}

// DeviceFreeHistory is a wrapper around gdk_device_free_history().
func DeviceFreeHistory(events []TimeCoord) {
	events0 := make([]*C.GdkTimeCoord, len(events))
	for idx, elemG := range events {
		events0[idx] = elemG.native()
	}
	var events0Ptr **C.GdkTimeCoord
	if len(events0) > 0 {
		events0Ptr = &events0[0]
	}
	C.gdk_device_free_history(events0Ptr, C.gint(len(events)))
}

// Object Window
type Window struct {
	gobject.Object
}

func (v Window) native() *C.GdkWindow {
	return (*C.GdkWindow)(v.Ptr)
}
func wrapWindow(p *C.GdkWindow) (v Window) {
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
	return gobject.Type(C.gdk_window_get_type())
}
func (v Window) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapWindow(unsafe.Pointer(ptr)), nil
	}
}

// Beep is a wrapper around gdk_window_beep().
func (window Window) Beep() {
	C.gdk_window_beep(window.native())
}

// BeginMoveDrag is a wrapper around gdk_window_begin_move_drag().
func (window Window) BeginMoveDrag(button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_move_drag(window.native(), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
}

// BeginMoveDragForDevice is a wrapper around gdk_window_begin_move_drag_for_device().
func (window Window) BeginMoveDragForDevice(device Device, button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_move_drag_for_device(window.native(), device.native(), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
}

// BeginResizeDrag is a wrapper around gdk_window_begin_resize_drag().
func (window Window) BeginResizeDrag(edge WindowEdge, button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_resize_drag(window.native(), C.GdkWindowEdge(edge), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
}

// BeginResizeDragForDevice is a wrapper around gdk_window_begin_resize_drag_for_device().
func (window Window) BeginResizeDragForDevice(edge WindowEdge, device Device, button int, root_x int, root_y int, timestamp uint32) {
	C.gdk_window_begin_resize_drag_for_device(window.native(), C.GdkWindowEdge(edge), device.native(), C.gint(button), C.gint(root_x), C.gint(root_y), C.guint32(timestamp))
}

// CoordsFromParent is a wrapper around gdk_window_coords_from_parent().
func (window Window) CoordsFromParent(parent_x float64, parent_y float64) (float64, float64) {
	var x0 C.gdouble
	var y0 C.gdouble
	C.gdk_window_coords_from_parent(window.native(), C.gdouble(parent_x), C.gdouble(parent_y), &x0, &y0)
	return float64(x0), float64(y0)
}

// CoordsToParent is a wrapper around gdk_window_coords_to_parent().
func (window Window) CoordsToParent(x float64, y float64) (float64, float64) {
	var parent_x0 C.gdouble
	var parent_y0 C.gdouble
	C.gdk_window_coords_to_parent(window.native(), C.gdouble(x), C.gdouble(y), &parent_x0, &parent_y0)
	return float64(parent_x0), float64(parent_y0)
}

// CreateGlContext is a wrapper around gdk_window_create_gl_context().
func (window Window) CreateGlContext() (GLContext, error) {
	var err glib.Error
	ret0 := C.gdk_window_create_gl_context(window.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return GLContext{}, err.GoValue()
	}
	return wrapGLContext(ret0), nil
}

// CreateSimilarSurface is a wrapper around gdk_window_create_similar_surface().
func (window Window) CreateSimilarSurface(content /*gir:cairo*/ cairo.Content, width int, height int) cairo.Surface {
	ret0 := C.gdk_window_create_similar_surface(window.native(), C.cairo_content_t(content), C.int(width), C.int(height))
	return cairo.WrapSurface(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// Deiconify is a wrapper around gdk_window_deiconify().
func (window Window) Deiconify() {
	C.gdk_window_deiconify(window.native())
}

// Destroy is a wrapper around gdk_window_destroy().
func (window Window) Destroy() {
	C.gdk_window_destroy(window.native())
}

// EndDrawFrame is a wrapper around gdk_window_end_draw_frame().
func (window Window) EndDrawFrame(context DrawingContext) {
	C.gdk_window_end_draw_frame(window.native(), context.native())
}

// EnsureNative is a wrapper around gdk_window_ensure_native().
func (window Window) EnsureNative() bool {
	ret0 := C.gdk_window_ensure_native(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Focus is a wrapper around gdk_window_focus().
func (window Window) Focus(timestamp uint32) {
	C.gdk_window_focus(window.native(), C.guint32(timestamp))
}

// FreezeUpdates is a wrapper around gdk_window_freeze_updates().
func (window Window) FreezeUpdates() {
	C.gdk_window_freeze_updates(window.native())
}

// Fullscreen is a wrapper around gdk_window_fullscreen().
func (window Window) Fullscreen() {
	C.gdk_window_fullscreen(window.native())
}

// FullscreenOnMonitor is a wrapper around gdk_window_fullscreen_on_monitor().
func (window Window) FullscreenOnMonitor(monitor int) {
	C.gdk_window_fullscreen_on_monitor(window.native(), C.gint(monitor))
}

// GeometryChanged is a wrapper around gdk_window_geometry_changed().
func (window Window) GeometryChanged() {
	C.gdk_window_geometry_changed(window.native())
}

// GetAcceptFocus is a wrapper around gdk_window_get_accept_focus().
func (window Window) GetAcceptFocus() bool {
	ret0 := C.gdk_window_get_accept_focus(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChildren is a wrapper around gdk_window_get_children().
func (window Window) GetChildren() glib.List {
	ret0 := C.gdk_window_get_children(window.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// GetChildrenWithUserData is a wrapper around gdk_window_get_children_with_user_data().
func (window Window) GetChildrenWithUserData(user_data unsafe.Pointer) glib.List {
	ret0 := C.gdk_window_get_children_with_user_data(window.native(), C.gpointer(user_data))
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// GetClipRegion is a wrapper around gdk_window_get_clip_region().
func (window Window) GetClipRegion() cairo.Region {
	ret0 := C.gdk_window_get_clip_region(window.native())
	return cairo.WrapRegion(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetCursor is a wrapper around gdk_window_get_cursor().
func (window Window) GetCursor() Cursor {
	ret0 := C.gdk_window_get_cursor(window.native())
	return wrapCursor(ret0)
}

// GetDecorations is a wrapper around gdk_window_get_decorations().
func (window Window) GetDecorations() (bool, WMDecoration) {
	var decorations0 C.GdkWMDecoration
	ret0 := C.gdk_window_get_decorations(window.native(), &decorations0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, WMDecoration(decorations0)
}

// GetDeviceCursor is a wrapper around gdk_window_get_device_cursor().
func (window Window) GetDeviceCursor(device Device) Cursor {
	ret0 := C.gdk_window_get_device_cursor(window.native(), device.native())
	return wrapCursor(ret0)
}

// GetDeviceEvents is a wrapper around gdk_window_get_device_events().
func (window Window) GetDeviceEvents(device Device) EventMask {
	ret0 := C.gdk_window_get_device_events(window.native(), device.native())
	return EventMask(ret0)
}

// GetDevicePosition is a wrapper around gdk_window_get_device_position().
func (window Window) GetDevicePosition(device Device) (Window, int, int, ModifierType) {
	var x0 C.gint
	var y0 C.gint
	var mask0 C.GdkModifierType
	ret0 := C.gdk_window_get_device_position(window.native(), device.native(), &x0, &y0, &mask0)
	return wrapWindow(ret0), int(x0), int(y0), ModifierType(mask0)
}

// GetDevicePositionDouble is a wrapper around gdk_window_get_device_position_double().
func (window Window) GetDevicePositionDouble(device Device) (Window, float64, float64, ModifierType) {
	var x0 C.gdouble
	var y0 C.gdouble
	var mask0 C.GdkModifierType
	ret0 := C.gdk_window_get_device_position_double(window.native(), device.native(), &x0, &y0, &mask0)
	return wrapWindow(ret0), float64(x0), float64(y0), ModifierType(mask0)
}

// GetDisplay is a wrapper around gdk_window_get_display().
func (window Window) GetDisplay() Display {
	ret0 := C.gdk_window_get_display(window.native())
	return wrapDisplay(ret0)
}

// GetDragProtocol is a wrapper around gdk_window_get_drag_protocol().
func (window Window) GetDragProtocol() (DragProtocol, Window) {
	var target0 *C.GdkWindow
	ret0 := C.gdk_window_get_drag_protocol(window.native(), &target0)
	return DragProtocol(ret0), wrapWindow(target0)
}

// GetEffectiveParent is a wrapper around gdk_window_get_effective_parent().
func (window Window) GetEffectiveParent() Window {
	ret0 := C.gdk_window_get_effective_parent(window.native())
	return wrapWindow(ret0)
}

// GetEffectiveToplevel is a wrapper around gdk_window_get_effective_toplevel().
func (window Window) GetEffectiveToplevel() Window {
	ret0 := C.gdk_window_get_effective_toplevel(window.native())
	return wrapWindow(ret0)
}

// GetEventCompression is a wrapper around gdk_window_get_event_compression().
func (window Window) GetEventCompression() bool {
	ret0 := C.gdk_window_get_event_compression(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetEvents is a wrapper around gdk_window_get_events().
func (window Window) GetEvents() EventMask {
	ret0 := C.gdk_window_get_events(window.native())
	return EventMask(ret0)
}

// GetFocusOnMap is a wrapper around gdk_window_get_focus_on_map().
func (window Window) GetFocusOnMap() bool {
	ret0 := C.gdk_window_get_focus_on_map(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFrameClock is a wrapper around gdk_window_get_frame_clock().
func (window Window) GetFrameClock() FrameClock {
	ret0 := C.gdk_window_get_frame_clock(window.native())
	return wrapFrameClock(ret0)
}

// GetFrameExtents is a wrapper around gdk_window_get_frame_extents().
func (window Window) GetFrameExtents() Rectangle {
	var rect0 C.GdkRectangle
	C.gdk_window_get_frame_extents(window.native(), &rect0)
	return wrapRectangle(&rect0)
}

// GetFullscreenMode is a wrapper around gdk_window_get_fullscreen_mode().
func (window Window) GetFullscreenMode() FullscreenMode {
	ret0 := C.gdk_window_get_fullscreen_mode(window.native())
	return FullscreenMode(ret0)
}

// GetGeometry is a wrapper around gdk_window_get_geometry().
func (window Window) GetGeometry() (int, int, int, int) {
	var x0 C.gint
	var y0 C.gint
	var width0 C.gint
	var height0 C.gint
	C.gdk_window_get_geometry(window.native(), &x0, &y0, &width0, &height0)
	return int(x0), int(y0), int(width0), int(height0)
}

// GetGroup is a wrapper around gdk_window_get_group().
func (window Window) GetGroup() Window {
	ret0 := C.gdk_window_get_group(window.native())
	return wrapWindow(ret0)
}

// GetHeight is a wrapper around gdk_window_get_height().
func (window Window) GetHeight() int {
	ret0 := C.gdk_window_get_height(window.native())
	return int(ret0)
}

// GetModalHint is a wrapper around gdk_window_get_modal_hint().
func (window Window) GetModalHint() bool {
	ret0 := C.gdk_window_get_modal_hint(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetOrigin is a wrapper around gdk_window_get_origin().
func (window Window) GetOrigin() (int, int, int) {
	var x0 C.gint
	var y0 C.gint
	ret0 := C.gdk_window_get_origin(window.native(), &x0, &y0)
	return int(ret0), int(x0), int(y0)
}

// GetParent is a wrapper around gdk_window_get_parent().
func (window Window) GetParent() Window {
	ret0 := C.gdk_window_get_parent(window.native())
	return wrapWindow(ret0)
}

// GetPassThrough is a wrapper around gdk_window_get_pass_through().
func (window Window) GetPassThrough() bool {
	ret0 := C.gdk_window_get_pass_through(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPosition is a wrapper around gdk_window_get_position().
func (window Window) GetPosition() (int, int) {
	var x0 C.gint
	var y0 C.gint
	C.gdk_window_get_position(window.native(), &x0, &y0)
	return int(x0), int(y0)
}

// GetRootCoords is a wrapper around gdk_window_get_root_coords().
func (window Window) GetRootCoords(x int, y int) (int, int) {
	var root_x0 C.gint
	var root_y0 C.gint
	C.gdk_window_get_root_coords(window.native(), C.gint(x), C.gint(y), &root_x0, &root_y0)
	return int(root_x0), int(root_y0)
}

// GetRootOrigin is a wrapper around gdk_window_get_root_origin().
func (window Window) GetRootOrigin() (int, int) {
	var x0 C.gint
	var y0 C.gint
	C.gdk_window_get_root_origin(window.native(), &x0, &y0)
	return int(x0), int(y0)
}

// GetScaleFactor is a wrapper around gdk_window_get_scale_factor().
func (window Window) GetScaleFactor() int {
	ret0 := C.gdk_window_get_scale_factor(window.native())
	return int(ret0)
}

// GetScreen is a wrapper around gdk_window_get_screen().
func (window Window) GetScreen() Screen {
	ret0 := C.gdk_window_get_screen(window.native())
	return wrapScreen(ret0)
}

// GetSourceEvents is a wrapper around gdk_window_get_source_events().
func (window Window) GetSourceEvents(source InputSource) EventMask {
	ret0 := C.gdk_window_get_source_events(window.native(), C.GdkInputSource(source))
	return EventMask(ret0)
}

// GetState is a wrapper around gdk_window_get_state().
func (window Window) GetState() WindowState {
	ret0 := C.gdk_window_get_state(window.native())
	return WindowState(ret0)
}

// GetSupportMultidevice is a wrapper around gdk_window_get_support_multidevice().
func (window Window) GetSupportMultidevice() bool {
	ret0 := C.gdk_window_get_support_multidevice(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetToplevel is a wrapper around gdk_window_get_toplevel().
func (window Window) GetToplevel() Window {
	ret0 := C.gdk_window_get_toplevel(window.native())
	return wrapWindow(ret0)
}

// GetTypeHint is a wrapper around gdk_window_get_type_hint().
func (window Window) GetTypeHint() WindowTypeHint {
	ret0 := C.gdk_window_get_type_hint(window.native())
	return WindowTypeHint(ret0)
}

// GetUpdateArea is a wrapper around gdk_window_get_update_area().
func (window Window) GetUpdateArea() cairo.Region {
	ret0 := C.gdk_window_get_update_area(window.native())
	return cairo.WrapRegion(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetUserData is a wrapper around gdk_window_get_user_data().
func (window Window) GetUserData() unsafe.Pointer {
	var data0 C.gpointer
	C.gdk_window_get_user_data(window.native(), &data0)
	return unsafe.Pointer(data0)
}

// GetVisibleRegion is a wrapper around gdk_window_get_visible_region().
func (window Window) GetVisibleRegion() cairo.Region {
	ret0 := C.gdk_window_get_visible_region(window.native())
	return cairo.WrapRegion(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetVisual is a wrapper around gdk_window_get_visual().
func (window Window) GetVisual() Visual {
	ret0 := C.gdk_window_get_visual(window.native())
	return wrapVisual(ret0)
}

// GetWidth is a wrapper around gdk_window_get_width().
func (window Window) GetWidth() int {
	ret0 := C.gdk_window_get_width(window.native())
	return int(ret0)
}

// GetWindowType is a wrapper around gdk_window_get_window_type().
func (window Window) GetWindowType() WindowType {
	ret0 := C.gdk_window_get_window_type(window.native())
	return WindowType(ret0)
}

// HasNative is a wrapper around gdk_window_has_native().
func (window Window) HasNative() bool {
	ret0 := C.gdk_window_has_native(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Hide is a wrapper around gdk_window_hide().
func (window Window) Hide() {
	C.gdk_window_hide(window.native())
}

// Iconify is a wrapper around gdk_window_iconify().
func (window Window) Iconify() {
	C.gdk_window_iconify(window.native())
}

// InputShapeCombineRegion is a wrapper around gdk_window_input_shape_combine_region().
func (window Window) InputShapeCombineRegion(shape_region cairo.Region, offset_x int, offset_y int) {
	C.gdk_window_input_shape_combine_region(window.native(), (*C.cairo_region_t)(shape_region.Ptr), C.gint(offset_x), C.gint(offset_y))
}

// InvalidateRect is a wrapper around gdk_window_invalidate_rect().
func (window Window) InvalidateRect(rect Rectangle, invalidate_children bool) {
	C.gdk_window_invalidate_rect(window.native(), rect.native(), C.gboolean(util.Bool2Int(invalidate_children)) /*go:.util*/)
}

// InvalidateRegion is a wrapper around gdk_window_invalidate_region().
func (window Window) InvalidateRegion(region cairo.Region, invalidate_children bool) {
	C.gdk_window_invalidate_region(window.native(), (*C.cairo_region_t)(region.Ptr), C.gboolean(util.Bool2Int(invalidate_children)) /*go:.util*/)
}

// IsDestroyed is a wrapper around gdk_window_is_destroyed().
func (window Window) IsDestroyed() bool {
	ret0 := C.gdk_window_is_destroyed(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsInputOnly is a wrapper around gdk_window_is_input_only().
func (window Window) IsInputOnly() bool {
	ret0 := C.gdk_window_is_input_only(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsShaped is a wrapper around gdk_window_is_shaped().
func (window Window) IsShaped() bool {
	ret0 := C.gdk_window_is_shaped(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsViewable is a wrapper around gdk_window_is_viewable().
func (window Window) IsViewable() bool {
	ret0 := C.gdk_window_is_viewable(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsVisible is a wrapper around gdk_window_is_visible().
func (window Window) IsVisible() bool {
	ret0 := C.gdk_window_is_visible(window.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Lower is a wrapper around gdk_window_lower().
func (window Window) Lower() {
	C.gdk_window_lower(window.native())
}

// MarkPaintFromClip is a wrapper around gdk_window_mark_paint_from_clip().
func (window Window) MarkPaintFromClip(cr cairo.Context) {
	C.gdk_window_mark_paint_from_clip(window.native(), (*C.cairo_t)(cr.Ptr))
}

// Maximize is a wrapper around gdk_window_maximize().
func (window Window) Maximize() {
	C.gdk_window_maximize(window.native())
}

// MergeChildInputShapes is a wrapper around gdk_window_merge_child_input_shapes().
func (window Window) MergeChildInputShapes() {
	C.gdk_window_merge_child_input_shapes(window.native())
}

// MergeChildShapes is a wrapper around gdk_window_merge_child_shapes().
func (window Window) MergeChildShapes() {
	C.gdk_window_merge_child_shapes(window.native())
}

// Move is a wrapper around gdk_window_move().
func (window Window) Move(x int, y int) {
	C.gdk_window_move(window.native(), C.gint(x), C.gint(y))
}

// MoveRegion is a wrapper around gdk_window_move_region().
func (window Window) MoveRegion(region cairo.Region, dx int, dy int) {
	C.gdk_window_move_region(window.native(), (*C.cairo_region_t)(region.Ptr), C.gint(dx), C.gint(dy))
}

// MoveResize is a wrapper around gdk_window_move_resize().
func (window Window) MoveResize(x int, y int, width int, height int) {
	C.gdk_window_move_resize(window.native(), C.gint(x), C.gint(y), C.gint(width), C.gint(height))
}

// PeekChildren is a wrapper around gdk_window_peek_children().
func (window Window) PeekChildren() glib.List {
	ret0 := C.gdk_window_peek_children(window.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapWindow(p) }) /*gir:GLib*/
}

// Raise is a wrapper around gdk_window_raise().
func (window Window) Raise() {
	C.gdk_window_raise(window.native())
}

// RegisterDnd is a wrapper around gdk_window_register_dnd().
func (window Window) RegisterDnd() {
	C.gdk_window_register_dnd(window.native())
}

// Reparent is a wrapper around gdk_window_reparent().
func (window Window) Reparent(new_parent Window, x int, y int) {
	C.gdk_window_reparent(window.native(), new_parent.native(), C.gint(x), C.gint(y))
}

// Resize is a wrapper around gdk_window_resize().
func (window Window) Resize(width int, height int) {
	C.gdk_window_resize(window.native(), C.gint(width), C.gint(height))
}

// Restack is a wrapper around gdk_window_restack().
func (window Window) Restack(sibling Window, above bool) {
	C.gdk_window_restack(window.native(), sibling.native(), C.gboolean(util.Bool2Int(above)) /*go:.util*/)
}

// Scroll is a wrapper around gdk_window_scroll().
func (window Window) Scroll(dx int, dy int) {
	C.gdk_window_scroll(window.native(), C.gint(dx), C.gint(dy))
}

// SetAcceptFocus is a wrapper around gdk_window_set_accept_focus().
func (window Window) SetAcceptFocus(accept_focus bool) {
	C.gdk_window_set_accept_focus(window.native(), C.gboolean(util.Bool2Int(accept_focus)) /*go:.util*/)
}

// SetChildInputShapes is a wrapper around gdk_window_set_child_input_shapes().
func (window Window) SetChildInputShapes() {
	C.gdk_window_set_child_input_shapes(window.native())
}

// SetChildShapes is a wrapper around gdk_window_set_child_shapes().
func (window Window) SetChildShapes() {
	C.gdk_window_set_child_shapes(window.native())
}

// SetCursor is a wrapper around gdk_window_set_cursor().
func (window Window) SetCursor(cursor Cursor) {
	C.gdk_window_set_cursor(window.native(), cursor.native())
}

// SetDecorations is a wrapper around gdk_window_set_decorations().
func (window Window) SetDecorations(decorations WMDecoration) {
	C.gdk_window_set_decorations(window.native(), C.GdkWMDecoration(decorations))
}

// SetDeviceCursor is a wrapper around gdk_window_set_device_cursor().
func (window Window) SetDeviceCursor(device Device, cursor Cursor) {
	C.gdk_window_set_device_cursor(window.native(), device.native(), cursor.native())
}

// SetDeviceEvents is a wrapper around gdk_window_set_device_events().
func (window Window) SetDeviceEvents(device Device, event_mask EventMask) {
	C.gdk_window_set_device_events(window.native(), device.native(), C.GdkEventMask(event_mask))
}

// SetEventCompression is a wrapper around gdk_window_set_event_compression().
func (window Window) SetEventCompression(event_compression bool) {
	C.gdk_window_set_event_compression(window.native(), C.gboolean(util.Bool2Int(event_compression)) /*go:.util*/)
}

// SetEvents is a wrapper around gdk_window_set_events().
func (window Window) SetEvents(event_mask EventMask) {
	C.gdk_window_set_events(window.native(), C.GdkEventMask(event_mask))
}

// SetFocusOnMap is a wrapper around gdk_window_set_focus_on_map().
func (window Window) SetFocusOnMap(focus_on_map bool) {
	C.gdk_window_set_focus_on_map(window.native(), C.gboolean(util.Bool2Int(focus_on_map)) /*go:.util*/)
}

// SetFullscreenMode is a wrapper around gdk_window_set_fullscreen_mode().
func (window Window) SetFullscreenMode(mode FullscreenMode) {
	C.gdk_window_set_fullscreen_mode(window.native(), C.GdkFullscreenMode(mode))
}

// SetFunctions is a wrapper around gdk_window_set_functions().
func (window Window) SetFunctions(functions WMFunction) {
	C.gdk_window_set_functions(window.native(), C.GdkWMFunction(functions))
}

// SetGeometryHints is a wrapper around gdk_window_set_geometry_hints().
func (window Window) SetGeometryHints(geometry Geometry, geom_mask WindowHints) {
	C.gdk_window_set_geometry_hints(window.native(), geometry.native(), C.GdkWindowHints(geom_mask))
}

// SetGroup is a wrapper around gdk_window_set_group().
func (window Window) SetGroup(leader Window) {
	C.gdk_window_set_group(window.native(), leader.native())
}

// SetIconList is a wrapper around gdk_window_set_icon_list().
func (window Window) SetIconList(pixbufs glib.List) {
	C.gdk_window_set_icon_list(window.native(), (*C.GList)(pixbufs.Ptr))
}

// SetIconName is a wrapper around gdk_window_set_icon_name().
func (window Window) SetIconName(name string) {
	name0 := (*C.gchar)(C.CString(name))
	C.gdk_window_set_icon_name(window.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetKeepAbove is a wrapper around gdk_window_set_keep_above().
func (window Window) SetKeepAbove(setting bool) {
	C.gdk_window_set_keep_above(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetKeepBelow is a wrapper around gdk_window_set_keep_below().
func (window Window) SetKeepBelow(setting bool) {
	C.gdk_window_set_keep_below(window.native(), C.gboolean(util.Bool2Int(setting)) /*go:.util*/)
}

// SetModalHint is a wrapper around gdk_window_set_modal_hint().
func (window Window) SetModalHint(modal bool) {
	C.gdk_window_set_modal_hint(window.native(), C.gboolean(util.Bool2Int(modal)) /*go:.util*/)
}

// SetOpacity is a wrapper around gdk_window_set_opacity().
func (window Window) SetOpacity(opacity float64) {
	C.gdk_window_set_opacity(window.native(), C.gdouble(opacity))
}

// SetOpaqueRegion is a wrapper around gdk_window_set_opaque_region().
func (window Window) SetOpaqueRegion(region cairo.Region) {
	C.gdk_window_set_opaque_region(window.native(), (*C.cairo_region_t)(region.Ptr))
}

// SetOverrideRedirect is a wrapper around gdk_window_set_override_redirect().
func (window Window) SetOverrideRedirect(override_redirect bool) {
	C.gdk_window_set_override_redirect(window.native(), C.gboolean(util.Bool2Int(override_redirect)) /*go:.util*/)
}

// SetPassThrough is a wrapper around gdk_window_set_pass_through().
func (window Window) SetPassThrough(pass_through bool) {
	C.gdk_window_set_pass_through(window.native(), C.gboolean(util.Bool2Int(pass_through)) /*go:.util*/)
}

// SetRole is a wrapper around gdk_window_set_role().
func (window Window) SetRole(role string) {
	role0 := (*C.gchar)(C.CString(role))
	C.gdk_window_set_role(window.native(), role0)
	C.free(unsafe.Pointer(role0)) /*ch:<stdlib.h>*/
}

// SetShadowWidth is a wrapper around gdk_window_set_shadow_width().
func (window Window) SetShadowWidth(left int, right int, top int, bottom int) {
	C.gdk_window_set_shadow_width(window.native(), C.gint(left), C.gint(right), C.gint(top), C.gint(bottom))
}

// SetSkipPagerHint is a wrapper around gdk_window_set_skip_pager_hint().
func (window Window) SetSkipPagerHint(skips_pager bool) {
	C.gdk_window_set_skip_pager_hint(window.native(), C.gboolean(util.Bool2Int(skips_pager)) /*go:.util*/)
}

// SetSkipTaskbarHint is a wrapper around gdk_window_set_skip_taskbar_hint().
func (window Window) SetSkipTaskbarHint(skips_taskbar bool) {
	C.gdk_window_set_skip_taskbar_hint(window.native(), C.gboolean(util.Bool2Int(skips_taskbar)) /*go:.util*/)
}

// SetSourceEvents is a wrapper around gdk_window_set_source_events().
func (window Window) SetSourceEvents(source InputSource, event_mask EventMask) {
	C.gdk_window_set_source_events(window.native(), C.GdkInputSource(source), C.GdkEventMask(event_mask))
}

// SetStartupId is a wrapper around gdk_window_set_startup_id().
func (window Window) SetStartupId(startup_id string) {
	startup_id0 := (*C.gchar)(C.CString(startup_id))
	C.gdk_window_set_startup_id(window.native(), startup_id0)
	C.free(unsafe.Pointer(startup_id0)) /*ch:<stdlib.h>*/
}

// SetSupportMultidevice is a wrapper around gdk_window_set_support_multidevice().
func (window Window) SetSupportMultidevice(support_multidevice bool) {
	C.gdk_window_set_support_multidevice(window.native(), C.gboolean(util.Bool2Int(support_multidevice)) /*go:.util*/)
}

// SetTitle is a wrapper around gdk_window_set_title().
func (window Window) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.gdk_window_set_title(window.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetTransientFor is a wrapper around gdk_window_set_transient_for().
func (window Window) SetTransientFor(parent Window) {
	C.gdk_window_set_transient_for(window.native(), parent.native())
}

// SetTypeHint is a wrapper around gdk_window_set_type_hint().
func (window Window) SetTypeHint(hint WindowTypeHint) {
	C.gdk_window_set_type_hint(window.native(), C.GdkWindowTypeHint(hint))
}

// SetUrgencyHint is a wrapper around gdk_window_set_urgency_hint().
func (window Window) SetUrgencyHint(urgent bool) {
	C.gdk_window_set_urgency_hint(window.native(), C.gboolean(util.Bool2Int(urgent)) /*go:.util*/)
}

// SetUserData is a wrapper around gdk_window_set_user_data().
func (window Window) SetUserData(user_data gobject.Object) {
	C.gdk_window_set_user_data(window.native(), C.gpointer((C.gpointer)(user_data.Ptr)))
}

// ShapeCombineRegion is a wrapper around gdk_window_shape_combine_region().
func (window Window) ShapeCombineRegion(shape_region cairo.Region, offset_x int, offset_y int) {
	C.gdk_window_shape_combine_region(window.native(), (*C.cairo_region_t)(shape_region.Ptr), C.gint(offset_x), C.gint(offset_y))
}

// Show is a wrapper around gdk_window_show().
func (window Window) Show() {
	C.gdk_window_show(window.native())
}

// ShowUnraised is a wrapper around gdk_window_show_unraised().
func (window Window) ShowUnraised() {
	C.gdk_window_show_unraised(window.native())
}

// Stick is a wrapper around gdk_window_stick().
func (window Window) Stick() {
	C.gdk_window_stick(window.native())
}

// ThawUpdates is a wrapper around gdk_window_thaw_updates().
func (window Window) ThawUpdates() {
	C.gdk_window_thaw_updates(window.native())
}

// Unfullscreen is a wrapper around gdk_window_unfullscreen().
func (window Window) Unfullscreen() {
	C.gdk_window_unfullscreen(window.native())
}

// Unmaximize is a wrapper around gdk_window_unmaximize().
func (window Window) Unmaximize() {
	C.gdk_window_unmaximize(window.native())
}

// Unstick is a wrapper around gdk_window_unstick().
func (window Window) Unstick() {
	C.gdk_window_unstick(window.native())
}

// Withdraw is a wrapper around gdk_window_withdraw().
func (window Window) Withdraw() {
	C.gdk_window_withdraw(window.native())
}

// WindowConstrainSize is a wrapper around gdk_window_constrain_size().
func WindowConstrainSize(geometry Geometry, flags WindowHints, width int, height int) (int, int) {
	var new_width0 C.gint
	var new_height0 C.gint
	C.gdk_window_constrain_size(geometry.native(), C.GdkWindowHints(flags), C.gint(width), C.gint(height), &new_width0, &new_height0)
	return int(new_width0), int(new_height0)
}

// Object GLContext
type GLContext struct {
	gobject.Object
}

func (v GLContext) native() *C.GdkGLContext {
	return (*C.GdkGLContext)(v.Ptr)
}
func wrapGLContext(p *C.GdkGLContext) (v GLContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapGLContext(p unsafe.Pointer) (v GLContext) {
	v.Ptr = p
	return
}
func (v GLContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapGLContext(p unsafe.Pointer) interface{} {
	return WrapGLContext(p)
}
func (v GLContext) GetType() gobject.Type {
	return gobject.Type(C.gdk_gl_context_get_type())
}
func (v GLContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapGLContext(unsafe.Pointer(ptr)), nil
	}
}

// GetDebugEnabled is a wrapper around gdk_gl_context_get_debug_enabled().
func (context GLContext) GetDebugEnabled() bool {
	ret0 := C.gdk_gl_context_get_debug_enabled(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDisplay is a wrapper around gdk_gl_context_get_display().
func (context GLContext) GetDisplay() Display {
	ret0 := C.gdk_gl_context_get_display(context.native())
	return wrapDisplay(ret0)
}

// GetForwardCompatible is a wrapper around gdk_gl_context_get_forward_compatible().
func (context GLContext) GetForwardCompatible() bool {
	ret0 := C.gdk_gl_context_get_forward_compatible(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetRequiredVersion is a wrapper around gdk_gl_context_get_required_version().
func (context GLContext) GetRequiredVersion() (int, int) {
	var major0 C.int
	var minor0 C.int
	C.gdk_gl_context_get_required_version(context.native(), &major0, &minor0)
	return int(major0), int(minor0)
}

// GetSharedContext is a wrapper around gdk_gl_context_get_shared_context().
func (context GLContext) GetSharedContext() GLContext {
	ret0 := C.gdk_gl_context_get_shared_context(context.native())
	return wrapGLContext(ret0)
}

// GetUseEs is a wrapper around gdk_gl_context_get_use_es().
func (context GLContext) GetUseEs() bool {
	ret0 := C.gdk_gl_context_get_use_es(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetVersion is a wrapper around gdk_gl_context_get_version().
func (context GLContext) GetVersion() (int, int) {
	var major0 C.int
	var minor0 C.int
	C.gdk_gl_context_get_version(context.native(), &major0, &minor0)
	return int(major0), int(minor0)
}

// GetWindow is a wrapper around gdk_gl_context_get_window().
func (context GLContext) GetWindow() Window {
	ret0 := C.gdk_gl_context_get_window(context.native())
	return wrapWindow(ret0)
}

// IsLegacy is a wrapper around gdk_gl_context_is_legacy().
func (context GLContext) IsLegacy() bool {
	ret0 := C.gdk_gl_context_is_legacy(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MakeCurrent is a wrapper around gdk_gl_context_make_current().
func (context GLContext) MakeCurrent() {
	C.gdk_gl_context_make_current(context.native())
}

// Realize is a wrapper around gdk_gl_context_realize().
func (context GLContext) Realize() (bool, error) {
	var err glib.Error
	ret0 := C.gdk_gl_context_realize(context.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetDebugEnabled is a wrapper around gdk_gl_context_set_debug_enabled().
func (context GLContext) SetDebugEnabled(enabled bool) {
	C.gdk_gl_context_set_debug_enabled(context.native(), C.gboolean(util.Bool2Int(enabled)) /*go:.util*/)
}

// SetForwardCompatible is a wrapper around gdk_gl_context_set_forward_compatible().
func (context GLContext) SetForwardCompatible(compatible bool) {
	C.gdk_gl_context_set_forward_compatible(context.native(), C.gboolean(util.Bool2Int(compatible)) /*go:.util*/)
}

// SetRequiredVersion is a wrapper around gdk_gl_context_set_required_version().
func (context GLContext) SetRequiredVersion(major int, minor int) {
	C.gdk_gl_context_set_required_version(context.native(), C.int(major), C.int(minor))
}

// SetUseEs is a wrapper around gdk_gl_context_set_use_es().
func (context GLContext) SetUseEs(use_es int) {
	C.gdk_gl_context_set_use_es(context.native(), C.int(use_es))
}

// GLContextClearCurrent is a wrapper around gdk_gl_context_clear_current().
func GLContextClearCurrent() {
	C.gdk_gl_context_clear_current()
}

// GLContextGetCurrent is a wrapper around gdk_gl_context_get_current().
func GLContextGetCurrent() GLContext {
	ret0 := C.gdk_gl_context_get_current()
	return wrapGLContext(ret0)
}

// Object DrawingContext
type DrawingContext struct {
	gobject.Object
}

func (v DrawingContext) native() *C.GdkDrawingContext {
	return (*C.GdkDrawingContext)(v.Ptr)
}
func wrapDrawingContext(p *C.GdkDrawingContext) (v DrawingContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDrawingContext(p unsafe.Pointer) (v DrawingContext) {
	v.Ptr = p
	return
}
func (v DrawingContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDrawingContext(p unsafe.Pointer) interface{} {
	return WrapDrawingContext(p)
}
func (v DrawingContext) GetType() gobject.Type {
	return gobject.Type(C.gdk_drawing_context_get_type())
}
func (v DrawingContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDrawingContext(unsafe.Pointer(ptr)), nil
	}
}

// GetCairoContext is a wrapper around gdk_drawing_context_get_cairo_context().
func (context DrawingContext) GetCairoContext() cairo.Context {
	ret0 := C.gdk_drawing_context_get_cairo_context(context.native())
	return cairo.WrapContext(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetClip is a wrapper around gdk_drawing_context_get_clip().
func (context DrawingContext) GetClip() cairo.Region {
	ret0 := C.gdk_drawing_context_get_clip(context.native())
	return cairo.WrapRegion(unsafe.Pointer(ret0)) /*gir:cairo*/
}

// GetWindow is a wrapper around gdk_drawing_context_get_window().
func (context DrawingContext) GetWindow() Window {
	ret0 := C.gdk_drawing_context_get_window(context.native())
	return wrapWindow(ret0)
}

// IsValid is a wrapper around gdk_drawing_context_is_valid().
func (context DrawingContext) IsValid() bool {
	ret0 := C.gdk_drawing_context_is_valid(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object Cursor
type Cursor struct {
	gobject.Object
}

func (v Cursor) native() *C.GdkCursor {
	return (*C.GdkCursor)(v.Ptr)
}
func wrapCursor(p *C.GdkCursor) (v Cursor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCursor(p unsafe.Pointer) (v Cursor) {
	v.Ptr = p
	return
}
func (v Cursor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCursor(p unsafe.Pointer) interface{} {
	return WrapCursor(p)
}
func (v Cursor) GetType() gobject.Type {
	return gobject.Type(C.gdk_cursor_get_type())
}
func (v Cursor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCursor(unsafe.Pointer(ptr)), nil
	}
}

// CursorNewForDisplay is a wrapper around gdk_cursor_new_for_display().
func CursorNewForDisplay(display Display, cursor_type CursorType) Cursor {
	ret0 := C.gdk_cursor_new_for_display(display.native(), C.GdkCursorType(cursor_type))
	return wrapCursor(ret0)
}

// CursorNewFromName is a wrapper around gdk_cursor_new_from_name().
func CursorNewFromName(display Display, name string) Cursor {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gdk_cursor_new_from_name(display.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapCursor(ret0)
}

// CursorNewFromSurface is a wrapper around gdk_cursor_new_from_surface().
func CursorNewFromSurface(display Display, surface cairo.Surface, x float64, y float64) Cursor {
	ret0 := C.gdk_cursor_new_from_surface(display.native(), (*C.cairo_surface_t)(surface.Ptr), C.gdouble(x), C.gdouble(y))
	return wrapCursor(ret0)
}

// GetCursorType is a wrapper around gdk_cursor_get_cursor_type().
func (cursor Cursor) GetCursorType() CursorType {
	ret0 := C.gdk_cursor_get_cursor_type(cursor.native())
	return CursorType(ret0)
}

// GetDisplay is a wrapper around gdk_cursor_get_display().
func (cursor Cursor) GetDisplay() Display {
	ret0 := C.gdk_cursor_get_display(cursor.native())
	return wrapDisplay(ret0)
}

// GetImage is a wrapper around gdk_cursor_get_image().
func (cursor Cursor) GetImage() gdkpixbuf.Pixbuf {
	ret0 := C.gdk_cursor_get_image(cursor.native())
	return gdkpixbuf.WrapPixbuf(unsafe.Pointer(ret0)) /*gir:GdkPixbuf*/
}

// GetSurface is a wrapper around gdk_cursor_get_surface().
func (cursor Cursor) GetSurface() (cairo.Surface, float64, float64) {
	var x_hot0 C.gdouble
	var y_hot0 C.gdouble
	ret0 := C.gdk_cursor_get_surface(cursor.native(), &x_hot0, &y_hot0)
	return cairo.WrapSurface(unsafe.Pointer(ret0)) /*gir:cairo*/, float64(x_hot0), float64(y_hot0)
}

// Object Seat
type Seat struct {
	gobject.Object
}

func (v Seat) native() *C.GdkSeat {
	return (*C.GdkSeat)(v.Ptr)
}
func wrapSeat(p *C.GdkSeat) (v Seat) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSeat(p unsafe.Pointer) (v Seat) {
	v.Ptr = p
	return
}
func (v Seat) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSeat(p unsafe.Pointer) interface{} {
	return WrapSeat(p)
}
func (v Seat) GetType() gobject.Type {
	return gobject.Type(C.gdk_seat_get_type())
}
func (v Seat) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSeat(unsafe.Pointer(ptr)), nil
	}
}

// GetCapabilities is a wrapper around gdk_seat_get_capabilities().
func (seat Seat) GetCapabilities() SeatCapabilities {
	ret0 := C.gdk_seat_get_capabilities(seat.native())
	return SeatCapabilities(ret0)
}

// GetDisplay is a wrapper around gdk_seat_get_display().
func (seat Seat) GetDisplay() Display {
	ret0 := C.gdk_seat_get_display(seat.native())
	return wrapDisplay(ret0)
}

// GetKeyboard is a wrapper around gdk_seat_get_keyboard().
func (seat Seat) GetKeyboard() Device {
	ret0 := C.gdk_seat_get_keyboard(seat.native())
	return wrapDevice(ret0)
}

// GetPointer is a wrapper around gdk_seat_get_pointer().
func (seat Seat) GetPointer() Device {
	ret0 := C.gdk_seat_get_pointer(seat.native())
	return wrapDevice(ret0)
}

// GetSlaves is a wrapper around gdk_seat_get_slaves().
func (seat Seat) GetSlaves(capabilities SeatCapabilities) glib.List {
	ret0 := C.gdk_seat_get_slaves(seat.native(), C.GdkSeatCapabilities(capabilities))
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDevice(p) }) /*gir:GLib*/
}

// Ungrab is a wrapper around gdk_seat_ungrab().
func (seat Seat) Ungrab() {
	C.gdk_seat_ungrab(seat.native())
}

// Struct Atom
type Atom struct {
	Ptr unsafe.Pointer
}

func (v Atom) native() *C.GdkAtom {
	return (*C.GdkAtom)(v.Ptr)
}
func wrapAtom(p *C.GdkAtom) Atom {
	return Atom{Ptr: unsafe.Pointer(p)}
}
func WrapAtom(p unsafe.Pointer) Atom {
	return Atom{Ptr: p}
}
func (v Atom) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAtom(p unsafe.Pointer) interface{} {
	return WrapAtom(p)
}

// AtomIntern is a wrapper around gdk_atom_intern().
func AtomIntern(atom_name string, only_if_exists bool) Atom {
	atom_name0 := (*C.gchar)(C.CString(atom_name))
	ret0 := C.gdk_atom_intern(atom_name0, C.gboolean(util.Bool2Int(only_if_exists)) /*go:.util*/)
	C.free(unsafe.Pointer(atom_name0)) /*ch:<stdlib.h>*/
	return wrapAtom(&ret0)
}

// AtomInternStaticString is a wrapper around gdk_atom_intern_static_string().
func AtomInternStaticString(atom_name string) Atom {
	atom_name0 := (*C.gchar)(C.CString(atom_name))
	ret0 := C.gdk_atom_intern_static_string(atom_name0)
	C.free(unsafe.Pointer(atom_name0)) /*ch:<stdlib.h>*/
	return wrapAtom(&ret0)
}

// Object DeviceManager
type DeviceManager struct {
	gobject.Object
}

func (v DeviceManager) native() *C.GdkDeviceManager {
	return (*C.GdkDeviceManager)(v.Ptr)
}
func wrapDeviceManager(p *C.GdkDeviceManager) (v DeviceManager) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDeviceManager(p unsafe.Pointer) (v DeviceManager) {
	v.Ptr = p
	return
}
func (v DeviceManager) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDeviceManager(p unsafe.Pointer) interface{} {
	return WrapDeviceManager(p)
}
func (v DeviceManager) GetType() gobject.Type {
	return gobject.Type(C.gdk_device_manager_get_type())
}
func (v DeviceManager) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDeviceManager(unsafe.Pointer(ptr)), nil
	}
}

// GetDisplay is a wrapper around gdk_device_manager_get_display().
func (device_manager DeviceManager) GetDisplay() Display {
	ret0 := C.gdk_device_manager_get_display(device_manager.native())
	return wrapDisplay(ret0)
}

// Object DeviceTool
type DeviceTool struct {
	gobject.Object
}

func (v DeviceTool) native() *C.GdkDeviceTool {
	return (*C.GdkDeviceTool)(v.Ptr)
}
func wrapDeviceTool(p *C.GdkDeviceTool) (v DeviceTool) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDeviceTool(p unsafe.Pointer) (v DeviceTool) {
	v.Ptr = p
	return
}
func (v DeviceTool) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDeviceTool(p unsafe.Pointer) interface{} {
	return WrapDeviceTool(p)
}
func (v DeviceTool) GetType() gobject.Type {
	return gobject.Type(C.gdk_device_tool_get_type())
}
func (v DeviceTool) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDeviceTool(unsafe.Pointer(ptr)), nil
	}
}

// GetHardwareId is a wrapper around gdk_device_tool_get_hardware_id().
func (tool DeviceTool) GetHardwareId() uint64 {
	ret0 := C.gdk_device_tool_get_hardware_id(tool.native())
	return uint64(ret0)
}

// GetSerial is a wrapper around gdk_device_tool_get_serial().
func (tool DeviceTool) GetSerial() uint64 {
	ret0 := C.gdk_device_tool_get_serial(tool.native())
	return uint64(ret0)
}

// GetToolType is a wrapper around gdk_device_tool_get_tool_type().
func (tool DeviceTool) GetToolType() DeviceToolType {
	ret0 := C.gdk_device_tool_get_tool_type(tool.native())
	return DeviceToolType(ret0)
}

// Object Monitor
type Monitor struct {
	gobject.Object
}

func (v Monitor) native() *C.GdkMonitor {
	return (*C.GdkMonitor)(v.Ptr)
}
func wrapMonitor(p *C.GdkMonitor) (v Monitor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMonitor(p unsafe.Pointer) (v Monitor) {
	v.Ptr = p
	return
}
func (v Monitor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMonitor(p unsafe.Pointer) interface{} {
	return WrapMonitor(p)
}
func (v Monitor) GetType() gobject.Type {
	return gobject.Type(C.gdk_monitor_get_type())
}
func (v Monitor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMonitor(unsafe.Pointer(ptr)), nil
	}
}

// GetDisplay is a wrapper around gdk_monitor_get_display().
func (monitor Monitor) GetDisplay() Display {
	ret0 := C.gdk_monitor_get_display(monitor.native())
	return wrapDisplay(ret0)
}

// GetGeometry is a wrapper around gdk_monitor_get_geometry().
func (monitor Monitor) GetGeometry() Rectangle {
	var geometry0 C.GdkRectangle
	C.gdk_monitor_get_geometry(monitor.native(), &geometry0)
	return wrapRectangle(&geometry0)
}

// GetHeightMm is a wrapper around gdk_monitor_get_height_mm().
func (monitor Monitor) GetHeightMm() int {
	ret0 := C.gdk_monitor_get_height_mm(monitor.native())
	return int(ret0)
}

// GetManufacturer is a wrapper around gdk_monitor_get_manufacturer().
func (monitor Monitor) GetManufacturer() string {
	ret0 := C.gdk_monitor_get_manufacturer(monitor.native())
	ret := C.GoString(ret0)
	return ret
}

// GetModel is a wrapper around gdk_monitor_get_model().
func (monitor Monitor) GetModel() string {
	ret0 := C.gdk_monitor_get_model(monitor.native())
	ret := C.GoString(ret0)
	return ret
}

// GetRefreshRate is a wrapper around gdk_monitor_get_refresh_rate().
func (monitor Monitor) GetRefreshRate() int {
	ret0 := C.gdk_monitor_get_refresh_rate(monitor.native())
	return int(ret0)
}

// GetScaleFactor is a wrapper around gdk_monitor_get_scale_factor().
func (monitor Monitor) GetScaleFactor() int {
	ret0 := C.gdk_monitor_get_scale_factor(monitor.native())
	return int(ret0)
}

// GetSubpixelLayout is a wrapper around gdk_monitor_get_subpixel_layout().
func (monitor Monitor) GetSubpixelLayout() SubpixelLayout {
	ret0 := C.gdk_monitor_get_subpixel_layout(monitor.native())
	return SubpixelLayout(ret0)
}

// GetWidthMm is a wrapper around gdk_monitor_get_width_mm().
func (monitor Monitor) GetWidthMm() int {
	ret0 := C.gdk_monitor_get_width_mm(monitor.native())
	return int(ret0)
}

// GetWorkarea is a wrapper around gdk_monitor_get_workarea().
func (monitor Monitor) GetWorkarea() Rectangle {
	var workarea0 C.GdkRectangle
	C.gdk_monitor_get_workarea(monitor.native(), &workarea0)
	return wrapRectangle(&workarea0)
}

// IsPrimary is a wrapper around gdk_monitor_is_primary().
func (monitor Monitor) IsPrimary() bool {
	ret0 := C.gdk_monitor_is_primary(monitor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object DisplayManager
type DisplayManager struct {
	gobject.Object
}

func (v DisplayManager) native() *C.GdkDisplayManager {
	return (*C.GdkDisplayManager)(v.Ptr)
}
func wrapDisplayManager(p *C.GdkDisplayManager) (v DisplayManager) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDisplayManager(p unsafe.Pointer) (v DisplayManager) {
	v.Ptr = p
	return
}
func (v DisplayManager) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDisplayManager(p unsafe.Pointer) interface{} {
	return WrapDisplayManager(p)
}
func (v DisplayManager) GetType() gobject.Type {
	return gobject.Type(C.gdk_display_manager_get_type())
}
func (v DisplayManager) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDisplayManager(unsafe.Pointer(ptr)), nil
	}
}

// GetDefaultDisplay is a wrapper around gdk_display_manager_get_default_display().
func (manager DisplayManager) GetDefaultDisplay() Display {
	ret0 := C.gdk_display_manager_get_default_display(manager.native())
	return wrapDisplay(ret0)
}

// OpenDisplay is a wrapper around gdk_display_manager_open_display().
func (manager DisplayManager) OpenDisplay(name string) Display {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.gdk_display_manager_open_display(manager.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapDisplay(ret0)
}

// SetDefaultDisplay is a wrapper around gdk_display_manager_set_default_display().
func (manager DisplayManager) SetDefaultDisplay(display Display) {
	C.gdk_display_manager_set_default_display(manager.native(), display.native())
}

// DisplayManagerGet is a wrapper around gdk_display_manager_get().
func DisplayManagerGet() DisplayManager {
	ret0 := C.gdk_display_manager_get()
	return wrapDisplayManager(ret0)
}

// Object DragContext
type DragContext struct {
	gobject.Object
}

func (v DragContext) native() *C.GdkDragContext {
	return (*C.GdkDragContext)(v.Ptr)
}
func wrapDragContext(p *C.GdkDragContext) (v DragContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDragContext(p unsafe.Pointer) (v DragContext) {
	v.Ptr = p
	return
}
func (v DragContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDragContext(p unsafe.Pointer) interface{} {
	return WrapDragContext(p)
}
func (v DragContext) GetType() gobject.Type {
	return gobject.Type(C.gdk_drag_context_get_type())
}
func (v DragContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDragContext(unsafe.Pointer(ptr)), nil
	}
}

// GetActions is a wrapper around gdk_drag_context_get_actions().
func (context DragContext) GetActions() DragAction {
	ret0 := C.gdk_drag_context_get_actions(context.native())
	return DragAction(ret0)
}

// GetDestWindow is a wrapper around gdk_drag_context_get_dest_window().
func (context DragContext) GetDestWindow() Window {
	ret0 := C.gdk_drag_context_get_dest_window(context.native())
	return wrapWindow(ret0)
}

// GetDevice is a wrapper around gdk_drag_context_get_device().
func (context DragContext) GetDevice() Device {
	ret0 := C.gdk_drag_context_get_device(context.native())
	return wrapDevice(ret0)
}

// GetDragWindow is a wrapper around gdk_drag_context_get_drag_window().
func (context DragContext) GetDragWindow() Window {
	ret0 := C.gdk_drag_context_get_drag_window(context.native())
	return wrapWindow(ret0)
}

// GetProtocol is a wrapper around gdk_drag_context_get_protocol().
func (context DragContext) GetProtocol() DragProtocol {
	ret0 := C.gdk_drag_context_get_protocol(context.native())
	return DragProtocol(ret0)
}

// GetSelectedAction is a wrapper around gdk_drag_context_get_selected_action().
func (context DragContext) GetSelectedAction() DragAction {
	ret0 := C.gdk_drag_context_get_selected_action(context.native())
	return DragAction(ret0)
}

// GetSourceWindow is a wrapper around gdk_drag_context_get_source_window().
func (context DragContext) GetSourceWindow() Window {
	ret0 := C.gdk_drag_context_get_source_window(context.native())
	return wrapWindow(ret0)
}

// GetSuggestedAction is a wrapper around gdk_drag_context_get_suggested_action().
func (context DragContext) GetSuggestedAction() DragAction {
	ret0 := C.gdk_drag_context_get_suggested_action(context.native())
	return DragAction(ret0)
}

// ListTargets is a wrapper around gdk_drag_context_list_targets().
func (context DragContext) ListTargets() glib.List {
	ret0 := C.gdk_drag_context_list_targets(context.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapAtom(p) }) /*gir:GLib*/
}

// ManageDnd is a wrapper around gdk_drag_context_manage_dnd().
func (context DragContext) ManageDnd(ipc_window Window, actions DragAction) bool {
	ret0 := C.gdk_drag_context_manage_dnd(context.native(), ipc_window.native(), C.GdkDragAction(actions))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetDevice is a wrapper around gdk_drag_context_set_device().
func (context DragContext) SetDevice(device Device) {
	C.gdk_drag_context_set_device(context.native(), device.native())
}

// SetHotspot is a wrapper around gdk_drag_context_set_hotspot().
func (context DragContext) SetHotspot(hot_x int, hot_y int) {
	C.gdk_drag_context_set_hotspot(context.native(), C.gint(hot_x), C.gint(hot_y))
}

// Object FrameClock
type FrameClock struct {
	gobject.Object
}

func (v FrameClock) native() *C.GdkFrameClock {
	return (*C.GdkFrameClock)(v.Ptr)
}
func wrapFrameClock(p *C.GdkFrameClock) (v FrameClock) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFrameClock(p unsafe.Pointer) (v FrameClock) {
	v.Ptr = p
	return
}
func (v FrameClock) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFrameClock(p unsafe.Pointer) interface{} {
	return WrapFrameClock(p)
}
func (v FrameClock) GetType() gobject.Type {
	return gobject.Type(C.gdk_frame_clock_get_type())
}
func (v FrameClock) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFrameClock(unsafe.Pointer(ptr)), nil
	}
}

// BeginUpdating is a wrapper around gdk_frame_clock_begin_updating().
func (frame_clock FrameClock) BeginUpdating() {
	C.gdk_frame_clock_begin_updating(frame_clock.native())
}

// EndUpdating is a wrapper around gdk_frame_clock_end_updating().
func (frame_clock FrameClock) EndUpdating() {
	C.gdk_frame_clock_end_updating(frame_clock.native())
}

// GetCurrentTimings is a wrapper around gdk_frame_clock_get_current_timings().
func (frame_clock FrameClock) GetCurrentTimings() FrameTimings {
	ret0 := C.gdk_frame_clock_get_current_timings(frame_clock.native())
	return wrapFrameTimings(ret0)
}

// GetFrameCounter is a wrapper around gdk_frame_clock_get_frame_counter().
func (frame_clock FrameClock) GetFrameCounter() int64 {
	ret0 := C.gdk_frame_clock_get_frame_counter(frame_clock.native())
	return int64(ret0)
}

// GetFrameTime is a wrapper around gdk_frame_clock_get_frame_time().
func (frame_clock FrameClock) GetFrameTime() int64 {
	ret0 := C.gdk_frame_clock_get_frame_time(frame_clock.native())
	return int64(ret0)
}

// GetHistoryStart is a wrapper around gdk_frame_clock_get_history_start().
func (frame_clock FrameClock) GetHistoryStart() int64 {
	ret0 := C.gdk_frame_clock_get_history_start(frame_clock.native())
	return int64(ret0)
}

// GetTimings is a wrapper around gdk_frame_clock_get_timings().
func (frame_clock FrameClock) GetTimings(frame_counter int64) FrameTimings {
	ret0 := C.gdk_frame_clock_get_timings(frame_clock.native(), C.gint64(frame_counter))
	return wrapFrameTimings(ret0)
}

// RequestPhase is a wrapper around gdk_frame_clock_request_phase().
func (frame_clock FrameClock) RequestPhase(phase FrameClockPhase) {
	C.gdk_frame_clock_request_phase(frame_clock.native(), C.GdkFrameClockPhase(phase))
}

// Object Keymap
type Keymap struct {
	gobject.Object
}

func (v Keymap) native() *C.GdkKeymap {
	return (*C.GdkKeymap)(v.Ptr)
}
func wrapKeymap(p *C.GdkKeymap) (v Keymap) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapKeymap(p unsafe.Pointer) (v Keymap) {
	v.Ptr = p
	return
}
func (v Keymap) IsNil() bool {
	return v.Ptr == nil
}
func IWrapKeymap(p unsafe.Pointer) interface{} {
	return WrapKeymap(p)
}
func (v Keymap) GetType() gobject.Type {
	return gobject.Type(C.gdk_keymap_get_type())
}
func (v Keymap) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapKeymap(unsafe.Pointer(ptr)), nil
	}
}

// GetCapsLockState is a wrapper around gdk_keymap_get_caps_lock_state().
func (keymap Keymap) GetCapsLockState() bool {
	ret0 := C.gdk_keymap_get_caps_lock_state(keymap.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetDirection is a wrapper around gdk_keymap_get_direction().
func (keymap Keymap) GetDirection() /*gir:Pango*/ pango.Direction {
	ret0 := C.gdk_keymap_get_direction(keymap.native())
	return /*gir:Pango*/ pango.Direction(ret0)
}

// GetModifierMask is a wrapper around gdk_keymap_get_modifier_mask().
func (keymap Keymap) GetModifierMask(intent ModifierIntent) ModifierType {
	ret0 := C.gdk_keymap_get_modifier_mask(keymap.native(), C.GdkModifierIntent(intent))
	return ModifierType(ret0)
}

// GetModifierState is a wrapper around gdk_keymap_get_modifier_state().
func (keymap Keymap) GetModifierState() uint {
	ret0 := C.gdk_keymap_get_modifier_state(keymap.native())
	return uint(ret0)
}

// GetNumLockState is a wrapper around gdk_keymap_get_num_lock_state().
func (keymap Keymap) GetNumLockState() bool {
	ret0 := C.gdk_keymap_get_num_lock_state(keymap.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetScrollLockState is a wrapper around gdk_keymap_get_scroll_lock_state().
func (keymap Keymap) GetScrollLockState() bool {
	ret0 := C.gdk_keymap_get_scroll_lock_state(keymap.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HaveBidiLayouts is a wrapper around gdk_keymap_have_bidi_layouts().
func (keymap Keymap) HaveBidiLayouts() bool {
	ret0 := C.gdk_keymap_have_bidi_layouts(keymap.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LookupKey is a wrapper around gdk_keymap_lookup_key().
func (keymap Keymap) LookupKey(key KeymapKey) uint {
	ret0 := C.gdk_keymap_lookup_key(keymap.native(), key.native())
	return uint(ret0)
}

// TranslateKeyboardState is a wrapper around gdk_keymap_translate_keyboard_state().
func (keymap Keymap) TranslateKeyboardState(hardware_keycode uint, state ModifierType, group int) (bool, uint, int, int, ModifierType) {
	var keyval0 C.guint
	var effective_group0 C.gint
	var level0 C.gint
	var consumed_modifiers0 C.GdkModifierType
	ret0 := C.gdk_keymap_translate_keyboard_state(keymap.native(), C.guint(hardware_keycode), C.GdkModifierType(state), C.gint(group), &keyval0, &effective_group0, &level0, &consumed_modifiers0)
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint(keyval0), int(effective_group0), int(level0), ModifierType(consumed_modifiers0)
}

// KeymapGetDefault is a wrapper around gdk_keymap_get_default().
func KeymapGetDefault() Keymap {
	ret0 := C.gdk_keymap_get_default()
	return wrapKeymap(ret0)
}

// KeymapGetForDisplay is a wrapper around gdk_keymap_get_for_display().
func KeymapGetForDisplay(display Display) Keymap {
	ret0 := C.gdk_keymap_get_for_display(display.native())
	return wrapKeymap(ret0)
}

// Object Visual
type Visual struct {
	gobject.Object
}

func (v Visual) native() *C.GdkVisual {
	return (*C.GdkVisual)(v.Ptr)
}
func wrapVisual(p *C.GdkVisual) (v Visual) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVisual(p unsafe.Pointer) (v Visual) {
	v.Ptr = p
	return
}
func (v Visual) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVisual(p unsafe.Pointer) interface{} {
	return WrapVisual(p)
}
func (v Visual) GetType() gobject.Type {
	return gobject.Type(C.gdk_visual_get_type())
}
func (v Visual) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVisual(unsafe.Pointer(ptr)), nil
	}
}

// GetBluePixelDetails is a wrapper around gdk_visual_get_blue_pixel_details().
func (visual Visual) GetBluePixelDetails() (uint32, int, int) {
	var mask0 C.guint32
	var shift0 C.gint
	var precision0 C.gint
	C.gdk_visual_get_blue_pixel_details(visual.native(), &mask0, &shift0, &precision0)
	return uint32(mask0), int(shift0), int(precision0)
}

// GetDepth is a wrapper around gdk_visual_get_depth().
func (visual Visual) GetDepth() int {
	ret0 := C.gdk_visual_get_depth(visual.native())
	return int(ret0)
}

// GetGreenPixelDetails is a wrapper around gdk_visual_get_green_pixel_details().
func (visual Visual) GetGreenPixelDetails() (uint32, int, int) {
	var mask0 C.guint32
	var shift0 C.gint
	var precision0 C.gint
	C.gdk_visual_get_green_pixel_details(visual.native(), &mask0, &shift0, &precision0)
	return uint32(mask0), int(shift0), int(precision0)
}

// GetRedPixelDetails is a wrapper around gdk_visual_get_red_pixel_details().
func (visual Visual) GetRedPixelDetails() (uint32, int, int) {
	var mask0 C.guint32
	var shift0 C.gint
	var precision0 C.gint
	C.gdk_visual_get_red_pixel_details(visual.native(), &mask0, &shift0, &precision0)
	return uint32(mask0), int(shift0), int(precision0)
}

// GetScreen is a wrapper around gdk_visual_get_screen().
func (visual Visual) GetScreen() Screen {
	ret0 := C.gdk_visual_get_screen(visual.native())
	return wrapScreen(ret0)
}

// GetVisualType is a wrapper around gdk_visual_get_visual_type().
func (visual Visual) GetVisualType() VisualType {
	ret0 := C.gdk_visual_get_visual_type(visual.native())
	return VisualType(ret0)
}

type XEvent unsafe.Pointer
type AxisUse int

const (
	AxisUseIgnore   AxisUse = 0
	AxisUseX                = 1
	AxisUseY                = 2
	AxisUsePressure         = 3
	AxisUseXtilt            = 4
	AxisUseYtilt            = 5
	AxisUseWheel            = 6
	AxisUseDistance         = 7
	AxisUseRotation         = 8
	AxisUseSlider           = 9
	AxisUseLast             = 10
)

type ByteOrder int

const (
	ByteOrderLsbFirst ByteOrder = 0
	ByteOrderMsbFirst           = 1
)

type CrossingMode int

const (
	CrossingModeNormal       CrossingMode = 0
	CrossingModeGrab                      = 1
	CrossingModeUngrab                    = 2
	CrossingModeGtkGrab                   = 3
	CrossingModeGtkUngrab                 = 4
	CrossingModeStateChanged              = 5
	CrossingModeTouchBegin                = 6
	CrossingModeTouchEnd                  = 7
	CrossingModeDeviceSwitch              = 8
)

type CursorType int

const (
	CursorTypeXCursor           CursorType = 0
	CursorTypeArrow                        = 2
	CursorTypeBasedArrowDown               = 4
	CursorTypeBasedArrowUp                 = 6
	CursorTypeBoat                         = 8
	CursorTypeBogosity                     = 10
	CursorTypeBottomLeftCorner             = 12
	CursorTypeBottomRightCorner            = 14
	CursorTypeBottomSide                   = 16
	CursorTypeBottomTee                    = 18
	CursorTypeBoxSpiral                    = 20
	CursorTypeCenterPtr                    = 22
	CursorTypeCircle                       = 24
	CursorTypeClock                        = 26
	CursorTypeCoffeeMug                    = 28
	CursorTypeCross                        = 30
	CursorTypeCrossReverse                 = 32
	CursorTypeCrosshair                    = 34
	CursorTypeDiamondCross                 = 36
	CursorTypeDot                          = 38
	CursorTypeDotbox                       = 40
	CursorTypeDoubleArrow                  = 42
	CursorTypeDraftLarge                   = 44
	CursorTypeDraftSmall                   = 46
	CursorTypeDrapedBox                    = 48
	CursorTypeExchange                     = 50
	CursorTypeFleur                        = 52
	CursorTypeGobbler                      = 54
	CursorTypeGumby                        = 56
	CursorTypeHand1                        = 58
	CursorTypeHand2                        = 60
	CursorTypeHeart                        = 62
	CursorTypeIcon                         = 64
	CursorTypeIronCross                    = 66
	CursorTypeLeftPtr                      = 68
	CursorTypeLeftSide                     = 70
	CursorTypeLeftTee                      = 72
	CursorTypeLeftbutton                   = 74
	CursorTypeLlAngle                      = 76
	CursorTypeLrAngle                      = 78
	CursorTypeMan                          = 80
	CursorTypeMiddlebutton                 = 82
	CursorTypeMouse                        = 84
	CursorTypePencil                       = 86
	CursorTypePirate                       = 88
	CursorTypePlus                         = 90
	CursorTypeQuestionArrow                = 92
	CursorTypeRightPtr                     = 94
	CursorTypeRightSide                    = 96
	CursorTypeRightTee                     = 98
	CursorTypeRightbutton                  = 100
	CursorTypeRtlLogo                      = 102
	CursorTypeSailboat                     = 104
	CursorTypeSbDownArrow                  = 106
	CursorTypeSbHDoubleArrow               = 108
	CursorTypeSbLeftArrow                  = 110
	CursorTypeSbRightArrow                 = 112
	CursorTypeSbUpArrow                    = 114
	CursorTypeSbVDoubleArrow               = 116
	CursorTypeShuttle                      = 118
	CursorTypeSizing                       = 120
	CursorTypeSpider                       = 122
	CursorTypeSpraycan                     = 124
	CursorTypeStar                         = 126
	CursorTypeTarget                       = 128
	CursorTypeTcross                       = 130
	CursorTypeTopLeftArrow                 = 132
	CursorTypeTopLeftCorner                = 134
	CursorTypeTopRightCorner               = 136
	CursorTypeTopSide                      = 138
	CursorTypeTopTee                       = 140
	CursorTypeTrek                         = 142
	CursorTypeUlAngle                      = 144
	CursorTypeUmbrella                     = 146
	CursorTypeUrAngle                      = 148
	CursorTypeWatch                        = 150
	CursorTypeXterm                        = 152
	CursorTypeLastCursor                   = 153
	CursorTypeBlankCursor                  = -2
	CursorTypeCursorIsPixmap               = -1
)

type DevicePadFeature int

const (
	DevicePadFeatureButton DevicePadFeature = 0
	DevicePadFeatureRing                    = 1
	DevicePadFeatureStrip                   = 2
)

type DeviceToolType int

const (
	DeviceToolTypeUnknown  DeviceToolType = 0
	DeviceToolTypePen                     = 1
	DeviceToolTypeEraser                  = 2
	DeviceToolTypeBrush                   = 3
	DeviceToolTypePencil                  = 4
	DeviceToolTypeAirbrush                = 5
	DeviceToolTypeMouse                   = 6
	DeviceToolTypeLens                    = 7
)

type DeviceType int

const (
	DeviceTypeMaster   DeviceType = 0
	DeviceTypeSlave               = 1
	DeviceTypeFloating            = 2
)

type DragCancelReason int

const (
	DragCancelReasonNoTarget      DragCancelReason = 0
	DragCancelReasonUserCancelled                  = 1
	DragCancelReasonError                          = 2
)

type DragProtocol int

const (
	DragProtocolNone           DragProtocol = 0
	DragProtocolMotif                       = 1
	DragProtocolXdnd                        = 2
	DragProtocolRootwin                     = 3
	DragProtocolWin32Dropfiles              = 4
	DragProtocolOle2                        = 5
	DragProtocolLocal                       = 6
	DragProtocolWayland                     = 7
)

type EventType int

const (
	EventTypeNothing           EventType = -1
	EventTypeDelete                      = 0
	EventTypeDestroy                     = 1
	EventTypeExpose                      = 2
	EventTypeMotionNotify                = 3
	EventTypeButtonPress                 = 4
	EventType2buttonPress                = 5
	EventTypeDoubleButtonPress           = 5
	EventType3buttonPress                = 6
	EventTypeTripleButtonPress           = 6
	EventTypeButtonRelease               = 7
	EventTypeKeyPress                    = 8
	EventTypeKeyRelease                  = 9
	EventTypeEnterNotify                 = 10
	EventTypeLeaveNotify                 = 11
	EventTypeFocusChange                 = 12
	EventTypeConfigure                   = 13
	EventTypeMap                         = 14
	EventTypeUnmap                       = 15
	EventTypePropertyNotify              = 16
	EventTypeSelectionClear              = 17
	EventTypeSelectionRequest            = 18
	EventTypeSelectionNotify             = 19
	EventTypeProximityIn                 = 20
	EventTypeProximityOut                = 21
	EventTypeDragEnter                   = 22
	EventTypeDragLeave                   = 23
	EventTypeDragMotion                  = 24
	EventTypeDragStatus                  = 25
	EventTypeDropStart                   = 26
	EventTypeDropFinished                = 27
	EventTypeClientEvent                 = 28
	EventTypeVisibilityNotify            = 29
	EventTypeScroll                      = 31
	EventTypeWindowState                 = 32
	EventTypeSetting                     = 33
	EventTypeOwnerChange                 = 34
	EventTypeGrabBroken                  = 35
	EventTypeDamage                      = 36
	EventTypeTouchBegin                  = 37
	EventTypeTouchUpdate                 = 38
	EventTypeTouchEnd                    = 39
	EventTypeTouchCancel                 = 40
	EventTypeTouchpadSwipe               = 41
	EventTypeTouchpadPinch               = 42
	EventTypePadButtonPress              = 43
	EventTypePadButtonRelease            = 44
	EventTypePadRing                     = 45
	EventTypePadStrip                    = 46
	EventTypePadGroupMode                = 47
	EventTypeEventLast                   = 48
)

type FilterReturn int

const (
	FilterReturnContinue  FilterReturn = 0
	FilterReturnTranslate              = 1
	FilterReturnRemove                 = 2
)

type FullscreenMode int

const (
	FullscreenModeCurrentMonitor FullscreenMode = 0
	FullscreenModeAllMonitors                   = 1
)

type GLError int

const (
	GLErrorNotAvailable       GLError = 0
	GLErrorUnsupportedFormat          = 1
	GLErrorUnsupportedProfile         = 2
)

type GrabOwnership int

const (
	GrabOwnershipNone        GrabOwnership = 0
	GrabOwnershipWindow                    = 1
	GrabOwnershipApplication               = 2
)

type GrabStatus int

const (
	GrabStatusSuccess        GrabStatus = 0
	GrabStatusAlreadyGrabbed            = 1
	GrabStatusInvalidTime               = 2
	GrabStatusNotViewable               = 3
	GrabStatusFrozen                    = 4
	GrabStatusFailed                    = 5
)

type Gravity int

const (
	GravityNorthWest Gravity = 1
	GravityNorth             = 2
	GravityNorthEast         = 3
	GravityWest              = 4
	GravityCenter            = 5
	GravityEast              = 6
	GravitySouthWest         = 7
	GravitySouth             = 8
	GravitySouthEast         = 9
	GravityStatic            = 10
)

type InputMode int

const (
	InputModeDisabled InputMode = 0
	InputModeScreen             = 1
	InputModeWindow             = 2
)

type InputSource int

const (
	InputSourceMouse       InputSource = 0
	InputSourcePen                     = 1
	InputSourceEraser                  = 2
	InputSourceCursor                  = 3
	InputSourceKeyboard                = 4
	InputSourceTouchscreen             = 5
	InputSourceTouchpad                = 6
	InputSourceTrackpoint              = 7
	InputSourceTabletPad               = 8
)

type ModifierIntent int

const (
	ModifierIntentPrimaryAccelerator ModifierIntent = 0
	ModifierIntentContextMenu                       = 1
	ModifierIntentExtendSelection                   = 2
	ModifierIntentModifySelection                   = 3
	ModifierIntentNoTextInput                       = 4
	ModifierIntentShiftGroup                        = 5
	ModifierIntentDefaultModMask                    = 6
)

type NotifyType int

const (
	NotifyTypeAncestor         NotifyType = 0
	NotifyTypeVirtual                     = 1
	NotifyTypeInferior                    = 2
	NotifyTypeNonlinear                   = 3
	NotifyTypeNonlinearVirtual            = 4
	NotifyTypeUnknown                     = 5
)

type OwnerChange int

const (
	OwnerChangeNewOwner OwnerChange = 0
	OwnerChangeDestroy              = 1
	OwnerChangeClose                = 2
)

type PropMode int

const (
	PropModeReplace PropMode = 0
	PropModePrepend          = 1
	PropModeAppend           = 2
)

type PropertyState int

const (
	PropertyStateNewValue PropertyState = 0
	PropertyStateDelete                 = 1
)

type ScrollDirection int

const (
	ScrollDirectionUp     ScrollDirection = 0
	ScrollDirectionDown                   = 1
	ScrollDirectionLeft                   = 2
	ScrollDirectionRight                  = 3
	ScrollDirectionSmooth                 = 4
)

type SettingAction int

const (
	SettingActionNew     SettingAction = 0
	SettingActionChanged               = 1
	SettingActionDeleted               = 2
)

type Status int

const (
	StatusOk         Status = 0
	StatusError             = -1
	StatusErrorParam        = -2
	StatusErrorFile         = -3
	StatusErrorMem          = -4
)

type SubpixelLayout int

const (
	SubpixelLayoutUnknown       SubpixelLayout = 0
	SubpixelLayoutNone                         = 1
	SubpixelLayoutHorizontalRgb                = 2
	SubpixelLayoutHorizontalBgr                = 3
	SubpixelLayoutVerticalRgb                  = 4
	SubpixelLayoutVerticalBgr                  = 5
)

type TouchpadGesturePhase int

const (
	TouchpadGesturePhaseBegin  TouchpadGesturePhase = 0
	TouchpadGesturePhaseUpdate                      = 1
	TouchpadGesturePhaseEnd                         = 2
	TouchpadGesturePhaseCancel                      = 3
)

type VisibilityState int

const (
	VisibilityStateUnobscured    VisibilityState = 0
	VisibilityStatePartial                       = 1
	VisibilityStateFullyObscured                 = 2
)

type VisualType int

const (
	VisualTypeStaticGray  VisualType = 0
	VisualTypeGrayscale              = 1
	VisualTypeStaticColor            = 2
	VisualTypePseudoColor            = 3
	VisualTypeTrueColor              = 4
	VisualTypeDirectColor            = 5
)

type WindowEdge int

const (
	WindowEdgeNorthWest WindowEdge = 0
	WindowEdgeNorth                = 1
	WindowEdgeNorthEast            = 2
	WindowEdgeWest                 = 3
	WindowEdgeEast                 = 4
	WindowEdgeSouthWest            = 5
	WindowEdgeSouth                = 6
	WindowEdgeSouthEast            = 7
)

type WindowType int

const (
	WindowTypeRoot       WindowType = 0
	WindowTypeToplevel              = 1
	WindowTypeChild                 = 2
	WindowTypeTemp                  = 3
	WindowTypeForeign               = 4
	WindowTypeOffscreen             = 5
	WindowTypeSubsurface            = 6
)

type WindowTypeHint int

const (
	WindowTypeHintNormal       WindowTypeHint = 0
	WindowTypeHintDialog                      = 1
	WindowTypeHintMenu                        = 2
	WindowTypeHintToolbar                     = 3
	WindowTypeHintSplashscreen                = 4
	WindowTypeHintUtility                     = 5
	WindowTypeHintDock                        = 6
	WindowTypeHintDesktop                     = 7
	WindowTypeHintDropdownMenu                = 8
	WindowTypeHintPopupMenu                   = 9
	WindowTypeHintTooltip                     = 10
	WindowTypeHintNotification                = 11
	WindowTypeHintCombo                       = 12
	WindowTypeHintDnd                         = 13
)

type WindowWindowClass int

const (
	WindowWindowClassInputOutput WindowWindowClass = 0
	WindowWindowClassInputOnly                     = 1
)

type AnchorHints int

const (
	AnchorHintsFlipX   AnchorHints = 1
	AnchorHintsFlipY               = 2
	AnchorHintsSlideX              = 4
	AnchorHintsSlideY              = 8
	AnchorHintsResizeX             = 16
	AnchorHintsResizeY             = 32
	AnchorHintsFlip                = 3
	AnchorHintsSlide               = 12
	AnchorHintsResize              = 48
)

type AxisFlags int

const (
	AxisFlagsX        AxisFlags = 2
	AxisFlagsY                  = 4
	AxisFlagsPressure           = 8
	AxisFlagsXtilt              = 16
	AxisFlagsYtilt              = 32
	AxisFlagsWheel              = 64
	AxisFlagsDistance           = 128
	AxisFlagsRotation           = 256
	AxisFlagsSlider             = 512
)

type DragAction int

const (
	DragActionDefault DragAction = 1
	DragActionCopy               = 2
	DragActionMove               = 4
	DragActionLink               = 8
	DragActionPrivate            = 16
	DragActionAsk                = 32
)

type EventMask int

const (
	EventMaskExposureMask          EventMask = 2
	EventMaskPointerMotionMask               = 4
	EventMaskPointerMotionHintMask           = 8
	EventMaskButtonMotionMask                = 16
	EventMaskButton1MotionMask               = 32
	EventMaskButton2MotionMask               = 64
	EventMaskButton3MotionMask               = 128
	EventMaskButtonPressMask                 = 256
	EventMaskButtonReleaseMask               = 512
	EventMaskKeyPressMask                    = 1024
	EventMaskKeyReleaseMask                  = 2048
	EventMaskEnterNotifyMask                 = 4096
	EventMaskLeaveNotifyMask                 = 8192
	EventMaskFocusChangeMask                 = 16384
	EventMaskStructureMask                   = 32768
	EventMaskPropertyChangeMask              = 65536
	EventMaskVisibilityNotifyMask            = 131072
	EventMaskProximityInMask                 = 262144
	EventMaskProximityOutMask                = 524288
	EventMaskSubstructureMask                = 1048576
	EventMaskScrollMask                      = 2097152
	EventMaskTouchMask                       = 4194304
	EventMaskSmoothScrollMask                = 8388608
	EventMaskTouchpadGestureMask             = 16777216
	EventMaskTabletPadMask                   = 33554432
	EventMaskAllEventsMask                   = 16777214
)

type FrameClockPhase int

const (
	FrameClockPhaseNone         FrameClockPhase = 0
	FrameClockPhaseFlushEvents                  = 1
	FrameClockPhaseBeforePaint                  = 2
	FrameClockPhaseUpdate                       = 4
	FrameClockPhaseLayout                       = 8
	FrameClockPhasePaint                        = 16
	FrameClockPhaseResumeEvents                 = 32
	FrameClockPhaseAfterPaint                   = 64
)

type ModifierType int

const (
	ModifierTypeShiftMask              ModifierType = 1
	ModifierTypeLockMask                            = 2
	ModifierTypeControlMask                         = 4
	ModifierTypeMod1Mask                            = 8
	ModifierTypeMod2Mask                            = 16
	ModifierTypeMod3Mask                            = 32
	ModifierTypeMod4Mask                            = 64
	ModifierTypeMod5Mask                            = 128
	ModifierTypeButton1Mask                         = 256
	ModifierTypeButton2Mask                         = 512
	ModifierTypeButton3Mask                         = 1024
	ModifierTypeButton4Mask                         = 2048
	ModifierTypeButton5Mask                         = 4096
	ModifierTypeModifierReserved13Mask              = 8192
	ModifierTypeModifierReserved14Mask              = 16384
	ModifierTypeModifierReserved15Mask              = 32768
	ModifierTypeModifierReserved16Mask              = 65536
	ModifierTypeModifierReserved17Mask              = 131072
	ModifierTypeModifierReserved18Mask              = 262144
	ModifierTypeModifierReserved19Mask              = 524288
	ModifierTypeModifierReserved20Mask              = 1048576
	ModifierTypeModifierReserved21Mask              = 2097152
	ModifierTypeModifierReserved22Mask              = 4194304
	ModifierTypeModifierReserved23Mask              = 8388608
	ModifierTypeModifierReserved24Mask              = 16777216
	ModifierTypeModifierReserved25Mask              = 33554432
	ModifierTypeSuperMask                           = 67108864
	ModifierTypeHyperMask                           = 134217728
	ModifierTypeMetaMask                            = 268435456
	ModifierTypeModifierReserved29Mask              = 536870912
	ModifierTypeReleaseMask                         = 1073741824
	ModifierTypeModifierMask                        = 1543512063
)

type SeatCapabilities int

const (
	SeatCapabilitiesNone         SeatCapabilities = 0
	SeatCapabilitiesPointer                       = 1
	SeatCapabilitiesTouch                         = 2
	SeatCapabilitiesTabletStylus                  = 4
	SeatCapabilitiesKeyboard                      = 8
	SeatCapabilitiesAllPointing                   = 7
	SeatCapabilitiesAll                           = 15
)

type WMDecoration int

const (
	WMDecorationAll      WMDecoration = 1
	WMDecorationBorder                = 2
	WMDecorationResizeh               = 4
	WMDecorationTitle                 = 8
	WMDecorationMenu                  = 16
	WMDecorationMinimize              = 32
	WMDecorationMaximize              = 64
)

type WMFunction int

const (
	WMFunctionAll      WMFunction = 1
	WMFunctionResize              = 2
	WMFunctionMove                = 4
	WMFunctionMinimize            = 8
	WMFunctionMaximize            = 16
	WMFunctionClose               = 32
)

type WindowAttributesType int

const (
	WindowAttributesTypeTitle    WindowAttributesType = 2
	WindowAttributesTypeX                             = 4
	WindowAttributesTypeY                             = 8
	WindowAttributesTypeCursor                        = 16
	WindowAttributesTypeVisual                        = 32
	WindowAttributesTypeWmclass                       = 64
	WindowAttributesTypeNoredir                       = 128
	WindowAttributesTypeTypeHint                      = 256
)

type WindowHints int

const (
	WindowHintsPos        WindowHints = 1
	WindowHintsMinSize                = 2
	WindowHintsMaxSize                = 4
	WindowHintsBaseSize               = 8
	WindowHintsAspect                 = 16
	WindowHintsResizeInc              = 32
	WindowHintsWinGravity             = 64
	WindowHintsUserPos                = 128
	WindowHintsUserSize               = 256
)

type WindowState int

const (
	WindowStateWithdrawn  WindowState = 1
	WindowStateIconified              = 2
	WindowStateMaximized              = 4
	WindowStateSticky                 = 8
	WindowStateFullscreen             = 16
	WindowStateAbove                  = 32
	WindowStateBelow                  = 64
	WindowStateFocused                = 128
	WindowStateTiled                  = 256
)
