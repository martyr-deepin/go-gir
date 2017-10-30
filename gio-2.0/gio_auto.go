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

type BusType C.GBusType

const (
	BusTypeStarter BusType = -1
	BusTypeNone            = 0
	BusTypeSystem          = 1
	BusTypeSession         = 2
)

type ConverterResult C.GConverterResult

const (
	ConverterResultError     ConverterResult = 0
	ConverterResultConverted                 = 1
	ConverterResultFinished                  = 2
	ConverterResultFlushed                   = 3
)

type CredentialsType C.GCredentialsType

const (
	CredentialsTypeInvalid             CredentialsType = 0
	CredentialsTypeLinuxUcred                          = 1
	CredentialsTypeFreebsdCmsgcred                     = 2
	CredentialsTypeOpenbsdSockpeercred                 = 3
	CredentialsTypeSolarisUcred                        = 4
	CredentialsTypeNetbsdUnpcbid                       = 5
)

type DBusError C.GDBusError

const (
	DBusErrorFailed                        DBusError = 0
	DBusErrorNoMemory                                = 1
	DBusErrorServiceUnknown                          = 2
	DBusErrorNameHasNoOwner                          = 3
	DBusErrorNoReply                                 = 4
	DBusErrorIoError                                 = 5
	DBusErrorBadAddress                              = 6
	DBusErrorNotSupported                            = 7
	DBusErrorLimitsExceeded                          = 8
	DBusErrorAccessDenied                            = 9
	DBusErrorAuthFailed                              = 10
	DBusErrorNoServer                                = 11
	DBusErrorTimeout                                 = 12
	DBusErrorNoNetwork                               = 13
	DBusErrorAddressInUse                            = 14
	DBusErrorDisconnected                            = 15
	DBusErrorInvalidArgs                             = 16
	DBusErrorFileNotFound                            = 17
	DBusErrorFileExists                              = 18
	DBusErrorUnknownMethod                           = 19
	DBusErrorTimedOut                                = 20
	DBusErrorMatchRuleNotFound                       = 21
	DBusErrorMatchRuleInvalid                        = 22
	DBusErrorSpawnExecFailed                         = 23
	DBusErrorSpawnForkFailed                         = 24
	DBusErrorSpawnChildExited                        = 25
	DBusErrorSpawnChildSignaled                      = 26
	DBusErrorSpawnFailed                             = 27
	DBusErrorSpawnSetupFailed                        = 28
	DBusErrorSpawnConfigInvalid                      = 29
	DBusErrorSpawnServiceInvalid                     = 30
	DBusErrorSpawnServiceNotFound                    = 31
	DBusErrorSpawnPermissionsInvalid                 = 32
	DBusErrorSpawnFileInvalid                        = 33
	DBusErrorSpawnNoMemory                           = 34
	DBusErrorUnixProcessIdUnknown                    = 35
	DBusErrorInvalidSignature                        = 36
	DBusErrorInvalidFileContent                      = 37
	DBusErrorSelinuxSecurityContextUnknown           = 38
	DBusErrorAdtAuditDataUnknown                     = 39
	DBusErrorObjectPathInUse                         = 40
	DBusErrorUnknownObject                           = 41
	DBusErrorUnknownInterface                        = 42
	DBusErrorUnknownProperty                         = 43
	DBusErrorPropertyReadOnly                        = 44
)

type DBusMessageByteOrder C.GDBusMessageByteOrder

const (
	DBusMessageByteOrderBigEndian    DBusMessageByteOrder = 66
	DBusMessageByteOrderLittleEndian                      = 108
)

type DBusMessageHeaderField C.GDBusMessageHeaderField

