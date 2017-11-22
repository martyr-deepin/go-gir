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
#include <stdio.h>
#include <strings.h>

static void destroy_closure(gpointer data) {
	g_closure_unref( (GClosure*)(data) );
}

// CallbackWapper?
static void callback_handler(gpointer instance, gpointer user_data) {
	GClosure* closure = user_data;
	g_closure_invoke(closure, NULL, 0, NULL, NULL);
}

static gulong _g_cancellable_connect(GCancellable *cancellable, GClosure* closure) {
	return g_cancellable_connect(cancellable, (GCallback)(callback_handler), closure, destroy_closure);
}
*/
import "C"
import (
	"github.com/linuxdeepin/go-gir/glib-2.0"
	"github.com/linuxdeepin/go-gir/gobject-2.0"
)

func (cancellable Cancellable) Connect(callback func()) uint {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr)
	ret0 := C._g_cancellable_connect(cancellable.native(), callback0)
	return uint(ret0)
}

//g_file_info_get_modification_time
func (info FileInfo) GetModificationTime() glib.TimeVal {
	result := glib.TimeValNew()
	C.g_file_info_get_modification_time(info.native(), (*C.GTimeVal)(result.Ptr))
	return result
}

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
