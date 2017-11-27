package gtk


/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
#include <stdlib.h>
*/
import "C"

func Init0() {
	C.gtk_init(nil, nil)
}