const (
	DBusMessageHeaderFieldInvalid     DBusMessageHeaderField = 0
	DBusMessageHeaderFieldPath                               = 1
	DBusMessageHeaderFieldInterface                          = 2
	DBusMessageHeaderFieldMember                             = 3
	DBusMessageHeaderFieldErrorName                          = 4
	DBusMessageHeaderFieldReplySerial                        = 5
	DBusMessageHeaderFieldDestination                        = 6
	DBusMessageHeaderFieldSender                             = 7
	DBusMessageHeaderFieldSignature                          = 8
	DBusMessageHeaderFieldNumUnixFds                         = 9
)

type DBusMessageType C.GDBusMessageType

const (
	DBusMessageTypeInvalid      DBusMessageType = 0
	DBusMessageTypeMethodCall                   = 1
	DBusMessageTypeMethodReturn                 = 2
	DBusMessageTypeError                        = 3
	DBusMessageTypeSignal                       = 4
)

type DataStreamByteOrder C.GDataStreamByteOrder

const (
	DataStreamByteOrderBigEndian    DataStreamByteOrder = 0
	DataStreamByteOrderLittleEndian                     = 1
	DataStreamByteOrderHostEndian                       = 2
)

type DataStreamNewlineType C.GDataStreamNewlineType

const (
	DataStreamNewlineTypeLf   DataStreamNewlineType = 0
	DataStreamNewlineTypeCr                         = 1
	DataStreamNewlineTypeCrLf                       = 2
	DataStreamNewlineTypeAny                        = 3
)

type DriveStartStopType C.GDriveStartStopType

const (
	DriveStartStopTypeUnknown   DriveStartStopType = 0
	DriveStartStopTypeShutdown                     = 1
	DriveStartStopTypeNetwork                      = 2
	DriveStartStopTypeMultidisk                    = 3
	DriveStartStopTypePassword                     = 4
)

type EmblemOrigin C.GEmblemOrigin

const (
	EmblemOriginUnknown      EmblemOrigin = 0
	EmblemOriginDevice                    = 1
	EmblemOriginLivemetadata              = 2
	EmblemOriginTag                       = 3
)

type FileAttributeStatus C.GFileAttributeStatus

const (
	FileAttributeStatusUnset        FileAttributeStatus = 0
	FileAttributeStatusSet                              = 1
	FileAttributeStatusErrorSetting                     = 2
)

type FileAttributeType C.GFileAttributeType

const (
	FileAttributeTypeInvalid    FileAttributeType = 0
	FileAttributeTypeString                       = 1
	FileAttributeTypeByteString                   = 2
	FileAttributeTypeBoolean                      = 3
	FileAttributeTypeUint32                       = 4
	FileAttributeTypeInt32                        = 5
	FileAttributeTypeUint64                       = 6
	FileAttributeTypeInt64                        = 7
	FileAttributeTypeObject                       = 8
	FileAttributeTypeStringv                      = 9
)

type FileMonitorEvent C.GFileMonitorEvent

const (
	FileMonitorEventChanged          FileMonitorEvent = 0
	FileMonitorEventChangesDoneHint                   = 1
	FileMonitorEventDeleted                           = 2
	FileMonitorEventCreated                           = 3
	FileMonitorEventAttributeChanged                  = 4
	FileMonitorEventPreUnmount                        = 5
	FileMonitorEventUnmounted                         = 6
	FileMonitorEventMoved                             = 7
	FileMonitorEventRenamed                           = 8
	FileMonitorEventMovedIn                           = 9
	FileMonitorEventMovedOut                          = 10
)

type FileType C.GFileType

const (
	FileTypeUnknown      FileType = 0
	FileTypeRegular               = 1
	FileTypeDirectory             = 2
	FileTypeSymbolicLink          = 3
	FileTypeSpecial               = 4
	FileTypeShortcut              = 5
	FileTypeMountable             = 6
)

type FilesystemPreviewType C.GFilesystemPreviewType

const (
	FilesystemPreviewTypeIfAlways FilesystemPreviewType = 0
	FilesystemPreviewTypeIfLocal                        = 1
	FilesystemPreviewTypeNever                          = 2
)

type IOErrorEnum C.GIOErrorEnum

