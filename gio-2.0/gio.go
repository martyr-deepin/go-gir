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


static void async_ready_callback(GObject *source_object,
                        GAsyncResult *res,
                        gpointer user_data);

static void _g_file_query_info_async (GFile *file,
                         const char *attributes,
                         GFileQueryInfoFlags flags,
                         int io_priority,
                         GCancellable *cancellable,
                         GClosure* closure) {

g_file_query_info_async (file,
                         attributes,
                         flags,
                         io_priority,
                         cancellable,
                         async_ready_callback,
                         closure);
}

static void async_ready_callback(GObject *source_object,
                        GAsyncResult *res,
                        gpointer user_data) {
	GClosure* closure = user_data;
	GValue params[2];
	bzero(params, 2*sizeof(GValue));

	g_value_init(&params[0], G_TYPE_OBJECT);
	g_value_set_object(&params[0], source_object);

	//g_value_init(&params[1], G_TYPE_POINTER);
	//g_value_set_pointer(&params[1], res);
	//g_value_init_from_instance(&params[1], res);
	g_value_init(&params[1], G_TYPE_ASYNC_RESULT);
	g_value_set_object(&params[1], res);

	printf("params 0 is value: %d\n", G_IS_VALUE(&params[0]) );
	printf("params 1 is value: %d\n", G_IS_VALUE(&params[1]) );

	g_closure_invoke(closure, NULL, 2, params, NULL);
}

*/
import "C"
import (
//"github.com/electricface/go-auto-gir/util"
//"unsafe"
)

import (
	"github.com/electricface/go-auto-gir/gobject-2.0"
	"unsafe"
)

//import "github.com/electricface/go-auto-gir/util"
//import "github.com/electricface/go-auto-gir/glib-2.0"

type AsyncReadyCallback func(sourceObject gobject.Object, res AsyncResult)

func (file File) QueryInfoAsync(attributes string, flags FileQueryInfoFlags, io_priority int, cancellable Cancellable, cb AsyncReadyCallback) {
	attributes0 := C.CString(attributes)
	closure := gobject.ClosureNew(cb)
	C._g_file_query_info_async(file.native(), attributes0, C.GFileQueryInfoFlags(flags), C.int(io_priority), cancellable.native(), (*C.GClosure)(closure.Ptr))
	C.free(unsafe.Pointer(attributes0))
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
