package effect

import "github.com/tomnz/glowpher/internal/variable"

// Effect defines the common interface that all effects must provide.
type Effect interface {
	Type() string
	Concrete() bool
	Params() map[string]variable.Param
	New(params map[string]variable.Param) (Effect, error)
}

type effect struct {
	concrete bool
	params   map[string]variable.Param
}

func (e *effect) Concrete() bool {
	return e.concrete
}

func (e *effect) Params() map[string]variable.Param {
	return e.params
}

// Registry is the globally available set of effect types.
var Registry = map[string]Effect{}
