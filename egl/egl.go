package egl

/*
#include "EGL/egl.h"
#include "EGL/eglext.h"
*/
import "C"

type Config *C.EGLConfig
type Display *C.EGLDisplay
type Context *C.EGLContext
type Surface *C.EGLSurface
type ClientBuffer *C.EGLClientBuffer

var NoDisplay = Display(nil)
var NoContext = Context(nil)
var NoSurface = Surface(nil)

type NativeDisplay *C.EGLNativeDisplayType

var DefaultDisplay = NativeDisplay(nil)

////////////////////////////////////////////////////////////////////////////////

type eglerror struct {
	msg string
}

func (e eglerror) Error() string {
	return e.msg
}

////////////////////////////////////////////////////////////////////////////////

func GetDisplay(nd NativeDisplay) (Display, error) {
	d := C.eglGetDisplay(C.EGLNativeDisplayType(nd))
	if d == C.EGL_NO_DISPLAY {
		return NoDisplay, eglerror{msg:"unable to get EGL display"}
	}
	return Display(d), nil
}

// Initialize initializes an EGL display connection.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglInitialize.xhtml
func Initialize(d Display) (maj, min int32, err error) {
	r := C.eglInitialize(C.EGLDisplay(d), (*C.int)(&maj), (*C.int)(&min))
	if r == C.EGL_FALSE {
		return maj, min, eglerror{msg:"unable to initialize EGL display"}
	}
	return maj, min, nil
}

