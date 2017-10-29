package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
*/
import "C"

// IdleSourceNew is a wrapper around g_idle_source_new().
func IdleSourceNew() Source {
	ret0 := C.g_idle_source_new()
	return wrapSource(ret0)
}

// TimeoutSourceNew is a wrapper around g_timeout_source_new().
func TimeoutSourceNew(interval uint) Source {

	// Var for Go: interval
	// Var for C: interval0
	// Type for Go: uint
	// Type for C: C.guint
	ret0 := C.g_timeout_source_new(C.guint(interval))
	return wrapSource(ret0)
}

// TimeoutSourceNewSeconds is a wrapper around g_timeout_source_new_seconds().
func TimeoutSourceNewSeconds(interval uint) Source {

	// Var for Go: interval
	// Var for C: interval0
	// Type for Go: uint
	// Type for C: C.guint
	ret0 := C.g_timeout_source_new_seconds(C.guint(interval))
	return wrapSource(ret0)
}
