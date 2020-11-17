package devices

import "image/color"

type Device interface {
	Setup(int) error
	ShowColors([]color.Color) error
	Clear() error
	Teardown()
}

var Registry = map[string]Device{}