const (
	IOErrorEnumFailed             IOErrorEnum = 0
	IOErrorEnumNotFound                       = 1
	IOErrorEnumExists                         = 2
	IOErrorEnumIsDirectory                    = 3
	IOErrorEnumNotDirectory                   = 4
	IOErrorEnumNotEmpty                       = 5
	IOErrorEnumNotRegularFile                 = 6
	IOErrorEnumNotSymbolicLink                = 7
	IOErrorEnumNotMountableFile               = 8
	IOErrorEnumFilenameTooLong                = 9
	IOErrorEnumInvalidFilename                = 10
	IOErrorEnumTooManyLinks                   = 11
	IOErrorEnumNoSpace                        = 12
	IOErrorEnumInvalidArgument                = 13
	IOErrorEnumPermissionDenied               = 14
	IOErrorEnumNotSupported                   = 15
	IOErrorEnumNotMounted                     = 16
	IOErrorEnumAlreadyMounted                 = 17
	IOErrorEnumClosed                         = 18
	IOErrorEnumCancelled                      = 19
	IOErrorEnumPending                        = 20
	IOErrorEnumReadOnly                       = 21
	IOErrorEnumCantCreateBackup               = 22
	IOErrorEnumWrongEtag                      = 23
	IOErrorEnumTimedOut                       = 24
	IOErrorEnumWouldRecurse                   = 25
	IOErrorEnumBusy                           = 26
	IOErrorEnumWouldBlock                     = 27
	IOErrorEnumHostNotFound                   = 28
	IOErrorEnumWouldMerge                     = 29
	IOErrorEnumFailedHandled                  = 30
	IOErrorEnumTooManyOpenFiles               = 31
	IOErrorEnumNotInitialized                 = 32
	IOErrorEnumAddressInUse                   = 33
	IOErrorEnumPartialInput                   = 34
	IOErrorEnumInvalidData                    = 35
	IOErrorEnumDbusError                      = 36
	IOErrorEnumHostUnreachable                = 37
	IOErrorEnumNetworkUnreachable             = 38
	IOErrorEnumConnectionRefused              = 39
	IOErrorEnumProxyFailed                    = 40
	IOErrorEnumProxyAuthFailed                = 41
	IOErrorEnumProxyNeedAuth                  = 42
	IOErrorEnumProxyNotAllowed                = 43
	IOErrorEnumBrokenPipe                     = 44
	IOErrorEnumConnectionClosed               = 44
	IOErrorEnumNotConnected                   = 45
	IOErrorEnumMessageTooLarge                = 46
)

type IOModuleScopeFlags C.GIOModuleScopeFlags

const (
	IOModuleScopeFlagsNone            IOModuleScopeFlags = 0
	IOModuleScopeFlagsBlockDuplicates                    = 1
)

type MountOperationResult C.GMountOperationResult

const (
	MountOperationResultHandled   MountOperationResult = 0
	MountOperationResultAborted                        = 1
	MountOperationResultUnhandled                      = 2
)

type NetworkConnectivity C.GNetworkConnectivity

const (
	NetworkConnectivityLocal   NetworkConnectivity = 1
	NetworkConnectivityLimited                     = 2
	NetworkConnectivityPortal                      = 3
	NetworkConnectivityFull                        = 4
)

type NotificationPriority C.GNotificationPriority

const (
	NotificationPriorityNormal NotificationPriority = 0
	NotificationPriorityLow                         = 1
	NotificationPriorityHigh                        = 2
	NotificationPriorityUrgent                      = 3
)

type PasswordSave C.GPasswordSave

const (
	PasswordSaveNever       PasswordSave = 0
	PasswordSaveForSession               = 1
	PasswordSavePermanently              = 2
)

type ResolverError C.GResolverError

const (
	ResolverErrorNotFound         ResolverError = 0
	ResolverErrorTemporaryFailure               = 1
	ResolverErrorInternal                       = 2
)

type ResolverRecordType C.GResolverRecordType

