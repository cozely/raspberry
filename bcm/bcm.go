package bcm

/*
#include "bcm_host.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

type errmsg string

func (e errmsg) Error() string {
	return string(e)
}

////////////////////////////////////////////////////////////////////////////////

func HostInit() {
	C.bcm_host_init()
}

func HostDeinit() {
	C.bcm_host_deinit()
}

func GetDisplaySize(display uint16) (w, h int32, err error) {
	var ww, hh C.uint32_t
	e := C.graphics_get_display_size(C.uint16_t(display), &ww, &hh)
	if e != 0 {
		return 0, 0, errmsg("unable to get display size")
	}
	w, h = int32(ww), int32(hh)
	return w, h, nil
}
