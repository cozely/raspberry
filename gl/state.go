package gl

/*
#include "GLES2/gl2.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

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

// Viewport sets the viewport.
//
// http://docs.gl/es2/glViewport
func Viewport(x, y, width, height int32) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

// GetError returns error information
//
// http://docs.gl/es2/glGetError
func GetError() Enum {
	return Enum(C.glGetError())
}