const (
	ResolverRecordTypeSrv ResolverRecordType = 1
	ResolverRecordTypeMx                     = 2
	ResolverRecordTypeTxt                    = 3
	ResolverRecordTypeSoa                    = 4
	ResolverRecordTypeNs                     = 5
)

type ResourceError C.GResourceError

const (
	ResourceErrorNotFound ResourceError = 0
	ResourceErrorInternal               = 1
)

type SocketClientEvent C.GSocketClientEvent

const (
	SocketClientEventResolving        SocketClientEvent = 0
	SocketClientEventResolved                           = 1
	SocketClientEventConnecting                         = 2
	SocketClientEventConnected                          = 3
	SocketClientEventProxyNegotiating                   = 4
	SocketClientEventProxyNegotiated                    = 5
	SocketClientEventTlsHandshaking                     = 6
	SocketClientEventTlsHandshaked                      = 7
	SocketClientEventComplete                           = 8
)

type SocketFamily C.GSocketFamily

const (
	SocketFamilyInvalid SocketFamily = 0
	SocketFamilyUnix                 = 1
	SocketFamilyIpv4                 = 2
	SocketFamilyIpv6                 = 10
)

type SocketListenerEvent C.GSocketListenerEvent

const (
	SocketListenerEventBinding   SocketListenerEvent = 0
	SocketListenerEventBound                         = 1
	SocketListenerEventListening                     = 2
	SocketListenerEventListened                      = 3
)

type SocketProtocol C.GSocketProtocol

const (
	SocketProtocolUnknown SocketProtocol = -1
	SocketProtocolDefault                = 0
	SocketProtocolTcp                    = 6
	SocketProtocolUdp                    = 17
	SocketProtocolSctp                   = 132
)

type SocketType C.GSocketType

const (
	SocketTypeInvalid   SocketType = 0
	SocketTypeStream               = 1
	SocketTypeDatagram             = 2
	SocketTypeSeqpacket            = 3
)

type TlsAuthenticationMode C.GTlsAuthenticationMode

const (
	TlsAuthenticationModeNone      TlsAuthenticationMode = 0
	TlsAuthenticationModeRequested                       = 1
	TlsAuthenticationModeRequired                        = 2
)

type TlsCertificateRequestFlags C.GTlsCertificateRequestFlags

const (
	TlsCertificateRequestFlagsNone TlsCertificateRequestFlags = 0
)

type TlsDatabaseLookupFlags C.GTlsDatabaseLookupFlags

const (
	TlsDatabaseLookupFlagsNone    TlsDatabaseLookupFlags = 0
	TlsDatabaseLookupFlagsKeypair                        = 1
)

type TlsError C.GTlsError

const (
	TlsErrorUnavailable         TlsError = 0
	TlsErrorMisc                         = 1
	TlsErrorBadCertificate               = 2
	TlsErrorNotTls                       = 3
	TlsErrorHandshake                    = 4
	TlsErrorCertificateRequired          = 5
	TlsErrorEof                          = 6
)

type TlsInteractionResult C.GTlsInteractionResult

const (
	TlsInteractionResultUnhandled TlsInteractionResult = 0
	TlsInteractionResultHandled                        = 1
	TlsInteractionResultFailed                         = 2
)

type TlsRehandshakeMode C.GTlsRehandshakeMode

const (
	TlsRehandshakeModeNever    TlsRehandshakeMode = 0
	TlsRehandshakeModeSafely                      = 1
	TlsRehandshakeModeUnsafely                    = 2
)

type UnixSocketAddressType C.GUnixSocketAddressType

const (
	UnixSocketAddressTypeInvalid        UnixSocketAddressType = 0
	UnixSocketAddressTypeAnonymous                            = 1
	UnixSocketAddressTypePath                                 = 2
	UnixSocketAddressTypeAbstract                             = 3
	UnixSocketAddressTypeAbstractPadded                       = 4
)

type ZlibCompressorFormat C.GZlibCompressorFormat

