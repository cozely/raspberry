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

// ActiveTexture selects active texture unit.
//
// http://docs.gl/es2/glActiveTexture
func ActiveTexture(target Enum) {
	C.glActiveTexture(C.GLenum(target))
}

// AttachShader attaches a shader object to a program object.
//
// http://docs.gl/es2/glAttachShader
func AttachShader(p Program, s Shader) {
	C.glAttachShader(C.GLuint(p), C.GLuint(s))
}

// BindAttribLocation associates a generic vertex attribute index with a named
// attribute variable.
//
// http://docs.gl/es2/glBindAttribLocation
func BindAttribLocation(p Program, index uint32, name string) {
	C.glBindAttribLocation(C.GLuint(p), C.GLuint(index), C.CString(name))
}

// BindBuffer binds a named buffer object.
//
// http://docs.gl/es2/glBindBuffer
func BindBuffer(target Enum, b Buffer) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(b))
}

// BindTexture binds a named texture to a texturing target.
//
// http://docs.gl/es2/glBindTexture
func BindTexture(target Enum, t Texture) {
	C.glBindTexture(C.GLenum(target), C.GLuint(t))
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
func BufferData(target Enum, data []byte, usage Enum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(len(data)), unsafe.Pointer(&data[0]),
		C.GLenum(usage))
}

// BufferDataUnsafe creates and initializes a buffer object's data store.
//
// http://docs.gl/es2/glBufferData
func BufferDataUnsafe(target Enum, size uintptr, data unsafe.Pointer, usage Enum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage))
}

// BufferSubData updates a subset of a buffer object's data store.
//
// http://docs.gl/es2/glBufferSubData
func BufferSubData(target Enum, offset int32, data []byte, usage Enum) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset),
		C.GLsizeiptr(len(data)), unsafe.Pointer(&data[0]))
}

// BufferSubDataUnsafe updates a subset of a buffer object's data store.
//
// http://docs.gl/es2/glBufferSubData
func BufferSubDataUnsafe(target Enum, offset, size uintptr, data unsafe.Pointer, usage Enum) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset),
		C.GLintptr(size), data)
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
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue),
		C.GLclampf(alpha))
}

// ClearDepthf specifies the clear value for the depth buffer.
//
// http://docs.gl/es2/glClearDepthf
func ClearDepthf(depth float32) {
	C.glClearDepthf(C.GLclampf(depth))
}

// ClearStencil specifies the clear value for the stencil buffer.
//
// http://docs.gl/es2/glClearStencil
func ClearStencil(s int32) {
	C.glClearStencil(C.GLint(s))
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

// CompressedTexImage2D specifies a two-dimensional texture image in a
// compressed format.
//
// http://docs.gl/es2/glCompressedTexImage2D
func CompressedTexImage2D(
	target Enum,
	level int32,
	internalFormat Enum,
	width, height int32,
	border int32,
	imageSize int32,
	data *byte,
) {
	C.glCompressedTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLenum(internalFormat),
		C.GLsizei(width), C.GLsizei(height),
		C.GLint(border),
		C.GLsizei(imageSize),
		unsafe.Pointer(data),
	)
}

// CompressedSubTexImage2D specifies a two-dimensional texture subimage in a
// compressed format.
//
// http://docs.gl/es2/glCompressedTexSubImage2D
func CompressedSubTexImage2D(
	target Enum,
	level int32,
	xoffset, yoffset int32,
	width, height int32,
	format Enum,
	imageSize int32,
	data *byte,
) {
	C.glCompressedTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format),
		C.GLsizei(imageSize),
		unsafe.Pointer(data),
	)
}

// CopyTexImage2D copies pixels into a 2D texture image.
//
// http://docs.gl/es2/glCopyTexImage2D
func CopyTexImage2D(
	target Enum,
	level int32,
	internalFormat Enum,
	x, y int32,
	width, height int32,
	border int32,
) {
	C.glCopyTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLenum(internalFormat),
		C.GLint(x), C.GLint(y),
		C.GLsizei(width), C.GLsizei(height),
		C.GLint(border),
	)
}

// CopyTexSubImage2D copies pixels into a 2D texture subimage.
//
// http://docs.gl/es2/glCopyTexSubImage2D
func CopyTexSubImage2D(
	target Enum,
	level int32,
	xoffset, yoffset int32,
	x, y int32,
	width, height int32,
) {
	C.glCopyTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLint(x), C.GLint(y),
		C.GLsizei(width), C.GLsizei(height),
	)
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

