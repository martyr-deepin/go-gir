package gudev

/*
#cgo pkg-config: gudev-1.0
#include <gudev/gudev.h>
#include <stdlib.h>
*/
import "C"
import "github.com/electricface/go-auto-gir/glib-2.0"
import "github.com/electricface/go-auto-gir/gobject-2.0"
import "github.com/electricface/go-auto-gir/util"
import "unsafe"

// Object Client
type Client struct {
	gobject.Object
}

func (v Client) native() *C.GUdevClient {
	return (*C.GUdevClient)(v.Ptr)
}
func wrapClient(p *C.GUdevClient) (v Client) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapClient(p unsafe.Pointer) (v Client) {
	v.Ptr = p
	return
}

// ClientNew is a wrapper around g_udev_client_new().
func ClientNew(subsystems []string) Client {
	subsystems0 := make([]*C.gchar, len(subsystems))
	for idx, elemG := range subsystems {
		elem := (*C.gchar)(C.CString(elemG))
		subsystems0[idx] = elem
	}
	var subsystems0Ptr **C.gchar
	if len(subsystems0) > 0 {
		subsystems0Ptr = &subsystems0[0]
	}
	ret0 := C.g_udev_client_new(subsystems0Ptr)
	for _, elem := range subsystems0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return wrapClient(ret0)
}

// QueryByDeviceFile is a wrapper around g_udev_client_query_by_device_file().
func (client Client) QueryByDeviceFile(device_file string) Device {
	device_file0 := (*C.gchar)(C.CString(device_file))
	ret0 := C.g_udev_client_query_by_device_file(client.native(), device_file0)
	C.free(unsafe.Pointer(device_file0)) /*ch:<stdlib.h>*/
	return wrapDevice(ret0)
}

// QueryByDeviceNumber is a wrapper around g_udev_client_query_by_device_number().
func (client Client) QueryByDeviceNumber(type_ DeviceType, number DeviceNumber) Device {
	ret0 := C.g_udev_client_query_by_device_number(client.native(), C.GUdevDeviceType(type_), C.GUdevDeviceNumber(number))
	return wrapDevice(ret0)
}

// QueryBySubsystem is a wrapper around g_udev_client_query_by_subsystem().
func (client Client) QueryBySubsystem(subsystem string) glib.List {
	subsystem0 := (*C.gchar)(C.CString(subsystem))
	ret0 := C.g_udev_client_query_by_subsystem(client.native(), subsystem0)
	C.free(unsafe.Pointer(subsystem0)) /*ch:<stdlib.h>*/
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDevice(p) }) /*gir:GLib*/
}

