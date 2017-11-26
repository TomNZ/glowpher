package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tomnz/glowpher/internal/devices"
)

func main() {
	fmt.Println("hello world!")
	device := devices.Registry["ws281x"]
	device.Setup(10)
	i := 0
	for {
		if i < 1000 {
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
			i++
		} else {
			time.Sleep(time.Second * 10)
			i = 0
		}
	}
}