// DeleteBuffers deletes named buffer objects.
//
// http://docs.gl/es2/glDeleteBuffers
func DeleteBuffers(b []Buffer) {
	C.glDeleteBuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
}

// DeleteProgram deletes a program object.
//
// http://docs.gl/es2/glDeleteProgram
func DeleteProgram(p Program) {
	C.glDeleteProgram(C.GLuint(p))
}

// DeleteShader deletes a shader object.
//
// http://docs.gl/es2/glDeleteShader
func DeleteShader(s Shader) {
	C.glDeleteShader(C.GLuint(s))
}

// DetachShader detaches a shader object from a program object.
//
// http://docs.gl/es2/glDetachShader
func DetachShader(p Program, s Shader) {
	C.glDetachShader(C.GLuint(p), C.GLuint(s))
}

// DeleteTextures deletes named textures.
//
// http://docs.gl/es2/glDeleteTextures
func DeleteTextures(textures []Texture) {
	C.glDeleteTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
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

// DisableVertexArray disables a generic vertex attribute array
//
// http://docs.gl/es2/glEnableVertexAttribArray
func DisableVertexArray(a Attrib) {
	C.glDisableVertexAttribArray(C.GLuint(a))
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

// EnableVertexAttribArray enables a generic vertex attribute array.
//
// http://docs.gl/es2/glEnableVertexAttribArray
func EnableVertexAttribArray(a Attrib) {
	C.glEnableVertexAttribArray(C.GLuint(a))
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

// GenBuffers generates buffer object names.
//
// http://docs.gl/es2/glGenBuffers
func GenBuffers(count int32) []Buffer {
	b := make([]Buffer, count)
	C.glGenBuffers(C.GLsizei(count), (*C.GLuint)(&b[0]))
	return b
}

// GenTextures generate texture names.
//
// http://docs.gl/es2/glGenTextures
func GenTextures(count int32) []Texture {
	t := make([]Texture, count)
	C.glGenTextures(C.GLsizei(count), (*C.GLuint)(&t[0]))
	return t
}

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

// GetBufferParameteriv returns parameters of a buffer object.
//
// http://docs.gl/es2/glGetBufferParameteriv
func GetBufferParameteriv(target Enum, value Enum) int32 {
	var v C.GLint
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(value), &v)
	return int32(v)
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
	//TODO: Use a slice return instead
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(&dst[0]))
}

// GetIntegerv returns the value or values of a selected parameter.
//
// http://docs.gl/es2/glGet
func GetIntegerv(pname Enum, dst []int32) {
	//TODO: Use a slice return instead
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(&dst[0]))
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

// GetString returns a string describing the current GL connection.
//
// http://docs.gl/es2/glGetString
func GetString(name Enum) string {
	return C.GoString((*C.char)(C.glGetString(C.GLenum(name))))
}

// GetTexParameteriv returns texture integer parameter values.
//
// http://docs.gl/es2/glGetTexParameter
func GetTexParameteriv(target Enum, pname Enum) []int32 {
	v := []int32{0} //TODO: double-check that all parameters returns a single value
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&v[0]))
	return v
}

// GetTexParameterfv returns texture float parameter values.
//
// http://docs.gl/es2/glGetTexParameter
func GetTexParameterfv(target Enum, pname Enum) []float32 {
	v := []float32{0} //TODO: double-check that all parameters returns a single value
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&v[0]))
	return v
}

// GetUniformfv returns the value of a float uniform variable.
//
// http://docs.gl/es2/glGetUniform
func GetUniformfv(p Program, u Uniform, dst []float32) {
	C.glGetUniformfv(C.GLuint(p), C.GLint(u), (*C.GLfloat)(&dst[0]))
}

// GetUniformiv returns the value of an integer uniform variable.
//
// http://docs.gl/es2/glGetUniform
func GetUniformiv(p Program, u Uniform, dst []int32) {
	C.glGetUniformiv(C.GLuint(p), C.GLint(u), (*C.GLint)(&dst[0]))
}

// GetUniformLocation returns the location of a uniform variable.
//
// http://docs.gl/es2/glGetUniformLocation
func GetUniformLocation(p Program, name string) Uniform {
	return Uniform(C.glGetUniformLocation(C.GLuint(p), C.CString(name)))
}

