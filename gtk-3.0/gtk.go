package gtk


/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Init0() {
	C.gtk_init(nil, nil)
}

// SetText is a wrapper around gtk_text_buffer_set_text().
func (buffer TextBuffer) SetText(text string) {
	text0 := (*C.gchar)(C.CString(text))
	C.gtk_text_buffer_set_text(buffer.native(), text0, C.gint(len(text)))
	C.free(unsafe.Pointer(text0))
}
