package main

import (
	"fmt"
	"time"

	"github.com/cozely/raspberry/gpio"
)

func main() {
	// fmt.Printf("%X\n", gpio.MultiGet())
	fmt.Println(gpio.GetMode(22))
	gpio.SetMode(22, gpio.Output)
	fmt.Println(gpio.GetMode(22))
	fmt.Println(gpio.Get(22))
	gpio.Set(22)
	fmt.Println(gpio.Get(22))
	<-time.After(2*time.Second)
	gpio.Clear(22)
	fmt.Println(gpio.Get(22))
}
