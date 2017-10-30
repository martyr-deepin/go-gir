package gio

/*
#cgo pkg-config: gio-2.0 gio-unix-2.0
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
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/gobject-2.0"
import "github.com/electricface/go-auto-gir/util"
import "github.com/electricface/go-auto-gir/glib-2.0"

// Interface AppInfo
type AppInfo struct {
	Ptr unsafe.Pointer
}

func (v AppInfo) native() *C.GAppInfo {
	return (*C.GAppInfo)(v.Ptr)
}
func wrapAppInfo(p *C.GAppInfo) AppInfo {
	return AppInfo{unsafe.Pointer(p)}
}
func WrapAppInfo(p unsafe.Pointer) AppInfo {
	return AppInfo{p}
}

// GetId is a wrapper around g_app_info_get_id().
func (appinfo AppInfo) GetId() string {

	// Var for Go: appinfo
	// Var for C: appinfo0
	// Type for Go: AppInfo
	// Type for C: *C.GAppInfo
	ret0 := C.g_app_info_get_id(appinfo.native())
	return C.GoString(ret0)
}

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

// Interface File
type File struct {
	Ptr unsafe.Pointer
}

func (v File) native() *C.GFile {
	return (*C.GFile)(v.Ptr)
}
func wrapFile(p *C.GFile) File {
	return File{unsafe.Pointer(p)}
}
func WrapFile(p unsafe.Pointer) File {
	return File{p}
}

// GetPath is a wrapper around g_file_get_path().
func (file File) GetPath() string {

	// Var for Go: file
	// Var for C: file0
	// Type for Go: File
	// Type for C: *C.GFile
	ret0 := C.g_file_get_path(file.native())
	defer C.g_free(C.gpointer(ret0))
	return C.GoString(ret0)
}

// Replace is a wrapper around g_file_replace().
func (file File) Replace(etag string, make_backup bool, flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {

	// Var for Go: file
	// Var for C: file0
	// Type for Go: File
	// Type for C: *C.GFile

	// Var for Go: etag
	// Var for C: etag0
	// Type for Go: string
	// Type for C: *C.char
	etag0 := C.CString(etag)
	defer C.free(unsafe.Pointer(etag0)) /*ch:<stdlib.h>*/

	// Var for Go: make_backup
	// Var for C: make_backup0
	// Type for Go: bool
	// Type for C: C.gboolean

	// Var for Go: flags
	// Var for C: flags0
	// Type for Go: FileCreateFlags
	// Type for C: C.GFileCreateFlags

	// Var for Go: cancellable
	// Var for C: cancellable0
	// Type for Go: Cancellable
	// Type for C: *C.GCancellable
	var err glib.Error
	ret0 := C.g_file_replace(file.native(), etag0, C.gboolean(util.Bool2Int(make_backup)) /*go:.util*/, C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// Object Cancellable
type Cancellable struct {
	gobject.Object
}

func (v Cancellable) native() *C.GCancellable {
	return (*C.GCancellable)(v.Ptr)
}
func wrapCancellable(p *C.GCancellable) (v Cancellable) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCancellable(p unsafe.Pointer) (v Cancellable) {
	v.Ptr = p
	return
}

// Object OutputStream
type OutputStream struct {
	gobject.Object
}

func (v OutputStream) native() *C.GOutputStream {
	return (*C.GOutputStream)(v.Ptr)
}
func wrapOutputStream(p *C.GOutputStream) (v OutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapOutputStream(p unsafe.Pointer) (v OutputStream) {
	v.Ptr = p
	return
}

// Object FileOutputStream
type FileOutputStream struct {
	OutputStream
}

func (v FileOutputStream) native() *C.GFileOutputStream {
	return (*C.GFileOutputStream)(v.Ptr)
}
func wrapFileOutputStream(p *C.GFileOutputStream) (v FileOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileOutputStream(p unsafe.Pointer) (v FileOutputStream) {
	v.Ptr = p
	return
}

// QueryInfo is a wrapper around g_file_output_stream_query_info().
func (stream FileOutputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {

	// Var for Go: stream
	// Var for C: stream0
	// Type for Go: FileOutputStream
	// Type for C: *C.GFileOutputStream

	// Var for Go: attributes
	// Var for C: attributes0
	// Type for Go: string
	// Type for C: *C.char
	attributes0 := C.CString(attributes)
	defer C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/

	// Var for Go: cancellable
	// Var for C: cancellable0
	// Type for Go: Cancellable
	// Type for C: *C.GCancellable
	var err glib.Error
	ret0 := C.g_file_output_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// Object FileInfo
type FileInfo struct {
	gobject.Object
}

func (v FileInfo) native() *C.GFileInfo {
	return (*C.GFileInfo)(v.Ptr)
}
func wrapFileInfo(p *C.GFileInfo) (v FileInfo) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileInfo(p unsafe.Pointer) (v FileInfo) {
	v.Ptr = p
	return
}

// Object Application
type Application struct {
	gobject.Object
}

func (v Application) native() *C.GApplication {
	return (*C.GApplication)(v.Ptr)
}
func wrapApplication(p *C.GApplication) (v Application) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapApplication(p unsafe.Pointer) (v Application) {
	v.Ptr = p
	return
}

// ApplicationNew is a wrapper around g_application_new().
func ApplicationNew(application_id string, flags ApplicationFlags) Application {

	// Var for Go: application_id
	// Var for C: application_id0
	// Type for Go: string
	// Type for C: *C.gchar
	application_id0 := (*C.gchar)(C.CString(application_id))
	defer C.free(unsafe.Pointer(application_id0)) /*ch:<stdlib.h>*/

	// Var for Go: flags
	// Var for C: flags0
	// Type for Go: ApplicationFlags
	// Type for C: C.GApplicationFlags
	ret0 := C.g_application_new(application_id0, C.GApplicationFlags(flags))
	return wrapApplication(ret0)
}
