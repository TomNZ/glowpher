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
	for {
		time.Sleep(time.Second * 10)
	}
}
