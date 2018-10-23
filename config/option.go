package config

// EffectOption defines an effect that is available to instantiate.
type EffectOption struct {
	Type          string `json:"type"`
	DefaultParams map[string]Param
}

// VariableOption defines a variable that is available to instantiate.
type VariableOption struct {
	Type          string `json:"type"`
	DefaultParams map[string]Param
}
