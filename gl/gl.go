package gl

import (
	"unsafe"
)

/*
#include "GLES2/gl2.h"

static inline void ShaderSource(GLuint s, const char *b) {
	const GLchar*bb[] = {b};
	glShaderSource(s, 1, bb, NULL);
}
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

var logbuf [1024]C.char

////////////////////////////////////////////////////////////////////////////////

// AttachShader attaches a shader object to a program object.
//
// http://docs.gl/es2/glAttachShader
func AttachShader(p Program, s Shader) {
	C.glAttachShader(C.GLuint(p), C.GLuint(s))
}

// BindBuffer binds a named buffer object.
//
// http://docs.gl/es2/glBindBuffer
func BindBuffer(target Enum, b Buffer) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(b))
}

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

// BufferData creates and initializes a buffer object's data store.
//
// http://docs.gl/es2/glBufferData
func BufferData(target Enum, size uintptr, data unsafe.Pointer, usage Enum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), data,
		C.GLenum(usage))
}

// Clear clears buffers to preset values.
//
// http://docs.gl/es2/glClear
func Clear(mask Enum) {
	C.glClear(C.GLbitfield(mask))
}

// ClearColor specifies clear values for the color buffers.
//
// http://docs.gl/es2/glClearColor
func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
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

// CompileShader compiles a shader object.
//
// http://docs.gl/es2/glCompileShader
func CompileShader(s Shader) {
	C.glCompileShader(C.GLuint(s))
}

// CreateProgram creates a program object.
//
// http://docs.gl/es2/glCreateProgram
func CreateProgram() Program {
	return Program(C.glCreateProgram())
}

// CreateShader creates a shader object
//
// http://docs.gl/es2/glCreateShader
func CreateShader(shaderType Enum) Shader {
	return Shader(C.glCreateShader(C.GLenum(shaderType)))
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

// DrawArrays renders primitives from array data.
//
// http://docs.gl/es2/glDrawArrays
func DrawArrays(mode Enum, first, count int32) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

// Enable enables server-side GL capabilities
//
// http://docs.gl/es2/glEnable
func Enable(cap Enum) {
	C.glEnable(C.GLenum(cap))
}

// EnableVertexAttribArray enables or disables a generic vertex attribute array.
//
// http://docs.gl/es2/glEnableVertexAttribArray
func EnableVertexAttribArray(index Attrib) {
	C.glEnableVertexAttribArray(C.GLuint(index))
}

// Finish blocks until all GL execution is complete.
//
// http://docs.gl/es2/glFinish
func Finish() {
	C.glFinish()
}

// Flush forces execution of GL commands in finite time.
//
// http://docs.gl/es2/glFlush
func Flush() {
	C.glFlush()
}

// FrontFace defines front- and back-facing polygons.
//
// http://docs.gl/es2/glFrontFace
func FrontFace(mode Enum) {
	C.glFrontFace(C.GLenum(mode))
}

// GenBuffer generates a buffer object name
//
// http://docs.gl/es2/glGenBuffers
func GenBuffer() Buffer {
	var b Buffer
	C.glGenBuffers(1, (*C.GLuint)(&b))
	return b
}

//TODO: func GenBuffers() []Buffer

// GetAttribLocation returns the location of an attribute variable.
//
// http://docs.gl/es2/glGetAttribLocation
func GetAttribLocation(p Program, name string) (a Attrib, ok bool) {
	aa := C.glGetAttribLocation(C.GLuint(p), C.CString(name))
	if aa == -1 {
		return 0, false
	}
	return Attrib(aa), true
}

// GetBooleanv returns the value or values of a selected parameter.
//
// http://docs.gl/es2/glGet
func GetBooleanv(pname Enum, dst []bool) {
	b := make([]C.GLboolean, len(dst))
	C.glGetBooleanv(C.GLenum(pname), &b[0])
	for i := range dst {
		dst[i] = b[i] == C.GL_TRUE
	}
}

// GetError returns error information
//
// http://docs.gl/es2/glGetError
func GetError() Enum {
	return Enum(C.glGetError())
}

// GetFloatv returns the value or values of a selected parameter.
//
// http://docs.gl/es2/glGet
func GetFloatv(pname Enum, dst []float32) {
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(&dst[0]))
}

// GetIntegerv returns the value or values of a selected parameter.
//
// http://docs.gl/es2/glGet
func GetIntegerv(pname Enum, dst []int32) {
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(&dst[0]))
}

// PixelStorei sets pixel storage modes.
//
// http://docs.gl/es2/glPixelStorei
func PixelStorei(pname Enum, param int32) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

// PolygonOffset sets the scale and units used to calculate depth values.
//
// http://docs.gl/es2/glPolygonOffset
func PolygonOffset(factor, units float32) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}

// GetProgramInfoLog returns the information log for a program object.
//
// http://docs.gl/es2/glGetProgramInfoLog
func GetProgramInfoLog(p Program) string {
	C.glGetProgramInfoLog(
		C.GLuint(p),
		C.int((unsafe.Sizeof(logbuf))),
		nil,
		&logbuf[0],
	)
	return C.GoString(&logbuf[0])
}

// SampleCoverage specifies multisample coverage parameters.
//
// http://docs.gl/es2/glSampleCoverage
func SampleCoverage(value float32, invert bool) {
	b := C.GL_FALSE
	if invert {
		b = C.GL_TRUE
	}
	C.glSampleCoverage(C.GLclampf(value), C.GLboolean(b))
}

// GetShaderInfoLog returns the information log for a shader object.
//
// http://docs.gl/es2/glGetShaderInfoLog
func GetShaderInfoLog(s Shader) string {
	C.glGetShaderInfoLog(
		C.GLuint(s),
		C.int((unsafe.Sizeof(logbuf))),
		nil,
		&logbuf[0],
	)
	return C.GoString(&logbuf[0])
}

// Hint specifies implementation-specific hints.
//
// http://docs.gl/es2/glHint
func Hint(target, mode Enum) {
	C.glHint(C.GLenum(target), C.GLenum(mode))
}

// IsEnabled tests whether a capability is enabled.
//
// http://docs.gl/es2/glIsEnabled
func IsEnabled(cap Enum) bool {
	r := C.glIsEnabled(C.GLenum(cap))
	return r == C.GL_TRUE
}

// LineWidth specifies the width of rasterized lines.
//
// http://docs.gl/es2/glLineWidth
func LineWidth(w float32) {
	C.glLineWidth(C.GLfloat(w))
}

// LinkProgram links a program object.
//
// http://docs.gl/es2/glLinkProgram
func LinkProgram(p Program) {
	C.glLinkProgram(C.GLuint(p))
}

// ShaderSource replaces the source code in a shader object.
//
// http://docs.gl/es2/glShaderSource
func ShaderSource(s Shader, source string) {
	C.ShaderSource(C.GLuint(s), C.CString(source))
}

// UseProgram installs a program object as part of current rendering state.
//
// http://docs.gl/es2/glUseProgram
func UseProgram(p Program) {
	C.glUseProgram(C.GLuint(p))
}

// VertexAttribPointer defines an array of generic vertex attribute data.
//
// http://docs.gl/es2/glVertexAttribPointer
func VertexAttribPointer(index Attrib, size int32, typ Enum, normalized bool, stride, offset uintptr) {
	n := C.GLboolean(FALSE)
	if normalized {
		n = C.GLboolean(TRUE)
	}
	C.glVertexAttribPointer(C.GLuint(index), C.GLint(size), C.GLenum(typ),
		n, C.GLsizei(stride), unsafe.Pointer(offset))
}

// Viewport sets the viewport.
//
// http://docs.gl/es2/glViewport
func Viewport(x, y, width, height int32) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
