package config

// ParamType indicates the type of a parameter.
type ParamType string

const (
	// ParamStringType indicates a string parameter.
	ParamStringType ParamType = "string"
	// ParamIntType indicates an integer parameter.
	ParamIntType ParamType = "int"
	// ParamFloatType indicates a float parameter.
	ParamFloatType ParamType = "float"
	// ParamColorType indicates a color parameter.
	ParamColorType ParamType = "color"
)

// Param declares the value or variable for a given parameter.
type Param struct {
	Type           ParamType       `json:"type"`
	Value          *ParamValue     `json:"value,omitempty"`
	Variable       *string         `json:"variable,omitempty"`
	VariableParams *VariableParams `json:"variableParams,omitempty"`
}

// VariableParams declares configuration options for a variable parameter.
type VariableParams struct {
	Multiply float32 `json:"multiply"`
	Add      float32 `json:"add"`
}

// Color describes a color.
type Color struct {
	R float32 `json:"r"`
	G float32 `json:"g"`
	B float32 `json:"b"`
}

// ParamValue stores the concrete value for a parameter. Only one of the values
// should be set, corresponding to the type indicated in the parameter.
type ParamValue struct {
	String *string  `json:"string,omitempty"`
	Int    *int     `json:"int,omitempty"`
	Float  *float32 `json:"float,omitempty"`
	Color  *Color   `json:"color,omitempty"`
}
