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
import (
//"github.com/electricface/go-auto-gir/util"
//"unsafe"
)

//import "github.com/electricface/go-auto-gir/gobject-2.0"
//import "github.com/electricface/go-auto-gir/util"
//import "github.com/electricface/go-auto-gir/glib-2.0"

//func (app Application) Run(argv []string) int {
//	argv0 := make([]*C.char, len(argv))
//	for i, arg := range argv {
//		arg0 := C.CString(arg)
//		defer C.free(unsafe.Pointer(arg0))
//		argv0[i] = arg0
//	}
//
//	var argv0Ptr **C.char
//	if len(argv0) > 0 {
//		argv0Ptr = &argv0[0]
//	}
//
//	ret0 := C.g_application_run(app.native(), C.int(len(argv)), argv0Ptr)
//	return int(ret0)
//}

//func (settings Settings) GetStrv(key string) []string {
//	key0 := (*C.gchar)(C.CString(key))
//	defer C.free(unsafe.Pointer(key0))
//	ret0 := C.g_settings_get_strv(settings.native(), key0)
//
//	// ret0 is nul term array
//	var ret0slice []*C.gchar
//	util.SetSliceDataLen(unsafe.Pointer(&ret0slice), unsafe.Pointer(ret0),
//		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))))
//	ret := make([]string, len(ret0slice))
//	for idx, varc := range ret0slice {
//		varg := C.GoString((*C.char)(varc))
//		// free element
//		defer C.g_free(C.gpointer(varc))
//		ret[idx] = varg
//	}
//
//	// free container
//	C.g_free(C.gpointer(ret0))
//	return ret
//}
