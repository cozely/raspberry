package framebuffer

import (
	"unsafe"

	"github.com/cozely/journal"
	"golang.org/x/sys/unix"
)

//------------------------------------------------------------------------------

// Info is the journal used to log informations.
var Info = journal.New(journal.Stdout, "")

var (
	fbfd     int
	originfo fb_var_screeninfo
	varinfo  fb_var_screeninfo
	fixinfo  fb_fix_screeninfo
	frames   [2]uint32
	memory   []byte
)

//------------------------------------------------------------------------------

// Setup prepares the framebuffer.
func Setup() error {
	var err error

	//TODO: support choosing framebuffer number
	fbfd, err = unix.Open("/dev/fb0", unix.O_RDWR, 0)
	if err != nil {
		return wrap("unable to open /dev/fb0", err)
	}

	// First, get the current state of the framebuffer
	err = ioctl(fbfd, _FBIOGET_VSCREENINFO, unsafe.Pointer(&originfo))
	if err != nil {
		return wrap("unable to get current framebuffer state", err)
	}
	frames = [2]uint32{0, originfo.yres}

	// Modify it to fit our needs
	varinfo = originfo
	varinfo.xres_virtual = varinfo.xres
	varinfo.yres_virtual = 2 * varinfo.yres
	varinfo.xoffset = 0
	varinfo.yoffset = 0
	varinfo.bits_per_pixel = 32
	err = ioctl(fbfd, _FBIOPUT_VSCREENINFO, unsafe.Pointer(&varinfo))
	if err != nil {
		return wrap("unable to change framebuffer state", err)
	}
	err = ioctl(fbfd, _FBIOGET_VSCREENINFO, unsafe.Pointer(&varinfo))
	if err != nil {
		return wrap("unable to get current framebuffer state", err)
	}
	//TODO: check if new state is correct

	// Findout framebuffer size and stride
	err = ioctl(fbfd, _FBIOGET_FSCREENINFO, unsafe.Pointer(&fixinfo))
	if err != nil {
		return wrap("unable to get framebuffer fixed info", err)
	}

	Info.Printf("\"%s\" resolution: %dx%d(%d), virtual: %dx%d +%d+%d, size %d, stride %d\n",
		fixinfo.id,
		varinfo.xres, varinfo.yres,
		varinfo.bits_per_pixel,
		varinfo.xres_virtual, varinfo.yres_virtual,
		varinfo.xoffset, varinfo.yoffset,
		fixinfo.smem_len, fixinfo.line_length)

	memory, err = unix.Mmap(
		fbfd,
		0,
		int(fixinfo.smem_len),
		unix.PROT_READ|unix.PROT_WRITE,
		unix.MAP_SHARED,
	)
	if err != nil {
		wrap("unable to memory-map the frambuffer", err)
	}
	// Registers = (*[41]uint32)(unsafe.Pointer(&m[0]))

	return nil
}

//------------------------------------------------------------------------------

func Swap() error {
	frames[0], frames[1] = frames[1], frames[0]
	varinfo.yoffset = frames[0]
	err := ioctl(fbfd, _FBIOPAN_DISPLAY, unsafe.Pointer(&varinfo))
	if err != nil {
		return wrap("unable to switch framebuffer page", err)
	}
	return nil
}

//------------------------------------------------------------------------------

func Cleanup() {
	err := ioctl(fbfd, _FBIOPUT_VSCREENINFO, unsafe.Pointer(&originfo))
	if err != nil {
		journal.Printf("Unable to clean up framebuffer state: %v", err)
	}

	unix.Close(fbfd)

	Info.Print("Previous state restored.")
}
