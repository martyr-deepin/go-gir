package gio

/*
#cgo pkg-config: gio-2.0 gio-unix-2.0
#include <stdlib.h>
#include <gio/gdesktopappinfo.h>
#include <gio/gfiledescriptorbased.h>
#include <gio/gio.h>
#include <gio/gunixconnection.h>
#include <gio/gunixcredentialsmessage.h>
#include <gio/gunixfdlist.h>
#include <gio/gunixfdmessage.h>
#include <gio/gunixinputstream.h>
#include <gio/gunixmounts.h>
#include <gio/gunixoutputstream.h>
#include <gio/gunixsocketaddress.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/gobject-2.0"
import "github.com/electricface/go-auto-gir/util"
import "github.com/electricface/go-auto-gir/glib-2.0"

// Object Settings
type Settings struct {
	gobject.Object
}

func (v Settings) native() *C.GSettings {
	return (*C.GSettings)(v.Ptr)
}
func wrapSettings(p *C.GSettings) (v Settings) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSettings(p unsafe.Pointer) (v Settings) {
	v.Ptr = p
	return
}

func (v Settings) GetType() gobject.Type {
	return gobject.Type(C.g_settings_get_type())
}

func (v Settings) GetGValueGetter() gobject.GValueGetter {
	return gvalueGetSettings
}

func gvalueGetSettings(p unsafe.Pointer) (interface{}, error) {
	ret := C.g_value_get_object((*C.GValue)(p))
	return WrapSettings(unsafe.Pointer(ret)), nil
}

// SettingsNew is a wrapper around g_settings_new().
func SettingsNew(schema_id string) Settings {

	// Var for Go: schema_id
	// Var for C: schema_id0
	// Type for Go: string
	// Type for C: *C.gchar
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	defer C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_new(schema_id0)
	return wrapSettings(ret0)
}

// Apply is a wrapper around g_settings_apply().
func (settings Settings) Apply() {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings
	C.g_settings_apply(settings.native())
}

// GetBoolean is a wrapper around g_settings_get_boolean().
func (settings Settings) GetBoolean(key string) bool {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_boolean(settings.native(), key0)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChild is a wrapper around g_settings_get_child().
func (settings Settings) GetChild(name string) Settings {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: name
	// Var for C: name0
	// Type for Go: string
	// Type for C: *C.gchar
	name0 := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_child(settings.native(), name0)
	return wrapSettings(ret0)
}

// GetDouble is a wrapper around g_settings_get_double().
func (settings Settings) GetDouble(key string) float64 {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_double(settings.native(), key0)
	return float64(ret0)
}

// GetEnum is a wrapper around g_settings_get_enum().
func (settings Settings) GetEnum(key string) int {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_enum(settings.native(), key0)
	return int(ret0)
}

// GetFlags is a wrapper around g_settings_get_flags().
func (settings Settings) GetFlags(key string) uint {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_flags(settings.native(), key0)
	return uint(ret0)
}

// GetInt is a wrapper around g_settings_get_int().
func (settings Settings) GetInt(key string) int {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_int(settings.native(), key0)
	return int(ret0)
}

// GetInt64 is a wrapper around g_settings_get_int64().
func (settings Settings) GetInt64(key string) int64 {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_int64(settings.native(), key0)
	return int64(ret0)
}

// GetString is a wrapper around g_settings_get_string().
func (settings Settings) GetString(key string) string {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_string(settings.native(), key0)
	defer C.g_free(C.gpointer(ret0))
	return C.GoString((*C.char)(ret0))
}

// GetUint is a wrapper around g_settings_get_uint().
func (settings Settings) GetUint(key string) uint {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_uint(settings.native(), key0)
	return uint(ret0)
}

// GetUint64 is a wrapper around g_settings_get_uint64().
func (settings Settings) GetUint64(key string) uint64 {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_uint64(settings.native(), key0)
	return uint64(ret0)
}

// GetValue is a wrapper around g_settings_get_value().
func (settings Settings) GetValue(key string) glib.Variant {

	// Var for Go: settings
	// Var for C: settings0
	// Type for Go: Settings
	// Type for C: *C.GSettings

	// Var for Go: key
	// Var for C: key0
	// Type for Go: string
	// Type for C: *C.gchar
	key0 := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret0 := C.g_settings_get_value(settings.native(), key0)
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:glib*/
}
