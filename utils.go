package main

import (
	"log"
	"path"
	"runtime"
)

/*
#include "GLES2/gl2.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

func checkgl() {
	e := C.glGetError()
	if e != 0 {
		_, f, l, ok := runtime.Caller(1)
		if !ok {
			log.Printf("[?:?] *** OpenGL error 0x%X", e)
		}
		log.Printf("[%s:%d] *** OpenGL error 0x%X", path.Base(f), l, e)
	}
}
