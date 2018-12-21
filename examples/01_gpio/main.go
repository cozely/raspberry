package main

import (
	"fmt"
	"time"

	"github.com/cozely/raspberry/gpio"
)

func main() {
	// fmt.Printf("%X\n", gpio.MultiGet())
	fmt.Println(gpio.Mode(22))
	gpio.Output(22)
	fmt.Println(gpio.Mode(22))
	fmt.Println(gpio.Get(22))
	gpio.Set(22)
	fmt.Println(gpio.Get(22))
	<-time.After(2*time.Second)
	gpio.Clear(22)
	fmt.Println(gpio.Get(22))

	gpio.PullUp(23)
	gpio.Input(23)
	for {
		fmt.Println(gpio.Get(23))
		<-time.After(time.Second / 4)
	}
}