const (
	ZlibCompressorFormatZlib ZlibCompressorFormat = 0
	ZlibCompressorFormatGzip                      = 1
	ZlibCompressorFormatRaw                       = 2
)

type AppInfoCreateFlags C.GAppInfoCreateFlags

const (
	AppInfoCreateFlagsNone                        AppInfoCreateFlags = 0
	AppInfoCreateFlagsNeedsTerminal                                  = 1
	AppInfoCreateFlagsSupportsUris                                   = 2
	AppInfoCreateFlagsSupportsStartupNotification                    = 4
)

type ApplicationFlags C.GApplicationFlags

const (
	ApplicationFlagsFlagsNone          ApplicationFlags = 0
	ApplicationFlagsIsService                           = 1
	ApplicationFlagsIsLauncher                          = 2
	ApplicationFlagsHandlesOpen                         = 4
	ApplicationFlagsHandlesCommandLine                  = 8
	ApplicationFlagsSendEnvironment                     = 16
	ApplicationFlagsNonUnique                           = 32
	ApplicationFlagsCanOverrideAppId                    = 64
)

type AskPasswordFlags C.GAskPasswordFlags

const (
	AskPasswordFlagsNeedPassword       AskPasswordFlags = 1
	AskPasswordFlagsNeedUsername                        = 2
	AskPasswordFlagsNeedDomain                          = 4
	AskPasswordFlagsSavingSupported                     = 8
	AskPasswordFlagsAnonymousSupported                  = 16
)

type BusNameOwnerFlags C.GBusNameOwnerFlags

const (
	BusNameOwnerFlagsNone             BusNameOwnerFlags = 0
	BusNameOwnerFlagsAllowReplacement                   = 1
	BusNameOwnerFlagsReplace                            = 2
)

type BusNameWatcherFlags C.GBusNameWatcherFlags

const (
	BusNameWatcherFlagsNone      BusNameWatcherFlags = 0
	BusNameWatcherFlagsAutoStart                     = 1
)

type ConverterFlags C.GConverterFlags

const (
	ConverterFlagsNone       ConverterFlags = 0
	ConverterFlagsInputAtEnd                = 1
	ConverterFlagsFlush                     = 2
)

type DBusCallFlags C.GDBusCallFlags

const (
	DBusCallFlagsNone                          DBusCallFlags = 0
	DBusCallFlagsNoAutoStart                                 = 1
	DBusCallFlagsAllowInteractiveAuthorization               = 2
)

type DBusCapabilityFlags C.GDBusCapabilityFlags

const (
	DBusCapabilityFlagsNone          DBusCapabilityFlags = 0
	DBusCapabilityFlagsUnixFdPassing                     = 1
)

type DBusConnectionFlags C.GDBusConnectionFlags

const (
	DBusConnectionFlagsNone                         DBusConnectionFlags = 0
	DBusConnectionFlagsAuthenticationClient                             = 1
	DBusConnectionFlagsAuthenticationServer                             = 2
	DBusConnectionFlagsAuthenticationAllowAnonymous                     = 4
	DBusConnectionFlagsMessageBusConnection                             = 8
	DBusConnectionFlagsDelayMessageProcessing                           = 16
)

type DBusInterfaceSkeletonFlags C.GDBusInterfaceSkeletonFlags

const (
	DBusInterfaceSkeletonFlagsNone                            DBusInterfaceSkeletonFlags = 0
	DBusInterfaceSkeletonFlagsHandleMethodInvocationsInThread                            = 1
)

type DBusMessageFlags C.GDBusMessageFlags

const (
	DBusMessageFlagsNone                          DBusMessageFlags = 0
	DBusMessageFlagsNoReplyExpected                                = 1
	DBusMessageFlagsNoAutoStart                                    = 2
	DBusMessageFlagsAllowInteractiveAuthorization                  = 4
)

type DBusObjectManagerClientFlags C.GDBusObjectManagerClientFlags

