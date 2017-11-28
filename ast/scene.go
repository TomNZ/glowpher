package ast

// Scene declares an overall animated configuration.
type Scene struct {
	Name   string
	Effect []Effect `json:"effects"`
}

// Effect declares the type and parameters of a given effect.
type Effect struct {
	Type   string           `json:"effect"`
	Params map[string]Param `json:"params"`
}
