package dispmanx

/*
#include "bcm_host.h"
*/

////////////////////////////////////////////////////////////////////////////////

type Window struct {
	Element       ElementHandle
	Width, Height int32
}

func (*Window) IsEGLNativeWindow() {}