// QueryBySubsystemAndName is a wrapper around g_udev_client_query_by_subsystem_and_name().
func (client Client) QueryBySubsystemAndName(subsystem string, name string) Device {
	subsystem0 := (*C.gchar)(C.CString(subsystem))
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_client_query_by_subsystem_and_name(client.native(), subsystem0, name0)
	C.free(unsafe.Pointer(subsystem0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(name0))      /*ch:<stdlib.h>*/
	return wrapDevice(ret0)
}

// QueryBySysfsPath is a wrapper around g_udev_client_query_by_sysfs_path().
func (client Client) QueryBySysfsPath(sysfs_path string) Device {
	sysfs_path0 := (*C.gchar)(C.CString(sysfs_path))
	ret0 := C.g_udev_client_query_by_sysfs_path(client.native(), sysfs_path0)
	C.free(unsafe.Pointer(sysfs_path0)) /*ch:<stdlib.h>*/
	return wrapDevice(ret0)
}

// Object Device
type Device struct {
	gobject.Object
}

func (v Device) native() *C.GUdevDevice {
	return (*C.GUdevDevice)(v.Ptr)
}
func wrapDevice(p *C.GUdevDevice) (v Device) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDevice(p unsafe.Pointer) (v Device) {
	v.Ptr = p
	return
}

// GetAction is a wrapper around g_udev_device_get_action().
func (device Device) GetAction() string {
	ret0 := C.g_udev_device_get_action(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDeviceFile is a wrapper around g_udev_device_get_device_file().
func (device Device) GetDeviceFile() string {
	ret0 := C.g_udev_device_get_device_file(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDeviceFileSymlinks is a wrapper around g_udev_device_get_device_file_symlinks().
func (device Device) GetDeviceFileSymlinks() []string {
	ret0 := C.g_udev_device_get_device_file_symlinks(device.native())
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

// GetDeviceNumber is a wrapper around g_udev_device_get_device_number().
func (device Device) GetDeviceNumber() DeviceNumber {
	ret0 := C.g_udev_device_get_device_number(device.native())
	return DeviceNumber(ret0)
}

// GetDeviceType is a wrapper around g_udev_device_get_device_type().
func (device Device) GetDeviceType() DeviceType {
	ret0 := C.g_udev_device_get_device_type(device.native())
	return DeviceType(ret0)
}

// GetDevtype is a wrapper around g_udev_device_get_devtype().
func (device Device) GetDevtype() string {
	ret0 := C.g_udev_device_get_devtype(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDriver is a wrapper around g_udev_device_get_driver().
func (device Device) GetDriver() string {
	ret0 := C.g_udev_device_get_driver(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetIsInitialized is a wrapper around g_udev_device_get_is_initialized().
func (device Device) GetIsInitialized() bool {
	ret0 := C.g_udev_device_get_is_initialized(device.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetName is a wrapper around g_udev_device_get_name().
func (device Device) GetName() string {
	ret0 := C.g_udev_device_get_name(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNumber is a wrapper around g_udev_device_get_number().
func (device Device) GetNumber() string {
	ret0 := C.g_udev_device_get_number(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetParent is a wrapper around g_udev_device_get_parent().
func (device Device) GetParent() Device {
	ret0 := C.g_udev_device_get_parent(device.native())
	return wrapDevice(ret0)
}

// GetParentWithSubsystem is a wrapper around g_udev_device_get_parent_with_subsystem().
func (device Device) GetParentWithSubsystem(subsystem string, devtype string) Device {
	subsystem0 := (*C.gchar)(C.CString(subsystem))
	devtype0 := (*C.gchar)(C.CString(devtype))
	ret0 := C.g_udev_device_get_parent_with_subsystem(device.native(), subsystem0, devtype0)
	C.free(unsafe.Pointer(subsystem0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(devtype0))   /*ch:<stdlib.h>*/
	return wrapDevice(ret0)
}

// GetProperty is a wrapper around g_udev_device_get_property().
func (device Device) GetProperty(key string) string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_get_property(device.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPropertyAsBoolean is a wrapper around g_udev_device_get_property_as_boolean().
func (device Device) GetPropertyAsBoolean(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_get_property_as_boolean(device.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPropertyAsDouble is a wrapper around g_udev_device_get_property_as_double().
func (device Device) GetPropertyAsDouble(key string) float64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_get_property_as_double(device.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return float64(ret0)
}

// GetPropertyAsInt is a wrapper around g_udev_device_get_property_as_int().
func (device Device) GetPropertyAsInt(key string) int {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_get_property_as_int(device.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetPropertyAsStrv is a wrapper around g_udev_device_get_property_as_strv().
func (device Device) GetPropertyAsStrv(key string) []string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_get_property_as_strv(device.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
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

// GetPropertyAsUint64 is a wrapper around g_udev_device_get_property_as_uint64().
func (device Device) GetPropertyAsUint64(key string) uint64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_get_property_as_uint64(device.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return uint64(ret0)
}

// GetPropertyKeys is a wrapper around g_udev_device_get_property_keys().
func (device Device) GetPropertyKeys() []string {
	ret0 := C.g_udev_device_get_property_keys(device.native())
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

// GetSeqnum is a wrapper around g_udev_device_get_seqnum().
func (device Device) GetSeqnum() uint64 {
	ret0 := C.g_udev_device_get_seqnum(device.native())
	return uint64(ret0)
}

// GetSubsystem is a wrapper around g_udev_device_get_subsystem().
func (device Device) GetSubsystem() string {
	ret0 := C.g_udev_device_get_subsystem(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSysfsAttr is a wrapper around g_udev_device_get_sysfs_attr().
func (device Device) GetSysfsAttr(name string) string {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_device_get_sysfs_attr(device.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSysfsAttrAsBoolean is a wrapper around g_udev_device_get_sysfs_attr_as_boolean().
func (device Device) GetSysfsAttrAsBoolean(name string) bool {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_device_get_sysfs_attr_as_boolean(device.native(), name0)
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSysfsAttrAsDouble is a wrapper around g_udev_device_get_sysfs_attr_as_double().
func (device Device) GetSysfsAttrAsDouble(name string) float64 {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_device_get_sysfs_attr_as_double(device.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return float64(ret0)
}

// GetSysfsAttrAsInt is a wrapper around g_udev_device_get_sysfs_attr_as_int().
func (device Device) GetSysfsAttrAsInt(name string) int {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_device_get_sysfs_attr_as_int(device.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetSysfsAttrAsStrv is a wrapper around g_udev_device_get_sysfs_attr_as_strv().
func (device Device) GetSysfsAttrAsStrv(name string) []string {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_device_get_sysfs_attr_as_strv(device.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
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

// GetSysfsAttrAsUint64 is a wrapper around g_udev_device_get_sysfs_attr_as_uint64().
func (device Device) GetSysfsAttrAsUint64(name string) uint64 {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_device_get_sysfs_attr_as_uint64(device.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return uint64(ret0)
}

// GetSysfsAttrKeys is a wrapper around g_udev_device_get_sysfs_attr_keys().
func (device Device) GetSysfsAttrKeys() []string {
	ret0 := C.g_udev_device_get_sysfs_attr_keys(device.native())
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

// GetSysfsPath is a wrapper around g_udev_device_get_sysfs_path().
func (device Device) GetSysfsPath() string {
	ret0 := C.g_udev_device_get_sysfs_path(device.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetTags is a wrapper around g_udev_device_get_tags().
func (device Device) GetTags() []string {
	ret0 := C.g_udev_device_get_tags(device.native())
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

// GetUsecSinceInitialized is a wrapper around g_udev_device_get_usec_since_initialized().
func (device Device) GetUsecSinceInitialized() uint64 {
	ret0 := C.g_udev_device_get_usec_since_initialized(device.native())
	return uint64(ret0)
}

// HasProperty is a wrapper around g_udev_device_has_property().
func (device Device) HasProperty(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_has_property(device.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasSysfsAttr is a wrapper around g_udev_device_has_sysfs_attr().
func (device Device) HasSysfsAttr(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_udev_device_has_sysfs_attr(device.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object Enumerator
type Enumerator struct {
	gobject.Object
}

func (v Enumerator) native() *C.GUdevEnumerator {
	return (*C.GUdevEnumerator)(v.Ptr)
}
func wrapEnumerator(p *C.GUdevEnumerator) (v Enumerator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEnumerator(p unsafe.Pointer) (v Enumerator) {
	v.Ptr = p
	return
}

// EnumeratorNew is a wrapper around g_udev_enumerator_new().
func EnumeratorNew(client Client) Enumerator {
	ret0 := C.g_udev_enumerator_new(client.native())
	return wrapEnumerator(ret0)
}

// AddMatchIsInitialized is a wrapper around g_udev_enumerator_add_match_is_initialized().
func (enumerator Enumerator) AddMatchIsInitialized() Enumerator {
	ret0 := C.g_udev_enumerator_add_match_is_initialized(enumerator.native())
	return wrapEnumerator(ret0)
}

// AddMatchName is a wrapper around g_udev_enumerator_add_match_name().
func (enumerator Enumerator) AddMatchName(name string) Enumerator {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_udev_enumerator_add_match_name(enumerator.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// AddMatchProperty is a wrapper around g_udev_enumerator_add_match_property().
func (enumerator Enumerator) AddMatchProperty(name string, value string) Enumerator {
	name0 := (*C.gchar)(C.CString(name))
	value0 := (*C.gchar)(C.CString(value))
	ret0 := C.g_udev_enumerator_add_match_property(enumerator.native(), name0, value0)
	C.free(unsafe.Pointer(name0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// AddMatchSubsystem is a wrapper around g_udev_enumerator_add_match_subsystem().
func (enumerator Enumerator) AddMatchSubsystem(subsystem string) Enumerator {
	subsystem0 := (*C.gchar)(C.CString(subsystem))
	ret0 := C.g_udev_enumerator_add_match_subsystem(enumerator.native(), subsystem0)
	C.free(unsafe.Pointer(subsystem0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// AddMatchSysfsAttr is a wrapper around g_udev_enumerator_add_match_sysfs_attr().
func (enumerator Enumerator) AddMatchSysfsAttr(name string, value string) Enumerator {
	name0 := (*C.gchar)(C.CString(name))
	value0 := (*C.gchar)(C.CString(value))
	ret0 := C.g_udev_enumerator_add_match_sysfs_attr(enumerator.native(), name0, value0)
	C.free(unsafe.Pointer(name0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// AddMatchTag is a wrapper around g_udev_enumerator_add_match_tag().
func (enumerator Enumerator) AddMatchTag(tag string) Enumerator {
	tag0 := (*C.gchar)(C.CString(tag))
	ret0 := C.g_udev_enumerator_add_match_tag(enumerator.native(), tag0)
	C.free(unsafe.Pointer(tag0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// AddNomatchSubsystem is a wrapper around g_udev_enumerator_add_nomatch_subsystem().
func (enumerator Enumerator) AddNomatchSubsystem(subsystem string) Enumerator {
	subsystem0 := (*C.gchar)(C.CString(subsystem))
	ret0 := C.g_udev_enumerator_add_nomatch_subsystem(enumerator.native(), subsystem0)
	C.free(unsafe.Pointer(subsystem0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// AddNomatchSysfsAttr is a wrapper around g_udev_enumerator_add_nomatch_sysfs_attr().
func (enumerator Enumerator) AddNomatchSysfsAttr(name string, value string) Enumerator {
	name0 := (*C.gchar)(C.CString(name))
	value0 := (*C.gchar)(C.CString(value))
	ret0 := C.g_udev_enumerator_add_nomatch_sysfs_attr(enumerator.native(), name0, value0)
	C.free(unsafe.Pointer(name0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// AddSysfsPath is a wrapper around g_udev_enumerator_add_sysfs_path().
func (enumerator Enumerator) AddSysfsPath(sysfs_path string) Enumerator {
	sysfs_path0 := (*C.gchar)(C.CString(sysfs_path))
	ret0 := C.g_udev_enumerator_add_sysfs_path(enumerator.native(), sysfs_path0)
	C.free(unsafe.Pointer(sysfs_path0)) /*ch:<stdlib.h>*/
	return wrapEnumerator(ret0)
}

// Execute is a wrapper around g_udev_enumerator_execute().
func (enumerator Enumerator) Execute() glib.List {
	ret0 := C.g_udev_enumerator_execute(enumerator.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDevice(p) }) /*gir:GLib*/
}

type DeviceNumber uint64
type DeviceType int

const (
	DeviceTypeNone  DeviceType = 0
	DeviceTypeBlock            = 98
	DeviceTypeChar             = 99
)
