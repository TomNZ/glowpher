package effect

import "github.com/tomnz/glowpher/internal/variable"

const brightnessType = "brightness"

type Brightness struct {
	effect
}

func (b *Brightness) Type() string {
	return brightnessType
}

func (b *Brightness) New(params map[string]variable.Param) (Effect, error) {
	b2 := Brightness(*b)
	b2.concrete = true
	b2.params = params
	return &b2, nil
}

func init() {
	Registry[brightnessType] = &Brightness{}
}
