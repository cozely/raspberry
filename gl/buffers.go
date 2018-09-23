package gl

import (
	"unsafe"
)

/*
#include "GLES2/gl2.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// A Buffer object
type Buffer uint32

// Attrib is a shader attribute location
type Attrib uint32

////////////////////////////////////////////////////////////////////////////////

// BindBuffer binds a named buffer object.
//
// http://docs.gl/es2/glBindBuffer
func BindBuffer(target Enum, b Buffer) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(b))
}

// BufferData creates and initializes a buffer object's data store.
//
// http://docs.gl/es2/glBufferData
func BufferData(target Enum, size uintptr, data unsafe.Pointer, usage Enum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), data,
		C.GLenum(usage))
}

// DrawArrays renders primitives from array data.
//
// http://docs.gl/es2/glDrawArrays
func DrawArrays(mode Enum, first, count int32) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

// EnableVertexAttribArray enables or disables a generic vertex attribute array.
//
// http://docs.gl/es2/glEnableVertexAttribArray
func EnableVertexAttribArray(index Attrib) {
	C.glEnableVertexAttribArray(C.GLuint(index))
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
