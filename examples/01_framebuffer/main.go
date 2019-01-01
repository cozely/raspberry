package main

import (
	"time"

	"github.com/cozely/raspberry/framebuffer"
)

//------------------------------------------------------------------------------

func main() {
	err := framebuffer.Setup()
	if err != nil {
		panic(err)
	}
	defer framebuffer.Cleanup()
	<-time.After(5 * time.Second)
}
