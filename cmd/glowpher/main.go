package main

import (
	"fmt"
	"math/rand"

	"github.com/tomnz/glowpher/internal/devices"
)

func main() {
	fmt.Println("hello world!")
	device := devices.Registry["ws281x"]
	device.Setup(10)
	for {
		device.ShowColors([]uint32{
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
			rand.Uint32(),
		})
	}
}
