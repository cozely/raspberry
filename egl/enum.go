package egl

/*
#include "EGL/egl.h"
#include "EGL/eglext.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// EGL Versioning
const (
	VERSION_1_0 Int = 1
	VERSION_1_1 Int = 1
	VERSION_1_2 Int = 1
	VERSION_1_3 Int = 1
	VERSION_1_4 Int = 1
)

// EGL Enumerants. Bitmasks and other exceptional cases aside, most
// enums are assigned unique values starting at 0x3000.

// EGL aliases
const (
	FALSE = 0
	TRUE  = 1
)

// Out-of-band handle values
var (
	DEFAULT_DISPLAY = NativeDisplay(nil)
	NO_DISPLAY      = Display(nil)
	NO_CONTEXT      = Context(nil)
	NO_SURFACE      = Surface(nil)
)

// Out-of-band attribute value
const DONT_CARE Int = 1

// Errors / GetError return values
const (
	SUCCESS             Int = 0x3000
	NOT_INITIALIZED     Int = 0x3001
	BAD_ACCESS          Int = 0x3002
	BAD_ALLOC           Int = 0x3003
	BAD_ATTRIBUTE       Int = 0x3004
	BAD_CONFIG          Int = 0x3005
	BAD_CONTEXT         Int = 0x3006
	BAD_CURRENT_SURFACE Int = 0x3007
	BAD_DISPLAY         Int = 0x3008
	BAD_MATCH           Int = 0x3009
	BAD_NATIVE_PIXMAP   Int = 0x300A
	BAD_NATIVE_WINDOW   Int = 0x300B
	BAD_PARAMETER       Int = 0x300C
	BAD_SURFACE         Int = 0x300D
	CONTEXT_LOST        Int = 0x300E // EGL 1.1 - IMG_power_management
)

// Reserved 0x300F-0x301F for additional errors

// Config attributes
const (
	BUFFER_SIZE             Int = 0x3020
	ALPHA_SIZE              Int = 0x3021
	BLUE_SIZE               Int = 0x3022
	GREEN_SIZE              Int = 0x3023
	RED_SIZE                Int = 0x3024
	DEPTH_SIZE              Int = 0x3025
	STENCIL_SIZE            Int = 0x3026
	CONFIG_CAVEAT           Int = 0x3027
	CONFIG_ID               Int = 0x3028
	LEVEL                   Int = 0x3029
	MAX_PBUFFER_HEIGHT      Int = 0x302A
	MAX_PBUFFER_PIXELS      Int = 0x302B
	MAX_PBUFFER_WIDTH       Int = 0x302C
	NATIVE_RENDERABLE       Int = 0x302D
	NATIVE_VISUAL_ID        Int = 0x302E
	NATIVE_VISUAL_TYPE      Int = 0x302F
	SAMPLES                 Int = 0x3031
	SAMPLE_BUFFERS          Int = 0x3032
	SURFACE_TYPE            Int = 0x3033
	TRANSPARENT_TYPE        Int = 0x3034
	TRANSPARENT_BLUE_VALUE  Int = 0x3035
	TRANSPARENT_GREEN_VALUE Int = 0x3036
	TRANSPARENT_RED_VALUE   Int = 0x3037
	NONE                    Int = 0x3038 // Attrib list terminator
	BIND_TO_TEXTURE_RGB     Int = 0x3039
	BIND_TO_TEXTURE_RGBA    Int = 0x303A
	MIN_SWAP_INTERVAL       Int = 0x303B
	MAX_SWAP_INTERVAL       Int = 0x303C
	LUMINANCE_SIZE          Int = 0x303D
	ALPHA_MASK_SIZE         Int = 0x303E
	COLOR_BUFFER_TYPE       Int = 0x303F
	RENDERABLE_TYPE         Int = 0x3040
	MATCH_NATIVE_PIXMAP     Int = 0x3041 // Pseudo-attribute (not queryable)
	CONFORMANT              Int = 0x3042
)

// Reserved 0x3041-0x304F for additional config attributes

// Config attribute values
const (
	SLOW_CONFIG           Int = 0x3050 // EGL_CONFIG_CAVEAT value
	NON_CONFORMANT_CONFIG Int = 0x3051 // EGL_CONFIG_CAVEAT value
	TRANSPARENT_RGB       Int = 0x3052 // EGL_TRANSPARENT_TYPE value
	RGB_BUFFER            Int = 0x308E // EGL_COLOR_BUFFER_TYPE value
	LUMINANCE_BUFFER      Int = 0x308F // EGL_COLOR_BUFFER_TYPE value
)

// More config attribute values, for EGL_TEXTURE_FORMAT
const (
	NO_TEXTURE   Int = 0x305C
	TEXTURE_RGB  Int = 0x305D
	TEXTURE_RGBA Int = 0x305E
	TEXTURE_2D   Int = 0x305F
)

// Config attribute mask bits
const (
	PBUFFER_BIT                 Int = 0x0001 // EGL_SURFACE_TYPE mask bits
	PIXMAP_BIT                  Int = 0x0002 // EGL_SURFACE_TYPE mask bits
	WINDOW_BIT                  Int = 0x0004 // EGL_SURFACE_TYPE mask bits
	VG_COLORSPACE_LINEAR_BIT    Int = 0x0020 // EGL_SURFACE_TYPE mask bits
	VG_ALPHA_FORMAT_PRE_BIT     Int = 0x0040 // EGL_SURFACE_TYPE mask bits
	MULTISAMPLE_RESOLVE_BOX_BIT Int = 0x0200 // EGL_SURFACE_TYPE mask bits
	SWAP_BEHAVIOR_PRESERVED_BIT Int = 0x0400 // EGL_SURFACE_TYPE mask bits

	OPENGL_ES_BIT  Int = 0x0001 // EGL_RENDERABLE_TYPE mask bits
	OPENVG_BIT     Int = 0x0002 // EGL_RENDERABLE_TYPE mask bits
	OPENGL_ES2_BIT Int = 0x0004 // EGL_RENDERABLE_TYPE mask bits
	OPENGL_BIT     Int = 0x0008 // EGL_RENDERABLE_TYPE mask bits
)

// QueryString targets
const (
	VENDOR      Int = 0x3053
	VERSION     Int = 0x3054
	EXTENSIONS  Int = 0x3055
	CLIENT_APIS Int = 0x308D
)

// QuerySurface / SurfaceAttrib / CreatePbufferSurface targets
const (
	HEIGHT                Int = 0x3056
	WIDTH                 Int = 0x3057
	LARGEST_PBUFFER       Int = 0x3058
	TEXTURE_FORMAT        Int = 0x3080
	TEXTURE_TARGET        Int = 0x3081
	MIPMAP_TEXTURE        Int = 0x3082
	MIPMAP_LEVEL          Int = 0x3083
	RENDER_BUFFER         Int = 0x3086
	VG_COLORSPACE         Int = 0x3087
	VG_ALPHA_FORMAT       Int = 0x3088
	HORIZONTAL_RESOLUTION Int = 0x3090
	VERTICAL_RESOLUTION   Int = 0x3091
	PIXEL_ASPECT_RATIO    Int = 0x3092
	SWAP_BEHAVIOR         Int = 0x3093
	MULTISAMPLE_RESOLVE   Int = 0x3099
)

// EGL_RENDER_BUFFER values / BindTexImage / ReleaseTexImage buffer targets
const (
	BACK_BUFFER   Int = 0x3084
	SINGLE_BUFFER Int = 0x3085
)

// OpenVG color spaces
const (
	VG_COLORSPACE_sRGB   Int = 0x3089 // EGL_VG_COLORSPACE value
	VG_COLORSPACE_LINEAR Int = 0x308A // EGL_VG_COLORSPACE value
)

// OpenVG alpha formats
const (
	VG_ALPHA_FORMAT_NONPRE Int = 0x308B // EGL_ALPHA_FORMAT value
	VG_ALPHA_FORMAT_PRE    Int = 0x308C // EGL_ALPHA_FORMAT value
)

// Constant scale factor by which fractional display resolutions &
// aspect ratio are scaled when queried as integer values.
const DISPLAY_SCALING = 10000

// Unknown display resolution/aspect ratio
const UNKNOWN Int = -1

// Back buffer swap behaviors
const (
	BUFFER_PRESERVED Int = 0x3094 // EGL_SWAP_BEHAVIOR value
	BUFFER_DESTROYED Int = 0x3095 // EGL_SWAP_BEHAVIOR value
)

// CreatePbufferFromClientBuffer buffer types
const OPENVG_IMAGE Int = 0x3096

// QueryContext targets
const CONTEXT_CLIENT_TYPE Int = 0x3097

// CreateContext attributes
const CONTEXT_CLIENT_VERSION Int = 0x3098

// Multisample resolution behaviors
const (
	MULTISAMPLE_RESOLVE_DEFAULT Int = 0x309A // EGL_MULTISAMPLE_RESOLVE value
	MULTISAMPLE_RESOLVE_BOX     Int = 0x309B // EGL_MULTISAMPLE_RESOLVE value
)

// BindAPI/QueryAPI targets
const (
	OPENGL_ES_API Int = 0x30A0
	OPENVG_API    Int = 0x30A1
	OPENGL_API    Int = 0x30A2
)

// GetCurrentSurface targets
const (
	DRAW Int = 0x3059
	READ Int = 0x305A
)

// WaitNative engines
const CORE_NATIVE_ENGINE Int = 0x305B

// EGL 1.2 tokens renamed for consistency in EGL 1.3
const (
	COLORSPACE          = VG_COLORSPACE
	ALPHA_FORMAT        = VG_ALPHA_FORMAT
	COLORSPACE_sRGB     = VG_COLORSPACE_sRGB
	COLORSPACE_LINEAR   = VG_COLORSPACE_LINEAR
	ALPHA_FORMAT_NONPRE = VG_ALPHA_FORMAT_NONPRE
	ALPHA_FORMAT_PRE    = VG_ALPHA_FORMAT_PRE
)
