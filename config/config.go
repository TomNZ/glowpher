package config

// Config defines the overall Glowpher configuration.
type Config struct {
	AvailableEffects   []AvailableEffect   `json:"availableEffects"`
	AvailableVariables []AvailableVariable `json:"availableVariables"`
	Playlist           Playlist            `json:"playlist"`
	Variables          []Variable          `json:"variables"`
}
