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

func Init(gpioPin int, ledCount int, brightness int) error {
	C.ledstring.channel[0].gpionum = C.int(gpioPin)
	C.ledstring.channel[0].count = C.int(ledCount)
	C.ledstring.channel[0].brightness = C.uint8_t(brightness)
	res := int(C.ws2811_init(&C.ledstring))
	if res == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error ws2811.init.%d", res))
	}
}

func Fini() {
	C.ws2811_fini(&C.ledstring)
}

func Render() error {
	res := int(C.ws2811_render(&C.ledstring))
	if res == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error ws2811.render.%d", res))
	}
}

func Wait() error {
	res := int(C.ws2811_wait(&C.ledstring))
	if res == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Error ws2811.wait.%d", res))
	}
}

func SetLed(index int, value uint32) {
	C.ws2811_set_led(&C.ledstring, C.int(index), C.uint32_t(value))
}

func Clear() {
	C.ws2811_clear(&C.ledstring)
}

func SetBitmap(a []uint32) {
	C.ws2811_set_bitmap(&C.ledstring, unsafe.Pointer(&a[0]), C.int(len(a)*4))
}