// GetVertexAttribfv returns a generic vertex attribute parameter.
//
// http://docs.gl/es2/glGetVertexAttrib
func GetVertexAttribfv(a Attrib, pname Enum, params []float32) {
	C.glGetVertexAttribfv(C.GLuint(a), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

// GetVertexAttribiv returns a generic vertex attribute parameter.
//
// http://docs.gl/es2/glGetVertexAttrib
func GetVertexAttribiv(a Attrib, pname Enum, params []int32) {
	C.glGetVertexAttribiv(C.GLuint(a), C.GLenum(pname), (*C.GLint)(&params[0]))
}

// GetVertexAttribPointerv returns the address of the specified generic vertex
// attribute pointer.
//
// http://docs.gl/es2/glGetVertexAttribPointerv
func GetVertexAttribPointerv(a Attrib, pname Enum) unsafe.Pointer {
	var p unsafe.Pointer
	C.glGetVertexAttribPointerv(C.GLuint(a), C.GLenum(pname), &p)
	return p
}

// Hint specifies implementation-specific hints.
//
// http://docs.gl/es2/glHint
func Hint(target, mode Enum) {
	C.glHint(C.GLenum(target), C.GLenum(mode))
}

// IsBuffer determines if a name corresponds to a buffer object.
//
// http://docs.gl/es2/glIsBuffer
func IsBuffer(b Buffer) bool {
	return C.glIsBuffer(C.GLuint(b)) == C.GL_TRUE
}

// IsEnabled tests whether a capability is enabled.
//
// http://docs.gl/es2/glIsEnabled
func IsEnabled(cap Enum) bool {
	r := C.glIsEnabled(C.GLenum(cap))
	return r == C.GL_TRUE
}

// IsTexture determines if a name corresponds to a texture.
//
// http://docs.gl/es2/glIsTexture
func IsTexture(t Texture) bool {
	return C.glIsTexture(C.GLuint(t)) == C.GL_TRUE
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

// ReadPixels reads a block of pixels from the frame buffer.
//
// http://docs.gl/es2/glReadPixels
func ReadPixels(
	x, y int32,
	width, height int32,
	format Enum,
	pixeltype Enum,
	data []byte,
) {
	C.glReadPixels(
		C.GLint(x), C.GLint(y),
		C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(pixeltype),
		unsafe.Pointer(&data[0]),
	)
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

// Scissor defines the scissor box
//
// http://docs.gl/es2/glScissor
func Scissor(x, y, width, height int32) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

// ShaderSource replaces the source code in a shader object.
//
// http://docs.gl/es2/glShaderSource
func ShaderSource(s Shader, source string) {
	C.ShaderSource(C.GLuint(s), C.CString(source))
}

// StencilFunc sets front and back function and reference value for stencil
// testing.
//
// http://docs.gl/es2/glStencilFunc
func StencilFunc(fn Enum, ref int32, mask uint32) {
	C.glStencilFunc(C.GLenum(fn), C.GLint(ref), C.GLuint(mask))
}

// StencilFuncSeparate sets front and/or back function and reference value for
// stencil testing.
//
// http://docs.gl/es2/glStencilFuncSeparate
func StencilFuncSeparate(face Enum, fn Enum, ref int32, mask uint32) {
	C.glStencilFuncSeparate(C.GLenum(face), C.GLenum(fn), C.GLint(ref), C.GLuint(mask))
}

// StencilMask controls the front and back writing of individual bits in the
// stencil planes.
//
// http://docs.gl/es2/glStencilMask
func StencilMask(mask uint32) {
	C.glStencilMask(C.GLuint(mask))
}

// StencilMaskSeparate controls the front and/or back writing of individual bits
// in the stencil planes.
//
// http://docs.gl/es2/glStencilMask
func StencilMaskSeparate(face Enum, mask uint32) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}

// StencilOp sets front and back stencil test actions.
//
// http://docs.gl/es2/glStencilOp
func StencilOp(sfail, dpfail, dppass Enum) {
	C.glStencilOp(C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
}

// StencilOpSeparate sets front and back stencil test actions.
//
// http://docs.gl/es2/glStencilOp
func StencilOpSeparate(face Enum, sfail, dpfail, dppass Enum) {
	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
}

// TexImage2D specifies a two-dimensional texture image.
//
// http://docs.gl/es2/glTexImage2D
func TexImage2D(
	target Enum,
	level int32,
	internalFormat Enum,
	width, height int32,
	border int32,
	format Enum,
	texeltype Enum,
	data []byte,
) {
	C.glTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(internalFormat),
		C.GLsizei(width), C.GLsizei(height),
		C.GLint(border),
		C.GLenum(format),
		C.GLenum(texeltype),
		unsafe.Pointer(&data[0]),
	)
}

// TexImage2DUnsafe specifies a two-dimensional texture image.
//
// http://docs.gl/es2/glTexImage2D
func TexImage2DUnsafe(
	target Enum,
	level int32,
	internalFormat Enum,
	width, height int32,
	border int32,
	format Enum,
	texeltype Enum,
	data unsafe.Pointer,
) {
	C.glTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(internalFormat),
		C.GLsizei(width), C.GLsizei(height),
		C.GLint(border),
		C.GLenum(format),
		C.GLenum(texeltype),
		data,
	)
}

// TexParameterf sets a float texture parameter.
//
// http://docs.gl/es2/glTexParameter
func TexParameterf(target Enum, pname Enum, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

// TexParameterfv sets float texture parameters.
//
// http://docs.gl/es2/glTexParameter
func TexParameterfv(target Enum, pname Enum, params []float32) {
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname),
		(*C.GLfloat)(&params[0]))
}

// TexParameteri sets an integer texture parameter.
//
// http://docs.gl/es2/glTexParameter
func TexParameteri(target Enum, pname Enum, param int32) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

// TexParameteriv sets integer texture parameters.
//
// http://docs.gl/es2/glTexParameter
func TexParameteriv(target Enum, pname Enum, params []int32) {
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname),
		(*C.GLint)(&params[0]))
}

