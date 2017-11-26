// +build rpiws281x

package rpiws281x

/*
#cgo CFLAGS: -std=c99
#cgo LDFLAGS: -lws2811
#include "rpiws281x.go.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type WS281x struct {
	strip *C.ws2811_t
}

func (w *WS281x) Setup(numLights int) {
	w.strip = 
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

func (w *WS281x) ShowColors(colors []uint32) {
	// Two-phase - set the colors, then render them
	C.ws2811_set_bitmap(&C.ledstring, unsafe.Pointer(&colors[0]), C.int(len(colors)*4))

	res := int(C.ws2811_render(&C.ledstring))
	if res == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error ws2811.render.%d", res))
	}
}

func (w *WS281x) Clear() {
	C.ws2811_clear(&C.ledstring)
}

func (w *WS281x) Teardown() {
	C.ws2811_fini(&C.ledstring)
}
