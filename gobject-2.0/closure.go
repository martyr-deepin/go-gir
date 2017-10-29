package gobject

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <glib-object.h>
#include <stdlib.h>
extern void	goMarshal(GClosure *, GValue *, guint, GValue *, gpointer, GValue *);

static GClosure *
_g_closure_new() {
	GClosure	*closure;

	closure = g_closure_new_simple(sizeof(GClosure), NULL);
	g_closure_set_marshal(closure, (GClosureMarshal)(goMarshal));
	return closure;
}

extern void	removeClosure(gpointer, GClosure *);

static void
_g_closure_add_finalize_notifier(GClosure *closure) {
	g_closure_add_finalize_notifier(closure, NULL, removeClosure);
}

*/
import "C"
import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"
	"unsafe"

	"github.com/electricface/go-auto-gir/util"
	"github.com/electricface/go-auto-gir/glib-2.0"
)

type closureContext struct {
	rf reflect.Value
}

var (
	errNilPtr = errors.New("cgo returned unexpected nil pointer")

	closures = struct {
		sync.RWMutex
		m map[*C.GClosure]closureContext
	}{
		m: make(map[*C.GClosure]closureContext),
	}

	signals = make(map[SignalHandle]Closure)
)

// removeClosure removes a closure from the internal closures map.  This is
// needed to prevent a leak where Go code can access the closure context
// (along with rf and userdata) even after an object has been destroyed and
// the GClosure is invalidated and will never run.
//
//export removeClosure
func removeClosure(_ C.gpointer, closure *C.GClosure) {
	closures.Lock()
	delete(closures.m, closure)
	closures.Unlock()
}

// goMarshal is called by the GLib runtime when a closure needs to be invoked.
// The closure will be invoked with as many arguments as it can take, from 0 to
// the full amount provided by the call. If the closure asks for more parameters
// than there are to give, a warning is printed to stderr and the closure is
// not run.
//
//export goMarshal
func goMarshal(closure *C.GClosure, retValue *C.GValue,
	nParams C.guint, params *C.GValue,
	invocationHint C.gpointer, marshalData *C.GValue) {

	// Get the context associated with this callback closure.
	closures.RLock()
	cc := closures.m[closure]
	closures.RUnlock()

	// Get number of parameters passed in.  If user data was saved with the
	// closure context, increment the total number of parameters.
	nGLibParams := int(nParams)
	nTotalParams := nGLibParams
	println("nGLibParams: ", nGLibParams)
	println("nTotalParams: ", nTotalParams)

	// Get number of parameters from the callback closure.  If this exceeds
	// the total number of marshaled parameters, a warning will be printed
	// to stderr, and the callback will not be run.
	nCbParams := cc.rf.Type().NumIn()
	println("nCbParams:", nCbParams)
	if nCbParams > nTotalParams {
		fmt.Fprintf(os.Stderr,
			"too many closure args: have %d, max allowed %d\n",
			nCbParams, nTotalParams)
		return
	}

	// Create a slice of reflect.Values as arguments to call the function.
	gValues := gValueSlice(params, nCbParams)
	args := make([]reflect.Value, 0, nCbParams)

	// Fill beginning of args, up to the minimum of the total number of callback
	// parameters and parameters from the glib runtime.
	for i := 0; i < nCbParams && i < nGLibParams; i++ {
		v := wrapValue(&gValues[i])
		val, err := v.GoValue()
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"no suitable Go value for arg %d: %v\n", i, err)
			return
		}
		fmt.Printf("i: %d, val: %#v\n", i, val)
		rv := reflect.ValueOf(val)
		args = append(args, rv.Convert(cc.rf.Type().In(i)))
	}

	// Call closure with args. If the callback returns one or more
	// values, save the GValue equivalent of the first.
	rv := cc.rf.Call(args)
	if retValue != nil && len(rv) > 0 {
		// set gRetValue to the return value of go function
		gRetValue := wrapValue(retValue)
		err := gRetValue.Set(rv[0].Interface())
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot save callback return value: %v", err)
		}
	}
}

// gValueSlice converts a C array of GValues to a Go slice.
func gValueSlice(values *C.GValue, nValues int) (slice []C.GValue) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Cap = nValues
	header.Len = nValues
	header.Data = uintptr(unsafe.Pointer(values))
	return
}

type CanGetTypeAndGValueMarshaler interface {
	GetType() Type
	GetGValueMarshaler() GValueMarshaler
}

var cgtm CanGetTypeAndGValueMarshaler
var cgtmType = reflect.TypeOf(&cgtm).Elem()

// ClosureNew creates a new GClosure and adds its callback function
// to the internally-maintained map. It's exported for visibility to other
// gotk3 packages and shouldn't be used in application code.
func ClosureNew(f interface{}) Closure {
	// Create a reflect.Value from f.  This is called when the
	// returned GClosure runs.
	rf := reflect.ValueOf(f)

	// Create closure context which points to the reflected func.
	cc := closureContext{rf: rf}

	// Closures can only be created from funcs.
	if rf.Type().Kind() != reflect.Func {
		panic("f is not a func")
	}

	fType := rf.Type()
	for i := 0; i < fType.NumIn(); i++ {
		argType := fType.In(i)
		fmt.Printf("i: %d, argType: %v\n", i, argType)
		if argType.Implements(cgtmType) {
			fmt.Println("it impl cctm")
			argEmptyValue := reflect.New(argType)
			cctm := argEmptyValue.Interface().(CanGetTypeAndGValueMarshaler)
			RegisterGValueMarshaler(cctm.GetType(), cctm.GetGValueMarshaler())
		}
	}

	c := C._g_closure_new()
	C._g_closure_add_finalize_notifier(c)

	// Associate the GClosure with rf.  rf will be looked up in this
	// map by the closure when the closure runs.
	closures.Lock()
	closures.m[c] = cc
	closures.Unlock()

	return wrapClosure(c)
}

type SignalHandle uint

func (v Object) connectClosure(after bool, detailedSignal string, f interface{}) SignalHandle {
	cstr := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(cstr))

	closure := ClosureNew(f)

	c := C.g_signal_connect_closure(C.gpointer(v.native()),
		(*C.gchar)(cstr), closure.native(), C.gboolean(util.Bool2Int(after)))
	handle := SignalHandle(c)

	// Map the signal handle to the closure.
	signals[handle] = closure

	return handle
}

func SourceSetClosure(src glib.Source, closure Closure) {
	C.g_source_set_closure( (*C.GSource)(src.Ptr), closure.native())
}

type SourceFunc func() bool

func IdleAdd(f SourceFunc) uint {
	src := glib.IdleSourceNew()
	id := setupSourceFunc(src, f)
	src.Unref()
	return id
}

func TimeoutAdd(interval uint, f SourceFunc) uint {
	src := glib.TimeoutSourceNew(interval)
	id := setupSourceFunc(src, f)
	src.Unref()
	return id
}

func TimeoutAddSeconds(interval uint, f SourceFunc) uint {
	src := glib.TimeoutSourceNewSeconds(interval)
	id := setupSourceFunc(src, f)
	src.Unref()
	return id
}

func setupSourceFunc(src glib.Source, f SourceFunc) uint {
	closure := ClosureNew(f)
	SourceSetClosure(src,closure)
	return src.Attach(glib.MainContextDefault())
}
