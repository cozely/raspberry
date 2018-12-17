package main

import (
	"fmt"

	"github.com/cozely/raspberry/bcm"
	"github.com/cozely/raspberry/dispmanx"
	"github.com/cozely/raspberry/egl"
)

////////////////////////////////////////////////////////////////////////////////

var screen struct {
	width, height int32

	display egl.Display
	surface egl.Surface
	context egl.Context
}

////////////////////////////////////////////////////////////////////////////////

func initScreen() error {
	var err error

	bcm.HostInit()

	screen.width, screen.height, err = bcm.GetDisplaySize(0)
	if err != nil {
		return fmt.Errorf("initScreen: %v", err)
	}

	// Establish a connection with the display

	screen.display, err = egl.GetDisplay(egl.DEFAULT_DISPLAY)
	if err != nil {
		return fmt.Errorf("initScreen: %v", err)
	}
	checkgl()

	_, _, err = egl.Initialize(screen.display)
	if err != nil {
		return fmt.Errorf("initScreen: %v", err)
	}
	checkgl()

	// Create and configure an EGL context

	afb := []egl.Int{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.SURFACE_TYPE, egl.WINDOW_BIT,
		egl.NONE,
	}
	conf, err := egl.ChooseConfig(screen.display, afb)
	if err != nil {
		return fmt.Errorf("initScreen: %v", err)
	}
	if len(conf) == 0 {
		return fmt.Errorf("initScreen: no framebuffer configuration available")
	}
	checkgl()

	err = egl.BindAPI(egl.OPENGL_ES_API)
	if err != nil {
		return fmt.Errorf("initScreen: %v", err)
	}
	checkgl()

	a := []egl.Int{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}
	screen.context, err = egl.CreateContext(screen.display, conf[0], egl.NO_CONTEXT, a)
	if err != nil {
		return fmt.Errorf("initScreen: %v", err)
	}
	checkgl()

	// Create an element (i.e. layer/sprite) with DispmanX

	dpy := dispmanx.DisplayOpen(0)
	upd := dispmanx.UpdateStart(0)

	src := dispmanx.Rect{
		X: 0, Y: 0,
		Width:  screen.width << 16,
		Height: screen.height << 16,
	}
	dst := dispmanx.Rect{
		X: 0, Y: 0,
		Width:  screen.width,
		Height: screen.height,
	}
	elm := dispmanx.ElementAdd(upd, dpy,
		0, dst, 0, src, dispmanx.PROTECTION_NONE,
		nil, nil, 0)

	dispmanx.UpdateSubmitSync(upd)
	checkgl()

	// Create an EGL window surface

	w := dispmanx.Window{
		Element: elm,
		Width:   screen.width,
		Height:  screen.height,
	}
	screen.surface, err = egl.CreateWindowSurface(screen.display, conf[0], &w, nil)
	if err != nil {
		return fmt.Errorf("screenInit: %v", err)
	}
	checkgl()

	err = egl.MakeCurrent(screen.display, screen.surface, screen.surface, screen.context)
	if err != nil {
		return fmt.Errorf("screenInit: %v", err)
	}
	checkgl()

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func swapBuffers() {
	egl.SwapBuffers(screen.display, screen.surface)
}
