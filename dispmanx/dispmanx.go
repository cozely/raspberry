package dispmanx

import (
	"unsafe"
)

/*
#include "bcm_host.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// A Device ID
type Device uint32

type DisplayHandle uint32
type UpdateHandle uint32
type ElementHandle uint32
type ResourceHandle uint32

const NoHandle = 0

type Protection uint32

const (
	ProtectionMax  Protection = 0x0F
	ProtectionNone Protection = 0
	ProtectionHDCP Protection = 11
)

type Rect struct {
	X, Y          int32
	Width, Height int32
}

type Transform uint32

const (
	// Bottom 2 bits sets the orientation

	NoRotate  Transform = 0
	Rotate90  Transform = 1
	Rotate180 Transform = 2
	Rotate270 Transform = 3

	FlipHriz Transform = 1 << 16
	FlipVert Transform = 1 << 17

	// extra flags for controlling 3d duplication behaviour

	StereoscopicInvert Transform = 1 << 19 // invert left/right images
	StereoscopicNone   Transform = 0 << 20
	StereoscopicMono   Transform = 1 << 20
	StereoscopicSBS    Transform = 2 << 20
	StereoscopicTB     Transform = 3 << 20
	StereoscopicMask   Transform = 15 << 20

	// extra flags for controlling snapshot behaviour

	SnapshotNoYUV       Transform = 1 << 24
	SnapshotNoRGB       Transform = 1 << 25
	SnapshotFill        Transform = 1 << 26
	SnapshotSwapRedBlue Transform = 1 << 27
	SnapshotPack        Transform = 1 << 28
)

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

// DisplayOpen opens a display on the given device.
func DisplayOpen(d Device) DisplayHandle {
	return DisplayHandle(C.vc_dispmanx_display_open(C.uint32_t(d)))
}

// UpdateStart starts a new update, returns NoHandle on error.
func UpdateStart(priority int32) UpdateHandle {
	return UpdateHandle(C.vc_dispmanx_update_start(C.int32_t(priority)))
}

// ElementAdd adds an elment to a display as part of an update.
func ElementAdd(
	u UpdateHandle, d DisplayHandle,
	layer int32, dstRect Rect,
	src ResourceHandle, srcRect Rect,
	p Protection,
	a *Alpha, // Must be nil
	c *Clamp, // Must be nil
	t Transform,
) ElementHandle {
	return ElementHandle(C.vc_dispmanx_element_add(
		C.DISPMANX_UPDATE_HANDLE_T(u), C.DISPMANX_DISPLAY_HANDLE_T(d),
		C.int32_t(layer), (*C.VC_RECT_T)(unsafe.Pointer(&dstRect)),
		C.DISPMANX_RESOURCE_HANDLE_T(src), (*C.VC_RECT_T)(unsafe.Pointer(&srcRect)),
		C.DISPMANX_PROTECTION_T(p),
		(*C.VC_DISPMANX_ALPHA_T)(unsafe.Pointer(a)),
		(*C.DISPMANX_CLAMP_T)(unsafe.Pointer(c)),
		C.DISPMANX_TRANSFORM_T(t),
	))
}

// UpdateSubmitSync ends an update and wait for it to complete.
func UpdateSubmitSync(u UpdateHandle) {
	C.vc_dispmanx_update_submit_sync(C.DISPMANX_UPDATE_HANDLE_T(u))
}
