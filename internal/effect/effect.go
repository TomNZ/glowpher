package effect

import (
	"time"

	"github.com/tomnz/glowpher/internal/color"
	"github.com/tomnz/glowpher/internal/variable"
)

// Effect defines the common interface that all effects must provide.
type Effect interface {
	Type() string
	Concrete() bool
	Params() map[string]variable.Param
	DefaultParams() map[string]variable.Param

	New(params map[string]variable.Param) (Effect, error)

	Step(duration time.Duration, colors []color.Color)
}

type effect struct {
	concrete bool
	params   map[string]variable.Param
}

func (e *effect) Concrete() bool {
	return e.concrete
}

func (e *effect) Params() map[string]variable.Param {
	if !e.Concrete() {
		panic("cannot get params for uninitialized effect")
	}
	return e.params
}

// Registry is the globally available set of effect types.
var Registry = map[string]Effect{}
