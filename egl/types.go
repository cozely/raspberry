package egl

/*
#include "EGL/egl.h"
#include "EGL/eglext.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

type Int C.EGLint

type Config C.EGLConfig
type Display C.EGLDisplay
type Context C.EGLContext
type Surface C.EGLSurface
type ClientBuffer C.EGLClientBuffer

type EGLConfig C.EGLConfig

type NativeDisplay C.EGLNativeDisplayType

type NativeWindow interface {
	IsEGLNativeWindow()
}
