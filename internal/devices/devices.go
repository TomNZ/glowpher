package devices

import "github.com/tomnz/glowpher/internal/color"

type Device interface {
	Setup(int) error
	ShowColors([]color.Color) error
	Clear() error
	Teardown()
}

var Registry = map[string]Device{}
