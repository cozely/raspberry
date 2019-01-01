package framebuffer

import (
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

//------------------------------------------------------------------------------

func ioctl(fd int, request int, argp unsafe.Pointer) error {
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(request), uintptr(argp))
	if errno != 0 {
		return os.NewSyscallError("ioctl", errno)
	}
	return nil
}

const (
	_FBIOGET_VSCREENINFO = 0x4600
	_FBIOPUT_VSCREENINFO = 0x4601
	_FBIOGET_FSCREENINFO = 0x4602
	_FBIOGETCMAP         = 0x4604
	_FBIOPUTCMAP         = 0x4605
	_FBIOPAN_DISPLAY     = 0x4606
)

type fb_fix_screeninfo struct {
	id           [16]byte  //  identification string eg "TT Builtin"
	smem_start   uint32    //  Start of frame buffer mem (physical address)
	smem_len     uint32    //  Length of frame buffer mem
	_type        uint32    //  see FB_TYPE_*
	type_aux     uint32    //  Interleave for interleaved Planes
	visual       uint32    //  see FB_VISUAL_*
	xpanstep     uint16    //  zero if no hardware panning
	ypanstep     uint16    //  zero if no hardware panning
	ywrapstep    uint16    //  zero if no hardware ywrap
	line_length  uint32    //  length of a line in bytes
	mmio_start   uint32    //  Start of Memory Mapped I/O (physical address)
	mmio_len     uint32    //  Length of Memory Mapped I/O
	accel        uint32    //  Indicate to driver which specific chip/card we have
	capabilities uint16    //  see FB_CAP_*
	reserved     [2]uint16 //  Reserved for future compatibility
}

type fb_bitfield struct {
	offset    uint32 // beginning of bitfield
	length    uint32 // length of bitfield
	msb_right uint32 // != 0 : Most significant bit is right
}

type fb_var_screeninfo struct {
	xres         uint32 //  visible resolution
	yres         uint32
	xres_virtual uint32 //  virtual resolution
	yres_virtual uint32
	xoffset      uint32 //  offset from virtual to visible
	yoffset      uint32 //  resolution

	bits_per_pixel uint32      //  guess what
	grayscale      uint32      //  0 = color, 1 = grayscale, >1 = FOURCC
	red            fb_bitfield //  bitfield in fb mem if true color,
	green          fb_bitfield //  else only length is significant
	blue           fb_bitfield
	transp         fb_bitfield //  transparency

	nonstd uint32 //  != 0 Non standard pixel format

	activate uint32 //  see FB_ACTIVATE_*

	height uint32 //  height of picture in mm
	width  uint32 //  width of picture in mm

	accel_flags uint32 //  (OBSOLETE) see fb_info.flags

	//  Timing: All values in pixclocks, except pixclock (of course)
	pixclock     uint32 //  pixel clock in ps (pico seconds)
	left_margin  uint32 //  time from sync to picture
	right_margin uint32 //  time from picture to sync
	upper_margin uint32 //  time from sync to picture
	lower_margin uint32
	hsync_len    uint32    //  length of horizontal sync
	vsync_len    uint32    //  length of vertical sync
	sync         uint32    //  see FB_SYNC_*
	vmode        uint32    //  see FB_VMODE_*
	rotate       uint32    //  angle we rotate counter clockwise
	colorspace   uint32    //  colorspace for FOURCC-based modes
	reserved     [4]uint32 //  Reserved for future compatibility
}
