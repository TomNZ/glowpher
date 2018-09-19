package variable

import (
	"errors"
	"math/rand"
)

const randomType = "random"

// Random is a variable that returns a value between 0 and 1.
type Random struct {
	variable
}

// Type returns the type of the variable.
func (r *Random) Type() string {
	return randomType
}

// New returns a new instance of the variable.
func (r *Random) New(name string, params map[string]Param) (Variable, error) {
	if r.concrete {
		return nil, errors.New("can't re-instantiate variables")
	}
	return &Random{
		variable{
			name:     name,
			concrete: true,
			params:   params,
		},
	}, nil
}

// Value returns the current value.
func (r *Random) Value() float32 {
	return rand.Float32()
}

func init() {
	Registry[randomType] = &Random{}
}
