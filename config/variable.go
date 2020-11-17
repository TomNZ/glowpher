package config

// Variable declares an available variable, along with its base parameters.
type Variable struct {
	Name   string           `json:"name"`
	Type   string           `json:"type"`
	Params map[string]Param `json:"params"`
}
