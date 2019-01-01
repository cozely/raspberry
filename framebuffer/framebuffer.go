package framebuffer

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

//------------------------------------------------------------------------------

var (
	originfo fb_var_screeninfo
	varinfo  fb_var_screeninfo
	fixinfo  fb_fix_screeninfo
)

//------------------------------------------------------------------------------

func Setup() error {
	f, err := unix.Open("/dev/fb0", unix.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer unix.Close(f)

	// First, get the current state of the framebuffer
	err = ioctl(f, _FBIOGET_VSCREENINFO, unsafe.Pointer(&originfo))
	if err != nil {
		panic(err) //TODO
	}
	fmt.Printf("Resolution: %dx%d(%d), virtual: %dx%d +%d+%d\n",
		originfo.xres, originfo.yres,
		originfo.bits_per_pixel,
		originfo.xres_virtual, originfo.yres_virtual,
		originfo.xoffset, originfo.yoffset)

	// Modify it to fit our needs
	varinfo = originfo
	varinfo.xres_virtual = varinfo.xres
	varinfo.yres_virtual = 2 * varinfo.yres
	varinfo.xoffset = 0
	varinfo.yoffset = 0
	err = ioctl(f, _FBIOPUT_VSCREENINFO, unsafe.Pointer(&varinfo))
	if err != nil {
		panic(err) //TODO
	}
	err = ioctl(f, _FBIOGET_VSCREENINFO, unsafe.Pointer(&varinfo))
	if err != nil {
		panic(err) //TODO
	}
	fmt.Printf("Resolution: %dx%d(%d), virtual: %dx%d +%d+%d\n",
		varinfo.xres, varinfo.yres,
		varinfo.bits_per_pixel,
		varinfo.xres_virtual, varinfo.yres_virtual,
		varinfo.xoffset, varinfo.yoffset)

	// Findout framebuffer size and stride
	err = ioctl(f, _FBIOGET_FSCREENINFO, unsafe.Pointer(&fixinfo))
	if err != nil {
		panic(err) //TODO
	}
	fmt.Printf("Framebuffer ID: %s\n", string(fixinfo.id[:]))
	fmt.Printf("fscreeninfo.smem_len: %d\n", fixinfo.smem_len)
	fmt.Printf("fscreeninfo.line_length: %d\n", fixinfo.line_length)

	varinfo.yoffset = varinfo.yres
	err = ioctl(f, _FBIOPAN_DISPLAY, unsafe.Pointer(&varinfo))
	if err != nil {
		panic(err) //TODO
	}
	err = ioctl(f, _FBIOGET_VSCREENINFO, unsafe.Pointer(&varinfo))
	if err != nil {
		panic(err) //TODO
	}
	fmt.Printf("Resolution: %dx%d(%d), virtual: %dx%d +%d+%d\n",
		varinfo.xres, varinfo.yres,
		varinfo.bits_per_pixel,
		varinfo.xres_virtual, varinfo.yres_virtual,
		varinfo.xoffset, varinfo.yoffset)

	// m, err := unix.Mmap(
	// 	f,
	// 	0,
	// 	int(unsafe.Sizeof(*Registers)),
	// 	unix.PROT_READ|unix.PROT_WRITE,
	// 	unix.MAP_SHARED,
	// )
	// Registers = (*[41]uint32)(unsafe.Pointer(&m[0]))

	return nil
}

//------------------------------------------------------------------------------

func Cleanup() error {
	f, err := unix.Open("/dev/fb0", unix.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer unix.Close(f)

	err = ioctl(f, _FBIOPUT_VSCREENINFO, unsafe.Pointer(&originfo))
	if err != nil {
		return fmt.Errorf("Unable to clean up framebuffer state: %v", err)
	}
	err = ioctl(f, _FBIOGET_VSCREENINFO, unsafe.Pointer(&varinfo))
	if err != nil {
		panic(err) //TODO
	}
	fmt.Printf("Resolution: %dx%d(%d), virtual: %dx%d +%d+%d\n",
		varinfo.xres, varinfo.yres,
		varinfo.bits_per_pixel,
		varinfo.xres_virtual, varinfo.yres_virtual,
		varinfo.xoffset, varinfo.yoffset)

	err = ioctl(f, _FBIOGET_FSCREENINFO, unsafe.Pointer(&fixinfo))
	if err != nil {
		panic(err) //TODO
	}
	fmt.Printf("Framebuffer ID: %s\n", string(fixinfo.id[:]))
	fmt.Printf("fscreeninfo.smem_start: %d\n", fixinfo.smem_start)
	fmt.Printf("fscreeninfo.smem_len: %d\n", fixinfo.smem_len)
	fmt.Printf("fscreeninfo.line_length: %d\n", fixinfo.line_length)

	return nil
}
