package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo CFLAGS: -I/opt/vc/include
#cgo LDFLAGS: -L/opt/vc/lib -lbcm_host -lbrcmEGL -lbrcmGLESv2

#include "bcm_host.h"
#include "GLES2/gl2.h"
#include "EGL/egl.h"
#include "EGL/eglext.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

var screen struct {
	width, height uint32

	display C.EGLDisplay
	surface C.EGLSurface
	context C.EGLContext
}

////////////////////////////////////////////////////////////////////////////////

func initScreen() error {
	var ec C.int // C error codes
	var e C.uint // C error bools

	C.bcm_host_init() //TODO: is there any error checking?

	ec = C.graphics_get_display_size(
		0,
		(*C.uint32_t)(unsafe.Pointer(&screen.width)),
		(*C.uint32_t)(unsafe.Pointer(&screen.height)),
	)
	if ec != 0 {
		return fmt.Errorf("unable to get display size (error code %d)", ec)
	}

	// Establish a connection with the display

	screen.display = C.eglGetDisplay(C.EGLNativeDisplayType(C.EGL_DEFAULT_DISPLAY))
	if screen.display == C.EGLDisplay(C.EGL_NO_DISPLAY) {
		return fmt.Errorf("unable to get EGL display connection")
	}
	checkgl()

	e = C.eglInitialize(screen.display, nil, nil)
	if e == C.EGL_FALSE {
		return fmt.Errorf("unable to initialize EGL display")
	}
	checkgl()

	// Create and configure an EGL context

	var conf C.EGLConfig

	afb := [...]C.EGLint{
		C.EGL_RED_SIZE, 8,
		C.EGL_GREEN_SIZE, 8,
		C.EGL_BLUE_SIZE, 8,
		C.EGL_SURFACE_TYPE, C.EGL_WINDOW_BIT,
		C.EGL_NONE,
	}
	var nc C.EGLint
	e = C.eglChooseConfig(screen.display, &afb[0], &conf, 1, &nc)
	if e == C.EGL_FALSE {
		return fmt.Errorf("unable to choose EGL framebuffer configuration")
	}
	checkgl()

	e = C.eglBindAPI(C.EGL_OPENGL_ES_API)
	if e == C.EGL_FALSE {
		return fmt.Errorf("unable to bind OpenGL ES API")
	}
	checkgl()

	a := [...]C.EGLint{
		C.EGL_CONTEXT_CLIENT_VERSION, 2,
		C.EGL_NONE,
	}
	screen.context = C.eglCreateContext(screen.display, conf,
		C.EGLContext(C.EGL_NO_CONTEXT), &a[0])
	if screen.context == C.EGLContext(C.EGL_NO_CONTEXT) {
		return fmt.Errorf("unable to create EGL context")
	}
	checkgl()

	// Create an element (i.e. layer/sprite) with DispmanX

	dpy := C.vc_dispmanx_display_open(0)
	upd := C.vc_dispmanx_update_start(0)

	src := C.VC_RECT_T{
		x: 0, y: 0,
		width:  C.int(screen.width << 16),
		height: C.int(screen.height << 16),
	}
	dst := C.VC_RECT_T{
		x: 0, y: 0,
		width:  C.int(screen.width),
		height: C.int(screen.height),
	}
	elm := C.vc_dispmanx_element_add(upd, dpy,
		0, &dst, 0, &src, C.DISPMANX_PROTECTION_NONE,
		nil, nil, 0)

	C.vc_dispmanx_update_submit_sync(upd)
	checkgl()

	// Create an EGL window surface

	w := C.EGL_DISPMANX_WINDOW_T{
		element: elm,
		width:   C.int(screen.width),
		height:  C.int(screen.height),
	}
	screen.surface = C.eglCreateWindowSurface(screen.display, conf,
		C.EGLNativeWindowType(unsafe.Pointer(&w)), nil)
	if screen.surface == C.EGLSurface(C.EGL_NO_SURFACE) {
		return fmt.Errorf("unable to create EGL window surface")
	}
	checkgl()

	e = C.eglMakeCurrent(screen.display, screen.surface, screen.surface, screen.context)
	if e == C.EGL_FALSE {
		return fmt.Errorf("unable to make EGL context current")
	}
	checkgl()

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func swapBuffers() {
	C.eglSwapBuffers(screen.display, screen.surface)
}
