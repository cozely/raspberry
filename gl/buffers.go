package gl

import (
	"unsafe"
)

/*
#include "GLES2/gl2.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

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

// EnableVertexAttribArray enables or disables a generic vertex attribute array.
//
// http://docs.gl/es2/glEnableVertexAttribArray
func EnableVertexAttribArray(index Attrib) {
	C.glEnableVertexAttribArray(C.GLuint(index))
}
