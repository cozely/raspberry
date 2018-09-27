package gl

/*
#include "GLES2/gl2.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// BlendColor sets the blend color.
//
// http://docs.gl/es2/glBlendColor
func BlendColor(red, green, blue, alpha float32) {
	C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

// BlendEquation specifies the equation used for both the RGB blend equation and
// the Alpha blend equation.
//
// http://docs.gl/es2/glBlendEquation
func BlendEquation(mode Enum) {
	C.glBlendEquation(C.GLenum(mode))
}

// BlendEquationSeparate sets the RGB blend equation and the alpha blend
// equation separately.
//
// http://docs.gl/es2/glBlendEquationSeparate
func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

// BlendFunc specifies pixel arithmetic.
//
// http://docs.gl/es2/glBlendFunc
func BlendFunc(srcFactor, dstFactor Enum) {
	C.glBlendFunc(C.GLenum(srcFactor), C.GLenum(dstFactor))
}

// BlendFuncSeparate specifies pixel arithmetic for RGB and alpha components
// separately.
//
// http://docs.gl/es2/glBlendFuncSeparate
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha Enum) {
	C.glBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB),
		C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

// ColorMask enables and disables writing of frame buffer color components.
//
// http://docs.gl/es2/glColorMask
func ColorMask(red, green, blue, alpha bool) {
	r, g, b, a := C.GLboolean(C.GL_FALSE), C.GLboolean(C.GL_FALSE),
		C.GLboolean(C.GL_FALSE), C.GLboolean(C.GL_FALSE)
	if red {
		r = C.GL_TRUE
	}
	if green {
		g = C.GL_TRUE
	}
	if blue {
		b = C.GL_TRUE
	}
	if alpha {
		a = C.GL_TRUE
	}
	C.glColorMask(r, g, b, a)
}

// CullFace specifies whether front- or back-facing polygons can be culled.
//
// http://docs.gl/es2/glCullFace
func CullFace(mode Enum) {
	C.glCullFace(C.GLenum(mode))
}

// DepthFunc specifies the value used for depth buffer comparisons.
//
// http://docs.gl/es2/glDepthFunc
func DepthFunc(fn Enum) {
	C.glDepthFunc(C.GLenum(fn))
}

// DepthMask enables or disables writing into the depth buffer.
//
// http://docs.gl/es2/glDepthMask
func DepthMask(flag bool) {
	if flag {
		C.glDepthMask(C.GL_TRUE)
		return
	}
	C.glDepthMask(C.GL_FALSE)
}

// DepthRangef specifies mapping of depth values from normalized device
// coordinates to window coordinates.
//
// http://docs.gl/es2/glDepthRangef
func DepthRangef(near, far float32) {
	C.glDepthRangef(C.GLclampf(near), C.GLclampf(far))
}

// Disable disables server-side GL capabilities
//
// http://docs.gl/es2/glEnable
func Disable(cap Enum) {
	C.glDisable(C.GLenum(cap))
}

// Enable enables server-side GL capabilities
//
// http://docs.gl/es2/glEnable
func Enable(cap Enum) {
	C.glEnable(C.GLenum(cap))
}

// FrontFace defines front- and back-facing polygons.
//
// http://docs.gl/es2/glFrontFace
func FrontFace(mode Enum) {
	C.glFrontFace(C.GLenum(mode))
}

// GetError returns error information
//
// http://docs.gl/es2/glGetError
func GetError() Enum {
	return Enum(C.glGetError())
}

// Viewport sets the viewport.
//
// http://docs.gl/es2/glViewport
func Viewport(x, y, width, height int32) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
