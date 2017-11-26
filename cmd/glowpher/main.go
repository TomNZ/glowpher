package main

import (
	"fmt"
	"time"

	"github.com/tomnz/glowpher/internal/rpiws281x"
)

func main() {
	fmt.Println("hello world!")
	rpiws281x.Init(18, 10, 255)
	rpiws281x.SetBitmap([]uint32{
		1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000,
	})
	rpiws281x.Render()
	time.Sleep(time.Second * 5)
	rpiws281x.Fini()
}
