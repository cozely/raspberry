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

// GetPeripheralAddress returns the ARM-side physical address where peripherals
// are mapped.
//
// This is 0x20000000 on the Pi Zero, Pi Zero W, and the first generation of the
// Raspberry Pi and Compute Module, and 0x3f000000 on the Pi 2, Pi 3 and Compute
// Module 3.
func GetPeripheralAddress() uintptr {
	return uintptr(C.bcm_host_get_peripheral_address())
}

// GetPeripheralSize returns the size of the peripheral's space.
//
// Thi is 0x01000000 for all models.
func GetPeripheralSize() uintptr {
	return uintptr(C.bcm_host_get_peripheral_size())
}

// GetSDRAMAddress returns the bus address of the SDRAM.
//
// This is 0x40000000 on the Pi Zero, Pi Zero W, and the first generation of the
// Raspberry Pi and Compute Module (GPU L2 cached), and 0xC0000000 on the Pi 2,
// Pi 3 and Compute Module 3 (uncached).
func GetSDRAMAddress() uintptr {
	return uintptr(C.bcm_host_get_sdram_address())
}

