package main

import (
	"log"
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

var pipeline struct {
	program     C.GLuint
	framebuffer C.GLuint
	texture     C.GLuint
	vbo         C.GLuint
}

var vshader = `
attribute vec4 vertex;
void main(void) {
	gl_Position = vertex;
}
`

var fshader = `
void main(void) {
	gl_FragColor = vec4(0.3, 0.1, 0.6, 1.0);
}
`

var vertices = [...]C.GLfloat{
	0, 0.65, 0.5, 1,
	-0.65, -0.475, 0.5, 1,
	0.65, -0.475, 0.5, 1,
	0, 0.65, 0.5, 1,
}

var logbuf [1024]C.char

////////////////////////////////////////////////////////////////////////////////

func createPipeline() error {
	// Compile and link the shaders

	vs := C.glCreateShader(C.GL_VERTEX_SHADER)
	C.ShaderSource(vs, C.CString(vshader))
	C.glCompileShader(vs)
	//TODO: the shader info log doesn't seem to contain any info, ever?
	C.glGetShaderInfoLog(
		vs,
		C.int((unsafe.Sizeof(logbuf))),
		nil,
		&logbuf[0],
	)
	log.Printf("Vertex Shader: %s\n", C.GoString(&logbuf[0]))
	checkgl()

	fs := C.glCreateShader(C.GL_FRAGMENT_SHADER)
	C.ShaderSource(fs, C.CString(fshader))
	C.glCompileShader(fs)
	//TODO: the shader info log doesn't seem to contain any info, ever?
	C.glGetShaderInfoLog(
		fs,
		C.int((unsafe.Sizeof(logbuf))),
		nil,
		&logbuf[0],
	)
	log.Printf("Fragment Shader: %s\n", C.GoString(&logbuf[0]))
	checkgl()

	pipeline.program = C.glCreateProgram()
	C.glAttachShader(pipeline.program, vs)
	C.glAttachShader(pipeline.program, fs)
	C.glLinkProgram(pipeline.program)
	C.glGetProgramInfoLog(
		pipeline.program,
		C.int((unsafe.Sizeof(logbuf))),
		nil,
		&logbuf[0],
	)
	log.Printf("Program Link: %s\n", C.GoString(&logbuf[0]))
	checkgl()

	// Create the framebuffer

	// C.glGenTextures(1, &pipeline.texture)
	// checkgl()
	// C.glBindTexture(C.GL_TEXTURE_2D, pipeline.texture)
	// checkgl()
	// C.glTexImage2D(
	// 	C.GL_TEXTURE_2D,
	// 	0, C.GL_RGB,
	// 	C.int(screen.width), C.int(screen.height),
	// 	0, C.GL_RGB,
	// 	C.GL_UNSIGNED_SHORT_5_6_5, nil,
	// )
	// checkgl()
	// C.glTexParameterf(C.GL_TEXTURE_2D, C.GL_TEXTURE_MIN_FILTER, C.GL_NEAREST)
	// C.glTexParameterf(C.GL_TEXTURE_2D, C.GL_TEXTURE_MAG_FILTER, C.GL_NEAREST)
	// checkgl()
	// C.glGenFramebuffers(1, &pipeline.framebuffer)
	// checkgl()
	// C.glBindFramebuffer(C.GL_FRAMEBUFFER, pipeline.framebuffer)
	// checkgl()
	// C.glFramebufferTexture2D(C.GL_FRAMEBUFFER, C.GL_COLOR_ATTACHMENT0,
	// 	C.GL_TEXTURE_2D, pipeline.texture, 0)
	// checkgl()
	// C.glBindFramebuffer(C.GL_FRAMEBUFFER, C.GLuint(0))
	// checkgl()

	C.glViewport(0, 0, C.int(screen.width), C.int(screen.height))

	C.glGenBuffers(1, &pipeline.vbo)
	checkgl()
	C.glBindBuffer(C.GL_ARRAY_BUFFER, pipeline.vbo)
	C.glBufferData(C.GL_ARRAY_BUFFER,
		C.long(unsafe.Sizeof(vertices)),
		unsafe.Pointer(&vertices[0]),
		C.GL_STATIC_DRAW)
	a := C.GLuint(C.glGetAttribLocation(pipeline.program, C.CString("vertex")))
	checkgl()
	C.glVertexAttribPointer(a, 4, C.GL_FLOAT, 0, 16, nil)
	C.glEnableVertexAttribArray(a)
	checkgl()

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func drawTriangle() error {
	C.glClearColor(0.99, 0.97, 0.90, 1)
	C.glClear(C.GL_COLOR_BUFFER_BIT)

	C.glUseProgram(pipeline.program)
	C.glBindBuffer(C.GL_ARRAY_BUFFER, pipeline.vbo)
	C.glDrawArrays(C.GL_TRIANGLES, 0, 3)

	C.glBindBuffer(C.GL_ARRAY_BUFFER, 0)
	C.glFlush()
	C.glFinish()

	swapBuffers()
	return nil
}
