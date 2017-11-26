package main

import (
	"fmt"
	"time"
	"ws2811"
)

func main() {
	fmt.Println("hello world!")
	ws2811.Init(18, 10, 255)
	ws2811.SetBitmap([]uint32{
		1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000,
	})
	ws2811.Render()
	time.Sleep(time.Second * 5)
	ws2811.Fini()
}