const (
	DBusObjectManagerClientFlagsNone           DBusObjectManagerClientFlags = 0
	DBusObjectManagerClientFlagsDoNotAutoStart                              = 1
)

type DBusPropertyInfoFlags C.GDBusPropertyInfoFlags

const (
	DBusPropertyInfoFlagsNone     DBusPropertyInfoFlags = 0
	DBusPropertyInfoFlagsReadable                       = 1
	DBusPropertyInfoFlagsWritable                       = 2
)

type DBusProxyFlags C.GDBusProxyFlags

const (
	DBusProxyFlagsNone                         DBusProxyFlags = 0
	DBusProxyFlagsDoNotLoadProperties                         = 1
	DBusProxyFlagsDoNotConnectSignals                         = 2
	DBusProxyFlagsDoNotAutoStart                              = 4
	DBusProxyFlagsGetInvalidatedProperties                    = 8
	DBusProxyFlagsDoNotAutoStartAtConstruction                = 16
)

type DBusSendMessageFlags C.GDBusSendMessageFlags

const (
	DBusSendMessageFlagsNone           DBusSendMessageFlags = 0
	DBusSendMessageFlagsPreserveSerial                      = 1
)

type DBusServerFlags C.GDBusServerFlags

const (
	DBusServerFlagsNone                         DBusServerFlags = 0
	DBusServerFlagsRunInThread                                  = 1
	DBusServerFlagsAuthenticationAllowAnonymous                 = 2
)

type DBusSignalFlags C.GDBusSignalFlags

const (
	DBusSignalFlagsNone               DBusSignalFlags = 0
	DBusSignalFlagsNoMatchRule                        = 1
	DBusSignalFlagsMatchArg0Namespace                 = 2
	DBusSignalFlagsMatchArg0Path                      = 4
)

type DBusSubtreeFlags C.GDBusSubtreeFlags

const (
	DBusSubtreeFlagsNone                        DBusSubtreeFlags = 0
	DBusSubtreeFlagsDispatchToUnenumeratedNodes                  = 1
)

type DriveStartFlags C.GDriveStartFlags

const (
	DriveStartFlagsNone DriveStartFlags = 0
)

type FileAttributeInfoFlags C.GFileAttributeInfoFlags

const (
	FileAttributeInfoFlagsNone          FileAttributeInfoFlags = 0
	FileAttributeInfoFlagsCopyWithFile                         = 1
	FileAttributeInfoFlagsCopyWhenMoved                        = 2
)

type FileCopyFlags C.GFileCopyFlags

const (
	FileCopyFlagsNone               FileCopyFlags = 0
	FileCopyFlagsOverwrite                        = 1
	FileCopyFlagsBackup                           = 2
	FileCopyFlagsNofollowSymlinks                 = 4
	FileCopyFlagsAllMetadata                      = 8
	FileCopyFlagsNoFallbackForMove                = 16
	FileCopyFlagsTargetDefaultPerms               = 32
)

type FileCreateFlags C.GFileCreateFlags

const (
	FileCreateFlagsNone               FileCreateFlags = 0
	FileCreateFlagsPrivate                            = 1
	FileCreateFlagsReplaceDestination                 = 2
)

type FileMeasureFlags C.GFileMeasureFlags

const (
	FileMeasureFlagsNone           FileMeasureFlags = 0
	FileMeasureFlagsReportAnyError                  = 2
	FileMeasureFlagsApparentSize                    = 4
	FileMeasureFlagsNoXdev                          = 8
)

type FileMonitorFlags C.GFileMonitorFlags

const (
	FileMonitorFlagsNone           FileMonitorFlags = 0
	FileMonitorFlagsWatchMounts                     = 1
	FileMonitorFlagsSendMoved                       = 2
	FileMonitorFlagsWatchHardLinks                  = 4
	FileMonitorFlagsWatchMoves                      = 8
)

type FileQueryInfoFlags C.GFileQueryInfoFlags

