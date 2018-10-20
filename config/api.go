package config

// API defines the overall Glowpher API.
type API struct {
	EffectOptions   []EffectOption   `json:"effectOptions"`
	VariableOptions []VariableOption `json:"variableOptions"`
	Config          Config           `json:"config"`
}

// Config defines the current Glowpher configuration.
type Config struct {
	Playlist  Playlist   `json:"playlist"`
	Variables []Variable `json:"variables"`
}
