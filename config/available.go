package config

// AvailableEffect defines an effect that is available to instantiate.
type AvailableEffect struct {
	Type       string `json:"type"`
	ParamTypes map[string]ParamType
}

// AvailableVariable defines a variable that is available to instantiate.
type AvailableVariable struct {
	Type       string `json:"type"`
	ParamTypes map[string]ParamType
}