const (
	FileQueryInfoFlagsNone             FileQueryInfoFlags = 0
	FileQueryInfoFlagsNofollowSymlinks                    = 1
)

type IOStreamSpliceFlags C.GIOStreamSpliceFlags

const (
	IOStreamSpliceFlagsNone         IOStreamSpliceFlags = 0
	IOStreamSpliceFlagsCloseStream1                     = 1
	IOStreamSpliceFlagsCloseStream2                     = 2
	IOStreamSpliceFlagsWaitForBoth                      = 4
)

type MountMountFlags C.GMountMountFlags

const (
	MountMountFlagsNone MountMountFlags = 0
)

type MountUnmountFlags C.GMountUnmountFlags

const (
	MountUnmountFlagsNone  MountUnmountFlags = 0
	MountUnmountFlagsForce                   = 1
)

type OutputStreamSpliceFlags C.GOutputStreamSpliceFlags

const (
	OutputStreamSpliceFlagsNone        OutputStreamSpliceFlags = 0
	OutputStreamSpliceFlagsCloseSource                         = 1
	OutputStreamSpliceFlagsCloseTarget                         = 2
)

type ResourceFlags C.GResourceFlags

const (
	ResourceFlagsNone       ResourceFlags = 0
	ResourceFlagsCompressed               = 1
)

type ResourceLookupFlags C.GResourceLookupFlags

const (
	ResourceLookupFlagsNone ResourceLookupFlags = 0
)

type SettingsBindFlags C.GSettingsBindFlags

const (
	SettingsBindFlagsDefault       SettingsBindFlags = 0
	SettingsBindFlagsGet                             = 1
	SettingsBindFlagsSet                             = 2
	SettingsBindFlagsNoSensitivity                   = 4
	SettingsBindFlagsGetNoChanges                    = 8
	SettingsBindFlagsInvertBoolean                   = 16
)

type SocketMsgFlags C.GSocketMsgFlags

const (
	SocketMsgFlagsNone      SocketMsgFlags = 0
	SocketMsgFlagsOob                      = 1
	SocketMsgFlagsPeek                     = 2
	SocketMsgFlagsDontroute                = 4
)

type SubprocessFlags C.GSubprocessFlags

const (
	SubprocessFlagsNone          SubprocessFlags = 0
	SubprocessFlagsStdinPipe                     = 1
	SubprocessFlagsStdinInherit                  = 2
	SubprocessFlagsStdoutPipe                    = 4
	SubprocessFlagsStdoutSilence                 = 8
	SubprocessFlagsStderrPipe                    = 16
	SubprocessFlagsStderrSilence                 = 32
	SubprocessFlagsStderrMerge                   = 64
	SubprocessFlagsInheritFds                    = 128
)

type TestDBusFlags C.GTestDBusFlags

const (
	TestDBusFlagsNone TestDBusFlags = 0
)

type TlsCertificateFlags C.GTlsCertificateFlags

const (
	TlsCertificateFlagsUnknownCa    TlsCertificateFlags = 1
	TlsCertificateFlagsBadIdentity                      = 2
	TlsCertificateFlagsNotActivated                     = 4
	TlsCertificateFlagsExpired                          = 8
	TlsCertificateFlagsRevoked                          = 16
	TlsCertificateFlagsInsecure                         = 32
	TlsCertificateFlagsGenericError                     = 64
	TlsCertificateFlagsValidateAll                      = 127
)

type TlsDatabaseVerifyFlags C.GTlsDatabaseVerifyFlags

const (
	TlsDatabaseVerifyFlagsNone TlsDatabaseVerifyFlags = 0
)

type TlsPasswordFlags C.GTlsPasswordFlags

const (
	TlsPasswordFlagsNone      TlsPasswordFlags = 0
	TlsPasswordFlagsRetry                      = 2
	TlsPasswordFlagsManyTries                  = 4
	TlsPasswordFlagsFinalTry                   = 8
)
