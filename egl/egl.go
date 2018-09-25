package egl

import (
	"reflect"
	"unsafe"
)

/*
#include "EGL/egl.h"
#include "EGL/eglext.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

type eglerror struct {
	msg string
}

func (e eglerror) Error() string {
	return e.msg
}

////////////////////////////////////////////////////////////////////////////////

// GetDisplay returns an EGL display connection.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglGetDisplay.xhtml
func GetDisplay(nd NativeDisplay) (Display, error) {
	d := C.eglGetDisplay(C.EGLNativeDisplayType(nd))
	if d == C.EGLDisplay(C.EGL_NO_DISPLAY) {
		return NO_DISPLAY, eglerror{msg: "unable to get EGL display"}
	}
	return Display(d), nil
}

// Initialize initializes an EGL display connection.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglInitialize.xhtml
func Initialize(d Display) (maj, min int32, err error) {
	r := C.eglInitialize(C.EGLDisplay(d), (*C.int)(&maj), (*C.int)(&min))
	if r == C.EGL_FALSE {
		return maj, min, eglerror{msg: "unable to initialize EGL display"}
	}
	return maj, min, nil
}

// ChooseConfig returns a list of EGL frame buffer configurations that match
// specified attributes.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglChooseConfig.xhtml
func ChooseConfig(d Display, a []Int) ([]Config, error) {
	if len(a) == 0 || a[len(a)-1] != NONE {
		a = append(a, NONE)
	}
	if len(a)%2 == 0 {
		a = append(a, NONE)
	}
	var nb C.EGLint
	r := C.eglChooseConfig(C.EGLDisplay(d), (*C.EGLint)(&a[0]),
		nil, 0, (*C.EGLint)(&nb))
	if r == C.EGL_FALSE {
		return []Config{}, eglerror{msg: "unable to find EGL framebuffer configuration"}
	}
	c := make([]Config, nb, nb)
	r = C.eglChooseConfig(C.EGLDisplay(d), (*C.EGLint)(&a[0]),
		(*C.EGLConfig)(&c[0]), nb, (*C.EGLint)(&nb))
	if r == C.EGL_FALSE {
		return []Config{}, eglerror{msg: "unable to get EGL framebuffer configuration"}
	}
	return c, nil
}

// BindAPI sets the current rendering API
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglBindAPI.xhtml
func BindAPI(api Int) error {
	r := C.eglBindAPI(C.EGLenum(api))
	if r == C.EGL_FALSE {
		return eglerror{msg: "unable to bind OpenGL ES API"}
	}
	return nil
}

// CreateContext creates a new EGL rendering context.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglCreateContext.xhtml
func CreateContext(d Display, c Config, share Context, a []Int) (Context, error) {
	if len(a) == 0 || a[len(a)-1] != NONE {
		a = append(a, NONE)
	}
	if len(a)%2 == 0 {
		a = append(a, NONE)
	}
	ctx := C.eglCreateContext(C.EGLDisplay(d), C.EGLConfig(c),
		C.EGLContext(share), (*C.EGLint)(&a[0]))
	if ctx == C.EGLContext(C.EGL_NO_CONTEXT) {
		return NO_CONTEXT, eglerror{msg: "unable to create EGL context"}
	}
	return Context(ctx), nil
}

// CreateWindowSurface creates a new EGL window surface.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglCreateWindowSurface.xhtml
func CreateWindowSurface(d Display, c Config, w NativeWindow, a []Int) (Surface, error) {
	if len(a) == 0 || a[len(a)-1] != NONE {
		a = append(a, NONE)
	}
	if len(a)%2 == 0 {
		a = append(a, NONE)
	}
	ww := unsafe.Pointer(reflect.ValueOf(w).Pointer()) //TODO: check if pointer
	s := C.eglCreateWindowSurface(C.EGLDisplay(d), C.EGLConfig(c),
		C.NativeWindowType(ww), (*C.EGLint)(&a[0]))
	if s == C.EGLSurface(C.EGL_NO_SURFACE) {
		return NO_SURFACE, eglerror{msg: "unable to create EGL window surface"}
	}
	return Surface(s), nil
}

// MakeCurrent attaches an EGL rendering context to EGL surfaces.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglMakeCurrent.xhtml
func MakeCurrent(d Display, draw Surface, read Surface, c Context) error {
	r := C.eglMakeCurrent(C.EGLDisplay(d), C.EGLSurface(draw), C.EGLSurface(read), C.EGLContext(c))
	if r == C.EGL_FALSE {
		return eglerror{msg: "unable to make EGL context current"}
	}
	return nil
}

// SwapBuffers posts EGL surface color buffer to a native window.
//
// https://www.khronos.org/registry/EGL/sdk/docs/man/html/eglSwapBuffers.xhtml
func SwapBuffers(d Display, s Surface) error {
	r := C.eglSwapBuffers(C.EGLDisplay(d), C.EGLSurface(s))
	if r == C.EGL_FALSE {
		return eglerror{msg: "unabel to swap EGL buffers"}
	}
	return nil
}
