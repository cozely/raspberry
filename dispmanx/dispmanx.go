package dispmanx

import (
	"unsafe"
)

/*
#include "bcm_host.h"
*/
import "C"

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
