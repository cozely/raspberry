package main

import (
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Lmicroseconds)

	err := initScreen()
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Screen size: %d x %d\n", screen.width, screen.height)

	err = createPipeline()
	if err != nil {
		log.Printf("createPipeline: %v", err)
		return
	}

	err = drawTriangle()
	if err != nil {
		log.Printf("drawTriangle: %v", err)
		return
	}

	time.Sleep(2 * time.Second)
}
