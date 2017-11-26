package main

import (
	"fmt"
	"time"

	"github.com/tomnz/glowpher/internal/devices"
)

func main() {
	fmt.Println("hello world!")
	device := devices.Registry["ws281x"]
	device.Setup(10)
	device.ShowColors([]uint32{
		1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000,
	})
	time.Sleep(time.Second * 5)
	device.Clear()
	time.Sleep(time.Second * 1)
	device.Teardown()
}
