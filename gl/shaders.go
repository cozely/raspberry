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

// A Shader object
type Shader uint32

// A Program object
type Program uint32

var logbuf [1024]C.char

////////////////////////////////////////////////////////////////////////////////

// CreateShader creates a shader object
//
// http://docs.gl/es2/glCreateShader
func CreateShader(shaderType Enum) Shader {
	return Shader(C.glCreateShader(C.GLenum(shaderType)))
}

// ShaderSource replaces the source code in a shader object.
//
// http://docs.gl/es2/glShaderSource
func ShaderSource(s Shader, source string) {
	C.ShaderSource(C.GLuint(s), C.CString(source))
}

// CompileShader compiles a shader object.
//
// http://docs.gl/es2/glCompileShader
func CompileShader(s Shader) {
	C.glCompileShader(C.GLuint(s))
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

// CreateProgram creates a program object.
//
// http://docs.gl/es2/glCreateProgram
func CreateProgram() Program {
	return Program(C.glCreateProgram())
}

// AttachShader attaches a shader object to a program object.
//
// http://docs.gl/es2/glAttachShader
func AttachShader(p Program, s Shader) {
	C.glAttachShader(C.GLuint(p), C.GLuint(s))
}

// LinkProgram links a program object.
//
// http://docs.gl/es2/glLinkProgram
func LinkProgram(p Program) {
	C.glLinkProgram(C.GLuint(p))
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

// GetAttribLocation returns the location of an attribute variable.
//
// http://docs.gl/es2/glGetAttribLocation
func GetAttribLocation(p Program, name string) (a Attrib, ok bool) {
	aa := C.glGetAttribLocation(C.GLuint(p), C.CString(name))
	if aa == -1  {
		return 0, false
	}
	return Attrib(aa), true
}

// UseProgram installs a program object as part of current rendering state.
//
// http://docs.gl/es2/glUseProgram
func UseProgram(p Program) {
	C.glUseProgram(C.GLuint(p))
}
