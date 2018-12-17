package main

import (
	"log"
	"unsafe"

	"github.com/cozely/raspberry/gl"
)

////////////////////////////////////////////////////////////////////////////////

var pipeline struct {
	program gl.Program
	// framebuffer C.GLuint
	// texture     C.GLuint
	vbo gl.Buffer
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

var vertices = [...]float32{
	0, 0.65, 0.5, 1,
	-0.65, -0.475, 0.5, 1,
	0.65, -0.475, 0.5, 1,
	0, 0.65, 0.5, 1,
}

////////////////////////////////////////////////////////////////////////////////

func createPipeline() error {
	// Compile and link the shaders

	vs := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vs, vshader)
	gl.CompileShader(vs)
	//TODO: the shader info log doesn't seem to contain any info, ever?
	log.Printf("Vertex Shader: %s\n", gl.GetShaderInfoLog(vs))
	checkgl()

	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fs, fshader)
	gl.CompileShader(fs)
	//TODO: the shader info log doesn't seem to contain any info, ever?
	log.Printf("Fragment Shader: %s\n", gl.GetShaderInfoLog(fs))
	checkgl()

	pipeline.program = gl.CreateProgram()
	gl.AttachShader(pipeline.program, vs)
	gl.AttachShader(pipeline.program, fs)
	gl.LinkProgram(pipeline.program)
	log.Printf("Program Link: %s\n", gl.GetProgramInfoLog(pipeline.program))
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

	gl.Viewport(0, 0, screen.width, screen.height)

	pipeline.vbo = gl.GenBuffers(1)[0]
	checkgl()
	gl.BindBuffer(gl.ARRAY_BUFFER, pipeline.vbo)
	gl.BufferDataUnsafe(gl.ARRAY_BUFFER, unsafe.Sizeof(vertices),
		unsafe.Pointer(&vertices[0]), gl.STATIC_DRAW)
	a, ok := gl.GetAttribLocation(pipeline.program, "vertex")
	if !ok {
		log.Printf("*** unable to get location of attribute \"vertex\"")
	}
	checkgl()
	gl.VertexAttribPointer(a, 4, gl.FLOAT, false, 16, 0)
	gl.EnableVertexAttribArray(a)
	checkgl()

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func drawTriangle() error {
	gl.ClearColor(0.99, 0.97, 0.90, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.UseProgram(pipeline.program)
	gl.BindBuffer(gl.ARRAY_BUFFER, pipeline.vbo)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.Flush()
	gl.Finish() // TODO: necessary?

	swapBuffers()
	return nil
}