// TexSubImage2D specifies a two-dimensional texture subimage.
//
// http://docs.gl/es2/glTexSubImage2D
func TexSubImage2D(
	target Enum,
	level int32,
	xoffset, yoffset int32,
	width, height int32,
	format Enum,
	texeltype Enum,
	data []byte,
) {
	C.glTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(texeltype),
		unsafe.Pointer(&data[0]),
	)
}

// TexSubImage2DUnsafe specifies a two-dimensional texture subimage.
//
// http://docs.gl/es2/glTexSubImage2D
func TexSubImage2DUnsafe(
	target Enum,
	level int32,
	xoffset, yoffset int32,
	width, height int32,
	format Enum,
	texeltype Enum,
	data unsafe.Pointer,
) {
	C.glTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(texeltype),
		data,
	)
}

// Uniform1f specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform1f(u Uniform, v0 float32) {
	C.glUniform1f(C.GLint(u), C.GLfloat(v0))
}

// Uniform1fv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform1fv(u Uniform, values []float32) {
	C.glUniform1fv(C.GLint(u), C.GLsizei(len(values)),
		(*C.GLfloat)(&values[0]))
}

// Uniform1i specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform1i(u Uniform, v0 int32) {
	C.glUniform1i(C.GLint(u), C.GLint(v0))
}

// Uniform1iv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform1iv(u Uniform, values []int32) {
	C.glUniform2iv(C.GLint(u), C.GLsizei(len(values)),
		(*C.GLint)(&values[0]))
}

// Uniform2f specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform2f(u Uniform, v [2]float32) {
	C.glUniform2f(C.GLint(u), C.GLfloat(v[0]), C.GLfloat(v[1]))
}

// Uniform2fv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform2fv(u Uniform, values [][2]float32) {
	C.glUniform2fv(C.GLint(u), C.GLsizei(2*len(values)),
		(*C.GLfloat)(&values[0][0]))
}

// Uniform2i specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform2i(u Uniform, v [2]int32) {
	C.glUniform2i(C.GLint(u), C.GLint(v[0]), C.GLint(v[1]))
}

// Uniform2iv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform2iv(u Uniform, values [][2]int32) {
	C.glUniform2iv(C.GLint(u), C.GLsizei(2*len(values)),
		(*C.GLint)(&values[0][0]))
}

// Uniform3f specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform3f(u Uniform, v [3]float32) {
	C.glUniform3f(C.GLint(u), C.GLfloat(v[0]), C.GLfloat(v[1]), C.GLfloat(v[2]))
}

// Uniform3fv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform3fv(u Uniform, values [][3]float32) {
	C.glUniform3fv(C.GLint(u), C.GLsizei(3*len(values)),
		(*C.GLfloat)(&values[0][0]))
}

// Uniform3i specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform3i(u Uniform, v [3]int32) {
	C.glUniform3i(C.GLint(u), C.GLint(v[0]), C.GLint(v[1]), C.GLint(v[2]))
}

// Uniform3iv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform3iv(u Uniform, values [][3]int32) {
	C.glUniform3iv(C.GLint(u), C.GLsizei(3*len(values)),
		(*C.GLint)(&values[0][0]))
}

// Uniform4f specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform4f(u Uniform, v [4]float32) {
	C.glUniform4f(C.GLint(u), C.GLfloat(v[0]), C.GLfloat(v[1]), C.GLfloat(v[2]),
		C.GLfloat(v[3]))
}

// Uniform4fv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform4fv(u Uniform, values [][4]float32) {
	C.glUniform4fv(C.GLint(u), C.GLsizei(4*len(values)),
		(*C.GLfloat)(&values[0][0]))
}

