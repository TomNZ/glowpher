// +build rpiws281x

package rpiws281x

/*
#cgo CFLAGS: -std=c99
#cgo LDFLAGS: -lws2811 -lm
#include "rpiws281x.go.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/tomnz/glowpher/internal/color"
)

type WS281x struct {
	strip *C.ws2811_t
}

func (w *WS281x) Setup(numLights int) error {
	C.ledstring.channel[0].gpionum = C.int(18)
	C.ledstring.channel[0].count = C.int(numLights)
	C.ledstring.channel[0].brightness = C.uint8_t(255)
	res := int(C.ws2811_init(&C.ledstring))
	if res == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error ws2811.init.%d", res))
	}
}

func (w *WS281x) ShowColors(colors []color.Color) error {
	colorInts := make([]uint32, len(colors))
	for idx := range colors {
		colorInts[idx] = colors[idx].Uint32()
	}
	// Two-phase - set the colors, then render them
	C.ws2811_set_bitmap(&C.ledstring, unsafe.Pointer(&colorInts[0]), C.int(len(colors)*4))
	return w.render()
}

func (w *WS281x) Clear() error {
	C.ws2811_clear(&C.ledstring)
	return w.render()
}

func (w *WS281x) render() error {
	res := int(C.ws2811_render(&C.ledstring))
	if res == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error ws2811.render.%d", res))
	}
	res = int(C.ws2811_wait(&C.ledstring))
	if res == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error ws2811.wait.%d", res))
	}
}

func (w *WS281x) Teardown() {
	C.ws2811_fini(&C.ledstring)
}
