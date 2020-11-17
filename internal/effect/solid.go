package effect

import (
	"time"

	"github.com/tomnz/glowpher/internal/color"
	"github.com/tomnz/glowpher/internal/variable"
)

const solidType = "solid"

type Solid struct {
	effect
}

func (s *Solid) Type() string {
	return solidType
}

func (s *Solid) DefaultParams() map[string]variable.Param {
	return map[string]variable.Param{}
}

func (s *Solid) New(params map[string]variable.Param) (Effect, error) {
	s2 := *s
	s2.concrete = true
	s2.params = params
	return &s2, nil
}

func (s *Solid) Step(duration time.Duration, colors []color.Color) {
	color := s.color()
	for idx := range colors {
		colors[idx] = color
	}
}

func (s *Solid) color() color.Color {
	colorParam := s.params["color"].(variable.ParamColor)
	return colorParam.Value()
}

func init() {
	Registry[solidType] = &Solid{}
}
