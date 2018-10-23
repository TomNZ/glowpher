package variable

// Variable defines the common interface that all variables must provide.
type Variable interface {
	Concrete() bool
	Type() string
	Name() string
	Params() map[string]Param
	DefaultParams() map[string]Param

	New(name string, params map[string]Param) (Variable, error)
}

// NumberVariable defines specific behavior for floating point variables.
type NumberVariable interface {
	Variable
	Value() float32
}

type variable struct {
	name     string
	concrete bool
	params   map[string]Param
}

func (v *variable) Concrete() bool {
	return v.concrete
}

func (v *variable) Name() string {
	return v.name
}

func (v *variable) Params() map[string]Param {
	if !v.Concrete() {
		panic("cannot get params for uninitialized variable")
	}
	return v.params
}

// Registry is the globally available set of variable types.
var Registry = map[string]Variable{}
