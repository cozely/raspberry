package dispmanx

/*
#include "bcm_host.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

const NO_HANDLE = 0

const (
	PROTECTION_MAX  Protection = 0x0F
	PROTECTION_NONE Protection = 0
	PROTECTION_HDCP Protection = 11
)

type Transform uint32

const (
	// Bottom 2 bits sets the orientation

	NO_ROTATE  Transform = 0
	ROTATE_90  Transform = 1
	ROTATE_180 Transform = 2
	ROTATE_270 Transform = 3

	FLIP_HRIZ Transform = 1 << 16
	FLIP_VERT Transform = 1 << 17

	// extra flags for controlling 3d duplication behaviour

	STEREOSCOPI_IINVERT Transform = 1 << 19 // invert left/right images
	STEREOSCOPI_NONE    Transform = 0 << 20
	STEREOSCOPI_MONO    Transform = 1 << 20
	STEREOSCOPI_SBS     Transform = 2 << 20
	STEREOSCOPI_TB      Transform = 3 << 20
	STEREOSCOPI_MASK    Transform = 15 << 20

	// extra flags for controlling snapshot behaviour

	SNAPSHOT_NO_YUV        Transform = 1 << 24
	SNAPSHOT_NO_RGB        Transform = 1 << 25
	SNAPSHOT_FILL          Transform = 1 << 26
	SNAPSHOT_SWAP_RED_BLUE Transform = 1 << 27
	SNAPSHOT_PACK          Transform = 1 << 28
)
