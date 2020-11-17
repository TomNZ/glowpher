package config

// API defines the overall Glowpher API.
type API struct {
	EffectOptions   []EffectOption   `json:"effectOptions"`
	VariableOptions []VariableOption `json:"variableOptions"`
	Config          Config           `json:"dsl"`
}

// Config defines the current Glowpher configuration.
type Config struct {
	Scenes    []Scene    `json:"scenes"`
	Playlists []Playlist `json:"playlists"`
	Variables []Variable `json:"variables"`
}
