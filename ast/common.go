package ast

// Param declares the value or variable for a given parameter.
type Param struct {
	Value    interface{}      `json:"value"`
	Variable string           `json:"variable"`
	Params   map[string]Param `json:"params"`
}