// Uniform4i specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform4i(u Uniform, v [4]int32) {
	C.glUniform4i(C.GLint(u), C.GLint(v[0]), C.GLint(v[1]), C.GLint(v[2]),
		C.GLint(v[3]))
}

// Uniform4iv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func Uniform4iv(u Uniform, values [][4]int32) {
	C.glUniform4iv(C.GLint(u), C.GLsizei(4*len(values)),
		(*C.GLint)(&values[0][0]))
}

// UniformMatrix2f specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func UniformMatrix2f(u Uniform, values [2][2]float32) {
	C.glUniformMatrix4fv(C.GLint(u), C.GLsizei(2*2), C.GL_FALSE,
		(*C.GLfloat)(&values[0][0]))
}

// UniformMatrix2fv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func UniformMatrix2fv(u Uniform, values [][2][2]float32) {
	C.glUniformMatrix2fv(C.GLint(u), C.GLsizei(2*2*len(values)), C.GL_FALSE,
		(*C.GLfloat)(&values[0][0][0]))
}

// UniformMatrix3f specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func UniformMatrix3f(u Uniform, values [3][3]float32) {
	C.glUniformMatrix4fv(C.GLint(u), C.GLsizei(3*3), C.GL_FALSE,
		(*C.GLfloat)(&values[0][0]))
}

// UniformMatrix3fv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func UniformMatrix3fv(u Uniform, values [][3][3]float32) {
	C.glUniformMatrix3fv(C.GLint(u), C.GLsizei(3*3*len(values)), C.GL_FALSE,
		(*C.GLfloat)(&values[0][0][0]))
}

// UniformMatrix4f specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func UniformMatrix4f(u Uniform, values [4][4]float32) {
	C.glUniformMatrix4fv(C.GLint(u), C.GLsizei(4*4), C.GL_FALSE,
		(*C.GLfloat)(&values[0][0]))
}

// UniformMatrix4fv specifies the value of a uniform variable for the current program
// object.
//
// http://docs.gl/es2/glUniform
func UniformMatrix4fv(u Uniform, values [][4][4]float32) {
	C.glUniformMatrix4fv(C.GLint(u), C.GLsizei(4*4*len(values)), C.GL_FALSE,
		(*C.GLfloat)(&values[0][0][0]))
}

// UseProgram installs a program object as part of current rendering state.
//
// http://docs.gl/es2/glUseProgram
func UseProgram(p Program) {
	C.glUseProgram(C.GLuint(p))
}

// VertexAttrib1f specifies the value of a generic vertex attribute.
//
// http://docs.gl/es2/glVertexAttrib
func VertexAttrib1f(a Attrib, v float32) {
	C.glVertexAttrib1f(C.GLuint(a), C.GLfloat(v))
}

// VertexAttrib2f specifies the value of a generic vertex attribute.
//
// http://docs.gl/es2/glVertexAttrib
func VertexAttrib2f(a Attrib, v [2]float32) {
	C.glVertexAttrib2f(C.GLuint(a), C.GLfloat(v[0]), C.GLfloat(v[1]))
}

// VertexAttrib3f specifies the value of a generic vertex attribute.
//
// http://docs.gl/es2/glVertexAttrib
func VertexAttrib3f(a Attrib, v [3]float32) {
	C.glVertexAttrib3f(C.GLuint(a), C.GLfloat(v[0]), C.GLfloat(v[1]),
		C.GLfloat(v[2]))
}

// VertexAttrib4f specifies the value of a generic vertex attribute.
//
// http://docs.gl/es2/glVertexAttrib
func VertexAttrib4f(a Attrib, v [4]float32) {
	C.glVertexAttrib4f(C.GLuint(a), C.GLfloat(v[0]), C.GLfloat(v[1]),
		C.GLfloat(v[2]), C.GLfloat(v[3]))
}

// VertexAttribPointer defines an array of generic vertex attribute data.
//
// http://docs.gl/es2/glVertexAttribPointer
func VertexAttribPointer(a Attrib, size int32, typ Enum, normalized bool, stride, offset uintptr) {
	n := C.GLboolean(FALSE)
	if normalized {
		n = C.GLboolean(TRUE)
	}
	C.glVertexAttribPointer(C.GLuint(a), C.GLint(size), C.GLenum(typ),
		n, C.GLsizei(stride), unsafe.Pointer(offset))
}

// Viewport sets the viewport.
//
// http://docs.gl/es2/glViewport
func Viewport(x, y, width, height int32) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
