package effect

import (
	"time"

	"github.com/tomnz/glowpher/internal/color"
	"github.com/tomnz/glowpher/internal/variable"
)

const brightnessType = "brightness"

type Brightness struct {
	effect
}

func (b *Brightness) Type() string {
	return brightnessType
}

func (b *Brightness) DefaultParams() map[string]variable.Param {
	return map[string]variable.Param{}
}

func (b *Brightness) New(params map[string]variable.Param) (Effect, error) {
	b2 := *b
	b2.concrete = true
	b2.params = params
	return &b2, nil
}

func (b *Brightness) Step(duration time.Duration, colors []color.Color) {
}

func init() {
	Registry[brightnessType] = &Brightness{}
}
