package dispmanx

import "unsafe"

/*
#include "bcm_host.h"
*/

////////////////////////////////////////////////////////////////////////////////

// A Device ID
type Device uint32

type DisplayHandle uint32
type UpdateHandle uint32
type ElementHandle uint32
type ResourceHandle uint32

type Protection uint32

type Rect struct {
	X, Y          int32
	Width, Height int32
}

type Alpha struct {
	flags   uint32
	opacity uint32
	mask    unsafe.Pointer
}

type Clamp struct {
	mode         uint32
	keyMask      uint32
	keys         [6]uint8
	replaveValue uint32
}

////////////////////////////////////////////////////////////////////////////////

type Window struct {
	Element       ElementHandle
	Width, Height int32
}

func (*Window) IsEGLNativeWindow() {}
