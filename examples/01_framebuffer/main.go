package main

import (
	"time"

	"github.com/cozely/journal"

	"github.com/cozely/raspberry/framebuffer"
)

//------------------------------------------------------------------------------

func main() {
	err := framebuffer.Setup()
	if err != nil {
		journal.Panic(err)
	}
	defer framebuffer.Cleanup()
	<-time.After(1 * time.Second)
	framebuffer.Swap()
	<-time.After(1 * time.Second)
	framebuffer.Swap()
	<-time.After(1 * time.Second)
	framebuffer.Swap()
	<-time.After(1 * time.Second)
}
